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

// NewUnreceivedPacketsParams creates a new UnreceivedPacketsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnreceivedPacketsParams() *UnreceivedPacketsParams {
	return &UnreceivedPacketsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnreceivedPacketsParamsWithTimeout creates a new UnreceivedPacketsParams object
// with the ability to set a timeout on a request.
func NewUnreceivedPacketsParamsWithTimeout(timeout time.Duration) *UnreceivedPacketsParams {
	return &UnreceivedPacketsParams{
		timeout: timeout,
	}
}

// NewUnreceivedPacketsParamsWithContext creates a new UnreceivedPacketsParams object
// with the ability to set a context for a request.
func NewUnreceivedPacketsParamsWithContext(ctx context.Context) *UnreceivedPacketsParams {
	return &UnreceivedPacketsParams{
		Context: ctx,
	}
}

// NewUnreceivedPacketsParamsWithHTTPClient creates a new UnreceivedPacketsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnreceivedPacketsParamsWithHTTPClient(client *http.Client) *UnreceivedPacketsParams {
	return &UnreceivedPacketsParams{
		HTTPClient: client,
	}
}

/* UnreceivedPacketsParams contains all the parameters to send to the API endpoint
   for the unreceived packets operation.

   Typically these are written to a http.Request.
*/
type UnreceivedPacketsParams struct {

	/* ChannelID.

	   channel unique identifier
	*/
	ChannelID string

	/* PacketCommitmentSequences.

	   list of packet sequences
	*/
	PacketCommitmentSequences []string

	/* PortID.

	   port unique identifier
	*/
	PortID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unreceived packets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnreceivedPacketsParams) WithDefaults() *UnreceivedPacketsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unreceived packets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnreceivedPacketsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unreceived packets params
func (o *UnreceivedPacketsParams) WithTimeout(timeout time.Duration) *UnreceivedPacketsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unreceived packets params
func (o *UnreceivedPacketsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unreceived packets params
func (o *UnreceivedPacketsParams) WithContext(ctx context.Context) *UnreceivedPacketsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unreceived packets params
func (o *UnreceivedPacketsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unreceived packets params
func (o *UnreceivedPacketsParams) WithHTTPClient(client *http.Client) *UnreceivedPacketsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unreceived packets params
func (o *UnreceivedPacketsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the unreceived packets params
func (o *UnreceivedPacketsParams) WithChannelID(channelID string) *UnreceivedPacketsParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the unreceived packets params
func (o *UnreceivedPacketsParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WithPacketCommitmentSequences adds the packetCommitmentSequences to the unreceived packets params
func (o *UnreceivedPacketsParams) WithPacketCommitmentSequences(packetCommitmentSequences []string) *UnreceivedPacketsParams {
	o.SetPacketCommitmentSequences(packetCommitmentSequences)
	return o
}

// SetPacketCommitmentSequences adds the packetCommitmentSequences to the unreceived packets params
func (o *UnreceivedPacketsParams) SetPacketCommitmentSequences(packetCommitmentSequences []string) {
	o.PacketCommitmentSequences = packetCommitmentSequences
}

// WithPortID adds the portID to the unreceived packets params
func (o *UnreceivedPacketsParams) WithPortID(portID string) *UnreceivedPacketsParams {
	o.SetPortID(portID)
	return o
}

// SetPortID adds the portId to the unreceived packets params
func (o *UnreceivedPacketsParams) SetPortID(portID string) {
	o.PortID = portID
}

// WriteToRequest writes these params to a swagger request
func (o *UnreceivedPacketsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channel_id
	if err := r.SetPathParam("channel_id", o.ChannelID); err != nil {
		return err
	}

	if o.PacketCommitmentSequences != nil {

		// binding items for packet_commitment_sequences
		joinedPacketCommitmentSequences := o.bindParamPacketCommitmentSequences(reg)

		// path array param packet_commitment_sequences
		// SetPathParam does not support variadic arguments, since we used JoinByFormat
		// we can send the first item in the array as it's all the items of the previous
		// array joined together
		if len(joinedPacketCommitmentSequences) > 0 {
			if err := r.SetPathParam("packet_commitment_sequences", joinedPacketCommitmentSequences[0]); err != nil {
				return err
			}
		}
	}

	// path param port_id
	if err := r.SetPathParam("port_id", o.PortID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUnreceivedPackets binds the parameter packet_commitment_sequences
func (o *UnreceivedPacketsParams) bindParamPacketCommitmentSequences(formats strfmt.Registry) []string {
	packetCommitmentSequencesIR := o.PacketCommitmentSequences

	var packetCommitmentSequencesIC []string
	for _, packetCommitmentSequencesIIR := range packetCommitmentSequencesIR { // explode []string

		packetCommitmentSequencesIIV := packetCommitmentSequencesIIR // string as string
		packetCommitmentSequencesIC = append(packetCommitmentSequencesIC, packetCommitmentSequencesIIV)
	}

	// items.CollectionFormat: "csv"
	packetCommitmentSequencesIS := swag.JoinByFormat(packetCommitmentSequencesIC, "csv")

	return packetCommitmentSequencesIS
}
