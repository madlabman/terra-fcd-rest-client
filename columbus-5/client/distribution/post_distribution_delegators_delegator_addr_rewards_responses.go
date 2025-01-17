// Code generated by go-swagger; DO NOT EDIT.

package distribution

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

// PostDistributionDelegatorsDelegatorAddrRewardsReader is a Reader for the PostDistributionDelegatorsDelegatorAddrRewards structure.
type PostDistributionDelegatorsDelegatorAddrRewardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDistributionDelegatorsDelegatorAddrRewardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDistributionDelegatorsDelegatorAddrRewardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDistributionDelegatorsDelegatorAddrRewardsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDistributionDelegatorsDelegatorAddrRewardsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDistributionDelegatorsDelegatorAddrRewardsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDistributionDelegatorsDelegatorAddrRewardsOK creates a PostDistributionDelegatorsDelegatorAddrRewardsOK with default headers values
func NewPostDistributionDelegatorsDelegatorAddrRewardsOK() *PostDistributionDelegatorsDelegatorAddrRewardsOK {
	return &PostDistributionDelegatorsDelegatorAddrRewardsOK{}
}

/* PostDistributionDelegatorsDelegatorAddrRewardsOK describes a response with status code 200, with default header values.

OK
*/
type PostDistributionDelegatorsDelegatorAddrRewardsOK struct {
	Payload *PostDistributionDelegatorsDelegatorAddrRewardsOKBody
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsOK) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/rewards][%d] postDistributionDelegatorsDelegatorAddrRewardsOK  %+v", 200, o.Payload)
}
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOK) GetPayload() *PostDistributionDelegatorsDelegatorAddrRewardsOKBody {
	return o.Payload
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostDistributionDelegatorsDelegatorAddrRewardsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDistributionDelegatorsDelegatorAddrRewardsBadRequest creates a PostDistributionDelegatorsDelegatorAddrRewardsBadRequest with default headers values
func NewPostDistributionDelegatorsDelegatorAddrRewardsBadRequest() *PostDistributionDelegatorsDelegatorAddrRewardsBadRequest {
	return &PostDistributionDelegatorsDelegatorAddrRewardsBadRequest{}
}

/* PostDistributionDelegatorsDelegatorAddrRewardsBadRequest describes a response with status code 400, with default header values.

Invalid delegator address
*/
type PostDistributionDelegatorsDelegatorAddrRewardsBadRequest struct {
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsBadRequest) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/rewards][%d] postDistributionDelegatorsDelegatorAddrRewardsBadRequest ", 400)
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDistributionDelegatorsDelegatorAddrRewardsUnauthorized creates a PostDistributionDelegatorsDelegatorAddrRewardsUnauthorized with default headers values
func NewPostDistributionDelegatorsDelegatorAddrRewardsUnauthorized() *PostDistributionDelegatorsDelegatorAddrRewardsUnauthorized {
	return &PostDistributionDelegatorsDelegatorAddrRewardsUnauthorized{}
}

/* PostDistributionDelegatorsDelegatorAddrRewardsUnauthorized describes a response with status code 401, with default header values.

Key password is wrong
*/
type PostDistributionDelegatorsDelegatorAddrRewardsUnauthorized struct {
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/rewards][%d] postDistributionDelegatorsDelegatorAddrRewardsUnauthorized ", 401)
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostDistributionDelegatorsDelegatorAddrRewardsInternalServerError creates a PostDistributionDelegatorsDelegatorAddrRewardsInternalServerError with default headers values
func NewPostDistributionDelegatorsDelegatorAddrRewardsInternalServerError() *PostDistributionDelegatorsDelegatorAddrRewardsInternalServerError {
	return &PostDistributionDelegatorsDelegatorAddrRewardsInternalServerError{}
}

/* PostDistributionDelegatorsDelegatorAddrRewardsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDistributionDelegatorsDelegatorAddrRewardsInternalServerError struct {
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /distribution/delegators/{delegatorAddr}/rewards][%d] postDistributionDelegatorsDelegatorAddrRewardsInternalServerError ", 500)
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostDistributionDelegatorsDelegatorAddrRewardsBody post distribution delegators delegator addr rewards body
swagger:model PostDistributionDelegatorsDelegatorAddrRewardsBody
*/
type PostDistributionDelegatorsDelegatorAddrRewardsBody struct {

	// base req
	BaseReq *PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReq `json:"base_req,omitempty"`
}

// Validate validates this post distribution delegators delegator addr rewards body
func (o *PostDistributionDelegatorsDelegatorAddrRewardsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsBody) validateBaseReq(formats strfmt.Registry) error {
	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Withdraw request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Withdraw request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post distribution delegators delegator addr rewards body based on the context it is used
func (o *PostDistributionDelegatorsDelegatorAddrRewardsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBaseReq(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsBody) contextValidateBaseReq(ctx context.Context, formats strfmt.Registry) error {

	if o.BaseReq != nil {
		if err := o.BaseReq.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Withdraw request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Withdraw request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsBody) UnmarshalBinary(b []byte) error {
	var res PostDistributionDelegatorsDelegatorAddrRewardsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostDistributionDelegatorsDelegatorAddrRewardsOKBody post distribution delegators delegator addr rewards o k body
swagger:model PostDistributionDelegatorsDelegatorAddrRewardsOKBody
*/
type PostDistributionDelegatorsDelegatorAddrRewardsOKBody struct {

	// fee
	Fee *PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFee `json:"fee,omitempty"`

	// memo
	Memo string `json:"memo,omitempty"`

	// msg
	Msg []string `json:"msg"`

	// signature
	Signature *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignature `json:"signature,omitempty"`
}

// Validate validates this post distribution delegators delegator addr rewards o k body
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBody) validateFee(formats strfmt.Registry) error {
	if swag.IsZero(o.Fee) { // not required
		return nil
	}

	if o.Fee != nil {
		if err := o.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBody) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(o.Signature) { // not required
		return nil
	}

	if o.Signature != nil {
		if err := o.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post distribution delegators delegator addr rewards o k body based on the context it is used
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBody) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if o.Fee != nil {
		if err := o.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBody) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if o.Signature != nil {
		if err := o.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBody) UnmarshalBinary(b []byte) error {
	var res PostDistributionDelegatorsDelegatorAddrRewardsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFee post distribution delegators delegator addr rewards o k body fee
swagger:model PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFee
*/
type PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFee struct {

	// amount
	Amount []*PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFeeAmountItems0 `json:"amount"`

	// gas
	Gas string `json:"gas,omitempty"`
}

// Validate validates this post distribution delegators delegator addr rewards o k body fee
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFee) validateAmount(formats strfmt.Registry) error {
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
					return ve.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post distribution delegators delegator addr rewards o k body fee based on the context it is used
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFee) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Amount); i++ {

		if o.Amount[i] != nil {
			if err := o.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFee) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFee) UnmarshalBinary(b []byte) error {
	var res PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFeeAmountItems0 post distribution delegators delegator addr rewards o k body fee amount items0
swagger:model PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFeeAmountItems0
*/
type PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFeeAmountItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post distribution delegators delegator addr rewards o k body fee amount items0
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFeeAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post distribution delegators delegator addr rewards o k body fee amount items0 based on context it is used
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFeeAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFeeAmountItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFeeAmountItems0) UnmarshalBinary(b []byte) error {
	var res PostDistributionDelegatorsDelegatorAddrRewardsOKBodyFeeAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignature post distribution delegators delegator addr rewards o k body signature
swagger:model PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignature
*/
type PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignature struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// pub key
	PubKey *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignaturePubKey `json:"pub_key,omitempty"`

	// sequence
	// Example: 0
	Sequence string `json:"sequence,omitempty"`

	// signature
	// Example: MEUCIQD02fsDPra8MtbRsyB1w7bqTM55Wu138zQbFcWx4+CFyAIge5WNPfKIuvzBZ69MyqHsqD8S1IwiEp+iUb6VSdtlpgY=
	Signature string `json:"signature,omitempty"`
}

// Validate validates this post distribution delegators delegator addr rewards o k body signature
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignature) validatePubKey(formats strfmt.Registry) error {
	if swag.IsZero(o.PubKey) { // not required
		return nil
	}

	if o.PubKey != nil {
		if err := o.PubKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post distribution delegators delegator addr rewards o k body signature based on the context it is used
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePubKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignature) contextValidatePubKey(ctx context.Context, formats strfmt.Registry) error {

	if o.PubKey != nil {
		if err := o.PubKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignature) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignature) UnmarshalBinary(b []byte) error {
	var res PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignaturePubKey post distribution delegators delegator addr rewards o k body signature pub key
swagger:model PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignaturePubKey
*/
type PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignaturePubKey struct {

	// type
	// Example: tendermint/PubKeySecp256k1
	Type string `json:"type,omitempty"`

	// value
	// Example: Avz04VhtKJh8ACCVzlI8aTosGy0ikFXKIVHQ3jKMrosH
	Value string `json:"value,omitempty"`
}

// Validate validates this post distribution delegators delegator addr rewards o k body signature pub key
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignaturePubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post distribution delegators delegator addr rewards o k body signature pub key based on context it is used
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignaturePubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignaturePubKey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignaturePubKey) UnmarshalBinary(b []byte) error {
	var res PostDistributionDelegatorsDelegatorAddrRewardsOKBodySignaturePubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReq post distribution delegators delegator addr rewards params body base req
swagger:model PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReq
*/
type PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReq struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// chain id
	// Example: Columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// fees
	Fees []*PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReqFeesItems0 `json:"fees"`

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

// Validate validates this post distribution delegators delegator addr rewards params body base req
func (o *PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReq) validateFees(formats strfmt.Registry) error {
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
					return ve.ValidateName("Withdraw request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Withdraw request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post distribution delegators delegator addr rewards params body base req based on the context it is used
func (o *PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReq) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Fees); i++ {

		if o.Fees[i] != nil {
			if err := o.Fees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Withdraw request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Withdraw request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReq) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReq) UnmarshalBinary(b []byte) error {
	var res PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReqFeesItems0 post distribution delegators delegator addr rewards params body base req fees items0
swagger:model PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReqFeesItems0
*/
type PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReqFeesItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post distribution delegators delegator addr rewards params body base req fees items0
func (o *PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReqFeesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post distribution delegators delegator addr rewards params body base req fees items0 based on context it is used
func (o *PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReqFeesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReqFeesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReqFeesItems0) UnmarshalBinary(b []byte) error {
	var res PostDistributionDelegatorsDelegatorAddrRewardsParamsBodyBaseReqFeesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
