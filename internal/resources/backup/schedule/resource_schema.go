/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

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
	backupcommon.BackupScopeKey: backupScopeSchema,
	backupcommon.SpecKey:        specSchema,
	common.MetaKey:              common.Meta,
}

var nameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "The name of the backup schedule",
	Required:    true,
	ForceNew:    true,
}

var scopeSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Scope block for Back up schedule (cluster/cluster group)",
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
						backupcommon.ClusterNameKey:           clusterNameSchema,
						backupcommon.ManagementClusterNameKey: managementClusterNameSchema,
						backupcommon.ProvisionerNameKey:       provisionerNameSchema,
					},
				},
			},
		},
	},
}

var managementClusterNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Management cluster name",
	Required:    true,
	ForceNew:    true,
}

var provisionerNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Cluster provisioner name",
	Required:    true,
	ForceNew:    true,
}

var clusterNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Cluster name",
	Required:    true,
	ForceNew:    true,
}

var backupScopeSchema = &schema.Schema{
	Type:             schema.TypeString,
	Description:      fmt.Sprintf("Scope for backup schedule.\nValid values are (%s)", strings.Join(backupcommon.ScopeValidValues, ", ")),
	Required:         true,
	ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(backupcommon.ScopeValidValues, false)),
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
			TemplateKey: templateSchema,
		},
	},
}

var templateSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Backup schedule template block, backup definition to be run on the provided schedule",
	MaxItems:    1,
	MinItems:    1,
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			backupcommon.CsiSnapshotTimeoutKey: {
				Description: "Specifies the time used to wait for CSI VolumeSnapshot status turns to ReadyToUse during creation, before returning error as timeout.\nThe default value is 10 minute.\nFormat is the time number and time sign, example: \"50s\" (50 seconds)",
				Type:        schema.TypeString,
				Optional:    true,
			},
			backupcommon.DefaultVolumesToFsBackupKey: {
				Type:        schema.TypeBool,
				Description: "Specifies whether all pod volumes should be backed up via file system backup by default.\n(Default: True)",
				Optional:    true,
				Default:     true,
			},
			backupcommon.DefaultVolumesToResticKey: {
				Type:        schema.TypeBool,
				Description: "Specifies whether restic should be used to take a backup of all pod volumes by default.\n(Default: False)",
				Optional:    true,
				Default:     false,
			},
			backupcommon.ExcludedNamespacesKey: {
				Type:        schema.TypeList,
				Description: "The namespaces to be excluded in the backup.\nCan't be used if scope is SET_NAMESPACES.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			backupcommon.SystemExcludedNamespacesKey: {
				Type:        schema.TypeList,
				Description: "System excluded namespaces for state.",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			backupcommon.ExcludedResourcesKey: {
				Type:        schema.TypeList,
				Description: "The name list for the resources to be excluded in backup.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			backupcommon.IncludedNamespacesKey: {
				Type:        schema.TypeList,
				Description: "The namespace to be included for backup from.\nIf empty, all namespaces are included.\nCan't be used if scope is FULL_CLUSTER.\nRequired if scope is SET_NAMESPACES.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			backupcommon.IncludedResourcesKey: {
				Type:        schema.TypeList,
				Description: "The name list for the resources to be included into backup. If empty, all resources are included.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			backupcommon.IncludeClusterResourcesKey: {
				Type:        schema.TypeBool,
				Description: "A flag which specifies whether cluster-scoped resources should be included for consideration in the backup.\nIf set to true, all cluster-scoped resources will be backed up. If set to false, all cluster-scoped resources will be excluded from the backup.\nIf unset, all cluster-scoped resources are included if and only if all namespaces are included and there are no excluded namespaces.\nOtherwise, only cluster-scoped resources associated with namespace-scoped resources included in the backup spec are backed up.\nFor example, if a PersistentVolumeClaim is included in the backup, its associated PersistentVolume (which is cluster-scoped) would also be backed up.\n(Default: False)",
				Optional:    true,
				Default:     false,
			},
			backupcommon.OrderedResourcesKey: {
				Type:        schema.TypeMap,
				Description: "Specifies the backup order of resources of specific Kind. The map key is the Kind name and value is a list of resource names separated by commas.\nEach resource name has format \"namespace/resourcename\".\nFor cluster resources, simply use \"resourcename\".",
				Optional:    true,
			},
			backupcommon.SnapshotVolumesKey: {
				Type:        schema.TypeBool,
				Description: "A flag which specifies whether to take cloud snapshots of any PV's referenced in the set of objects included in the Backup.\nIf set to true, snapshots will be taken, otherwise, snapshots will be skipped.\nIf left unset, snapshots will be attempted if volume snapshots are configured for the cluster.",
				Optional:    true,
				Default:     false,
			},
			backupcommon.StorageLocationKey: {
				Type:        schema.TypeString,
				Description: "The name of a BackupStorageLocation where the backup should be stored.",
				Optional:    true,
			},
			backupcommon.BackupTTLKey: {
				Type:        schema.TypeString,
				Description: "The backup retention period.",
				Optional:    true,
			},
			backupcommon.VolumeSnapshotLocationsKey: {
				Type:        schema.TypeList,
				Description: "A list containing names of VolumeSnapshotLocations associated with this backup.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			backupcommon.LabelSelectorKey: {
				Type:        schema.TypeList,
				Description: "The label selector to selectively adding individual objects to the backup schedule.\nIf not specified, all objects are included.\nCan't be used if scope is FULL_CLUSTER or SET_NAMESPACES.\nRequired if scope is LABEL_SELECTOR and Or Label Selectors are not defined",
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: labelSelectorResource,
				},
			},
			backupcommon.OrLabelSelectorKey: {
				Type:        schema.TypeList,
				Description: "(Repeatable Block) A list of label selectors to filter with when adding individual objects to the backup.\nIf multiple provided they will be joined by the OR operator.\nLabelSelector as well as OrLabelSelectors cannot co-exist in backup request, only one of them can be used.\nCan't be used if scope is FULL_CLUSTER or SET_NAMESPACES.\nRequired if scope is LABEL_SELECTOR and Label Selector is not defined",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: labelSelectorResource,
				},
			},
			backupcommon.HooksKey: hooksSchema,
		},
	},
}

var hooksSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Hooks block represent custom actions that should be executed at different phases of the backup.",
	MaxItems:    1,
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			backupcommon.ResourceKey: {
				Type:        schema.TypeList,
				Description: "(Repeatable Block) Resources are hooks that should be executed when backing up individual instances of a resource.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						backupcommon.NameKey: {
							Type:        schema.TypeString,
							Description: "The name of the hook resource.",
							Required:    true,
						},
						backupcommon.ExcludedNamespacesKey: {
							Type:        schema.TypeList,
							Description: "Specifies the namespaces to which this hook spec does not apply.",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						backupcommon.IncludedNamespacesKey: {
							Type:        schema.TypeList,
							Description: "Specifies the namespaces to which this hook spec applies.\nIf empty, it applies to all namespaces.",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						backupcommon.LabelSelectorKey: {
							Type:        schema.TypeList,
							Description: "The label selector to selectively adding individual objects to the hook resource.\nIf not specified, all objects are included.",
							MaxItems:    1,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: labelSelectorResource,
							},
						},
						backupcommon.PostHookKey: pHookSchema,
						backupcommon.PreHookKey:  pHookSchema,
					},
				},
			},
		},
	},
}

var labelSelectorResource = map[string]*schema.Schema{
	backupcommon.MatchExpressionKey: {
		Type:        schema.TypeList,
		Description: "(Repeatable Block) A list of label selector requirements. The requirements are ANDed.",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				backupcommon.MeKey: {
					Type:        schema.TypeString,
					Description: "Key is the label key that the selector applies to.",
					Required:    true,
				},
				backupcommon.MeOperatorKey: {
					Type:        schema.TypeString,
					Description: "Operator represents a key's relationship to a set of values.\nValid operators are \"In\", \"NotIn\", \"Exists\" and \"DoesNotExist\".",
					Required:    true,
				},
				backupcommon.MeValuesKey: {
					Type:        schema.TypeList,
					Description: "Values is an array of string values.\nIf the operator is \"In\" or \"NotIn\", the values array must be non-empty.\nIf the operator is \"Exists\" or \"DoesNotExist\", the values array must be empty.\nThis array is replaced during a strategic merge patch.",
					Optional:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	},
	backupcommon.MatchLabelsKey: {
		Type:        schema.TypeMap,
		Description: "A map of {key,value} pairs. A single {key,value} in the map is equivalent to an element of match_expressions, whose key field is \"key\", the operator is \"In\" and the values array contains only \"value\".\nThe requirements are ANDed.",
		Optional:    true,
	},
}

var pHookSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "(Repeatable Block) A list of backup hooks to execute after storing the item in the backup.\nThese are executed after all \"additional items\" from item actions are processed.",
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			backupcommon.ExecKey: {
				Type:        schema.TypeList,
				Description: "Exec block defines an exec hook.",
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						backupcommon.CommandKey: {
							Type:        schema.TypeList,
							Description: "The command and arguments to execute.",
							Required:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						backupcommon.ContainerKey: {
							Type:        schema.TypeString,
							Description: "The container in the pod where the command should be executed.\nIf not specified, the pod's first container is used.",
							Required:    true,
						},
						backupcommon.OnErrorKey: {
							Type:             schema.TypeString,
							Description:      fmt.Sprintf("Specifies how Velero should behave if it encounters an error executing this hook.\nValid values are (%s)", strings.Join(backupcommon.OnErrorValidValues, ", ")),
							Optional:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(backupcommon.OnErrorValidValues, false)),
						},
						backupcommon.TimeoutKey: {
							Type:        schema.TypeString,
							Description: "Defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.",
							Optional:    true,
						},
					},
				},
			},
		},
	},
}
