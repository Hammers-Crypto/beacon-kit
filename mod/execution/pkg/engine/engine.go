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

package engine

import (
	"context"

	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/berachain/beacon-kit/mod/execution/pkg/client"
	"github.com/berachain/beacon-kit/mod/log"
	engineprimitives "github.com/berachain/beacon-kit/mod/primitives-engine"
	engineerrors "github.com/berachain/beacon-kit/mod/primitives-engine/pkg/errors"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	jsonrpc "github.com/berachain/beacon-kit/mod/primitives/pkg/net/json-rpc"
)

// Engine is Beacon-Kit's implementation of the `ExecutionEngine`
// from the Ethereum 2.0 Specification.
type Engine[
	ExecutionPayloadT ExecutionPayload,
	ExecutionPayloadDenebT engineprimitives.ExecutionPayload,
] struct {
	ec      *client.EngineClient[ExecutionPayloadDenebT]
	logger  log.Logger[any]
	metrics *engineMetrics
}

// New creates a new Engine.
func New[
	ExecutionPayloadT ExecutionPayload,
	ExecutionPayloadDenebT engineprimitives.ExecutionPayload,
](
	ec *client.EngineClient[ExecutionPayloadDenebT],
	logger log.Logger[any],
	ts TelemetrySink,
) *Engine[ExecutionPayloadT, ExecutionPayloadDenebT] {
	return &Engine[ExecutionPayloadT, ExecutionPayloadDenebT]{
		ec:      ec,
		logger:  logger,
		metrics: newEngineMetrics(ts, logger),
	}
}

// Start spawns any goroutines required by the service.
func (ee *Engine[
	ExecutionPayloadT, ExecutionPayloadDenebT,
]) Start(
	ctx context.Context,
) error {
	go func() {
		// TODO: handle better
		if err := ee.ec.Start(ctx); err != nil {
			panic(err)
		}
	}()
	return nil
}

// Status returns error if the service is not considered healthy.
func (ee *Engine[
	ExecutionPayloadT, ExecutionPayloadDenebT,
]) Status() error {
	return ee.ec.Status()
}

// GetPayload returns the payload and blobs bundle for the given slot.
func (ee *Engine[
	ExecutionPayloadT, ExecutionPayloadDenebT,
]) GetPayload(
	ctx context.Context,
	req *engineprimitives.GetPayloadRequest,
) (engineprimitives.BuiltExecutionPayloadEnv, error) {
	return ee.ec.GetPayload(
		ctx, req.PayloadID,
		req.ForkVersion,
	)
}

// NotifyForkchoiceUpdate notifies the execution client of a forkchoice update.
func (ee *Engine[
	ExecutionPayloadT, ExecutionPayloadDenebT,
]) NotifyForkchoiceUpdate(
	ctx context.Context,
	req *engineprimitives.ForkchoiceUpdateRequest,
) (*engineprimitives.PayloadID, *common.ExecutionHash, error) {
	ee.logger.Info("notifying forkchoice update",
		"head_eth1_hash", req.State.HeadBlockHash,
		"safe_eth1_hash", req.State.SafeBlockHash,
		"finalized_eth1_hash", req.State.FinalizedBlockHash,
		"has_attributes", req.PayloadAttributes != nil,
	)

	// Notify the execution engine of the forkchoice update.
	payloadID, latestValidHash, err := ee.ec.ForkchoiceUpdated(
		ctx,
		req.State,
		req.PayloadAttributes,
		req.ForkVersion,
	)
	switch {
	case errors.IsAny(
		err,
		engineerrors.ErrAcceptedPayloadStatus,
		engineerrors.ErrSyncingPayloadStatus,
	):
		ee.metrics.MarkForkchoiceUpdateAcceptedSyncing(req.State)
		return payloadID, nil, nil

	// If we get invalid payload status, we will need to find a valid
	// ancestor block and force a recovery.
	//
	// These two cases are semantically the same:
	// https://github.com/ethereum/execution-apis/issues/270
	case errors.IsAny(
		err,
		engineerrors.ErrInvalidPayloadStatus,
		engineerrors.ErrInvalidBlockHashPayloadStatus,
	):
		ee.metrics.MarkForkchoiceUpdateInvalid(req.State)
		req.State.HeadBlockHash = req.State.SafeBlockHash
		payloadID, latestValidHash, err = ee.NotifyForkchoiceUpdate(ctx, req)
		if err != nil {
			// We have to return the error here since this function
			// is recursive.
			return nil, nil, err
		}
		return payloadID, latestValidHash, ErrBadBlockProduced
	case jsonrpc.IsPreDefinedError(err):
		ee.metrics.MarkForkchoiceUpdateJSONRPCError(err)
		return nil, nil, errors.Join(err, engineerrors.ErrPreDefinedJSONRPC)
	case err != nil:
		ee.metrics.MarkForkchoiceUpdateUndefinedError(err)
		return nil, nil, err
	}

	return payloadID, latestValidHash, nil
}

// VerifyAndNotifyNewPayload verifies the new payload and notifies the
// execution client.
func (ee *Engine[
	ExecutionPayloadT, ExecutionPayloadDenebT,
]) VerifyAndNotifyNewPayload(
	ctx context.Context,
	req *engineprimitives.NewPayloadRequest[ExecutionPayloadT],
) error {
	// First we verify the block hash and versioned hashes are valid.
	//
	// TODO: is this required? Or will the EL handle this for us during
	// new payload?
	if err := req.HasValidVersionedAndBlockHashes(); err != nil {
		return err
	}

	// If the block already exists on our execution client
	// we can skip sending the payload to speed things up a bit.
	if req.SkipIfExists {
		header, err := ee.ec.HeaderByHash(
			ctx,
			req.ExecutionPayload.GetBlockHash(),
		)

		// If we find the header and there is no error, we can
		// skip any payload verification, since this block must've
		// been validated at some point in the past.
		if header != nil && err == nil {
			ee.logger.Info("skipping new payload, block already available",
				"block_hash", req.ExecutionPayload.GetBlockHash(),
			)
			return nil
		}
	}

	ee.logger.Info(
		"calling new payload",
		"payload-block-hash", req.ExecutionPayload.GetBlockHash(),
		"payload-parent-block-hash", req.ExecutionPayload.GetParentHash(),
		"is-optimistic", req.Optimistic,
	)

	// Otherwise we will send the payload to the execution client.
	lastValidHash, err := ee.ec.NewPayload(
		ctx,
		req.ExecutionPayload,
		req.VersionedHashes,
		req.ParentBeaconBlockRoot,
	)

	// We abstract away some of the complexity and categorize status codes
	// to make it easier to reason about.
	switch {
	// If we get accepted or syncing, we are going to optimistically
	// say that the block is valid, this is utilized during syncing
	// to allow the beacon-chain to continue processing blocks, while
	// its execution client is fetching things over it's p2p layer.
	case errors.IsAny(
		err,
		engineerrors.ErrAcceptedPayloadStatus,
		engineerrors.ErrSyncingPayloadStatus,
	):
		ee.metrics.MarkNewPayloadAcceptedSyncingPayloadStatus(
			req.ExecutionPayload.GetBlockHash(),
			req.Optimistic,
		)

	// These two cases are semantically the same:
	// https://github.com/ethereum/execution-apis/issues/270
	case errors.IsAny(
		err,
		engineerrors.ErrInvalidPayloadStatus,
		engineerrors.ErrInvalidBlockHashPayloadStatus,
	):
		ee.metrics.MarkNewPayloadInvalidPayloadStatus(
			req.ExecutionPayload.GetBlockHash(),
			req.Optimistic,
		)

		return ErrBadBlockProduced

	case jsonrpc.IsPreDefinedError(err):
		// Protect against possible nil value.
		if lastValidHash == nil {
			lastValidHash = &common.ExecutionHash{}
		}

		ee.metrics.MarkNewPayloadJSONRPCError(
			req.ExecutionPayload.GetBlockHash(),
			*lastValidHash,
			req.Optimistic,
			err,
		)

		err = errors.Join(err, engineerrors.ErrPreDefinedJSONRPC)
	case err != nil:
		ee.metrics.MarkNewPayloadUndefinedError(
			req.ExecutionPayload.GetBlockHash(),
			req.Optimistic,
			err,
		)
	}

	// Under the optimistic condition, we are fine ignoring the error. This
	// is mainly to allow us to safely call the execution client
	// during abci.FinalizeBlock. If we are in abci.FinalizeBlock and
	// we get an error here, we make the assumption that
	// abci.ProcessProposal
	// has deemed that the BeaconBlock containing the given ExecutionPayload
	// was marked as valid by an honest majority of validators, and we
	// don't want to halt the chain because of an error here.
	//
	// The pratical reason we want to handle this edge case
	// is to protect against an awkward shutdown condition in which an
	// execution client dies between the end of abci.ProcessProposal
	// and the beginning of abci.FinalizeBlock. Without handling this case
	// it would cause a failure of abci.FinalizeBlock and a
	// "CONSENSUS FAILURE!!!!" at the CometBFT layer.
	if req.Optimistic {
		return nil
	}
	return err
}
