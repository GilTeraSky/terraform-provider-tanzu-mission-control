/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupcommon

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode Specifies how Velero should behave if it encounters an error executing this hook.
//
//   - MODE_UNSPECIFIED: The default mode.
//   - CONTINUE: Means that an error from a hook is acceptable, and the operation can proceed.
//   - FAIL: Means that an error from a hook is problematic, and the operation should be in error.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.HookErrorMode.
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode string

func NewVmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode(value VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode) *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode.
func (m VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode) Pointer() *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorModeMODEUNSPECIFIED captures enum value "MODE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorModeMODEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode = "MODE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorModeCONTINUE captures enum value "CONTINUE".
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorModeCONTINUE VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode = "CONTINUE"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorModeFAIL captures enum value "FAIL".
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorModeFAIL VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode = "FAIL"
)

// for schema.
var VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorModeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorMode

	if err := json.Unmarshal([]byte(`["MODE_UNSPECIFIED","CONTINUE","FAIL"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorModeEnum = append(VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorModeEnum, v)
	}
}
