/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplateclient

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	custompolicytemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/custompolicytemplate"
)

const (
	// API Paths.
	customPolicyTemplateAPIVersionAndGroup = "v1alpha1/policy/templates"

	// Params.
	templateNameQueryParam      = "searchScope.name"
	sortedByQueryParam          = "sortBy"
	queryQueryParam             = "query"
	includeTotalCountQueryParam = "includeTotalCount"
)

// New creates a new custom policy template resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for custom policy template resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	CustomPolicyTemplateResourceServiceCreate(request *custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData) (*custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error)

	CustomPolicyTemplateResourceServiceUpdate(request *custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData) (*custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error)

	CustomPolicyTemplateResourceServiceDelete(fn *custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName) error

	CustomPolicyTemplateResourceServiceGet(fn *custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName) (*custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error)

	CustomPolicyTemplateResourceServiceList(request *custompolicytemplatemodels.ListCustomTemplatesRequest) (*custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateListData, error)
}

/*
CustomPolicyTemplateResourceServiceGet gets an custom policy template.
*/
func (c *Client) CustomPolicyTemplateResourceServiceGet(fullName *custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName) (*custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error) {
	requestURL := helper.ConstructRequestURL(customPolicyTemplateAPIVersionAndGroup, fullName.Name).String()
	resp := &custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData{}
	err := c.Get(requestURL, resp)

	return resp, err
}

/*
CustomPolicyTemplateResourceServiceCreate creates an custom policy template.
*/
func (c *Client) CustomPolicyTemplateResourceServiceCreate(request *custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData) (*custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error) {
	response := &custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData{}
	requestURL := helper.ConstructRequestURL(customPolicyTemplateAPIVersionAndGroup).String()
	err := c.Create(requestURL, request, response)

	return response, err
}

/*
CustomPolicyTemplateResourceServiceUpdate updates an custom policy template.
*/
func (c *Client) CustomPolicyTemplateResourceServiceUpdate(request *custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData) (*custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData, error) {
	response := &custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateData{}
	requestURL := helper.ConstructRequestURL(customPolicyTemplateAPIVersionAndGroup, request.Template.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}

/*
CustomPolicyTemplateResourceServiceDelete deletes an custom policy template.
*/
func (c *Client) CustomPolicyTemplateResourceServiceDelete(fullName *custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateFullName) error {
	requestURL := helper.ConstructRequestURL(customPolicyTemplateAPIVersionAndGroup, fullName.Name).String()

	return c.Delete(requestURL)
}

/*
CustomPolicyTemplateResourceServiceList lists custom policy templates.
*/
func (c *Client) CustomPolicyTemplateResourceServiceList(request *custompolicytemplatemodels.ListCustomTemplatesRequest) (*custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateListData, error) {
	resp := &custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplateListData{}
	requestURL := helper.ConstructRequestURL(customPolicyTemplateAPIVersionAndGroup)
	queryParams := url.Values{}

	queryParams.Add(includeTotalCountQueryParam, helper.ConvertToString(request.IncludeTotalCount, ""))

	if request.TemplateName != "" {
		queryParams.Add(templateNameQueryParam, request.TemplateName)
	}

	if request.SortBy != "" {
		queryParams.Add(sortedByQueryParam, request.SortBy)
	}

	if request.Query != "" {
		queryParams.Add(queryQueryParam, request.Query)
	}

	if len(queryParams) > 0 {
		requestURL = requestURL.AppendQueryParams(queryParams)
	}

	err := c.Get(requestURL.String(), resp)

	return resp, err
}
