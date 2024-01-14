//go:build custompolicytemplate
// +build custompolicytemplate

/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplate

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	custompolicytemplateres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/custompolicytemplate"
)

func initTestProvider(t *testing.T) *schema.Provider {
	testAccProvider := &schema.Provider{
		Schema: authctx.ProviderAuthSchema(),
		ResourcesMap: map[string]*schema.Resource{
			custompolicytemplateres.ResourceName: custompolicytemplateres.ResourceCustomPolicyTemplate(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			custompolicytemplateres.ListDataSourceName: custompolicytemplateres.DataSourceListCustomTemplates(),
			custompolicytemplateres.DataSourceName:     custompolicytemplateres.DataSourceCustomTemplate(),
		},
		ConfigureContextFunc: authctx.ProviderConfigureContext,
	}

	if err := testAccProvider.InternalValidate(); err != nil {
		require.NoError(t, err)
	}

	return testAccProvider
}
