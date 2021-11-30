// Code generated by go-swagger; DO NOT EDIT.

package oracle

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

// NewGetOracleVotersValidatorAggregatePrevoteParams creates a new GetOracleVotersValidatorAggregatePrevoteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOracleVotersValidatorAggregatePrevoteParams() *GetOracleVotersValidatorAggregatePrevoteParams {
	return &GetOracleVotersValidatorAggregatePrevoteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOracleVotersValidatorAggregatePrevoteParamsWithTimeout creates a new GetOracleVotersValidatorAggregatePrevoteParams object
// with the ability to set a timeout on a request.
func NewGetOracleVotersValidatorAggregatePrevoteParamsWithTimeout(timeout time.Duration) *GetOracleVotersValidatorAggregatePrevoteParams {
	return &GetOracleVotersValidatorAggregatePrevoteParams{
		timeout: timeout,
	}
}

// NewGetOracleVotersValidatorAggregatePrevoteParamsWithContext creates a new GetOracleVotersValidatorAggregatePrevoteParams object
// with the ability to set a context for a request.
func NewGetOracleVotersValidatorAggregatePrevoteParamsWithContext(ctx context.Context) *GetOracleVotersValidatorAggregatePrevoteParams {
	return &GetOracleVotersValidatorAggregatePrevoteParams{
		Context: ctx,
	}
}

// NewGetOracleVotersValidatorAggregatePrevoteParamsWithHTTPClient creates a new GetOracleVotersValidatorAggregatePrevoteParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOracleVotersValidatorAggregatePrevoteParamsWithHTTPClient(client *http.Client) *GetOracleVotersValidatorAggregatePrevoteParams {
	return &GetOracleVotersValidatorAggregatePrevoteParams{
		HTTPClient: client,
	}
}

/* GetOracleVotersValidatorAggregatePrevoteParams contains all the parameters to send to the API endpoint
   for the get oracle voters validator aggregate prevote operation.

   Typically these are written to a http.Request.
*/
type GetOracleVotersValidatorAggregatePrevoteParams struct {

	/* Validator.

	   oracle operator
	*/
	Validator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get oracle voters validator aggregate prevote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOracleVotersValidatorAggregatePrevoteParams) WithDefaults() *GetOracleVotersValidatorAggregatePrevoteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get oracle voters validator aggregate prevote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOracleVotersValidatorAggregatePrevoteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get oracle voters validator aggregate prevote params
func (o *GetOracleVotersValidatorAggregatePrevoteParams) WithTimeout(timeout time.Duration) *GetOracleVotersValidatorAggregatePrevoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oracle voters validator aggregate prevote params
func (o *GetOracleVotersValidatorAggregatePrevoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oracle voters validator aggregate prevote params
func (o *GetOracleVotersValidatorAggregatePrevoteParams) WithContext(ctx context.Context) *GetOracleVotersValidatorAggregatePrevoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oracle voters validator aggregate prevote params
func (o *GetOracleVotersValidatorAggregatePrevoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oracle voters validator aggregate prevote params
func (o *GetOracleVotersValidatorAggregatePrevoteParams) WithHTTPClient(client *http.Client) *GetOracleVotersValidatorAggregatePrevoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oracle voters validator aggregate prevote params
func (o *GetOracleVotersValidatorAggregatePrevoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidator adds the validator to the get oracle voters validator aggregate prevote params
func (o *GetOracleVotersValidatorAggregatePrevoteParams) WithValidator(validator string) *GetOracleVotersValidatorAggregatePrevoteParams {
	o.SetValidator(validator)
	return o
}

// SetValidator adds the validator to the get oracle voters validator aggregate prevote params
func (o *GetOracleVotersValidatorAggregatePrevoteParams) SetValidator(validator string) {
	o.Validator = validator
}

// WriteToRequest writes these params to a swagger request
func (o *GetOracleVotersValidatorAggregatePrevoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param validator
	if err := r.SetPathParam("validator", o.Validator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
