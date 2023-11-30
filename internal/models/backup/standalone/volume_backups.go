/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackup VolumeBackup contains metadata about a particular volume backup taken.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.VolumeBackup
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackup struct {

	// The method used to perform the volume backup.
	Method *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod `json:"method,omitempty"`

	// Additional metadata about the pod where the volume was mounted.
	// Only present for file system backups.
	PodInfo *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupPodInfo `json:"podInfo,omitempty"`

	// The name of the persistent volume.
	PvName string `json:"pvName,omitempty"`

	// The name of the persistent volume claim.
	PvcName string `json:"pvcName,omitempty"`

	// The namespace of the persistent volume claim.
	PvcNamespace string `json:"pvcNamespace,omitempty"`

	// The name of the storage class used by the persistent volume.
	ScName string `json:"scName,omitempty"`

	// The complete size of the snapshot in bytes.
	Size string `json:"size,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackup) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackup

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
