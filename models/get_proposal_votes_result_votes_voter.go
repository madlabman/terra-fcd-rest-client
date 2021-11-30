// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetProposalVotesResultVotesVoter get proposal votes result votes voter
//
// swagger:model getProposalVotesResult.votes.voter
type GetProposalVotesResultVotesVoter struct {

	// account address
	// Required: true
	AccountAddress *string `json:"accountAddress"`

	// moniker
	Moniker string `json:"moniker,omitempty"`

	// operator address
	OperatorAddress string `json:"operatorAddress,omitempty"`
}

// Validate validates this get proposal votes result votes voter
func (m *GetProposalVotesResultVotesVoter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalVotesResultVotesVoter) validateAccountAddress(formats strfmt.Registry) error {

	if err := validate.Required("accountAddress", "body", m.AccountAddress); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get proposal votes result votes voter based on context it is used
func (m *GetProposalVotesResultVotesVoter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetProposalVotesResultVotesVoter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProposalVotesResultVotesVoter) UnmarshalBinary(b []byte) error {
	var res GetProposalVotesResultVotesVoter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
