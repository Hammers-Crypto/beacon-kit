// Code generated by fastssz. DO NOT EDIT.
// Hash: 2a6118608869bb95d912ae478e6ed7502440ea71a3176bd917744722f4575ea6
// Version: 0.1.3
package ssz

import (
	ssz "github.com/ferranbt/fastssz"
)

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
