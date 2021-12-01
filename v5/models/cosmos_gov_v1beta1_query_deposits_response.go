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

// CosmosGovV1beta1QueryDepositsResponse QueryDepositsResponse is the response type for the Query/Deposits RPC method.
//
// swagger:model cosmos.gov.v1beta1.QueryDepositsResponse
type CosmosGovV1beta1QueryDepositsResponse struct {

	// deposits
	Deposits []*CosmosGovV1beta1QueryDepositsResponseDepositsItems0 `json:"deposits"`

	// pagination
	Pagination *CosmosGovV1beta1QueryDepositsResponsePagination `json:"pagination,omitempty"`
}

// Validate validates this cosmos gov v1beta1 query deposits response
func (m *CosmosGovV1beta1QueryDepositsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeposits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosGovV1beta1QueryDepositsResponse) validateDeposits(formats strfmt.Registry) error {
	if swag.IsZero(m.Deposits) { // not required
		return nil
	}

	for i := 0; i < len(m.Deposits); i++ {
		if swag.IsZero(m.Deposits[i]) { // not required
			continue
		}

		if m.Deposits[i] != nil {
			if err := m.Deposits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deposits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deposits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosGovV1beta1QueryDepositsResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos gov v1beta1 query deposits response based on the context it is used
func (m *CosmosGovV1beta1QueryDepositsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeposits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosGovV1beta1QueryDepositsResponse) contextValidateDeposits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Deposits); i++ {

		if m.Deposits[i] != nil {
			if err := m.Deposits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deposits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deposits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosGovV1beta1QueryDepositsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryDepositsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryDepositsResponse) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryDepositsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosGovV1beta1QueryDepositsResponseDepositsItems0 Deposit defines an amount deposited by an account address to an active
// proposal.
//
// swagger:model CosmosGovV1beta1QueryDepositsResponseDepositsItems0
type CosmosGovV1beta1QueryDepositsResponseDepositsItems0 struct {

	// amount
	Amount []*CosmosGovV1beta1QueryDepositsResponseDepositsItems0AmountItems0 `json:"amount"`

	// depositor
	Depositor string `json:"depositor,omitempty"`

	// proposal id
	ProposalID string `json:"proposal_id,omitempty"`
}

// Validate validates this cosmos gov v1beta1 query deposits response deposits items0
func (m *CosmosGovV1beta1QueryDepositsResponseDepositsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosGovV1beta1QueryDepositsResponseDepositsItems0) validateAmount(formats strfmt.Registry) error {
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
					return ve.ValidateName("amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos gov v1beta1 query deposits response deposits items0 based on the context it is used
func (m *CosmosGovV1beta1QueryDepositsResponseDepositsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosGovV1beta1QueryDepositsResponseDepositsItems0) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Amount); i++ {

		if m.Amount[i] != nil {
			if err := m.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryDepositsResponseDepositsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryDepositsResponseDepositsItems0) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryDepositsResponseDepositsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosGovV1beta1QueryDepositsResponseDepositsItems0AmountItems0 Coin defines a token with a denomination and an amount.
//
// NOTE: The amount field is an Int which implements the custom method
// signatures required by gogoproto.
//
// swagger:model CosmosGovV1beta1QueryDepositsResponseDepositsItems0AmountItems0
type CosmosGovV1beta1QueryDepositsResponseDepositsItems0AmountItems0 struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this cosmos gov v1beta1 query deposits response deposits items0 amount items0
func (m *CosmosGovV1beta1QueryDepositsResponseDepositsItems0AmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos gov v1beta1 query deposits response deposits items0 amount items0 based on context it is used
func (m *CosmosGovV1beta1QueryDepositsResponseDepositsItems0AmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryDepositsResponseDepositsItems0AmountItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryDepositsResponseDepositsItems0AmountItems0) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryDepositsResponseDepositsItems0AmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosGovV1beta1QueryDepositsResponsePagination pagination defines the pagination in the response.
//
// swagger:model CosmosGovV1beta1QueryDepositsResponsePagination
type CosmosGovV1beta1QueryDepositsResponsePagination struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this cosmos gov v1beta1 query deposits response pagination
func (m *CosmosGovV1beta1QueryDepositsResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos gov v1beta1 query deposits response pagination based on context it is used
func (m *CosmosGovV1beta1QueryDepositsResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryDepositsResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryDepositsResponsePagination) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryDepositsResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}