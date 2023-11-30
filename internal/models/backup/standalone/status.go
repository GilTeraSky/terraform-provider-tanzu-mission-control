/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatus Status of the backup resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.backup.Status
type VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatus struct {

	// A list of available phases for backup object.
	AvailablePhases []*VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase `json:"availablePhases"`

	// The URL to download the backup logs.
	BackupLogsURL string `json:"backupLogsUrl,omitempty"`

	// The timestamp when a backup was completed.
	// Format: date-time
	CompletionTimestamp strfmt.DateTime `json:"completionTimestamp,omitempty"`

	// The conditions attached to this backup object.
	// The description of the conditions is as follows:
	// - "Scheduled" with status 'Unknown' indicates the backup request has not been applied to the cluster yet
	// - "Scheduled" with status 'False' indicates the request could not be forwarded to the cluster (e.g. intent generation failure)
	// - "Scheduled" with status 'True' and "Ready" with status 'Unknown' indicates the backup create / delete intent has been applied / deleted but not yet acted upon
	// - "Ready" with status 'True' indicates the the creation of data protection is complete
	// - "Ready" with status 'False' indicates the the creation of data protection is in error state.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// The total number of attempted CSI volume snapshots for this backup.
	CsiVolumeSnapshotsAttempted int32 `json:"csiVolumeSnapshotsAttempted,omitempty"`

	// The total number of successfully completed CSI volume snapshots for this backup.
	CsiVolumeSnapshotsCompleted int32 `json:"csiVolumeSnapshotsCompleted,omitempty"`

	// The expiration time associated with this Backup object if it is eligible for garbage-collection.
	// Format: date-time
	Expiration strfmt.DateTime `json:"expiration,omitempty"`

	// The error that caused the entire backup to fail.
	FailureReason string `json:"failureReason,omitempty"`

	// The backup's format version, including major, minor, and patch version.
	FormatVersion string `json:"formatVersion,omitempty"`

	// This holds the status of ListBackupResources action which is triggered on the cluster once the backup is completed.
	// This will specify the state of that process such that the user is indicated, if the action to get the resources is processed or failed.
	GatherBackupResourcesStatus *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupGatherBackupResourcesStatus `json:"gatherBackupResourcesStatus,omitempty"`

	// The resource generation the current status applies to.
	ObservedGeneration string `json:"observedGeneration,omitempty"`

	// The current state of the Backup.
	Phase *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatusPhase `json:"phase,omitempty"`

	// Additional info about the phase.
	PhaseInfo string `json:"phaseInfo,omitempty"`

	// The backup progress so far.
	Progress *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupProgress `json:"progress,omitempty"`

	// The URL to download the list of resources that were backed up.
	ResourceListURL string `json:"resourceListUrl,omitempty"`

	// Map of important resources kind with their names. Map<k8s_Kind, list of names>.
	// Currently the expected kinds in a backup are Namespace, PersistentVolumeClaim and PersistentVolume.
	// Other resources which are part of the backup will be saved in an appropriate format in TMC S3 bucket and will be accessible
	// to the user on demand.
	Resources map[string]VmwareTanzuManageV1alpha1ClusterDataProtectionBackupResources `json:"resources,omitempty"`

	// A list of all the volume backups attempted by restic.
	ResticBackupsAttempted []string `json:"resticBackupsAttempted"`

	// A list of all the volume backups completed by restic.
	ResticBackupsCompleted []string `json:"resticBackupsCompleted"`

	// The timestamp when a backup was started.
	// Format: date-time
	StartTimestamp strfmt.DateTime `json:"startTimestamp,omitempty"`

	// A list of all validation errors (if applicable).
	ValidationErrors []string `json:"validationErrors"`

	// Information about volumes backed up.
	VolumeBackups []*VmwareTanzuManageV1alpha1ClusterDataProtectionBackupVolumeBackup `json:"volumeBackups"`

	// The total number of attempted volume snapshots for this backup.
	VolumeSnapshotsAttempted int32 `json:"volumeSnapshotsAttempted,omitempty"`

	// The total number of successfully completed volume snapshots for this backup.
	VolumeSnapshotsCompleted int32 `json:"volumeSnapshotsCompleted,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataProtectionBackupStatus

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
