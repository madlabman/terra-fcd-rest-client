// Code generated by go-swagger; DO NOT EDIT.

package governance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetGovProposalsProposalIDVotesReader is a Reader for the GetGovProposalsProposalIDVotes structure.
type GetGovProposalsProposalIDVotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGovProposalsProposalIDVotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGovProposalsProposalIDVotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGovProposalsProposalIDVotesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGovProposalsProposalIDVotesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGovProposalsProposalIDVotesOK creates a GetGovProposalsProposalIDVotesOK with default headers values
func NewGetGovProposalsProposalIDVotesOK() *GetGovProposalsProposalIDVotesOK {
	return &GetGovProposalsProposalIDVotesOK{}
}

/* GetGovProposalsProposalIDVotesOK describes a response with status code 200, with default header values.

OK
*/
type GetGovProposalsProposalIDVotesOK struct {
	Payload []*GetGovProposalsProposalIDVotesOKBodyItems0
}

func (o *GetGovProposalsProposalIDVotesOK) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/votes][%d] getGovProposalsProposalIdVotesOK  %+v", 200, o.Payload)
}
func (o *GetGovProposalsProposalIDVotesOK) GetPayload() []*GetGovProposalsProposalIDVotesOKBodyItems0 {
	return o.Payload
}

func (o *GetGovProposalsProposalIDVotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGovProposalsProposalIDVotesBadRequest creates a GetGovProposalsProposalIDVotesBadRequest with default headers values
func NewGetGovProposalsProposalIDVotesBadRequest() *GetGovProposalsProposalIDVotesBadRequest {
	return &GetGovProposalsProposalIDVotesBadRequest{}
}

/* GetGovProposalsProposalIDVotesBadRequest describes a response with status code 400, with default header values.

Invalid proposal id
*/
type GetGovProposalsProposalIDVotesBadRequest struct {
}

func (o *GetGovProposalsProposalIDVotesBadRequest) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/votes][%d] getGovProposalsProposalIdVotesBadRequest ", 400)
}

func (o *GetGovProposalsProposalIDVotesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovProposalsProposalIDVotesInternalServerError creates a GetGovProposalsProposalIDVotesInternalServerError with default headers values
func NewGetGovProposalsProposalIDVotesInternalServerError() *GetGovProposalsProposalIDVotesInternalServerError {
	return &GetGovProposalsProposalIDVotesInternalServerError{}
}

/* GetGovProposalsProposalIDVotesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetGovProposalsProposalIDVotesInternalServerError struct {
}

func (o *GetGovProposalsProposalIDVotesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/votes][%d] getGovProposalsProposalIdVotesInternalServerError ", 500)
}

func (o *GetGovProposalsProposalIDVotesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetGovProposalsProposalIDVotesOKBodyItems0 get gov proposals proposal ID votes o k body items0
swagger:model GetGovProposalsProposalIDVotesOKBodyItems0
*/
type GetGovProposalsProposalIDVotesOKBodyItems0 struct {

	// options
	Options []*GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0 `json:"options"`

	// proposal id
	ProposalID string `json:"proposal_id,omitempty"`

	// voter
	Voter string `json:"voter,omitempty"`
}

// Validate validates this get gov proposals proposal ID votes o k body items0
func (o *GetGovProposalsProposalIDVotesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGovProposalsProposalIDVotesOKBodyItems0) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(o.Options) { // not required
		return nil
	}

	for i := 0; i < len(o.Options); i++ {
		if swag.IsZero(o.Options[i]) { // not required
			continue
		}

		if o.Options[i] != nil {
			if err := o.Options[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get gov proposals proposal ID votes o k body items0 based on the context it is used
func (o *GetGovProposalsProposalIDVotesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGovProposalsProposalIDVotesOKBodyItems0) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Options); i++ {

		if o.Options[i] != nil {
			if err := o.Options[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetGovProposalsProposalIDVotesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGovProposalsProposalIDVotesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetGovProposalsProposalIDVotesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0 get gov proposals proposal ID votes o k body items0 options items0
swagger:model GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0
*/
type GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0 struct {

	// VoteOption enumerates the valid vote options for a given governance proposal.
	// - VOTE_OPTION_UNSPECIFIED: VOTE_OPTION_UNSPECIFIED defines a no-op vote option. - VOTE_OPTION_YES: VOTE_OPTION_YES defines a yes vote option. - VOTE_OPTION_ABSTAIN: VOTE_OPTION_ABSTAIN defines an abstain vote option. - VOTE_OPTION_NO: VOTE_OPTION_NO defines a no vote option. - VOTE_OPTION_NO_WITH_VETO: VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
	// Enum: [VOTE_OPTION_UNSPECIFIED VOTE_OPTION_YES VOTE_OPTION_ABSTAIN VOTE_OPTION_NO VOTE_OPTION_NO_WITH_VETO]
	Option *string `json:"option,omitempty"`

	// weight
	Weight string `json:"weight,omitempty"`
}

// Validate validates this get gov proposals proposal ID votes o k body items0 options items0
func (o *GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getGovProposalsProposalIdVotesOKBodyItems0OptionsItems0TypeOptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VOTE_OPTION_UNSPECIFIED","VOTE_OPTION_YES","VOTE_OPTION_ABSTAIN","VOTE_OPTION_NO","VOTE_OPTION_NO_WITH_VETO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getGovProposalsProposalIdVotesOKBodyItems0OptionsItems0TypeOptionPropEnum = append(getGovProposalsProposalIdVotesOKBodyItems0OptionsItems0TypeOptionPropEnum, v)
	}
}

const (

	// GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0OptionVOTEOPTIONUNSPECIFIED captures enum value "VOTE_OPTION_UNSPECIFIED"
	GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0OptionVOTEOPTIONUNSPECIFIED string = "VOTE_OPTION_UNSPECIFIED"

	// GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0OptionVOTEOPTIONYES captures enum value "VOTE_OPTION_YES"
	GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0OptionVOTEOPTIONYES string = "VOTE_OPTION_YES"

	// GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0OptionVOTEOPTIONABSTAIN captures enum value "VOTE_OPTION_ABSTAIN"
	GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0OptionVOTEOPTIONABSTAIN string = "VOTE_OPTION_ABSTAIN"

	// GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0OptionVOTEOPTIONNO captures enum value "VOTE_OPTION_NO"
	GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0OptionVOTEOPTIONNO string = "VOTE_OPTION_NO"

	// GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0OptionVOTEOPTIONNOWITHVETO captures enum value "VOTE_OPTION_NO_WITH_VETO"
	GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0OptionVOTEOPTIONNOWITHVETO string = "VOTE_OPTION_NO_WITH_VETO"
)

// prop value enum
func (o *GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0) validateOptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getGovProposalsProposalIdVotesOKBodyItems0OptionsItems0TypeOptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0) validateOption(formats strfmt.Registry) error {
	if swag.IsZero(o.Option) { // not required
		return nil
	}

	// value enum
	if err := o.validateOptionEnum("option", "body", *o.Option); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get gov proposals proposal ID votes o k body items0 options items0 based on context it is used
func (o *GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0) UnmarshalBinary(b []byte) error {
	var res GetGovProposalsProposalIDVotesOKBodyItems0OptionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
