/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterbackupschedule

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase The lifecycle phase of a schedule backup.
//
//   - PHASE_UNSPECIFIED: Phase_unspecified is the default phase.
//   - PENDING: Pending phase is set when the schedule object is being processed by the service (TMC).
//   - CREATING: Creating phase is set when schedule is being created on the cluster.
//   - NEW: The schedule has been created but not yet processed by velero.
//   - ENABLED: The schedule has been validated and will now be triggering backups according to the schedule spec.
//   - FAILEDVALIDATION: The schedule has failed the velero controller's validations and therefore will not trigger backups.
//   - PENDING_DELETE: Pending delete is set when the object deletion is being processed by the service.
//   - DELETING: The phase when schedule is being deleted.
//   - UPDATING: The phase when schedule is being updated.
//   - PAUSED: The phase when schedule is being paused.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.schedule.Status.Phase.
type VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase string

func NewVmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase(value VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase) *VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase.
func (m VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase) Pointer() *VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhasePENDING VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase = "PENDING"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseCREATING VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase = "CREATING"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseNEW captures enum value "NEW".
	VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseNEW VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase = "NEW"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseENABLED captures enum value "ENABLED".
	VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseENABLED VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase = "ENABLED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseFAILEDVALIDATION captures enum value "FAILEDVALIDATION".
	VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseFAILEDVALIDATION VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase = "FAILEDVALIDATION"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhasePENDINGDELETE captures enum value "PENDING_DELETE".
	VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhasePENDINGDELETE VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase = "PENDING_DELETE"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseDELETING VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase = "DELETING"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseUPDATING captures enum value "UPDATING".
	VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseUPDATING VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase = "UPDATING"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhasePAUSED captures enum value "PAUSED".
	VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhasePAUSED VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase = "PAUSED"
)

// for schema.
var VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhase

	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","CREATING","NEW","ENABLED","FAILEDVALIDATION","PENDING_DELETE","DELETING","UPDATING","PAUSED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseEnum = append(VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleStatusPhaseEnum, v)
	}
}
