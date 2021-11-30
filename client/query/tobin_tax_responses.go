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

// TobinTaxReader is a Reader for the TobinTax structure.
type TobinTaxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TobinTaxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTobinTaxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTobinTaxDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTobinTaxOK creates a TobinTaxOK with default headers values
func NewTobinTaxOK() *TobinTaxOK {
	return &TobinTaxOK{}
}

/* TobinTaxOK describes a response with status code 200, with default header values.

A successful response.
*/
type TobinTaxOK struct {
	Payload *TobinTaxOKBody
}

func (o *TobinTaxOK) Error() string {
	return fmt.Sprintf("[GET /terra/oracle/v1beta1/denoms/{denom}/tobin_tax][%d] tobinTaxOK  %+v", 200, o.Payload)
}
func (o *TobinTaxOK) GetPayload() *TobinTaxOKBody {
	return o.Payload
}

func (o *TobinTaxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TobinTaxOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTobinTaxDefault creates a TobinTaxDefault with default headers values
func NewTobinTaxDefault(code int) *TobinTaxDefault {
	return &TobinTaxDefault{
		_statusCode: code,
	}
}

/* TobinTaxDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type TobinTaxDefault struct {
	_statusCode int

	Payload *TobinTaxDefaultBody
}

// Code gets the status code for the tobin tax default response
func (o *TobinTaxDefault) Code() int {
	return o._statusCode
}

func (o *TobinTaxDefault) Error() string {
	return fmt.Sprintf("[GET /terra/oracle/v1beta1/denoms/{denom}/tobin_tax][%d] TobinTax default  %+v", o._statusCode, o.Payload)
}
func (o *TobinTaxDefault) GetPayload() *TobinTaxDefaultBody {
	return o.Payload
}

func (o *TobinTaxDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TobinTaxDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TobinTaxDefaultBody tobin tax default body
swagger:model TobinTaxDefaultBody
*/
type TobinTaxDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*TobinTaxDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this tobin tax default body
func (o *TobinTaxDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TobinTaxDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("TobinTax default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TobinTax default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tobin tax default body based on the context it is used
func (o *TobinTaxDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TobinTaxDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TobinTax default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TobinTax default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TobinTaxDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TobinTaxDefaultBody) UnmarshalBinary(b []byte) error {
	var res TobinTaxDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TobinTaxDefaultBodyDetailsItems0 tobin tax default body details items0
swagger:model TobinTaxDefaultBodyDetailsItems0
*/
type TobinTaxDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this tobin tax default body details items0
func (o *TobinTaxDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tobin tax default body details items0 based on context it is used
func (o *TobinTaxDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TobinTaxDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TobinTaxDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res TobinTaxDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TobinTaxOKBody QueryTobinTaxResponse is response type for the
// Query/TobinTax RPC method.
swagger:model TobinTaxOKBody
*/
type TobinTaxOKBody struct {

	// tobin_taxe defines the tobin tax of a denom
	TobinTax string `json:"tobin_tax,omitempty"`
}

// Validate validates this tobin tax o k body
func (o *TobinTaxOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tobin tax o k body based on context it is used
func (o *TobinTaxOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TobinTaxOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TobinTaxOKBody) UnmarshalBinary(b []byte) error {
	var res TobinTaxOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
