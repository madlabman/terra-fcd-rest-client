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

// GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrReader is a Reader for the GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddr structure.
type GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK creates a GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK with default headers values
func NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK() *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK {
	return &GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK{}
}

/* GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK describes a response with status code 200, with default header values.

OK
*/
type GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK struct {
	Payload *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/validators/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK  %+v", 200, o.Payload)
}
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK) GetPayload() *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody {
	return o.Payload
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrBadRequest creates a GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrBadRequest with default headers values
func NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrBadRequest() *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrBadRequest {
	return &GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrBadRequest{}
}

/* GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrBadRequest describes a response with status code 400, with default header values.

Invalid delegator address or validator address
*/
type GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrBadRequest struct {
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrBadRequest) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/validators/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrBadRequest ", 400)
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrInternalServerError creates a GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrInternalServerError with default headers values
func NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrInternalServerError() *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrInternalServerError {
	return &GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrInternalServerError{}
}

/* GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrInternalServerError struct {
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/validators/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrInternalServerError ", 500)
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody get staking delegators delegator addr validators validator addr o k body
swagger:model GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody
*/
type GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody struct {

	// bond height
	// Example: 0
	BondHeight string `json:"bond_height,omitempty"`

	// bond intra tx counter
	// Example: 0
	BondIntraTxCounter int64 `json:"bond_intra_tx_counter,omitempty"`

	// commission
	Commission *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyCommission `json:"commission,omitempty"`

	// consensus pubkey
	ConsensusPubkey *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyConsensusPubkey `json:"consensus_pubkey,omitempty"`

	// delegator shares
	DelegatorShares string `json:"delegator_shares,omitempty"`

	// description
	Description *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyDescription `json:"description,omitempty"`

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

// Validate validates this get staking delegators delegator addr validators validator addr o k body
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody) validateCommission(formats strfmt.Registry) error {
	if swag.IsZero(o.Commission) { // not required
		return nil
	}

	if o.Commission != nil {
		if err := o.Commission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "commission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "commission")
			}
			return err
		}
	}

	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody) validateConsensusPubkey(formats strfmt.Registry) error {
	if swag.IsZero(o.ConsensusPubkey) { // not required
		return nil
	}

	if o.ConsensusPubkey != nil {
		if err := o.ConsensusPubkey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "consensus_pubkey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "consensus_pubkey")
			}
			return err
		}
	}

	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if o.Description != nil {
		if err := o.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "description")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "description")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get staking delegators delegator addr validators validator addr o k body based on the context it is used
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody) contextValidateCommission(ctx context.Context, formats strfmt.Registry) error {

	if o.Commission != nil {
		if err := o.Commission.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "commission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "commission")
			}
			return err
		}
	}

	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody) contextValidateConsensusPubkey(ctx context.Context, formats strfmt.Registry) error {

	if o.ConsensusPubkey != nil {
		if err := o.ConsensusPubkey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "consensus_pubkey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "consensus_pubkey")
			}
			return err
		}
	}

	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if o.Description != nil {
		if err := o.Description.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "description")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK" + "." + "description")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyCommission get staking delegators delegator addr validators validator addr o k body commission
swagger:model GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyCommission
*/
type GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyCommission struct {

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

// Validate validates this get staking delegators delegator addr validators validator addr o k body commission
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyCommission) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get staking delegators delegator addr validators validator addr o k body commission based on context it is used
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyCommission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyCommission) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyCommission) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyCommission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyConsensusPubkey get staking delegators delegator addr validators validator addr o k body consensus pubkey
swagger:model GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyConsensusPubkey
*/
type GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyConsensusPubkey struct {

	// type
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this get staking delegators delegator addr validators validator addr o k body consensus pubkey
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyConsensusPubkey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get staking delegators delegator addr validators validator addr o k body consensus pubkey based on context it is used
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyConsensusPubkey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyConsensusPubkey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyConsensusPubkey) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyConsensusPubkey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyDescription get staking delegators delegator addr validators validator addr o k body description
swagger:model GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyDescription
*/
type GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyDescription struct {

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

// Validate validates this get staking delegators delegator addr validators validator addr o k body description
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get staking delegators delegator addr validators validator addr o k body description based on context it is used
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyDescription) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyDescription) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOKBodyDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}