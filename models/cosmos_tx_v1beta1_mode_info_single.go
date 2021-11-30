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

// CosmosTxV1beta1ModeInfoSingle Single is the mode info for a single signer. It is structured as a message
// to allow for additional fields such as locale for SIGN_MODE_TEXTUAL in the
// future
//
// swagger:model cosmos.tx.v1beta1.ModeInfo.Single
type CosmosTxV1beta1ModeInfoSingle struct {

	// mode is the signing mode of the single signer
	//
	// SignMode represents a signing mode with its own security guarantees.
	//
	//  - SIGN_MODE_UNSPECIFIED: SIGN_MODE_UNSPECIFIED specifies an unknown signing mode and will be
	// rejected
	//  - SIGN_MODE_DIRECT: SIGN_MODE_DIRECT specifies a signing mode which uses SignDoc and is
	// verified with raw bytes from Tx
	//  - SIGN_MODE_TEXTUAL: SIGN_MODE_TEXTUAL is a future signing mode that will verify some
	// human-readable textual representation on top of the binary representation
	// from SIGN_MODE_DIRECT
	//  - SIGN_MODE_LEGACY_AMINO_JSON: SIGN_MODE_LEGACY_AMINO_JSON is a backwards compatibility mode which uses
	// Amino JSON and will be removed in the future
	// Enum: [SIGN_MODE_UNSPECIFIED SIGN_MODE_DIRECT SIGN_MODE_TEXTUAL SIGN_MODE_LEGACY_AMINO_JSON]
	Mode *string `json:"mode,omitempty"`
}

// Validate validates this cosmos tx v1beta1 mode info single
func (m *CosmosTxV1beta1ModeInfoSingle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cosmosTxV1beta1ModeInfoSingleTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SIGN_MODE_UNSPECIFIED","SIGN_MODE_DIRECT","SIGN_MODE_TEXTUAL","SIGN_MODE_LEGACY_AMINO_JSON"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cosmosTxV1beta1ModeInfoSingleTypeModePropEnum = append(cosmosTxV1beta1ModeInfoSingleTypeModePropEnum, v)
	}
}

const (

	// CosmosTxV1beta1ModeInfoSingleModeSIGNMODEUNSPECIFIED captures enum value "SIGN_MODE_UNSPECIFIED"
	CosmosTxV1beta1ModeInfoSingleModeSIGNMODEUNSPECIFIED string = "SIGN_MODE_UNSPECIFIED"

	// CosmosTxV1beta1ModeInfoSingleModeSIGNMODEDIRECT captures enum value "SIGN_MODE_DIRECT"
	CosmosTxV1beta1ModeInfoSingleModeSIGNMODEDIRECT string = "SIGN_MODE_DIRECT"

	// CosmosTxV1beta1ModeInfoSingleModeSIGNMODETEXTUAL captures enum value "SIGN_MODE_TEXTUAL"
	CosmosTxV1beta1ModeInfoSingleModeSIGNMODETEXTUAL string = "SIGN_MODE_TEXTUAL"

	// CosmosTxV1beta1ModeInfoSingleModeSIGNMODELEGACYAMINOJSON captures enum value "SIGN_MODE_LEGACY_AMINO_JSON"
	CosmosTxV1beta1ModeInfoSingleModeSIGNMODELEGACYAMINOJSON string = "SIGN_MODE_LEGACY_AMINO_JSON"
)

// prop value enum
func (m *CosmosTxV1beta1ModeInfoSingle) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cosmosTxV1beta1ModeInfoSingleTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CosmosTxV1beta1ModeInfoSingle) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cosmos tx v1beta1 mode info single based on context it is used
func (m *CosmosTxV1beta1ModeInfoSingle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosTxV1beta1ModeInfoSingle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosTxV1beta1ModeInfoSingle) UnmarshalBinary(b []byte) error {
	var res CosmosTxV1beta1ModeInfoSingle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
