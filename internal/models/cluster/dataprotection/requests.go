/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionCreateDataProtectionRequest Request to create a DataProtection.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.CreateDataProtectionRequest.
type VmwareTanzuManageV1alpha1ClusterDataProtectionCreateDataProtectionRequest struct {

	// DataProtection to create.
	DataProtection *VmwareTanzuManageV1alpha1ClusterDataProtectionDataProtection `json:"dataProtection,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionCreateDataProtectionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionCreateDataProtectionRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionCreateDataProtectionRequest

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterDataProtectionCreateDataProtectionResponse Response from creating a DataProtection.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.CreateDataProtectionResponse.
type VmwareTanzuManageV1alpha1ClusterDataProtectionCreateDataProtectionResponse struct {

	// DataProtection created.
	DataProtection *VmwareTanzuManageV1alpha1ClusterDataProtectionDataProtection `json:"dataProtection,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionCreateDataProtectionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionCreateDataProtectionResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionCreateDataProtectionResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterDataProtectionListDataProtectionsResponse Response from listing DataProtections.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.ListDataProtectionsResponse.
type VmwareTanzuManageV1alpha1ClusterDataProtectionListDataProtectionsResponse struct {

	// List of dataprotections.
	DataProtections []*VmwareTanzuManageV1alpha1ClusterDataProtectionDataProtection `json:"dataProtections"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionListDataProtectionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionListDataProtectionsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionListDataProtectionsResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
