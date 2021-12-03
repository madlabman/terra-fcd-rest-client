// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosStakingV1beta1QueryPoolResponse QueryPoolResponse is response type for the Query/Pool RPC method.
//
// swagger:model cosmos.staking.v1beta1.QueryPoolResponse
type CosmosStakingV1beta1QueryPoolResponse struct {

	// pool
	Pool *CosmosStakingV1beta1QueryPoolResponsePool `json:"pool,omitempty"`
}

// Validate validates this cosmos staking v1beta1 query pool response
func (m *CosmosStakingV1beta1QueryPoolResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1QueryPoolResponse) validatePool(formats strfmt.Registry) error {
	if swag.IsZero(m.Pool) { // not required
		return nil
	}

	if m.Pool != nil {
		if err := m.Pool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pool")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos staking v1beta1 query pool response based on the context it is used
func (m *CosmosStakingV1beta1QueryPoolResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1QueryPoolResponse) contextValidatePool(ctx context.Context, formats strfmt.Registry) error {

	if m.Pool != nil {
		if err := m.Pool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pool")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryPoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryPoolResponse) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1QueryPoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosStakingV1beta1QueryPoolResponsePool pool defines the pool info.
//
// swagger:model CosmosStakingV1beta1QueryPoolResponsePool
type CosmosStakingV1beta1QueryPoolResponsePool struct {

	// bonded tokens
	BondedTokens string `json:"bonded_tokens,omitempty"`

	// not bonded tokens
	NotBondedTokens string `json:"not_bonded_tokens,omitempty"`
}

// Validate validates this cosmos staking v1beta1 query pool response pool
func (m *CosmosStakingV1beta1QueryPoolResponsePool) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos staking v1beta1 query pool response pool based on context it is used
func (m *CosmosStakingV1beta1QueryPoolResponsePool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryPoolResponsePool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryPoolResponsePool) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1QueryPoolResponsePool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}