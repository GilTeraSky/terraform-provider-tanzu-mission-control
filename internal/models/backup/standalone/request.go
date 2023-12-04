/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupCreateBackupRequest Request to create a Backup.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.CreateBackupRequest
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData struct {

	// Backup to create.
	Backup *VmwareTanzuManageV1alpha1ClusterDataProtectionBackup `json:"backup,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterDataprotectionBackupListBackupsResponse Response from listing Backups.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.ListBackupsResponse
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupListResponse struct {

	// List of backups.
	Backups []*VmwareTanzuManageV1alpha1ClusterDataProtectionBackup `json:"backups"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupListResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionBackupListResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}