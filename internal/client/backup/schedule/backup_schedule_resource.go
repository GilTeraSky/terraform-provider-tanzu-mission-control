/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	"net/url"

	"github.com/pkg/errors"

	backupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/backup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	backupcommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/common"
	clusterbackupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/schedule/cluster"
)

const (
	// API Paths.
	dataProtectionSchedulePath = "dataprotection/schedules"

	// Query Params.
	managementClusterNameSearchQueryParam = "searchScope.managementClusterName"
	provisionerNameSearchQueryParam       = "searchScope.provisionerName"
	scheduleNameSearchQueryParam          = "searchScope.name"
)

// New creates a new backup schedule resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for backup schedule resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {

	// Cluster Backup Schedule API.
	BackupScheduleResourceServiceCreate(request *clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleRequest) (*clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleResponse, error)

	BackupScheduleResourceServiceUpdate(request *clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleRequest) (*clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleResponse, error)

	BackupScheduleResourceServiceDelete(fn *backupcommon.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName) error

	BackupScheduleResourceServiceGet(fn *backupcommon.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName) (*clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleResponse, error)

	BackupScheduleResourceServiceList(request *clusterbackupschedulemodels.ListBackupSchedulesRequest) (*clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleListSchedulesResponse, error)
}

/*
BackupScheduleResourceServiceCreate creates a backup schedule.
*/
func (c *Client) BackupScheduleResourceServiceCreate(request *clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleRequest) (*clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleResponse, error) {
	response := &clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleResponse{}
	requestURL := helper.ConstructRequestURL(backupclient.ClusterApiVersionAndGroup, request.Schedule.FullName.ClusterName, dataProtectionSchedulePath).String()
	err := c.Create(requestURL, request, response)

	return response, err
}

/*
BackupScheduleResourceServiceUpdate updates a backup schedule.
*/
func (c *Client) BackupScheduleResourceServiceUpdate(request *clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleRequest) (*clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleResponse, error) {
	response := &clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleResponse{}
	requestURL := helper.ConstructRequestURL(backupclient.ClusterApiVersionAndGroup, request.Schedule.FullName.ClusterName, dataProtectionSchedulePath, request.Schedule.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}

/*
BackupScheduleResourceServiceDelete deletes a backup schedule.
*/
func (c *Client) BackupScheduleResourceServiceDelete(fullName *backupcommon.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName) error {
	requestURL := helper.ConstructRequestURL(backupclient.ClusterApiVersionAndGroup, fullName.ClusterName, dataProtectionSchedulePath, fullName.Name)
	queryParams := url.Values{}

	queryParams.Add(backupclient.ManagementClusterNameQueryParam, fullName.ManagementClusterName)
	queryParams.Add(backupclient.ProvisionerNameQueryParam, fullName.ProvisionerName)

	requestURL = requestURL.AppendQueryParams(queryParams)

	return c.Delete(requestURL.String())
}

/*
BackupScheduleResourceServiceGet gets a backup schedule.
*/
func (c *Client) BackupScheduleResourceServiceGet(fullName *backupcommon.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName) (*clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleResponse, error) {
	requestURL := helper.ConstructRequestURL(backupclient.ClusterApiVersionAndGroup, fullName.ClusterName, dataProtectionSchedulePath, fullName.Name)
	queryParams := url.Values{}

	queryParams.Add(backupclient.ManagementClusterNameQueryParam, fullName.ManagementClusterName)
	queryParams.Add(backupclient.ProvisionerNameQueryParam, fullName.ProvisionerName)
	requestURL = requestURL.AppendQueryParams(queryParams)

	resp := &clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleResponse{}
	err := c.Get(requestURL.String(), resp)

	return resp, err
}

/*
BackupScheduleResourceServiceList lists backup schedules.
*/
func (c *Client) BackupScheduleResourceServiceList(request *clusterbackupschedulemodels.ListBackupSchedulesRequest) (*clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleListSchedulesResponse, error) {
	resp := &clusterbackupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataProtectionScheduleListSchedulesResponse{}

	if request.SearchScope == nil || request.SearchScope.ClusterName == "" || request.SearchScope.ManagementClusterName == "" || request.SearchScope.ProvisionerName == "" {
		return nil, errors.New("scope must be set with management cluster name, provisioner name and cluster name.")
	}

	requestURL := helper.ConstructRequestURL(backupclient.ClusterApiVersionAndGroup, request.SearchScope.ClusterName, dataProtectionSchedulePath)
	queryParams := url.Values{}
	queryParams.Add(managementClusterNameSearchQueryParam, request.SearchScope.ManagementClusterName)
	queryParams.Add(provisionerNameSearchQueryParam, request.SearchScope.ProvisionerName)

	if request.SearchScope.Name != "" {
		queryParams.Add(scheduleNameSearchQueryParam, request.SearchScope.Name)
	}

	if len(queryParams) > 0 {
		requestURL = requestURL.AppendQueryParams(queryParams)
	}

	err := c.Get(requestURL.String(), resp)

	return resp, err
}
