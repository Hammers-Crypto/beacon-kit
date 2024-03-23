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

package beacon

import (
	sdkcollections "cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
	beacontypes "github.com/berachain/beacon-kit/beacon/core/types"
	bls "github.com/cosmos/cosmos-sdk/crypto/keys/bls12_381"
)

// validatorsIndex is a structure that holds a unique index for validators based
// on their public key.
type validatorsIndex struct {
	// Pubkey is a unique index mapping a validator's public key to their
	// numeric ID and vice versa.
	Pubkey   *indexes.Unique[[]byte, uint64, *beacontypes.Validator]
	ConsAddr *indexes.Unique[[]byte, uint64, *beacontypes.Validator]
}

// IndexesList returns a list of all indexes associated with the
// validatorsIndex.
func (a validatorsIndex) IndexesList() []sdkcollections.Index[
	uint64, *beacontypes.Validator,
] {
	return []sdkcollections.Index[uint64, *beacontypes.Validator]{
		a.Pubkey,
		a.ConsAddr,
	}
}

// NewValidatorsIndex creates a new validatorsIndex with a unique index for
// validator public keys.
func newValidatorsIndex(sb *sdkcollections.SchemaBuilder) validatorsIndex {
	return validatorsIndex{
		Pubkey: indexes.NewUnique(
			sb,
			sdkcollections.NewPrefix(validatorPubkeyToIndexPrefix),
			validatorPubkeyToIndexPrefix,
			sdkcollections.BytesKey,
			sdkcollections.Uint64Key,
			// The mapping function simply returns the public key as the index
			// key.
			func(_ uint64, validator *beacontypes.Validator) ([]byte, error) {
				return validator.Pubkey[:], nil
			},
		),
		ConsAddr: indexes.NewUnique(
			sb,
			sdkcollections.NewPrefix(validatorConsAddrToIndexPrefix),
			validatorConsAddrToIndexPrefix,
			sdkcollections.BytesKey,
			sdkcollections.Uint64Key,
			// The mapping function simply returns the consensus address as the
			// index key.
			func(_ uint64, validator *beacontypes.Validator) ([]byte, error) {
				addr := (&bls.PubKey{Key: validator.Pubkey[:]}).
					Address()
				return addr, nil
			},
		),
	}
}
