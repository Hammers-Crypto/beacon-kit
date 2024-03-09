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

package state

import (
	beacontypes "github.com/berachain/beacon-kit/beacon/core/types"
	enginetypes "github.com/berachain/beacon-kit/engine/types"
)

// BeaconState is the interface for the beacon state. It
// is a combination of the read-only and write-only beacon state consensus.
type BeaconState interface {
	ReadOnlyBeaconState
	WriteOnlyBeaconState
}

// ReadOnlyBeaconState is the interface for a read-only beacon state.
type ReadOnlyBeaconState interface {
	// TODO: fill these in as we develop impl
	ReadWriteDepositQueue
	ReadOnlyWithdrawals

	GetParentBlockRoot() [32]byte

	// TODO: Actually decouple epocha nd slot
	// GetEpochBySlot(primitives.Slot) primitives.Epoch
}

// WriteOnlyBeaconState is the interface for a write-only beacon state.
type WriteOnlyBeaconState interface {
	SetParentBlockRoot([32]byte)
}

// ReadWriteDepositQueue has read and write access to deposit queue.
type ReadWriteDepositQueue interface {
	EnqueueDeposits([]*beacontypes.Deposit) error
	DequeueDeposits(n uint64) ([]*beacontypes.Deposit, error)
}

// ReadOnlyWithdrawals only has read access to withdrawal methods.
type ReadOnlyWithdrawals interface {
	ExpectedWithdrawals() ([]*enginetypes.Withdrawal, error)
}
