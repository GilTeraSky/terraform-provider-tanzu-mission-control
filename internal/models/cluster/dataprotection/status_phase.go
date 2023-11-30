/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase Available phases for data protection object.
//
//   - PHASE_UNSPECIFIED: Phase_unspecified is the default phase.
//   - PENDING: Pending phase is set when the data protection object is being processed by the service (TMC).
//   - CREATING: Creating phase is set when data protection is being enabled on the cluster.
//   - PENDING_DELETE: Pending delete is set when the data protection delete is being processed by the service.
//   - DELETING: Deleting the set when the data protection delete is in progress on the the cluster.
//   - READY: Ready phase is set when the data protection is successfully enabled.
//   - ERROR: Error phase is set when there was a failure while creating/deleting data protection.
//   - UPDATING: Updating is set when the data protection is being updated.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.Status.Phase.
type VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase string

func NewVmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase(value VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase) *VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase.
func (m VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase) Pointer() *VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhasePENDING captures enum value "PENDING".
	VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhasePENDING VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase = "PENDING"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseCREATING captures enum value "CREATING".
	VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseCREATING VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase = "CREATING"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhasePENDINGDELETE captures enum value "PENDING_DELETE".
	VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhasePENDINGDELETE VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase = "PENDING_DELETE"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseDELETING captures enum value "DELETING".
	VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseDELETING VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase = "DELETING"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseREADY captures enum value "READY".
	VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseREADY VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase = "READY"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseERROR captures enum value "ERROR".
	VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseERROR VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase = "ERROR"

	// VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseUPDATING captures enum value "UPDATING".
	VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseUPDATING VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase = "UPDATING"
)

// for schema.
var VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhase

	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","PENDING","CREATING","PENDING_DELETE","DELETING","READY","ERROR","UPDATING"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseEnum = append(VmwareTanzuManageV1alpha1ClusterDataProtectionStatusPhaseEnum, v)
	}
}
