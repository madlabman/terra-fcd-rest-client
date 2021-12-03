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
)

// CosmosTxV1beta1AuthInfo AuthInfo describes the fee and signer modes that are used to sign a
// transaction.
//
// swagger:model cosmos.tx.v1beta1.AuthInfo
type CosmosTxV1beta1AuthInfo struct {

	// fee
	Fee *CosmosTxV1beta1AuthInfoFee `json:"fee,omitempty"`

	// signer_infos defines the signing modes for the required signers. The number
	// and order of elements must match the required signers from TxBody's
	// messages. The first element is the primary signer and the one which pays
	// the fee.
	SignerInfos []*CosmosTxV1beta1SignerInfo `json:"signer_infos"`
}

// Validate validates this cosmos tx v1beta1 auth info
func (m *CosmosTxV1beta1AuthInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignerInfos(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosTxV1beta1AuthInfo) validateFee(formats strfmt.Registry) error {
	if swag.IsZero(m.Fee) { // not required
		return nil
	}

	if m.Fee != nil {
		if err := m.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fee")
			}
			return err
		}
	}

	return nil
}

func (m *CosmosTxV1beta1AuthInfo) validateSignerInfos(formats strfmt.Registry) error {
	if swag.IsZero(m.SignerInfos) { // not required
		return nil
	}

	for i := 0; i < len(m.SignerInfos); i++ {
		if swag.IsZero(m.SignerInfos[i]) { // not required
			continue
		}

		if m.SignerInfos[i] != nil {
			if err := m.SignerInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signer_infos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signer_infos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos tx v1beta1 auth info based on the context it is used
func (m *CosmosTxV1beta1AuthInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignerInfos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosTxV1beta1AuthInfo) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if m.Fee != nil {
		if err := m.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fee")
			}
			return err
		}
	}

	return nil
}

func (m *CosmosTxV1beta1AuthInfo) contextValidateSignerInfos(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SignerInfos); i++ {

		if m.SignerInfos[i] != nil {
			if err := m.SignerInfos[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signer_infos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signer_infos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosTxV1beta1AuthInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosTxV1beta1AuthInfo) UnmarshalBinary(b []byte) error {
	var res CosmosTxV1beta1AuthInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosTxV1beta1AuthInfoFee Fee is the fee and gas limit for the transaction. The first signer is the
// primary signer and the one which pays the fee. The fee can be calculated
// based on the cost of evaluating the body and doing signature verification
// of the signers. This can be estimated via simulation.
//
// swagger:model CosmosTxV1beta1AuthInfoFee
type CosmosTxV1beta1AuthInfoFee struct {

	// amount is the amount of coins to be paid as a fee
	Amount []*CosmosTxV1beta1AuthInfoFeeAmountItems0 `json:"amount"`

	// gas_limit is the maximum gas that can be used in transaction processing
	// before an out of gas error occurs
	GasLimit string `json:"gas_limit,omitempty"`

	// if set, the fee payer (either the first signer or the value of the payer field) requests that a fee grant be used
	// to pay fees instead of the fee payer's own balance. If an appropriate fee grant does not exist or the chain does
	// not support fee grants, this will fail
	Granter string `json:"granter,omitempty"`

	// if unset, the first signer is responsible for paying the fees. If set, the specified account must pay the fees.
	// the payer must be a tx signer (and thus have signed this field in AuthInfo).
	// setting this field does *not* change the ordering of required signers for the transaction.
	Payer string `json:"payer,omitempty"`
}

// Validate validates this cosmos tx v1beta1 auth info fee
func (m *CosmosTxV1beta1AuthInfoFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosTxV1beta1AuthInfoFee) validateAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	for i := 0; i < len(m.Amount); i++ {
		if swag.IsZero(m.Amount[i]) { // not required
			continue
		}

		if m.Amount[i] != nil {
			if err := m.Amount[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos tx v1beta1 auth info fee based on the context it is used
func (m *CosmosTxV1beta1AuthInfoFee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosTxV1beta1AuthInfoFee) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Amount); i++ {

		if m.Amount[i] != nil {
			if err := m.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosTxV1beta1AuthInfoFee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosTxV1beta1AuthInfoFee) UnmarshalBinary(b []byte) error {
	var res CosmosTxV1beta1AuthInfoFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosTxV1beta1AuthInfoFeeAmountItems0 Coin defines a token with a denomination and an amount.
//
// NOTE: The amount field is an Int which implements the custom method
// signatures required by gogoproto.
//
// swagger:model CosmosTxV1beta1AuthInfoFeeAmountItems0
type CosmosTxV1beta1AuthInfoFeeAmountItems0 struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this cosmos tx v1beta1 auth info fee amount items0
func (m *CosmosTxV1beta1AuthInfoFeeAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos tx v1beta1 auth info fee amount items0 based on context it is used
func (m *CosmosTxV1beta1AuthInfoFeeAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosTxV1beta1AuthInfoFeeAmountItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosTxV1beta1AuthInfoFeeAmountItems0) UnmarshalBinary(b []byte) error {
	var res CosmosTxV1beta1AuthInfoFeeAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}