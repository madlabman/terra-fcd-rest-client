// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TerraTreasuryV1beta1QueryTaxCapsResponse QueryTaxCapsResponse is response type for the
// Query/TaxCaps RPC method.
//
// swagger:model terra.treasury.v1beta1.QueryTaxCapsResponse
type TerraTreasuryV1beta1QueryTaxCapsResponse struct {

	// tax caps
	TaxCaps []*TerraTreasuryV1beta1QueryTaxCapsResponseTaxCapsItems0 `json:"tax_caps"`
}

// Validate validates this terra treasury v1beta1 query tax caps response
func (m *TerraTreasuryV1beta1QueryTaxCapsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaxCaps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraTreasuryV1beta1QueryTaxCapsResponse) validateTaxCaps(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxCaps) { // not required
		return nil
	}

	for i := 0; i < len(m.TaxCaps); i++ {
		if swag.IsZero(m.TaxCaps[i]) { // not required
			continue
		}

		if m.TaxCaps[i] != nil {
			if err := m.TaxCaps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tax_caps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tax_caps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this terra treasury v1beta1 query tax caps response based on the context it is used
func (m *TerraTreasuryV1beta1QueryTaxCapsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaxCaps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraTreasuryV1beta1QueryTaxCapsResponse) contextValidateTaxCaps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TaxCaps); i++ {

		if m.TaxCaps[i] != nil {
			if err := m.TaxCaps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tax_caps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tax_caps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerraTreasuryV1beta1QueryTaxCapsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraTreasuryV1beta1QueryTaxCapsResponse) UnmarshalBinary(b []byte) error {
	var res TerraTreasuryV1beta1QueryTaxCapsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TerraTreasuryV1beta1QueryTaxCapsResponseTaxCapsItems0 QueryTaxCapsResponseItem is response item type for the
// Query/TaxCaps RPC method.
//
// swagger:model TerraTreasuryV1beta1QueryTaxCapsResponseTaxCapsItems0
type TerraTreasuryV1beta1QueryTaxCapsResponseTaxCapsItems0 struct {

	// denom
	Denom string `json:"denom,omitempty"`

	// tax cap
	TaxCap string `json:"tax_cap,omitempty"`
}

// Validate validates this terra treasury v1beta1 query tax caps response tax caps items0
func (m *TerraTreasuryV1beta1QueryTaxCapsResponseTaxCapsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this terra treasury v1beta1 query tax caps response tax caps items0 based on context it is used
func (m *TerraTreasuryV1beta1QueryTaxCapsResponseTaxCapsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TerraTreasuryV1beta1QueryTaxCapsResponseTaxCapsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraTreasuryV1beta1QueryTaxCapsResponseTaxCapsItems0) UnmarshalBinary(b []byte) error {
	var res TerraTreasuryV1beta1QueryTaxCapsResponseTaxCapsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}