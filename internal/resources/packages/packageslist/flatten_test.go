/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package listpackages

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	packageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/package/cluster"
	packagespec "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/package/spec"
)

func TestFlattenSpecForClusterScope(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageListPackagesResponse
		expected    []interface{}
	}{
		{
			description: "check for nil cluster list package response",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario cluster list package response",
			input: &packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageListPackagesResponse{
				Packages: []*packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackagePackage{
					{
						FullName: &packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageFullName{
							Name: "someName",
						},
						Spec: &packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec{
							CapacityRequirementsDescription: "someCapacityRequirementsDescription",
							Licenses: []string{
								"some1",
							},
							RepositoryName: "testrepo1",
							ReleaseNotes:   "someReleaseNotes",
							ReleasedAt:     strfmt.DateTime{},
							ValuesSchema: &packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageValuesSchema{
								Template: &packageclustermodel.K8sIoApimachineryPkgRuntimeRawExtension{
									Raw: []byte("somevalue"),
								},
							},
						},
					},
					{
						FullName: &packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageFullName{
							Name: "someName3",
						},
						Spec: &packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec{
							CapacityRequirementsDescription: "someCapacityRequirementsDescription",
							Licenses: []string{
								"some3",
							},
							ReleaseNotes:   "someReleaseNotes",
							RepositoryName: "testrepo2",
							ReleasedAt:     strfmt.DateTime{},
							ValuesSchema: &packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageValuesSchema{
								Template: &packageclustermodel.K8sIoApimachineryPkgRuntimeRawExtension{
									Raw: []byte("somevalue"),
								},
							},
						},
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					nameKey: "someName",
					SpecKey: []interface{}{
						map[string]interface{}{
							packagespec.CapacityRequirementsDescriptionKey: "someCapacityRequirementsDescription",
							packagespec.LicensesKey: []string{
								"some1",
							},
							packagespec.RepositoryNameKey: "testrepo1",
							packagespec.ReleaseNotesKey:   "someReleaseNotes",
							packagespec.ReleasedAtKey:     strfmt.DateTime{}.String(),
							packagespec.ValuesSchemaKey: []interface{}{
								map[string]interface{}{
									packagespec.TemplateKey: []interface{}{
										map[string]interface{}{
											packagespec.RawKey: "somevalue",
										},
									},
								},
							},
						},
					},
				},
				map[string]interface{}{
					nameKey: "someName3",
					SpecKey: []interface{}{
						map[string]interface{}{
							packagespec.CapacityRequirementsDescriptionKey: "someCapacityRequirementsDescription",
							packagespec.LicensesKey: []string{
								"some3",
							},
							packagespec.RepositoryNameKey: "testrepo2",
							packagespec.ReleaseNotesKey:   "someReleaseNotes",
							packagespec.ReleasedAtKey:     strfmt.DateTime{}.String(),
							packagespec.ValuesSchemaKey: []interface{}{
								map[string]interface{}{
									packagespec.TemplateKey: []interface{}{
										map[string]interface{}{
											packagespec.RawKey: "somevalue",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenSpecForClusterScope(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
