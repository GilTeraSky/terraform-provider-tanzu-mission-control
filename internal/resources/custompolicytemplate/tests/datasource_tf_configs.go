/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplate

import (
	"fmt"

	custompolicytemplateres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/custompolicytemplate"
)

const (
	ListDataSourceName = "list_custom_policy_templates"
	DataSourceName     = "custom_policy_template"
)

var (
	DataSourceFullName     = fmt.Sprintf("data.%s.%s", custompolicytemplateres.DataSourceName, DataSourceName)
	ListDataSourceFullName = fmt.Sprintf("data.%s.%s", custompolicytemplateres.ListDataSourceName, ListDataSourceName)
)

func GetDataSourceConfig() string {
	return fmt.Sprintf(`
		data "%s" "%s" {
		  name = "%s"
		}
		`,
		custompolicytemplateres.DataSourceName,
		DataSourceName,
		CustomPolicyTemplateName)
}

func GetListDataSourceConfig() string {
	return fmt.Sprintf(`
		data "%s" "%s" {
		}
		`,
		custompolicytemplateres.ListDataSourceName,
		ListDataSourceName)
}
