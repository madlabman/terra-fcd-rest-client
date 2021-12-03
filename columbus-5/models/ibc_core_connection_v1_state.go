// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// IbcCoreConnectionV1State State defines if a connection is in one of the following states:
// INIT, TRYOPEN, OPEN or UNINITIALIZED.
//
//  - STATE_UNINITIALIZED_UNSPECIFIED: Default State
//  - STATE_INIT: A connection end has just started the opening handshake.
//  - STATE_TRYOPEN: A connection end has acknowledged the handshake step on the counterparty
// chain.
//  - STATE_OPEN: A connection end has completed the handshake.
//
// swagger:model ibc.core.connection.v1.State
type IbcCoreConnectionV1State string

func NewIbcCoreConnectionV1State(value IbcCoreConnectionV1State) *IbcCoreConnectionV1State {
	v := value
	return &v
}

const (

	// IbcCoreConnectionV1StateSTATEUNINITIALIZEDUNSPECIFIED captures enum value "STATE_UNINITIALIZED_UNSPECIFIED"
	IbcCoreConnectionV1StateSTATEUNINITIALIZEDUNSPECIFIED IbcCoreConnectionV1State = "STATE_UNINITIALIZED_UNSPECIFIED"

	// IbcCoreConnectionV1StateSTATEINIT captures enum value "STATE_INIT"
	IbcCoreConnectionV1StateSTATEINIT IbcCoreConnectionV1State = "STATE_INIT"

	// IbcCoreConnectionV1StateSTATETRYOPEN captures enum value "STATE_TRYOPEN"
	IbcCoreConnectionV1StateSTATETRYOPEN IbcCoreConnectionV1State = "STATE_TRYOPEN"

	// IbcCoreConnectionV1StateSTATEOPEN captures enum value "STATE_OPEN"
	IbcCoreConnectionV1StateSTATEOPEN IbcCoreConnectionV1State = "STATE_OPEN"
)

// for schema
var ibcCoreConnectionV1StateEnum []interface{}

func init() {
	var res []IbcCoreConnectionV1State
	if err := json.Unmarshal([]byte(`["STATE_UNINITIALIZED_UNSPECIFIED","STATE_INIT","STATE_TRYOPEN","STATE_OPEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ibcCoreConnectionV1StateEnum = append(ibcCoreConnectionV1StateEnum, v)
	}
}

func (m IbcCoreConnectionV1State) validateIbcCoreConnectionV1StateEnum(path, location string, value IbcCoreConnectionV1State) error {
	if err := validate.EnumCase(path, location, value, ibcCoreConnectionV1StateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this ibc core connection v1 state
func (m IbcCoreConnectionV1State) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIbcCoreConnectionV1StateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this ibc core connection v1 state based on context it is used
func (m IbcCoreConnectionV1State) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}