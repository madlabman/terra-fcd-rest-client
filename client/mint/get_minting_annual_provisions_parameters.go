// Code generated by go-swagger; DO NOT EDIT.

package mint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetMintingAnnualProvisionsParams creates a new GetMintingAnnualProvisionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMintingAnnualProvisionsParams() *GetMintingAnnualProvisionsParams {
	return &GetMintingAnnualProvisionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMintingAnnualProvisionsParamsWithTimeout creates a new GetMintingAnnualProvisionsParams object
// with the ability to set a timeout on a request.
func NewGetMintingAnnualProvisionsParamsWithTimeout(timeout time.Duration) *GetMintingAnnualProvisionsParams {
	return &GetMintingAnnualProvisionsParams{
		timeout: timeout,
	}
}

// NewGetMintingAnnualProvisionsParamsWithContext creates a new GetMintingAnnualProvisionsParams object
// with the ability to set a context for a request.
func NewGetMintingAnnualProvisionsParamsWithContext(ctx context.Context) *GetMintingAnnualProvisionsParams {
	return &GetMintingAnnualProvisionsParams{
		Context: ctx,
	}
}

// NewGetMintingAnnualProvisionsParamsWithHTTPClient creates a new GetMintingAnnualProvisionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMintingAnnualProvisionsParamsWithHTTPClient(client *http.Client) *GetMintingAnnualProvisionsParams {
	return &GetMintingAnnualProvisionsParams{
		HTTPClient: client,
	}
}

/* GetMintingAnnualProvisionsParams contains all the parameters to send to the API endpoint
   for the get minting annual provisions operation.

   Typically these are written to a http.Request.
*/
type GetMintingAnnualProvisionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get minting annual provisions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMintingAnnualProvisionsParams) WithDefaults() *GetMintingAnnualProvisionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get minting annual provisions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMintingAnnualProvisionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get minting annual provisions params
func (o *GetMintingAnnualProvisionsParams) WithTimeout(timeout time.Duration) *GetMintingAnnualProvisionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get minting annual provisions params
func (o *GetMintingAnnualProvisionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get minting annual provisions params
func (o *GetMintingAnnualProvisionsParams) WithContext(ctx context.Context) *GetMintingAnnualProvisionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get minting annual provisions params
func (o *GetMintingAnnualProvisionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get minting annual provisions params
func (o *GetMintingAnnualProvisionsParams) WithHTTPClient(client *http.Client) *GetMintingAnnualProvisionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get minting annual provisions params
func (o *GetMintingAnnualProvisionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMintingAnnualProvisionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
