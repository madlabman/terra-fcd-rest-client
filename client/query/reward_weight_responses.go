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

// RewardWeightReader is a Reader for the RewardWeight structure.
type RewardWeightReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RewardWeightReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRewardWeightOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRewardWeightDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRewardWeightOK creates a RewardWeightOK with default headers values
func NewRewardWeightOK() *RewardWeightOK {
	return &RewardWeightOK{}
}

/* RewardWeightOK describes a response with status code 200, with default header values.

A successful response.
*/
type RewardWeightOK struct {
	Payload *RewardWeightOKBody
}

func (o *RewardWeightOK) Error() string {
	return fmt.Sprintf("[GET /terra/treasury/v1beta1/reward_weight][%d] rewardWeightOK  %+v", 200, o.Payload)
}
func (o *RewardWeightOK) GetPayload() *RewardWeightOKBody {
	return o.Payload
}

func (o *RewardWeightOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RewardWeightOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRewardWeightDefault creates a RewardWeightDefault with default headers values
func NewRewardWeightDefault(code int) *RewardWeightDefault {
	return &RewardWeightDefault{
		_statusCode: code,
	}
}

/* RewardWeightDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type RewardWeightDefault struct {
	_statusCode int

	Payload *RewardWeightDefaultBody
}

// Code gets the status code for the reward weight default response
func (o *RewardWeightDefault) Code() int {
	return o._statusCode
}

func (o *RewardWeightDefault) Error() string {
	return fmt.Sprintf("[GET /terra/treasury/v1beta1/reward_weight][%d] RewardWeight default  %+v", o._statusCode, o.Payload)
}
func (o *RewardWeightDefault) GetPayload() *RewardWeightDefaultBody {
	return o.Payload
}

func (o *RewardWeightDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RewardWeightDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RewardWeightDefaultBody reward weight default body
swagger:model RewardWeightDefaultBody
*/
type RewardWeightDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*RewardWeightDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this reward weight default body
func (o *RewardWeightDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RewardWeightDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("RewardWeight default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RewardWeight default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this reward weight default body based on the context it is used
func (o *RewardWeightDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RewardWeightDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RewardWeight default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RewardWeight default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *RewardWeightDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RewardWeightDefaultBody) UnmarshalBinary(b []byte) error {
	var res RewardWeightDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RewardWeightDefaultBodyDetailsItems0 reward weight default body details items0
swagger:model RewardWeightDefaultBodyDetailsItems0
*/
type RewardWeightDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this reward weight default body details items0
func (o *RewardWeightDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this reward weight default body details items0 based on context it is used
func (o *RewardWeightDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RewardWeightDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RewardWeightDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RewardWeightDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RewardWeightOKBody QueryRewardWeightResponse is response type for the
// Query/RewardWeight RPC method.
swagger:model RewardWeightOKBody
*/
type RewardWeightOKBody struct {

	// reward weight
	RewardWeight string `json:"reward_weight,omitempty"`
}

// Validate validates this reward weight o k body
func (o *RewardWeightOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this reward weight o k body based on context it is used
func (o *RewardWeightOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RewardWeightOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RewardWeightOKBody) UnmarshalBinary(b []byte) error {
	var res RewardWeightOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
