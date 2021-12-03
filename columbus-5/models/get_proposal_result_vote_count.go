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

// GetProposalResultVoteCount get proposal result vote count
//
// swagger:model getProposalResult.vote.count
type GetProposalResultVoteCount struct {

	// vote count
	// Required: true
	Abstain *float64 `json:"Abstain"`

	// vote count
	// Required: true
	No *float64 `json:"No"`

	// vote count
	// Required: true
	NoWithVeto *float64 `json:"NoWithVeto"`

	// vote count
	// Required: true
	Yes *float64 `json:"Yes"`
}

// Validate validates this get proposal result vote count
func (m *GetProposalResultVoteCount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbstain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoWithVeto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalResultVoteCount) validateAbstain(formats strfmt.Registry) error {

	if err := validate.Required("Abstain", "body", m.Abstain); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalResultVoteCount) validateNo(formats strfmt.Registry) error {

	if err := validate.Required("No", "body", m.No); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalResultVoteCount) validateNoWithVeto(formats strfmt.Registry) error {

	if err := validate.Required("NoWithVeto", "body", m.NoWithVeto); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalResultVoteCount) validateYes(formats strfmt.Registry) error {

	if err := validate.Required("Yes", "body", m.Yes); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get proposal result vote count based on context it is used
func (m *GetProposalResultVoteCount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetProposalResultVoteCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProposalResultVoteCount) UnmarshalBinary(b []byte) error {
	var res GetProposalResultVoteCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}