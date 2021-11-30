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

// NewDelegatorValidatorParams creates a new DelegatorValidatorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDelegatorValidatorParams() *DelegatorValidatorParams {
	return &DelegatorValidatorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDelegatorValidatorParamsWithTimeout creates a new DelegatorValidatorParams object
// with the ability to set a timeout on a request.
func NewDelegatorValidatorParamsWithTimeout(timeout time.Duration) *DelegatorValidatorParams {
	return &DelegatorValidatorParams{
		timeout: timeout,
	}
}

// NewDelegatorValidatorParamsWithContext creates a new DelegatorValidatorParams object
// with the ability to set a context for a request.
func NewDelegatorValidatorParamsWithContext(ctx context.Context) *DelegatorValidatorParams {
	return &DelegatorValidatorParams{
		Context: ctx,
	}
}

// NewDelegatorValidatorParamsWithHTTPClient creates a new DelegatorValidatorParams object
// with the ability to set a custom HTTPClient for a request.
func NewDelegatorValidatorParamsWithHTTPClient(client *http.Client) *DelegatorValidatorParams {
	return &DelegatorValidatorParams{
		HTTPClient: client,
	}
}

/* DelegatorValidatorParams contains all the parameters to send to the API endpoint
   for the delegator validator operation.

   Typically these are written to a http.Request.
*/
type DelegatorValidatorParams struct {

	/* DelegatorAddr.

	   delegator_addr defines the delegator address to query for.
	*/
	DelegatorAddr string

	/* ValidatorAddr.

	   validator_addr defines the validator address to query for.
	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delegator validator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DelegatorValidatorParams) WithDefaults() *DelegatorValidatorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delegator validator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DelegatorValidatorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delegator validator params
func (o *DelegatorValidatorParams) WithTimeout(timeout time.Duration) *DelegatorValidatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delegator validator params
func (o *DelegatorValidatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delegator validator params
func (o *DelegatorValidatorParams) WithContext(ctx context.Context) *DelegatorValidatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delegator validator params
func (o *DelegatorValidatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delegator validator params
func (o *DelegatorValidatorParams) WithHTTPClient(client *http.Client) *DelegatorValidatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delegator validator params
func (o *DelegatorValidatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegatorAddr adds the delegatorAddr to the delegator validator params
func (o *DelegatorValidatorParams) WithDelegatorAddr(delegatorAddr string) *DelegatorValidatorParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the delegator validator params
func (o *DelegatorValidatorParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WithValidatorAddr adds the validatorAddr to the delegator validator params
func (o *DelegatorValidatorParams) WithValidatorAddr(validatorAddr string) *DelegatorValidatorParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the delegator validator params
func (o *DelegatorValidatorParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *DelegatorValidatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param delegator_addr
	if err := r.SetPathParam("delegator_addr", o.DelegatorAddr); err != nil {
		return err
	}

	// path param validator_addr
	if err := r.SetPathParam("validator_addr", o.ValidatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
