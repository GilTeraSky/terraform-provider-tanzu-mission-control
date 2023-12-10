/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backup

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/proxy"
	commonbackupmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/common"
	testhelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/testing"
)

var (
	context = authctx.TanzuContext{
		ServerEndpoint:   os.Getenv(authctx.ServerEndpointEnvVar),
		Token:            os.Getenv(authctx.VMWCloudAPITokenEnvVar),
		VMWCloudEndPoint: os.Getenv(authctx.VMWCloudEndpointEnvVar),
		TLSConfig:        &proxy.TLSConfig{},
	}
)

func TestAcceptanceBackupResource(t *testing.T) {
	err := context.Setup()

	if err != nil {
		t.Error(errors.Wrap(err, "unable to set the context"))
		t.FailNow()
	}

	backupTargetLocationName, backupTargetLocationNameExist := os.LookupEnv(backupTargetLocationNameEnv)

	if !backupTargetLocationNameExist {
		t.Error("Backup target location name is missing!")
		t.FailNow()
	}

	var (
		tfResourceConfigBuilder = InitResourceTFConfigBuilder(testScopeHelper)
		provider                = initTestProvider(t)
	)

	resource.Test(t, resource.TestCase{
		PreCheck: testhelper.TestPreCheck(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		ProviderFactories: testhelper.GetTestProviderFactories(provider),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: tfResourceConfigBuilder.GetFullClusterBackupConfig(backupTargetLocationName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(FullClusterBackupResourceFullName, "name", FullClusterBackupName),
					verifyBackupResourceCreation(provider, FullClusterBackupResourceFullName, FullClusterBackupName),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetNamespacesBackupConfig(backupTargetLocationName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(NamespacesBackupResourceFullName, "name", NamespacesBackupName),
					verifyBackupResourceCreation(provider, NamespacesBackupResourceFullName, NamespacesBackupName),
				),
			},
			{
				Config: tfResourceConfigBuilder.GetLabelsBackupConfig(backupTargetLocationName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(LabelsBackupResourceFullName, "name", LabelsBackupName),
					verifyBackupResourceCreation(provider, LabelsBackupResourceFullName, LabelsBackupName),
				),
			},
		},
	},
	)

	t.Log("backup resource acceptance test complete!")
}

func verifyBackupResourceCreation(
	provider *schema.Provider,
	resourceName string,
	backupName string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if provider == nil {
			return fmt.Errorf("provider not initialised")
		}

		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("could not found resource %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID not set, resource %s", resourceName)
		}

		fn := &commonbackupmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName{
			Name:                  backupName,
			ManagementClusterName: testScopeHelper.Cluster.ManagementClusterName,
			ClusterName:           testScopeHelper.Cluster.Name,
			ProvisionerName:       testScopeHelper.Cluster.ProvisionerName,
		}

		resp, err := context.TMCConnection.BackupResourceService.BackupResourceServiceGet(fn)

		if err != nil {
			return fmt.Errorf("backup resource not found, resource: %s | err: %s", resourceName, err)
		}

		if resp == nil {
			return fmt.Errorf("backup resource is empty, resource: %s", resourceName)
		}

		return nil
	}
}
