/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase Phase of the ListBackupResources action in the associated cluster.
//
//   - PHASE_UNSPECIFIED: Default phase.
//   - PROCESSED: The internal backupResources state which specifies if the resources have been populated to the CR.
//   - FAILED: Failure state while processing the backup resources tar from velero or while uploading.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.GatherBackupResourcesStatus.Phase
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase string

func NewVmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase(value VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase) *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase.
func (m VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase) Pointer() *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhasePROCESSED captures enum value "PROCESSED"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhasePROCESSED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase = "PROCESSED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhaseFAILED captures enum value "FAILED"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhaseFAILED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase = "FAILED"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterDataprotectionBackupGatherBackupResourcesStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase

	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PROCESSED","FAILED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterDataprotectionBackupGatherBackupResourcesStatusPhaseEnum = append(vmwareTanzuManageV1alpha1ClusterDataprotectionBackupGatherBackupResourcesStatusPhaseEnum, v)
	}
}
