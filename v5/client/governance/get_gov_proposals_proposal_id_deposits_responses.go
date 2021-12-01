// Code generated by go-swagger; DO NOT EDIT.

package governance

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

// GetGovProposalsProposalIDDepositsReader is a Reader for the GetGovProposalsProposalIDDeposits structure.
type GetGovProposalsProposalIDDepositsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGovProposalsProposalIDDepositsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGovProposalsProposalIDDepositsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGovProposalsProposalIDDepositsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGovProposalsProposalIDDepositsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGovProposalsProposalIDDepositsOK creates a GetGovProposalsProposalIDDepositsOK with default headers values
func NewGetGovProposalsProposalIDDepositsOK() *GetGovProposalsProposalIDDepositsOK {
	return &GetGovProposalsProposalIDDepositsOK{}
}

/* GetGovProposalsProposalIDDepositsOK describes a response with status code 200, with default header values.

OK
*/
type GetGovProposalsProposalIDDepositsOK struct {
	Payload []*GetGovProposalsProposalIDDepositsOKBodyItems0
}

func (o *GetGovProposalsProposalIDDepositsOK) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/deposits][%d] getGovProposalsProposalIdDepositsOK  %+v", 200, o.Payload)
}
func (o *GetGovProposalsProposalIDDepositsOK) GetPayload() []*GetGovProposalsProposalIDDepositsOKBodyItems0 {
	return o.Payload
}

func (o *GetGovProposalsProposalIDDepositsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGovProposalsProposalIDDepositsBadRequest creates a GetGovProposalsProposalIDDepositsBadRequest with default headers values
func NewGetGovProposalsProposalIDDepositsBadRequest() *GetGovProposalsProposalIDDepositsBadRequest {
	return &GetGovProposalsProposalIDDepositsBadRequest{}
}

/* GetGovProposalsProposalIDDepositsBadRequest describes a response with status code 400, with default header values.

Invalid proposal id
*/
type GetGovProposalsProposalIDDepositsBadRequest struct {
}

func (o *GetGovProposalsProposalIDDepositsBadRequest) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/deposits][%d] getGovProposalsProposalIdDepositsBadRequest ", 400)
}

func (o *GetGovProposalsProposalIDDepositsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovProposalsProposalIDDepositsInternalServerError creates a GetGovProposalsProposalIDDepositsInternalServerError with default headers values
func NewGetGovProposalsProposalIDDepositsInternalServerError() *GetGovProposalsProposalIDDepositsInternalServerError {
	return &GetGovProposalsProposalIDDepositsInternalServerError{}
}

/* GetGovProposalsProposalIDDepositsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetGovProposalsProposalIDDepositsInternalServerError struct {
}

func (o *GetGovProposalsProposalIDDepositsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/deposits][%d] getGovProposalsProposalIdDepositsInternalServerError ", 500)
}

func (o *GetGovProposalsProposalIDDepositsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetGovProposalsProposalIDDepositsOKBodyItems0 get gov proposals proposal ID deposits o k body items0
swagger:model GetGovProposalsProposalIDDepositsOKBodyItems0
*/
type GetGovProposalsProposalIDDepositsOKBodyItems0 struct {

	// amount
	Amount []*GetGovProposalsProposalIDDepositsOKBodyItems0AmountItems0 `json:"amount"`

	// bech32 encoded address
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	Depositor string `json:"depositor,omitempty"`

	// proposal id
	ProposalID string `json:"proposal_id,omitempty"`
}

// Validate validates this get gov proposals proposal ID deposits o k body items0
func (o *GetGovProposalsProposalIDDepositsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGovProposalsProposalIDDepositsOKBodyItems0) validateAmount(formats strfmt.Registry) error {
	if swag.IsZero(o.Amount) { // not required
		return nil
	}

	for i := 0; i < len(o.Amount); i++ {
		if swag.IsZero(o.Amount[i]) { // not required
			continue
		}

		if o.Amount[i] != nil {
			if err := o.Amount[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get gov proposals proposal ID deposits o k body items0 based on the context it is used
func (o *GetGovProposalsProposalIDDepositsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGovProposalsProposalIDDepositsOKBodyItems0) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Amount); i++ {

		if o.Amount[i] != nil {
			if err := o.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetGovProposalsProposalIDDepositsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGovProposalsProposalIDDepositsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetGovProposalsProposalIDDepositsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetGovProposalsProposalIDDepositsOKBodyItems0AmountItems0 get gov proposals proposal ID deposits o k body items0 amount items0
swagger:model GetGovProposalsProposalIDDepositsOKBodyItems0AmountItems0
*/
type GetGovProposalsProposalIDDepositsOKBodyItems0AmountItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this get gov proposals proposal ID deposits o k body items0 amount items0
func (o *GetGovProposalsProposalIDDepositsOKBodyItems0AmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get gov proposals proposal ID deposits o k body items0 amount items0 based on context it is used
func (o *GetGovProposalsProposalIDDepositsOKBodyItems0AmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetGovProposalsProposalIDDepositsOKBodyItems0AmountItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGovProposalsProposalIDDepositsOKBodyItems0AmountItems0) UnmarshalBinary(b []byte) error {
	var res GetGovProposalsProposalIDDepositsOKBodyItems0AmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}