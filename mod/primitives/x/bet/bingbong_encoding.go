// Code generated by fastssz. DO NOT EDIT.
// Hash: 59f5e4ff3293a212556e203d4c97e59dfbd85853cd19156e69605b2a017e189f
// Version: 0.1.3
package bet

import (
	"github.com/berachain/beacon-kit/mod/primitives/math"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the ItemB object
func (i *ItemB) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(i)
}

// MarshalSSZTo ssz marshals the ItemB object to a target array
func (i *ItemB) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Field (0) 'Val1'
	dst = ssz.MarshalUint64(dst, uint64(i.Val1))

	// Offset (1) 'Val2'
	dst = ssz.WriteOffset(dst, offset)

	// Field (1) 'Val2'
	if size := len(i.Val2); size > 2 {
		err = ssz.ErrListTooBigFn("ItemB.Val2", size, 2)
		return
	}
	for ii := 0; ii < len(i.Val2); ii++ {
		dst = ssz.MarshalUint64(dst, i.Val2[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the ItemB object
func (i *ItemB) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'Val1'
	i.Val1 = math.U64(ssz.UnmarshallUint64(buf[0:8]))

	// Offset (1) 'Val2'
	if o1 = ssz.ReadOffset(buf[8:12]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Val2'
	{
		buf = tail[o1:]
		num, err := ssz.DivideInt2(len(buf), 8, 2)
		if err != nil {
			return err
		}
		i.Val2 = ssz.ExtendUint64(i.Val2, num)
		for ii := 0; ii < num; ii++ {
			i.Val2[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ItemB object
func (i *ItemB) SizeSSZ() (size int) {
	size = 12

	// Field (1) 'Val2'
	size += len(i.Val2) * 8

	return
}

// HashTreeRoot ssz hashes the ItemB object
func (i *ItemB) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(i)
}

// HashTreeRootWith ssz hashes the ItemB object with a hasher
func (i *ItemB) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Val1'
	hh.PutUint64(uint64(i.Val1))

	// Field (1) 'Val2'
	{
		if size := len(i.Val2); size > 2 {
			err = ssz.ErrListTooBigFn("ItemB.Val2", size, 2)
			return
		}
		subIndx := hh.Index()
		for _, i := range i.Val2 {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(i.Val2))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(2, numItems, 8))
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ItemB object
func (i *ItemB) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(i)
}
