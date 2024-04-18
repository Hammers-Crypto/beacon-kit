// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package ssz

import (
	"reflect"

	"github.com/berachain/beacon-kit/mod/merkle"
)

// Merkleize hashes the packed value and returns the HTR.
func MerkleizeBasic[B Basic, RootT ~[32]byte](value B) (RootT, error) {
	return MerkleizeVecBasic[B, RootT]([]B{value})
}

// func MerkleizeVecBasic[B Basic, RootT ~[32]byte](value B) (RootT, error) {
func MerkleizeVecBasic[B Basic, RootT ~[32]byte](value []B) (RootT, error) {
	packed, err := Pack[B, RootT](value)
	if err != nil {
		return [32]byte{}, err
	}
	return Merkleize[RootT, RootT](packed)
}

func MerkleizeListBasic[B Basic, RootT ~[32]byte](value []B) (RootT, error) {
	packed, err := Pack[B, RootT](value)
	if err != nil {
		return [32]byte{}, err
	}
	root, err := Merkleize[RootT, RootT](packed, ChunkCountBasicListVec(value))
	if err != nil {
		return [32]byte{}, err
	}
	return merkle.MixinLength(root, uint64(len(value))), nil
}

// TODO bitlist
func MerkleizeContainer[C Composite[RootT], RootT ~[32]byte](value C) (RootT, error) {
	htrs := make([]RootT, reflect.ValueOf(value).NumField())
	var err error
	for i := 0; i < reflect.ValueOf(value).NumField(); i++ {
		field := reflect.ValueOf(value).Field(i).Interface().(Composite[RootT])
		htrs[i], err = field.HashTreeRoot()
		if err != nil {
			return RootT{}, err
		}
	}
	return Merkleize[RootT, RootT](htrs)
}

func MerkleizeVecComposite[C Composite[RootT], RootT ~[32]byte](value []C) (RootT, error) {
	htrs := make([]RootT, len(value))
	var err error
	for i, el := range value {
		htrs[i], err = el.HashTreeRoot()
		if err != nil {
			return RootT{}, err
		}
	}
	return Merkleize[RootT, RootT](htrs)
}

func MerkleizeListComposite[C Composite[RootT], RootT ~[32]byte](value []C) (RootT, error) {
	htrs := make([]RootT, len(value))
	var err error
	for i, el := range value {
		htrs[i], err = el.HashTreeRoot()
		if err != nil {
			return RootT{}, err
		}
	}
	root, err := Merkleize[RootT, RootT](htrs, ChunkCountCompositeList[RootT, C](value))
	if err != nil {
		return RootT{}, err
	}
	return merkle.MixinLength(root, uint64(len(value))), nil
}

// -----------------------------------------------------------------------------

// MerkleizeList hashes each element in the list and then returns the HTR of
// the list of corresponding roots, with the length mixed in.
func MerkleizeList[T Hashable[[32]byte]](
	elements []T, limit uint64,
) ([32]byte, error) {
	body, err := MerkleizeVector(elements, limit)
	if err != nil {
		return [32]byte{}, err
	}
	return merkle.MixinLength(body, uint64(len(elements))), nil
}

// MerkleizeVector hashes each element in the list and then returns the HTR
// of the corresponding list of roots.
func MerkleizeVector[T Hashable[[32]byte]](
	elements []T, length uint64,
) ([32]byte, error) {
	roots := make([][32]byte, len(elements))
	var err error
	for i, el := range elements {
		roots[i], err = el.HashTreeRoot()
		if err != nil {
			return [32]byte{}, err
		}
	}
	return merkle.NewRootWithMaxLeaves[[32]byte, [32]byte](roots, length)
}
