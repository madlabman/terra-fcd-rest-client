// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosBaseQueryV1beta1PageResponse PageResponse is to be embedded in gRPC response messages where the
// corresponding request message has used PageRequest.
//
//  message SomeResponse {
//          repeated Bar results = 1;
//          PageResponse page = 2;
//  }
//
// swagger:model cosmos.base.query.v1beta1.PageResponse
type CosmosBaseQueryV1beta1PageResponse struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this cosmos base query v1beta1 page response
func (m *CosmosBaseQueryV1beta1PageResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos base query v1beta1 page response based on context it is used
func (m *CosmosBaseQueryV1beta1PageResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseQueryV1beta1PageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseQueryV1beta1PageResponse) UnmarshalBinary(b []byte) error {
	var res CosmosBaseQueryV1beta1PageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}