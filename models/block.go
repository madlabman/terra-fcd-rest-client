// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Block block
//
// swagger:model Block
type Block struct {

	// evidence
	Evidence []string `json:"evidence"`

	// header
	Header *BlockHeader `json:"header,omitempty"`

	// last commit
	LastCommit *BlockLastCommit `json:"last_commit,omitempty"`

	// txs
	Txs []string `json:"txs"`
}

// Validate validates this block
func (m *Block) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastCommit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Block) validateHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.Header) { // not required
		return nil
	}

	if m.Header != nil {
		if err := m.Header.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

func (m *Block) validateLastCommit(formats strfmt.Registry) error {
	if swag.IsZero(m.LastCommit) { // not required
		return nil
	}

	if m.LastCommit != nil {
		if err := m.LastCommit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_commit")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this block based on the context it is used
func (m *Block) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastCommit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Block) contextValidateHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.Header != nil {
		if err := m.Header.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

func (m *Block) contextValidateLastCommit(ctx context.Context, formats strfmt.Registry) error {

	if m.LastCommit != nil {
		if err := m.LastCommit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_commit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Block) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Block) UnmarshalBinary(b []byte) error {
	var res Block
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockHeader block header
//
// swagger:model BlockHeader
type BlockHeader struct {

	// app hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	AppHash string `json:"app_hash,omitempty"`

	// chain id
	// Example: columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// consensus hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	ConsensusHash string `json:"consensus_hash,omitempty"`

	// data hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	DataHash string `json:"data_hash,omitempty"`

	// evidence hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	EvidenceHash string `json:"evidence_hash,omitempty"`

	// height
	// Example: 1
	Height float64 `json:"height,omitempty"`

	// last block id
	LastBlockID *BlockHeaderLastBlockID `json:"last_block_id,omitempty"`

	// last commit hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	LastCommitHash string `json:"last_commit_hash,omitempty"`

	// last results hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	LastResultsHash string `json:"last_results_hash,omitempty"`

	// next validators hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	NextValidatorsHash string `json:"next_validators_hash,omitempty"`

	// num txs
	// Example: 0
	NumTxs float64 `json:"num_txs,omitempty"`

	// bech32 encoded address
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	ProposerAddress string `json:"proposer_address,omitempty"`

	// time
	// Example: 2017-12-30T05:53:09.287+01:00
	Time string `json:"time,omitempty"`

	// total txs
	// Example: 35
	TotalTxs float64 `json:"total_txs,omitempty"`

	// validators hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	ValidatorsHash string `json:"validators_hash,omitempty"`

	// version
	Version *BlockHeaderVersion `json:"version,omitempty"`
}

// Validate validates this block header
func (m *BlockHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastBlockID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockHeader) validateLastBlockID(formats strfmt.Registry) error {
	if swag.IsZero(m.LastBlockID) { // not required
		return nil
	}

	if m.LastBlockID != nil {
		if err := m.LastBlockID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "last_block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "last_block_id")
			}
			return err
		}
	}

	return nil
}

func (m *BlockHeader) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this block header based on the context it is used
func (m *BlockHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastBlockID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockHeader) contextValidateLastBlockID(ctx context.Context, formats strfmt.Registry) error {

	if m.LastBlockID != nil {
		if err := m.LastBlockID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "last_block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "last_block_id")
			}
			return err
		}
	}

	return nil
}

func (m *BlockHeader) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {
		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockHeader) UnmarshalBinary(b []byte) error {
	var res BlockHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockHeaderLastBlockID block header last block ID
//
// swagger:model BlockHeaderLastBlockID
type BlockHeaderLastBlockID struct {

	// hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	Hash string `json:"hash,omitempty"`

	// parts
	Parts *BlockHeaderLastBlockIDParts `json:"parts,omitempty"`
}

// Validate validates this block header last block ID
func (m *BlockHeaderLastBlockID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockHeaderLastBlockID) validateParts(formats strfmt.Registry) error {
	if swag.IsZero(m.Parts) { // not required
		return nil
	}

	if m.Parts != nil {
		if err := m.Parts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "last_block_id" + "." + "parts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "last_block_id" + "." + "parts")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this block header last block ID based on the context it is used
func (m *BlockHeaderLastBlockID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockHeaderLastBlockID) contextValidateParts(ctx context.Context, formats strfmt.Registry) error {

	if m.Parts != nil {
		if err := m.Parts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "last_block_id" + "." + "parts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "last_block_id" + "." + "parts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockHeaderLastBlockID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockHeaderLastBlockID) UnmarshalBinary(b []byte) error {
	var res BlockHeaderLastBlockID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockHeaderLastBlockIDParts block header last block ID parts
//
// swagger:model BlockHeaderLastBlockIDParts
type BlockHeaderLastBlockIDParts struct {

	// hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	Hash string `json:"hash,omitempty"`

	// total
	// Example: 0
	Total float64 `json:"total,omitempty"`
}

// Validate validates this block header last block ID parts
func (m *BlockHeaderLastBlockIDParts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this block header last block ID parts based on context it is used
func (m *BlockHeaderLastBlockIDParts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlockHeaderLastBlockIDParts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockHeaderLastBlockIDParts) UnmarshalBinary(b []byte) error {
	var res BlockHeaderLastBlockIDParts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockHeaderVersion block header version
//
// swagger:model BlockHeaderVersion
type BlockHeaderVersion struct {

	// app
	// Example: 0
	App string `json:"app,omitempty"`

	// block
	// Example: 10
	Block string `json:"block,omitempty"`
}

// Validate validates this block header version
func (m *BlockHeaderVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this block header version based on context it is used
func (m *BlockHeaderVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlockHeaderVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockHeaderVersion) UnmarshalBinary(b []byte) error {
	var res BlockHeaderVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockLastCommit block last commit
//
// swagger:model BlockLastCommit
type BlockLastCommit struct {

	// block id
	BlockID *BlockLastCommitBlockID `json:"block_id,omitempty"`

	// precommits
	Precommits []*BlockLastCommitPrecommitsItems0 `json:"precommits"`
}

// Validate validates this block last commit
func (m *BlockLastCommit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrecommits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockLastCommit) validateBlockID(formats strfmt.Registry) error {
	if swag.IsZero(m.BlockID) { // not required
		return nil
	}

	if m.BlockID != nil {
		if err := m.BlockID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_commit" + "." + "block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_commit" + "." + "block_id")
			}
			return err
		}
	}

	return nil
}

func (m *BlockLastCommit) validatePrecommits(formats strfmt.Registry) error {
	if swag.IsZero(m.Precommits) { // not required
		return nil
	}

	for i := 0; i < len(m.Precommits); i++ {
		if swag.IsZero(m.Precommits[i]) { // not required
			continue
		}

		if m.Precommits[i] != nil {
			if err := m.Precommits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("last_commit" + "." + "precommits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("last_commit" + "." + "precommits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this block last commit based on the context it is used
func (m *BlockLastCommit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlockID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrecommits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockLastCommit) contextValidateBlockID(ctx context.Context, formats strfmt.Registry) error {

	if m.BlockID != nil {
		if err := m.BlockID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_commit" + "." + "block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_commit" + "." + "block_id")
			}
			return err
		}
	}

	return nil
}

func (m *BlockLastCommit) contextValidatePrecommits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Precommits); i++ {

		if m.Precommits[i] != nil {
			if err := m.Precommits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("last_commit" + "." + "precommits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("last_commit" + "." + "precommits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockLastCommit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockLastCommit) UnmarshalBinary(b []byte) error {
	var res BlockLastCommit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockLastCommitBlockID block last commit block ID
//
// swagger:model BlockLastCommitBlockID
type BlockLastCommitBlockID struct {

	// hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	Hash string `json:"hash,omitempty"`

	// parts
	Parts *BlockLastCommitBlockIDParts `json:"parts,omitempty"`
}

// Validate validates this block last commit block ID
func (m *BlockLastCommitBlockID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockLastCommitBlockID) validateParts(formats strfmt.Registry) error {
	if swag.IsZero(m.Parts) { // not required
		return nil
	}

	if m.Parts != nil {
		if err := m.Parts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_commit" + "." + "block_id" + "." + "parts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_commit" + "." + "block_id" + "." + "parts")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this block last commit block ID based on the context it is used
func (m *BlockLastCommitBlockID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockLastCommitBlockID) contextValidateParts(ctx context.Context, formats strfmt.Registry) error {

	if m.Parts != nil {
		if err := m.Parts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_commit" + "." + "block_id" + "." + "parts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_commit" + "." + "block_id" + "." + "parts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockLastCommitBlockID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockLastCommitBlockID) UnmarshalBinary(b []byte) error {
	var res BlockLastCommitBlockID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockLastCommitBlockIDParts block last commit block ID parts
//
// swagger:model BlockLastCommitBlockIDParts
type BlockLastCommitBlockIDParts struct {

	// hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	Hash string `json:"hash,omitempty"`

	// total
	// Example: 0
	Total float64 `json:"total,omitempty"`
}

// Validate validates this block last commit block ID parts
func (m *BlockLastCommitBlockIDParts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this block last commit block ID parts based on context it is used
func (m *BlockLastCommitBlockIDParts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlockLastCommitBlockIDParts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockLastCommitBlockIDParts) UnmarshalBinary(b []byte) error {
	var res BlockLastCommitBlockIDParts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockLastCommitPrecommitsItems0 block last commit precommits items0
//
// swagger:model BlockLastCommitPrecommitsItems0
type BlockLastCommitPrecommitsItems0 struct {

	// block id
	BlockID *BlockLastCommitPrecommitsItems0BlockID `json:"block_id,omitempty"`

	// height
	// Example: 0
	Height string `json:"height,omitempty"`

	// round
	// Example: 0
	Round string `json:"round,omitempty"`

	// signature
	// Example: 7uTC74QlknqYWEwg7Vn6M8Om7FuZ0EO4bjvuj6rwH1mTUJrRuMMZvAAqT9VjNgP0RA/TDp6u/92AqrZfXJSpBQ==
	Signature string `json:"signature,omitempty"`

	// timestamp
	// Example: 2017-12-30T05:53:09.287+01:00
	Timestamp string `json:"timestamp,omitempty"`

	// type
	// Example: 2
	Type float64 `json:"type,omitempty"`

	// validator address
	ValidatorAddress string `json:"validator_address,omitempty"`

	// validator index
	// Example: 0
	ValidatorIndex string `json:"validator_index,omitempty"`
}

// Validate validates this block last commit precommits items0
func (m *BlockLastCommitPrecommitsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockLastCommitPrecommitsItems0) validateBlockID(formats strfmt.Registry) error {
	if swag.IsZero(m.BlockID) { // not required
		return nil
	}

	if m.BlockID != nil {
		if err := m.BlockID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block_id")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this block last commit precommits items0 based on the context it is used
func (m *BlockLastCommitPrecommitsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlockID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockLastCommitPrecommitsItems0) contextValidateBlockID(ctx context.Context, formats strfmt.Registry) error {

	if m.BlockID != nil {
		if err := m.BlockID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block_id")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockLastCommitPrecommitsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockLastCommitPrecommitsItems0) UnmarshalBinary(b []byte) error {
	var res BlockLastCommitPrecommitsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockLastCommitPrecommitsItems0BlockID block last commit precommits items0 block ID
//
// swagger:model BlockLastCommitPrecommitsItems0BlockID
type BlockLastCommitPrecommitsItems0BlockID struct {

	// hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	Hash string `json:"hash,omitempty"`

	// parts
	Parts *BlockLastCommitPrecommitsItems0BlockIDParts `json:"parts,omitempty"`
}

// Validate validates this block last commit precommits items0 block ID
func (m *BlockLastCommitPrecommitsItems0BlockID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockLastCommitPrecommitsItems0BlockID) validateParts(formats strfmt.Registry) error {
	if swag.IsZero(m.Parts) { // not required
		return nil
	}

	if m.Parts != nil {
		if err := m.Parts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block_id" + "." + "parts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block_id" + "." + "parts")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this block last commit precommits items0 block ID based on the context it is used
func (m *BlockLastCommitPrecommitsItems0BlockID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockLastCommitPrecommitsItems0BlockID) contextValidateParts(ctx context.Context, formats strfmt.Registry) error {

	if m.Parts != nil {
		if err := m.Parts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block_id" + "." + "parts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block_id" + "." + "parts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockLastCommitPrecommitsItems0BlockID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockLastCommitPrecommitsItems0BlockID) UnmarshalBinary(b []byte) error {
	var res BlockLastCommitPrecommitsItems0BlockID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockLastCommitPrecommitsItems0BlockIDParts block last commit precommits items0 block ID parts
//
// swagger:model BlockLastCommitPrecommitsItems0BlockIDParts
type BlockLastCommitPrecommitsItems0BlockIDParts struct {

	// hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	Hash string `json:"hash,omitempty"`

	// total
	// Example: 0
	Total float64 `json:"total,omitempty"`
}

// Validate validates this block last commit precommits items0 block ID parts
func (m *BlockLastCommitPrecommitsItems0BlockIDParts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this block last commit precommits items0 block ID parts based on context it is used
func (m *BlockLastCommitPrecommitsItems0BlockIDParts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlockLastCommitPrecommitsItems0BlockIDParts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockLastCommitPrecommitsItems0BlockIDParts) UnmarshalBinary(b []byte) error {
	var res BlockLastCommitPrecommitsItems0BlockIDParts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
