// Code generated by fastssz. DO NOT EDIT.
// Hash: 65cea7852bfc707eb9b47c31c9911417e24762e4cf1ccb14c4b7a57a7f4a8402
// Version: 0.1.3
package types

import (
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the ExecutionPayloadHeaderDeneb object
func (e *ExecutionPayloadHeaderDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ExecutionPayloadHeaderDeneb object to a target array
func (e *ExecutionPayloadHeaderDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(584)

	// Field (0) 'ParentHash'
	dst = append(dst, e.ParentHash[:]...)

	// Field (1) 'FeeRecipient'
	dst = append(dst, e.FeeRecipient[:]...)

	// Field (2) 'StateRoot'
	dst = append(dst, e.StateRoot[:]...)

	// Field (3) 'ReceiptsRoot'
	dst = append(dst, e.ReceiptsRoot[:]...)

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("ExecutionPayloadHeaderDeneb.LogsBloom", size, 256)
		return
	}
	dst = append(dst, e.LogsBloom...)

	// Field (5) 'Random'
	dst = append(dst, e.Random[:]...)

	// Field (6) 'Number'
	dst = ssz.MarshalUint64(dst, uint64(e.Number))

	// Field (7) 'GasLimit'
	dst = ssz.MarshalUint64(dst, uint64(e.GasLimit))

	// Field (8) 'GasUsed'
	dst = ssz.MarshalUint64(dst, uint64(e.GasUsed))

	// Field (9) 'Timestamp'
	dst = ssz.MarshalUint64(dst, uint64(e.Timestamp))

	// Offset (10) 'ExtraData'
	dst = ssz.WriteOffset(dst, offset)

	// Field (11) 'BaseFeePerGas'
	dst = append(dst, e.BaseFeePerGas[:]...)

	// Field (12) 'BlockHash'
	dst = append(dst, e.BlockHash[:]...)

	// Field (13) 'TransactionsRoot'
	dst = append(dst, e.TransactionsRoot[:]...)

	// Field (14) 'WithdrawalsRoot'
	dst = append(dst, e.WithdrawalsRoot[:]...)

	// Field (15) 'BlobGasUsed'
	dst = ssz.MarshalUint64(dst, uint64(e.BlobGasUsed))

	// Field (16) 'ExcessBlobGas'
	dst = ssz.MarshalUint64(dst, uint64(e.ExcessBlobGas))

	// Field (10) 'ExtraData'
	if size := len(e.ExtraData); size > 32 {
		err = ssz.ErrBytesLengthFn("ExecutionPayloadHeaderDeneb.ExtraData", size, 32)
		return
	}
	dst = append(dst, e.ExtraData...)

	return
}

// UnmarshalSSZ ssz unmarshals the ExecutionPayloadHeaderDeneb object
func (e *ExecutionPayloadHeaderDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 584 {
		return ssz.ErrSize
	}

	tail := buf
	var o10 uint64

	// Field (0) 'ParentHash'
	copy(e.ParentHash[:], buf[0:32])

	// Field (1) 'FeeRecipient'
	copy(e.FeeRecipient[:], buf[32:52])

	// Field (2) 'StateRoot'
	copy(e.StateRoot[:], buf[52:84])

	// Field (3) 'ReceiptsRoot'
	copy(e.ReceiptsRoot[:], buf[84:116])

	// Field (4) 'LogsBloom'
	if cap(e.LogsBloom) == 0 {
		e.LogsBloom = make([]byte, 0, len(buf[116:372]))
	}
	e.LogsBloom = append(e.LogsBloom, buf[116:372]...)

	// Field (5) 'Random'
	copy(e.Random[:], buf[372:404])

	// Field (6) 'Number'
	e.Number = math.U64(ssz.UnmarshallUint64(buf[404:412]))

	// Field (7) 'GasLimit'
	e.GasLimit = math.U64(ssz.UnmarshallUint64(buf[412:420]))

	// Field (8) 'GasUsed'
	e.GasUsed = math.U64(ssz.UnmarshallUint64(buf[420:428]))

	// Field (9) 'Timestamp'
	e.Timestamp = math.U64(ssz.UnmarshallUint64(buf[428:436]))

	// Offset (10) 'ExtraData'
	if o10 = ssz.ReadOffset(buf[436:440]); o10 > size {
		return ssz.ErrOffset
	}

	if o10 < 584 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (11) 'BaseFeePerGas'
	copy(e.BaseFeePerGas[:], buf[440:472])

	// Field (12) 'BlockHash'
	copy(e.BlockHash[:], buf[472:504])

	// Field (13) 'TransactionsRoot'
	copy(e.TransactionsRoot[:], buf[504:536])

	// Field (14) 'WithdrawalsRoot'
	copy(e.WithdrawalsRoot[:], buf[536:568])

	// Field (15) 'BlobGasUsed'
	e.BlobGasUsed = math.U64(ssz.UnmarshallUint64(buf[568:576]))

	// Field (16) 'ExcessBlobGas'
	e.ExcessBlobGas = math.U64(ssz.UnmarshallUint64(buf[576:584]))

	// Field (10) 'ExtraData'
	{
		buf = tail[o10:]
		if len(buf) > 32 {
			return ssz.ErrBytesLength
		}
		if cap(e.ExtraData) == 0 {
			e.ExtraData = make([]byte, 0, len(buf))
		}
		e.ExtraData = append(e.ExtraData, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ExecutionPayloadHeaderDeneb object
func (e *ExecutionPayloadHeaderDeneb) SizeSSZ() (size int) {
	size = 584

	// Field (10) 'ExtraData'
	size += len(e.ExtraData)

	return
}

// HashTreeRoot ssz hashes the ExecutionPayloadHeaderDeneb object
func (e *ExecutionPayloadHeaderDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ExecutionPayloadHeaderDeneb object with a hasher
func (e *ExecutionPayloadHeaderDeneb) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ParentHash'
	hh.PutBytes(e.ParentHash[:])

	// Field (1) 'FeeRecipient'
	hh.PutBytes(e.FeeRecipient[:])

	// Field (2) 'StateRoot'
	hh.PutBytes(e.StateRoot[:])

	// Field (3) 'ReceiptsRoot'
	hh.PutBytes(e.ReceiptsRoot[:])

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("ExecutionPayloadHeaderDeneb.LogsBloom", size, 256)
		return
	}
	hh.PutBytes(e.LogsBloom)

	// Field (5) 'Random'
	hh.PutBytes(e.Random[:])

	// Field (6) 'Number'
	hh.PutUint64(uint64(e.Number))

	// Field (7) 'GasLimit'
	hh.PutUint64(uint64(e.GasLimit))

	// Field (8) 'GasUsed'
	hh.PutUint64(uint64(e.GasUsed))

	// Field (9) 'Timestamp'
	hh.PutUint64(uint64(e.Timestamp))

	// Field (10) 'ExtraData'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(e.ExtraData))
		if byteLen > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(e.ExtraData)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (32+31)/32)
	}

	// Field (11) 'BaseFeePerGas'
	hh.PutBytes(e.BaseFeePerGas[:])

	// Field (12) 'BlockHash'
	hh.PutBytes(e.BlockHash[:])

	// Field (13) 'TransactionsRoot'
	hh.PutBytes(e.TransactionsRoot[:])

	// Field (14) 'WithdrawalsRoot'
	hh.PutBytes(e.WithdrawalsRoot[:])

	// Field (15) 'BlobGasUsed'
	hh.PutUint64(uint64(e.BlobGasUsed))

	// Field (16) 'ExcessBlobGas'
	hh.PutUint64(uint64(e.ExcessBlobGas))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ExecutionPayloadHeaderDeneb object
func (e *ExecutionPayloadHeaderDeneb) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(e)
}
