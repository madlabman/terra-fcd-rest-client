// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IbcCoreClientV1QueryClientStatesResponse QueryClientStatesResponse is the response type for the Query/ClientStates RPC
// method.
//
// swagger:model ibc.core.client.v1.QueryClientStatesResponse
type IbcCoreClientV1QueryClientStatesResponse struct {

	// list of stored ClientStates of the chain.
	ClientStates []*IbcCoreClientV1QueryClientStatesResponseClientStatesItems0 `json:"client_states"`

	// pagination
	Pagination *IbcCoreClientV1QueryClientStatesResponsePagination `json:"pagination,omitempty"`
}

// Validate validates this ibc core client v1 query client states response
func (m *IbcCoreClientV1QueryClientStatesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientStates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreClientV1QueryClientStatesResponse) validateClientStates(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientStates) { // not required
		return nil
	}

	for i := 0; i < len(m.ClientStates); i++ {
		if swag.IsZero(m.ClientStates[i]) { // not required
			continue
		}

		if m.ClientStates[i] != nil {
			if err := m.ClientStates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("client_states" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("client_states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IbcCoreClientV1QueryClientStatesResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ibc core client v1 query client states response based on the context it is used
func (m *IbcCoreClientV1QueryClientStatesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClientStates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreClientV1QueryClientStatesResponse) contextValidateClientStates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClientStates); i++ {

		if m.ClientStates[i] != nil {
			if err := m.ClientStates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("client_states" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("client_states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IbcCoreClientV1QueryClientStatesResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreClientV1QueryClientStatesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreClientV1QueryClientStatesResponse) UnmarshalBinary(b []byte) error {
	var res IbcCoreClientV1QueryClientStatesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreClientV1QueryClientStatesResponseClientStatesItems0 IdentifiedClientState defines a client state with an additional client
// identifier field.
//
// swagger:model IbcCoreClientV1QueryClientStatesResponseClientStatesItems0
type IbcCoreClientV1QueryClientStatesResponseClientStatesItems0 struct {

	// client identifier
	ClientID string `json:"client_id,omitempty"`

	// client state
	ClientState *IbcCoreClientV1QueryClientStatesResponseClientStatesItems0ClientState `json:"client_state,omitempty"`
}

// Validate validates this ibc core client v1 query client states response client states items0
func (m *IbcCoreClientV1QueryClientStatesResponseClientStatesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreClientV1QueryClientStatesResponseClientStatesItems0) validateClientState(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientState) { // not required
		return nil
	}

	if m.ClientState != nil {
		if err := m.ClientState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client_state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ibc core client v1 query client states response client states items0 based on the context it is used
func (m *IbcCoreClientV1QueryClientStatesResponseClientStatesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClientState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreClientV1QueryClientStatesResponseClientStatesItems0) contextValidateClientState(ctx context.Context, formats strfmt.Registry) error {

	if m.ClientState != nil {
		if err := m.ClientState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client_state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreClientV1QueryClientStatesResponseClientStatesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreClientV1QueryClientStatesResponseClientStatesItems0) UnmarshalBinary(b []byte) error {
	var res IbcCoreClientV1QueryClientStatesResponseClientStatesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreClientV1QueryClientStatesResponseClientStatesItems0ClientState client state
//
// `Any` contains an arbitrary serialized protocol buffer message along with a
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
//
// swagger:model IbcCoreClientV1QueryClientStatesResponseClientStatesItems0ClientState
type IbcCoreClientV1QueryClientStatesResponseClientStatesItems0ClientState struct {

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

// Validate validates this ibc core client v1 query client states response client states items0 client state
func (m *IbcCoreClientV1QueryClientStatesResponseClientStatesItems0ClientState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core client v1 query client states response client states items0 client state based on context it is used
func (m *IbcCoreClientV1QueryClientStatesResponseClientStatesItems0ClientState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreClientV1QueryClientStatesResponseClientStatesItems0ClientState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreClientV1QueryClientStatesResponseClientStatesItems0ClientState) UnmarshalBinary(b []byte) error {
	var res IbcCoreClientV1QueryClientStatesResponseClientStatesItems0ClientState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreClientV1QueryClientStatesResponsePagination pagination response
//
// PageResponse is to be embedded in gRPC response messages where the
// corresponding request message has used PageRequest.
//
//  message SomeResponse {
//          repeated Bar results = 1;
//          PageResponse page = 2;
//  }
//
// swagger:model IbcCoreClientV1QueryClientStatesResponsePagination
type IbcCoreClientV1QueryClientStatesResponsePagination struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this ibc core client v1 query client states response pagination
func (m *IbcCoreClientV1QueryClientStatesResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core client v1 query client states response pagination based on context it is used
func (m *IbcCoreClientV1QueryClientStatesResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreClientV1QueryClientStatesResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreClientV1QueryClientStatesResponsePagination) UnmarshalBinary(b []byte) error {
	var res IbcCoreClientV1QueryClientStatesResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
