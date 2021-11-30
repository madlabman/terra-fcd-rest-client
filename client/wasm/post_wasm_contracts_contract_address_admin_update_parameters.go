// Code generated by go-swagger; DO NOT EDIT.

package wasm

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

// NewPostWasmContractsContractAddressAdminUpdateParams creates a new PostWasmContractsContractAddressAdminUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWasmContractsContractAddressAdminUpdateParams() *PostWasmContractsContractAddressAdminUpdateParams {
	return &PostWasmContractsContractAddressAdminUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWasmContractsContractAddressAdminUpdateParamsWithTimeout creates a new PostWasmContractsContractAddressAdminUpdateParams object
// with the ability to set a timeout on a request.
func NewPostWasmContractsContractAddressAdminUpdateParamsWithTimeout(timeout time.Duration) *PostWasmContractsContractAddressAdminUpdateParams {
	return &PostWasmContractsContractAddressAdminUpdateParams{
		timeout: timeout,
	}
}

// NewPostWasmContractsContractAddressAdminUpdateParamsWithContext creates a new PostWasmContractsContractAddressAdminUpdateParams object
// with the ability to set a context for a request.
func NewPostWasmContractsContractAddressAdminUpdateParamsWithContext(ctx context.Context) *PostWasmContractsContractAddressAdminUpdateParams {
	return &PostWasmContractsContractAddressAdminUpdateParams{
		Context: ctx,
	}
}

// NewPostWasmContractsContractAddressAdminUpdateParamsWithHTTPClient creates a new PostWasmContractsContractAddressAdminUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWasmContractsContractAddressAdminUpdateParamsWithHTTPClient(client *http.Client) *PostWasmContractsContractAddressAdminUpdateParams {
	return &PostWasmContractsContractAddressAdminUpdateParams{
		HTTPClient: client,
	}
}

/* PostWasmContractsContractAddressAdminUpdateParams contains all the parameters to send to the API endpoint
   for the post wasm contracts contract address admin update operation.

   Typically these are written to a http.Request.
*/
type PostWasmContractsContractAddressAdminUpdateParams struct {

	/* ContractAddress.

	   contract address you want to update admin
	*/
	ContractAddress string

	// UpdateContractAdminRequestBody.
	UpdateContractAdminRequestBody PostWasmContractsContractAddressAdminUpdateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post wasm contracts contract address admin update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWasmContractsContractAddressAdminUpdateParams) WithDefaults() *PostWasmContractsContractAddressAdminUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post wasm contracts contract address admin update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWasmContractsContractAddressAdminUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post wasm contracts contract address admin update params
func (o *PostWasmContractsContractAddressAdminUpdateParams) WithTimeout(timeout time.Duration) *PostWasmContractsContractAddressAdminUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post wasm contracts contract address admin update params
func (o *PostWasmContractsContractAddressAdminUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post wasm contracts contract address admin update params
func (o *PostWasmContractsContractAddressAdminUpdateParams) WithContext(ctx context.Context) *PostWasmContractsContractAddressAdminUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post wasm contracts contract address admin update params
func (o *PostWasmContractsContractAddressAdminUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post wasm contracts contract address admin update params
func (o *PostWasmContractsContractAddressAdminUpdateParams) WithHTTPClient(client *http.Client) *PostWasmContractsContractAddressAdminUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post wasm contracts contract address admin update params
func (o *PostWasmContractsContractAddressAdminUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContractAddress adds the contractAddress to the post wasm contracts contract address admin update params
func (o *PostWasmContractsContractAddressAdminUpdateParams) WithContractAddress(contractAddress string) *PostWasmContractsContractAddressAdminUpdateParams {
	o.SetContractAddress(contractAddress)
	return o
}

// SetContractAddress adds the contractAddress to the post wasm contracts contract address admin update params
func (o *PostWasmContractsContractAddressAdminUpdateParams) SetContractAddress(contractAddress string) {
	o.ContractAddress = contractAddress
}

// WithUpdateContractAdminRequestBody adds the updateContractAdminRequestBody to the post wasm contracts contract address admin update params
func (o *PostWasmContractsContractAddressAdminUpdateParams) WithUpdateContractAdminRequestBody(updateContractAdminRequestBody PostWasmContractsContractAddressAdminUpdateBody) *PostWasmContractsContractAddressAdminUpdateParams {
	o.SetUpdateContractAdminRequestBody(updateContractAdminRequestBody)
	return o
}

// SetUpdateContractAdminRequestBody adds the updateContractAdminRequestBody to the post wasm contracts contract address admin update params
func (o *PostWasmContractsContractAddressAdminUpdateParams) SetUpdateContractAdminRequestBody(updateContractAdminRequestBody PostWasmContractsContractAddressAdminUpdateBody) {
	o.UpdateContractAdminRequestBody = updateContractAdminRequestBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostWasmContractsContractAddressAdminUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contractAddress
	if err := r.SetPathParam("contractAddress", o.ContractAddress); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateContractAdminRequestBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
