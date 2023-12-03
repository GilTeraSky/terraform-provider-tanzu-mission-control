/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	clusterbackupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/schedule/cluster"
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
	backupcommon.NameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	ScopeKey: &tfModelConverterHelper.BlockToStruct{
		ClusterScopeKey: &tfModelConverterHelper.BlockToStruct{
			backupcommon.ClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterName"),
			backupcommon.ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
			backupcommon.ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("fullName", "provisionerName"),
		},
	},
	common.MetaKey: common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	backupcommon.SpecKey: &tfModelConverterHelper.BlockToStruct{
		PausedKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "paused"),
		ScheduleKey: &tfModelConverterHelper.BlockToStruct{
			RateKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "schedule", "rate"),
		},
		TemplateKey: &tfModelConverterHelper.BlockToStruct{
			backupcommon.CsiSnapshotTimeoutKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "csiSnapshotTimeout"),
			backupcommon.DefaultVolumesToFsBackupKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "defaultVolumesToFsBackup"),
			backupcommon.DefaultVolumesToResticKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "defaultVolumesToRestic"),
			backupcommon.ExcludedNamespacesKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "excludedNamespaces"),
			backupcommon.IncludedNamespacesKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "includedNamespaces"),
			backupcommon.ExcludedResourcesKey:        tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "excludedResources"),
			backupcommon.IncludedResourcesKey:        tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "includedResources"),
			backupcommon.IncludeClusterResourcesKey:  tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "includeClusterResources"),
			backupcommon.SnapshotVolumesKey:          tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "snapshotVolumes"),
			backupcommon.StorageLocationKey:          tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "storageLocation"),
			backupcommon.BackupTTLKey:                tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "ttl"),
			backupcommon.VolumeSnapshotLocationsKey:  tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "volumeSnapshotLocations"),
			backupcommon.HooksKey: &tfModelConverterHelper.BlockToStruct{
				backupcommon.ResourceKey: &tfModelConverterHelper.BlockSliceToStructSlice{
					{
						backupcommon.ExcludedNamespacesKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "excludedNamespaces"),
						backupcommon.IncludedNamespacesKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "includedNamespaces"),
						backupcommon.NameKey:               tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "name"),
						backupcommon.PostHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
							{
								backupcommon.ExecKey: &tfModelConverterHelper.BlockToStruct{
									backupcommon.CommandKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "command"),
									backupcommon.ContainerKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "container"),
									backupcommon.OnErrorKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "onError"),
									backupcommon.TimeoutKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "timeout"),
								},
							},
						},
						backupcommon.PreHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
							{
								backupcommon.ExecKey: &tfModelConverterHelper.BlockToStruct{
									backupcommon.CommandKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "command"),
									backupcommon.ContainerKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "container"),
									backupcommon.OnErrorKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "onError"),
									backupcommon.TimeoutKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "timeout"),
								},
							},
						},
						backupcommon.LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
							backupcommon.MatchExpressionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
								{
									backupcommon.MeKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "key"),
									backupcommon.MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "operator"),
									backupcommon.MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "values"),
								},
							},
							backupcommon.MatchLabelsKey: &tfModelConverterHelper.Map{
								tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "labelSelector", "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
							},
						},
					},
				},
			},
			backupcommon.LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
				backupcommon.MatchExpressionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
					{
						backupcommon.MeKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "labelSelector", matchExpressionsArrayField, "key"),
						backupcommon.MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "labelSelector", matchExpressionsArrayField, "operator"),
						backupcommon.MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "labelSelector", matchExpressionsArrayField, "values"),
					},
				},
				backupcommon.MatchLabelsKey: &tfModelConverterHelper.Map{
					tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "labelSelector", "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
				},
			},
			backupcommon.OrLabelSelectorKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					backupcommon.MatchExpressionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
						{
							backupcommon.MeKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "template", orLabelSelectorsArrayField, matchExpressionsArrayField, "key"),
							backupcommon.MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", orLabelSelectorsArrayField, matchExpressionsArrayField, "operator"),
							backupcommon.MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", orLabelSelectorsArrayField, matchExpressionsArrayField, "values"),
						},
					},
					backupcommon.MatchLabelsKey: &tfModelConverterHelper.Map{
						tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", orLabelSelectorsArrayField, "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
					},
				},
			},
			backupcommon.OrderedResourcesKey: &tfModelConverterHelper.Map{
				tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "orderedResources", tfModelConverterHelper.AllMapKeysFieldMarker),
			},
		},
	},
}

var tfModelDataSourceRequestMap = &tfModelConverterHelper.BlockToStruct{
	SortByKey:            "sortBy",
	QueryKey:             "query",
	IncludeTotalCountKey: "includeTotal",
	backupcommon.NameKey: tfModelConverterHelper.BuildDefaultModelPath("searchScope", "name"),
	ScopeKey: &tfModelConverterHelper.BlockToStruct{
		ClusterGroupScopeKey: &tfModelConverterHelper.BlockToStruct{
			ClusterGroupNameKey: tfModelConverterHelper.BuildDefaultModelPath("searchScope", "clusterGroupName"),
		},
		ClusterScopeKey: &tfModelConverterHelper.BlockToStruct{
			backupcommon.ClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("searchScope", "clusterName"),
			backupcommon.ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("searchScope", "managementClusterName"),
			backupcommon.ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("searchScope", "provisionerName"),
		},
	},
}

var tfModelDataSourceResponseMap = &tfModelConverterHelper.BlockToStruct{
	SchedulesKey: &tfModelConverterHelper.BlockSliceToStructSlice{
		// UNPACK tfModelResourceMap HERE.
	},
	TotalCountKey: "totalCount",
}

var tfModelResourceConverter = tfModelConverterHelper.TFSchemaModelConverter[*clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupSchedule]{
	TFModelMap: tfModelResourceMap,
}

var tfModelDataSourceRequestConverter = tfModelConverterHelper.TFSchemaModelConverter[*clusterbackupschedulemodels.ListBackupSchedulesRequest]{
	TFModelMap: tfModelDataSourceRequestMap,
}

var tfModelDataSourceResponseConverter = tfModelConverterHelper.TFSchemaModelConverter[*clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleListSchedulesResponse]{
	TFModelMap: tfModelDataSourceResponseMap,
}

func constructTFModelDataSourceResponseMap() {
	targetLocationDataSourceSchema := tfModelResourceConverter.UnpackSchema(tfModelConverterHelper.BuildArrayField("schedules"))

	*(*tfModelDataSourceResponseMap)[SchedulesKey].(*tfModelConverterHelper.BlockSliceToStructSlice) = append(
		*(*tfModelDataSourceResponseMap)[SchedulesKey].(*tfModelConverterHelper.BlockSliceToStructSlice),
		targetLocationDataSourceSchema,
	)
}
