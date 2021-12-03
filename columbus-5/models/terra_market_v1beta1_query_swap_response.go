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

// TerraMarketV1beta1QuerySwapResponse QuerySwapResponse is the response type for the Query/Swap RPC method.
//
// swagger:model terra.market.v1beta1.QuerySwapResponse
type TerraMarketV1beta1QuerySwapResponse struct {

	// return coin
	ReturnCoin *TerraMarketV1beta1QuerySwapResponseReturnCoin `json:"return_coin,omitempty"`
}

// Validate validates this terra market v1beta1 query swap response
func (m *TerraMarketV1beta1QuerySwapResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReturnCoin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraMarketV1beta1QuerySwapResponse) validateReturnCoin(formats strfmt.Registry) error {
	if swag.IsZero(m.ReturnCoin) { // not required
		return nil
	}

	if m.ReturnCoin != nil {
		if err := m.ReturnCoin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("return_coin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("return_coin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this terra market v1beta1 query swap response based on the context it is used
func (m *TerraMarketV1beta1QuerySwapResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReturnCoin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraMarketV1beta1QuerySwapResponse) contextValidateReturnCoin(ctx context.Context, formats strfmt.Registry) error {

	if m.ReturnCoin != nil {
		if err := m.ReturnCoin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("return_coin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("return_coin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerraMarketV1beta1QuerySwapResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraMarketV1beta1QuerySwapResponse) UnmarshalBinary(b []byte) error {
	var res TerraMarketV1beta1QuerySwapResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TerraMarketV1beta1QuerySwapResponseReturnCoin return_coin defines the coin returned as a result of the swap simulation.
//
// swagger:model TerraMarketV1beta1QuerySwapResponseReturnCoin
type TerraMarketV1beta1QuerySwapResponseReturnCoin struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this terra market v1beta1 query swap response return coin
func (m *TerraMarketV1beta1QuerySwapResponseReturnCoin) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this terra market v1beta1 query swap response return coin based on context it is used
func (m *TerraMarketV1beta1QuerySwapResponseReturnCoin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TerraMarketV1beta1QuerySwapResponseReturnCoin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraMarketV1beta1QuerySwapResponseReturnCoin) UnmarshalBinary(b []byte) error {
	var res TerraMarketV1beta1QuerySwapResponseReturnCoin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}