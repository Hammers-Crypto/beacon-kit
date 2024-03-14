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

package core

import (
	"fmt"

	"github.com/berachain/beacon-kit/beacon/core/state"
	"github.com/berachain/beacon-kit/beacon/core/types"
	"github.com/berachain/beacon-kit/config"
	enginetypes "github.com/berachain/beacon-kit/engine/types"
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
type StateProcessor struct {
	cfg *config.Beacon
	st  state.BeaconState
}

// NewStateProcessor creates a new state processor.
func NewStateProcessor(
	cfg *config.Beacon,
	st state.BeaconState,
) *StateProcessor {
	return &StateProcessor{
		cfg: cfg,
		st:  st,
	}
}

// ProcessSlot processes the slot and ensures it matches the local state.
func (sp *StateProcessor) ProcessSlot(
	_ uint64,
) error {
	return nil
}

// ProcessBlock processes the block and ensures it matches the local state.
func (sp *StateProcessor) ProcessBlock(
	blk types.BeaconBlock,
) error {
	// Ensure Body is non nil.
	body := blk.GetBody()
	if body.IsNil() {
		return types.ErrNilBlkBody
	}

	// process the eth1 vote.
	payload := body.GetExecutionPayload()
	if payload.IsNil() {
		return types.ErrNilPayloadInBlk
	}

	// common.ProcessHeader

	// process the withdrawals.
	if err := sp.processWithdrawals(payload.GetWithdrawals()); err != nil {
		return err
	}

	// phase0.ProcessProposerSlashings
	// phase0.ProcessAttesterSlashings

	// process the randao reveal.
	if err := sp.processRandaoReveal(); err != nil {
		return err
	}

	// phase0.ProcessEth1Vote ? forkchoice?

	// process the deposits and ensure they match the local state.
	if err := sp.processDeposits(body.GetDeposits()); err != nil {
		return err
	}

	// ProcessVoluntaryExits

	return nil
}

// ProcessBlob processes a blob.
func (sp *StateProcessor) ProcessBlob() error {
	// TODO: 4844.
	return nil
}

// ProcessDeposits processes the deposits and ensures they match the
// local state.
func (sp *StateProcessor) processDeposits(
	deposits []*types.Deposit,
) error {
	if uint64(len(deposits)) > sp.cfg.Limits.MaxDepositsPerBlock {
		return fmt.Errorf(
			"too many deposits, expected: %d, got: %d",
			sp.cfg.Limits.MaxDepositsPerBlock, len(deposits),
		)
	}

	// Dequeue and verify the logs.
	localDeposits, err := sp.st.ExpectedDeposits(uint64(len(deposits)))
	if err != nil {
		return err
	}

	// Ensure the deposits match the local state.
	for i, dep := range deposits {
		if dep == nil {
			return types.ErrNilDeposit
		}
		if dep.Index != localDeposits[i].Index {
			return fmt.Errorf(
				"deposit index does not match, expected: %d, got: %d",
				localDeposits[i].Index, dep.Index)
		}
	}
	return nil
}

// processWithdrawals processes the withdrawals and ensures they match the
// local state.
func (sp *StateProcessor) processWithdrawals(
	withdrawals []*enginetypes.Withdrawal,
) error {
	if uint64(len(withdrawals)) > sp.cfg.Limits.MaxWithdrawalsPerPayload {
		return fmt.Errorf(
			"too many withdrawals, expected: %d, got: %d",
			sp.cfg.Limits.MaxWithdrawalsPerPayload, len(withdrawals),
		)
	}

	// Dequeue and verify the withdrawals.
	localWithdrawals, err := sp.st.DequeueWithdrawals(uint64(len(withdrawals)))
	if err != nil {
		return err
	}

	// Ensure the deposits match the local state.
	for i, dep := range withdrawals {
		if dep == nil {
			return types.ErrNilWithdrawal
		}
		if dep.Index != localWithdrawals[i].Index {
			return fmt.Errorf(
				"deposit index does not match, expected: %d, got: %d",
				localWithdrawals[i].Index, dep.Index)
		}
	}
	return nil
}

func (sp *StateProcessor) processRandaoReveal() error {
	return nil
}
