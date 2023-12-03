/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backup

import (
	commonbackupmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/common"
)

type BackupScope string

const (
	FullClusterBackupScope   BackupScope = "FULL_CLUSTER"
	NamespacesBackupScope    BackupScope = "SET_NAMESPACES"
	LabelSelectorBackupScope BackupScope = "LABEL_SELECTOR"
)

const (

	// Root Keys.
	NameKey                  = "name"
	ClusterNameKey           = "cluster_name"
	SpecKey                  = "spec"
	ProvisionerNameKey       = "provisioner_name"
	ManagementClusterNameKey = "management_cluster_name"

	// Template Directive Keys.
	BackupTTLKey                = "backup_ttl"
	SystemExcludedNamespacesKey = "sys_excluded_namespaces"
	ExcludedNamespacesKey       = "excluded_namespaces"
	IncludedNamespacesKey       = "included_namespaces"
	ExcludedResourcesKey        = "excluded_resources"
	IncludedResourcesKey        = "included_resources"
	IncludeClusterResourcesKey  = "include_cluster_resources"
	DefaultVolumesToResticKey   = "default_volumes_to_restic"
	SnapshotVolumesKey          = "snapshot_volumes"
	CsiSnapshotTimeoutKey       = "csi_snapshot_timeout"
	DefaultVolumesToFsBackupKey = "default_volumes_to_fs_backup"
	StorageLocationKey          = "storage_location"
	VolumeSnapshotLocationsKey  = "volume_snapshot_locations"
	OrderedResourcesKey         = "ordered_resources"
	HooksKey                    = "hooks"
	LabelSelectorKey            = "label_selector"
	OrLabelSelectorKey          = "or_label_selector"

	// Hooks Directive Keys.
	ResourceKey = "resource"

	// Resource Directive Keys.
	PreHookKey  = "pre_hook"
	PostHookKey = "post_hook"

	// Pre/Post Hook Directive Keys.
	ExecKey = "exec"

	// Exec Directive Keys.
	CommandKey   = "command"
	ContainerKey = "container"
	OnErrorKey   = "on_error"
	TimeoutKey   = "timeout"

	// (Or)Label Selector Directive Keys.
	MatchLabelsKey     = "match_labels"
	MatchExpressionKey = "match_expression"

	// Match Expressions Directive Keys.
	MeKey         = "key"
	MeOperatorKey = "operator"
	MeValuesKey   = "values"
)

var (
	ScopeValidValues   = []string{string(FullClusterBackupScope), string(NamespacesBackupScope), string(LabelSelectorBackupScope)}
	OnErrorValidValues = []string{
		string(commonbackupmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorModeFAIL),
		string(commonbackupmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupHookErrorModeCONTINUE),
	}
)
