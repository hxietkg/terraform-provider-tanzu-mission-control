/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedulemodels

import (
	"github.com/go-openapi/swag"
)

// !!! NOT GENERATED BY SWAGGER !!!.

type ListBackupSchedulesSearchScope struct {
	ClusterName string `json:"clusterName,omitempty"`

	ManagementClusterName string `json:"managementClusterName,omitempty"`

	ProvisionerName string `json:"provisionerName,omitempty"`

	Name string `json:"name,omitempty"`
}

// MarshalBinary interface implementation.
func (m *ListBackupSchedulesSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *ListBackupSchedulesSearchScope) UnmarshalBinary(b []byte) error {
	var res ListBackupSchedulesSearchScope

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
