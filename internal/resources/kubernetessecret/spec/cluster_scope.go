/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	secretmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
)

func ConstructSpecForClusterScope(d *schema.ResourceData) (spec *secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec) {
	spec = &secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec{}

	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec
	}

	data := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	spec.SecretType = secretmodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEDOCKERCONFIGJSON)

	if v, ok := specData[DockerConfigjsonKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			auth := v1[0].(map[string]interface{})

			var serverURL, username, password string

			if v, ok := auth[ImageRegistryURLKey]; ok {
				serverURL = v.(string)
			}

			if v, ok := auth[UsernameKey]; ok {
				username = v.(string)
			}

			if v, ok := auth[PasswordKey]; ok {
				password = v.(string)
			}

			secretSpecData, err := GetEncodedSpecData(serverURL, username, password)
			if err != nil {
				return spec
			}

			spec.Data = map[string]strfmt.Base64{
				DockerconfigKey: secretSpecData,
			}
		}
	}

	return spec
}

func FlattenSpecForClusterScope(spec *secretmodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec, pass string) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	secretData, ok := spec.Data[DockerconfigKey]
	if !ok {
		return data
	}

	dockerConfigJSON, err := getDecodedSpecData(secretData)
	if err != nil {
		return data
	}

	var dockerConfigJSONData = make(map[string]interface{})

	for serverURL, creds := range dockerConfigJSON.Auths {
		for attribute, value := range creds {
			dockerConfigJSONData[ImageRegistryURLKey] = serverURL

			if attribute == UsernameKey {
				stringValue, ok := value.(string)
				if !ok {
					return data
				}

				dockerConfigJSONData[UsernameKey] = stringValue
			}
		}
	}

	dockerConfigJSONData[PasswordKey] = pass

	flattenSpecData[DockerConfigjsonKey] = []interface{}{dockerConfigJSONData}

	return []interface{}{flattenSpecData}
}

func GetEncodedSpecData(serverURL, username, password string) (strfmt.Base64, error) {
	var secretspecdata strfmt.Base64

	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	err := secretspecdata.Scan(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"auths":{"%s":{"username":"%s","password":"%s","auth":"%s"}}}`, serverURL, username, password, auth))))
	if err != nil {
		return nil, err
	}

	return secretspecdata, nil
}

func getDecodedSpecData(data strfmt.Base64) (*dockerConfigJSON, error) {
	rawData, err := base64.StdEncoding.DecodeString(data.String())
	if err != nil {
		return nil, err
	}

	dockerConfigJSON := &dockerConfigJSON{}

	err = json.Unmarshal(rawData, dockerConfigJSON)
	if err != nil {
		return nil, err
	}

	return dockerConfigJSON, nil
}
