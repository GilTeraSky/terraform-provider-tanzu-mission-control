/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase The possible phases pod volume backup can be in.
//
//   - PHASE_UNSPECIFIED: Default phase.
//   - IN_PROGRESS: The pod volume backup is in progress.
//   - COMPLETED: The pod volume backup has completed successfully.
//   - FAILED: The pod volume backup has completed unsuccessfully.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.VolumeBackupPodInfo.PodVolumeBackupPhase
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase string

func NewVmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase(value VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase) *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase.
func (m VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase) Pointer() *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhaseINPROGRESS captures enum value "IN_PROGRESS".
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhaseINPROGRESS VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase = "IN_PROGRESS"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhaseCOMPLETED captures enum value "COMPLETED".
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhaseCOMPLETED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase = "COMPLETED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhaseFAILED captures enum value "FAILED".
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhaseFAILED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase = "FAILED"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterDataprotectionBackupVolumeBackupPodInfoPodVolumeBackupPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase

	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","IN_PROGRESS","COMPLETED","FAILED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterDataprotectionBackupVolumeBackupPodInfoPodVolumeBackupPhaseEnum = append(vmwareTanzuManageV1alpha1ClusterDataprotectionBackupVolumeBackupPodInfoPodVolumeBackupPhaseEnum, v)
	}
}
