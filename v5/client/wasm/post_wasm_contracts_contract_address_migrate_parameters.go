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

// NewPostWasmContractsContractAddressMigrateParams creates a new PostWasmContractsContractAddressMigrateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWasmContractsContractAddressMigrateParams() *PostWasmContractsContractAddressMigrateParams {
	return &PostWasmContractsContractAddressMigrateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWasmContractsContractAddressMigrateParamsWithTimeout creates a new PostWasmContractsContractAddressMigrateParams object
// with the ability to set a timeout on a request.
func NewPostWasmContractsContractAddressMigrateParamsWithTimeout(timeout time.Duration) *PostWasmContractsContractAddressMigrateParams {
	return &PostWasmContractsContractAddressMigrateParams{
		timeout: timeout,
	}
}

// NewPostWasmContractsContractAddressMigrateParamsWithContext creates a new PostWasmContractsContractAddressMigrateParams object
// with the ability to set a context for a request.
func NewPostWasmContractsContractAddressMigrateParamsWithContext(ctx context.Context) *PostWasmContractsContractAddressMigrateParams {
	return &PostWasmContractsContractAddressMigrateParams{
		Context: ctx,
	}
}

// NewPostWasmContractsContractAddressMigrateParamsWithHTTPClient creates a new PostWasmContractsContractAddressMigrateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWasmContractsContractAddressMigrateParamsWithHTTPClient(client *http.Client) *PostWasmContractsContractAddressMigrateParams {
	return &PostWasmContractsContractAddressMigrateParams{
		HTTPClient: client,
	}
}

/* PostWasmContractsContractAddressMigrateParams contains all the parameters to send to the API endpoint
   for the post wasm contracts contract address migrate operation.

   Typically these are written to a http.Request.
*/
type PostWasmContractsContractAddressMigrateParams struct {

	/* ContractAddress.

	   contract address you want to migrate
	*/
	ContractAddress string

	// MigrateContractRequestBody.
	MigrateContractRequestBody PostWasmContractsContractAddressMigrateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post wasm contracts contract address migrate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWasmContractsContractAddressMigrateParams) WithDefaults() *PostWasmContractsContractAddressMigrateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post wasm contracts contract address migrate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWasmContractsContractAddressMigrateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post wasm contracts contract address migrate params
func (o *PostWasmContractsContractAddressMigrateParams) WithTimeout(timeout time.Duration) *PostWasmContractsContractAddressMigrateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post wasm contracts contract address migrate params
func (o *PostWasmContractsContractAddressMigrateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post wasm contracts contract address migrate params
func (o *PostWasmContractsContractAddressMigrateParams) WithContext(ctx context.Context) *PostWasmContractsContractAddressMigrateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post wasm contracts contract address migrate params
func (o *PostWasmContractsContractAddressMigrateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post wasm contracts contract address migrate params
func (o *PostWasmContractsContractAddressMigrateParams) WithHTTPClient(client *http.Client) *PostWasmContractsContractAddressMigrateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post wasm contracts contract address migrate params
func (o *PostWasmContractsContractAddressMigrateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContractAddress adds the contractAddress to the post wasm contracts contract address migrate params
func (o *PostWasmContractsContractAddressMigrateParams) WithContractAddress(contractAddress string) *PostWasmContractsContractAddressMigrateParams {
	o.SetContractAddress(contractAddress)
	return o
}

// SetContractAddress adds the contractAddress to the post wasm contracts contract address migrate params
func (o *PostWasmContractsContractAddressMigrateParams) SetContractAddress(contractAddress string) {
	o.ContractAddress = contractAddress
}

// WithMigrateContractRequestBody adds the migrateContractRequestBody to the post wasm contracts contract address migrate params
func (o *PostWasmContractsContractAddressMigrateParams) WithMigrateContractRequestBody(migrateContractRequestBody PostWasmContractsContractAddressMigrateBody) *PostWasmContractsContractAddressMigrateParams {
	o.SetMigrateContractRequestBody(migrateContractRequestBody)
	return o
}

// SetMigrateContractRequestBody adds the migrateContractRequestBody to the post wasm contracts contract address migrate params
func (o *PostWasmContractsContractAddressMigrateParams) SetMigrateContractRequestBody(migrateContractRequestBody PostWasmContractsContractAddressMigrateBody) {
	o.MigrateContractRequestBody = migrateContractRequestBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostWasmContractsContractAddressMigrateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contractAddress
	if err := r.SetPathParam("contractAddress", o.ContractAddress); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.MigrateContractRequestBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}