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
	"github.com/go-openapi/validate"
)

// GetTxListResultTxsEvents get tx list result txs events
//
// swagger:model getTxListResult.txs.events
type GetTxListResultTxsEvents struct {

	// attributes
	// Required: true
	Attributes []*GetTxListResultTxsEventsAttributes `json:"attributes"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this get tx list result txs events
func (m *GetTxListResultTxsEvents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxListResultTxsEvents) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	for i := 0; i < len(m.Attributes); i++ {
		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetTxListResultTxsEvents) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get tx list result txs events based on the context it is used
func (m *GetTxListResultTxsEvents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxListResultTxsEvents) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attributes); i++ {

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTxListResultTxsEvents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTxListResultTxsEvents) UnmarshalBinary(b []byte) error {
	var res GetTxListResultTxsEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
