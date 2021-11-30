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

// NewPostGovProposalsProposalIDDepositsParams creates a new PostGovProposalsProposalIDDepositsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostGovProposalsProposalIDDepositsParams() *PostGovProposalsProposalIDDepositsParams {
	return &PostGovProposalsProposalIDDepositsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostGovProposalsProposalIDDepositsParamsWithTimeout creates a new PostGovProposalsProposalIDDepositsParams object
// with the ability to set a timeout on a request.
func NewPostGovProposalsProposalIDDepositsParamsWithTimeout(timeout time.Duration) *PostGovProposalsProposalIDDepositsParams {
	return &PostGovProposalsProposalIDDepositsParams{
		timeout: timeout,
	}
}

// NewPostGovProposalsProposalIDDepositsParamsWithContext creates a new PostGovProposalsProposalIDDepositsParams object
// with the ability to set a context for a request.
func NewPostGovProposalsProposalIDDepositsParamsWithContext(ctx context.Context) *PostGovProposalsProposalIDDepositsParams {
	return &PostGovProposalsProposalIDDepositsParams{
		Context: ctx,
	}
}

// NewPostGovProposalsProposalIDDepositsParamsWithHTTPClient creates a new PostGovProposalsProposalIDDepositsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostGovProposalsProposalIDDepositsParamsWithHTTPClient(client *http.Client) *PostGovProposalsProposalIDDepositsParams {
	return &PostGovProposalsProposalIDDepositsParams{
		HTTPClient: client,
	}
}

/* PostGovProposalsProposalIDDepositsParams contains all the parameters to send to the API endpoint
   for the post gov proposals proposal ID deposits operation.

   Typically these are written to a http.Request.
*/
type PostGovProposalsProposalIDDepositsParams struct {

	// PostDepositBody.
	PostDepositBody PostGovProposalsProposalIDDepositsBody

	/* ProposalID.

	   proposal id
	*/
	ProposalID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post gov proposals proposal ID deposits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGovProposalsProposalIDDepositsParams) WithDefaults() *PostGovProposalsProposalIDDepositsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post gov proposals proposal ID deposits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGovProposalsProposalIDDepositsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post gov proposals proposal ID deposits params
func (o *PostGovProposalsProposalIDDepositsParams) WithTimeout(timeout time.Duration) *PostGovProposalsProposalIDDepositsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post gov proposals proposal ID deposits params
func (o *PostGovProposalsProposalIDDepositsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post gov proposals proposal ID deposits params
func (o *PostGovProposalsProposalIDDepositsParams) WithContext(ctx context.Context) *PostGovProposalsProposalIDDepositsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post gov proposals proposal ID deposits params
func (o *PostGovProposalsProposalIDDepositsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post gov proposals proposal ID deposits params
func (o *PostGovProposalsProposalIDDepositsParams) WithHTTPClient(client *http.Client) *PostGovProposalsProposalIDDepositsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post gov proposals proposal ID deposits params
func (o *PostGovProposalsProposalIDDepositsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostDepositBody adds the postDepositBody to the post gov proposals proposal ID deposits params
func (o *PostGovProposalsProposalIDDepositsParams) WithPostDepositBody(postDepositBody PostGovProposalsProposalIDDepositsBody) *PostGovProposalsProposalIDDepositsParams {
	o.SetPostDepositBody(postDepositBody)
	return o
}

// SetPostDepositBody adds the postDepositBody to the post gov proposals proposal ID deposits params
func (o *PostGovProposalsProposalIDDepositsParams) SetPostDepositBody(postDepositBody PostGovProposalsProposalIDDepositsBody) {
	o.PostDepositBody = postDepositBody
}

// WithProposalID adds the proposalID to the post gov proposals proposal ID deposits params
func (o *PostGovProposalsProposalIDDepositsParams) WithProposalID(proposalID string) *PostGovProposalsProposalIDDepositsParams {
	o.SetProposalID(proposalID)
	return o
}

// SetProposalID adds the proposalId to the post gov proposals proposal ID deposits params
func (o *PostGovProposalsProposalIDDepositsParams) SetProposalID(proposalID string) {
	o.ProposalID = proposalID
}

// WriteToRequest writes these params to a swagger request
func (o *PostGovProposalsProposalIDDepositsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.PostDepositBody); err != nil {
		return err
	}

	// path param proposalId
	if err := r.SetPathParam("proposalId", o.ProposalID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
