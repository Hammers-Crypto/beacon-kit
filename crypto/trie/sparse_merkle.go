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

// Package trie defines utilities for sparse merkle tries for Ethereum
// consensus.
package trie

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	byteslib "github.com/berachain/beacon-kit/lib/bytes"
	"github.com/pkg/errors"
)

const (
	// 2^63 would overflow.
	MaxTrieDepth = 62
)

// SparseMerkleTrie implements a sparse, general purpose Merkle trie
// to be used across Ethereum consensus functionality.
type SparseMerkleTrie struct {
	depth    uint
	branches [][][]byte
	// list of provided items before hashing them into leaves.
	originalItems [][]byte
}

// NewTrie returns a new merkle trie filled with zerohashes to use.
func NewTrie(depth uint64) (*SparseMerkleTrie, error) {
	var zeroBytes [32]byte
	items := [][]byte{zeroBytes[:]}
	return GenerateTrieFromItems(items, depth)
}

// GenerateTrieFromItems constructs a Merkle trie
// from a sequence of byte slices.
func GenerateTrieFromItems(
	items [][]byte,
	depth uint64,
) (*SparseMerkleTrie, error) {
	if len(items) == 0 {
		return nil, errors.New("no items provided to generate Merkle trie")
	}
	if depth > MaxTrieDepth {
		// PowerOf2 would overflow
		return nil, errors.New(
			"supported merkle trie depth exceeded (max uint64 depth is 63, " +
				"theoretical max sparse merkle trie depth is 64)")
	}

	leaves := items
	layers := make([][][]byte, depth+1)
	transformedLeaves := make([][]byte, len(leaves))
	for i := range leaves {
		arr := byteslib.ToBytes32(leaves[i])
		transformedLeaves[i] = arr[:]
	}
	layers[0] = transformedLeaves
	for i := uint64(0); i < depth; i++ {
		if len(layers[i])%2 == 1 {
			layers[i] = append(layers[i], ZeroHashes[i][:])
		}
		updatedValues := make([][]byte, 0)
		for j := 0; j < len(layers[i]); j += 2 {
			concat := sha256.Sum256(append(layers[i][j], layers[i][j+1]...))
			updatedValues = append(updatedValues, concat[:])
		}
		layers[i+1] = updatedValues
	}
	return &SparseMerkleTrie{
		branches:      layers,
		originalItems: items,
		depth:         uint(depth),
	}, nil
}

// Items returns the original items passed in when creating the Merkle trie.
func (m *SparseMerkleTrie) Items() [][]byte {
	return m.originalItems
}

// HashTreeRoot returns the hash root of the Merkle trie
// defined in the deposit contract.
func (m *SparseMerkleTrie) HashTreeRoot() ([32]byte, error) {
	var enc [32]byte
	numItems := uint64(len(m.originalItems))
	if len(m.originalItems) == 1 &&
		bytes.Equal(m.originalItems[0], ZeroHashes[0][:]) {
		// Accounting for empty tries
		numItems = 0
	}
	binary.LittleEndian.PutUint64(enc[:], numItems)
	return sha256.Sum256(
		append(m.branches[len(m.branches)-1][0], enc[:]...),
	), nil
}

// Insert an item into the trie.
func (m *SparseMerkleTrie) Insert(item []byte, index int) error {
	if index < 0 {
		return fmt.Errorf("negative index provided: %d", index)
	}
	for index >= len(m.branches[0]) {
		m.branches[0] = append(m.branches[0], ZeroHashes[0][:])
	}
	someItem := byteslib.ToBytes32(item)
	m.branches[0][index] = someItem[:]
	if index >= len(m.originalItems) {
		m.originalItems = append(m.originalItems, someItem[:])
	} else {
		m.originalItems[index] = someItem[:]
	}
	currentIndex := index
	root := byteslib.ToBytes32(item)
	two := 2
	for i := 0; i < int(m.depth); i++ {
		isLeft := currentIndex%two == 0
		neighborIdx := currentIndex ^ 1
		var neighbor []byte
		if neighborIdx >= len(m.branches[i]) {
			neighbor = ZeroHashes[i][:]
		} else {
			neighbor = m.branches[i][neighborIdx]
		}
		if isLeft {
			parentHash := sha256.Sum256(append(root[:], neighbor...))
			root = parentHash
		} else {
			parentHash := sha256.Sum256(append(neighbor, root[:]...))
			root = parentHash
		}
		parentIdx := currentIndex / two
		if len(m.branches[i+1]) == 0 || parentIdx >= len(m.branches[i+1]) {
			newItem := root
			m.branches[i+1] = append(m.branches[i+1], newItem[:])
		} else {
			newItem := root
			m.branches[i+1][parentIdx] = newItem[:]
		}
		currentIndex = parentIdx
	}
	return nil
}

// MerkleProof computes a proof from a trie's branches using a Merkle index.
func (m *SparseMerkleTrie) MerkleProof(index int) ([][]byte, error) {
	if index < 0 {
		return nil, fmt.Errorf("merkle index is negative: %d", index)
	}
	leaves := m.branches[0]
	if index >= len(leaves) {
		return nil, fmt.Errorf(
			"merkle index out of range in trie, max range: %d, received: %d",
			len(leaves),
			index,
		)
	}
	merkleIndex := uint(index)
	proof := make([][]byte, m.depth+1)
	for i := uint(0); i < m.depth; i++ {
		subIndex := (merkleIndex / (1 << i)) ^ 1
		if subIndex < uint(len(m.branches[i])) {
			item := byteslib.ToBytes32(m.branches[i][subIndex])
			proof[i] = item[:]
		} else {
			proof[i] = ZeroHashes[i][:]
		}
	}
	var enc [32]byte
	binary.LittleEndian.PutUint64(enc[:], uint64(len(m.originalItems)))
	proof[len(proof)-1] = enc[:]
	return proof, nil
}

// VerifyMerkleProofWithDepth verifies a Merkle branch against a root of a trie.
func VerifyMerkleProofWithDepth(
	root, item []byte,
	merkleIndex uint64,
	proof [][]byte,
	depth uint64,
) bool {
	if uint64(len(proof)) != depth+1 {
		return false
	}
	node := byteslib.ToBytes32(item)
	for i := uint64(0); i <= depth; i++ {
		if (merkleIndex & 1) == 1 {
			node = sha256.Sum256(append(proof[i], node[:]...))
		} else {
			node = sha256.Sum256(append(node[:], proof[i]...))
		}
		merkleIndex /= 2
	}
	return bytes.Equal(root, node[:])
}

// VerifyMerkleProof given a trie root, a leaf, the generalized merkle index
// of the leaf in the trie, and the proof itself.
func VerifyMerkleProof(
	root, leaf []byte,
	merkleIndex uint64,
	proof [][]byte,
) bool {
	if len(proof) == 0 {
		return false
	}
	return VerifyMerkleProofWithDepth(
		root,
		leaf,
		merkleIndex,
		proof,
		uint64(len(proof)-1),
	)
}

// Copy performs a deep copy of the trie.
func (m *SparseMerkleTrie) Copy() *SparseMerkleTrie {
	dstBranches := make([][][]byte, len(m.branches))
	for i1, srcB1 := range m.branches {
		dstBranches[i1] = byteslib.SafeCopy2d(srcB1)
	}

	return &SparseMerkleTrie{
		depth:         m.depth,
		branches:      dstBranches,
		originalItems: byteslib.SafeCopy2d(m.originalItems),
	}
}

// NumOfItems returns the num of items stored in
// the sparse merkle trie. We handle a special case
// where if there is only one item stored and it is an
// empty 32-byte root.
func (m *SparseMerkleTrie) NumOfItems() int {
	var zeroBytes [32]byte
	if len(m.originalItems) == 1 &&
		bytes.Equal(m.originalItems[0], zeroBytes[:]) {
		return 0
	}
	return len(m.originalItems)
}
