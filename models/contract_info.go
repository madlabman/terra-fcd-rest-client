// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContractInfo contract info
//
// swagger:model ContractInfo
type ContractInfo struct {

	// bech32 encoded address
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	Address string `json:"address,omitempty"`

	// bech32 encoded address
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	Admin string `json:"admin,omitempty"`

	// code id
	CodeID string `json:"code_id,omitempty"`

	// bech32 encoded address
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	Creator string `json:"creator,omitempty"`

	// init msg
	InitMsg string `json:"init_msg,omitempty"`
}

// Validate validates this contract info
func (m *ContractInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this contract info based on context it is used
func (m *ContractInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContractInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractInfo) UnmarshalBinary(b []byte) error {
	var res ContractInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
