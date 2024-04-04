// Code generated by fastssz. DO NOT EDIT.
// Hash: 3e68b2f68581ebe6884c67c4c8e431156850eb77cb8a7d58602fae68e542e4c1
// Version: 0.1.3
package types

import (
	primitives "github.com/berachain/beacon-kit/mod/primitives"
	ssz "github.com/ferranbt/fastssz"
)

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

	// Field (0) 'Sidecars'
	if size := len(b.Sidecars); size > 6 {
		err = ssz.ErrListTooBigFn("BlobSidecars.Sidecars", size, 6)
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
		num, err := ssz.DivideInt2(len(buf), 131544, 6)
		if err != nil {
			return err
		}
		b.Sidecars = make([]*BlobSidecar, num)
		for ii := 0; ii < num; ii++ {
			if b.Sidecars[ii] == nil {
				b.Sidecars[ii] = new(BlobSidecar)
			}
			if err = b.Sidecars[ii].UnmarshalSSZ(buf[ii*131544 : (ii+1)*131544]); err != nil {
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
	size += len(b.Sidecars) * 131544

	return
}

// HashTreeRoot ssz hashes the BlobSidecars object
func (b *BlobSidecars) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlobSidecars object with a hasher
func (b *BlobSidecars) HashTreeRootWith(hh ssz.HashWalker) (err error) {
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
		hh.MerkleizeWithMixin(subIndx, num, 6)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BlobSidecars object
func (b *BlobSidecars) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
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
		err = ssz.ErrBytesLengthFn("BlobSidecar.Blob", size, 131072)
		return
	}
	dst = append(dst, b.Blob...)

	// Field (2) 'KzgCommitment'
	dst = append(dst, b.KzgCommitment[:]...)

	// Field (3) 'KzgProof'
	if size := len(b.KzgProof); size != 48 {
		err = ssz.ErrBytesLengthFn("BlobSidecar.KzgProof", size, 48)
		return
	}
	dst = append(dst, b.KzgProof...)

	// Field (4) 'BeaconBlockHeader'
	if b.BeaconBlockHeader == nil {
		b.BeaconBlockHeader = new(primitives.BeaconBlockHeader)
	}
	if dst, err = b.BeaconBlockHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'InclusionProof'
	if size := len(b.InclusionProof); size != 8 {
		err = ssz.ErrVectorLengthFn("BlobSidecar.InclusionProof", size, 8)
		return
	}
	for ii := 0; ii < 8; ii++ {
		if size := len(b.InclusionProof[ii]); size != 32 {
			err = ssz.ErrBytesLengthFn("BlobSidecar.InclusionProof[ii]", size, 32)
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
	if size != 131544 {
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
	copy(b.KzgCommitment[:], buf[131080:131128])

	// Field (3) 'KzgProof'
	if cap(b.KzgProof) == 0 {
		b.KzgProof = make([]byte, 0, len(buf[131128:131176]))
	}
	b.KzgProof = append(b.KzgProof, buf[131128:131176]...)

	// Field (4) 'BeaconBlockHeader'
	if b.BeaconBlockHeader == nil {
		b.BeaconBlockHeader = new(primitives.BeaconBlockHeader)
	}
	if err = b.BeaconBlockHeader.UnmarshalSSZ(buf[131176:131288]); err != nil {
		return err
	}

	// Field (5) 'InclusionProof'
	b.InclusionProof = make([][]byte, 8)
	for ii := 0; ii < 8; ii++ {
		if cap(b.InclusionProof[ii]) == 0 {
			b.InclusionProof[ii] = make([]byte, 0, len(buf[131288:131544][ii*32:(ii+1)*32]))
		}
		b.InclusionProof[ii] = append(b.InclusionProof[ii], buf[131288:131544][ii*32:(ii+1)*32]...)
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlobSidecar object
func (b *BlobSidecar) SizeSSZ() (size int) {
	size = 131544
	return
}

// HashTreeRoot ssz hashes the BlobSidecar object
func (b *BlobSidecar) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlobSidecar object with a hasher
func (b *BlobSidecar) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Index'
	hh.PutUint64(b.Index)

	// Field (1) 'Blob'
	if size := len(b.Blob); size != 131072 {
		err = ssz.ErrBytesLengthFn("BlobSidecar.Blob", size, 131072)
		return
	}
	hh.PutBytes(b.Blob)

	// Field (2) 'KzgCommitment'
	hh.PutBytes(b.KzgCommitment[:])

	// Field (3) 'KzgProof'
	if size := len(b.KzgProof); size != 48 {
		err = ssz.ErrBytesLengthFn("BlobSidecar.KzgProof", size, 48)
		return
	}
	hh.PutBytes(b.KzgProof)

	// Field (4) 'BeaconBlockHeader'
	if b.BeaconBlockHeader == nil {
		b.BeaconBlockHeader = new(primitives.BeaconBlockHeader)
	}
	if err = b.BeaconBlockHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'InclusionProof'
	{
		if size := len(b.InclusionProof); size != 8 {
			err = ssz.ErrVectorLengthFn("BlobSidecar.InclusionProof", size, 8)
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
		hh.Merkleize(subIndx)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BlobSidecar object
func (b *BlobSidecar) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
