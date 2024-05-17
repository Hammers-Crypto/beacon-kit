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
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/constants"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/state-transition/pkg/core/state"
)

// StateProcessor is a basic Processor, which takes care of the
// main state transition for the beacon chain.
type StateProcessor[
	BeaconBlockT types.BeaconBlock,
	BeaconStateT state.BeaconState,
	BlobSidecarsT interface{ Len() int },
] struct {
	cs              primitives.ChainSpec
	bp              BlobProcessor[BlobSidecarsT]
	rp              RandaoProcessor[BeaconBlockT, BeaconStateT]
	signer          crypto.BLSSigner
	logger          log.Logger[any]
	executionEngine ExecutionEngine
	// DepositProcessor
	// WithdrawalProcessor
}

// NewStateProcessor creates a new state processor.
func NewStateProcessor[
	BeaconBlockT types.BeaconBlock,
	BeaconStateT state.BeaconState,
	BlobSidecarsT interface{ Len() int },
](
	cs primitives.ChainSpec,
	bp BlobProcessor[BlobSidecarsT],
	rp RandaoProcessor[BeaconBlockT, BeaconStateT],
	executionEngine ExecutionEngine,
	signer crypto.BLSSigner,
	logger log.Logger[any],
) *StateProcessor[BeaconBlockT, BeaconStateT, BlobSidecarsT] {
	return &StateProcessor[BeaconBlockT, BeaconStateT, BlobSidecarsT]{
		cs:              cs,
		bp:              bp,
		rp:              rp,
		executionEngine: executionEngine,
		signer:          signer,
		logger:          logger,
	}
}

// Transition is the main function for processing a state transition.
func (sp *StateProcessor[
	BeaconBlockT, BeaconStateT, BlobSidecarsT,
]) Transition(
	ctx Context,
	st BeaconStateT,
	blk BeaconBlockT,
	/*validateSignature bool, */
) error {
	// Process the slot.
	if err := sp.ProcessSlot(st); err != nil {
		return err
	}

	// Process the block.
	if err := sp.ProcessBlock(ctx, st, blk); err != nil {
		return err
	}

	return nil
}

// ProcessSlot is run when a slot is missed.
func (sp *StateProcessor[
	BeaconBlockT, BeaconStateT, BlobSidecarsT,
]) ProcessSlot(
	st BeaconStateT,
) error {
	slot, err := st.GetSlot()
	if err != nil {
		return err
	}

	// Before we make any changes, we calculate the previous state root.
	prevStateRoot, err := st.HashTreeRoot()
	if err != nil {
		return err
	}

	// We update our state roots and block roots.
	if err = st.UpdateStateRootAtIndex(
		uint64(slot)%sp.cs.SlotsPerHistoricalRoot(),
		prevStateRoot,
	); err != nil {
		return err
	}

	// We get the latest block header, this will not have
	// a state root on it.
	latestHeader, err := st.GetLatestBlockHeader()
	if err != nil {
		return err
	}

	// We set the "rawHeader" in the StateProcessor, but cannot fill in
	// the StateRoot until the following block.
	if (latestHeader.StateRoot == primitives.Root{}) {
		latestHeader.StateRoot = prevStateRoot
		if err = st.SetLatestBlockHeader(latestHeader); err != nil {
			return err
		}
	}

	// We update the block root.
	var prevBlockRoot primitives.Root
	prevBlockRoot, err = latestHeader.HashTreeRoot()
	if err != nil {
		return err
	}

	if err = st.UpdateBlockRootAtIndex(
		uint64(slot)%sp.cs.SlotsPerHistoricalRoot(), prevBlockRoot,
	); err != nil {
		return err
	}

	// Process the Epoch Boundary.
	if uint64(slot+1)%sp.cs.SlotsPerEpoch() == 0 {
		if err = sp.processEpoch(st); err != nil {
			return err
		}
		sp.logger.Info(
			"processed epoch transition 🔃",
			"old", uint64(slot)/sp.cs.SlotsPerEpoch(),
			"new", uint64(slot+1)/sp.cs.SlotsPerEpoch(),
		)
	}

	return st.SetSlot(slot + 1)
}

// ProcessBlock processes the block and ensures it matches the local state.
//
//nolint:funlen // todo fix.
func (sp *StateProcessor[
	BeaconBlockT, BeaconStateT, BlobSidecarsT,
]) ProcessBlock(
	ctx Context,
	st BeaconStateT,
	blk BeaconBlockT,
) error {
	// process the freshly created header.
	if err := sp.processBlockHeader(st, blk); err != nil {
		return err
	}

	body := blk.GetBody()

	// process the execution payload.
	if err := sp.processExecutionPayload(
		ctx, st, blk,
	); err != nil {
		return err
	}

	// process the withdrawals.
	if err := sp.processWithdrawals(
		st, body.GetExecutionPayload(),
	); err != nil {
		return err
	}

	// phase0.ProcessProposerSlashings
	// phase0.ProcessAttesterSlashings

	// process the randao reveal.
	if err := sp.processRandaoReveal(st, blk); err != nil {
		return err
	}

	// phase0.ProcessEth1Vote ? forkchoice?

	// TODO: LOOK HERE
	//
	// process the deposits and ensure they match the local state.
	if err := sp.processOperations(st, body); err != nil {
		return err
	}

	if ctx.GetValidateResult() {
		// Ensure the state root matches the block.
		//
		// TODO: We need to validate this in ProcessProposal as well.
		if stateRoot, err := st.HashTreeRoot(); err != nil {
			return err
		} else if blk.GetStateRoot() != stateRoot {
			return errors.Wrapf(
				ErrStateRootMismatch, "expected %s, got %s",
				primitives.Root(stateRoot), blk.GetStateRoot(),
			)
		}
	}
	return nil
}

// processEpoch processes the epoch and ensures it matches the local state.
func (sp *StateProcessor[
	BeaconBlockT, BeaconStateT, BlobSidecarsT,
]) processEpoch(
	st BeaconStateT,
) error {
	var err error
	if err = sp.processRewardsAndPenalties(st); err != nil {
		return err
	}
	if err = sp.processSlashingsReset(st); err != nil {
		return err
	}
	if err = sp.processRandaoMixesReset(st); err != nil {
		return err
	}
	return nil
}

// processBlockHeader processes the header and ensures it matches the local
// state.
func (sp *StateProcessor[
	BeaconBlockT, BeaconStateT, BlobSidecarsT,
]) processBlockHeader(
	st BeaconStateT,
	blk BeaconBlockT,
) error {
	// Get the current slot.
	slot, err := st.GetSlot()
	if err != nil {
		return err
	}

	// Ensure the block slot matches the state slot.
	if blk.GetSlot() != slot {
		return errors.Newf(
			"slot does not match, expected: %d, got: %d",
			slot,
			blk.GetSlot(),
		)
	}

	latestBlockHeader, err := st.GetLatestBlockHeader()
	if err != nil {
		return err
	}

	if blk.GetSlot() <= latestBlockHeader.GetSlot() {
		return errors.Newf(
			"block slot is too low, expected: > %d, got: %d",
			latestBlockHeader.GetSlot(),
			blk.GetSlot(),
		)
	}

	// Ensure the parent root matches the latest block header.
	parentBlockRoot, err := latestBlockHeader.HashTreeRoot()
	if err != nil {
		return err
	}

	// Ensure the parent root matches the latest block header.
	if parentBlockRoot != blk.GetParentBlockRoot() {
		return errors.Newf(
			"parent root does not match, expected: %x, got: %x",
			parentBlockRoot,
			blk.GetParentBlockRoot(),
		)
	}

	// Ensure the block is within the acceptable range.
	// TODO: move this is in the wrong spot.
	if deposits := blk.GetBody().GetDeposits(); uint64(
		len(deposits),
	) > sp.cs.MaxDepositsPerBlock() {
		return errors.Newf(
			"too many deposits, expected: %d, got: %d",
			sp.cs.MaxDepositsPerBlock(), len(deposits),
		)
	}

	// Ensure the proposer is not slashed.
	bodyRoot, err := blk.GetBody().HashTreeRoot()
	if err != nil {
		return err
	}

	// Store as the new latest block
	if err = st.SetLatestBlockHeader(
		types.NewBeaconBlockHeader(
			blk.GetSlot(),
			blk.GetProposerIndex(),
			blk.GetParentBlockRoot(),
			// state_root is zeroed and overwritten
			// in the next `process_slot` call.
			[32]byte{},
			bodyRoot),
	); err != nil {
		return err
	}

	proposer, err := st.ValidatorByIndex(blk.GetProposerIndex())
	if err != nil {
		return err
	}

	// Verify the proposer is not slashed.
	if proposer.Slashed {
		return errors.Newf(
			"proposer is slashed, index: %d",
			blk.GetProposerIndex(),
		)
	}
	return nil
}

// getAttestationDeltas as defined in the Ethereum 2.0 specification.
// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#get_attestation_deltas
//
//nolint:lll
func (sp *StateProcessor[
	BeaconBlockT, BeaconStateT, BlobSidecarsT,
]) getAttestationDeltas(
	st BeaconStateT,
) ([]math.Gwei, []math.Gwei, error) {
	// TODO: implement this function forreal
	validators, err := st.GetValidators()
	if err != nil {
		return nil, nil, err
	}
	placeholder := make([]math.Gwei, len(validators))
	return placeholder, placeholder, nil
}

// processRewardsAndPenalties as defined in the Ethereum 2.0 specification.
// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#process_rewards_and_penalties
//
//nolint:lll
func (sp *StateProcessor[
	BeaconBlockT, BeaconStateT, BlobSidecarsT,
]) processRewardsAndPenalties(
	st BeaconStateT,
) error {
	slot, err := st.GetSlot()
	if err != nil {
		return err
	}

	if sp.cs.SlotToEpoch(slot) == math.U64(constants.GenesisEpoch) {
		return nil
	}

	rewards, penalties, err := sp.getAttestationDeltas(st)
	if err != nil {
		return err
	}

	validators, err := st.GetValidators()
	if err != nil {
		return err
	}

	if len(validators) != len(rewards) || len(validators) != len(penalties) {
		return errors.Newf(
			"mismatched rewards and penalties lengths: %d, %d, %d",
			len(validators), len(rewards), len(penalties),
		)
	}

	for i := range validators {
		// Increase the balance of the validator.
		if err = st.IncreaseBalance(
			math.ValidatorIndex(i),
			rewards[i],
		); err != nil {
			return err
		}

		// Decrease the balance of the validator.
		if err = st.DecreaseBalance(
			math.ValidatorIndex(i),
			penalties[i],
		); err != nil {
			return err
		}
	}

	return nil
}
