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

// ExecuteGrantReq execute grant req
//
// swagger:model ExecuteGrantReq
type ExecuteGrantReq struct {

	// base req
	// Required: true
	BaseReq *ExecuteGrantReqBaseReq `json:"base_req"`

	// msgs
	// Required: true
	Msgs []string `json:"msgs"`
}

// Validate validates this execute grant req
func (m *ExecuteGrantReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsgs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecuteGrantReq) validateBaseReq(formats strfmt.Registry) error {

	if err := validate.Required("base_req", "body", m.BaseReq); err != nil {
		return err
	}

	if m.BaseReq != nil {
		if err := m.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("base_req")
			}
			return err
		}
	}

	return nil
}

func (m *ExecuteGrantReq) validateMsgs(formats strfmt.Registry) error {

	if err := validate.Required("msgs", "body", m.Msgs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this execute grant req based on the context it is used
func (m *ExecuteGrantReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseReq(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecuteGrantReq) contextValidateBaseReq(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseReq != nil {
		if err := m.BaseReq.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("base_req")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExecuteGrantReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecuteGrantReq) UnmarshalBinary(b []byte) error {
	var res ExecuteGrantReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExecuteGrantReqBaseReq execute grant req base req
//
// swagger:model ExecuteGrantReqBaseReq
type ExecuteGrantReqBaseReq struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// chain id
	// Example: Columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// fees
	Fees []*ExecuteGrantReqBaseReqFeesItems0 `json:"fees"`

	// Sender address or Keybase name to generate a transaction
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	From string `json:"from,omitempty"`

	// gas
	// Example: 200000
	Gas string `json:"gas,omitempty"`

	// gas adjustment
	// Example: 1.2
	GasAdjustment string `json:"gas_adjustment,omitempty"`

	// memo
	// Example: Sent via Terra Station 🚀
	Memo string `json:"memo,omitempty"`

	// sequence
	// Example: 1
	Sequence string `json:"sequence,omitempty"`

	// Estimate gas for a transaction (cannot be used in conjunction with generate_only)
	// Example: false
	Simulate bool `json:"simulate,omitempty"`
}

// Validate validates this execute grant req base req
func (m *ExecuteGrantReqBaseReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecuteGrantReqBaseReq) validateFees(formats strfmt.Registry) error {
	if swag.IsZero(m.Fees) { // not required
		return nil
	}

	for i := 0; i < len(m.Fees); i++ {
		if swag.IsZero(m.Fees[i]) { // not required
			continue
		}

		if m.Fees[i] != nil {
			if err := m.Fees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this execute grant req base req based on the context it is used
func (m *ExecuteGrantReqBaseReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecuteGrantReqBaseReq) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fees); i++ {

		if m.Fees[i] != nil {
			if err := m.Fees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExecuteGrantReqBaseReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecuteGrantReqBaseReq) UnmarshalBinary(b []byte) error {
	var res ExecuteGrantReqBaseReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExecuteGrantReqBaseReqFeesItems0 execute grant req base req fees items0
//
// swagger:model ExecuteGrantReqBaseReqFeesItems0
type ExecuteGrantReqBaseReqFeesItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this execute grant req base req fees items0
func (m *ExecuteGrantReqBaseReqFeesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this execute grant req base req fees items0 based on context it is used
func (m *ExecuteGrantReqBaseReqFeesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExecuteGrantReqBaseReqFeesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecuteGrantReqBaseReqFeesItems0) UnmarshalBinary(b []byte) error {
	var res ExecuteGrantReqBaseReqFeesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}