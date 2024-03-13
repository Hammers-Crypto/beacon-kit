// Code generated by fastssz. DO NOT EDIT.
// Hash: 8f80215fc76cf4d53c909643e4512e5340c5207543e62ba7d1243ab43c0b51b4
package signing

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the ForkData object
func (f *ForkData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(f)
}

// MarshalSSZTo ssz marshals the ForkData object to a target array
func (f *ForkData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'CurrentVersion'
	dst = append(dst, f.CurrentVersion[:]...)

	// Field (1) 'GenesisValidatorsRoot'
	dst = append(dst, f.GenesisValidatorsRoot[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the ForkData object
func (f *ForkData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 36 {
		return ssz.ErrSize
	}

	// Field (0) 'CurrentVersion'
	copy(f.CurrentVersion[:], buf[0:4])

	// Field (1) 'GenesisValidatorsRoot'
	copy(f.GenesisValidatorsRoot[:], buf[4:36])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ForkData object
func (f *ForkData) SizeSSZ() (size int) {
	size = 36
	return
}

// HashTreeRoot ssz hashes the ForkData object
func (f *ForkData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(f)
}

// HashTreeRootWith ssz hashes the ForkData object with a hasher
func (f *ForkData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentVersion'
	hh.PutBytes(f.CurrentVersion[:])

	// Field (1) 'GenesisValidatorsRoot'
	hh.PutBytes(f.GenesisValidatorsRoot[:])

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Data object
func (d *Data) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the Data object to a target array
func (d *Data) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ObjectRoot'
	if size := len(d.ObjectRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ObjectRoot", size, 32)
		return
	}
	dst = append(dst, d.ObjectRoot...)

	// Field (1) 'Domain'
	dst = append(dst, d.Domain[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the Data object
func (d *Data) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'ObjectRoot'
	if cap(d.ObjectRoot) == 0 {
		d.ObjectRoot = make([]byte, 0, len(buf[0:32]))
	}
	d.ObjectRoot = append(d.ObjectRoot, buf[0:32]...)

	// Field (1) 'Domain'
	copy(d.Domain[:], buf[32:64])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Data object
func (d *Data) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the Data object
func (d *Data) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the Data object with a hasher
func (d *Data) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ObjectRoot'
	if size := len(d.ObjectRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ObjectRoot", size, 32)
		return
	}
	hh.PutBytes(d.ObjectRoot)

	// Field (1) 'Domain'
	hh.PutBytes(d.Domain[:])

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}
