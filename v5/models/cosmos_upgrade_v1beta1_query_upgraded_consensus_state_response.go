// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosUpgradeV1beta1QueryUpgradedConsensusStateResponse QueryUpgradedConsensusStateResponse is the response type for the Query/UpgradedConsensusState
// RPC method.
//
// swagger:model cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse
type CosmosUpgradeV1beta1QueryUpgradedConsensusStateResponse struct {

	// upgraded consensus state
	// Format: byte
	UpgradedConsensusState strfmt.Base64 `json:"upgraded_consensus_state,omitempty"`
}

// Validate validates this cosmos upgrade v1beta1 query upgraded consensus state response
func (m *CosmosUpgradeV1beta1QueryUpgradedConsensusStateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos upgrade v1beta1 query upgraded consensus state response based on context it is used
func (m *CosmosUpgradeV1beta1QueryUpgradedConsensusStateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosUpgradeV1beta1QueryUpgradedConsensusStateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosUpgradeV1beta1QueryUpgradedConsensusStateResponse) UnmarshalBinary(b []byte) error {
	var res CosmosUpgradeV1beta1QueryUpgradedConsensusStateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}