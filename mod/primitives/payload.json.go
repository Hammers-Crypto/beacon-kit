// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package primitives

import (
	"encoding/json"
	"errors"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*executableDataDenebMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (e ExecutableDataDeneb) MarshalJSON() ([]byte, error) {
	type ExecutableDataDeneb struct {
		ParentHash    common.Hash     `json:"parentHash"    ssz-size:"32"  gencodec:"required"`
		FeeRecipient  common.Address  `json:"feeRecipient"  ssz-size:"20"  gencodec:"required"`
		StateRoot     Bytes32         `json:"stateRoot"     ssz-size:"32"  gencodec:"required"`
		ReceiptsRoot  Bytes32         `json:"receiptsRoot"  ssz-size:"32"  gencodec:"required"`
		LogsBloom     hexutil.Bytes   `json:"logsBloom"     ssz-size:"256" gencodec:"required"`
		Random        Bytes32         `json:"prevRandao"    ssz-size:"32"  gencodec:"required"`
		Number        math.U64        `json:"blockNumber"                  gencodec:"required"`
		GasLimit      math.U64        `json:"gasLimit"                     gencodec:"required"`
		GasUsed       math.U64        `json:"gasUsed"                      gencodec:"required"`
		Timestamp     math.U64        `json:"timestamp"                    gencodec:"required"`
		ExtraData     hexutil.Bytes   `json:"extraData"                    gencodec:"required" ssz-max:"32"`
		BaseFeePerGas math.U256L      `json:"baseFeePerGas" ssz-size:"32"  gencodec:"required"`
		BlockHash     common.Hash     `json:"blockHash"     ssz-size:"32"  gencodec:"required"`
		Transactions  []hexutil.Bytes `json:"transactions"  ssz-size:"?,?" gencodec:"required" ssz-max:"1048576,1073741824"`
		Withdrawals   []*Withdrawal   `json:"withdrawals"                                      ssz-max:"16"`
		BlobGasUsed   math.U64        `json:"blobGasUsed"`
		ExcessBlobGas math.U64        `json:"excessBlobGas"`
	}
	var enc ExecutableDataDeneb
	enc.ParentHash = e.ParentHash
	enc.FeeRecipient = e.FeeRecipient
	enc.StateRoot = e.StateRoot
	enc.ReceiptsRoot = e.ReceiptsRoot
	enc.LogsBloom = e.LogsBloom
	enc.Random = e.Random
	enc.Number = e.Number
	enc.GasLimit = e.GasLimit
	enc.GasUsed = e.GasUsed
	enc.Timestamp = e.Timestamp
	enc.ExtraData = e.ExtraData
	enc.BaseFeePerGas = e.BaseFeePerGas
	enc.BlockHash = e.BlockHash
	if e.Transactions != nil {
		enc.Transactions = make([]hexutil.Bytes, len(e.Transactions))
		for k, v := range e.Transactions {
			enc.Transactions[k] = v
		}
	}
	enc.Withdrawals = e.Withdrawals
	enc.BlobGasUsed = e.BlobGasUsed
	enc.ExcessBlobGas = e.ExcessBlobGas
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (e *ExecutableDataDeneb) UnmarshalJSON(input []byte) error {
	type ExecutableDataDeneb struct {
		ParentHash    *common.Hash    `json:"parentHash"    ssz-size:"32"  gencodec:"required"`
		FeeRecipient  *common.Address `json:"feeRecipient"  ssz-size:"20"  gencodec:"required"`
		StateRoot     *Bytes32        `json:"stateRoot"     ssz-size:"32"  gencodec:"required"`
		ReceiptsRoot  *Bytes32        `json:"receiptsRoot"  ssz-size:"32"  gencodec:"required"`
		LogsBloom     *hexutil.Bytes  `json:"logsBloom"     ssz-size:"256" gencodec:"required"`
		Random        *Bytes32        `json:"prevRandao"    ssz-size:"32"  gencodec:"required"`
		Number        *math.U64       `json:"blockNumber"                  gencodec:"required"`
		GasLimit      *math.U64       `json:"gasLimit"                     gencodec:"required"`
		GasUsed       *math.U64       `json:"gasUsed"                      gencodec:"required"`
		Timestamp     *math.U64       `json:"timestamp"                    gencodec:"required"`
		ExtraData     *hexutil.Bytes  `json:"extraData"                    gencodec:"required" ssz-max:"32"`
		BaseFeePerGas *math.U256L     `json:"baseFeePerGas" ssz-size:"32"  gencodec:"required"`
		BlockHash     *common.Hash    `json:"blockHash"     ssz-size:"32"  gencodec:"required"`
		Transactions  []hexutil.Bytes `json:"transactions"  ssz-size:"?,?" gencodec:"required" ssz-max:"1048576,1073741824"`
		Withdrawals   []*Withdrawal   `json:"withdrawals"                                      ssz-max:"16"`
		BlobGasUsed   *math.U64       `json:"blobGasUsed"`
		ExcessBlobGas *math.U64       `json:"excessBlobGas"`
	}
	var dec ExecutableDataDeneb
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for ExecutableDataDeneb")
	}
	e.ParentHash = *dec.ParentHash
	if dec.FeeRecipient == nil {
		return errors.New("missing required field 'feeRecipient' for ExecutableDataDeneb")
	}
	e.FeeRecipient = *dec.FeeRecipient
	if dec.StateRoot == nil {
		return errors.New("missing required field 'stateRoot' for ExecutableDataDeneb")
	}
	e.StateRoot = *dec.StateRoot
	if dec.ReceiptsRoot == nil {
		return errors.New("missing required field 'receiptsRoot' for ExecutableDataDeneb")
	}
	e.ReceiptsRoot = *dec.ReceiptsRoot
	if dec.LogsBloom == nil {
		return errors.New("missing required field 'logsBloom' for ExecutableDataDeneb")
	}
	e.LogsBloom = *dec.LogsBloom
	if dec.Random == nil {
		return errors.New("missing required field 'prevRandao' for ExecutableDataDeneb")
	}
	e.Random = *dec.Random
	if dec.Number == nil {
		return errors.New("missing required field 'blockNumber' for ExecutableDataDeneb")
	}
	e.Number = *dec.Number
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for ExecutableDataDeneb")
	}
	e.GasLimit = *dec.GasLimit
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for ExecutableDataDeneb")
	}
	e.GasUsed = *dec.GasUsed
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for ExecutableDataDeneb")
	}
	e.Timestamp = *dec.Timestamp
	if dec.ExtraData == nil {
		return errors.New("missing required field 'extraData' for ExecutableDataDeneb")
	}
	e.ExtraData = *dec.ExtraData
	if dec.BaseFeePerGas == nil {
		return errors.New("missing required field 'baseFeePerGas' for ExecutableDataDeneb")
	}
	e.BaseFeePerGas = *dec.BaseFeePerGas
	if dec.BlockHash == nil {
		return errors.New("missing required field 'blockHash' for ExecutableDataDeneb")
	}
	e.BlockHash = *dec.BlockHash
	if dec.Transactions == nil {
		return errors.New("missing required field 'transactions' for ExecutableDataDeneb")
	}
	e.Transactions = make([][]byte, len(dec.Transactions))
	for k, v := range dec.Transactions {
		e.Transactions[k] = v
	}
	if dec.Withdrawals != nil {
		e.Withdrawals = dec.Withdrawals
	}
	if dec.BlobGasUsed != nil {
		e.BlobGasUsed = *dec.BlobGasUsed
	}
	if dec.ExcessBlobGas != nil {
		e.ExcessBlobGas = *dec.ExcessBlobGas
	}
	return nil
}
