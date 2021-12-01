// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lidofinance/terra-fcd-rest-client/models"
)

// GetV1MempoolReader is a Reader for the GetV1Mempool structure.
type GetV1MempoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1MempoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1MempoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1MempoolOK creates a GetV1MempoolOK with default headers values
func NewGetV1MempoolOK() *GetV1MempoolOK {
	return &GetV1MempoolOK{}
}

/* GetV1MempoolOK describes a response with status code 200, with default header values.

Success
*/
type GetV1MempoolOK struct {
	Payload []*models.Txs
}

func (o *GetV1MempoolOK) Error() string {
	return fmt.Sprintf("[GET /v1/mempool][%d] getV1MempoolOK  %+v", 200, o.Payload)
}
func (o *GetV1MempoolOK) GetPayload() []*models.Txs {
	return o.Payload
}

func (o *GetV1MempoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}