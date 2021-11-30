// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetValidatorDetailResultSelfDelegation get validator detail result self delegation
//
// swagger:model getValidatorDetailResult.selfDelegation
type GetValidatorDetailResultSelfDelegation struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// weight
	// Required: true
	Weight *string `json:"weight"`
}

// Validate validates this get validator detail result self delegation
func (m *GetValidatorDetailResultSelfDelegation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetValidatorDetailResultSelfDelegation) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GetValidatorDetailResultSelfDelegation) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("weight", "body", m.Weight); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get validator detail result self delegation based on context it is used
func (m *GetValidatorDetailResultSelfDelegation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetValidatorDetailResultSelfDelegation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetValidatorDetailResultSelfDelegation) UnmarshalBinary(b []byte) error {
	var res GetValidatorDetailResultSelfDelegation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
