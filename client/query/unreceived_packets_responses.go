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

// UnreceivedPacketsReader is a Reader for the UnreceivedPackets structure.
type UnreceivedPacketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnreceivedPacketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnreceivedPacketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnreceivedPacketsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnreceivedPacketsOK creates a UnreceivedPacketsOK with default headers values
func NewUnreceivedPacketsOK() *UnreceivedPacketsOK {
	return &UnreceivedPacketsOK{}
}

/* UnreceivedPacketsOK describes a response with status code 200, with default header values.

A successful response.
*/
type UnreceivedPacketsOK struct {
	Payload *UnreceivedPacketsOKBody
}

func (o *UnreceivedPacketsOK) Error() string {
	return fmt.Sprintf("[GET /ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_commitments/{packet_commitment_sequences}/unreceived_packets][%d] unreceivedPacketsOK  %+v", 200, o.Payload)
}
func (o *UnreceivedPacketsOK) GetPayload() *UnreceivedPacketsOKBody {
	return o.Payload
}

func (o *UnreceivedPacketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UnreceivedPacketsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnreceivedPacketsDefault creates a UnreceivedPacketsDefault with default headers values
func NewUnreceivedPacketsDefault(code int) *UnreceivedPacketsDefault {
	return &UnreceivedPacketsDefault{
		_statusCode: code,
	}
}

/* UnreceivedPacketsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type UnreceivedPacketsDefault struct {
	_statusCode int

	Payload *UnreceivedPacketsDefaultBody
}

// Code gets the status code for the unreceived packets default response
func (o *UnreceivedPacketsDefault) Code() int {
	return o._statusCode
}

func (o *UnreceivedPacketsDefault) Error() string {
	return fmt.Sprintf("[GET /ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_commitments/{packet_commitment_sequences}/unreceived_packets][%d] UnreceivedPackets default  %+v", o._statusCode, o.Payload)
}
func (o *UnreceivedPacketsDefault) GetPayload() *UnreceivedPacketsDefaultBody {
	return o.Payload
}

func (o *UnreceivedPacketsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UnreceivedPacketsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UnreceivedPacketsDefaultBody unreceived packets default body
swagger:model UnreceivedPacketsDefaultBody
*/
type UnreceivedPacketsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*UnreceivedPacketsDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this unreceived packets default body
func (o *UnreceivedPacketsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnreceivedPacketsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UnreceivedPackets default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UnreceivedPackets default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this unreceived packets default body based on the context it is used
func (o *UnreceivedPacketsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnreceivedPacketsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UnreceivedPackets default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UnreceivedPackets default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UnreceivedPacketsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnreceivedPacketsDefaultBody) UnmarshalBinary(b []byte) error {
	var res UnreceivedPacketsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UnreceivedPacketsDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
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
swagger:model UnreceivedPacketsDefaultBodyDetailsItems0
*/
type UnreceivedPacketsDefaultBodyDetailsItems0 struct {

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

// Validate validates this unreceived packets default body details items0
func (o *UnreceivedPacketsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this unreceived packets default body details items0 based on context it is used
func (o *UnreceivedPacketsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UnreceivedPacketsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnreceivedPacketsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UnreceivedPacketsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UnreceivedPacketsOKBody QueryUnreceivedPacketsResponse is the response type for the
// Query/UnreceivedPacketCommitments RPC method
swagger:model UnreceivedPacketsOKBody
*/
type UnreceivedPacketsOKBody struct {

	// height
	Height *UnreceivedPacketsOKBodyHeight `json:"height,omitempty"`

	// list of unreceived packet sequences
	Sequences []string `json:"sequences"`
}

// Validate validates this unreceived packets o k body
func (o *UnreceivedPacketsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnreceivedPacketsOKBody) validateHeight(formats strfmt.Registry) error {
	if swag.IsZero(o.Height) { // not required
		return nil
	}

	if o.Height != nil {
		if err := o.Height.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unreceivedPacketsOK" + "." + "height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unreceivedPacketsOK" + "." + "height")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this unreceived packets o k body based on the context it is used
func (o *UnreceivedPacketsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnreceivedPacketsOKBody) contextValidateHeight(ctx context.Context, formats strfmt.Registry) error {

	if o.Height != nil {
		if err := o.Height.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unreceivedPacketsOK" + "." + "height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unreceivedPacketsOK" + "." + "height")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UnreceivedPacketsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnreceivedPacketsOKBody) UnmarshalBinary(b []byte) error {
	var res UnreceivedPacketsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UnreceivedPacketsOKBodyHeight query block height
//
// Normally the RevisionHeight is incremented at each height while keeping
// RevisionNumber the same. However some consensus algorithms may choose to
// reset the height in certain conditions e.g. hard forks, state-machine
// breaking changes In these cases, the RevisionNumber is incremented so that
// height continues to be monitonically increasing even as the RevisionHeight
// gets reset
swagger:model UnreceivedPacketsOKBodyHeight
*/
type UnreceivedPacketsOKBodyHeight struct {

	// the height within the given revision
	RevisionHeight string `json:"revision_height,omitempty"`

	// the revision that the client is currently on
	RevisionNumber string `json:"revision_number,omitempty"`
}

// Validate validates this unreceived packets o k body height
func (o *UnreceivedPacketsOKBodyHeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this unreceived packets o k body height based on context it is used
func (o *UnreceivedPacketsOKBodyHeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UnreceivedPacketsOKBodyHeight) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnreceivedPacketsOKBodyHeight) UnmarshalBinary(b []byte) error {
	var res UnreceivedPacketsOKBodyHeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
