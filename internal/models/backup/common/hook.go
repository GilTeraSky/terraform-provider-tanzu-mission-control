/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupcommon

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupBackupResourceHook BackupResourceHook defines a hook for a resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.BackupResourceHook.
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupBackupResourceHook struct {

	// Exec defines an exec hook.
	Exec *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupExecHook `json:"exec,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupBackupResourceHook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupBackupResourceHook) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionBackupBackupResourceHook

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
