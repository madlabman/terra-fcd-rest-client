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

// IbcCoreChannelV1QueryChannelResponse QueryChannelResponse is the response type for the Query/Channel RPC method.
// Besides the Channel end, it includes a proof and the height from which the
// proof was retrieved.
//
// swagger:model ibc.core.channel.v1.QueryChannelResponse
type IbcCoreChannelV1QueryChannelResponse struct {

	// channel
	Channel *IbcCoreChannelV1QueryChannelResponseChannel `json:"channel,omitempty"`

	// merkle proof of existence
	// Format: byte
	Proof strfmt.Base64 `json:"proof,omitempty"`

	// proof height
	ProofHeight *IbcCoreChannelV1QueryChannelResponseProofHeight `json:"proof_height,omitempty"`
}

// Validate validates this ibc core channel v1 query channel response
func (m *IbcCoreChannelV1QueryChannelResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProofHeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1QueryChannelResponse) validateChannel(formats strfmt.Registry) error {
	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	if m.Channel != nil {
		if err := m.Channel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreChannelV1QueryChannelResponse) validateProofHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.ProofHeight) { // not required
		return nil
	}

	if m.ProofHeight != nil {
		if err := m.ProofHeight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proof_height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proof_height")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ibc core channel v1 query channel response based on the context it is used
func (m *IbcCoreChannelV1QueryChannelResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChannel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProofHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1QueryChannelResponse) contextValidateChannel(ctx context.Context, formats strfmt.Registry) error {

	if m.Channel != nil {
		if err := m.Channel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreChannelV1QueryChannelResponse) contextValidateProofHeight(ctx context.Context, formats strfmt.Registry) error {

	if m.ProofHeight != nil {
		if err := m.ProofHeight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proof_height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proof_height")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelResponse) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryChannelResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreChannelV1QueryChannelResponseChannel channel associated with the request identifiers
//
// Channel defines pipeline for exactly-once packet delivery between specific
// modules on separate blockchains, which has at least one end capable of
// sending packets and one end capable of receiving packets.
//
// swagger:model IbcCoreChannelV1QueryChannelResponseChannel
type IbcCoreChannelV1QueryChannelResponseChannel struct {

	// list of connection identifiers, in order, along which packets sent on
	// this channel will travel
	ConnectionHops []string `json:"connection_hops"`

	// counterparty
	Counterparty *IbcCoreChannelV1QueryChannelResponseChannelCounterparty `json:"counterparty,omitempty"`

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

// Validate validates this ibc core channel v1 query channel response channel
func (m *IbcCoreChannelV1QueryChannelResponseChannel) Validate(formats strfmt.Registry) error {
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

func (m *IbcCoreChannelV1QueryChannelResponseChannel) validateCounterparty(formats strfmt.Registry) error {
	if swag.IsZero(m.Counterparty) { // not required
		return nil
	}

	if m.Counterparty != nil {
		if err := m.Counterparty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel" + "." + "counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channel" + "." + "counterparty")
			}
			return err
		}
	}

	return nil
}

var ibcCoreChannelV1QueryChannelResponseChannelTypeOrderingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ORDER_NONE_UNSPECIFIED","ORDER_UNORDERED","ORDER_ORDERED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ibcCoreChannelV1QueryChannelResponseChannelTypeOrderingPropEnum = append(ibcCoreChannelV1QueryChannelResponseChannelTypeOrderingPropEnum, v)
	}
}

const (

	// IbcCoreChannelV1QueryChannelResponseChannelOrderingORDERNONEUNSPECIFIED captures enum value "ORDER_NONE_UNSPECIFIED"
	IbcCoreChannelV1QueryChannelResponseChannelOrderingORDERNONEUNSPECIFIED string = "ORDER_NONE_UNSPECIFIED"

	// IbcCoreChannelV1QueryChannelResponseChannelOrderingORDERUNORDERED captures enum value "ORDER_UNORDERED"
	IbcCoreChannelV1QueryChannelResponseChannelOrderingORDERUNORDERED string = "ORDER_UNORDERED"

	// IbcCoreChannelV1QueryChannelResponseChannelOrderingORDERORDERED captures enum value "ORDER_ORDERED"
	IbcCoreChannelV1QueryChannelResponseChannelOrderingORDERORDERED string = "ORDER_ORDERED"
)

// prop value enum
func (m *IbcCoreChannelV1QueryChannelResponseChannel) validateOrderingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ibcCoreChannelV1QueryChannelResponseChannelTypeOrderingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IbcCoreChannelV1QueryChannelResponseChannel) validateOrdering(formats strfmt.Registry) error {
	if swag.IsZero(m.Ordering) { // not required
		return nil
	}

	// value enum
	if err := m.validateOrderingEnum("channel"+"."+"ordering", "body", *m.Ordering); err != nil {
		return err
	}

	return nil
}

var ibcCoreChannelV1QueryChannelResponseChannelTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATE_UNINITIALIZED_UNSPECIFIED","STATE_INIT","STATE_TRYOPEN","STATE_OPEN","STATE_CLOSED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ibcCoreChannelV1QueryChannelResponseChannelTypeStatePropEnum = append(ibcCoreChannelV1QueryChannelResponseChannelTypeStatePropEnum, v)
	}
}

const (

	// IbcCoreChannelV1QueryChannelResponseChannelStateSTATEUNINITIALIZEDUNSPECIFIED captures enum value "STATE_UNINITIALIZED_UNSPECIFIED"
	IbcCoreChannelV1QueryChannelResponseChannelStateSTATEUNINITIALIZEDUNSPECIFIED string = "STATE_UNINITIALIZED_UNSPECIFIED"

	// IbcCoreChannelV1QueryChannelResponseChannelStateSTATEINIT captures enum value "STATE_INIT"
	IbcCoreChannelV1QueryChannelResponseChannelStateSTATEINIT string = "STATE_INIT"

	// IbcCoreChannelV1QueryChannelResponseChannelStateSTATETRYOPEN captures enum value "STATE_TRYOPEN"
	IbcCoreChannelV1QueryChannelResponseChannelStateSTATETRYOPEN string = "STATE_TRYOPEN"

	// IbcCoreChannelV1QueryChannelResponseChannelStateSTATEOPEN captures enum value "STATE_OPEN"
	IbcCoreChannelV1QueryChannelResponseChannelStateSTATEOPEN string = "STATE_OPEN"

	// IbcCoreChannelV1QueryChannelResponseChannelStateSTATECLOSED captures enum value "STATE_CLOSED"
	IbcCoreChannelV1QueryChannelResponseChannelStateSTATECLOSED string = "STATE_CLOSED"
)

// prop value enum
func (m *IbcCoreChannelV1QueryChannelResponseChannel) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ibcCoreChannelV1QueryChannelResponseChannelTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IbcCoreChannelV1QueryChannelResponseChannel) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("channel"+"."+"state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ibc core channel v1 query channel response channel based on the context it is used
func (m *IbcCoreChannelV1QueryChannelResponseChannel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCounterparty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1QueryChannelResponseChannel) contextValidateCounterparty(ctx context.Context, formats strfmt.Registry) error {

	if m.Counterparty != nil {
		if err := m.Counterparty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel" + "." + "counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channel" + "." + "counterparty")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelResponseChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelResponseChannel) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryChannelResponseChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreChannelV1QueryChannelResponseChannelCounterparty counterparty channel end
//
// swagger:model IbcCoreChannelV1QueryChannelResponseChannelCounterparty
type IbcCoreChannelV1QueryChannelResponseChannelCounterparty struct {

	// channel end on the counterparty chain
	ChannelID string `json:"channel_id,omitempty"`

	// port on the counterparty chain which owns the other end of the channel.
	PortID string `json:"port_id,omitempty"`
}

// Validate validates this ibc core channel v1 query channel response channel counterparty
func (m *IbcCoreChannelV1QueryChannelResponseChannelCounterparty) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core channel v1 query channel response channel counterparty based on context it is used
func (m *IbcCoreChannelV1QueryChannelResponseChannelCounterparty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelResponseChannelCounterparty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelResponseChannelCounterparty) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryChannelResponseChannelCounterparty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreChannelV1QueryChannelResponseProofHeight height at which the proof was retrieved
//
// Normally the RevisionHeight is incremented at each height while keeping
// RevisionNumber the same. However some consensus algorithms may choose to
// reset the height in certain conditions e.g. hard forks, state-machine
// breaking changes In these cases, the RevisionNumber is incremented so that
// height continues to be monitonically increasing even as the RevisionHeight
// gets reset
//
// swagger:model IbcCoreChannelV1QueryChannelResponseProofHeight
type IbcCoreChannelV1QueryChannelResponseProofHeight struct {

	// the height within the given revision
	RevisionHeight string `json:"revision_height,omitempty"`

	// the revision that the client is currently on
	RevisionNumber string `json:"revision_number,omitempty"`
}

// Validate validates this ibc core channel v1 query channel response proof height
func (m *IbcCoreChannelV1QueryChannelResponseProofHeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core channel v1 query channel response proof height based on context it is used
func (m *IbcCoreChannelV1QueryChannelResponseProofHeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelResponseProofHeight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelResponseProofHeight) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryChannelResponseProofHeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
