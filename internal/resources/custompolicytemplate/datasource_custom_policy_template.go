/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	openapiv3 "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/openapi_v3_utils"
	custompolicytemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/custompolicytemplate"
	recipemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/recipe"
)

func DataSourceCustomTemplate() *schema.Resource {
	// Unpack resource map to datasource map.
	constructTFModelListDataSourceResponseMap()

	return &schema.Resource{
		ReadContext: dataSourceCustomTemplatesRead,
		Schema:      templatesDataSourceSchema,
	}
}

func dataSourceCustomTemplatesRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	customTemplateName := data.Get(NameKey).(string)
	customTemplateFn := &custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName{
		Name: customTemplateName,
	}

	templateResp, err := config.TMCConnection.CustomPolicyTemplateResourceService.CustomPolicyTemplateResourceServiceGet(customTemplateFn)

	switch {
	case err != nil:
		return diag.FromErr(errors.Wrapf(err, "Couldn't get custom template '%s'.", customTemplateName))
	case templateResp.Template == nil:
		data.SetId(fmt.Sprintf("NO_DATA_%s", customTemplateName))
	default:
		_ = data.Set(TemplateManifestKey, templateResp.Template.Spec.Object)

		recipeData, err := getRecipeData(&config, customTemplateName)

		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "Couldn't get recipe for custom template '%s'.", customTemplateName))
		}

		recipeSchemaJSON := make(map[string]interface{})
		_ = json.Unmarshal([]byte(recipeData.Recipe.Spec.InputSchema), &recipeSchemaJSON)

		// This is a fix due to missing types in array definition for TargetKubernetesResources field.
		recipeSchemaJSON["properties"].(map[string]interface{})["targetKubernetesResources"].(map[string]interface{})["items"].(map[string]interface{})["type"] = "object"

		recipeTemplateJSON, _ := openapiv3.BuildOpenAPIV3Template(recipeSchemaJSON)
		recipeTemplateBytes, _ := json.Marshal(recipeTemplateJSON)
		_ = data.Set(RecipeTemplateKey, helper.ConvertToString(recipeTemplateBytes, ""))
		_ = data.Set(RecipeSchemaKey, recipeData.Recipe.Spec.InputSchema)

		data.SetId(customTemplateName)
	}

	return diags
}

func getRecipeData(config *authctx.TanzuContext, customTemplateName string) (*recipemodels.VmwareTanzuManageV1alpha1PolicyTypeRecipeData, error) {
	recipeFullName := &recipemodels.VmwareTanzuManageV1alpha1PolicyTypeRecipeFullName{
		TypeName: "custom-policy",
		Name:     customTemplateName,
	}

	recipeResp, err := config.TMCConnection.RecipeResourceService.RecipeResourceServiceGet(recipeFullName)

	return recipeResp, err
}
