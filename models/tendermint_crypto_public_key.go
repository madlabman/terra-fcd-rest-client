// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TendermintCryptoPublicKey PublicKey defines the keys available for use with Tendermint Validators
//
// swagger:model tendermint.crypto.PublicKey
type TendermintCryptoPublicKey struct {

	// ed25519
	// Format: byte
	Ed25519 strfmt.Base64 `json:"ed25519,omitempty"`

	// secp256k1
	// Format: byte
	Secp256k1 strfmt.Base64 `json:"secp256k1,omitempty"`
}

// Validate validates this tendermint crypto public key
func (m *TendermintCryptoPublicKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tendermint crypto public key based on context it is used
func (m *TendermintCryptoPublicKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TendermintCryptoPublicKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintCryptoPublicKey) UnmarshalBinary(b []byte) error {
	var res TendermintCryptoPublicKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
