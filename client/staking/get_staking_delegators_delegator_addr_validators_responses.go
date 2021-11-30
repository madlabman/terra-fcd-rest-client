// Code generated by go-swagger; DO NOT EDIT.

package staking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetStakingDelegatorsDelegatorAddrValidatorsReader is a Reader for the GetStakingDelegatorsDelegatorAddrValidators structure.
type GetStakingDelegatorsDelegatorAddrValidatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStakingDelegatorsDelegatorAddrValidatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStakingDelegatorsDelegatorAddrValidatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStakingDelegatorsDelegatorAddrValidatorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStakingDelegatorsDelegatorAddrValidatorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStakingDelegatorsDelegatorAddrValidatorsOK creates a GetStakingDelegatorsDelegatorAddrValidatorsOK with default headers values
func NewGetStakingDelegatorsDelegatorAddrValidatorsOK() *GetStakingDelegatorsDelegatorAddrValidatorsOK {
	return &GetStakingDelegatorsDelegatorAddrValidatorsOK{}
}

/* GetStakingDelegatorsDelegatorAddrValidatorsOK describes a response with status code 200, with default header values.

OK
*/
type GetStakingDelegatorsDelegatorAddrValidatorsOK struct {
	Payload []*GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsOK) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/validators][%d] getStakingDelegatorsDelegatorAddrValidatorsOK  %+v", 200, o.Payload)
}
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOK) GetPayload() []*GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0 {
	return o.Payload
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrValidatorsBadRequest creates a GetStakingDelegatorsDelegatorAddrValidatorsBadRequest with default headers values
func NewGetStakingDelegatorsDelegatorAddrValidatorsBadRequest() *GetStakingDelegatorsDelegatorAddrValidatorsBadRequest {
	return &GetStakingDelegatorsDelegatorAddrValidatorsBadRequest{}
}

/* GetStakingDelegatorsDelegatorAddrValidatorsBadRequest describes a response with status code 400, with default header values.

Invalid delegator address
*/
type GetStakingDelegatorsDelegatorAddrValidatorsBadRequest struct {
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/validators][%d] getStakingDelegatorsDelegatorAddrValidatorsBadRequest ", 400)
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrValidatorsInternalServerError creates a GetStakingDelegatorsDelegatorAddrValidatorsInternalServerError with default headers values
func NewGetStakingDelegatorsDelegatorAddrValidatorsInternalServerError() *GetStakingDelegatorsDelegatorAddrValidatorsInternalServerError {
	return &GetStakingDelegatorsDelegatorAddrValidatorsInternalServerError{}
}

/* GetStakingDelegatorsDelegatorAddrValidatorsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetStakingDelegatorsDelegatorAddrValidatorsInternalServerError struct {
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/validators][%d] getStakingDelegatorsDelegatorAddrValidatorsInternalServerError ", 500)
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0 get staking delegators delegator addr validators o k body items0
swagger:model GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0
*/
type GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0 struct {

	// bond height
	// Example: 0
	BondHeight string `json:"bond_height,omitempty"`

	// bond intra tx counter
	// Example: 0
	BondIntraTxCounter int64 `json:"bond_intra_tx_counter,omitempty"`

	// commission
	Commission *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Commission `json:"commission,omitempty"`

	// consensus pubkey
	ConsensusPubkey *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0ConsensusPubkey `json:"consensus_pubkey,omitempty"`

	// delegator shares
	DelegatorShares string `json:"delegator_shares,omitempty"`

	// description
	Description *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Description `json:"description,omitempty"`

	// jailed
	Jailed bool `json:"jailed,omitempty"`

	// bech32 encoded address
	// Example: terravaloper1wg2mlrxdmnnkkykgqg4znky86nyrtc45q7a85l
	OperatorAddress string `json:"operator_address,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`

	// tokens
	Tokens string `json:"tokens,omitempty"`

	// unbonding height
	// Example: 0
	UnbondingHeight string `json:"unbonding_height,omitempty"`

	// unbonding time
	// Example: 1970-01-01T00:00:00Z
	UnbondingTime string `json:"unbonding_time,omitempty"`
}

// Validate validates this get staking delegators delegator addr validators o k body items0
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommission(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateConsensusPubkey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0) validateCommission(formats strfmt.Registry) error {
	if swag.IsZero(o.Commission) { // not required
		return nil
	}

	if o.Commission != nil {
		if err := o.Commission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commission")
			}
			return err
		}
	}

	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0) validateConsensusPubkey(formats strfmt.Registry) error {
	if swag.IsZero(o.ConsensusPubkey) { // not required
		return nil
	}

	if o.ConsensusPubkey != nil {
		if err := o.ConsensusPubkey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consensus_pubkey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consensus_pubkey")
			}
			return err
		}
	}

	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if o.Description != nil {
		if err := o.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get staking delegators delegator addr validators o k body items0 based on the context it is used
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCommission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateConsensusPubkey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0) contextValidateCommission(ctx context.Context, formats strfmt.Registry) error {

	if o.Commission != nil {
		if err := o.Commission.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commission")
			}
			return err
		}
	}

	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0) contextValidateConsensusPubkey(ctx context.Context, formats strfmt.Registry) error {

	if o.ConsensusPubkey != nil {
		if err := o.ConsensusPubkey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consensus_pubkey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consensus_pubkey")
			}
			return err
		}
	}

	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if o.Description != nil {
		if err := o.Description.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Commission get staking delegators delegator addr validators o k body items0 commission
swagger:model GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Commission
*/
type GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Commission struct {

	// max change rate
	// Example: 0
	MaxChangeRate string `json:"max_change_rate,omitempty"`

	// max rate
	// Example: 0
	MaxRate string `json:"max_rate,omitempty"`

	// rate
	// Example: 0
	Rate string `json:"rate,omitempty"`

	// update time
	// Example: 1970-01-01T00:00:00Z
	UpdateTime string `json:"update_time,omitempty"`
}

// Validate validates this get staking delegators delegator addr validators o k body items0 commission
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Commission) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get staking delegators delegator addr validators o k body items0 commission based on context it is used
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Commission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Commission) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Commission) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Commission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0ConsensusPubkey get staking delegators delegator addr validators o k body items0 consensus pubkey
swagger:model GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0ConsensusPubkey
*/
type GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0ConsensusPubkey struct {

	// type
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this get staking delegators delegator addr validators o k body items0 consensus pubkey
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0ConsensusPubkey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get staking delegators delegator addr validators o k body items0 consensus pubkey based on context it is used
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0ConsensusPubkey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0ConsensusPubkey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0ConsensusPubkey) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0ConsensusPubkey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Description get staking delegators delegator addr validators o k body items0 description
swagger:model GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Description
*/
type GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Description struct {

	// details
	Details string `json:"details,omitempty"`

	// identity
	Identity string `json:"identity,omitempty"`

	// moniker
	Moniker string `json:"moniker,omitempty"`

	// security contact
	SecurityContact string `json:"security_contact,omitempty"`

	// website
	Website string `json:"website,omitempty"`
}

// Validate validates this get staking delegators delegator addr validators o k body items0 description
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Description) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get staking delegators delegator addr validators o k body items0 description based on context it is used
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Description) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Description) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Description) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrValidatorsOKBodyItems0Description
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
