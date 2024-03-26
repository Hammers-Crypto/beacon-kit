// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package enginetypes

import (
	"encoding/json"

	"github.com/berachain/beacon-kit/primitives"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*withdrawalJSONMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (w Withdrawal) MarshalJSON() ([]byte, error) {
	type Withdrawal struct {
		Index     hexutil.Uint64 `json:"index"          ssz-size:"8"`
		Validator hexutil.Uint64 `json:"validatorIndex" ssz-size:"8"`
		Address   common.Address `json:"address"        ssz-size:"20"`
		Amount    hexutil.Uint64 `json:"amount"         ssz-size:"8"`
	}
	var enc Withdrawal
	enc.Index = hexutil.Uint64(w.Index)
	enc.Validator = hexutil.Uint64(w.Validator)
	enc.Address = w.Address
	enc.Amount = hexutil.Uint64(w.Amount)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (w *Withdrawal) UnmarshalJSON(input []byte) error {
	type Withdrawal struct {
		Index     *hexutil.Uint64 `json:"index"          ssz-size:"8"`
		Validator *hexutil.Uint64 `json:"validatorIndex" ssz-size:"8"`
		Address   *common.Address `json:"address"        ssz-size:"20"`
		Amount    *hexutil.Uint64 `json:"amount"         ssz-size:"8"`
	}
	var dec Withdrawal
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Index != nil {
		w.Index = uint64(*dec.Index)
	}
	if dec.Validator != nil {
		w.Validator = primitives.ValidatorIndex(*dec.Validator)
	}
	if dec.Address != nil {
		w.Address = *dec.Address
	}
	if dec.Amount != nil {
		w.Amount = uint64(*dec.Amount)
	}
	return nil
}
