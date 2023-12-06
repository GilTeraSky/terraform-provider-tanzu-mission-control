/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backup

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	commonbackupmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/common"
	backupsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/standalone"
	backupcommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backup"
)

func ResourceBackup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBackupCreate,
		ReadContext:   resourceBackupRead,
		UpdateContext: resourceBackupUpdate,
		DeleteContext: resourceBackupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceBackupImporter,
		},
		CustomizeDiff: validateSchema,
		Schema:        backupResourceSchema,
	}
}

func resourceBackupCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create Tanzu Mission Control backup."))
	}

	request := &backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData{
		Backup: model,
	}

	_, err = config.TMCConnection.BackupsResourceService.BackupResourceServiceCreate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create Tanzu Mission Control backup.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s,  Name: %s",
			model.FullName.ManagementClusterName, model.FullName.ProvisionerName, model.FullName.ClusterName, model.FullName.Name))
	}

	return resourceBackupRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func resourceBackupRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var resp *backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData

	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{backupcommon.ClusterNameKey, backupcommon.ManagementClusterNameKey, backupcommon.ProvisionerNameKey, backupcommon.NameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read Tanzu Mission Control backup."))
	}

	backupFn := model.FullName
	resp, err = readResourceWait(ctx, &config, backupFn)

	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			if !helper.IsContextCallerSet(ctx) {
				*data = schema.ResourceData{}

				return diags
			} else if helper.IsDeleteState(ctx) {
				// d.SetId("") is automatically called assuming delete returns no errors, but
				// it is added here for explicitness.
				_ = schema.RemoveFromState(data, m)

				return diags
			}
		}

		return diag.FromErr(errors.Wrapf(err, "Couldn't read backup.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s,  Name: %s",
			backupFn.ManagementClusterName, backupFn.ProvisionerName, backupFn.ClusterName, backupFn.Name))
	} else if resp != nil {
		fullNameList := []string{backupFn.ManagementClusterName, backupFn.ProvisionerName, backupFn.ClusterName, backupFn.Name}

		data.SetId(strings.Join(fullNameList, "/"))

		if *resp.Backup.Status.Phase == backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseFAILEDVALIDATION || *resp.Backup.Status.Phase == backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseFAILED {
			diags = diag.Errorf("Backup has failed.")
		} else if *resp.Backup.Status.Phase == backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePARTIALLYFAILED {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Detail:   "Backup has partially failed.",
			})
		}
	}

	return diags
}

func resourceBackupDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{backupcommon.ClusterNameKey, backupcommon.ManagementClusterNameKey, backupcommon.ProvisionerNameKey, backupcommon.NameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete Tanzu Mission Control backup."))
	}

	backupFn := model.FullName
	err = config.TMCConnection.BackupsResourceService.BackupResourceServiceDelete(backupFn)

	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete Tanzu Mission Control backup.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s,  Name: %s",
			backupFn.ManagementClusterName, backupFn.ProvisionerName, backupFn.ClusterName, backupFn.Name))
	}

	return resourceBackupRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func resourceBackupUpdate(_ context.Context, _ *schema.ResourceData, _ interface{}) (diags diag.Diagnostics) {
	return diag.FromErr(errors.New("update of Tanzu Mission Control backup is not supported"))
}

func resourceBackupImporter(ctx context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	config := m.(authctx.TanzuContext)
	backupID := data.Id()

	if backupID == "" {
		return nil, errors.New("ID is needed to import a backup.")
	}

	namesArray := strings.Split(backupID, "/")

	if len(namesArray) != 4 {
		return nil, errors.Errorf("Invalid backup ID.\nBackup id should consists of a full cluster name and the backup name separated by '/'.\nProvided ID: %s", backupID)
	}

	backupFn := &commonbackupmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName{
		ManagementClusterName: namesArray[0],
		ProvisionerName:       namesArray[1],
		ClusterName:           namesArray[2],
		Name:                  namesArray[3],
	}
	resp, err := readResourceWait(ctx, &config, backupFn)

	if err != nil || resp.Backup == nil {
		return nil, errors.Errorf("Couldn't import backup.\nManagement Cluster Name: %s, Provisioner Name: %s, Cluster Name: %s,  Name: %s",
			backupFn.ManagementClusterName, backupFn.ProvisionerName, backupFn.ClusterName, backupFn.Name)
	} else {
		err = tfModelResourceConverter.FillTFSchema(resp.Backup, data)

		if err != nil {
			return nil, err
		}
	}

	return []*schema.ResourceData{data}, err
}

func validateSchema(ctx context.Context, diff *schema.ResourceDiff, m interface{}) error {
	backupSpec := diff.Get(backupcommon.SpecKey).([]interface{})[0].(map[string]interface{})
	backupScope := backupcommon.BackupScope(diff.Get(backupcommon.BackupScopeKey).(string))

	return backupcommon.ValidateSchema(backupSpec, backupScope)
}

func readResourceWait(ctx context.Context, config *authctx.TanzuContext, resourceFullName *commonbackupmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName) (resp *backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData, err error) {
	stopStatuses := map[backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase]bool{
		backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseCOMPLETED:        true,
		backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseFAILEDVALIDATION: true,
		backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseFAILED:           true,
		backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePARTIALLYFAILED:  true,
	}

	responseStatus := backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePHASEUNSPECIFIED
	_, isStopStatus := stopStatuses[responseStatus]
	isCtxCallerSet := helper.IsContextCallerSet(ctx)

	for !isStopStatus {
		if isCtxCallerSet || (!isCtxCallerSet && responseStatus != backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePHASEUNSPECIFIED) {
			time.Sleep(5 * time.Second)
		}

		resp, err = config.TMCConnection.BackupsResourceService.BackupResourceServiceGet(resourceFullName)

		if err != nil || resp == nil || resp.Backup == nil {
			return nil, err
		}

		responseStatus = *resp.Backup.Status.Phase
		_, isStopStatus = stopStatuses[responseStatus]
	}

	return resp, err
}
