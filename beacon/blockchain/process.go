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
	"fmt"

	"github.com/berachain/beacon-kit/beacon/core"
	beacontypes "github.com/berachain/beacon-kit/beacon/core/types"
	bls12381 "github.com/berachain/beacon-kit/crypto/bls12-381"
	"github.com/berachain/beacon-kit/crypto/kzg"
)

// ProcessBeaconBlock receives an incoming beacon block, it first validates
// and then processes the block.
func (s *Service) ProcessBeaconBlock(
	ctx context.Context,
	blk beacontypes.ReadOnlyBeaconBlock,
	proposerPubkey [bls12381.PubKeyLength]byte,
	blockHash [32]byte,
) error {
	// If we get any sort of error from the execution client, we bubble
	// it up and reject the proposal, as we do not want to write a block
	// finalization to the consensus layer that is invalid.
	var (
		// eg, groupCtx   = errgroup.WithContext(ctx)
		isValidPayload bool
	)

	// If the block is nil, We have to abort.
	if blk == nil || blk.IsNil() {
		return beacontypes.ErrNilBlk
	}

	// TODO:
	// expectedProposer, err := epc.GetBeaconProposer(benv.Slot)

	// This go rountine validates the execution level aspects of the block.
	// i.e: does newPayload return VALID?
	if _, err := s.validateExecutionOnBlock(
		ctx, blk,
	); err != nil {
		s.Logger().
			Error("failed to notify engine of new payload", "error", err)
		return err
	}

	// This go routine validates the consensus level aspects of the block.
	// i.e: does it have a valid ancestor?
	err := s.validateStateTransition(ctx, blk, proposerPubkey)
	if err != nil {
		s.Logger().
			Error("failed to validate state transition", "error", err)
		return err
	}

	// TODO: This is very much the wrong spot for this.
	if err = s.rp.MixinNewReveal(ctx, blk); err != nil {
		return err
	}

	// daStartTime := time.Now()
	// if avs != nil {
	// avs.IsDataAvailable(ctx, s.CurrentSlot(), rob); err != nil {
	// 		return errors.Wrap(err, "could not validate blob data availability
	// (AvailabilityStore.IsDataAvailable)")
	// 	}
	// } else {
	// s.isDataAvailable(ctx, blockRoot, blockCopy); err != nil {
	// 		return errors.Wrap(err, "could not validate blob data availability")
	// 	}
	// }

	// // Wait for the goroutines to finish.
	// if err := eg.Wait(); err != nil {
	// 	return err
	// }

	// Perform post block processing.
	return s.postBlockProcess(
		ctx, blk, blockHash, isValidPayload,
	)
}

// validateStateTransition checks a block's state transition.
// TODO: Expand rules, consider modularity. Current implementation
// is hardcoded for single slot finality, which works but lacks flexibility.
func (s *Service) validateStateTransition(
	ctx context.Context, blk beacontypes.ReadOnlyBeaconBlock,
	proposerPubKey [bls12381.PubKeyLength]byte,
) error {
	// Ensure the parent block root matches what we have locally.
	// TODO: get rid of CometBFT stuff.
	parentBlockRoot := s.BeaconState(ctx).GetParentBlockRoot()
	if parentBlockRoot != blk.GetParentBlockRoot() {
		return fmt.Errorf(
			"parent root does not match, expected: %x, got: %x",
			parentBlockRoot,
			blk.GetParentBlockRoot(),
		)
	}

	// Create a new state processor.
	sp := core.NewStateProcessor(
		s.BeaconCfg(),
		s.BeaconState(ctx),
	)

	// Verify the RANDAO Reveal.
	// TODO: move into state processor.
	if err := s.rp.VerifyReveal(
		proposerPubKey,
		s.BeaconCfg().SlotToEpoch(blk.GetSlot()),
		blk.GetRandaoReveal(),
	); err != nil {
		return err
	}

	// ---------------------///
	//   VALIDATE KZG HERE  ///
	// ---------------------///

	// ---------------------///
	//   Process Deposits   ///
	// ---------------------///

	return sp.ProcessBlock(
		blk,
	)
}

// validateExecutionOnBlock checks the validity of a the execution payload
// on the beacon block.
func (s *Service) validateExecutionOnBlock(
	// todo: parentRoot hashs should be on blk.
	ctx context.Context,
	blk beacontypes.ReadOnlyBeaconBlock,
) (bool, error) {
	body := blk.GetBody()
	payload := body.GetExecutionPayload()
	if payload.IsNil() {
		return false, beacontypes.ErrNilPayloadInBlk
	}

	// In BeaconKit, since we are currently operating on SingleSlot Finality
	// we purposefully reject any block that is not a child of the last
	// finalized block.
	safeHash := s.ForkchoiceStore(ctx).JustifiedPayloadBlockHash()
	if safeHash != payload.GetParentHash() {
		return false, fmt.Errorf(
			"parent block with hash %x is not finalized, expected finalized hash %x",
			payload.GetParentHash(),
			safeHash,
		)
	}

	expectedMix, err := s.BeaconState(ctx).RandaoMix()
	if err != nil {
		return false, err
	}

	// Ensure the prev randao matches the local state.
	if payload.GetPrevRandao() != expectedMix {
		return false, fmt.Errorf(
			"prev randao does not match, expected: %x, got: %x",
			expectedMix, payload.GetPrevRandao(),
		)
	}

	// if expectedTime, err := spec.TimeAtSlot(slot, genesisTime); err != nil {
	// 	return fmt.Errorf("slot or genesis time in state is corrupt, cannot
	// compute time: %v", err)
	// } else if payload.Timestamp != expectedTime {
	// 	return fmt.Errorf("state at slot %d, genesis time %d, expected execution
	// payload time %d, but got %d",
	// 		slot, genesisTime, expectedTime, payload.Timestamp)
	// }

	// TODO: add some more safety checks here.
	return s.es.NotifyNewPayload(
		ctx,
		blk.GetSlot(),
		payload,
		kzg.ConvertCommitmentsToVersionedHashes(
			body.GetBlobKzgCommitments(),
		),
		blk.GetParentBlockRoot(),
	)
}
