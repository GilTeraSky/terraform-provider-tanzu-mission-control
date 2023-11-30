/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase The phase for the current backup.
//
//   - PHASE_UNSPECIFIED: Phase_unspecified is the default phase.
//   - PENDING: Pending phase is set when the backup object is being processed by the service (TMC).
//   - CREATING: Creating phase is set when backup is being created on the cluster.
//   - NEW: The phase when backup has been created but not yet processed by velero.
//   - FAILEDVALIDATION: The phase when backup has failed the velero controller's validations and therefore will not run.
//   - INPROGRESS: The phase when backup is currently executing by velero.
//   - COMPLETED: The phase when backup has run successfully without errors.
//   - FAILED: The phase when backup ran but encountered an error that prevented it from completing successfully.
//   - PENDING_DELETE: Pending delete is set when the object deletion is being processed by the service.
//   - DELETING: The phase when backup and all its associated data are being deleted.
//   - PARTIALLY_FAILED: The phase when the backup has run to completion but encountered errors backing up individual items.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.Status.Phase
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase string

func NewVmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase(value VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase) *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase.
func (m VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase) Pointer() *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePENDING captures enum value "PENDING"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePENDING VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase = "PENDING"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseCREATING captures enum value "CREATING"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseCREATING VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase = "CREATING"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseNEW captures enum value "NEW"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseNEW VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase = "NEW"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseFAILEDVALIDATION captures enum value "FAILEDVALIDATION"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseFAILEDVALIDATION VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase = "FAILEDVALIDATION"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseINPROGRESS captures enum value "INPROGRESS"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseINPROGRESS VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase = "INPROGRESS"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseCOMPLETED captures enum value "COMPLETED"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseCOMPLETED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase = "COMPLETED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseFAILED captures enum value "FAILED"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseFAILED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase = "FAILED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePENDINGDELETE captures enum value "PENDING_DELETE"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePENDINGDELETE VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase = "PENDING_DELETE"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseDELETING captures enum value "DELETING"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhaseDELETING VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase = "DELETING"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePARTIALLYFAILED captures enum value "PARTIALLY_FAILED"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhasePARTIALLYFAILED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase = "PARTIALLY_FAILED"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterDataprotectionBackupStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase

	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","CREATING","NEW","FAILEDVALIDATION","INPROGRESS","COMPLETED","FAILED","PENDING_DELETE","DELETING","PARTIALLY_FAILED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterDataprotectionBackupStatusPhaseEnum = append(vmwareTanzuManageV1alpha1ClusterDataprotectionBackupStatusPhaseEnum, v)
	}
}
