/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupProgress Progress provides additional fields to track backup completion progress.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.Progress
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupProgress struct {

	// The number of items that have actually been written to the
	// backup tarball so far.
	ItemsBackedUp int32 `json:"itemsBackedUp,omitempty"`

	// The total number of items to be backed up. This number may change
	// throughout the execution of the backup due to plugins that return additional related
	// items to back up, the velero.io/exclude-from-backup label, and various other
	// filters that happen as items are processed.
	TotalItems int32 `json:"totalItems,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupProgress) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionBackupProgress

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
