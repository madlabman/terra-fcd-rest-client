// Code generated by go-swagger; DO NOT EDIT.

package query

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

// NewPoolParams creates a new PoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPoolParams() *PoolParams {
	return &PoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPoolParamsWithTimeout creates a new PoolParams object
// with the ability to set a timeout on a request.
func NewPoolParamsWithTimeout(timeout time.Duration) *PoolParams {
	return &PoolParams{
		timeout: timeout,
	}
}

// NewPoolParamsWithContext creates a new PoolParams object
// with the ability to set a context for a request.
func NewPoolParamsWithContext(ctx context.Context) *PoolParams {
	return &PoolParams{
		Context: ctx,
	}
}

// NewPoolParamsWithHTTPClient creates a new PoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewPoolParamsWithHTTPClient(client *http.Client) *PoolParams {
	return &PoolParams{
		HTTPClient: client,
	}
}

/* PoolParams contains all the parameters to send to the API endpoint
   for the pool operation.

   Typically these are written to a http.Request.
*/
type PoolParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PoolParams) WithDefaults() *PoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pool params
func (o *PoolParams) WithTimeout(timeout time.Duration) *PoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pool params
func (o *PoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pool params
func (o *PoolParams) WithContext(ctx context.Context) *PoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pool params
func (o *PoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pool params
func (o *PoolParams) WithHTTPClient(client *http.Client) *PoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pool params
func (o *PoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
