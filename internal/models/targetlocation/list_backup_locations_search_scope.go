// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package targetlocationmodels

import "github.com/go-openapi/swag"

// !!! NOT GENERATED BY SWAGGER !!!.

type ListBackupLocationsSearchScope struct {
	ProviderName string `json:"providerName,omitempty"`

	Name string `json:"name,omitempty"`

	CredentialName string `json:"credentialName,omitempty"`

	AssignedGroupName string `json:"assignedGroupName,omitempty"`

	ClusterName string `json:"clusterName,omitempty"`

	ManagementClusterName string `json:"managementClusterName,omitempty"`

	ProvisionerName string `json:"provisionerName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *ListBackupLocationsSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *ListBackupLocationsSearchScope) UnmarshalBinary(b []byte) error {
	var res ListBackupLocationsSearchScope

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
