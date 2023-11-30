/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessMode The permissions for a BackupStorageLocation.
//
//   - READONLY: The read only access.
//   - READWRITE: Read and write access.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backuplocation.Status.BackupStorageLocationAccessMode.
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessMode string

func NewVmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessMode(value VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessMode) *VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessMode.
func (m VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessMode) Pointer() *VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessMode {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessModeREADONLY captures enum value "READONLY".
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessModeREADONLY VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessMode = "READONLY"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessModeREADWRITE captures enum value "READWRITE".
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessModeREADWRITE VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessMode = "READWRITE"
)

// for schema.
var VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessModeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessMode

	if err := json.Unmarshal([]byte(`["READONLY","READWRITE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessModeEnum = append(VmwareTanzuManageV1alpha1ClusterDataProtectionBackuplocationStatusBackupStorageLocationAccessModeEnum, v)
	}
}
