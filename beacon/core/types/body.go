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

package types

import (
	randaotypes "github.com/berachain/beacon-kit/beacon/core/randao/types"
	"github.com/berachain/beacon-kit/crypto/sha256"
	enginetypes "github.com/berachain/beacon-kit/engine/types"
	"github.com/berachain/beacon-kit/lib/encoding/ssz"
	"github.com/cockroachdb/errors"
)

// BeaconBlockBodyDeneb represents the body of a beacon block in the Deneb
// chain.
type BeaconBlockBodyDeneb struct {
	// RandaoReveal is the reveal of the RANDAO.
	RandaoReveal [96]byte `ssz-size:"96"`
	// Graffiti is for a fun message or meme.
	Graffiti [32]byte `ssz-size:"32"`
	// Deposits is the list of deposits included in the body.
	Deposits []*Deposit `                ssz-max:"16"`
	// ExecutionPayload is the execution payload of the body.
	ExecutionPayload *enginetypes.ExecutableDataDeneb
	// BlobKzgCommitments is the list of KZG commitments for the EIP-4844 blobs.
	BlobKzgCommitments [][48]byte `ssz-size:"?,48" ssz-max:"16"`
}

// IsNil checks if the BeaconBlockBodyDeneb is nil.
func (b *BeaconBlockBodyDeneb) IsNil() bool {
	return b == nil
}

// GetBlobKzgCommitments returns the BlobKzgCommitments of the Body.
func (b *BeaconBlockBodyDeneb) GetBlobKzgCommitments() [][48]byte {
	return b.BlobKzgCommitments
}

// GetRandaoReveal returns the RandaoReveal of the Body.
func (b *BeaconBlockBodyDeneb) GetRandaoReveal() randaotypes.Reveal {
	return b.RandaoReveal
}

// GetExecutionPayload returns the ExecutionPayload of the Body.
//
//nolint:lll
func (b *BeaconBlockBodyDeneb) GetExecutionPayload() enginetypes.ExecutionPayload {
	return b.ExecutionPayload
}

// GetDeposits returns the Deposits of the BeaconBlockBodyDeneb.
func (b *BeaconBlockBodyDeneb) GetDeposits() []*Deposit {
	return b.Deposits
}

// SetDeposits sets the Deposits of the BeaconBlockBodyDeneb.
func (b *BeaconBlockBodyDeneb) SetDeposits(deposits []*Deposit) {
	b.Deposits = deposits
}

// SetExecutionData sets the ExecutionData of the BeaconBlockBodyDeneb.
func (b *BeaconBlockBodyDeneb) SetExecutionData(
	executionData enginetypes.ExecutionPayload,
) error {
	var ok bool
	b.ExecutionPayload, ok = executionData.(*enginetypes.ExecutableDataDeneb)
	if !ok {
		return errors.New("invalid execution data type")
	}
	return nil
}

// SetBlobKzgCommitments sets the BlobKzgCommitments of the
// BeaconBlockBodyDeneb.
func (b *BeaconBlockBodyDeneb) SetBlobKzgCommitments(commitments [][48]byte) {
	b.BlobKzgCommitments = commitments
}

// If you are adding values to the BeaconBlockBodyDeneb struct,
// the body length must be increased and GetTopLevelRoots updated.
const bodyLength = 5

func (b *BeaconBlockBodyDeneb) GetTopLevelRoots() ([][]byte, error) {
	layer := make([][]byte, bodyLength)
	for i := range layer {
		layer[i] = make([]byte, 32)
	}

	randao := b.RandaoReveal
	root, err := ssz.MerkleizeByteSliceSSZ(randao[:])
	if err != nil {
		return nil, err
	}
	copy(layer[0], root[:])

	// graffiti
	root = b.Graffiti
	copy(layer[1], root[:])

	// Deposits
	dep := b.Deposits
	root, err = sha256.BuildMerkleRoot(dep, 16)
	if err != nil {
		return nil, err
	}
	copy(layer[3], root[:])

	// Execution Payload
	rt, err := b.ExecutionPayload.HashTreeRoot()
	if err != nil {
		return nil, err
	}

	copy(layer[4], rt[:])

	return layer, nil
}

func (b *BeaconBlockBodyDeneb) AttachExecution(
	executionData enginetypes.ExecutionPayload,
) error {
	var ok bool
	b.ExecutionPayload, ok = executionData.(*enginetypes.ExecutableDataDeneb)
	if !ok {
		return errors.New("invalid execution data type")
	}
	return nil
}

func (b *BeaconBlockBodyDeneb) GetKzgCommitments() [][48]byte {
	return b.BlobKzgCommitments
}
