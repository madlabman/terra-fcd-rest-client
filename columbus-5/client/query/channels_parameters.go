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
	"github.com/go-openapi/swag"
)

// NewChannelsParams creates a new ChannelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChannelsParams() *ChannelsParams {
	return &ChannelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChannelsParamsWithTimeout creates a new ChannelsParams object
// with the ability to set a timeout on a request.
func NewChannelsParamsWithTimeout(timeout time.Duration) *ChannelsParams {
	return &ChannelsParams{
		timeout: timeout,
	}
}

// NewChannelsParamsWithContext creates a new ChannelsParams object
// with the ability to set a context for a request.
func NewChannelsParamsWithContext(ctx context.Context) *ChannelsParams {
	return &ChannelsParams{
		Context: ctx,
	}
}

// NewChannelsParamsWithHTTPClient creates a new ChannelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewChannelsParamsWithHTTPClient(client *http.Client) *ChannelsParams {
	return &ChannelsParams{
		HTTPClient: client,
	}
}

/* ChannelsParams contains all the parameters to send to the API endpoint
   for the channels operation.

   Typically these are written to a http.Request.
*/
type ChannelsParams struct {

	/* PaginationCountTotal.

	     count_total is set to true  to indicate that the result set should include
	a count of the total number of items available for pagination in UIs.
	count_total is only respected when offset is used. It is ignored when key
	is set.

	     Format: boolean
	*/
	PaginationCountTotal *bool

	/* PaginationKey.

	     key is a value returned in PageResponse.next_key to begin
	querying the next page most efficiently. Only one of offset or key
	should be set.

	     Format: byte
	*/
	PaginationKey *strfmt.Base64

	/* PaginationLimit.

	     limit is the total number of results to be returned in the result page.
	If left empty it will default to a value to be set by each app.

	     Format: uint64
	*/
	PaginationLimit *string

	/* PaginationOffset.

	     offset is a numeric offset that can be used when key is unavailable.
	It is less efficient than using key. Only one of offset or key should
	be set.

	     Format: uint64
	*/
	PaginationOffset *string

	/* PaginationReverse.

	   reverse is set to true if results are to be returned in the descending order.

	   Format: boolean
	*/
	PaginationReverse *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the channels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChannelsParams) WithDefaults() *ChannelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the channels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChannelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the channels params
func (o *ChannelsParams) WithTimeout(timeout time.Duration) *ChannelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the channels params
func (o *ChannelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the channels params
func (o *ChannelsParams) WithContext(ctx context.Context) *ChannelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the channels params
func (o *ChannelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the channels params
func (o *ChannelsParams) WithHTTPClient(client *http.Client) *ChannelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the channels params
func (o *ChannelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaginationCountTotal adds the paginationCountTotal to the channels params
func (o *ChannelsParams) WithPaginationCountTotal(paginationCountTotal *bool) *ChannelsParams {
	o.SetPaginationCountTotal(paginationCountTotal)
	return o
}

// SetPaginationCountTotal adds the paginationCountTotal to the channels params
func (o *ChannelsParams) SetPaginationCountTotal(paginationCountTotal *bool) {
	o.PaginationCountTotal = paginationCountTotal
}

// WithPaginationKey adds the paginationKey to the channels params
func (o *ChannelsParams) WithPaginationKey(paginationKey *strfmt.Base64) *ChannelsParams {
	o.SetPaginationKey(paginationKey)
	return o
}

// SetPaginationKey adds the paginationKey to the channels params
func (o *ChannelsParams) SetPaginationKey(paginationKey *strfmt.Base64) {
	o.PaginationKey = paginationKey
}

// WithPaginationLimit adds the paginationLimit to the channels params
func (o *ChannelsParams) WithPaginationLimit(paginationLimit *string) *ChannelsParams {
	o.SetPaginationLimit(paginationLimit)
	return o
}

// SetPaginationLimit adds the paginationLimit to the channels params
func (o *ChannelsParams) SetPaginationLimit(paginationLimit *string) {
	o.PaginationLimit = paginationLimit
}

// WithPaginationOffset adds the paginationOffset to the channels params
func (o *ChannelsParams) WithPaginationOffset(paginationOffset *string) *ChannelsParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the channels params
func (o *ChannelsParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationReverse adds the paginationReverse to the channels params
func (o *ChannelsParams) WithPaginationReverse(paginationReverse *bool) *ChannelsParams {
	o.SetPaginationReverse(paginationReverse)
	return o
}

// SetPaginationReverse adds the paginationReverse to the channels params
func (o *ChannelsParams) SetPaginationReverse(paginationReverse *bool) {
	o.PaginationReverse = paginationReverse
}

// WriteToRequest writes these params to a swagger request
func (o *ChannelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PaginationCountTotal != nil {

		// query param pagination.count_total
		var qrPaginationCountTotal bool

		if o.PaginationCountTotal != nil {
			qrPaginationCountTotal = *o.PaginationCountTotal
		}
		qPaginationCountTotal := swag.FormatBool(qrPaginationCountTotal)
		if qPaginationCountTotal != "" {

			if err := r.SetQueryParam("pagination.count_total", qPaginationCountTotal); err != nil {
				return err
			}
		}
	}

	if o.PaginationKey != nil {

		// query param pagination.key
		var qrPaginationKey strfmt.Base64

		if o.PaginationKey != nil {
			qrPaginationKey = *o.PaginationKey
		}
		qPaginationKey := qrPaginationKey.String()
		if qPaginationKey != "" {

			if err := r.SetQueryParam("pagination.key", qPaginationKey); err != nil {
				return err
			}
		}
	}

	if o.PaginationLimit != nil {

		// query param pagination.limit
		var qrPaginationLimit string

		if o.PaginationLimit != nil {
			qrPaginationLimit = *o.PaginationLimit
		}
		qPaginationLimit := qrPaginationLimit
		if qPaginationLimit != "" {

			if err := r.SetQueryParam("pagination.limit", qPaginationLimit); err != nil {
				return err
			}
		}
	}

	if o.PaginationOffset != nil {

		// query param pagination.offset
		var qrPaginationOffset string

		if o.PaginationOffset != nil {
			qrPaginationOffset = *o.PaginationOffset
		}
		qPaginationOffset := qrPaginationOffset
		if qPaginationOffset != "" {

			if err := r.SetQueryParam("pagination.offset", qPaginationOffset); err != nil {
				return err
			}
		}
	}

	if o.PaginationReverse != nil {

		// query param pagination.reverse
		var qrPaginationReverse bool

		if o.PaginationReverse != nil {
			qrPaginationReverse = *o.PaginationReverse
		}
		qPaginationReverse := swag.FormatBool(qrPaginationReverse)
		if qPaginationReverse != "" {

			if err := r.SetQueryParam("pagination.reverse", qPaginationReverse); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}