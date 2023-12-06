/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	backupcommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const (
	ResourceName = "tanzu-mission-control_backup"
)

var backupResourceSchema = map[string]*schema.Schema{
	backupcommon.NameKey:                  nameSchema,
	backupcommon.ClusterNameKey:           backupcommon.ClusterNameSchema,
	backupcommon.ManagementClusterNameKey: backupcommon.ManagementClusterNameSchema,
	backupcommon.ProvisionerNameKey:       backupcommon.ProvisionerNameSchema,
	backupcommon.BackupScopeKey:           backupcommon.BackupScopeSchema,
	backupcommon.SpecKey:                  backupcommon.SpecSchema,
	common.MetaKey:                        common.Meta,
}

var nameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "The name of the backup.",
	Required:    true,
	ForceNew:    true,
}
