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

package builder

import (
	"context"
	"fmt"
	"time"

	"github.com/berachain/beacon-kit/mod/core/state"
	"github.com/berachain/beacon-kit/mod/execution"
	"github.com/berachain/beacon-kit/mod/primitives"
	engineprimitives "github.com/berachain/beacon-kit/mod/primitives-engine"
	"github.com/berachain/beacon-kit/mod/primitives/math"
)

// BuildLocalPayload builds a payload for the given slot and
// returns the payload ID.
func (pb *PayloadBuilder) BuildLocalPayload(
	ctx context.Context,
	st state.BeaconState,
	parentEth1Hash primitives.ExecutionHash,
	slot math.Slot,
	timestamp uint64,
	parentBlockRoot primitives.Root,
) (*engineprimitives.PayloadID, error) {
	// Assemble the payload attributes.
	attrs, err := pb.getPayloadAttribute(st, slot, timestamp, parentBlockRoot)
	if err != nil {
		return nil, fmt.Errorf("%w error when getting payload attributes", err)
	}

	// Notify the execution client of the forkchoice update.
	var payloadID *engineprimitives.PayloadID
	pb.logger.Info(
		"bob the builder; can we fix it; bob the builder; yes we can 🚧",
		"for_slot", slot,
		"parent_eth1_hash", parentEth1Hash,
		"parent_block_root", parentBlockRoot,
	)

	latestExecutionPayload, err := st.GetLatestExecutionPayload()
	if err != nil {
		return nil, err
	}
	parentEth1BlockHash := latestExecutionPayload.GetBlockHash()

	payloadID, _, err = pb.ee.NotifyForkchoiceUpdate(
		ctx, &execution.ForkchoiceUpdateRequest{
			State: &engineprimitives.ForkchoiceState{
				HeadBlockHash:      parentEth1Hash,
				SafeBlockHash:      parentEth1BlockHash,
				FinalizedBlockHash: parentEth1BlockHash,
			},
			PayloadAttributes: attrs,
			ForkVersion:       pb.chainSpec.ActiveForkVersionForSlot(slot),
		},
	)
	if err != nil {
		return nil, err
	} else if payloadID == nil {
		pb.logger.Warn("received nil payload ID on VALID engine response",
			"head_eth1_hash", parentEth1Hash,
			"for_slot", slot,
		)

		return payloadID, ErrNilPayloadOnValidResponse
	}

	pb.logger.Info("forkchoice updated with payload attributes",
		"head_eth1_hash", parentEth1Hash,
		"for_slot", slot,
		"payload_id", payloadID,
	)

	pb.pc.Set(
		slot,
		parentBlockRoot,
		*payloadID,
	)

	return payloadID, nil
}

// GetBestPayload attempts to pull a previously built payload
// by reading a payloadID from the builder's cache. If it fails to
// retrieve a payload, it will build a new payload and wait for the
// execution client to return the payload.
func (pb *PayloadBuilder) GetBestPayload(
	ctx context.Context,
	st state.BeaconState,
	slot math.Slot,
	parentBlockRoot primitives.Root,
	parentEth1Hash primitives.ExecutionHash,
) (engineprimitives.ExecutionPayload, engineprimitives.BlobsBundle, bool, error) {
	// TODO: Proposer-Builder Separation Improvements Later.
	// val, tracked := s.TrackedValidatorsCache.Validator(vIdx)
	// if !tracked {
	// 	logrus.WithFields(logFields).Warn("could not find tracked proposer
	// index")
	// }

	// We first attempt to see if we previously fired off a payload built for
	// this particular slot and parent block root. If we have, and we are able
	// to
	// retrieve it from our execution client, we can return it immediately.
	payload, blobsBundle, overrideBuilder, err := pb.
		requestBuiltPayloadFromExecutionClient(
			ctx,
			parentBlockRoot,
			slot,
		)

	// If there was no error we can simply return the payload that we
	// just retrieved.
	if err == nil {
		return payload, blobsBundle, overrideBuilder, nil
	}

	// Otherwise we will fall back to triggering a payload build.
	return pb.buildAndWaitForLocalPayload(
		ctx,
		st,
		parentEth1Hash,
		slot,
		// TODO: we need to do the proper timestamp math here for EIP4788.
		//#nosec:G701 // won't realistically overflow.
		uint64(time.Now().Unix()),
		parentBlockRoot,
	)
}

// requestBuiltPayloadFromExecutionClient retrieves the payload and blobs
// bundle.
func (pb *PayloadBuilder) requestBuiltPayloadFromExecutionClient(
	ctx context.Context,
	parentBlockRoot primitives.Root,
	slot math.Slot,
) (engineprimitives.ExecutionPayload, engineprimitives.BlobsBundle, bool, error) {
	// See if we have a payload ID for this slot and parent block root.
	payloadID, found := pb.pc.Get(slot, parentBlockRoot)
	if !found || (payloadID == engineprimitives.PayloadID{}) {
		// If we don't have a payload ID, we can't retrieve the payload.
		return nil, nil, false, ErrPayloadIDNotFound
	}

	// Request the payload from the execution client.
	payload, blobsBundle, overrideBuilder, err := pb.getPayloadFromExecutionClient(
		ctx,
		&payloadID,
		slot,
	)
	if err != nil {
		return nil, nil, false, err
	} else if payload == nil {
		return nil, nil, false, ErrNilPayload
	}

	// Cache the payload and return.
	pb.pc.Set(slot, payload.GetParentHash(), payloadID)
	return payload, blobsBundle, overrideBuilder, nil
}

// buildAndWaitForLocalPayload, triggers a payload build process, waits
// for a configuration specified period, and then retrieves the built
// payload from the execution client.
func (pb *PayloadBuilder) buildAndWaitForLocalPayload(
	ctx context.Context,
	st state.BeaconState,
	parentEth1Hash primitives.ExecutionHash,
	slot math.Slot,
	timestamp uint64,
	parentBlockRoot primitives.Root,
) (engineprimitives.ExecutionPayload,
	engineprimitives.BlobsBundle, bool, error) {
	// Build the payload and wait for the execution client to return the payload
	// ID.
	payloadID, err := pb.BuildLocalPayload(
		ctx, st, parentEth1Hash, slot, timestamp, parentBlockRoot,
	)
	if err != nil {
		return nil, nil, false, err
	}

	// Wait for the payload to be delivered to the execution client.
	pb.logger.Info(
		"waiting for local payload to be delivered to execution client",
		"for_slot", slot, "timeout", pb.cfg.LocalBuildPayloadTimeout.String(),
	)
	select {
	case <-time.After(pb.cfg.LocalBuildPayloadTimeout):
		// We want to trigger delivery of the payload to the execution client
		// before the timestamp expires.
		break
	case <-ctx.Done():
		return nil, nil, false, ctx.Err()
	}

	// Get the payload from the execution client.
	return pb.getPayloadFromExecutionClient(
		ctx, payloadID, slot)
}

// getPayloadAttributes returns the payload attributes for the given state and
// slot. The attribute is required to initiate a payload build process in the
// context of an `engine_forkchoiceUpdated` call.
func (pb *PayloadBuilder) getPayloadAttribute(
	st state.BeaconState,
	slot math.Slot,
	timestamp uint64,
	prevHeadRoot [32]byte,
) (engineprimitives.PayloadAttributer, error) {
	var (
		prevRandao [32]byte
	)

	// Get the expected withdrawals to include in this payload.
	withdrawals, err := st.ExpectedWithdrawals()
	if err != nil {
		pb.logger.Error(
			"Could not get expected withdrawals to get payload attribute",
			"error",
			err,
		)
		return nil, err
	}

	epoch := pb.chainSpec.SlotToEpoch(slot)

	// Get the previous randao mix.
	prevRandao, err = st.GetRandaoMixAtIndex(
		uint64(epoch) % pb.chainSpec.EpochsPerHistoricalVector(),
	)
	if err != nil {
		return nil, err
	}

	return engineprimitives.NewPayloadAttributes[*engineprimitives.Withdrawal](
		pb.chainSpec.ActiveForkVersionForEpoch(epoch),
		timestamp,
		prevRandao,
		pb.cfg.SuggestedFeeRecipient,
		withdrawals,
		prevHeadRoot,
	)
}

// getPayloadFromExecutionClient retrieves the payload and blobs bundle for the
// given slot.
func (pb *PayloadBuilder) getPayloadFromExecutionClient(
	ctx context.Context,
	payloadID *engineprimitives.PayloadID,
	slot math.Slot,
) (engineprimitives.ExecutionPayload,
	engineprimitives.BlobsBundle, bool, error) {
	if payloadID == nil {
		return nil, nil, false, ErrNilPayloadID
	}

	envelope, err := pb.ee.GetPayload(
		ctx,
		&execution.GetPayloadRequest{
			PayloadID:   *payloadID,
			ForkVersion: pb.chainSpec.ActiveForkVersionForSlot(slot),
		},
	)
	if err != nil {
		return nil, nil, false, err
	} else if envelope == nil {
		return nil, nil, false, ErrNilPayloadEnvelope
	}

	overrideBuilder := envelope.ShouldOverrideBuilder()
	args := []any{
		"for_slot", slot,
		"override_builder", overrideBuilder,
	}

	payload := envelope.GetExecutionPayload()
	if payload != nil && !payload.IsNil() {
		args = append(args,
			"payload_block_hash", payload.GetBlockHash(),
			"parent_hash", payload.GetParentHash(),
		)
	}

	blobsBundle := envelope.GetBlobsBundle()
	if blobsBundle != nil {
		args = append(args, "num_blobs", len(blobsBundle.GetBlobs()))
	}

	pb.logger.Info("payload retrieved from local builder 🏗️ ", args...)
	return payload, blobsBundle, overrideBuilder, err
}
