// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lidofinance/terra-fcd-rest-client/v5/models"
)

// GetV1MempoolTxhashReader is a Reader for the GetV1MempoolTxhash structure.
type GetV1MempoolTxhashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1MempoolTxhashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1MempoolTxhashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1MempoolTxhashOK creates a GetV1MempoolTxhashOK with default headers values
func NewGetV1MempoolTxhashOK() *GetV1MempoolTxhashOK {
	return &GetV1MempoolTxhashOK{}
}

/* GetV1MempoolTxhashOK describes a response with status code 200, with default header values.

Success
*/
type GetV1MempoolTxhashOK struct {
	Payload *models.GetMempoolByHashResult
}

func (o *GetV1MempoolTxhashOK) Error() string {
	return fmt.Sprintf("[GET /v1/mempool/{txhash}][%d] getV1MempoolTxhashOK  %+v", 200, o.Payload)
}
func (o *GetV1MempoolTxhashOK) GetPayload() *models.GetMempoolByHashResult {
	return o.Payload
}

func (o *GetV1MempoolTxhashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMempoolByHashResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
