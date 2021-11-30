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

// IbcCoreChannelV1QueryUnreceivedPacketsResponse QueryUnreceivedPacketsResponse is the response type for the
// Query/UnreceivedPacketCommitments RPC method
//
// swagger:model ibc.core.channel.v1.QueryUnreceivedPacketsResponse
type IbcCoreChannelV1QueryUnreceivedPacketsResponse struct {

	// height
	Height *IbcCoreChannelV1QueryUnreceivedPacketsResponseHeight `json:"height,omitempty"`

	// list of unreceived packet sequences
	Sequences []string `json:"sequences"`
}

// Validate validates this ibc core channel v1 query unreceived packets response
func (m *IbcCoreChannelV1QueryUnreceivedPacketsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1QueryUnreceivedPacketsResponse) validateHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Height) { // not required
		return nil
	}

	if m.Height != nil {
		if err := m.Height.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("height")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ibc core channel v1 query unreceived packets response based on the context it is used
func (m *IbcCoreChannelV1QueryUnreceivedPacketsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1QueryUnreceivedPacketsResponse) contextValidateHeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Height != nil {
		if err := m.Height.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("height")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryUnreceivedPacketsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryUnreceivedPacketsResponse) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryUnreceivedPacketsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreChannelV1QueryUnreceivedPacketsResponseHeight query block height
//
// Normally the RevisionHeight is incremented at each height while keeping
// RevisionNumber the same. However some consensus algorithms may choose to
// reset the height in certain conditions e.g. hard forks, state-machine
// breaking changes In these cases, the RevisionNumber is incremented so that
// height continues to be monitonically increasing even as the RevisionHeight
// gets reset
//
// swagger:model IbcCoreChannelV1QueryUnreceivedPacketsResponseHeight
type IbcCoreChannelV1QueryUnreceivedPacketsResponseHeight struct {

	// the height within the given revision
	RevisionHeight string `json:"revision_height,omitempty"`

	// the revision that the client is currently on
	RevisionNumber string `json:"revision_number,omitempty"`
}

// Validate validates this ibc core channel v1 query unreceived packets response height
func (m *IbcCoreChannelV1QueryUnreceivedPacketsResponseHeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core channel v1 query unreceived packets response height based on context it is used
func (m *IbcCoreChannelV1QueryUnreceivedPacketsResponseHeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryUnreceivedPacketsResponseHeight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryUnreceivedPacketsResponseHeight) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryUnreceivedPacketsResponseHeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}