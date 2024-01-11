/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplate

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	custompolicytemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/custompolicytemplate"
)

func DataSourceListCustomTemplates() *schema.Resource {
	// Unpack resource map to datasource map.
	constructTFModelListDataSourceResponseMap()

	return &schema.Resource{
		ReadContext: dataSourceListCustomTemplatesRead,
		Schema:      listTemplatesDataSourceSchema,
	}
}

func dataSourceListCustomTemplatesRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var resp *custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateListData

	config := m.(authctx.TanzuContext)
	request, err := tfModelListDataSourceRequestConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't list custom templates."))
	}

	resp, err = config.TMCConnection.CustomPolicyTemplateResourceService.CustomPolicyTemplateResourceServiceList(request)

	switch {
	case err != nil:
		return diag.FromErr(errors.Wrap(err, "Couldn't list custom templates."))
	case resp.Templates == nil:
		data.SetId("NO_DATA")
	default:
		err = tfModelListDataSourceResponseConverter.FillTFSchema(resp, data)

		if err != nil {
			diags = diag.FromErr(err)
		}

		requestData := []string{
			helper.ConvertToString(request.IncludeTotalCount, ""),
		}

		if request.TemplateName != "" {
			requestData = append(requestData, request.TemplateName)
		}

		if request.SortBy != "" {
			requestData = append(requestData, request.TemplateName)
		}

		if request.Query != "" {
			requestData = append(requestData, request.TemplateName)
		}

		dataSourceHash, _ := helper.Hash(strings.Join(requestData, ","))
		dataSourceID := fmt.Sprintf("list_custom_templates_%s", dataSourceHash)
		data.SetId(dataSourceID)
	}

	return diags
}
