/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfo VolumeBackupPodInfo contains additional metadata about the pod where a backed up volume was mounted.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.VolumeBackupPodInfo
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfo struct {

	// The phase of the pod volume backup.
	BackupPhase *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfoPodVolumeBackupPhase `json:"backupPhase,omitempty"`

	// The name of the pod where the volume was mounted.
	PodName string `json:"podName,omitempty"`

	// The namespace of the pod where the volume was mounted.
	PodNamespace string `json:"podNamespace,omitempty"`

	// The name of the volume as depicted in the pod manifest.
	PodVolumeName string `json:"podVolumeName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfo) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfo

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
