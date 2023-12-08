/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package standalonebackup

import (
	"net/url"

	backupclient "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/backup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	backupcommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/common"
	backupsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backup/standalone"
)

const (
	// API Paths.
	dataProtectionBackupsPath = "dataprotection/backups"
)

// New creates a new backup resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for backup resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	BackupResourceServiceCreate(request *backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData) (*backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData, error)

	BackupResourceServiceDelete(fn *backupcommon.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName) error

	BackupResourceServiceGet(fn *backupcommon.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName) (*backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData, error)
}

/*
BackupResourceServiceCreate creates a backup.
*/
func (c *Client) BackupResourceServiceCreate(request *backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData) (*backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData, error) {
	response := &backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData{}
	requestURL := helper.ConstructRequestURL(backupclient.ClusterApiVersionAndGroup, request.Backup.FullName.ClusterName, dataProtectionBackupsPath).String()
	err := c.Create(requestURL, request, response)

	return response, err
}

/*
BackupResourceServiceDelete deletes a backup schedule.
*/
func (c *Client) BackupResourceServiceDelete(fullName *backupcommon.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName) error {
	requestURL := helper.ConstructRequestURL(backupclient.ClusterApiVersionAndGroup, fullName.ClusterName, dataProtectionBackupsPath, fullName.Name)
	queryParams := url.Values{}

	queryParams.Add(backupclient.ManagementClusterNameQueryParam, fullName.ManagementClusterName)
	queryParams.Add(backupclient.ProvisionerNameQueryParam, fullName.ProvisionerName)

	requestURL = requestURL.AppendQueryParams(queryParams)

	return c.Delete(requestURL.String())
}

/*
BackupResourceServiceGet gets a backup.
*/
func (c *Client) BackupResourceServiceGet(fullName *backupcommon.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupFullName) (*backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData, error) {
	requestURL := helper.ConstructRequestURL(backupclient.ClusterApiVersionAndGroup, fullName.ClusterName, dataProtectionBackupsPath, fullName.Name)
	queryParams := url.Values{}

	queryParams.Add(backupclient.ManagementClusterNameQueryParam, fullName.ManagementClusterName)
	queryParams.Add(backupclient.ProvisionerNameQueryParam, fullName.ProvisionerName)

	requestURL = requestURL.AppendQueryParams(queryParams)
	resp := &backupsmodels.VmwareTanzuManageV1alpha1ClusterDataProtectionBackupData{}
	err := c.Get(requestURL.String(), resp)

	return resp, err
}
