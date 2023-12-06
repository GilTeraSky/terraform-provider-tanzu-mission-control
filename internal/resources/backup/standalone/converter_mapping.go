/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backup

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	backupsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/standalone"
	backupcommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var (
	resourcesArrayField        = tfModelConverterHelper.BuildArrayField("resources")
	postHooksArrayField        = tfModelConverterHelper.BuildArrayField("postHooks")
	preHooksArrayField         = tfModelConverterHelper.BuildArrayField("preHooks")
	matchExpressionsArrayField = tfModelConverterHelper.BuildArrayField("matchExpressions")
	orLabelSelectorsArrayField = tfModelConverterHelper.BuildArrayField("orLabelSelectors")
)

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	backupcommon.NameKey:                  tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	backupcommon.ClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterName"),
	backupcommon.ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
	backupcommon.ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("fullName", "provisionerName"),
	common.MetaKey:                        common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	backupcommon.SpecKey: &tfModelConverterHelper.BlockToStruct{
		backupcommon.CsiSnapshotTimeoutKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "csiSnapshotTimeout"),
		backupcommon.DefaultVolumesToFsBackupKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "defaultVolumesToFsBackup"),
		backupcommon.DefaultVolumesToResticKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "defaultVolumesToRestic"),
		backupcommon.ExcludedNamespacesKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "excludedNamespaces"),
		backupcommon.IncludedNamespacesKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "includedNamespaces"),
		backupcommon.ExcludedResourcesKey:        tfModelConverterHelper.BuildDefaultModelPath("spec", "excludedResources"),
		backupcommon.IncludedResourcesKey:        tfModelConverterHelper.BuildDefaultModelPath("spec", "includedResources"),
		backupcommon.IncludeClusterResourcesKey:  tfModelConverterHelper.BuildDefaultModelPath("spec", "includeClusterResources"),
		backupcommon.SnapshotVolumesKey:          tfModelConverterHelper.BuildDefaultModelPath("spec", "snapshotVolumes"),
		backupcommon.StorageLocationKey:          tfModelConverterHelper.BuildDefaultModelPath("spec", "storageLocation"),
		backupcommon.BackupTTLKey:                tfModelConverterHelper.BuildDefaultModelPath("spec", "ttl"),
		backupcommon.VolumeSnapshotLocationsKey:  tfModelConverterHelper.BuildDefaultModelPath("spec", "volumeSnapshotLocations"),
		backupcommon.HooksKey: &tfModelConverterHelper.BlockToStruct{
			backupcommon.ResourceKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					backupcommon.ExcludedNamespacesKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, "excludedNamespaces"),
					backupcommon.IncludedNamespacesKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, "includedNamespaces"),
					backupcommon.NameKey:               tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, "name"),
					backupcommon.PostHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
						{
							backupcommon.ExecKey: &tfModelConverterHelper.BlockToStruct{
								backupcommon.CommandKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, postHooksArrayField, "exec", "command"),
								backupcommon.ContainerKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, postHooksArrayField, "exec", "container"),
								backupcommon.OnErrorKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, postHooksArrayField, "exec", "onError"),
								backupcommon.TimeoutKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, postHooksArrayField, "exec", "timeout"),
							},
						},
					},
					backupcommon.PreHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
						{
							backupcommon.ExecKey: &tfModelConverterHelper.BlockToStruct{
								backupcommon.CommandKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, preHooksArrayField, "exec", "command"),
								backupcommon.ContainerKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, preHooksArrayField, "exec", "container"),
								backupcommon.OnErrorKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, preHooksArrayField, "exec", "onError"),
								backupcommon.TimeoutKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, preHooksArrayField, "exec", "timeout"),
							},
						},
					},
					backupcommon.LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
						backupcommon.MatchExpressionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
							{
								backupcommon.MeKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "key"),
								backupcommon.MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "operator"),
								backupcommon.MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "values"),
							},
						},
						backupcommon.MatchLabelsKey: &tfModelConverterHelper.Map{
							tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "hooks", resourcesArrayField, "labelSelector", "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
						},
					},
				},
			},
		},
		backupcommon.LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
			backupcommon.MatchExpressionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					backupcommon.MeKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "labelSelector", matchExpressionsArrayField, "key"),
					backupcommon.MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "labelSelector", matchExpressionsArrayField, "operator"),
					backupcommon.MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "labelSelector", matchExpressionsArrayField, "values"),
				},
			},
			backupcommon.MatchLabelsKey: &tfModelConverterHelper.Map{
				tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "labelSelector", "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
			},
		},
		backupcommon.OrLabelSelectorKey: &tfModelConverterHelper.BlockSliceToStructSlice{
			{
				backupcommon.MatchExpressionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
					{
						backupcommon.MeKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", orLabelSelectorsArrayField, matchExpressionsArrayField, "key"),
						backupcommon.MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", orLabelSelectorsArrayField, matchExpressionsArrayField, "operator"),
						backupcommon.MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", orLabelSelectorsArrayField, matchExpressionsArrayField, "values"),
					},
				},
				backupcommon.MatchLabelsKey: &tfModelConverterHelper.Map{
					tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", orLabelSelectorsArrayField, "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
				},
			},
		},
		backupcommon.OrderedResourcesKey: &tfModelConverterHelper.Map{
			tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "orderedResources", tfModelConverterHelper.AllMapKeysFieldMarker),
		},
	},
}
var tfModelResourceConverter = tfModelConverterHelper.TFSchemaModelConverter[*backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackup]{
	TFModelMap: tfModelResourceMap,
}
