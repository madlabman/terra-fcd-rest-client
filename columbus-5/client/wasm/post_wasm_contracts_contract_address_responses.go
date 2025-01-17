// Code generated by go-swagger; DO NOT EDIT.

package wasm

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

// PostWasmContractsContractAddressReader is a Reader for the PostWasmContractsContractAddress structure.
type PostWasmContractsContractAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWasmContractsContractAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWasmContractsContractAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWasmContractsContractAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWasmContractsContractAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWasmContractsContractAddressOK creates a PostWasmContractsContractAddressOK with default headers values
func NewPostWasmContractsContractAddressOK() *PostWasmContractsContractAddressOK {
	return &PostWasmContractsContractAddressOK{}
}

/* PostWasmContractsContractAddressOK describes a response with status code 200, with default header values.

OK
*/
type PostWasmContractsContractAddressOK struct {
	Payload *PostWasmContractsContractAddressOKBody
}

func (o *PostWasmContractsContractAddressOK) Error() string {
	return fmt.Sprintf("[POST /wasm/contracts/{contractAddress}][%d] postWasmContractsContractAddressOK  %+v", 200, o.Payload)
}
func (o *PostWasmContractsContractAddressOK) GetPayload() *PostWasmContractsContractAddressOKBody {
	return o.Payload
}

func (o *PostWasmContractsContractAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostWasmContractsContractAddressOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWasmContractsContractAddressBadRequest creates a PostWasmContractsContractAddressBadRequest with default headers values
func NewPostWasmContractsContractAddressBadRequest() *PostWasmContractsContractAddressBadRequest {
	return &PostWasmContractsContractAddressBadRequest{}
}

/* PostWasmContractsContractAddressBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostWasmContractsContractAddressBadRequest struct {
}

func (o *PostWasmContractsContractAddressBadRequest) Error() string {
	return fmt.Sprintf("[POST /wasm/contracts/{contractAddress}][%d] postWasmContractsContractAddressBadRequest ", 400)
}

func (o *PostWasmContractsContractAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWasmContractsContractAddressInternalServerError creates a PostWasmContractsContractAddressInternalServerError with default headers values
func NewPostWasmContractsContractAddressInternalServerError() *PostWasmContractsContractAddressInternalServerError {
	return &PostWasmContractsContractAddressInternalServerError{}
}

/* PostWasmContractsContractAddressInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostWasmContractsContractAddressInternalServerError struct {
}

func (o *PostWasmContractsContractAddressInternalServerError) Error() string {
	return fmt.Sprintf("[POST /wasm/contracts/{contractAddress}][%d] postWasmContractsContractAddressInternalServerError ", 500)
}

func (o *PostWasmContractsContractAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostWasmContractsContractAddressBody post wasm contracts contract address body
swagger:model PostWasmContractsContractAddressBody
*/
type PostWasmContractsContractAddressBody struct {

	// base req
	BaseReq *PostWasmContractsContractAddressParamsBodyBaseReq `json:"base_req,omitempty"`

	// coins
	Coins []*PostWasmContractsContractAddressParamsBodyCoinsItems0 `json:"coins"`

	// exec msg
	// Example: {}
	ExecMsg string `json:"exec_msg,omitempty"`
}

// Validate validates this post wasm contracts contract address body
func (o *PostWasmContractsContractAddressBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCoins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressBody) validateBaseReq(formats strfmt.Registry) error {
	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execute contract request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("execute contract request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *PostWasmContractsContractAddressBody) validateCoins(formats strfmt.Registry) error {
	if swag.IsZero(o.Coins) { // not required
		return nil
	}

	for i := 0; i < len(o.Coins); i++ {
		if swag.IsZero(o.Coins[i]) { // not required
			continue
		}

		if o.Coins[i] != nil {
			if err := o.Coins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("execute contract request body" + "." + "coins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("execute contract request body" + "." + "coins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post wasm contracts contract address body based on the context it is used
func (o *PostWasmContractsContractAddressBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBaseReq(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateCoins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressBody) contextValidateBaseReq(ctx context.Context, formats strfmt.Registry) error {

	if o.BaseReq != nil {
		if err := o.BaseReq.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execute contract request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("execute contract request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *PostWasmContractsContractAddressBody) contextValidateCoins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Coins); i++ {

		if o.Coins[i] != nil {
			if err := o.Coins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("execute contract request body" + "." + "coins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("execute contract request body" + "." + "coins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressBody) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressOKBody post wasm contracts contract address o k body
swagger:model PostWasmContractsContractAddressOKBody
*/
type PostWasmContractsContractAddressOKBody struct {

	// fee
	Fee *PostWasmContractsContractAddressOKBodyFee `json:"fee,omitempty"`

	// memo
	Memo string `json:"memo,omitempty"`

	// msg
	Msg []string `json:"msg"`

	// signature
	Signature *PostWasmContractsContractAddressOKBodySignature `json:"signature,omitempty"`
}

// Validate validates this post wasm contracts contract address o k body
func (o *PostWasmContractsContractAddressOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostWasmContractsContractAddressOKBody) validateFee(formats strfmt.Registry) error {
	if swag.IsZero(o.Fee) { // not required
		return nil
	}

	if o.Fee != nil {
		if err := o.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostWasmContractsContractAddressOKBody) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(o.Signature) { // not required
		return nil
	}

	if o.Signature != nil {
		if err := o.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post wasm contracts contract address o k body based on the context it is used
func (o *PostWasmContractsContractAddressOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostWasmContractsContractAddressOKBody) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if o.Fee != nil {
		if err := o.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostWasmContractsContractAddressOKBody) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if o.Signature != nil {
		if err := o.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressOKBody) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressOKBodyFee post wasm contracts contract address o k body fee
swagger:model PostWasmContractsContractAddressOKBodyFee
*/
type PostWasmContractsContractAddressOKBodyFee struct {

	// amount
	Amount []*PostWasmContractsContractAddressOKBodyFeeAmountItems0 `json:"amount"`

	// gas
	Gas string `json:"gas,omitempty"`
}

// Validate validates this post wasm contracts contract address o k body fee
func (o *PostWasmContractsContractAddressOKBodyFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressOKBodyFee) validateAmount(formats strfmt.Registry) error {
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
					return ve.ValidateName("postWasmContractsContractAddressOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postWasmContractsContractAddressOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post wasm contracts contract address o k body fee based on the context it is used
func (o *PostWasmContractsContractAddressOKBodyFee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressOKBodyFee) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Amount); i++ {

		if o.Amount[i] != nil {
			if err := o.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postWasmContractsContractAddressOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postWasmContractsContractAddressOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressOKBodyFee) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressOKBodyFee) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressOKBodyFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressOKBodyFeeAmountItems0 post wasm contracts contract address o k body fee amount items0
swagger:model PostWasmContractsContractAddressOKBodyFeeAmountItems0
*/
type PostWasmContractsContractAddressOKBodyFeeAmountItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post wasm contracts contract address o k body fee amount items0
func (o *PostWasmContractsContractAddressOKBodyFeeAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post wasm contracts contract address o k body fee amount items0 based on context it is used
func (o *PostWasmContractsContractAddressOKBodyFeeAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressOKBodyFeeAmountItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressOKBodyFeeAmountItems0) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressOKBodyFeeAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressOKBodySignature post wasm contracts contract address o k body signature
swagger:model PostWasmContractsContractAddressOKBodySignature
*/
type PostWasmContractsContractAddressOKBodySignature struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// pub key
	PubKey *PostWasmContractsContractAddressOKBodySignaturePubKey `json:"pub_key,omitempty"`

	// sequence
	// Example: 0
	Sequence string `json:"sequence,omitempty"`

	// signature
	// Example: MEUCIQD02fsDPra8MtbRsyB1w7bqTM55Wu138zQbFcWx4+CFyAIge5WNPfKIuvzBZ69MyqHsqD8S1IwiEp+iUb6VSdtlpgY=
	Signature string `json:"signature,omitempty"`
}

// Validate validates this post wasm contracts contract address o k body signature
func (o *PostWasmContractsContractAddressOKBodySignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressOKBodySignature) validatePubKey(formats strfmt.Registry) error {
	if swag.IsZero(o.PubKey) { // not required
		return nil
	}

	if o.PubKey != nil {
		if err := o.PubKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post wasm contracts contract address o k body signature based on the context it is used
func (o *PostWasmContractsContractAddressOKBodySignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePubKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressOKBodySignature) contextValidatePubKey(ctx context.Context, formats strfmt.Registry) error {

	if o.PubKey != nil {
		if err := o.PubKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressOKBodySignature) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressOKBodySignature) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressOKBodySignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressOKBodySignaturePubKey post wasm contracts contract address o k body signature pub key
swagger:model PostWasmContractsContractAddressOKBodySignaturePubKey
*/
type PostWasmContractsContractAddressOKBodySignaturePubKey struct {

	// type
	// Example: tendermint/PubKeySecp256k1
	Type string `json:"type,omitempty"`

	// value
	// Example: Avz04VhtKJh8ACCVzlI8aTosGy0ikFXKIVHQ3jKMrosH
	Value string `json:"value,omitempty"`
}

// Validate validates this post wasm contracts contract address o k body signature pub key
func (o *PostWasmContractsContractAddressOKBodySignaturePubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post wasm contracts contract address o k body signature pub key based on context it is used
func (o *PostWasmContractsContractAddressOKBodySignaturePubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressOKBodySignaturePubKey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressOKBodySignaturePubKey) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressOKBodySignaturePubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressParamsBodyBaseReq post wasm contracts contract address params body base req
swagger:model PostWasmContractsContractAddressParamsBodyBaseReq
*/
type PostWasmContractsContractAddressParamsBodyBaseReq struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// chain id
	// Example: Columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// fees
	Fees []*PostWasmContractsContractAddressParamsBodyBaseReqFeesItems0 `json:"fees"`

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

// Validate validates this post wasm contracts contract address params body base req
func (o *PostWasmContractsContractAddressParamsBodyBaseReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressParamsBodyBaseReq) validateFees(formats strfmt.Registry) error {
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
					return ve.ValidateName("execute contract request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("execute contract request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post wasm contracts contract address params body base req based on the context it is used
func (o *PostWasmContractsContractAddressParamsBodyBaseReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressParamsBodyBaseReq) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Fees); i++ {

		if o.Fees[i] != nil {
			if err := o.Fees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("execute contract request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("execute contract request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressParamsBodyBaseReq) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressParamsBodyBaseReq) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressParamsBodyBaseReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressParamsBodyBaseReqFeesItems0 post wasm contracts contract address params body base req fees items0
swagger:model PostWasmContractsContractAddressParamsBodyBaseReqFeesItems0
*/
type PostWasmContractsContractAddressParamsBodyBaseReqFeesItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post wasm contracts contract address params body base req fees items0
func (o *PostWasmContractsContractAddressParamsBodyBaseReqFeesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post wasm contracts contract address params body base req fees items0 based on context it is used
func (o *PostWasmContractsContractAddressParamsBodyBaseReqFeesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressParamsBodyBaseReqFeesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressParamsBodyBaseReqFeesItems0) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressParamsBodyBaseReqFeesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressParamsBodyCoinsItems0 post wasm contracts contract address params body coins items0
swagger:model PostWasmContractsContractAddressParamsBodyCoinsItems0
*/
type PostWasmContractsContractAddressParamsBodyCoinsItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post wasm contracts contract address params body coins items0
func (o *PostWasmContractsContractAddressParamsBodyCoinsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post wasm contracts contract address params body coins items0 based on context it is used
func (o *PostWasmContractsContractAddressParamsBodyCoinsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressParamsBodyCoinsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressParamsBodyCoinsItems0) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressParamsBodyCoinsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
