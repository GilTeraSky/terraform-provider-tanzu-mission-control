/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backup

import (
	"fmt"
	"strings"

	backupcommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backup"
	backupres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/backup/standalone"
	clusterres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster"
	dataprotectiontests "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/dataprotection/tests"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

const (
	FullClusterBackupResourceName = "test_full_cluster_backup"
	FullClusterBackupName         = "full-cluster-backup"

	NamespacesBackupResourceName = "test_namespaces_backup"
	NamespacesBackupName         = "namespaces-backup"

	LabelsBackupResourceName = "test_labels_backup"
	LabelsBackupName         = "labels-backup"

	// In some cases, even though the cluster is healthy and enabled with data protection it is still not ready to be backed up,
	// therefore using a delay of X seconds to make sure the tests pass.
	backupDelaySeconds = 30
)

var (
	FullClusterBackupResourceFullName = fmt.Sprintf("%s.%s", backupres.ResourceName, FullClusterBackupResourceName)
	NamespacesBackupResourceFullName  = fmt.Sprintf("%s.%s", backupres.ResourceName, NamespacesBackupResourceName)
	LabelsBackupResourceFullName      = fmt.Sprintf("%s.%s", backupres.ResourceName, LabelsBackupResourceName)
)

type ResourceTFConfigBuilder struct {
	DataProtectionRequiredResource string
	ClusterInfo                    string
}

func InitResourceTFConfigBuilder(scopeHelper *commonscope.ScopeHelperResources) *ResourceTFConfigBuilder {
	dataProtectionConfigBuilder := dataprotectiontests.InitResourceTFConfigBuilder(scopeHelper, dataprotectiontests.RsFullBuild)
	dataProtectionRequiredResource := dataProtectionConfigBuilder.GetEnableDataProtectionConfig()

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

	tfConfigBuilder := &ResourceTFConfigBuilder{
		DataProtectionRequiredResource: strings.Trim(dataProtectionRequiredResource, " "),
		ClusterInfo:                    clusterInfo,
	}

	return tfConfigBuilder
}

func (builder *ResourceTFConfigBuilder) GetFullClusterBackupConfig(backupTargetLocationName string) string {
	return fmt.Sprintf(`
		%s

		resource "%s" "%s" {
		  name = "%s"
          %s

          backup_scope = "%s"

		  spec {		    
            storage_location = "%s"
            backup_ttl = "2592000s"
		    excluded_namespaces = [
			  "app-01",
			  "app-02",
			  "app-03",
			  "app-04"
		    ]
		    excluded_resources = [
			  "secrets",
			  "configmaps"
		    ]
		  }

          depends_on = [time_sleep.wait_%v_seconds]
		}
		
		resource "time_sleep" "wait_%[8]v_seconds" {
		  depends_on =  [%[9]s]
		
		  create_duration = "%[8]vs"
		}
		`,
		builder.DataProtectionRequiredResource,
		backupres.ResourceName,
		FullClusterBackupResourceName,
		FullClusterBackupName,
		builder.ClusterInfo,
		backupcommon.FullClusterBackupScope,
		backupTargetLocationName,
		backupDelaySeconds,
		dataprotectiontests.EnableDataProtectionResourceFullName)
}

func (builder *ResourceTFConfigBuilder) GetNamespacesBackupConfig(backupTargetLocationName string) string {
	return fmt.Sprintf(`
		%s
		
		resource "%s" "%s" {
          name = "%s"
		  %s
          backup_scope = "%s"
		
		  spec {		
			
		    included_namespaces = [
		  	"app-01",
		  	"app-02",
		  	"app-03",
		  	"app-04"
		    ]
		    backup_ttl = "86400s"
		    excluded_resources = [
		  	"secrets",
		  	"configmaps"
		    ]
		    include_cluster_resources    = true
		    storage_location = "%s"
		    hooks {
		  	resource {
		  	  name = "sample-config"
		  	  pre_hook {
		  		exec {
		  		  command = ["echo 'hello'"]
		  		  container = "workload"
		  		  on_error  = "CONTINUE"
		  		  timeout   = "10s"
		  		}
		  	  }
		  	  pre_hook {
		  		exec {
		  		  command = ["echo 'hello'"]
		  		  container = "db"
		  		  on_error  = "CONTINUE"
		  		  timeout   = "30s"
		  		}
		  	  }
		  	  post_hook {
		  		exec {
		  		  command = ["echo 'goodbye'"]
		  		  container = "db"
		  		  on_error  = "CONTINUE"
		  		  timeout   = "60s"
		  		}
		  	  }
		  	  post_hook {
		  	  	exec {
		  	  	  command = ["echo 'goodbye'"]
		  	  	  container = "workload"
		  	  	  on_error  = "FAIL"
		        	  timeout   = "20s"
                  }
		        }
		      }
		    }
		  }

          depends_on = [time_sleep.wait_%v_seconds]
		}
		
		resource "time_sleep" "wait_%[8]v_seconds" {
		  depends_on =  [%[9]s]
		
		  create_duration = "%[8]vs"
		}
		`,
		builder.DataProtectionRequiredResource,
		backupres.ResourceName,
		NamespacesBackupResourceName,
		NamespacesBackupName,
		builder.ClusterInfo,
		backupcommon.NamespacesBackupScope,
		backupTargetLocationName,
		backupDelaySeconds,
		dataprotectiontests.EnableDataProtectionResourceFullName)
}

func (builder *ResourceTFConfigBuilder) GetLabelsBackupConfig(backupTargetLocationName string) string {
	return fmt.Sprintf(`
		%s

		resource "%s" "%s" {
		  name = "%s"
		  %s

		  backup_scope = "%s"

          spec {
			  default_volumes_to_fs_backup = false
			  include_cluster_resources = true
			  backup_ttl = "604800s"
			  storage_location = "%s"
			  label_selector {
				match_expression {
				  key      = "apps.tanzu.vmware.com/tap-ns"
				  operator = "Exists"
				}
				match_expression {
				  key      = "apps.tanzu.vmware.com/exclude-from-backup"
				  operator = "DoesNotExist"
				}
			  }
		  }

          depends_on = [time_sleep.wait_%v_seconds]
		}
		
		resource "time_sleep" "wait_%[8]v_seconds" {
		  depends_on =  [%[9]s]
		
		  create_duration = "%[8]vs"
		}
		`,
		builder.DataProtectionRequiredResource,
		backupres.ResourceName,
		LabelsBackupResourceName,
		LabelsBackupName,
		builder.ClusterInfo,
		backupcommon.LabelSelectorBackupScope,
		backupTargetLocationName,
		backupDelaySeconds,
		dataprotectiontests.EnableDataProtectionResourceFullName)
}
