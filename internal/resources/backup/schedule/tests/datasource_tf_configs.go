/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupscheduletests

import (
	"fmt"

	backupcommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backup"
	backupscheduleres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backup/schedule"
	clusterres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

type DataSourceBuildMode string

const (
	DsFullBuild  DataSourceBuildMode = "FULL"
	DsNoParentRs DataSourceBuildMode = "NO_PARENT_RESOURCE"
)

const (
	DataSourceName = "test_cluster_scope"
)

var (
	DataSourceFullName = fmt.Sprintf("data.%s.%s", backupscheduleres.ResourceName, DataSourceName)
)

type DataSourceTFConfigBuilder struct {
	BackupScheduleRequiredResource string
	ClusterInfo                    string
}

func InitDataSourceTFConfigBuilder(scopeHelper *commonscope.ScopeHelperResources, resourceConfigBuilder *ResourceTFConfigBuilder, bMode DataSourceBuildMode) *DataSourceTFConfigBuilder {
	var backupScheduleRequiredResource string

	if bMode != DsNoParentRs {
		backupScheduleRequiredResource = resourceConfigBuilder.GetLabelsBackupScheduleConfig()
	}

	mgmtClusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ManagementClusterNameKey)
	clusterName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.NameKey)
	provisionerName := fmt.Sprintf("%s.%s", scopeHelper.Cluster.ResourceName, clusterres.ProvisionerNameKey)
	clusterInfo := fmt.Sprintf(`
		%s = %s
		%s = %s        
		%s = %s  
		`,
		backupcommon.ClusterNameKey, clusterName,
		backupcommon.ManagementClusterNameKey, mgmtClusterName,
		backupcommon.ProvisionerNameKey, provisionerName)

	tfConfigBuilder := &DataSourceTFConfigBuilder{
		BackupScheduleRequiredResource: backupScheduleRequiredResource,
		ClusterInfo:                    clusterInfo,
	}

	return tfConfigBuilder
}

func (builder *DataSourceTFConfigBuilder) GetDataSourceConfig() string {
	return fmt.Sprintf(`
		%s

		data "%s" "%s" {
          name = "%s"
		  scope {
			cluster {
				%s
			}
		  }

          depends_on = [%s]
		}
		`,
		builder.BackupScheduleRequiredResource,
		backupscheduleres.ResourceName,
		DataSourceName,
		LabelsBackupScheduleName,
		builder.ClusterInfo,
		LabelsBackupScheduleResourceFullName)
}
