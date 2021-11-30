// Code generated by go-swagger; DO NOT EDIT.

package treasury

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetTreasuryTaxProceedsReader is a Reader for the GetTreasuryTaxProceeds structure.
type GetTreasuryTaxProceedsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTreasuryTaxProceedsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTreasuryTaxProceedsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetTreasuryTaxProceedsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTreasuryTaxProceedsOK creates a GetTreasuryTaxProceedsOK with default headers values
func NewGetTreasuryTaxProceedsOK() *GetTreasuryTaxProceedsOK {
	return &GetTreasuryTaxProceedsOK{}
}

/* GetTreasuryTaxProceedsOK describes a response with status code 200, with default header values.

OK
*/
type GetTreasuryTaxProceedsOK struct {
	Payload []*GetTreasuryTaxProceedsOKBodyItems0
}

func (o *GetTreasuryTaxProceedsOK) Error() string {
	return fmt.Sprintf("[GET /treasury/tax_proceeds][%d] getTreasuryTaxProceedsOK  %+v", 200, o.Payload)
}
func (o *GetTreasuryTaxProceedsOK) GetPayload() []*GetTreasuryTaxProceedsOKBodyItems0 {
	return o.Payload
}

func (o *GetTreasuryTaxProceedsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTreasuryTaxProceedsInternalServerError creates a GetTreasuryTaxProceedsInternalServerError with default headers values
func NewGetTreasuryTaxProceedsInternalServerError() *GetTreasuryTaxProceedsInternalServerError {
	return &GetTreasuryTaxProceedsInternalServerError{}
}

/* GetTreasuryTaxProceedsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetTreasuryTaxProceedsInternalServerError struct {
}

func (o *GetTreasuryTaxProceedsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /treasury/tax_proceeds][%d] getTreasuryTaxProceedsInternalServerError ", 500)
}

func (o *GetTreasuryTaxProceedsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetTreasuryTaxProceedsOKBodyItems0 get treasury tax proceeds o k body items0
swagger:model GetTreasuryTaxProceedsOKBodyItems0
*/
type GetTreasuryTaxProceedsOKBodyItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this get treasury tax proceeds o k body items0
func (o *GetTreasuryTaxProceedsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get treasury tax proceeds o k body items0 based on context it is used
func (o *GetTreasuryTaxProceedsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTreasuryTaxProceedsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTreasuryTaxProceedsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetTreasuryTaxProceedsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}