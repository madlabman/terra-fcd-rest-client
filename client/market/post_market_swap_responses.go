// Code generated by go-swagger; DO NOT EDIT.

package market

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
	"github.com/go-openapi/validate"
)

// PostMarketSwapReader is a Reader for the PostMarketSwap structure.
type PostMarketSwapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMarketSwapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostMarketSwapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostMarketSwapInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostMarketSwapOK creates a PostMarketSwapOK with default headers values
func NewPostMarketSwapOK() *PostMarketSwapOK {
	return &PostMarketSwapOK{}
}

/* PostMarketSwapOK describes a response with status code 200, with default header values.

OK
*/
type PostMarketSwapOK struct {
	Payload *PostMarketSwapOKBody
}

func (o *PostMarketSwapOK) Error() string {
	return fmt.Sprintf("[POST /market/swap][%d] postMarketSwapOK  %+v", 200, o.Payload)
}
func (o *PostMarketSwapOK) GetPayload() *PostMarketSwapOKBody {
	return o.Payload
}

func (o *PostMarketSwapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostMarketSwapOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMarketSwapInternalServerError creates a PostMarketSwapInternalServerError with default headers values
func NewPostMarketSwapInternalServerError() *PostMarketSwapInternalServerError {
	return &PostMarketSwapInternalServerError{}
}

/* PostMarketSwapInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostMarketSwapInternalServerError struct {
}

func (o *PostMarketSwapInternalServerError) Error() string {
	return fmt.Sprintf("[POST /market/swap][%d] postMarketSwapInternalServerError ", 500)
}

func (o *PostMarketSwapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostMarketSwapBody post market swap body
swagger:model PostMarketSwapBody
*/
type PostMarketSwapBody struct {

	// ask denom
	// Example: uluna
	// Required: true
	AskDenom *string `json:"ask_denom"`

	// base req
	// Required: true
	BaseReq *PostMarketSwapParamsBodyBaseReq `json:"base_req"`

	// offer coin
	// Required: true
	OfferCoin *PostMarketSwapParamsBodyOfferCoin `json:"offer_coin"`

	// the `receiver` field can be skipped when the receiver is trader
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	Receiver string `json:"receiver,omitempty"`
}

// Validate validates this post market swap body
func (o *PostMarketSwapBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAskDenom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOfferCoin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostMarketSwapBody) validateAskDenom(formats strfmt.Registry) error {

	if err := validate.Required("Swap coin request body"+"."+"ask_denom", "body", o.AskDenom); err != nil {
		return err
	}

	return nil
}

func (o *PostMarketSwapBody) validateBaseReq(formats strfmt.Registry) error {

	if err := validate.Required("Swap coin request body"+"."+"base_req", "body", o.BaseReq); err != nil {
		return err
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Swap coin request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Swap coin request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *PostMarketSwapBody) validateOfferCoin(formats strfmt.Registry) error {

	if err := validate.Required("Swap coin request body"+"."+"offer_coin", "body", o.OfferCoin); err != nil {
		return err
	}

	if o.OfferCoin != nil {
		if err := o.OfferCoin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Swap coin request body" + "." + "offer_coin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Swap coin request body" + "." + "offer_coin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post market swap body based on the context it is used
func (o *PostMarketSwapBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBaseReq(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOfferCoin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostMarketSwapBody) contextValidateBaseReq(ctx context.Context, formats strfmt.Registry) error {

	if o.BaseReq != nil {
		if err := o.BaseReq.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Swap coin request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Swap coin request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *PostMarketSwapBody) contextValidateOfferCoin(ctx context.Context, formats strfmt.Registry) error {

	if o.OfferCoin != nil {
		if err := o.OfferCoin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Swap coin request body" + "." + "offer_coin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Swap coin request body" + "." + "offer_coin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostMarketSwapBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMarketSwapBody) UnmarshalBinary(b []byte) error {
	var res PostMarketSwapBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostMarketSwapOKBody post market swap o k body
swagger:model PostMarketSwapOKBody
*/
type PostMarketSwapOKBody struct {

	// fee
	Fee *PostMarketSwapOKBodyFee `json:"fee,omitempty"`

	// memo
	Memo string `json:"memo,omitempty"`

	// msg
	Msg []string `json:"msg"`

	// signature
	Signature *PostMarketSwapOKBodySignature `json:"signature,omitempty"`
}

// Validate validates this post market swap o k body
func (o *PostMarketSwapOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostMarketSwapOKBody) validateFee(formats strfmt.Registry) error {
	if swag.IsZero(o.Fee) { // not required
		return nil
	}

	if o.Fee != nil {
		if err := o.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postMarketSwapOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postMarketSwapOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostMarketSwapOKBody) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(o.Signature) { // not required
		return nil
	}

	if o.Signature != nil {
		if err := o.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postMarketSwapOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postMarketSwapOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post market swap o k body based on the context it is used
func (o *PostMarketSwapOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostMarketSwapOKBody) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if o.Fee != nil {
		if err := o.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postMarketSwapOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postMarketSwapOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostMarketSwapOKBody) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if o.Signature != nil {
		if err := o.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postMarketSwapOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postMarketSwapOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostMarketSwapOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMarketSwapOKBody) UnmarshalBinary(b []byte) error {
	var res PostMarketSwapOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostMarketSwapOKBodyFee post market swap o k body fee
swagger:model PostMarketSwapOKBodyFee
*/
type PostMarketSwapOKBodyFee struct {

	// amount
	Amount []*PostMarketSwapOKBodyFeeAmountItems0 `json:"amount"`

	// gas
	Gas string `json:"gas,omitempty"`
}

// Validate validates this post market swap o k body fee
func (o *PostMarketSwapOKBodyFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostMarketSwapOKBodyFee) validateAmount(formats strfmt.Registry) error {
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
					return ve.ValidateName("postMarketSwapOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postMarketSwapOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post market swap o k body fee based on the context it is used
func (o *PostMarketSwapOKBodyFee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostMarketSwapOKBodyFee) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Amount); i++ {

		if o.Amount[i] != nil {
			if err := o.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postMarketSwapOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postMarketSwapOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostMarketSwapOKBodyFee) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMarketSwapOKBodyFee) UnmarshalBinary(b []byte) error {
	var res PostMarketSwapOKBodyFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostMarketSwapOKBodyFeeAmountItems0 post market swap o k body fee amount items0
swagger:model PostMarketSwapOKBodyFeeAmountItems0
*/
type PostMarketSwapOKBodyFeeAmountItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post market swap o k body fee amount items0
func (o *PostMarketSwapOKBodyFeeAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post market swap o k body fee amount items0 based on context it is used
func (o *PostMarketSwapOKBodyFeeAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostMarketSwapOKBodyFeeAmountItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMarketSwapOKBodyFeeAmountItems0) UnmarshalBinary(b []byte) error {
	var res PostMarketSwapOKBodyFeeAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostMarketSwapOKBodySignature post market swap o k body signature
swagger:model PostMarketSwapOKBodySignature
*/
type PostMarketSwapOKBodySignature struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// pub key
	PubKey *PostMarketSwapOKBodySignaturePubKey `json:"pub_key,omitempty"`

	// sequence
	// Example: 0
	Sequence string `json:"sequence,omitempty"`

	// signature
	// Example: MEUCIQD02fsDPra8MtbRsyB1w7bqTM55Wu138zQbFcWx4+CFyAIge5WNPfKIuvzBZ69MyqHsqD8S1IwiEp+iUb6VSdtlpgY=
	Signature string `json:"signature,omitempty"`
}

// Validate validates this post market swap o k body signature
func (o *PostMarketSwapOKBodySignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostMarketSwapOKBodySignature) validatePubKey(formats strfmt.Registry) error {
	if swag.IsZero(o.PubKey) { // not required
		return nil
	}

	if o.PubKey != nil {
		if err := o.PubKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postMarketSwapOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postMarketSwapOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post market swap o k body signature based on the context it is used
func (o *PostMarketSwapOKBodySignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePubKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostMarketSwapOKBodySignature) contextValidatePubKey(ctx context.Context, formats strfmt.Registry) error {

	if o.PubKey != nil {
		if err := o.PubKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postMarketSwapOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postMarketSwapOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostMarketSwapOKBodySignature) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMarketSwapOKBodySignature) UnmarshalBinary(b []byte) error {
	var res PostMarketSwapOKBodySignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostMarketSwapOKBodySignaturePubKey post market swap o k body signature pub key
swagger:model PostMarketSwapOKBodySignaturePubKey
*/
type PostMarketSwapOKBodySignaturePubKey struct {

	// type
	// Example: tendermint/PubKeySecp256k1
	Type string `json:"type,omitempty"`

	// value
	// Example: Avz04VhtKJh8ACCVzlI8aTosGy0ikFXKIVHQ3jKMrosH
	Value string `json:"value,omitempty"`
}

// Validate validates this post market swap o k body signature pub key
func (o *PostMarketSwapOKBodySignaturePubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post market swap o k body signature pub key based on context it is used
func (o *PostMarketSwapOKBodySignaturePubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostMarketSwapOKBodySignaturePubKey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMarketSwapOKBodySignaturePubKey) UnmarshalBinary(b []byte) error {
	var res PostMarketSwapOKBodySignaturePubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostMarketSwapParamsBodyBaseReq post market swap params body base req
swagger:model PostMarketSwapParamsBodyBaseReq
*/
type PostMarketSwapParamsBodyBaseReq struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// chain id
	// Example: Columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// fees
	Fees []*PostMarketSwapParamsBodyBaseReqFeesItems0 `json:"fees"`

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

// Validate validates this post market swap params body base req
func (o *PostMarketSwapParamsBodyBaseReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostMarketSwapParamsBodyBaseReq) validateFees(formats strfmt.Registry) error {
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
					return ve.ValidateName("Swap coin request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Swap coin request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post market swap params body base req based on the context it is used
func (o *PostMarketSwapParamsBodyBaseReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostMarketSwapParamsBodyBaseReq) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Fees); i++ {

		if o.Fees[i] != nil {
			if err := o.Fees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Swap coin request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Swap coin request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostMarketSwapParamsBodyBaseReq) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMarketSwapParamsBodyBaseReq) UnmarshalBinary(b []byte) error {
	var res PostMarketSwapParamsBodyBaseReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostMarketSwapParamsBodyBaseReqFeesItems0 post market swap params body base req fees items0
swagger:model PostMarketSwapParamsBodyBaseReqFeesItems0
*/
type PostMarketSwapParamsBodyBaseReqFeesItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post market swap params body base req fees items0
func (o *PostMarketSwapParamsBodyBaseReqFeesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post market swap params body base req fees items0 based on context it is used
func (o *PostMarketSwapParamsBodyBaseReqFeesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostMarketSwapParamsBodyBaseReqFeesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMarketSwapParamsBodyBaseReqFeesItems0) UnmarshalBinary(b []byte) error {
	var res PostMarketSwapParamsBodyBaseReqFeesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostMarketSwapParamsBodyOfferCoin post market swap params body offer coin
swagger:model PostMarketSwapParamsBodyOfferCoin
*/
type PostMarketSwapParamsBodyOfferCoin struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post market swap params body offer coin
func (o *PostMarketSwapParamsBodyOfferCoin) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post market swap params body offer coin based on context it is used
func (o *PostMarketSwapParamsBodyOfferCoin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostMarketSwapParamsBodyOfferCoin) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostMarketSwapParamsBodyOfferCoin) UnmarshalBinary(b []byte) error {
	var res PostMarketSwapParamsBodyOfferCoin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
