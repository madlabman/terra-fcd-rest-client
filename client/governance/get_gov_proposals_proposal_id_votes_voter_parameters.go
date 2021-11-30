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

// NewGetGovProposalsProposalIDVotesVoterParams creates a new GetGovProposalsProposalIDVotesVoterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGovProposalsProposalIDVotesVoterParams() *GetGovProposalsProposalIDVotesVoterParams {
	return &GetGovProposalsProposalIDVotesVoterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGovProposalsProposalIDVotesVoterParamsWithTimeout creates a new GetGovProposalsProposalIDVotesVoterParams object
// with the ability to set a timeout on a request.
func NewGetGovProposalsProposalIDVotesVoterParamsWithTimeout(timeout time.Duration) *GetGovProposalsProposalIDVotesVoterParams {
	return &GetGovProposalsProposalIDVotesVoterParams{
		timeout: timeout,
	}
}

// NewGetGovProposalsProposalIDVotesVoterParamsWithContext creates a new GetGovProposalsProposalIDVotesVoterParams object
// with the ability to set a context for a request.
func NewGetGovProposalsProposalIDVotesVoterParamsWithContext(ctx context.Context) *GetGovProposalsProposalIDVotesVoterParams {
	return &GetGovProposalsProposalIDVotesVoterParams{
		Context: ctx,
	}
}

// NewGetGovProposalsProposalIDVotesVoterParamsWithHTTPClient creates a new GetGovProposalsProposalIDVotesVoterParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGovProposalsProposalIDVotesVoterParamsWithHTTPClient(client *http.Client) *GetGovProposalsProposalIDVotesVoterParams {
	return &GetGovProposalsProposalIDVotesVoterParams{
		HTTPClient: client,
	}
}

/* GetGovProposalsProposalIDVotesVoterParams contains all the parameters to send to the API endpoint
   for the get gov proposals proposal ID votes voter operation.

   Typically these are written to a http.Request.
*/
type GetGovProposalsProposalIDVotesVoterParams struct {

	/* ProposalID.

	   proposal id
	*/
	ProposalID string

	/* Voter.

	   Bech32 voter address
	*/
	Voter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get gov proposals proposal ID votes voter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGovProposalsProposalIDVotesVoterParams) WithDefaults() *GetGovProposalsProposalIDVotesVoterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gov proposals proposal ID votes voter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGovProposalsProposalIDVotesVoterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get gov proposals proposal ID votes voter params
func (o *GetGovProposalsProposalIDVotesVoterParams) WithTimeout(timeout time.Duration) *GetGovProposalsProposalIDVotesVoterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gov proposals proposal ID votes voter params
func (o *GetGovProposalsProposalIDVotesVoterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gov proposals proposal ID votes voter params
func (o *GetGovProposalsProposalIDVotesVoterParams) WithContext(ctx context.Context) *GetGovProposalsProposalIDVotesVoterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gov proposals proposal ID votes voter params
func (o *GetGovProposalsProposalIDVotesVoterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gov proposals proposal ID votes voter params
func (o *GetGovProposalsProposalIDVotesVoterParams) WithHTTPClient(client *http.Client) *GetGovProposalsProposalIDVotesVoterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gov proposals proposal ID votes voter params
func (o *GetGovProposalsProposalIDVotesVoterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProposalID adds the proposalID to the get gov proposals proposal ID votes voter params
func (o *GetGovProposalsProposalIDVotesVoterParams) WithProposalID(proposalID string) *GetGovProposalsProposalIDVotesVoterParams {
	o.SetProposalID(proposalID)
	return o
}

// SetProposalID adds the proposalId to the get gov proposals proposal ID votes voter params
func (o *GetGovProposalsProposalIDVotesVoterParams) SetProposalID(proposalID string) {
	o.ProposalID = proposalID
}

// WithVoter adds the voter to the get gov proposals proposal ID votes voter params
func (o *GetGovProposalsProposalIDVotesVoterParams) WithVoter(voter string) *GetGovProposalsProposalIDVotesVoterParams {
	o.SetVoter(voter)
	return o
}

// SetVoter adds the voter to the get gov proposals proposal ID votes voter params
func (o *GetGovProposalsProposalIDVotesVoterParams) SetVoter(voter string) {
	o.Voter = voter
}

// WriteToRequest writes these params to a swagger request
func (o *GetGovProposalsProposalIDVotesVoterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param proposalId
	if err := r.SetPathParam("proposalId", o.ProposalID); err != nil {
		return err
	}

	// path param voter
	if err := r.SetPathParam("voter", o.Voter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}