// Code generated by go-swagger; DO NOT EDIT.

package staking

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

// NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams creates a new GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams() *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams {
	return &GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsParamsWithTimeout creates a new GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams object
// with the ability to set a timeout on a request.
func NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsParamsWithTimeout(timeout time.Duration) *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams {
	return &GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams{
		timeout: timeout,
	}
}

// NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsParamsWithContext creates a new GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams object
// with the ability to set a context for a request.
func NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsParamsWithContext(ctx context.Context) *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams {
	return &GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams{
		Context: ctx,
	}
}

// NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsParamsWithHTTPClient creates a new GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsParamsWithHTTPClient(client *http.Client) *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams {
	return &GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams{
		HTTPClient: client,
	}
}

/* GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams contains all the parameters to send to the API endpoint
   for the get staking delegators delegator addr unbonding delegations operation.

   Typically these are written to a http.Request.
*/
type GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams struct {

	/* DelegatorAddr.

	   Bech32 AccAddress of Delegator
	*/
	DelegatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get staking delegators delegator addr unbonding delegations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) WithDefaults() *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get staking delegators delegator addr unbonding delegations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get staking delegators delegator addr unbonding delegations params
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) WithTimeout(timeout time.Duration) *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get staking delegators delegator addr unbonding delegations params
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get staking delegators delegator addr unbonding delegations params
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) WithContext(ctx context.Context) *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get staking delegators delegator addr unbonding delegations params
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get staking delegators delegator addr unbonding delegations params
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) WithHTTPClient(client *http.Client) *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get staking delegators delegator addr unbonding delegations params
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegatorAddr adds the delegatorAddr to the get staking delegators delegator addr unbonding delegations params
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) WithDelegatorAddr(delegatorAddr string) *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the get staking delegators delegator addr unbonding delegations params
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param delegatorAddr
	if err := r.SetPathParam("delegatorAddr", o.DelegatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}