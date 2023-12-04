/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package pbackup

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
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
	ScopeKey       = "scope"
	NameKey        = "name"
	BackupScopeKey = "backup_scope"
	SpecKey        = "spec"

	// Spec/Template Directive Keys.
	ClusterNameKey           = "cluster_name"
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

func ValidateSchema(backupSpecData map[string]interface{}, scope BackupScope) (err error) {
	errMsgs := make([]string, 0)
	includedNamespaces := backupSpecData[IncludedNamespacesKey].([]interface{})
	excludedNamespaces := backupSpecData[ExcludedNamespacesKey].([]interface{})
	orLabelSelector := backupSpecData[OrLabelSelectorKey].([]interface{})
	labelSelector := backupSpecData[LabelSelectorKey].([]interface{})

	switch scope {
	case FullClusterBackupScope:
		if len(includedNamespaces) > 0 {
			errMsgs = append(errMsgs, fmt.Sprintf("Included namespaces can't be configured when scope is %s", scope))
		}

		if len(labelSelector) > 0 {
			errMsgs = append(errMsgs, fmt.Sprintf("Lablel selectors can't be configured when scope is %s", scope))
		}

		if len(orLabelSelector) > 0 {
			errMsgs = append(errMsgs, fmt.Sprintf("Or lables selectors can't be configured when scope is %s", scope))
		}
	case NamespacesBackupScope:
		if len(includedNamespaces) == 0 {
			errMsgs = append(errMsgs, fmt.Sprintf("Included namespaces must be configured when scope is %s", scope))
		}

		if len(excludedNamespaces) > 0 {
			errMsgs = append(errMsgs, fmt.Sprintf("Excluded namespaces can't be configured when scope is %s", scope))
		}

		if len(labelSelector) > 0 {
			errMsgs = append(errMsgs, fmt.Sprintf("Lable selectors can't be configured when scope is %s", scope))
		}

		if len(orLabelSelector) > 0 {
			errMsgs = append(errMsgs, fmt.Sprintf("Or lables selectors can't be configured when scope is %s", scope))
		}

	case LabelSelectorBackupScope:
		if len(labelSelector) == 0 && len(orLabelSelector) == 0 {
			errMsgs = append(errMsgs, fmt.Sprintf("Or/Lablel selectors must be configured when scope is %s", scope))
		}

		if len(includedNamespaces) > 0 {
			errMsgs = append(errMsgs, fmt.Sprintf("Included namespaces can't be configured when scope is %s", scope))
		}
	}

	if len(errMsgs) > 0 {
		err = errors.New(fmt.Sprintf("Schema validation failed:\n%s", strings.Join(errMsgs, "\n")))
	}

	return err
}
