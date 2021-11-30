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
	"github.com/go-openapi/validate"
)

// UnbondingDelegationReader is a Reader for the UnbondingDelegation structure.
type UnbondingDelegationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnbondingDelegationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnbondingDelegationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnbondingDelegationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnbondingDelegationOK creates a UnbondingDelegationOK with default headers values
func NewUnbondingDelegationOK() *UnbondingDelegationOK {
	return &UnbondingDelegationOK{}
}

/* UnbondingDelegationOK describes a response with status code 200, with default header values.

A successful response.
*/
type UnbondingDelegationOK struct {
	Payload *UnbondingDelegationOKBody
}

func (o *UnbondingDelegationOK) Error() string {
	return fmt.Sprintf("[GET /cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}/unbonding_delegation][%d] unbondingDelegationOK  %+v", 200, o.Payload)
}
func (o *UnbondingDelegationOK) GetPayload() *UnbondingDelegationOKBody {
	return o.Payload
}

func (o *UnbondingDelegationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UnbondingDelegationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnbondingDelegationDefault creates a UnbondingDelegationDefault with default headers values
func NewUnbondingDelegationDefault(code int) *UnbondingDelegationDefault {
	return &UnbondingDelegationDefault{
		_statusCode: code,
	}
}

/* UnbondingDelegationDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type UnbondingDelegationDefault struct {
	_statusCode int

	Payload *UnbondingDelegationDefaultBody
}

// Code gets the status code for the unbonding delegation default response
func (o *UnbondingDelegationDefault) Code() int {
	return o._statusCode
}

func (o *UnbondingDelegationDefault) Error() string {
	return fmt.Sprintf("[GET /cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}/unbonding_delegation][%d] UnbondingDelegation default  %+v", o._statusCode, o.Payload)
}
func (o *UnbondingDelegationDefault) GetPayload() *UnbondingDelegationDefaultBody {
	return o.Payload
}

func (o *UnbondingDelegationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UnbondingDelegationDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UnbondingDelegationDefaultBody unbonding delegation default body
swagger:model UnbondingDelegationDefaultBody
*/
type UnbondingDelegationDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*UnbondingDelegationDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this unbonding delegation default body
func (o *UnbondingDelegationDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnbondingDelegationDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UnbondingDelegation default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UnbondingDelegation default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this unbonding delegation default body based on the context it is used
func (o *UnbondingDelegationDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnbondingDelegationDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UnbondingDelegation default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UnbondingDelegation default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UnbondingDelegationDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnbondingDelegationDefaultBody) UnmarshalBinary(b []byte) error {
	var res UnbondingDelegationDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UnbondingDelegationDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := ptypes.MarshalAny(foo)
//      ...
//      foo := &pb.Foo{}
//      if err := ptypes.UnmarshalAny(any, foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
swagger:model UnbondingDelegationDefaultBodyDetailsItems0
*/
type UnbondingDelegationDefaultBodyDetailsItems0 struct {

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	TypeURL string `json:"type_url,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this unbonding delegation default body details items0
func (o *UnbondingDelegationDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this unbonding delegation default body details items0 based on context it is used
func (o *UnbondingDelegationDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UnbondingDelegationDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnbondingDelegationDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UnbondingDelegationDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UnbondingDelegationOKBody QueryDelegationResponse is response type for the Query/UnbondingDelegation
// RPC method.
swagger:model UnbondingDelegationOKBody
*/
type UnbondingDelegationOKBody struct {

	// unbond
	Unbond *UnbondingDelegationOKBodyUnbond `json:"unbond,omitempty"`
}

// Validate validates this unbonding delegation o k body
func (o *UnbondingDelegationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUnbond(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnbondingDelegationOKBody) validateUnbond(formats strfmt.Registry) error {
	if swag.IsZero(o.Unbond) { // not required
		return nil
	}

	if o.Unbond != nil {
		if err := o.Unbond.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unbondingDelegationOK" + "." + "unbond")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unbondingDelegationOK" + "." + "unbond")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this unbonding delegation o k body based on the context it is used
func (o *UnbondingDelegationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUnbond(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnbondingDelegationOKBody) contextValidateUnbond(ctx context.Context, formats strfmt.Registry) error {

	if o.Unbond != nil {
		if err := o.Unbond.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unbondingDelegationOK" + "." + "unbond")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unbondingDelegationOK" + "." + "unbond")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UnbondingDelegationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnbondingDelegationOKBody) UnmarshalBinary(b []byte) error {
	var res UnbondingDelegationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UnbondingDelegationOKBodyUnbond UnbondingDelegation stores all of a single delegator's unbonding bonds
// for a single validator in an time-ordered list.
swagger:model UnbondingDelegationOKBodyUnbond
*/
type UnbondingDelegationOKBodyUnbond struct {

	// delegator_address is the bech32-encoded address of the delegator.
	DelegatorAddress string `json:"delegator_address,omitempty"`

	// entries are the unbonding delegation entries.
	Entries []*UnbondingDelegationOKBodyUnbondEntriesItems0 `json:"entries"`

	// validator_address is the bech32-encoded address of the validator.
	ValidatorAddress string `json:"validator_address,omitempty"`
}

// Validate validates this unbonding delegation o k body unbond
func (o *UnbondingDelegationOKBodyUnbond) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnbondingDelegationOKBodyUnbond) validateEntries(formats strfmt.Registry) error {
	if swag.IsZero(o.Entries) { // not required
		return nil
	}

	for i := 0; i < len(o.Entries); i++ {
		if swag.IsZero(o.Entries[i]) { // not required
			continue
		}

		if o.Entries[i] != nil {
			if err := o.Entries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unbondingDelegationOK" + "." + "unbond" + "." + "entries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("unbondingDelegationOK" + "." + "unbond" + "." + "entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this unbonding delegation o k body unbond based on the context it is used
func (o *UnbondingDelegationOKBodyUnbond) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEntries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnbondingDelegationOKBodyUnbond) contextValidateEntries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Entries); i++ {

		if o.Entries[i] != nil {
			if err := o.Entries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unbondingDelegationOK" + "." + "unbond" + "." + "entries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("unbondingDelegationOK" + "." + "unbond" + "." + "entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UnbondingDelegationOKBodyUnbond) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnbondingDelegationOKBodyUnbond) UnmarshalBinary(b []byte) error {
	var res UnbondingDelegationOKBodyUnbond
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UnbondingDelegationOKBodyUnbondEntriesItems0 UnbondingDelegationEntry defines an unbonding object with relevant metadata.
swagger:model UnbondingDelegationOKBodyUnbondEntriesItems0
*/
type UnbondingDelegationOKBodyUnbondEntriesItems0 struct {

	// balance defines the tokens to receive at completion.
	Balance string `json:"balance,omitempty"`

	// completion_time is the unix time for unbonding completion.
	// Format: date-time
	CompletionTime strfmt.DateTime `json:"completion_time,omitempty"`

	// creation_height is the height which the unbonding took place.
	CreationHeight string `json:"creation_height,omitempty"`

	// initial_balance defines the tokens initially scheduled to receive at completion.
	InitialBalance string `json:"initial_balance,omitempty"`
}

// Validate validates this unbonding delegation o k body unbond entries items0
func (o *UnbondingDelegationOKBodyUnbondEntriesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCompletionTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnbondingDelegationOKBodyUnbondEntriesItems0) validateCompletionTime(formats strfmt.Registry) error {
	if swag.IsZero(o.CompletionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("completion_time", "body", "date-time", o.CompletionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this unbonding delegation o k body unbond entries items0 based on context it is used
func (o *UnbondingDelegationOKBodyUnbondEntriesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UnbondingDelegationOKBodyUnbondEntriesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnbondingDelegationOKBodyUnbondEntriesItems0) UnmarshalBinary(b []byte) error {
	var res UnbondingDelegationOKBodyUnbondEntriesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
