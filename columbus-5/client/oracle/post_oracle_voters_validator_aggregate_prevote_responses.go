// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostOracleVotersValidatorAggregatePrevoteReader is a Reader for the PostOracleVotersValidatorAggregatePrevote structure.
type PostOracleVotersValidatorAggregatePrevoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOracleVotersValidatorAggregatePrevoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOracleVotersValidatorAggregatePrevoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOracleVotersValidatorAggregatePrevoteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOracleVotersValidatorAggregatePrevoteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOracleVotersValidatorAggregatePrevoteOK creates a PostOracleVotersValidatorAggregatePrevoteOK with default headers values
func NewPostOracleVotersValidatorAggregatePrevoteOK() *PostOracleVotersValidatorAggregatePrevoteOK {
	return &PostOracleVotersValidatorAggregatePrevoteOK{}
}

/* PostOracleVotersValidatorAggregatePrevoteOK describes a response with status code 200, with default header values.

OK
*/
type PostOracleVotersValidatorAggregatePrevoteOK struct {
	Payload *PostOracleVotersValidatorAggregatePrevoteOKBody
}

func (o *PostOracleVotersValidatorAggregatePrevoteOK) Error() string {
	return fmt.Sprintf("[POST /oracle/voters/{validator}/aggregate_prevote][%d] postOracleVotersValidatorAggregatePrevoteOK  %+v", 200, o.Payload)
}
func (o *PostOracleVotersValidatorAggregatePrevoteOK) GetPayload() *PostOracleVotersValidatorAggregatePrevoteOKBody {
	return o.Payload
}

func (o *PostOracleVotersValidatorAggregatePrevoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostOracleVotersValidatorAggregatePrevoteOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOracleVotersValidatorAggregatePrevoteBadRequest creates a PostOracleVotersValidatorAggregatePrevoteBadRequest with default headers values
func NewPostOracleVotersValidatorAggregatePrevoteBadRequest() *PostOracleVotersValidatorAggregatePrevoteBadRequest {
	return &PostOracleVotersValidatorAggregatePrevoteBadRequest{}
}

/* PostOracleVotersValidatorAggregatePrevoteBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostOracleVotersValidatorAggregatePrevoteBadRequest struct {
}

func (o *PostOracleVotersValidatorAggregatePrevoteBadRequest) Error() string {
	return fmt.Sprintf("[POST /oracle/voters/{validator}/aggregate_prevote][%d] postOracleVotersValidatorAggregatePrevoteBadRequest ", 400)
}

func (o *PostOracleVotersValidatorAggregatePrevoteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostOracleVotersValidatorAggregatePrevoteInternalServerError creates a PostOracleVotersValidatorAggregatePrevoteInternalServerError with default headers values
func NewPostOracleVotersValidatorAggregatePrevoteInternalServerError() *PostOracleVotersValidatorAggregatePrevoteInternalServerError {
	return &PostOracleVotersValidatorAggregatePrevoteInternalServerError{}
}

/* PostOracleVotersValidatorAggregatePrevoteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostOracleVotersValidatorAggregatePrevoteInternalServerError struct {
}

func (o *PostOracleVotersValidatorAggregatePrevoteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /oracle/voters/{validator}/aggregate_prevote][%d] postOracleVotersValidatorAggregatePrevoteInternalServerError ", 500)
}

func (o *PostOracleVotersValidatorAggregatePrevoteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostOracleVotersValidatorAggregatePrevoteBody post oracle voters validator aggregate prevote body
swagger:model PostOracleVotersValidatorAggregatePrevoteBody
*/
type PostOracleVotersValidatorAggregatePrevoteBody struct {

	// base req
	BaseReq *PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReq `json:"base_req,omitempty"`

	// exchange rates of Luna in denom currencies are to make aggregate prevote hash; this field is required to submit prevote in case absence of hash
	// Example: 1000.0ukrw,1.2uusd,0.99usdr
	ExchangeRates string `json:"exchange_rates,omitempty"`

	// hex string; hash of next vote; empty == skip prevote
	// Example: 061bf1e27dfff121f40c826e593c8a28ec299a02
	Hash string `json:"hash,omitempty"`

	// salt is to make prevote hash; this field is required to submit prevote in case  absence of hash
	// Example: abcd
	Salt string `json:"salt,omitempty"`
}

// Validate validates this post oracle voters validator aggregate prevote body
func (o *PostOracleVotersValidatorAggregatePrevoteBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteBody) validateBaseReq(formats strfmt.Registry) error {
	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Aggregate Prevote request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Aggregate Prevote request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post oracle voters validator aggregate prevote body based on the context it is used
func (o *PostOracleVotersValidatorAggregatePrevoteBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBaseReq(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteBody) contextValidateBaseReq(ctx context.Context, formats strfmt.Registry) error {

	if o.BaseReq != nil {
		if err := o.BaseReq.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Aggregate Prevote request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Aggregate Prevote request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteBody) UnmarshalBinary(b []byte) error {
	var res PostOracleVotersValidatorAggregatePrevoteBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostOracleVotersValidatorAggregatePrevoteOKBody post oracle voters validator aggregate prevote o k body
swagger:model PostOracleVotersValidatorAggregatePrevoteOKBody
*/
type PostOracleVotersValidatorAggregatePrevoteOKBody struct {

	// fee
	Fee *PostOracleVotersValidatorAggregatePrevoteOKBodyFee `json:"fee,omitempty"`

	// memo
	Memo string `json:"memo,omitempty"`

	// msg
	Msg []string `json:"msg"`

	// signature
	Signature *PostOracleVotersValidatorAggregatePrevoteOKBodySignature `json:"signature,omitempty"`
}

// Validate validates this post oracle voters validator aggregate prevote o k body
func (o *PostOracleVotersValidatorAggregatePrevoteOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteOKBody) validateFee(formats strfmt.Registry) error {
	if swag.IsZero(o.Fee) { // not required
		return nil
	}

	if o.Fee != nil {
		if err := o.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteOKBody) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(o.Signature) { // not required
		return nil
	}

	if o.Signature != nil {
		if err := o.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post oracle voters validator aggregate prevote o k body based on the context it is used
func (o *PostOracleVotersValidatorAggregatePrevoteOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteOKBody) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if o.Fee != nil {
		if err := o.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteOKBody) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if o.Signature != nil {
		if err := o.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteOKBody) UnmarshalBinary(b []byte) error {
	var res PostOracleVotersValidatorAggregatePrevoteOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostOracleVotersValidatorAggregatePrevoteOKBodyFee post oracle voters validator aggregate prevote o k body fee
swagger:model PostOracleVotersValidatorAggregatePrevoteOKBodyFee
*/
type PostOracleVotersValidatorAggregatePrevoteOKBodyFee struct {

	// amount
	Amount []*PostOracleVotersValidatorAggregatePrevoteOKBodyFeeAmountItems0 `json:"amount"`

	// gas
	Gas string `json:"gas,omitempty"`
}

// Validate validates this post oracle voters validator aggregate prevote o k body fee
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodyFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteOKBodyFee) validateAmount(formats strfmt.Registry) error {
	if swag.IsZero(o.Amount) { // not required
		return nil
	}

	for i := 0; i < len(o.Amount); i++ {
		if swag.IsZero(o.Amount[i]) { // not required
			continue
		}

		if o.Amount[i] != nil {
			if err := o.Amount[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post oracle voters validator aggregate prevote o k body fee based on the context it is used
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodyFee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteOKBodyFee) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Amount); i++ {

		if o.Amount[i] != nil {
			if err := o.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodyFee) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodyFee) UnmarshalBinary(b []byte) error {
	var res PostOracleVotersValidatorAggregatePrevoteOKBodyFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostOracleVotersValidatorAggregatePrevoteOKBodyFeeAmountItems0 post oracle voters validator aggregate prevote o k body fee amount items0
swagger:model PostOracleVotersValidatorAggregatePrevoteOKBodyFeeAmountItems0
*/
type PostOracleVotersValidatorAggregatePrevoteOKBodyFeeAmountItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post oracle voters validator aggregate prevote o k body fee amount items0
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodyFeeAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post oracle voters validator aggregate prevote o k body fee amount items0 based on context it is used
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodyFeeAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodyFeeAmountItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodyFeeAmountItems0) UnmarshalBinary(b []byte) error {
	var res PostOracleVotersValidatorAggregatePrevoteOKBodyFeeAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostOracleVotersValidatorAggregatePrevoteOKBodySignature post oracle voters validator aggregate prevote o k body signature
swagger:model PostOracleVotersValidatorAggregatePrevoteOKBodySignature
*/
type PostOracleVotersValidatorAggregatePrevoteOKBodySignature struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// pub key
	PubKey *PostOracleVotersValidatorAggregatePrevoteOKBodySignaturePubKey `json:"pub_key,omitempty"`

	// sequence
	// Example: 0
	Sequence string `json:"sequence,omitempty"`

	// signature
	// Example: MEUCIQD02fsDPra8MtbRsyB1w7bqTM55Wu138zQbFcWx4+CFyAIge5WNPfKIuvzBZ69MyqHsqD8S1IwiEp+iUb6VSdtlpgY=
	Signature string `json:"signature,omitempty"`
}

// Validate validates this post oracle voters validator aggregate prevote o k body signature
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodySignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteOKBodySignature) validatePubKey(formats strfmt.Registry) error {
	if swag.IsZero(o.PubKey) { // not required
		return nil
	}

	if o.PubKey != nil {
		if err := o.PubKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post oracle voters validator aggregate prevote o k body signature based on the context it is used
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodySignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePubKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteOKBodySignature) contextValidatePubKey(ctx context.Context, formats strfmt.Registry) error {

	if o.PubKey != nil {
		if err := o.PubKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postOracleVotersValidatorAggregatePrevoteOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodySignature) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodySignature) UnmarshalBinary(b []byte) error {
	var res PostOracleVotersValidatorAggregatePrevoteOKBodySignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostOracleVotersValidatorAggregatePrevoteOKBodySignaturePubKey post oracle voters validator aggregate prevote o k body signature pub key
swagger:model PostOracleVotersValidatorAggregatePrevoteOKBodySignaturePubKey
*/
type PostOracleVotersValidatorAggregatePrevoteOKBodySignaturePubKey struct {

	// type
	// Example: tendermint/PubKeySecp256k1
	Type string `json:"type,omitempty"`

	// value
	// Example: Avz04VhtKJh8ACCVzlI8aTosGy0ikFXKIVHQ3jKMrosH
	Value string `json:"value,omitempty"`
}

// Validate validates this post oracle voters validator aggregate prevote o k body signature pub key
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodySignaturePubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post oracle voters validator aggregate prevote o k body signature pub key based on context it is used
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodySignaturePubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodySignaturePubKey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteOKBodySignaturePubKey) UnmarshalBinary(b []byte) error {
	var res PostOracleVotersValidatorAggregatePrevoteOKBodySignaturePubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReq post oracle voters validator aggregate prevote params body base req
swagger:model PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReq
*/
type PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReq struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// chain id
	// Example: Columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// fees
	Fees []*PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReqFeesItems0 `json:"fees"`

	// Sender address or Keybase name to generate a transaction
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	From string `json:"from,omitempty"`

	// gas
	// Example: 200000
	Gas string `json:"gas,omitempty"`

	// gas adjustment
	// Example: 1.2
	GasAdjustment string `json:"gas_adjustment,omitempty"`

	// memo
	// Example: Sent via Terra Station 🚀
	Memo string `json:"memo,omitempty"`

	// sequence
	// Example: 1
	Sequence string `json:"sequence,omitempty"`

	// Estimate gas for a transaction (cannot be used in conjunction with generate_only)
	// Example: false
	Simulate bool `json:"simulate,omitempty"`
}

// Validate validates this post oracle voters validator aggregate prevote params body base req
func (o *PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReq) validateFees(formats strfmt.Registry) error {
	if swag.IsZero(o.Fees) { // not required
		return nil
	}

	for i := 0; i < len(o.Fees); i++ {
		if swag.IsZero(o.Fees[i]) { // not required
			continue
		}

		if o.Fees[i] != nil {
			if err := o.Fees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Aggregate Prevote request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Aggregate Prevote request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post oracle voters validator aggregate prevote params body base req based on the context it is used
func (o *PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReq) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Fees); i++ {

		if o.Fees[i] != nil {
			if err := o.Fees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Aggregate Prevote request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Aggregate Prevote request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReq) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReq) UnmarshalBinary(b []byte) error {
	var res PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReqFeesItems0 post oracle voters validator aggregate prevote params body base req fees items0
swagger:model PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReqFeesItems0
*/
type PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReqFeesItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post oracle voters validator aggregate prevote params body base req fees items0
func (o *PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReqFeesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post oracle voters validator aggregate prevote params body base req fees items0 based on context it is used
func (o *PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReqFeesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReqFeesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReqFeesItems0) UnmarshalBinary(b []byte) error {
	var res PostOracleVotersValidatorAggregatePrevoteParamsBodyBaseReqFeesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
