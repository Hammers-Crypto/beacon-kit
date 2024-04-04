// Code generated by fastssz. DO NOT EDIT.
// Hash: 27a4e3f646235b7127022efaa6bfdd849a9f7ff57289f358f611050bb152319c
// Version: 0.1.3
package primitives

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BlockHeader object
func (b *BlockHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlockHeader object to a target array
func (b *BlockHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Field (1) 'ProposerIndex'
	dst = ssz.MarshalUint64(dst, uint64(b.ProposerIndex))

	// Field (2) 'ParentRoot'
	dst = append(dst, b.ParentRoot[:]...)

	// Field (3) 'StateRoot'
	dst = append(dst, b.StateRoot[:]...)

	// Field (4) 'BodyRoot'
	dst = append(dst, b.BodyRoot[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the BlockHeader object
func (b *BlockHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 112 {
		return ssz.ErrSize
	}

	// Field (0) 'Slot'
	b.Slot = Slot(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'ProposerIndex'
	b.ProposerIndex = ValidatorIndex(ssz.UnmarshallUint64(buf[8:16]))

	// Field (2) 'ParentRoot'
	copy(b.ParentRoot[:], buf[16:48])

	// Field (3) 'StateRoot'
	copy(b.StateRoot[:], buf[48:80])

	// Field (4) 'BodyRoot'
	copy(b.BodyRoot[:], buf[80:112])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlockHeader object
func (b *BlockHeader) SizeSSZ() (size int) {
	size = 112
	return
}

// HashTreeRoot ssz hashes the BlockHeader object
func (b *BlockHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlockHeader object with a hasher
func (b *BlockHeader) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (1) 'ProposerIndex'
	hh.PutUint64(uint64(b.ProposerIndex))

	// Field (2) 'ParentRoot'
	hh.PutBytes(b.ParentRoot[:])

	// Field (3) 'StateRoot'
	hh.PutBytes(b.StateRoot[:])

	// Field (4) 'BodyRoot'
	hh.PutBytes(b.BodyRoot[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BlockHeader object
func (b *BlockHeader) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}

// MarshalSSZ ssz marshals the SigningData object
func (s *SigningData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SigningData object to a target array
func (s *SigningData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ObjectRoot'
	dst = append(dst, s.ObjectRoot[:]...)

	// Field (1) 'Domain'
	dst = append(dst, s.Domain[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SigningData object
func (s *SigningData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'ObjectRoot'
	copy(s.ObjectRoot[:], buf[0:32])

	// Field (1) 'Domain'
	copy(s.Domain[:], buf[32:64])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SigningData object
func (s *SigningData) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the SigningData object
func (s *SigningData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SigningData object with a hasher
func (s *SigningData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ObjectRoot'
	hh.PutBytes(s.ObjectRoot[:])

	// Field (1) 'Domain'
	hh.PutBytes(s.Domain[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SigningData object
func (s *SigningData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the Withdrawal object
func (w *Withdrawal) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(w)
}

// MarshalSSZTo ssz marshals the Withdrawal object to a target array
func (w *Withdrawal) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Index'
	dst = ssz.MarshalUint64(dst, w.Index)

	// Field (1) 'Validator'
	dst = ssz.MarshalUint64(dst, uint64(w.Validator))

	// Field (2) 'Address'
	dst = append(dst, w.Address[:]...)

	// Field (3) 'Amount'
	dst = ssz.MarshalUint64(dst, uint64(w.Amount))

	return
}

// UnmarshalSSZ ssz unmarshals the Withdrawal object
func (w *Withdrawal) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 44 {
		return ssz.ErrSize
	}

	// Field (0) 'Index'
	w.Index = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Validator'
	w.Validator = ValidatorIndex(ssz.UnmarshallUint64(buf[8:16]))

	// Field (2) 'Address'
	copy(w.Address[:], buf[16:36])

	// Field (3) 'Amount'
	w.Amount = Gwei(ssz.UnmarshallUint64(buf[36:44]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Withdrawal object
func (w *Withdrawal) SizeSSZ() (size int) {
	size = 44
	return
}

// HashTreeRoot ssz hashes the Withdrawal object
func (w *Withdrawal) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(w)
}

// HashTreeRootWith ssz hashes the Withdrawal object with a hasher
func (w *Withdrawal) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Index'
	hh.PutUint64(w.Index)

	// Field (1) 'Validator'
	hh.PutUint64(uint64(w.Validator))

	// Field (2) 'Address'
	hh.PutBytes(w.Address[:])

	// Field (3) 'Amount'
	hh.PutUint64(uint64(w.Amount))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Withdrawal object
func (w *Withdrawal) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(w)
}
