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

// TendermintTypesSignedMsgType SignedMsgType is a type of signed message in the consensus.
//
//  - SIGNED_MSG_TYPE_PREVOTE: Votes
//  - SIGNED_MSG_TYPE_PROPOSAL: Proposals
//
// swagger:model tendermint.types.SignedMsgType
type TendermintTypesSignedMsgType string

func NewTendermintTypesSignedMsgType(value TendermintTypesSignedMsgType) *TendermintTypesSignedMsgType {
	v := value
	return &v
}

const (

	// TendermintTypesSignedMsgTypeSIGNEDMSGTYPEUNKNOWN captures enum value "SIGNED_MSG_TYPE_UNKNOWN"
	TendermintTypesSignedMsgTypeSIGNEDMSGTYPEUNKNOWN TendermintTypesSignedMsgType = "SIGNED_MSG_TYPE_UNKNOWN"

	// TendermintTypesSignedMsgTypeSIGNEDMSGTYPEPREVOTE captures enum value "SIGNED_MSG_TYPE_PREVOTE"
	TendermintTypesSignedMsgTypeSIGNEDMSGTYPEPREVOTE TendermintTypesSignedMsgType = "SIGNED_MSG_TYPE_PREVOTE"

	// TendermintTypesSignedMsgTypeSIGNEDMSGTYPEPRECOMMIT captures enum value "SIGNED_MSG_TYPE_PRECOMMIT"
	TendermintTypesSignedMsgTypeSIGNEDMSGTYPEPRECOMMIT TendermintTypesSignedMsgType = "SIGNED_MSG_TYPE_PRECOMMIT"

	// TendermintTypesSignedMsgTypeSIGNEDMSGTYPEPROPOSAL captures enum value "SIGNED_MSG_TYPE_PROPOSAL"
	TendermintTypesSignedMsgTypeSIGNEDMSGTYPEPROPOSAL TendermintTypesSignedMsgType = "SIGNED_MSG_TYPE_PROPOSAL"
)

// for schema
var tendermintTypesSignedMsgTypeEnum []interface{}

func init() {
	var res []TendermintTypesSignedMsgType
	if err := json.Unmarshal([]byte(`["SIGNED_MSG_TYPE_UNKNOWN","SIGNED_MSG_TYPE_PREVOTE","SIGNED_MSG_TYPE_PRECOMMIT","SIGNED_MSG_TYPE_PROPOSAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tendermintTypesSignedMsgTypeEnum = append(tendermintTypesSignedMsgTypeEnum, v)
	}
}

func (m TendermintTypesSignedMsgType) validateTendermintTypesSignedMsgTypeEnum(path, location string, value TendermintTypesSignedMsgType) error {
	if err := validate.EnumCase(path, location, value, tendermintTypesSignedMsgTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this tendermint types signed msg type
func (m TendermintTypesSignedMsgType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTendermintTypesSignedMsgTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this tendermint types signed msg type based on context it is used
func (m TendermintTypesSignedMsgType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}