// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new oracle API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for oracle API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetOracleDenomsActives(params *GetOracleDenomsActivesParams, opts ...ClientOption) (*GetOracleDenomsActivesOK, error)

	GetOracleDenomsDenomExchangeRate(params *GetOracleDenomsDenomExchangeRateParams, opts ...ClientOption) (*GetOracleDenomsDenomExchangeRateOK, error)

	GetOracleDenomsExchangeRates(params *GetOracleDenomsExchangeRatesParams, opts ...ClientOption) (*GetOracleDenomsExchangeRatesOK, error)

	GetOracleParameters(params *GetOracleParametersParams, opts ...ClientOption) (*GetOracleParametersOK, error)

	GetOracleVotersValidatorAggregatePrevote(params *GetOracleVotersValidatorAggregatePrevoteParams, opts ...ClientOption) (*GetOracleVotersValidatorAggregatePrevoteOK, error)

	GetOracleVotersValidatorAggregateVote(params *GetOracleVotersValidatorAggregateVoteParams, opts ...ClientOption) (*GetOracleVotersValidatorAggregateVoteOK, error)

	GetOracleVotersValidatorFeeder(params *GetOracleVotersValidatorFeederParams, opts ...ClientOption) (*GetOracleVotersValidatorFeederOK, error)

	GetOracleVotersValidatorMiss(params *GetOracleVotersValidatorMissParams, opts ...ClientOption) (*GetOracleVotersValidatorMissOK, error)

	PostOracleVotersValidatorAggregatePrevote(params *PostOracleVotersValidatorAggregatePrevoteParams, opts ...ClientOption) (*PostOracleVotersValidatorAggregatePrevoteOK, error)

	PostOracleVotersValidatorAggregateVote(params *PostOracleVotersValidatorAggregateVoteParams, opts ...ClientOption) (*PostOracleVotersValidatorAggregateVoteOK, error)

	PostOracleVotersValidatorFeeder(params *PostOracleVotersValidatorFeederParams, opts ...ClientOption) (*PostOracleVotersValidatorFeederOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetOracleDenomsActives gets all activated denoms
*/
func (a *Client) GetOracleDenomsActives(params *GetOracleDenomsActivesParams, opts ...ClientOption) (*GetOracleDenomsActivesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleDenomsActivesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOracleDenomsActives",
		Method:             "GET",
		PathPattern:        "/oracle/denoms/actives",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOracleDenomsActivesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleDenomsActivesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleDenomsActives: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleDenomsDenomExchangeRate gets the current effective exchange rate in luna for the asset
*/
func (a *Client) GetOracleDenomsDenomExchangeRate(params *GetOracleDenomsDenomExchangeRateParams, opts ...ClientOption) (*GetOracleDenomsDenomExchangeRateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleDenomsDenomExchangeRateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOracleDenomsDenomExchangeRate",
		Method:             "GET",
		PathPattern:        "/oracle/denoms/{denom}/exchange_rate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOracleDenomsDenomExchangeRateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleDenomsDenomExchangeRateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleDenomsDenomExchangeRate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleDenomsExchangeRates gets all activated exchange rates
*/
func (a *Client) GetOracleDenomsExchangeRates(params *GetOracleDenomsExchangeRatesParams, opts ...ClientOption) (*GetOracleDenomsExchangeRatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleDenomsExchangeRatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOracleDenomsExchangeRates",
		Method:             "GET",
		PathPattern:        "/oracle/denoms/exchange_rates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOracleDenomsExchangeRatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleDenomsExchangeRatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleDenomsExchangeRates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleParameters gets oracle params
*/
func (a *Client) GetOracleParameters(params *GetOracleParametersParams, opts ...ClientOption) (*GetOracleParametersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleParametersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOracleParameters",
		Method:             "GET",
		PathPattern:        "/oracle/parameters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOracleParametersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleParametersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleParameters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleVotersValidatorAggregatePrevote gets the currently outstanding aggregate exchange rate oracle prevote of a validator
*/
func (a *Client) GetOracleVotersValidatorAggregatePrevote(params *GetOracleVotersValidatorAggregatePrevoteParams, opts ...ClientOption) (*GetOracleVotersValidatorAggregatePrevoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleVotersValidatorAggregatePrevoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOracleVotersValidatorAggregatePrevote",
		Method:             "GET",
		PathPattern:        "/oracle/voters/{validator}/aggregate_prevote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOracleVotersValidatorAggregatePrevoteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleVotersValidatorAggregatePrevoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleVotersValidatorAggregatePrevote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleVotersValidatorAggregateVote gets the currently outstanding aggregate exchange rate oracle vote of a validator
*/
func (a *Client) GetOracleVotersValidatorAggregateVote(params *GetOracleVotersValidatorAggregateVoteParams, opts ...ClientOption) (*GetOracleVotersValidatorAggregateVoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleVotersValidatorAggregateVoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOracleVotersValidatorAggregateVote",
		Method:             "GET",
		PathPattern:        "/oracle/voters/{validator}/aggregate_vote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOracleVotersValidatorAggregateVoteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleVotersValidatorAggregateVoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleVotersValidatorAggregateVote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleVotersValidatorFeeder gets delegated oracle feeder of a validator
*/
func (a *Client) GetOracleVotersValidatorFeeder(params *GetOracleVotersValidatorFeederParams, opts ...ClientOption) (*GetOracleVotersValidatorFeederOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleVotersValidatorFeederParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOracleVotersValidatorFeeder",
		Method:             "GET",
		PathPattern:        "/oracle/voters/{validator}/feeder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOracleVotersValidatorFeederReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleVotersValidatorFeederOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleVotersValidatorFeeder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOracleVotersValidatorMiss gets the number of vote periods missed in this oracle slash window
*/
func (a *Client) GetOracleVotersValidatorMiss(params *GetOracleVotersValidatorMissParams, opts ...ClientOption) (*GetOracleVotersValidatorMissOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOracleVotersValidatorMissParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOracleVotersValidatorMiss",
		Method:             "GET",
		PathPattern:        "/oracle/voters/{validator}/miss",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOracleVotersValidatorMissReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOracleVotersValidatorMissOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOracleVotersValidatorMiss: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOracleVotersValidatorAggregatePrevote generates oracle aggregate exchange rate prevote message containing a hash
*/
func (a *Client) PostOracleVotersValidatorAggregatePrevote(params *PostOracleVotersValidatorAggregatePrevoteParams, opts ...ClientOption) (*PostOracleVotersValidatorAggregatePrevoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOracleVotersValidatorAggregatePrevoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostOracleVotersValidatorAggregatePrevote",
		Method:             "POST",
		PathPattern:        "/oracle/voters/{validator}/aggregate_prevote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostOracleVotersValidatorAggregatePrevoteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOracleVotersValidatorAggregatePrevoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOracleVotersValidatorAggregatePrevote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOracleVotersValidatorAggregateVote generates oracle aggregate exchange rate vote message containing exchange rates and salt to prove the aggregate prevote
*/
func (a *Client) PostOracleVotersValidatorAggregateVote(params *PostOracleVotersValidatorAggregateVoteParams, opts ...ClientOption) (*PostOracleVotersValidatorAggregateVoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOracleVotersValidatorAggregateVoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostOracleVotersValidatorAggregateVote",
		Method:             "POST",
		PathPattern:        "/oracle/voters/{validator}/aggregate_vote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostOracleVotersValidatorAggregateVoteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOracleVotersValidatorAggregateVoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOracleVotersValidatorAggregateVote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOracleVotersValidatorFeeder generates oracle feeder right delegation message
*/
func (a *Client) PostOracleVotersValidatorFeeder(params *PostOracleVotersValidatorFeederParams, opts ...ClientOption) (*PostOracleVotersValidatorFeederOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOracleVotersValidatorFeederParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostOracleVotersValidatorFeeder",
		Method:             "POST",
		PathPattern:        "/oracle/voters/{validator}/feeder",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostOracleVotersValidatorFeederReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOracleVotersValidatorFeederOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOracleVotersValidatorFeeder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}