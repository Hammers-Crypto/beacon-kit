// Code generated by fastssz. DO NOT EDIT.
// Hash: fb60a498e0a808ec5a6729e9dc2fe06660c9a0d4bd899a2988253290538a99dc
// Version: 0.1.3
package primitives

import (
	eip4844 "github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlockBodyDeneb object to a target array
func (b *BeaconBlockBodyDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(212)

	// Field (0) 'RandaoReveal'
	dst = append(dst, b.RandaoReveal[:]...)

	// Field (1) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(Eth1Data)
	}
	if dst, err = b.Eth1Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Graffiti'
	dst = append(dst, b.Graffiti[:]...)

	// Offset (3) 'Deposits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Deposits) * 192

	// Offset (4) 'ExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(ExecutableDataDeneb)
	}
	offset += b.ExecutionPayload.SizeSSZ()

	// Offset (5) 'BlobKzgCommitments'
	dst = ssz.WriteOffset(dst, offset)

	// Field (3) 'Deposits'
	if size := len(b.Deposits); size > 16 {
		err = ssz.ErrListTooBigFn("BeaconBlockBodyDeneb.Deposits", size, 16)
		return
	}
	for ii := 0; ii < len(b.Deposits); ii++ {
		if dst, err = b.Deposits[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (4) 'ExecutionPayload'
	if dst, err = b.ExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'BlobKzgCommitments'
	if size := len(b.BlobKzgCommitments); size > 16 {
		err = ssz.ErrListTooBigFn("BeaconBlockBodyDeneb.BlobKzgCommitments", size, 16)
		return
	}
	for ii := 0; ii < len(b.BlobKzgCommitments); ii++ {
		dst = append(dst, b.BlobKzgCommitments[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 212 {
		return ssz.ErrSize
	}

	tail := buf
	var o3, o4, o5 uint64

	// Field (0) 'RandaoReveal'
	copy(b.RandaoReveal[:], buf[0:96])

	// Field (1) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(Eth1Data)
	}
	if err = b.Eth1Data.UnmarshalSSZ(buf[96:168]); err != nil {
		return err
	}

	// Field (2) 'Graffiti'
	copy(b.Graffiti[:], buf[168:200])

	// Offset (3) 'Deposits'
	if o3 = ssz.ReadOffset(buf[200:204]); o3 > size {
		return ssz.ErrOffset
	}

	if o3 < 212 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (4) 'ExecutionPayload'
	if o4 = ssz.ReadOffset(buf[204:208]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Offset (5) 'BlobKzgCommitments'
	if o5 = ssz.ReadOffset(buf[208:212]); o5 > size || o4 > o5 {
		return ssz.ErrOffset
	}

	// Field (3) 'Deposits'
	{
		buf = tail[o3:o4]
		num, err := ssz.DivideInt2(len(buf), 192, 16)
		if err != nil {
			return err
		}
		b.Deposits = make([]*Deposit, num)
		for ii := 0; ii < num; ii++ {
			if b.Deposits[ii] == nil {
				b.Deposits[ii] = new(Deposit)
			}
			if err = b.Deposits[ii].UnmarshalSSZ(buf[ii*192 : (ii+1)*192]); err != nil {
				return err
			}
		}
	}

	// Field (4) 'ExecutionPayload'
	{
		buf = tail[o4:o5]
		if b.ExecutionPayload == nil {
			b.ExecutionPayload = new(ExecutableDataDeneb)
		}
		if err = b.ExecutionPayload.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (5) 'BlobKzgCommitments'
	{
		buf = tail[o5:]
		num, err := ssz.DivideInt2(len(buf), 48, 16)
		if err != nil {
			return err
		}
		b.BlobKzgCommitments = make([]eip4844.KZGCommitment, num)
		for ii := 0; ii < num; ii++ {
			copy(b.BlobKzgCommitments[ii][:], buf[ii*48:(ii+1)*48])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) SizeSSZ() (size int) {
	size = 212

	// Field (3) 'Deposits'
	size += len(b.Deposits) * 192

	// Field (4) 'ExecutionPayload'
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(ExecutableDataDeneb)
	}
	size += b.ExecutionPayload.SizeSSZ()

	// Field (5) 'BlobKzgCommitments'
	size += len(b.BlobKzgCommitments) * 48

	return
}

// HashTreeRoot ssz hashes the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlockBodyDeneb object with a hasher
func (b *BeaconBlockBodyDeneb) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'RandaoReveal'
	hh.PutBytes(b.RandaoReveal[:])

	// Field (1) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(Eth1Data)
	}
	if err = b.Eth1Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Graffiti'
	hh.PutBytes(b.Graffiti[:])

	// Field (3) 'Deposits'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Deposits))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Deposits {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 16)
	}

	// Field (4) 'ExecutionPayload'
	if err = b.ExecutionPayload.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'BlobKzgCommitments'
	{
		if size := len(b.BlobKzgCommitments); size > 16 {
			err = ssz.ErrListTooBigFn("BeaconBlockBodyDeneb.BlobKzgCommitments", size, 16)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlobKzgCommitments {
			hh.PutBytes(i[:])
		}
		numItems := uint64(len(b.BlobKzgCommitments))
		hh.MerkleizeWithMixin(subIndx, numItems, 16)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
