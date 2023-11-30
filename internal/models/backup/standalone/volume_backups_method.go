/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod The possible methods used to perform a volume backup.
//
//   - METHOD_UNSPECIFIED: Unspecified method.
//   - RESTIC: A file system backup using velero's restic integration.
//   - CSI_VOLUME_SNAPSHOT: A volume snapshot performed by the CSI driver.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.VolumeBackup.VolumeBackupMethod
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod string

func NewVmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod(value VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod) *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod.
func (m VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod) Pointer() *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethodMETHODUNSPECIFIED captures enum value "METHOD_UNSPECIFIED"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethodMETHODUNSPECIFIED VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod = "METHOD_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethodRESTIC captures enum value "RESTIC"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethodRESTIC VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod = "RESTIC"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethodCSIVOLUMESNAPSHOT captures enum value "CSI_VOLUME_SNAPSHOT"
	VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethodCSIVOLUMESNAPSHOT VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod = "CSI_VOLUME_SNAPSHOT"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterDataprotectionBackupVolumeBackupVolumeBackupMethodEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackupVolumeBackupMethod

	if err := json.Unmarshal([]byte(`["METHOD_UNSPECIFIED","RESTIC","CSI_VOLUME_SNAPSHOT"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterDataprotectionBackupVolumeBackupVolumeBackupMethodEnum = append(vmwareTanzuManageV1alpha1ClusterDataprotectionBackupVolumeBackupVolumeBackupMethodEnum, v)
	}
}
