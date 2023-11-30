/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatus Defines the status of ListBackupResources action.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.GatherBackupResourcesStatus
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatus struct {

	// Conditions attached to sub-operations like uploads during the collection of backup resources.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Phase is the overall state of the ListBackupResources operation.
	Phase *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatusPhase `json:"phase,omitempty"`

	// PhaseInfo defines the message for the state which maybe non-empty for the failed state.
	PhaseInfo string `json:"phaseInfo,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatus

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
