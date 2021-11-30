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

// GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrReader is a Reader for the GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddr structure.
type GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK creates a GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK with default headers values
func NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK() *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK {
	return &GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK{}
}

/* GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK describes a response with status code 200, with default header values.

OK
*/
type GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK struct {
	Payload *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/delegations/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK  %+v", 200, o.Payload)
}
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK) GetPayload() *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody {
	return o.Payload
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest creates a GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest with default headers values
func NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest() *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest {
	return &GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest{}
}

/* GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest describes a response with status code 400, with default header values.

Invalid delegator address or validator address
*/
type GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest struct {
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/delegations/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest ", 400)
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError creates a GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError with default headers values
func NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError() *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError {
	return &GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError{}
}

/* GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError struct {
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/delegations/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError ", 500)
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody get staking delegators delegator addr delegations validator addr o k body
swagger:model GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody
*/
type GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody struct {

	// balance
	Balance *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyBalance `json:"balance,omitempty"`

	// delegation
	Delegation *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyDelegation `json:"delegation,omitempty"`
}

// Validate validates this get staking delegators delegator addr delegations validator addr o k body
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDelegation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody) validateBalance(formats strfmt.Registry) error {
	if swag.IsZero(o.Balance) { // not required
		return nil
	}

	if o.Balance != nil {
		if err := o.Balance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK" + "." + "balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK" + "." + "balance")
			}
			return err
		}
	}

	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody) validateDelegation(formats strfmt.Registry) error {
	if swag.IsZero(o.Delegation) { // not required
		return nil
	}

	if o.Delegation != nil {
		if err := o.Delegation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK" + "." + "delegation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK" + "." + "delegation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get staking delegators delegator addr delegations validator addr o k body based on the context it is used
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBalance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDelegation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody) contextValidateBalance(ctx context.Context, formats strfmt.Registry) error {

	if o.Balance != nil {
		if err := o.Balance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK" + "." + "balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK" + "." + "balance")
			}
			return err
		}
	}

	return nil
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody) contextValidateDelegation(ctx context.Context, formats strfmt.Registry) error {

	if o.Delegation != nil {
		if err := o.Delegation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK" + "." + "delegation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK" + "." + "delegation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyBalance get staking delegators delegator addr delegations validator addr o k body balance
swagger:model GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyBalance
*/
type GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyBalance struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this get staking delegators delegator addr delegations validator addr o k body balance
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyBalance) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get staking delegators delegator addr delegations validator addr o k body balance based on context it is used
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyBalance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyBalance) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyBalance) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyDelegation get staking delegators delegator addr delegations validator addr o k body delegation
swagger:model GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyDelegation
*/
type GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyDelegation struct {

	// delegator address
	DelegatorAddress string `json:"delegator_address,omitempty"`

	// shares
	Shares string `json:"shares,omitempty"`

	// validator address
	ValidatorAddress string `json:"validator_address,omitempty"`
}

// Validate validates this get staking delegators delegator addr delegations validator addr o k body delegation
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyDelegation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get staking delegators delegator addr delegations validator addr o k body delegation based on context it is used
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyDelegation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyDelegation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyDelegation) UnmarshalBinary(b []byte) error {
	var res GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOKBodyDelegation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
