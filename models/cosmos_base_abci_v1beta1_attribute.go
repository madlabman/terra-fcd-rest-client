// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosBaseAbciV1beta1Attribute Attribute defines an attribute wrapper where the key and value are
// strings instead of raw bytes.
//
// swagger:model cosmos.base.abci.v1beta1.Attribute
type CosmosBaseAbciV1beta1Attribute struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this cosmos base abci v1beta1 attribute
func (m *CosmosBaseAbciV1beta1Attribute) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos base abci v1beta1 attribute based on context it is used
func (m *CosmosBaseAbciV1beta1Attribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1Attribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1Attribute) UnmarshalBinary(b []byte) error {
	var res CosmosBaseAbciV1beta1Attribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
