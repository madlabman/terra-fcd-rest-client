// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CosmosGovV1beta1QueryProposalResponse QueryProposalResponse is the response type for the Query/Proposal RPC method.
//
// swagger:model cosmos.gov.v1beta1.QueryProposalResponse
type CosmosGovV1beta1QueryProposalResponse struct {

	// proposal
	Proposal *CosmosGovV1beta1QueryProposalResponseProposal `json:"proposal,omitempty"`
}

// Validate validates this cosmos gov v1beta1 query proposal response
func (m *CosmosGovV1beta1QueryProposalResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProposal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponse) validateProposal(formats strfmt.Registry) error {
	if swag.IsZero(m.Proposal) { // not required
		return nil
	}

	if m.Proposal != nil {
		if err := m.Proposal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proposal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proposal")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos gov v1beta1 query proposal response based on the context it is used
func (m *CosmosGovV1beta1QueryProposalResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProposal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponse) contextValidateProposal(ctx context.Context, formats strfmt.Registry) error {

	if m.Proposal != nil {
		if err := m.Proposal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proposal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proposal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryProposalResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryProposalResponse) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryProposalResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosGovV1beta1QueryProposalResponseProposal Proposal defines the core field members of a governance proposal.
//
// swagger:model CosmosGovV1beta1QueryProposalResponseProposal
type CosmosGovV1beta1QueryProposalResponseProposal struct {

	// content
	Content *CosmosGovV1beta1QueryProposalResponseProposalContent `json:"content,omitempty"`

	// deposit end time
	// Format: date-time
	DepositEndTime strfmt.DateTime `json:"deposit_end_time,omitempty"`

	// final tally result
	FinalTallyResult *CosmosGovV1beta1QueryProposalResponseProposalFinalTallyResult `json:"final_tally_result,omitempty"`

	// proposal id
	ProposalID string `json:"proposal_id,omitempty"`

	// ProposalStatus enumerates the valid statuses of a proposal.
	//
	//  - PROPOSAL_STATUS_UNSPECIFIED: PROPOSAL_STATUS_UNSPECIFIED defines the default propopsal status.
	//  - PROPOSAL_STATUS_DEPOSIT_PERIOD: PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit
	// period.
	//  - PROPOSAL_STATUS_VOTING_PERIOD: PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting
	// period.
	//  - PROPOSAL_STATUS_PASSED: PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has
	// passed.
	//  - PROPOSAL_STATUS_REJECTED: PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has
	// been rejected.
	//  - PROPOSAL_STATUS_FAILED: PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has
	// failed.
	// Enum: [PROPOSAL_STATUS_UNSPECIFIED PROPOSAL_STATUS_DEPOSIT_PERIOD PROPOSAL_STATUS_VOTING_PERIOD PROPOSAL_STATUS_PASSED PROPOSAL_STATUS_REJECTED PROPOSAL_STATUS_FAILED]
	Status *string `json:"status,omitempty"`

	// submit time
	// Format: date-time
	SubmitTime strfmt.DateTime `json:"submit_time,omitempty"`

	// total deposit
	TotalDeposit []*CosmosGovV1beta1QueryProposalResponseProposalTotalDepositItems0 `json:"total_deposit"`

	// voting end time
	// Format: date-time
	VotingEndTime strfmt.DateTime `json:"voting_end_time,omitempty"`

	// voting start time
	// Format: date-time
	VotingStartTime strfmt.DateTime `json:"voting_start_time,omitempty"`
}

// Validate validates this cosmos gov v1beta1 query proposal response proposal
func (m *CosmosGovV1beta1QueryProposalResponseProposal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepositEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinalTallyResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmitTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalDeposit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVotingEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVotingStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponseProposal) validateContent(formats strfmt.Registry) error {
	if swag.IsZero(m.Content) { // not required
		return nil
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proposal" + "." + "content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proposal" + "." + "content")
			}
			return err
		}
	}

	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponseProposal) validateDepositEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DepositEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("proposal"+"."+"deposit_end_time", "body", "date-time", m.DepositEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponseProposal) validateFinalTallyResult(formats strfmt.Registry) error {
	if swag.IsZero(m.FinalTallyResult) { // not required
		return nil
	}

	if m.FinalTallyResult != nil {
		if err := m.FinalTallyResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proposal" + "." + "final_tally_result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proposal" + "." + "final_tally_result")
			}
			return err
		}
	}

	return nil
}

var cosmosGovV1beta1QueryProposalResponseProposalTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROPOSAL_STATUS_UNSPECIFIED","PROPOSAL_STATUS_DEPOSIT_PERIOD","PROPOSAL_STATUS_VOTING_PERIOD","PROPOSAL_STATUS_PASSED","PROPOSAL_STATUS_REJECTED","PROPOSAL_STATUS_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cosmosGovV1beta1QueryProposalResponseProposalTypeStatusPropEnum = append(cosmosGovV1beta1QueryProposalResponseProposalTypeStatusPropEnum, v)
	}
}

const (

	// CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSUNSPECIFIED captures enum value "PROPOSAL_STATUS_UNSPECIFIED"
	CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSUNSPECIFIED string = "PROPOSAL_STATUS_UNSPECIFIED"

	// CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSDEPOSITPERIOD captures enum value "PROPOSAL_STATUS_DEPOSIT_PERIOD"
	CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSDEPOSITPERIOD string = "PROPOSAL_STATUS_DEPOSIT_PERIOD"

	// CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSVOTINGPERIOD captures enum value "PROPOSAL_STATUS_VOTING_PERIOD"
	CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSVOTINGPERIOD string = "PROPOSAL_STATUS_VOTING_PERIOD"

	// CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSPASSED captures enum value "PROPOSAL_STATUS_PASSED"
	CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSPASSED string = "PROPOSAL_STATUS_PASSED"

	// CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSREJECTED captures enum value "PROPOSAL_STATUS_REJECTED"
	CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSREJECTED string = "PROPOSAL_STATUS_REJECTED"

	// CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSFAILED captures enum value "PROPOSAL_STATUS_FAILED"
	CosmosGovV1beta1QueryProposalResponseProposalStatusPROPOSALSTATUSFAILED string = "PROPOSAL_STATUS_FAILED"
)

// prop value enum
func (m *CosmosGovV1beta1QueryProposalResponseProposal) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cosmosGovV1beta1QueryProposalResponseProposalTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponseProposal) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("proposal"+"."+"status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponseProposal) validateSubmitTime(formats strfmt.Registry) error {
	if swag.IsZero(m.SubmitTime) { // not required
		return nil
	}

	if err := validate.FormatOf("proposal"+"."+"submit_time", "body", "date-time", m.SubmitTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponseProposal) validateTotalDeposit(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalDeposit) { // not required
		return nil
	}

	for i := 0; i < len(m.TotalDeposit); i++ {
		if swag.IsZero(m.TotalDeposit[i]) { // not required
			continue
		}

		if m.TotalDeposit[i] != nil {
			if err := m.TotalDeposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proposal" + "." + "total_deposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proposal" + "." + "total_deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponseProposal) validateVotingEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.VotingEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("proposal"+"."+"voting_end_time", "body", "date-time", m.VotingEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponseProposal) validateVotingStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.VotingStartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("proposal"+"."+"voting_start_time", "body", "date-time", m.VotingStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cosmos gov v1beta1 query proposal response proposal based on the context it is used
func (m *CosmosGovV1beta1QueryProposalResponseProposal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFinalTallyResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalDeposit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponseProposal) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	if m.Content != nil {
		if err := m.Content.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proposal" + "." + "content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proposal" + "." + "content")
			}
			return err
		}
	}

	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponseProposal) contextValidateFinalTallyResult(ctx context.Context, formats strfmt.Registry) error {

	if m.FinalTallyResult != nil {
		if err := m.FinalTallyResult.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proposal" + "." + "final_tally_result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proposal" + "." + "final_tally_result")
			}
			return err
		}
	}

	return nil
}

func (m *CosmosGovV1beta1QueryProposalResponseProposal) contextValidateTotalDeposit(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TotalDeposit); i++ {

		if m.TotalDeposit[i] != nil {
			if err := m.TotalDeposit[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proposal" + "." + "total_deposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proposal" + "." + "total_deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryProposalResponseProposal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryProposalResponseProposal) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryProposalResponseProposal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosGovV1beta1QueryProposalResponseProposalContent `Any` contains an arbitrary serialized protocol buffer message along with a
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
// swagger:model CosmosGovV1beta1QueryProposalResponseProposalContent
type CosmosGovV1beta1QueryProposalResponseProposalContent struct {

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

// Validate validates this cosmos gov v1beta1 query proposal response proposal content
func (m *CosmosGovV1beta1QueryProposalResponseProposalContent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos gov v1beta1 query proposal response proposal content based on context it is used
func (m *CosmosGovV1beta1QueryProposalResponseProposalContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryProposalResponseProposalContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryProposalResponseProposalContent) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryProposalResponseProposalContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosGovV1beta1QueryProposalResponseProposalFinalTallyResult TallyResult defines a standard tally for a governance proposal.
//
// swagger:model CosmosGovV1beta1QueryProposalResponseProposalFinalTallyResult
type CosmosGovV1beta1QueryProposalResponseProposalFinalTallyResult struct {

	// abstain
	Abstain string `json:"abstain,omitempty"`

	// no
	No string `json:"no,omitempty"`

	// no with veto
	NoWithVeto string `json:"no_with_veto,omitempty"`

	// yes
	Yes string `json:"yes,omitempty"`
}

// Validate validates this cosmos gov v1beta1 query proposal response proposal final tally result
func (m *CosmosGovV1beta1QueryProposalResponseProposalFinalTallyResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos gov v1beta1 query proposal response proposal final tally result based on context it is used
func (m *CosmosGovV1beta1QueryProposalResponseProposalFinalTallyResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryProposalResponseProposalFinalTallyResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryProposalResponseProposalFinalTallyResult) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryProposalResponseProposalFinalTallyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosGovV1beta1QueryProposalResponseProposalTotalDepositItems0 Coin defines a token with a denomination and an amount.
//
// NOTE: The amount field is an Int which implements the custom method
// signatures required by gogoproto.
//
// swagger:model CosmosGovV1beta1QueryProposalResponseProposalTotalDepositItems0
type CosmosGovV1beta1QueryProposalResponseProposalTotalDepositItems0 struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this cosmos gov v1beta1 query proposal response proposal total deposit items0
func (m *CosmosGovV1beta1QueryProposalResponseProposalTotalDepositItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos gov v1beta1 query proposal response proposal total deposit items0 based on context it is used
func (m *CosmosGovV1beta1QueryProposalResponseProposalTotalDepositItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryProposalResponseProposalTotalDepositItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryProposalResponseProposalTotalDepositItems0) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryProposalResponseProposalTotalDepositItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
