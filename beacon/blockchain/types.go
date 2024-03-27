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

package blockchain

import (
	"context"

	"github.com/berachain/beacon-kit/beacon/core/state"
	enginetypes "github.com/berachain/beacon-kit/engine/types"
	"github.com/berachain/beacon-kit/primitives"
)

// LocalBuilder is the interface for the builder service.
type LocalBuilder interface {
	BuildLocalPayload(
		ctx context.Context,
		parentEth1Hash primitives.ExecutionHash,
		slot primitives.Slot,
		timestamp uint64,
		parentBlockRoot primitives.Root,
	) (*enginetypes.PayloadID, error)
}

// RandaoProcessor is the interface for the randao processor.
type RandaoProcessor interface {
	BuildReveal(
		st state.BeaconState,
	) (primitives.BLSSignature, error)
	MixinNewReveal(
		st state.BeaconState,
		reveal primitives.BLSSignature,
	) error
	VerifyReveal(
		st state.BeaconState,
		proposerPubkey primitives.BLSPubkey,
		reveal primitives.BLSSignature,
	) error
}

// StakingService is the interface for the staking service.
type StakingService interface {
	// ProcessLogsInETH1Block processes logs in an eth1 block.
	ProcessLogsInETH1Block(
		ctx context.Context,
		blockHash primitives.ExecutionHash,
	) error
}

type SyncService interface {
	IsInitSync() bool
	Status() error
}
