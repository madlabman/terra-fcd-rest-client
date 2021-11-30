// Code generated by go-swagger; DO NOT EDIT.

package governance

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

// NewGetGovProposalsProposalIDVotesParams creates a new GetGovProposalsProposalIDVotesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGovProposalsProposalIDVotesParams() *GetGovProposalsProposalIDVotesParams {
	return &GetGovProposalsProposalIDVotesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGovProposalsProposalIDVotesParamsWithTimeout creates a new GetGovProposalsProposalIDVotesParams object
// with the ability to set a timeout on a request.
func NewGetGovProposalsProposalIDVotesParamsWithTimeout(timeout time.Duration) *GetGovProposalsProposalIDVotesParams {
	return &GetGovProposalsProposalIDVotesParams{
		timeout: timeout,
	}
}

// NewGetGovProposalsProposalIDVotesParamsWithContext creates a new GetGovProposalsProposalIDVotesParams object
// with the ability to set a context for a request.
func NewGetGovProposalsProposalIDVotesParamsWithContext(ctx context.Context) *GetGovProposalsProposalIDVotesParams {
	return &GetGovProposalsProposalIDVotesParams{
		Context: ctx,
	}
}

// NewGetGovProposalsProposalIDVotesParamsWithHTTPClient creates a new GetGovProposalsProposalIDVotesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGovProposalsProposalIDVotesParamsWithHTTPClient(client *http.Client) *GetGovProposalsProposalIDVotesParams {
	return &GetGovProposalsProposalIDVotesParams{
		HTTPClient: client,
	}
}

/* GetGovProposalsProposalIDVotesParams contains all the parameters to send to the API endpoint
   for the get gov proposals proposal ID votes operation.

   Typically these are written to a http.Request.
*/
type GetGovProposalsProposalIDVotesParams struct {

	/* ProposalID.

	   proposal id
	*/
	ProposalID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get gov proposals proposal ID votes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGovProposalsProposalIDVotesParams) WithDefaults() *GetGovProposalsProposalIDVotesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gov proposals proposal ID votes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGovProposalsProposalIDVotesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get gov proposals proposal ID votes params
func (o *GetGovProposalsProposalIDVotesParams) WithTimeout(timeout time.Duration) *GetGovProposalsProposalIDVotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gov proposals proposal ID votes params
func (o *GetGovProposalsProposalIDVotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gov proposals proposal ID votes params
func (o *GetGovProposalsProposalIDVotesParams) WithContext(ctx context.Context) *GetGovProposalsProposalIDVotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gov proposals proposal ID votes params
func (o *GetGovProposalsProposalIDVotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gov proposals proposal ID votes params
func (o *GetGovProposalsProposalIDVotesParams) WithHTTPClient(client *http.Client) *GetGovProposalsProposalIDVotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gov proposals proposal ID votes params
func (o *GetGovProposalsProposalIDVotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProposalID adds the proposalID to the get gov proposals proposal ID votes params
func (o *GetGovProposalsProposalIDVotesParams) WithProposalID(proposalID string) *GetGovProposalsProposalIDVotesParams {
	o.SetProposalID(proposalID)
	return o
}

// SetProposalID adds the proposalId to the get gov proposals proposal ID votes params
func (o *GetGovProposalsProposalIDVotesParams) SetProposalID(proposalID string) {
	o.ProposalID = proposalID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGovProposalsProposalIDVotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param proposalId
	if err := r.SetPathParam("proposalId", o.ProposalID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
