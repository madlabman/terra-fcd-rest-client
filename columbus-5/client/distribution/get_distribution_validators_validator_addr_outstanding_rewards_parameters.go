// Code generated by go-swagger; DO NOT EDIT.

package distribution

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

// NewGetDistributionValidatorsValidatorAddrOutstandingRewardsParams creates a new GetDistributionValidatorsValidatorAddrOutstandingRewardsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDistributionValidatorsValidatorAddrOutstandingRewardsParams() *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams {
	return &GetDistributionValidatorsValidatorAddrOutstandingRewardsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDistributionValidatorsValidatorAddrOutstandingRewardsParamsWithTimeout creates a new GetDistributionValidatorsValidatorAddrOutstandingRewardsParams object
// with the ability to set a timeout on a request.
func NewGetDistributionValidatorsValidatorAddrOutstandingRewardsParamsWithTimeout(timeout time.Duration) *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams {
	return &GetDistributionValidatorsValidatorAddrOutstandingRewardsParams{
		timeout: timeout,
	}
}

// NewGetDistributionValidatorsValidatorAddrOutstandingRewardsParamsWithContext creates a new GetDistributionValidatorsValidatorAddrOutstandingRewardsParams object
// with the ability to set a context for a request.
func NewGetDistributionValidatorsValidatorAddrOutstandingRewardsParamsWithContext(ctx context.Context) *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams {
	return &GetDistributionValidatorsValidatorAddrOutstandingRewardsParams{
		Context: ctx,
	}
}

// NewGetDistributionValidatorsValidatorAddrOutstandingRewardsParamsWithHTTPClient creates a new GetDistributionValidatorsValidatorAddrOutstandingRewardsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDistributionValidatorsValidatorAddrOutstandingRewardsParamsWithHTTPClient(client *http.Client) *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams {
	return &GetDistributionValidatorsValidatorAddrOutstandingRewardsParams{
		HTTPClient: client,
	}
}

/* GetDistributionValidatorsValidatorAddrOutstandingRewardsParams contains all the parameters to send to the API endpoint
   for the get distribution validators validator addr outstanding rewards operation.

   Typically these are written to a http.Request.
*/
type GetDistributionValidatorsValidatorAddrOutstandingRewardsParams struct {

	/* ValidatorAddr.

	   Bech32 OperatorAddress of validator
	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get distribution validators validator addr outstanding rewards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams) WithDefaults() *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get distribution validators validator addr outstanding rewards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get distribution validators validator addr outstanding rewards params
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams) WithTimeout(timeout time.Duration) *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get distribution validators validator addr outstanding rewards params
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get distribution validators validator addr outstanding rewards params
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams) WithContext(ctx context.Context) *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get distribution validators validator addr outstanding rewards params
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get distribution validators validator addr outstanding rewards params
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams) WithHTTPClient(client *http.Client) *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get distribution validators validator addr outstanding rewards params
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidatorAddr adds the validatorAddr to the get distribution validators validator addr outstanding rewards params
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams) WithValidatorAddr(validatorAddr string) *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the get distribution validators validator addr outstanding rewards params
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param validatorAddr
	if err := r.SetPathParam("validatorAddr", o.ValidatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}