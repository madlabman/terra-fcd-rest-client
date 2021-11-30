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

// NewGetWasmContractsContractAddressStoreRawParams creates a new GetWasmContractsContractAddressStoreRawParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWasmContractsContractAddressStoreRawParams() *GetWasmContractsContractAddressStoreRawParams {
	return &GetWasmContractsContractAddressStoreRawParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWasmContractsContractAddressStoreRawParamsWithTimeout creates a new GetWasmContractsContractAddressStoreRawParams object
// with the ability to set a timeout on a request.
func NewGetWasmContractsContractAddressStoreRawParamsWithTimeout(timeout time.Duration) *GetWasmContractsContractAddressStoreRawParams {
	return &GetWasmContractsContractAddressStoreRawParams{
		timeout: timeout,
	}
}

// NewGetWasmContractsContractAddressStoreRawParamsWithContext creates a new GetWasmContractsContractAddressStoreRawParams object
// with the ability to set a context for a request.
func NewGetWasmContractsContractAddressStoreRawParamsWithContext(ctx context.Context) *GetWasmContractsContractAddressStoreRawParams {
	return &GetWasmContractsContractAddressStoreRawParams{
		Context: ctx,
	}
}

// NewGetWasmContractsContractAddressStoreRawParamsWithHTTPClient creates a new GetWasmContractsContractAddressStoreRawParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWasmContractsContractAddressStoreRawParamsWithHTTPClient(client *http.Client) *GetWasmContractsContractAddressStoreRawParams {
	return &GetWasmContractsContractAddressStoreRawParams{
		HTTPClient: client,
	}
}

/* GetWasmContractsContractAddressStoreRawParams contains all the parameters to send to the API endpoint
   for the get wasm contracts contract address store raw operation.

   Typically these are written to a http.Request.
*/
type GetWasmContractsContractAddressStoreRawParams struct {

	/* ContractAddress.

	   contract address you want to lookup
	*/
	ContractAddress string

	/* Key.

	   base64 encoded raw key to access
	*/
	Key string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get wasm contracts contract address store raw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWasmContractsContractAddressStoreRawParams) WithDefaults() *GetWasmContractsContractAddressStoreRawParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get wasm contracts contract address store raw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWasmContractsContractAddressStoreRawParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get wasm contracts contract address store raw params
func (o *GetWasmContractsContractAddressStoreRawParams) WithTimeout(timeout time.Duration) *GetWasmContractsContractAddressStoreRawParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wasm contracts contract address store raw params
func (o *GetWasmContractsContractAddressStoreRawParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wasm contracts contract address store raw params
func (o *GetWasmContractsContractAddressStoreRawParams) WithContext(ctx context.Context) *GetWasmContractsContractAddressStoreRawParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wasm contracts contract address store raw params
func (o *GetWasmContractsContractAddressStoreRawParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wasm contracts contract address store raw params
func (o *GetWasmContractsContractAddressStoreRawParams) WithHTTPClient(client *http.Client) *GetWasmContractsContractAddressStoreRawParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wasm contracts contract address store raw params
func (o *GetWasmContractsContractAddressStoreRawParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContractAddress adds the contractAddress to the get wasm contracts contract address store raw params
func (o *GetWasmContractsContractAddressStoreRawParams) WithContractAddress(contractAddress string) *GetWasmContractsContractAddressStoreRawParams {
	o.SetContractAddress(contractAddress)
	return o
}

// SetContractAddress adds the contractAddress to the get wasm contracts contract address store raw params
func (o *GetWasmContractsContractAddressStoreRawParams) SetContractAddress(contractAddress string) {
	o.ContractAddress = contractAddress
}

// WithKey adds the key to the get wasm contracts contract address store raw params
func (o *GetWasmContractsContractAddressStoreRawParams) WithKey(key string) *GetWasmContractsContractAddressStoreRawParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the get wasm contracts contract address store raw params
func (o *GetWasmContractsContractAddressStoreRawParams) SetKey(key string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *GetWasmContractsContractAddressStoreRawParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contractAddress
	if err := r.SetPathParam("contractAddress", o.ContractAddress); err != nil {
		return err
	}

	// query param key
	qrKey := o.Key
	qKey := qrKey
	if qKey != "" {

		if err := r.SetQueryParam("key", qKey); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
