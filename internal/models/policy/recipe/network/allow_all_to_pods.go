/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
Code generated by go-swagger; DO NOT EDIT.
*/

package policyrecipenetworkmodel

import (
	"github.com/go-openapi/swag"

	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

// V1alpha1CommonPolicySpecNetworkV1AllowAllToPods allow-all-to-pods schema
//
// swagger:model V1alpha1CommonPolicySpecNetworkV1AllowAllToPods
type V1alpha1CommonPolicySpecNetworkV1AllowAllToPods struct {

	// Allow traffic only from own namespace
	// Allow traffic only from pods in the same namespace as the destination pod. Defaults to false (allow from all namespaces).
	FromOwnNamespace *bool `json:"fromOwnNamespace,omitempty"`

	// Pod Labels on which traffic should be allowed
	// Use a label selector to identify the pods to which the policy applies
	// Required: true
	ToPodLabels []*policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels `json:"toPodLabels"`
}

// MarshalBinary interface implementation
func (m *V1alpha1CommonPolicySpecNetworkV1AllowAllToPods) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1CommonPolicySpecNetworkV1AllowAllToPods) UnmarshalBinary(b []byte) error {
	var res V1alpha1CommonPolicySpecNetworkV1AllowAllToPods
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}