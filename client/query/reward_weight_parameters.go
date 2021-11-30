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

// NewRewardWeightParams creates a new RewardWeightParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRewardWeightParams() *RewardWeightParams {
	return &RewardWeightParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRewardWeightParamsWithTimeout creates a new RewardWeightParams object
// with the ability to set a timeout on a request.
func NewRewardWeightParamsWithTimeout(timeout time.Duration) *RewardWeightParams {
	return &RewardWeightParams{
		timeout: timeout,
	}
}

// NewRewardWeightParamsWithContext creates a new RewardWeightParams object
// with the ability to set a context for a request.
func NewRewardWeightParamsWithContext(ctx context.Context) *RewardWeightParams {
	return &RewardWeightParams{
		Context: ctx,
	}
}

// NewRewardWeightParamsWithHTTPClient creates a new RewardWeightParams object
// with the ability to set a custom HTTPClient for a request.
func NewRewardWeightParamsWithHTTPClient(client *http.Client) *RewardWeightParams {
	return &RewardWeightParams{
		HTTPClient: client,
	}
}

/* RewardWeightParams contains all the parameters to send to the API endpoint
   for the reward weight operation.

   Typically these are written to a http.Request.
*/
type RewardWeightParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reward weight params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RewardWeightParams) WithDefaults() *RewardWeightParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reward weight params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RewardWeightParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reward weight params
func (o *RewardWeightParams) WithTimeout(timeout time.Duration) *RewardWeightParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reward weight params
func (o *RewardWeightParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reward weight params
func (o *RewardWeightParams) WithContext(ctx context.Context) *RewardWeightParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reward weight params
func (o *RewardWeightParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reward weight params
func (o *RewardWeightParams) WithHTTPClient(client *http.Client) *RewardWeightParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reward weight params
func (o *RewardWeightParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RewardWeightParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
