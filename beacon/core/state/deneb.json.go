// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package state

import (
	"encoding/json"
	"errors"

	"github.com/berachain/beacon-kit/beacon/core/types"
	"github.com/berachain/beacon-kit/primitives"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*beaconStateDenebJSONMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (b BeaconStateDeneb) MarshalJSON() ([]byte, error) {
	type BeaconStateDeneb struct {
		GenesisValidatorsRoot        hexutil.Bytes            `json:"genesisValidatorsRoot" ssz-size:"32"`
		Slot                         primitives.Slot          `json:"slot"`
		LatestBlockHeader            *types.BeaconBlockHeader `json:"latestBlockHeader"`
		BlockRoots                   []primitives.HashRoot    `json:"blockRoots"        ssz-size:"?,32" ssz-max:"8192"`
		StateRoots                   []primitives.HashRoot    `json:"stateRoots"        ssz-size:"?,32" ssz-max:"8192"`
		Eth1GenesisHash              common.Hash              `json:"eth1GenesisHash"  ssz-size:"32"`
		Eth1DepositIndex             uint64                   `json:"eth1DepositIndex"`
		Validators                   []*types.Validator       `json:"validators" ssz-max:"1099511627776"`
		RandaoMixes                  []primitives.HashRoot    `json:"randaoMixes" ssz-size:"?,32" ssz-max:"65536"`
		NextWithdrawalIndex          uint64                   `json:"nextWithdrawalIndex"`
		NextWithdrawalValidatorIndex uint64                   `json:"nextWithdrawalValidatorIndex"`
	}
	var enc BeaconStateDeneb
	enc.GenesisValidatorsRoot = b.GenesisValidatorsRoot[:]
	enc.Slot = b.Slot
	enc.LatestBlockHeader = b.LatestBlockHeader
	if b.BlockRoots != nil {
		enc.BlockRoots = make([]primitives.HashRoot, len(b.BlockRoots))
		for k, v := range b.BlockRoots {
			enc.BlockRoots[k] = v
		}
	}
	if b.StateRoots != nil {
		enc.StateRoots = make([]primitives.HashRoot, len(b.StateRoots))
		for k, v := range b.StateRoots {
			enc.StateRoots[k] = v
		}
	}
	enc.Eth1GenesisHash = b.Eth1GenesisHash
	enc.Eth1DepositIndex = b.Eth1DepositIndex
	enc.Validators = b.Validators
	if b.RandaoMixes != nil {
		enc.RandaoMixes = make([]primitives.HashRoot, len(b.RandaoMixes))
		for k, v := range b.RandaoMixes {
			enc.RandaoMixes[k] = v
		}
	}
	enc.NextWithdrawalIndex = b.NextWithdrawalIndex
	enc.NextWithdrawalValidatorIndex = b.NextWithdrawalValidatorIndex
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (b *BeaconStateDeneb) UnmarshalJSON(input []byte) error {
	type BeaconStateDeneb struct {
		GenesisValidatorsRoot        *hexutil.Bytes           `json:"genesisValidatorsRoot" ssz-size:"32"`
		Slot                         *primitives.Slot         `json:"slot"`
		LatestBlockHeader            *types.BeaconBlockHeader `json:"latestBlockHeader"`
		BlockRoots                   []primitives.HashRoot    `json:"blockRoots"        ssz-size:"?,32" ssz-max:"8192"`
		StateRoots                   []primitives.HashRoot    `json:"stateRoots"        ssz-size:"?,32" ssz-max:"8192"`
		Eth1GenesisHash              *common.Hash             `json:"eth1GenesisHash"  ssz-size:"32"`
		Eth1DepositIndex             *uint64                  `json:"eth1DepositIndex"`
		Validators                   []*types.Validator       `json:"validators" ssz-max:"1099511627776"`
		RandaoMixes                  []primitives.HashRoot    `json:"randaoMixes" ssz-size:"?,32" ssz-max:"65536"`
		NextWithdrawalIndex          *uint64                  `json:"nextWithdrawalIndex"`
		NextWithdrawalValidatorIndex *uint64                  `json:"nextWithdrawalValidatorIndex"`
	}
	var dec BeaconStateDeneb
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.GenesisValidatorsRoot != nil {
		if len(*dec.GenesisValidatorsRoot) != len(b.GenesisValidatorsRoot) {
			return errors.New("field 'genesisValidatorsRoot' has wrong length, need 32 items")
		}
		copy(b.GenesisValidatorsRoot[:], *dec.GenesisValidatorsRoot)
	}
	if dec.Slot != nil {
		b.Slot = *dec.Slot
	}
	if dec.LatestBlockHeader != nil {
		b.LatestBlockHeader = dec.LatestBlockHeader
	}
	if dec.BlockRoots != nil {
		b.BlockRoots = make([][32]byte, len(dec.BlockRoots))
		for k, v := range dec.BlockRoots {
			b.BlockRoots[k] = v
		}
	}
	if dec.StateRoots != nil {
		b.StateRoots = make([][32]byte, len(dec.StateRoots))
		for k, v := range dec.StateRoots {
			b.StateRoots[k] = v
		}
	}
	if dec.Eth1GenesisHash != nil {
		b.Eth1GenesisHash = *dec.Eth1GenesisHash
	}
	if dec.Eth1DepositIndex != nil {
		b.Eth1DepositIndex = *dec.Eth1DepositIndex
	}
	if dec.Validators != nil {
		b.Validators = dec.Validators
	}
	if dec.RandaoMixes != nil {
		b.RandaoMixes = make([][32]byte, len(dec.RandaoMixes))
		for k, v := range dec.RandaoMixes {
			b.RandaoMixes[k] = v
		}
	}
	if dec.NextWithdrawalIndex != nil {
		b.NextWithdrawalIndex = *dec.NextWithdrawalIndex
	}
	if dec.NextWithdrawalValidatorIndex != nil {
		b.NextWithdrawalValidatorIndex = *dec.NextWithdrawalValidatorIndex
	}
	return nil
}
