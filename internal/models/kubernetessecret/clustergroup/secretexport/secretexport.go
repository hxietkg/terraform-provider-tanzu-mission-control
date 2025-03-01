// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustergroupsecretexport

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportSecretExport Represents Tanzu Secret Export.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.secretexport.SecretExport
type VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportSecretExport struct {

	// Full name for the Secret Export.
	FullName *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportFullName `json:"fullName,omitempty"`

	// Metadata for the Secret Export object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Status for the Secret Export.
	Status *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportSecretExport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportSecretExport) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportSecretExport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
