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

// TerraPoolDeltaReader is a Reader for the TerraPoolDelta structure.
type TerraPoolDeltaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TerraPoolDeltaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTerraPoolDeltaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTerraPoolDeltaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTerraPoolDeltaOK creates a TerraPoolDeltaOK with default headers values
func NewTerraPoolDeltaOK() *TerraPoolDeltaOK {
	return &TerraPoolDeltaOK{}
}

/* TerraPoolDeltaOK describes a response with status code 200, with default header values.

A successful response.
*/
type TerraPoolDeltaOK struct {
	Payload *TerraPoolDeltaOKBody
}

func (o *TerraPoolDeltaOK) Error() string {
	return fmt.Sprintf("[GET /terra/market/v1beta1/terra_pool_delta][%d] terraPoolDeltaOK  %+v", 200, o.Payload)
}
func (o *TerraPoolDeltaOK) GetPayload() *TerraPoolDeltaOKBody {
	return o.Payload
}

func (o *TerraPoolDeltaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TerraPoolDeltaOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTerraPoolDeltaDefault creates a TerraPoolDeltaDefault with default headers values
func NewTerraPoolDeltaDefault(code int) *TerraPoolDeltaDefault {
	return &TerraPoolDeltaDefault{
		_statusCode: code,
	}
}

/* TerraPoolDeltaDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type TerraPoolDeltaDefault struct {
	_statusCode int

	Payload *TerraPoolDeltaDefaultBody
}

// Code gets the status code for the terra pool delta default response
func (o *TerraPoolDeltaDefault) Code() int {
	return o._statusCode
}

func (o *TerraPoolDeltaDefault) Error() string {
	return fmt.Sprintf("[GET /terra/market/v1beta1/terra_pool_delta][%d] TerraPoolDelta default  %+v", o._statusCode, o.Payload)
}
func (o *TerraPoolDeltaDefault) GetPayload() *TerraPoolDeltaDefaultBody {
	return o.Payload
}

func (o *TerraPoolDeltaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TerraPoolDeltaDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TerraPoolDeltaDefaultBody terra pool delta default body
swagger:model TerraPoolDeltaDefaultBody
*/
type TerraPoolDeltaDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*TerraPoolDeltaDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this terra pool delta default body
func (o *TerraPoolDeltaDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TerraPoolDeltaDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("TerraPoolDelta default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TerraPoolDelta default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this terra pool delta default body based on the context it is used
func (o *TerraPoolDeltaDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TerraPoolDeltaDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TerraPoolDelta default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TerraPoolDelta default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TerraPoolDeltaDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TerraPoolDeltaDefaultBody) UnmarshalBinary(b []byte) error {
	var res TerraPoolDeltaDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TerraPoolDeltaDefaultBodyDetailsItems0 terra pool delta default body details items0
swagger:model TerraPoolDeltaDefaultBodyDetailsItems0
*/
type TerraPoolDeltaDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this terra pool delta default body details items0
func (o *TerraPoolDeltaDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this terra pool delta default body details items0 based on context it is used
func (o *TerraPoolDeltaDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TerraPoolDeltaDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TerraPoolDeltaDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res TerraPoolDeltaDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TerraPoolDeltaOKBody QueryTerraPoolDeltaResponse is the response type for the Query/TerraPoolDelta RPC method.
swagger:model TerraPoolDeltaOKBody
*/
type TerraPoolDeltaOKBody struct {

	// terra_pool_delta defines the gap between the TerraPool and the TerraBasePool
	// Format: byte
	TerraPoolDelta strfmt.Base64 `json:"terra_pool_delta,omitempty"`
}

// Validate validates this terra pool delta o k body
func (o *TerraPoolDeltaOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this terra pool delta o k body based on context it is used
func (o *TerraPoolDeltaOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TerraPoolDeltaOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TerraPoolDeltaOKBody) UnmarshalBinary(b []byte) error {
	var res TerraPoolDeltaOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}