/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupResources Resources represents the list of names of k8s resources of one kind.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.Resources
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupResources struct {

	// Names will be the list of resource name under that GVK.
	Names []string `json:"names"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupResources) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionBackupResources

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
