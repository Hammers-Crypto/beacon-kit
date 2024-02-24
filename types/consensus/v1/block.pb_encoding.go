// Code generated by fastssz. DO NOT EDIT.
// Hash: 94a250ed22aa64cbedba3a4d48e971e39dc69e082d7ac894773324c77ebc0383
package consensusv1

import (
	github_com_itsdevbear_bolaris_types_consensus_primitives "github.com/itsdevbear/bolaris/types/consensus/primitives"
	ssz "github.com/prysmaticlabs/fastssz"
	v1 "github.com/prysmaticlabs/prysm/v5/proto/engine/v1"
)

// MarshalSSZ ssz marshals the BeaconKitBlockDeneb object
func (b *BeaconKitBlockDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconKitBlockDeneb object to a target array
func (b *BeaconKitBlockDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(44)

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Offset (1) 'Body'
	dst = ssz.WriteOffset(dst, offset)
	if b.Body == nil {
		b.Body = new(BeaconKitBlockBodyDeneb)
	}
	offset += b.Body.SizeSSZ()

	// Field (2) 'PayloadValue'
	if size := len(b.PayloadValue); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PayloadValue", size, 32)
		return
	}
	dst = append(dst, b.PayloadValue...)

	// Field (1) 'Body'
	if dst, err = b.Body.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconKitBlockDeneb object
func (b *BeaconKitBlockDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 44 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'Slot'
	b.Slot = github_com_itsdevbear_bolaris_types_consensus_primitives.Slot(ssz.UnmarshallUint64(buf[0:8]))

	// Offset (1) 'Body'
	if o1 = ssz.ReadOffset(buf[8:12]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 44 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (2) 'PayloadValue'
	if cap(b.PayloadValue) == 0 {
		b.PayloadValue = make([]byte, 0, len(buf[12:44]))
	}
	b.PayloadValue = append(b.PayloadValue, buf[12:44]...)

	// Field (1) 'Body'
	{
		buf = tail[o1:]
		if b.Body == nil {
			b.Body = new(BeaconKitBlockBodyDeneb)
		}
		if err = b.Body.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconKitBlockDeneb object
func (b *BeaconKitBlockDeneb) SizeSSZ() (size int) {
	size = 44

	// Field (1) 'Body'
	if b.Body == nil {
		b.Body = new(BeaconKitBlockBodyDeneb)
	}
	size += b.Body.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BeaconKitBlockDeneb object
func (b *BeaconKitBlockDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconKitBlockDeneb object with a hasher
func (b *BeaconKitBlockDeneb) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (1) 'Body'
	if err = b.Body.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'PayloadValue'
	if size := len(b.PayloadValue); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PayloadValue", size, 32)
		return
	}
	hh.PutBytes(b.PayloadValue)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the BeaconKitBlockBodyDeneb object
func (b *BeaconKitBlockBodyDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconKitBlockBodyDeneb object to a target array
func (b *BeaconKitBlockBodyDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(140)

	// Field (0) 'RandaoReveal'
	if size := len(b.RandaoReveal); size != 96 {
		err = ssz.ErrBytesLengthFn("--.RandaoReveal", size, 96)
		return
	}
	dst = append(dst, b.RandaoReveal...)

	// Field (1) 'Graffiti'
	if size := len(b.Graffiti); size != 32 {
		err = ssz.ErrBytesLengthFn("--.Graffiti", size, 32)
		return
	}
	dst = append(dst, b.Graffiti...)

	// Offset (2) 'Deposits'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.Deposits); ii++ {
		offset += 4
		offset += b.Deposits[ii].SizeSSZ()
	}

	// Offset (3) 'ExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(v1.ExecutionPayloadDeneb)
	}
	offset += b.ExecutionPayload.SizeSSZ()

	// Offset (4) 'BlobKzgCommitments'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.BlobKzgCommitments) * 48

	// Field (2) 'Deposits'
	if size := len(b.Deposits); size > 16 {
		err = ssz.ErrListTooBigFn("--.Deposits", size, 16)
		return
	}
	{
		offset = 4 * len(b.Deposits)
		for ii := 0; ii < len(b.Deposits); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.Deposits[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.Deposits); ii++ {
		if dst, err = b.Deposits[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (3) 'ExecutionPayload'
	if dst, err = b.ExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'BlobKzgCommitments'
	if size := len(b.BlobKzgCommitments); size > 16 {
		err = ssz.ErrListTooBigFn("--.BlobKzgCommitments", size, 16)
		return
	}
	for ii := 0; ii < len(b.BlobKzgCommitments); ii++ {
		if size := len(b.BlobKzgCommitments[ii]); size != 48 {
			err = ssz.ErrBytesLengthFn("--.BlobKzgCommitments[ii]", size, 48)
			return
		}
		dst = append(dst, b.BlobKzgCommitments[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconKitBlockBodyDeneb object
func (b *BeaconKitBlockBodyDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 140 {
		return ssz.ErrSize
	}

	tail := buf
	var o2, o3, o4 uint64

	// Field (0) 'RandaoReveal'
	if cap(b.RandaoReveal) == 0 {
		b.RandaoReveal = make([]byte, 0, len(buf[0:96]))
	}
	b.RandaoReveal = append(b.RandaoReveal, buf[0:96]...)

	// Field (1) 'Graffiti'
	if cap(b.Graffiti) == 0 {
		b.Graffiti = make([]byte, 0, len(buf[96:128]))
	}
	b.Graffiti = append(b.Graffiti, buf[96:128]...)

	// Offset (2) 'Deposits'
	if o2 = ssz.ReadOffset(buf[128:132]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 < 140 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (3) 'ExecutionPayload'
	if o3 = ssz.ReadOffset(buf[132:136]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Offset (4) 'BlobKzgCommitments'
	if o4 = ssz.ReadOffset(buf[136:140]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Field (2) 'Deposits'
	{
		buf = tail[o2:o3]
		num, err := ssz.DecodeDynamicLength(buf, 16)
		if err != nil {
			return err
		}
		b.Deposits = make([]*Deposit, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.Deposits[indx] == nil {
				b.Deposits[indx] = new(Deposit)
			}
			if err = b.Deposits[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (3) 'ExecutionPayload'
	{
		buf = tail[o3:o4]
		if b.ExecutionPayload == nil {
			b.ExecutionPayload = new(v1.ExecutionPayloadDeneb)
		}
		if err = b.ExecutionPayload.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (4) 'BlobKzgCommitments'
	{
		buf = tail[o4:]
		num, err := ssz.DivideInt2(len(buf), 48, 16)
		if err != nil {
			return err
		}
		b.BlobKzgCommitments = make([][]byte, num)
		for ii := 0; ii < num; ii++ {
			if cap(b.BlobKzgCommitments[ii]) == 0 {
				b.BlobKzgCommitments[ii] = make([]byte, 0, len(buf[ii*48:(ii+1)*48]))
			}
			b.BlobKzgCommitments[ii] = append(b.BlobKzgCommitments[ii], buf[ii*48:(ii+1)*48]...)
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconKitBlockBodyDeneb object
func (b *BeaconKitBlockBodyDeneb) SizeSSZ() (size int) {
	size = 140

	// Field (2) 'Deposits'
	for ii := 0; ii < len(b.Deposits); ii++ {
		size += 4
		size += b.Deposits[ii].SizeSSZ()
	}

	// Field (3) 'ExecutionPayload'
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(v1.ExecutionPayloadDeneb)
	}
	size += b.ExecutionPayload.SizeSSZ()

	// Field (4) 'BlobKzgCommitments'
	size += len(b.BlobKzgCommitments) * 48

	return
}

// HashTreeRoot ssz hashes the BeaconKitBlockBodyDeneb object
func (b *BeaconKitBlockBodyDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconKitBlockBodyDeneb object with a hasher
func (b *BeaconKitBlockBodyDeneb) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'RandaoReveal'
	if size := len(b.RandaoReveal); size != 96 {
		err = ssz.ErrBytesLengthFn("--.RandaoReveal", size, 96)
		return
	}
	hh.PutBytes(b.RandaoReveal)

	// Field (1) 'Graffiti'
	if size := len(b.Graffiti); size != 32 {
		err = ssz.ErrBytesLengthFn("--.Graffiti", size, 32)
		return
	}
	hh.PutBytes(b.Graffiti)

	// Field (2) 'Deposits'
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
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, num, 16)
		} else {
			hh.MerkleizeWithMixin(subIndx, num, 16)
		}
	}

	// Field (3) 'ExecutionPayload'
	if err = b.ExecutionPayload.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'BlobKzgCommitments'
	{
		if size := len(b.BlobKzgCommitments); size > 16 {
			err = ssz.ErrListTooBigFn("--.BlobKzgCommitments", size, 16)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlobKzgCommitments {
			if len(i) != 48 {
				err = ssz.ErrBytesLength
				return
			}
			hh.PutBytes(i)
		}

		numItems := uint64(len(b.BlobKzgCommitments))
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, numItems, 16)
		} else {
			hh.MerkleizeWithMixin(subIndx, numItems, 16)
		}
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}
