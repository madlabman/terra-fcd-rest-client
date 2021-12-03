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

// GetProposalDepositsResultDeposits get proposal deposits result deposits
//
// swagger:model getProposalDepositsResult.deposits
type GetProposalDepositsResultDeposits struct {

	// deposit
	// Required: true
	Deposit []*GetProposalDepositsResultDepositsDeposit `json:"deposit"`

	// Depositor information
	// Required: true
	Depositor []*GetProposalDepositsResultDepositsDepositor `json:"depositor"`

	// Txhash of the deposit transaction
	// Required: true
	Txhash *string `json:"txhash"`
}

// Validate validates this get proposal deposits result deposits
func (m *GetProposalDepositsResultDeposits) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeposit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepositor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxhash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalDepositsResultDeposits) validateDeposit(formats strfmt.Registry) error {

	if err := validate.Required("deposit", "body", m.Deposit); err != nil {
		return err
	}

	for i := 0; i < len(m.Deposit); i++ {
		if swag.IsZero(m.Deposit[i]) { // not required
			continue
		}

		if m.Deposit[i] != nil {
			if err := m.Deposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetProposalDepositsResultDeposits) validateDepositor(formats strfmt.Registry) error {

	if err := validate.Required("depositor", "body", m.Depositor); err != nil {
		return err
	}

	for i := 0; i < len(m.Depositor); i++ {
		if swag.IsZero(m.Depositor[i]) { // not required
			continue
		}

		if m.Depositor[i] != nil {
			if err := m.Depositor[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("depositor" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("depositor" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetProposalDepositsResultDeposits) validateTxhash(formats strfmt.Registry) error {

	if err := validate.Required("txhash", "body", m.Txhash); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get proposal deposits result deposits based on the context it is used
func (m *GetProposalDepositsResultDeposits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeposit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDepositor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalDepositsResultDeposits) contextValidateDeposit(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Deposit); i++ {

		if m.Deposit[i] != nil {
			if err := m.Deposit[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetProposalDepositsResultDeposits) contextValidateDepositor(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Depositor); i++ {

		if m.Depositor[i] != nil {
			if err := m.Depositor[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("depositor" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("depositor" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetProposalDepositsResultDeposits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProposalDepositsResultDeposits) UnmarshalBinary(b []byte) error {
	var res GetProposalDepositsResultDeposits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}