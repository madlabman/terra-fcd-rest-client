// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MultiSignPubKey multi sign pub key
//
// swagger:model MultiSignPubKey
type MultiSignPubKey struct {

	// pubkeys
	Pubkeys []string `json:"pubkeys"`

	// threshold
	// Example: 1
	Threshold float64 `json:"threshold,omitempty"`
}

// Validate validates this multi sign pub key
func (m *MultiSignPubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this multi sign pub key based on context it is used
func (m *MultiSignPubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MultiSignPubKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiSignPubKey) UnmarshalBinary(b []byte) error {
	var res MultiSignPubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
