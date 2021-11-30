// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TreasuryParams treasury params
//
// swagger:model TreasuryParams
type TreasuryParams struct {

	// mining increment
	// Example: 1.07
	MiningIncrement float32 `json:"mining_increment,omitempty"`

	// reward policy
	RewardPolicy *TreasuryParamsRewardPolicy `json:"reward_policy,omitempty"`

	// 67%
	// Example: 0.67
	SeigniorageBurdenTarget float32 `json:"seigniorage_burden_target,omitempty"`

	// tax policy
	TaxPolicy *TreasuryParamsTaxPolicy `json:"tax_policy,omitempty"`

	// window long
	// Example: 52
	WindowLong int64 `json:"window_long,omitempty"`

	// window probation
	// Example: 12
	WindowProbation int64 `json:"window_probation,omitempty"`

	// window short
	// Example: 4
	WindowShort int64 `json:"window_short,omitempty"`
}

// Validate validates this treasury params
func (m *TreasuryParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRewardPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreasuryParams) validateRewardPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.RewardPolicy) { // not required
		return nil
	}

	if m.RewardPolicy != nil {
		if err := m.RewardPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reward_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reward_policy")
			}
			return err
		}
	}

	return nil
}

func (m *TreasuryParams) validateTaxPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxPolicy) { // not required
		return nil
	}

	if m.TaxPolicy != nil {
		if err := m.TaxPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tax_policy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this treasury params based on the context it is used
func (m *TreasuryParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRewardPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreasuryParams) contextValidateRewardPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.RewardPolicy != nil {
		if err := m.RewardPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reward_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reward_policy")
			}
			return err
		}
	}

	return nil
}

func (m *TreasuryParams) contextValidateTaxPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxPolicy != nil {
		if err := m.TaxPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tax_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TreasuryParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreasuryParams) UnmarshalBinary(b []byte) error {
	var res TreasuryParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TreasuryParamsRewardPolicy treasury params reward policy
//
// swagger:model TreasuryParamsRewardPolicy
type TreasuryParamsRewardPolicy struct {

	// cap
	Cap *TreasuryParamsRewardPolicyCap `json:"cap,omitempty"`

	// 0.025%
	// Example: 0.00025
	ChangeMax float32 `json:"change_max,omitempty"`

	// 1%
	// Example: 0.01
	RateMax float32 `json:"rate_max,omitempty"`

	// 0.05%
	// Example: 0.0005
	RateMin float32 `json:"rate_min,omitempty"`
}

// Validate validates this treasury params reward policy
func (m *TreasuryParamsRewardPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreasuryParamsRewardPolicy) validateCap(formats strfmt.Registry) error {
	if swag.IsZero(m.Cap) { // not required
		return nil
	}

	if m.Cap != nil {
		if err := m.Cap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reward_policy" + "." + "cap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reward_policy" + "." + "cap")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this treasury params reward policy based on the context it is used
func (m *TreasuryParamsRewardPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreasuryParamsRewardPolicy) contextValidateCap(ctx context.Context, formats strfmt.Registry) error {

	if m.Cap != nil {
		if err := m.Cap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reward_policy" + "." + "cap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reward_policy" + "." + "cap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TreasuryParamsRewardPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreasuryParamsRewardPolicy) UnmarshalBinary(b []byte) error {
	var res TreasuryParamsRewardPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TreasuryParamsRewardPolicyCap treasury params reward policy cap
//
// swagger:model TreasuryParamsRewardPolicyCap
type TreasuryParamsRewardPolicyCap struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this treasury params reward policy cap
func (m *TreasuryParamsRewardPolicyCap) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this treasury params reward policy cap based on context it is used
func (m *TreasuryParamsRewardPolicyCap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TreasuryParamsRewardPolicyCap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreasuryParamsRewardPolicyCap) UnmarshalBinary(b []byte) error {
	var res TreasuryParamsRewardPolicyCap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TreasuryParamsTaxPolicy treasury params tax policy
//
// swagger:model TreasuryParamsTaxPolicy
type TreasuryParamsTaxPolicy struct {

	// cap
	Cap *TreasuryParamsTaxPolicyCap `json:"cap,omitempty"`

	// 0.025%
	// Example: 0.00025
	ChangeMax float32 `json:"change_max,omitempty"`

	// 1%
	// Example: 0.01
	RateMax float32 `json:"rate_max,omitempty"`

	// 0.05%
	// Example: 0.0005
	RateMin float32 `json:"rate_min,omitempty"`
}

// Validate validates this treasury params tax policy
func (m *TreasuryParamsTaxPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreasuryParamsTaxPolicy) validateCap(formats strfmt.Registry) error {
	if swag.IsZero(m.Cap) { // not required
		return nil
	}

	if m.Cap != nil {
		if err := m.Cap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_policy" + "." + "cap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tax_policy" + "." + "cap")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this treasury params tax policy based on the context it is used
func (m *TreasuryParamsTaxPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreasuryParamsTaxPolicy) contextValidateCap(ctx context.Context, formats strfmt.Registry) error {

	if m.Cap != nil {
		if err := m.Cap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_policy" + "." + "cap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tax_policy" + "." + "cap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TreasuryParamsTaxPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreasuryParamsTaxPolicy) UnmarshalBinary(b []byte) error {
	var res TreasuryParamsTaxPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TreasuryParamsTaxPolicyCap treasury params tax policy cap
//
// swagger:model TreasuryParamsTaxPolicyCap
type TreasuryParamsTaxPolicyCap struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this treasury params tax policy cap
func (m *TreasuryParamsTaxPolicyCap) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this treasury params tax policy cap based on context it is used
func (m *TreasuryParamsTaxPolicyCap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TreasuryParamsTaxPolicyCap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreasuryParamsTaxPolicyCap) UnmarshalBinary(b []byte) error {
	var res TreasuryParamsTaxPolicyCap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}