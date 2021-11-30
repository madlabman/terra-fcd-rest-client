// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IbcCoreChannelV1PacketState PacketState defines the generic type necessary to retrieve and store
// packet commitments, acknowledgements, and receipts.
// Caller is responsible for knowing the context necessary to interpret this
// state as a commitment, acknowledgement, or a receipt.
//
// swagger:model ibc.core.channel.v1.PacketState
type IbcCoreChannelV1PacketState struct {

	// channel unique identifier.
	ChannelID string `json:"channel_id,omitempty"`

	// embedded data that represents packet state.
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// channel port identifier.
	PortID string `json:"port_id,omitempty"`

	// packet sequence.
	Sequence string `json:"sequence,omitempty"`
}

// Validate validates this ibc core channel v1 packet state
func (m *IbcCoreChannelV1PacketState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core channel v1 packet state based on context it is used
func (m *IbcCoreChannelV1PacketState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1PacketState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1PacketState) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1PacketState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
