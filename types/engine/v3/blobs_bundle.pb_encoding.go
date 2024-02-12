// Code generated by fastssz. DO NOT EDIT.
// Hash: 974f360ea9cf490b0d408638386b1f5c324b5dcf055fbfd05a87257c9b6c952a
package v3

import (
	ssz "github.com/itsdevbear/fastssz"
)

// MarshalSSZ ssz marshals the BlobsBundle object
func (b *BlobsBundle) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlobsBundle object to a target array
func (b *BlobsBundle) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Offset (0) 'KzgCommitments'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.KzgCommitments) * 48

	// Offset (1) 'Proofs'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Proofs) * 48

	// Offset (2) 'Blobs'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Blobs) * 131072

	// Field (0) 'KzgCommitments'
	if size := len(b.KzgCommitments); size > 16 {
		err = ssz.ErrListTooBigFn("--.KzgCommitments", size, 16)
		return
	}
	for ii := 0; ii < len(b.KzgCommitments); ii++ {
		if size := len(b.KzgCommitments[ii]); size != 48 {
			err = ssz.ErrBytesLengthFn("--.KzgCommitments[ii]", size, 48)
			return
		}
		dst = append(dst, b.KzgCommitments[ii]...)
	}

	// Field (1) 'Proofs'
	if size := len(b.Proofs); size > 16 {
		err = ssz.ErrListTooBigFn("--.Proofs", size, 16)
		return
	}
	for ii := 0; ii < len(b.Proofs); ii++ {
		if size := len(b.Proofs[ii]); size != 48 {
			err = ssz.ErrBytesLengthFn("--.Proofs[ii]", size, 48)
			return
		}
		dst = append(dst, b.Proofs[ii]...)
	}

	// Field (2) 'Blobs'
	if size := len(b.Blobs); size > 16 {
		err = ssz.ErrListTooBigFn("--.Blobs", size, 16)
		return
	}
	for ii := 0; ii < len(b.Blobs); ii++ {
		if size := len(b.Blobs[ii]); size != 131072 {
			err = ssz.ErrBytesLengthFn("--.Blobs[ii]", size, 131072)
			return
		}
		dst = append(dst, b.Blobs[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BlobsBundle object
func (b *BlobsBundle) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2 uint64

	// Offset (0) 'KzgCommitments'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'Proofs'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'Blobs'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'KzgCommitments'
	{
		buf = tail[o0:o1]
		num, err := ssz.DivideInt2(len(buf), 48, 16)
		if err != nil {
			return err
		}
		b.KzgCommitments = make([][]byte, num)
		for ii := 0; ii < num; ii++ {
			if cap(b.KzgCommitments[ii]) == 0 {
				b.KzgCommitments[ii] = make([]byte, 0, len(buf[ii*48:(ii+1)*48]))
			}
			b.KzgCommitments[ii] = append(b.KzgCommitments[ii], buf[ii*48:(ii+1)*48]...)
		}
	}

	// Field (1) 'Proofs'
	{
		buf = tail[o1:o2]
		num, err := ssz.DivideInt2(len(buf), 48, 16)
		if err != nil {
			return err
		}
		b.Proofs = make([][]byte, num)
		for ii := 0; ii < num; ii++ {
			if cap(b.Proofs[ii]) == 0 {
				b.Proofs[ii] = make([]byte, 0, len(buf[ii*48:(ii+1)*48]))
			}
			b.Proofs[ii] = append(b.Proofs[ii], buf[ii*48:(ii+1)*48]...)
		}
	}

	// Field (2) 'Blobs'
	{
		buf = tail[o2:]
		num, err := ssz.DivideInt2(len(buf), 131072, 16)
		if err != nil {
			return err
		}
		b.Blobs = make([][]byte, num)
		for ii := 0; ii < num; ii++ {
			if cap(b.Blobs[ii]) == 0 {
				b.Blobs[ii] = make([]byte, 0, len(buf[ii*131072:(ii+1)*131072]))
			}
			b.Blobs[ii] = append(b.Blobs[ii], buf[ii*131072:(ii+1)*131072]...)
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlobsBundle object
func (b *BlobsBundle) SizeSSZ() (size int) {
	size = 12

	// Field (0) 'KzgCommitments'
	size += len(b.KzgCommitments) * 48

	// Field (1) 'Proofs'
	size += len(b.Proofs) * 48

	// Field (2) 'Blobs'
	size += len(b.Blobs) * 131072

	return
}

// HashTreeRoot ssz hashes the BlobsBundle object
func (b *BlobsBundle) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlobsBundle object with a hasher
func (b *BlobsBundle) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'KzgCommitments'
	{
		if size := len(b.KzgCommitments); size > 16 {
			err = ssz.ErrListTooBigFn("--.KzgCommitments", size, 16)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.KzgCommitments {
			if len(i) != 48 {
				err = ssz.ErrBytesLength
				return
			}
			hh.PutBytes(i)
		}

		numItems := uint64(len(b.KzgCommitments))
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, numItems, 16)
		} else {
			hh.MerkleizeWithMixin(subIndx, numItems, 16)
		}
	}

	// Field (1) 'Proofs'
	{
		if size := len(b.Proofs); size > 16 {
			err = ssz.ErrListTooBigFn("--.Proofs", size, 16)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Proofs {
			if len(i) != 48 {
				err = ssz.ErrBytesLength
				return
			}
			hh.PutBytes(i)
		}

		numItems := uint64(len(b.Proofs))
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, numItems, 16)
		} else {
			hh.MerkleizeWithMixin(subIndx, numItems, 16)
		}
	}

	// Field (2) 'Blobs'
	{
		if size := len(b.Blobs); size > 16 {
			err = ssz.ErrListTooBigFn("--.Blobs", size, 16)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Blobs {
			if len(i) != 131072 {
				err = ssz.ErrBytesLength
				return
			}
			hh.PutBytes(i)
		}

		numItems := uint64(len(b.Blobs))
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
