// Code generated by go-swagger; DO NOT EDIT.

package governance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lidofinance/terra-fcd-rest-client/v5/models"
)

// GetV1GovProposalsReader is a Reader for the GetV1GovProposals structure.
type GetV1GovProposalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1GovProposalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1GovProposalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1GovProposalsOK creates a GetV1GovProposalsOK with default headers values
func NewGetV1GovProposalsOK() *GetV1GovProposalsOK {
	return &GetV1GovProposalsOK{}
}

/* GetV1GovProposalsOK describes a response with status code 200, with default header values.

Success
*/
type GetV1GovProposalsOK struct {
	Payload *models.GetProposalListResult
}

func (o *GetV1GovProposalsOK) Error() string {
	return fmt.Sprintf("[GET /v1/gov/proposals][%d] getV1GovProposalsOK  %+v", 200, o.Payload)
}
func (o *GetV1GovProposalsOK) GetPayload() *models.GetProposalListResult {
	return o.Payload
}

func (o *GetV1GovProposalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetProposalListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
