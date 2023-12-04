/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	backupcommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const (
	ResourceName = "tanzu-mission-control_backup_schedule"

	// Spec Directive Keys.
	PausedKey   = "paused"
	ScheduleKey = "schedule"
	TemplateKey = "template"

	// Schedule Directive Keys.
	RateKey = "rate"
)

var backupScheduleResourceSchema = map[string]*schema.Schema{
	backupcommon.NameKey:        nameSchema,
	backupcommon.ScopeKey:       scopeSchema,
	backupcommon.BackupScopeKey: backupcommon.BackupScopeSchema,
	backupcommon.SpecKey:        specSchema,
	common.MetaKey:              common.Meta,
}

var nameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "The name of the backup schedule.",
	Required:    true,
	ForceNew:    true,
}

var scopeSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Scope block for Back up schedule (cluster/cluster group).",
	Required:    true,
	MaxItems:    1,
	Optional:    false,
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

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Backup schedule spec block",
	Required:    true,
	MaxItems:    1,
	MinItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			PausedKey: {
				Type:        schema.TypeBool,
				Description: "Paused specifies whether the schedule is paused or not. (Default: False)",
				Optional:    true,
				Default:     false,
			},
			ScheduleKey: {
				Type:        schema.TypeList,
				Description: "Schedule block",
				MaxItems:    1,
				MinItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						RateKey: {
							Type:        schema.TypeString,
							Description: "Cron expression of backup schedule rate/interval",
							Required:    true,
						},
					},
				},
			},
			TemplateKey: backupcommon.SpecSchema,
		},
	},
}
