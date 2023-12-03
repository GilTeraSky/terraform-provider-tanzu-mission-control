/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterbackupschedule

import (
	"github.com/go-openapi/swag"

	backupcommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/common"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleSpec The schedule spec.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.schedule.Spec.
type VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleSpec struct {

	// Paused specifies whether the schedule is paused or not.
	Paused bool `json:"paused"`

	// Rate at which the backup is to be run.
	Schedule *VmwareTanzuManageV1alpha1CommonScheduleSchedule `json:"schedule,omitempty"`

	// The definition of the Backup to be run on the provided schedule.
	Template *backupcommon.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupSpec `json:"template,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
