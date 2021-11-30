// Code generated by go-swagger; DO NOT EDIT.

package wasm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lidofinance/terra-fcd-rest-client/models"
)

// GetV1WasmCodeCodeIDReader is a Reader for the GetV1WasmCodeCodeID structure.
type GetV1WasmCodeCodeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WasmCodeCodeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WasmCodeCodeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1WasmCodeCodeIDOK creates a GetV1WasmCodeCodeIDOK with default headers values
func NewGetV1WasmCodeCodeIDOK() *GetV1WasmCodeCodeIDOK {
	return &GetV1WasmCodeCodeIDOK{}
}

/* GetV1WasmCodeCodeIDOK describes a response with status code 200, with default header values.

Success
*/
type GetV1WasmCodeCodeIDOK struct {
	Payload *models.GetIndividualWasmCodeResult
}

func (o *GetV1WasmCodeCodeIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/wasm/code/{code_id}][%d] getV1WasmCodeCodeIdOK  %+v", 200, o.Payload)
}
func (o *GetV1WasmCodeCodeIDOK) GetPayload() *models.GetIndividualWasmCodeResult {
	return o.Payload
}

func (o *GetV1WasmCodeCodeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIndividualWasmCodeResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
