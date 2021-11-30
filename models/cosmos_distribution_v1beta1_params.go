// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosDistributionV1beta1Params Params defines the set of params for the distribution module.
//
// swagger:model cosmos.distribution.v1beta1.Params
type CosmosDistributionV1beta1Params struct {

	// base proposer reward
	BaseProposerReward string `json:"base_proposer_reward,omitempty"`

	// bonus proposer reward
	BonusProposerReward string `json:"bonus_proposer_reward,omitempty"`

	// community tax
	CommunityTax string `json:"community_tax,omitempty"`

	// withdraw addr enabled
	WithdrawAddrEnabled bool `json:"withdraw_addr_enabled,omitempty"`
}

// Validate validates this cosmos distribution v1beta1 params
func (m *CosmosDistributionV1beta1Params) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos distribution v1beta1 params based on context it is used
func (m *CosmosDistributionV1beta1Params) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosDistributionV1beta1Params) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosDistributionV1beta1Params) UnmarshalBinary(b []byte) error {
	var res CosmosDistributionV1beta1Params
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
