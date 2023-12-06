/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	backupcommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backup"
)

const (
	// Root Directive Keys.
	SortByKey            = "sort_by"
	QueryKey             = "query"
	IncludeTotalCountKey = "include_total_count"
	SchedulesKey         = "schedules"
	TotalCountKey        = "total_count"
	ClusterScopeKey      = "cluster"
	ClusterGroupScopeKey = "cluster_group"
	ClusterGroupNameKey  = "cluster_group_name"
)

var backupScheduleDataSourceSchema = map[string]*schema.Schema{
	backupcommon.NameKey:  nameSchema,
	backupcommon.ScopeKey: searchScopeSchema,
	SortByKey:             sortBySchema,
	QueryKey:              querySchema,
	IncludeTotalCountKey:  includeTotalSchema,
	SchedulesKey:          schedulesSchema,
	TotalCountKey:         totalCountSchema,
}

var (
	sortBySchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Sort backups by field.",
		Optional:    true,
	}

	querySchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Define a query for listing backups",
		Optional:    true,
	}

	includeTotalSchema = &schema.Schema{
		Type:        schema.TypeBool,
		Description: "Whether to include total count of backups.\n(Default: True)",
		Optional:    true,
		Default:     true,
	}

	schedulesSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "A list of schedules returned",
		Computed:    true,
		Elem: &schema.Resource{
			Schema: backupScheduleResourceSchema,
		},
	}

	totalCountSchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Total count of schedules returned",
		Computed:    true,
	}

	searchScopeSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Search scope block",
		MaxItems:    1,
		Required:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				ClusterScopeKey: {
					Type:        schema.TypeList,
					Optional:    true,
					Description: "Cluster scope block",
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							backupcommon.ClusterNameKey:           backupcommon.ClusterNameSchema,
							backupcommon.ManagementClusterNameKey: backupcommon.ManagementClusterNameSchema,
							backupcommon.ProvisionerNameKey:       backupcommon.ProvisionerNameSchema,
						},
					},
				},
			},
		},
	}
)
