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

// GetDashboardResultCommunityPool get dashboard result community pool
//
// swagger:model getDashboardResult.communityPool
type GetDashboardResultCommunityPool struct {

	// ukrw amount
	// Required: true
	Ukrw *string `json:"ukrw"`

	// uluna amount
	// Required: true
	Uluna *string `json:"uluna"`

	// umnt amount
	// Required: true
	Umnt *string `json:"umnt"`

	// usdr amount
	// Required: true
	Usdr *string `json:"usdr"`

	// uusd amount
	// Required: true
	Uusd *string `json:"uusd"`
}

// Validate validates this get dashboard result community pool
func (m *GetDashboardResultCommunityPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUkrw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUluna(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUmnt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsdr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUusd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDashboardResultCommunityPool) validateUkrw(formats strfmt.Registry) error {

	if err := validate.Required("ukrw", "body", m.Ukrw); err != nil {
		return err
	}

	return nil
}

func (m *GetDashboardResultCommunityPool) validateUluna(formats strfmt.Registry) error {

	if err := validate.Required("uluna", "body", m.Uluna); err != nil {
		return err
	}

	return nil
}

func (m *GetDashboardResultCommunityPool) validateUmnt(formats strfmt.Registry) error {

	if err := validate.Required("umnt", "body", m.Umnt); err != nil {
		return err
	}

	return nil
}

func (m *GetDashboardResultCommunityPool) validateUsdr(formats strfmt.Registry) error {

	if err := validate.Required("usdr", "body", m.Usdr); err != nil {
		return err
	}

	return nil
}

func (m *GetDashboardResultCommunityPool) validateUusd(formats strfmt.Registry) error {

	if err := validate.Required("uusd", "body", m.Uusd); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get dashboard result community pool based on context it is used
func (m *GetDashboardResultCommunityPool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetDashboardResultCommunityPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDashboardResultCommunityPool) UnmarshalBinary(b []byte) error {
	var res GetDashboardResultCommunityPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}