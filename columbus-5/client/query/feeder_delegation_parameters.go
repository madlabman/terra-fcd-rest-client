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

// NewFeederDelegationParams creates a new FeederDelegationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFeederDelegationParams() *FeederDelegationParams {
	return &FeederDelegationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFeederDelegationParamsWithTimeout creates a new FeederDelegationParams object
// with the ability to set a timeout on a request.
func NewFeederDelegationParamsWithTimeout(timeout time.Duration) *FeederDelegationParams {
	return &FeederDelegationParams{
		timeout: timeout,
	}
}

// NewFeederDelegationParamsWithContext creates a new FeederDelegationParams object
// with the ability to set a context for a request.
func NewFeederDelegationParamsWithContext(ctx context.Context) *FeederDelegationParams {
	return &FeederDelegationParams{
		Context: ctx,
	}
}

// NewFeederDelegationParamsWithHTTPClient creates a new FeederDelegationParams object
// with the ability to set a custom HTTPClient for a request.
func NewFeederDelegationParamsWithHTTPClient(client *http.Client) *FeederDelegationParams {
	return &FeederDelegationParams{
		HTTPClient: client,
	}
}

/* FeederDelegationParams contains all the parameters to send to the API endpoint
   for the feeder delegation operation.

   Typically these are written to a http.Request.
*/
type FeederDelegationParams struct {

	/* ValidatorAddr.

	   validator defines the validator address to query for.
	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the feeder delegation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FeederDelegationParams) WithDefaults() *FeederDelegationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the feeder delegation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FeederDelegationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the feeder delegation params
func (o *FeederDelegationParams) WithTimeout(timeout time.Duration) *FeederDelegationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the feeder delegation params
func (o *FeederDelegationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the feeder delegation params
func (o *FeederDelegationParams) WithContext(ctx context.Context) *FeederDelegationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the feeder delegation params
func (o *FeederDelegationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the feeder delegation params
func (o *FeederDelegationParams) WithHTTPClient(client *http.Client) *FeederDelegationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the feeder delegation params
func (o *FeederDelegationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidatorAddr adds the validatorAddr to the feeder delegation params
func (o *FeederDelegationParams) WithValidatorAddr(validatorAddr string) *FeederDelegationParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the feeder delegation params
func (o *FeederDelegationParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *FeederDelegationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param validator_addr
	if err := r.SetPathParam("validator_addr", o.ValidatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}