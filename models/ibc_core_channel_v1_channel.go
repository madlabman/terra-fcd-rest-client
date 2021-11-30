// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IbcCoreChannelV1Channel Channel defines pipeline for exactly-once packet delivery between specific
// modules on separate blockchains, which has at least one end capable of
// sending packets and one end capable of receiving packets.
//
// swagger:model ibc.core.channel.v1.Channel
type IbcCoreChannelV1Channel struct {

	// list of connection identifiers, in order, along which packets sent on
	// this channel will travel
	ConnectionHops []string `json:"connection_hops"`

	// counterparty
	Counterparty *IbcCoreChannelV1ChannelCounterparty `json:"counterparty,omitempty"`

	// whether the channel is ordered or unordered
	//
	// - ORDER_NONE_UNSPECIFIED: zero-value for channel ordering
	//  - ORDER_UNORDERED: packets can be delivered in any order, which may differ from the order in
	// which they were sent.
	//  - ORDER_ORDERED: packets are delivered exactly in the order which they were sent
	// Enum: [ORDER_NONE_UNSPECIFIED ORDER_UNORDERED ORDER_ORDERED]
	Ordering *string `json:"ordering,omitempty"`

	// current state of the channel end
	//
	// State defines if a channel is in one of the following states:
	// CLOSED, INIT, TRYOPEN, OPEN or UNINITIALIZED.
	//
	//  - STATE_UNINITIALIZED_UNSPECIFIED: Default State
	//  - STATE_INIT: A channel has just started the opening handshake.
	//  - STATE_TRYOPEN: A channel has acknowledged the handshake step on the counterparty chain.
	//  - STATE_OPEN: A channel has completed the handshake. Open channels are
	// ready to send and receive packets.
	//  - STATE_CLOSED: A channel has been closed and can no longer be used to send or receive
	// packets.
	// Enum: [STATE_UNINITIALIZED_UNSPECIFIED STATE_INIT STATE_TRYOPEN STATE_OPEN STATE_CLOSED]
	State *string `json:"state,omitempty"`

	// opaque channel version, which is agreed upon during the handshake
	Version string `json:"version,omitempty"`
}

// Validate validates this ibc core channel v1 channel
func (m *IbcCoreChannelV1Channel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounterparty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdering(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1Channel) validateCounterparty(formats strfmt.Registry) error {
	if swag.IsZero(m.Counterparty) { // not required
		return nil
	}

	if m.Counterparty != nil {
		if err := m.Counterparty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counterparty")
			}
			return err
		}
	}

	return nil
}

var ibcCoreChannelV1ChannelTypeOrderingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ORDER_NONE_UNSPECIFIED","ORDER_UNORDERED","ORDER_ORDERED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ibcCoreChannelV1ChannelTypeOrderingPropEnum = append(ibcCoreChannelV1ChannelTypeOrderingPropEnum, v)
	}
}

const (

	// IbcCoreChannelV1ChannelOrderingORDERNONEUNSPECIFIED captures enum value "ORDER_NONE_UNSPECIFIED"
	IbcCoreChannelV1ChannelOrderingORDERNONEUNSPECIFIED string = "ORDER_NONE_UNSPECIFIED"

	// IbcCoreChannelV1ChannelOrderingORDERUNORDERED captures enum value "ORDER_UNORDERED"
	IbcCoreChannelV1ChannelOrderingORDERUNORDERED string = "ORDER_UNORDERED"

	// IbcCoreChannelV1ChannelOrderingORDERORDERED captures enum value "ORDER_ORDERED"
	IbcCoreChannelV1ChannelOrderingORDERORDERED string = "ORDER_ORDERED"
)

// prop value enum
func (m *IbcCoreChannelV1Channel) validateOrderingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ibcCoreChannelV1ChannelTypeOrderingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IbcCoreChannelV1Channel) validateOrdering(formats strfmt.Registry) error {
	if swag.IsZero(m.Ordering) { // not required
		return nil
	}

	// value enum
	if err := m.validateOrderingEnum("ordering", "body", *m.Ordering); err != nil {
		return err
	}

	return nil
}

var ibcCoreChannelV1ChannelTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATE_UNINITIALIZED_UNSPECIFIED","STATE_INIT","STATE_TRYOPEN","STATE_OPEN","STATE_CLOSED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ibcCoreChannelV1ChannelTypeStatePropEnum = append(ibcCoreChannelV1ChannelTypeStatePropEnum, v)
	}
}

const (

	// IbcCoreChannelV1ChannelStateSTATEUNINITIALIZEDUNSPECIFIED captures enum value "STATE_UNINITIALIZED_UNSPECIFIED"
	IbcCoreChannelV1ChannelStateSTATEUNINITIALIZEDUNSPECIFIED string = "STATE_UNINITIALIZED_UNSPECIFIED"

	// IbcCoreChannelV1ChannelStateSTATEINIT captures enum value "STATE_INIT"
	IbcCoreChannelV1ChannelStateSTATEINIT string = "STATE_INIT"

	// IbcCoreChannelV1ChannelStateSTATETRYOPEN captures enum value "STATE_TRYOPEN"
	IbcCoreChannelV1ChannelStateSTATETRYOPEN string = "STATE_TRYOPEN"

	// IbcCoreChannelV1ChannelStateSTATEOPEN captures enum value "STATE_OPEN"
	IbcCoreChannelV1ChannelStateSTATEOPEN string = "STATE_OPEN"

	// IbcCoreChannelV1ChannelStateSTATECLOSED captures enum value "STATE_CLOSED"
	IbcCoreChannelV1ChannelStateSTATECLOSED string = "STATE_CLOSED"
)

// prop value enum
func (m *IbcCoreChannelV1Channel) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ibcCoreChannelV1ChannelTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IbcCoreChannelV1Channel) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ibc core channel v1 channel based on the context it is used
func (m *IbcCoreChannelV1Channel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCounterparty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1Channel) contextValidateCounterparty(ctx context.Context, formats strfmt.Registry) error {

	if m.Counterparty != nil {
		if err := m.Counterparty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counterparty")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1Channel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1Channel) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1Channel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreChannelV1ChannelCounterparty counterparty channel end
//
// swagger:model IbcCoreChannelV1ChannelCounterparty
type IbcCoreChannelV1ChannelCounterparty struct {

	// channel end on the counterparty chain
	ChannelID string `json:"channel_id,omitempty"`

	// port on the counterparty chain which owns the other end of the channel.
	PortID string `json:"port_id,omitempty"`
}

// Validate validates this ibc core channel v1 channel counterparty
func (m *IbcCoreChannelV1ChannelCounterparty) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core channel v1 channel counterparty based on context it is used
func (m *IbcCoreChannelV1ChannelCounterparty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1ChannelCounterparty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1ChannelCounterparty) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1ChannelCounterparty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
