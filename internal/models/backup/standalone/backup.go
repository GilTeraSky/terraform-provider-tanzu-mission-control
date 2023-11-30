/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"github.com/go-openapi/swag"

	backupcommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/common"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionBackupBackup The Kubernetes cluster state related to backup (API objects and associated volume state).
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.Backup
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackup struct {

	// Full name for the Backup Config.
	FullName *backupcommon.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName `json:"fullName,omitempty"`

	// Metadata for the backup object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the backup.
	Spec *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupSpec `json:"spec,omitempty"`

	// Status of the backup.
	Status *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackup) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionBackup

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
