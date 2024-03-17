// Code generated by fastssz. DO NOT EDIT.
// Hash: e45e703708af3ad58a0ef779daa3148cb61ef9c8d5e59e023fc97711dbaf7102
package types

import (
	enginetypes "github.com/berachain/beacon-kit/engine/types"
	"github.com/berachain/beacon-kit/primitives"
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlockDeneb object to a target array
func (b *BeaconBlockDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(52)

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Field (1) 'ProposerIndex'
	dst = ssz.MarshalUint64(dst, uint64(b.ProposerIndex))

	// Field (2) 'ParentBlockRoot'
	dst = append(dst, b.ParentBlockRoot[:]...)

	// Offset (3) 'Body'
	dst = ssz.WriteOffset(dst, offset)
	if b.Body == nil {
		b.Body = new(BeaconBlockBodyDeneb)
	}
	offset += b.Body.SizeSSZ()

	// Field (3) 'Body'
	if dst, err = b.Body.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 52 {
		return ssz.ErrSize
	}

	tail := buf
	var o3 uint64

	// Field (0) 'Slot'
	b.Slot = primitives.Slot(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'ProposerIndex'
	b.ProposerIndex = primitives.ValidatorIndex(ssz.UnmarshallUint64(buf[8:16]))

	// Field (2) 'ParentBlockRoot'
	copy(b.ParentBlockRoot[:], buf[16:48])

	// Offset (3) 'Body'
	if o3 = ssz.ReadOffset(buf[48:52]); o3 > size {
		return ssz.ErrOffset
	}

	if o3 < 52 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (3) 'Body'
	{
		buf = tail[o3:]
		if b.Body == nil {
			b.Body = new(BeaconBlockBodyDeneb)
		}
		if err = b.Body.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) SizeSSZ() (size int) {
	size = 52

	// Field (3) 'Body'
	if b.Body == nil {
		b.Body = new(BeaconBlockBodyDeneb)
	}
	size += b.Body.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlockDeneb object with a hasher
func (b *BeaconBlockDeneb) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (1) 'ProposerIndex'
	hh.PutUint64(uint64(b.ProposerIndex))

	// Field (2) 'ParentBlockRoot'
	hh.PutBytes(b.ParentBlockRoot[:])

	// Field (3) 'Body'
	if err = b.Body.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlockBodyDeneb object to a target array
func (b *BeaconBlockBodyDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(140)

	// Field (0) 'RandaoReveal'
	dst = append(dst, b.RandaoReveal[:]...)

	// Field (1) 'Graffiti'
	dst = append(dst, b.Graffiti[:]...)

	// Offset (2) 'Deposits'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.Deposits); ii++ {
		offset += 4
		offset += b.Deposits[ii].SizeSSZ()
	}

	// Offset (3) 'ExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(enginetypes.ExecutableDataDeneb)
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
		dst = append(dst, b.BlobKzgCommitments[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 140 {
		return ssz.ErrSize
	}

	tail := buf
	var o2, o3, o4 uint64

	// Field (0) 'RandaoReveal'
	copy(b.RandaoReveal[:], buf[0:96])

	// Field (1) 'Graffiti'
	copy(b.Graffiti[:], buf[96:128])

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
			b.ExecutionPayload = new(enginetypes.ExecutableDataDeneb)
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
		b.BlobKzgCommitments = make([][48]byte, num)
		for ii := 0; ii < num; ii++ {
			copy(b.BlobKzgCommitments[ii][:], buf[ii*48:(ii+1)*48])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) SizeSSZ() (size int) {
	size = 140

	// Field (2) 'Deposits'
	for ii := 0; ii < len(b.Deposits); ii++ {
		size += 4
		size += b.Deposits[ii].SizeSSZ()
	}

	// Field (3) 'ExecutionPayload'
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(enginetypes.ExecutableDataDeneb)
	}
	size += b.ExecutionPayload.SizeSSZ()

	// Field (4) 'BlobKzgCommitments'
	size += len(b.BlobKzgCommitments) * 48

	return
}

// HashTreeRoot ssz hashes the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlockBodyDeneb object with a hasher
func (b *BeaconBlockBodyDeneb) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'RandaoReveal'
	hh.PutBytes(b.RandaoReveal[:])

	// Field (1) 'Graffiti'
	hh.PutBytes(b.Graffiti[:])

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
			hh.PutBytes(i[:])
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

// MarshalSSZ ssz marshals the Deposit object
func (d *Deposit) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the Deposit object to a target array
func (d *Deposit) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(56)

	// Offset (0) 'Pubkey'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(d.Pubkey)

	// Field (1) 'Credentials'
	if size := len(d.Credentials); size != 32 {
		err = ssz.ErrBytesLengthFn("--.Credentials", size, 32)
		return
	}
	dst = append(dst, d.Credentials...)

	// Field (2) 'Amount'
	dst = ssz.MarshalUint64(dst, d.Amount)

	// Offset (3) 'Signature'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(d.Signature)

	// Field (4) 'Index'
	dst = ssz.MarshalUint64(dst, d.Index)

	// Field (0) 'Pubkey'
	if size := len(d.Pubkey); size > 48 {
		err = ssz.ErrBytesLengthFn("--.Pubkey", size, 48)
		return
	}
	dst = append(dst, d.Pubkey...)

	// Field (3) 'Signature'
	if size := len(d.Signature); size > 96 {
		err = ssz.ErrBytesLengthFn("--.Signature", size, 96)
		return
	}
	dst = append(dst, d.Signature...)

	return
}

// UnmarshalSSZ ssz unmarshals the Deposit object
func (d *Deposit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 56 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o3 uint64

	// Offset (0) 'Pubkey'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 56 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Credentials'
	if cap(d.Credentials) == 0 {
		d.Credentials = make([]byte, 0, len(buf[4:36]))
	}
	d.Credentials = append(d.Credentials, buf[4:36]...)

	// Field (2) 'Amount'
	d.Amount = ssz.UnmarshallUint64(buf[36:44])

	// Offset (3) 'Signature'
	if o3 = ssz.ReadOffset(buf[44:48]); o3 > size || o0 > o3 {
		return ssz.ErrOffset
	}

	// Field (4) 'Index'
	d.Index = ssz.UnmarshallUint64(buf[48:56])

	// Field (0) 'Pubkey'
	{
		buf = tail[o0:o3]
		if len(buf) > 48 {
			return ssz.ErrBytesLength
		}
		if cap(d.Pubkey) == 0 {
			d.Pubkey = make([]byte, 0, len(buf))
		}
		d.Pubkey = append(d.Pubkey, buf...)
	}

	// Field (3) 'Signature'
	{
		buf = tail[o3:]
		if len(buf) > 96 {
			return ssz.ErrBytesLength
		}
		if cap(d.Signature) == 0 {
			d.Signature = make([]byte, 0, len(buf))
		}
		d.Signature = append(d.Signature, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Deposit object
func (d *Deposit) SizeSSZ() (size int) {
	size = 56

	// Field (0) 'Pubkey'
	size += len(d.Pubkey)

	// Field (3) 'Signature'
	size += len(d.Signature)

	return
}

// HashTreeRoot ssz hashes the Deposit object
func (d *Deposit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the Deposit object with a hasher
func (d *Deposit) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkey'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(d.Pubkey))
		if byteLen > 48 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(d.Pubkey)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (48+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (48+31)/32)
		}
	}

	// Field (1) 'Credentials'
	if size := len(d.Credentials); size != 32 {
		err = ssz.ErrBytesLengthFn("--.Credentials", size, 32)
		return
	}
	hh.PutBytes(d.Credentials)

	// Field (2) 'Amount'
	hh.PutUint64(d.Amount)

	// Field (3) 'Signature'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(d.Signature))
		if byteLen > 96 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(d.Signature)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (96+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (96+31)/32)
		}
	}

	// Field (4) 'Index'
	hh.PutUint64(d.Index)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the BlobSidecars object
func (b *BlobSidecars) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlobSidecars object to a target array
func (b *BlobSidecars) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Sidecars'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Sidecars) * 131432

	// Field (0) 'Sidecars'
	if size := len(b.Sidecars); size > 6 {
		err = ssz.ErrListTooBigFn("--.Sidecars", size, 6)
		return
	}
	for ii := 0; ii < len(b.Sidecars); ii++ {
		if dst, err = b.Sidecars[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BlobSidecars object
func (b *BlobSidecars) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Sidecars'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Sidecars'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 131432, 6)
		if err != nil {
			return err
		}
		b.Sidecars = make([]*BlobSidecar, num)
		for ii := 0; ii < num; ii++ {
			if b.Sidecars[ii] == nil {
				b.Sidecars[ii] = new(BlobSidecar)
			}
			if err = b.Sidecars[ii].UnmarshalSSZ(buf[ii*131432 : (ii+1)*131432]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlobSidecars object
func (b *BlobSidecars) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Sidecars'
	size += len(b.Sidecars) * 131432

	return
}

// HashTreeRoot ssz hashes the BlobSidecars object
func (b *BlobSidecars) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlobSidecars object with a hasher
func (b *BlobSidecars) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Sidecars'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Sidecars))
		if num > 6 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Sidecars {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, num, 6)
		} else {
			hh.MerkleizeWithMixin(subIndx, num, 6)
		}
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the BlobSidecar object
func (b *BlobSidecar) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlobSidecar object to a target array
func (b *BlobSidecar) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Index'
	dst = ssz.MarshalUint64(dst, b.Index)

	// Field (1) 'Blob'
	if size := len(b.Blob); size != 131072 {
		err = ssz.ErrBytesLengthFn("--.Blob", size, 131072)
		return
	}
	dst = append(dst, b.Blob...)

	// Field (2) 'KzgCommitment'
	if size := len(b.KzgCommitment); size != 48 {
		err = ssz.ErrBytesLengthFn("--.KzgCommitment", size, 48)
		return
	}
	dst = append(dst, b.KzgCommitment...)

	// Field (3) 'KzgProof'
	if size := len(b.KzgProof); size != 48 {
		err = ssz.ErrBytesLengthFn("--.KzgProof", size, 48)
		return
	}
	dst = append(dst, b.KzgProof...)

	// Field (4) 'InclusionProof'
	if size := len(b.InclusionProof); size != 8 {
		err = ssz.ErrVectorLengthFn("--.InclusionProof", size, 8)
		return
	}
	for ii := 0; ii < 8; ii++ {
		if size := len(b.InclusionProof[ii]); size != 32 {
			err = ssz.ErrBytesLengthFn("--.InclusionProof[ii]", size, 32)
			return
		}
		dst = append(dst, b.InclusionProof[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BlobSidecar object
func (b *BlobSidecar) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 131432 {
		return ssz.ErrSize
	}

	// Field (0) 'Index'
	b.Index = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Blob'
	if cap(b.Blob) == 0 {
		b.Blob = make([]byte, 0, len(buf[8:131080]))
	}
	b.Blob = append(b.Blob, buf[8:131080]...)

	// Field (2) 'KzgCommitment'
	if cap(b.KzgCommitment) == 0 {
		b.KzgCommitment = make([]byte, 0, len(buf[131080:131128]))
	}
	b.KzgCommitment = append(b.KzgCommitment, buf[131080:131128]...)

	// Field (3) 'KzgProof'
	if cap(b.KzgProof) == 0 {
		b.KzgProof = make([]byte, 0, len(buf[131128:131176]))
	}
	b.KzgProof = append(b.KzgProof, buf[131128:131176]...)

	// Field (4) 'InclusionProof'
	b.InclusionProof = make([][]byte, 8)
	for ii := 0; ii < 8; ii++ {
		if cap(b.InclusionProof[ii]) == 0 {
			b.InclusionProof[ii] = make([]byte, 0, len(buf[131176:131432][ii*32:(ii+1)*32]))
		}
		b.InclusionProof[ii] = append(b.InclusionProof[ii], buf[131176:131432][ii*32:(ii+1)*32]...)
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlobSidecar object
func (b *BlobSidecar) SizeSSZ() (size int) {
	size = 131432
	return
}

// HashTreeRoot ssz hashes the BlobSidecar object
func (b *BlobSidecar) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlobSidecar object with a hasher
func (b *BlobSidecar) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Index'
	hh.PutUint64(b.Index)

	// Field (1) 'Blob'
	if size := len(b.Blob); size != 131072 {
		err = ssz.ErrBytesLengthFn("--.Blob", size, 131072)
		return
	}
	hh.PutBytes(b.Blob)

	// Field (2) 'KzgCommitment'
	if size := len(b.KzgCommitment); size != 48 {
		err = ssz.ErrBytesLengthFn("--.KzgCommitment", size, 48)
		return
	}
	hh.PutBytes(b.KzgCommitment)

	// Field (3) 'KzgProof'
	if size := len(b.KzgProof); size != 48 {
		err = ssz.ErrBytesLengthFn("--.KzgProof", size, 48)
		return
	}
	hh.PutBytes(b.KzgProof)

	// Field (4) 'InclusionProof'
	{
		if size := len(b.InclusionProof); size != 8 {
			err = ssz.ErrVectorLengthFn("--.InclusionProof", size, 8)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.InclusionProof {
			if len(i) != 32 {
				err = ssz.ErrBytesLength
				return
			}
			hh.Append(i)
		}

		if ssz.EnableVectorizedHTR {
			hh.MerkleizeVectorizedHTR(subIndx)
		} else {
			hh.Merkleize(subIndx)
		}
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}
