// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lidofinance/terra-fcd-rest-client/models"
)

// GetV1DashboardSeigniorageProceedsReader is a Reader for the GetV1DashboardSeigniorageProceeds structure.
type GetV1DashboardSeigniorageProceedsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1DashboardSeigniorageProceedsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1DashboardSeigniorageProceedsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1DashboardSeigniorageProceedsOK creates a GetV1DashboardSeigniorageProceedsOK with default headers values
func NewGetV1DashboardSeigniorageProceedsOK() *GetV1DashboardSeigniorageProceedsOK {
	return &GetV1DashboardSeigniorageProceedsOK{}
}

/* GetV1DashboardSeigniorageProceedsOK describes a response with status code 200, with default header values.

Success
*/
type GetV1DashboardSeigniorageProceedsOK struct {
	Payload []*models.Seigniorage
}

func (o *GetV1DashboardSeigniorageProceedsOK) Error() string {
	return fmt.Sprintf("[GET /v1/dashboard/seigniorage_proceeds][%d] getV1DashboardSeigniorageProceedsOK  %+v", 200, o.Payload)
}
func (o *GetV1DashboardSeigniorageProceedsOK) GetPayload() []*models.Seigniorage {
	return o.Payload
}

func (o *GetV1DashboardSeigniorageProceedsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
