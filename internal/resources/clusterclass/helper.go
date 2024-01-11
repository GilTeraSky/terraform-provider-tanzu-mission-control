package clusterclass

import (
	"encoding/base64"
	"encoding/json"

	openapiv3 "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/openapi_v3_utils"
	clusterclassmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clusterclass"
)

func BuildClusterClassMap(spec *clusterclassmodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassSpec) map[string]interface{} {
	openAPIV3Schema := make(map[string]interface{})

	for _, v := range spec.Variables {
		decodedTemplate, _ := base64.StdEncoding.DecodeString(v.Schema.Template.Raw.String())
		templateJSON := make(map[string]interface{})
		_ = json.Unmarshal(decodedTemplate, &templateJSON)
		templateSchema := templateJSON["openAPIV3Schema"].(map[string]interface{})

		_, defaultExist := templateSchema[string(openapiv3.DefaultKey)]
		_, requiredExist := templateSchema[string(openapiv3.RequiredKey)]

		if !requiredExist && !defaultExist && v.Required {
			templateSchema[string(openapiv3.RequiredKey)] = true
		}

		openAPIV3Schema[v.Name] = templateSchema
	}

	return openAPIV3Schema
}

func generateClusterVariablesTemplate(clusterVariables map[string]interface{}) (schemaTemplate map[string]interface{}) {
	schemaTemplate = make(map[string]interface{})

	for k, v := range clusterVariables {
		schemaTemplate[k], _ = openapiv3.BuildOpenAPIV3Template(v.(map[string]interface{}))
	}

	return schemaTemplate
}
