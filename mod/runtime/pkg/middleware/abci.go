// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	appmodulev2 "cosmossdk.io/core/appmodule/v2"
	asynctypes "github.com/berachain/beacon-kit/mod/async/pkg/types"
	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/events"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/runtime/pkg/encoding"
	cmtabci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcegraph/conc/iter"
	"golang.org/x/sync/errgroup"
)

/* -------------------------------------------------------------------------- */
/*                                 InitGenesis                                */
/* -------------------------------------------------------------------------- */

// InitGenesis is called by the base app to initialize the state of the.
func (h *ABCIMiddleware[
	_, _, _, _, _, _, GenesisT,
]) InitGenesis(
	ctx context.Context,
	bz []byte,
) ([]appmodulev2.ValidatorUpdate, error) {
	data := new(GenesisT)
	if err := json.Unmarshal(bz, data); err != nil {
		return nil, err
	}
	updates, err := h.chainService.ProcessGenesisData(
		ctx,
		*data,
	)
	if err != nil {
		return nil, err
	}

	// Convert updates into the Cosmos SDK format.
	return iter.MapErr(updates, convertValidatorUpdate)
}

/* -------------------------------------------------------------------------- */
/*                               PrepareProposal                              */
/* -------------------------------------------------------------------------- */

// prepareProposal is the internal handler for preparing proposals.
func (h *ABCIMiddleware[
	_, _, _, _, _, _, _,
]) PrepareProposal(
	ctx sdk.Context,
	req *cmtabci.PrepareProposalRequest,
) (*cmtabci.PrepareProposalResponse, error) {
	var (
		wg                          sync.WaitGroup
		startTime                   = time.Now()
		beaconBlockErr, sidecarsErr error
		beaconBlockBz, sidecarsBz   []byte
	)
	defer h.metrics.measurePrepareProposalDuration(startTime)

	// Send a request to the validator service to give us a beacon block
	// and blob sidecards to pass to ABCI.
	if err := h.slotBroker.Publish(asynctypes.NewEvent(
		ctx, events.NewSlot, math.Slot(req.Height),
	)); err != nil {
		return nil, err
	}

	// Using a wait group instead of an errgroup to ensure we drain
	// the associated channels for the beacon block and sidecars.
	//nolint:mnd // bet.
	wg.Add(2)
	go func() {
		defer wg.Done()
		beaconBlockBz, beaconBlockErr = h.waitforBeaconBlk(ctx)
	}()

	go func() {
		defer wg.Done()
		sidecarsBz, sidecarsErr = h.waitForSidecars(ctx)
	}()

	wg.Wait()
	if beaconBlockErr != nil {
		return nil, beaconBlockErr
	} else if sidecarsErr != nil {
		return nil, sidecarsErr
	}

	return &cmtabci.PrepareProposalResponse{
		Txs: [][]byte{beaconBlockBz, sidecarsBz},
	}, nil
}

// waitForSidecars waits for the sidecars to be built and returns them.
func (h *ABCIMiddleware[
	_, _, _, _, _, _, _,
]) waitForSidecars(ctx context.Context) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case msg := <-h.sidecarsCh:
		if msg.Error() != nil {
			return nil, msg.Error()
		}
		return h.blobGossiper.Publish(ctx, msg.Data())
	}
}

// waitforBeaconBlk waits for the beacon block to be built and returns it.
func (h *ABCIMiddleware[
	_, _, _, _, _, _, _,
]) waitforBeaconBlk(ctx context.Context) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case beaconBlock := <-h.blkCh:
		if beaconBlock.Error() != nil {
			return nil, beaconBlock.Error()
		}
		return h.beaconBlockGossiper.Publish(
			ctx,
			beaconBlock.Data(),
		)
	}
}

/* -------------------------------------------------------------------------- */
/*                               ProcessProposal                              */
/* -------------------------------------------------------------------------- */

// ProcessProposal processes the proposal for the ABCI middleware.
// It handles both the beacon block and blob sidecars concurrently.
func (h *ABCIMiddleware[
	_, BeaconBlockT, _, BlobSidecarsT, _, _, _,
]) ProcessProposal(
	ctx sdk.Context,
	req *cmtabci.ProcessProposalRequest,
) (*cmtabci.ProcessProposalResponse, error) {
	var (
		blk      BeaconBlockT
		sidecars BlobSidecarsT
		err      error
		g, _     = errgroup.WithContext(ctx)
	)

	startTime := time.Now()
	defer h.metrics.measureProcessProposalDuration(startTime)

	if blk, err = h.beaconBlockGossiper.Request(ctx, req); err != nil {
		return h.createResponse(errors.WrapNonFatal(err))
	}

	if sidecars, err = h.blobGossiper.Request(ctx, req); err != nil {
		return h.createResponse(errors.WrapNonFatal(err))
	}

	g.Go(func() error {
		return h.processBeaconBlock(ctx, blk)
	})

	g.Go(func() error {
		return h.processBlobSidecars(ctx, sidecars)
	})

	return h.createResponse(g.Wait())
}

// processBeaconBlock handles the processing of the beacon block.
// It requests the block, publishes a received event, and waits for
// verification.
func (h *ABCIMiddleware[
	_, BeaconBlockT, _, BlobSidecarsT, _, _, _,
]) processBeaconBlock(
	ctx context.Context,
	blk BeaconBlockT,
) error {
	// Publish the received event.
	if err := h.blkBroker.Publish(
		asynctypes.NewEvent(ctx, events.BeaconBlockReceived, blk, nil),
	); err != nil {
		return err
	}

	// Wait for a response.
	select {
	case <-ctx.Done():
		return ctx.Err()
	case msg := <-h.blkCh:
		if msg.Type() != events.BeaconBlockVerified {
			return errors.Wrapf(
				ErrUnexpectedEvent, "unexpected event type: %s", msg.Type(),
			)
		}
		return msg.Error()
	}
}

// processBlobSidecars handles the processing of blob sidecars.
// It requests the sidecars, publishes a received event, and waits for
// processing.
func (h *ABCIMiddleware[
	_, BeaconBlockT, _, BlobSidecarsT, _, _, _,
]) processBlobSidecars(
	ctx context.Context,
	sidecars BlobSidecarsT,
) error {
	// Publish the received event.
	if err := h.sidecarsBroker.Publish(
		asynctypes.NewEvent(ctx, events.BlobSidecarsReceived, sidecars, nil),
	); err != nil {
		return err
	}

	// Wait for a response.
	select {
	case <-ctx.Done():
		return ctx.Err()
	case msg := <-h.sidecarsCh:
		if msg.Type() != events.BlobSidecarsProcessed {
			return errors.Wrapf(
				ErrUnexpectedEvent, "unexpected event type: %s", msg.Type(),
			)
		}
		return msg.Error()
	}
}

// createResponse generates the appropriate ProcessProposalResponse based on the
// error.
func (h *ABCIMiddleware[
	_, BeaconBlockT, _, BlobSidecarsT, _, _, _,
]) createResponse(err error) (*cmtabci.ProcessProposalResponse, error) {
	status := cmtabci.PROCESS_PROPOSAL_STATUS_REJECT
	if !errors.IsFatal(err) {
		status = cmtabci.PROCESS_PROPOSAL_STATUS_ACCEPT
		err = nil
	}
	return &cmtabci.ProcessProposalResponse{Status: status}, err
}

/* -------------------------------------------------------------------------- */
/*                                FinalizeBlock                               */
/* -------------------------------------------------------------------------- */

// PreBlock is called by the base app before the block is finalized. It
// is responsible for aggregating oracle data from each validator and writing
// the oracle data to the store.
func (h *ABCIMiddleware[
	_, _, _, _, _, _, _,
]) PreBlock(
	_ sdk.Context, req *cmtabci.FinalizeBlockRequest,
) error {
	h.req = req
	return nil
}

// EndBlock returns the validator set updates from the beacon state.
func (h *ABCIMiddleware[
	_, BeaconBlockT, _, BlobSidecarsT, _, _, _,
]) EndBlock(
	ctx context.Context,
) ([]appmodulev2.ValidatorUpdate, error) {
	blk, blobs, err := encoding.
		ExtractBlobsAndBlockFromRequest[BeaconBlockT, BlobSidecarsT](
		h.req,
		BeaconBlockTxIndex,
		BlobSidecarsTxIndex,
		h.chainSpec.ActiveForkVersionForSlot(
			math.Slot(h.req.Height),
		))
	if err != nil {
		// If we don't have a block, we can't do anything.
		//nolint:nilerr // by design.
		return nil, nil
	}

	// Send the sidecars to the sidecars feed, we know at this point
	// That the blobs have been successfully verified in process proposal.
	if err = h.sidecarsBroker.Publish(asynctypes.NewEvent(
		ctx, events.BlobSidecarsVerified, blobs,
	)); err != nil {
		return nil, err
	}

	// Wait for a response from the da service, with the current codepaths
	// we can't parallelize retrieving the DA service response and the
	// validator updates, since we need to check for IsDataAvailable in
	// `ProcessBeaconBlock`, we should improve this though.
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case sidecars := <-h.sidecarsCh:
		if sidecars.Type() != events.BlobSidecarsProcessed {
			return nil, fmt.Errorf(
				"unexpected event type: %s", sidecars.Type())
		}
		if sidecars.Error() != nil {
			return nil, sidecars.Error()
		}
	}

	// TODO: Move to Async.
	valUpdates, err := h.chainService.ProcessBeaconBlock(
		ctx, blk,
	)
	if err != nil {
		return nil, err
	}

	return iter.MapErr(
		valUpdates.RemoveDuplicates().Sort(), convertValidatorUpdate,
	)
}
