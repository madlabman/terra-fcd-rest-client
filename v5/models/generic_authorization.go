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

// GenericAuthorization generic authorization
//
// swagger:model GenericAuthorization
type GenericAuthorization struct {

	// type
	// Example: msgauth/GenericAuthorization
	Type string `json:"type,omitempty"`

	// value
	Value *GenericAuthorizationValue `json:"value,omitempty"`
}

// Validate validates this generic authorization
func (m *GenericAuthorization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericAuthorization) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this generic authorization based on the context it is used
func (m *GenericAuthorization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericAuthorization) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {
		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenericAuthorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericAuthorization) UnmarshalBinary(b []byte) error {
	var res GenericAuthorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GenericAuthorizationValue generic authorization value
//
// swagger:model GenericAuthorizationValue
type GenericAuthorizationValue struct {

	// msg type
	// Example: send
	MsgType string `json:"msg_type,omitempty"`
}

// Validate validates this generic authorization value
func (m *GenericAuthorizationValue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this generic authorization value based on context it is used
func (m *GenericAuthorizationValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenericAuthorizationValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericAuthorizationValue) UnmarshalBinary(b []byte) error {
	var res GenericAuthorizationValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}