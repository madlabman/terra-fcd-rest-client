// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BalanceReader is a Reader for the Balance structure.
type BalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBalanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBalanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBalanceOK creates a BalanceOK with default headers values
func NewBalanceOK() *BalanceOK {
	return &BalanceOK{}
}

/* BalanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type BalanceOK struct {
	Payload *BalanceOKBody
}

func (o *BalanceOK) Error() string {
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/balances/{address}/{denom}][%d] balanceOK  %+v", 200, o.Payload)
}
func (o *BalanceOK) GetPayload() *BalanceOKBody {
	return o.Payload
}

func (o *BalanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(BalanceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBalanceDefault creates a BalanceDefault with default headers values
func NewBalanceDefault(code int) *BalanceDefault {
	return &BalanceDefault{
		_statusCode: code,
	}
}

/* BalanceDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type BalanceDefault struct {
	_statusCode int

	Payload *BalanceDefaultBody
}

// Code gets the status code for the balance default response
func (o *BalanceDefault) Code() int {
	return o._statusCode
}

func (o *BalanceDefault) Error() string {
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/balances/{address}/{denom}][%d] Balance default  %+v", o._statusCode, o.Payload)
}
func (o *BalanceDefault) GetPayload() *BalanceDefaultBody {
	return o.Payload
}

func (o *BalanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(BalanceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*BalanceDefaultBody balance default body
swagger:model BalanceDefaultBody
*/
type BalanceDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*BalanceDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this balance default body
func (o *BalanceDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BalanceDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Balance default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Balance default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this balance default body based on the context it is used
func (o *BalanceDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BalanceDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Balance default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Balance default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *BalanceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BalanceDefaultBody) UnmarshalBinary(b []byte) error {
	var res BalanceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*BalanceDefaultBodyDetailsItems0 balance default body details items0
swagger:model BalanceDefaultBodyDetailsItems0
*/
type BalanceDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this balance default body details items0
func (o *BalanceDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this balance default body details items0 based on context it is used
func (o *BalanceDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BalanceDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BalanceDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res BalanceDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*BalanceOKBody QueryBalanceResponse is the response type for the Query/Balance RPC method.
swagger:model BalanceOKBody
*/
type BalanceOKBody struct {

	// balance
	Balance *BalanceOKBodyBalance `json:"balance,omitempty"`
}

// Validate validates this balance o k body
func (o *BalanceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BalanceOKBody) validateBalance(formats strfmt.Registry) error {
	if swag.IsZero(o.Balance) { // not required
		return nil
	}

	if o.Balance != nil {
		if err := o.Balance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balanceOK" + "." + "balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("balanceOK" + "." + "balance")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this balance o k body based on the context it is used
func (o *BalanceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBalance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BalanceOKBody) contextValidateBalance(ctx context.Context, formats strfmt.Registry) error {

	if o.Balance != nil {
		if err := o.Balance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balanceOK" + "." + "balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("balanceOK" + "." + "balance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *BalanceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BalanceOKBody) UnmarshalBinary(b []byte) error {
	var res BalanceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*BalanceOKBodyBalance Coin defines a token with a denomination and an amount.
//
// NOTE: The amount field is an Int which implements the custom method
// signatures required by gogoproto.
swagger:model BalanceOKBodyBalance
*/
type BalanceOKBodyBalance struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this balance o k body balance
func (o *BalanceOKBodyBalance) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this balance o k body balance based on context it is used
func (o *BalanceOKBodyBalance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BalanceOKBodyBalance) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BalanceOKBodyBalance) UnmarshalBinary(b []byte) error {
	var res BalanceOKBodyBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
