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

// GetV1TxsGasPricesReader is a Reader for the GetV1TxsGasPrices structure.
type GetV1TxsGasPricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TxsGasPricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TxsGasPricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1TxsGasPricesOK creates a GetV1TxsGasPricesOK with default headers values
func NewGetV1TxsGasPricesOK() *GetV1TxsGasPricesOK {
	return &GetV1TxsGasPricesOK{}
}

/* GetV1TxsGasPricesOK describes a response with status code 200, with default header values.

Success
*/
type GetV1TxsGasPricesOK struct {
	Payload *models.GetGasPricesResult
}

func (o *GetV1TxsGasPricesOK) Error() string {
	return fmt.Sprintf("[GET /v1/txs/gas_prices][%d] getV1TxsGasPricesOK  %+v", 200, o.Payload)
}
func (o *GetV1TxsGasPricesOK) GetPayload() *models.GetGasPricesResult {
	return o.Payload
}

func (o *GetV1TxsGasPricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetGasPricesResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
