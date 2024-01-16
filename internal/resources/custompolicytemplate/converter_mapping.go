/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplate

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	custompolicytemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/custompolicytemplate"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var (
	dataInventoryArrayField = tfModelConverterHelper.BuildArrayField("dataInventory")
)

// Resource Converter Objects.
var (
	tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
		NameKey:        tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
		common.MetaKey: common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
		SpecKey: &tfModelConverterHelper.BlockToStruct{
			IsDeprecatedKey:     tfModelConverterHelper.BuildDefaultModelPath("spec", "deprecated"),
			TemplateManifestKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "object"),
			ObjectTypeKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "objectType"),
			TemplateTypeKey:     tfModelConverterHelper.BuildDefaultModelPath("spec", "templateType"),
			DataInventoryKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					GroupKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", dataInventoryArrayField, "group"),
					VersionKey: tfModelConverterHelper.BuildDefaultModelPath("spec", dataInventoryArrayField, "kind"),
					KindKey:    tfModelConverterHelper.BuildDefaultModelPath("spec", dataInventoryArrayField, "version"),
				},
			},
		},
	}

	tfModelResourceConverter = tfModelConverterHelper.TFSchemaModelConverter[*custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplate]{
		TFModelMap: tfModelResourceMap,
	}
)

// Data Source List Converter Objects.
var (
	tfModelListDataSourceRequestMap = &tfModelConverterHelper.BlockToStruct{
		NameKey:              "templateName",
		SortByKey:            "sortBy",
		QueryKey:             "query",
		IncludeTotalCountKey: "includeTotal",
	}

	tfModelListDataSourceResponseMap = &tfModelConverterHelper.BlockToStruct{
		TemplatesKey: &tfModelConverterHelper.BlockSliceToStructSlice{
			// UNPACK tfModelResourceMap HERE.
		},
		TotalCountKey: "totalCount",
	}

	tfModelListDataSourceRequestConverter = tfModelConverterHelper.TFSchemaModelConverter[*custompolicytemplatemodels.ListCustomTemplatesRequest]{
		TFModelMap: tfModelListDataSourceRequestMap,
	}

	tfModelListDataSourceResponseConverter = tfModelConverterHelper.TFSchemaModelConverter[*custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateListData]{
		TFModelMap: tfModelListDataSourceResponseMap,
	}
)

// Unpacks & populates tfModelResourceMap to tfModelListDataSourceResponseMap.
func constructTFModelListDataSourceResponseMap() {
	targetLocationDataSourceSchema := tfModelResourceConverter.UnpackSchema(tfModelConverterHelper.BuildArrayField("templates"))

	*(*tfModelListDataSourceResponseMap)[TemplatesKey].(*tfModelConverterHelper.BlockSliceToStructSlice) = append(
		*(*tfModelListDataSourceResponseMap)[TemplatesKey].(*tfModelConverterHelper.BlockSliceToStructSlice),
		targetLocationDataSourceSchema,
	)
}
