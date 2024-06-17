// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/ethereum/go-ethereum/common"
)

var _ = (*executionPayloadHeaderDenebMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (e ExecutionPayloadHeaderDeneb) MarshalJSON() ([]byte, error) {
	type ExecutionPayloadHeaderDeneb struct {
		ParentHash       common.Hash    `json:"parentHash"       ssz-size:"32"  gencodec:"required"`
		FeeRecipient     common.Address `json:"feeRecipient"     ssz-size:"20"  gencodec:"required"`
		StateRoot        bytes.B32      `json:"stateRoot"        ssz-size:"32"  gencodec:"required"`
		ReceiptsRoot     bytes.B32      `json:"receiptsRoot"     ssz-size:"32"  gencodec:"required"`
		LogsBloom        bytes.Bytes    `json:"logsBloom"        ssz-size:"256" gencodec:"required"`
		Random           bytes.B32      `json:"prevRandao"       ssz-size:"32"  gencodec:"required"`
		Number           math.U64       `json:"blockNumber"                     gencodec:"required"`
		GasLimit         math.U64       `json:"gasLimit"                        gencodec:"required"`
		GasUsed          math.U64       `json:"gasUsed"                         gencodec:"required"`
		Timestamp        math.U64       `json:"timestamp"                       gencodec:"required"`
		ExtraData        bytes.Bytes    `json:"extraData"                       gencodec:"required" ssz-max:"32"`
		BaseFeePerGas    math.U256L     `json:"baseFeePerGas"    ssz-size:"32"  gencodec:"required"`
		BlockHash        common.Hash    `json:"blockHash"        ssz-size:"32"  gencodec:"required"`
		TransactionsRoot bytes.B32      `json:"transactionsRoot" ssz-size:"32"  gencodec:"required"`
		WithdrawalsRoot  bytes.B32      `json:"withdrawalsRoot"  ssz-size:"32"`
		BlobGasUsed      math.U64       `json:"blobGasUsed"`
		ExcessBlobGas    math.U64       `json:"excessBlobGas"`
	}
	var enc ExecutionPayloadHeaderDeneb
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
	enc.TransactionsRoot = e.TransactionsRoot
	enc.WithdrawalsRoot = e.WithdrawalsRoot
	enc.BlobGasUsed = e.BlobGasUsed
	enc.ExcessBlobGas = e.ExcessBlobGas
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (e *ExecutionPayloadHeaderDeneb) UnmarshalJSON(input []byte) error {
	type ExecutionPayloadHeaderDeneb struct {
		ParentHash       *common.Hash    `json:"parentHash"       ssz-size:"32"  gencodec:"required"`
		FeeRecipient     *common.Address `json:"feeRecipient"     ssz-size:"20"  gencodec:"required"`
		StateRoot        *bytes.B32      `json:"stateRoot"        ssz-size:"32"  gencodec:"required"`
		ReceiptsRoot     *bytes.B32      `json:"receiptsRoot"     ssz-size:"32"  gencodec:"required"`
		LogsBloom        *bytes.Bytes    `json:"logsBloom"        ssz-size:"256" gencodec:"required"`
		Random           *bytes.B32      `json:"prevRandao"       ssz-size:"32"  gencodec:"required"`
		Number           *math.U64       `json:"blockNumber"                     gencodec:"required"`
		GasLimit         *math.U64       `json:"gasLimit"                        gencodec:"required"`
		GasUsed          *math.U64       `json:"gasUsed"                         gencodec:"required"`
		Timestamp        *math.U64       `json:"timestamp"                       gencodec:"required"`
		ExtraData        *bytes.Bytes    `json:"extraData"                       gencodec:"required" ssz-max:"32"`
		BaseFeePerGas    *math.U256L     `json:"baseFeePerGas"    ssz-size:"32"  gencodec:"required"`
		BlockHash        *common.Hash    `json:"blockHash"        ssz-size:"32"  gencodec:"required"`
		TransactionsRoot *bytes.B32      `json:"transactionsRoot" ssz-size:"32"  gencodec:"required"`
		WithdrawalsRoot  *bytes.B32      `json:"withdrawalsRoot"  ssz-size:"32"`
		BlobGasUsed      *math.U64       `json:"blobGasUsed"`
		ExcessBlobGas    *math.U64       `json:"excessBlobGas"`
	}

	fmt.Println("HENLO")
	var dec ExecutionPayloadHeaderDeneb
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for ExecutionPayloadHeaderDeneb")
	}
	e.ParentHash = *dec.ParentHash
	if dec.FeeRecipient == nil {
		return errors.New("missing required field 'feeRecipient' for ExecutionPayloadHeaderDeneb")
	}
	e.FeeRecipient = *dec.FeeRecipient
	if dec.StateRoot == nil {
		return errors.New("missing required field 'stateRoot' for ExecutionPayloadHeaderDeneb")
	}
	e.StateRoot = *dec.StateRoot
	if dec.ReceiptsRoot == nil {
		return errors.New("missing required field 'receiptsRoot' for ExecutionPayloadHeaderDeneb")
	}
	e.ReceiptsRoot = *dec.ReceiptsRoot
	if dec.LogsBloom == nil {
		return errors.New("missing required field 'logsBloom' for ExecutionPayloadHeaderDeneb")
	}
	e.LogsBloom = *dec.LogsBloom
	if dec.Random == nil {
		return errors.New("missing required field 'prevRandao' for ExecutionPayloadHeaderDeneb")
	}
	e.Random = *dec.Random
	if dec.Number == nil {
		return errors.New("missing required field 'blockNumber' for ExecutionPayloadHeaderDeneb")
	}
	e.Number = *dec.Number
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for ExecutionPayloadHeaderDeneb")
	}
	e.GasLimit = *dec.GasLimit
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for ExecutionPayloadHeaderDeneb")
	}
	e.GasUsed = *dec.GasUsed
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for ExecutionPayloadHeaderDeneb")
	}
	e.Timestamp = *dec.Timestamp
	if dec.ExtraData == nil {
		return errors.New("missing required field 'extraData' for ExecutionPayloadHeaderDeneb")
	}
	e.ExtraData = *dec.ExtraData
	if dec.BaseFeePerGas == nil {
		return errors.New("missing required field 'baseFeePerGas' for ExecutionPayloadHeaderDeneb")
	}
	e.BaseFeePerGas = *dec.BaseFeePerGas
	if dec.BlockHash == nil {
		return errors.New("missing required field 'blockHash' for ExecutionPayloadHeaderDeneb")
	}
	e.BlockHash = *dec.BlockHash
	if dec.TransactionsRoot == nil {
		return errors.New("missing required field 'transactionsRoot' for ExecutionPayloadHeaderDeneb")
	}
	e.TransactionsRoot = *dec.TransactionsRoot
	if dec.WithdrawalsRoot != nil {
		e.WithdrawalsRoot = *dec.WithdrawalsRoot
	}
	if dec.BlobGasUsed != nil {
		e.BlobGasUsed = *dec.BlobGasUsed
	}
	if dec.ExcessBlobGas != nil {
		e.ExcessBlobGas = *dec.ExcessBlobGas
	}

	fmt.Println("POINTER DIFF?")
	fmt.Println(e)
	return nil
}
