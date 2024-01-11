/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplate

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	ListDataSourceName = "tanzu-mission-control_custom_policy_template_list"
	DataSourceName     = "tanzu-mission-control_custom_policy_template"

	// List Data Source Directives.
	SortByKey            = "sort_by"
	QueryKey             = "query"
	IncludeTotalCountKey = "include_total_count"
	TemplatesKey         = "templates"
	TotalCountKey        = "total_count"

	// Data Source Directives.
	RecipeTemplateKey = "recipe_template"
	RecipeSchemaKey   = "recipe_schema"
)

// List Data Source Schema.
var (
	listTemplatesDataSourceSchema = map[string]*schema.Schema{
		NameKey:              nameDSSchema,
		SortByKey:            sortBySchema,
		QueryKey:             querySchema,
		IncludeTotalCountKey: includeTotalSchema,
		TemplatesKey:         templatesSchema,
		TotalCountKey:        totalCountSchema,
	}

	nameDSSchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "The name of the template, supports globbing; default (*).",
		Optional:    true,
	}

	sortBySchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Sort templates by field.",
		Optional:    true,
	}

	querySchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Define a query for listing templates",
		Optional:    true,
	}

	includeTotalSchema = &schema.Schema{
		Type:        schema.TypeBool,
		Description: "Whether to include total count of templates listed.\n(Default: True)",
		Optional:    true,
		Default:     true,
	}

	templatesSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "A list of templates returned",
		Computed:    true,
		Elem: &schema.Resource{
			Schema: customPolicyTemplateResourceSchema,
		},
	}

	totalCountSchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Total count of schedules returned",
		Computed:    true,
	}
)

// templatesDataSourceSchema Data Source Schema.
var templatesDataSourceSchema = map[string]*schema.Schema{
	NameKey: nameSchema,
	TemplateManifestKey: {
		Type:        schema.TypeString,
		Description: "YAML formatted Kubernetes resource.",
		Computed:    true,
	},
	RecipeTemplateKey: {
		Type:        schema.TypeString,
		Description: "JSON encoded OpenAPI template.",
		Computed:    true,
	},
	RecipeSchemaKey: {
		Type:        schema.TypeString,
		Description: "JSON encoded OpenAPI schema of the template.",
		Computed:    true,
	},
}
