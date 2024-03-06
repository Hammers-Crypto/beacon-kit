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

package preblock

import (
	"context"
	randaotypes "github.com/itsdevbear/bolaris/beacon/core/randao/types"

	"cosmossdk.io/log"
	cometabci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/itsdevbear/bolaris/beacon/blockchain"
	"github.com/itsdevbear/bolaris/beacon/core/state"
	beacontypes "github.com/itsdevbear/bolaris/beacon/core/types"
	"github.com/itsdevbear/bolaris/beacon/sync"
	"github.com/itsdevbear/bolaris/config"
	byteslib "github.com/itsdevbear/bolaris/lib/bytes"
	"github.com/itsdevbear/bolaris/primitives"
	abcitypes "github.com/itsdevbear/bolaris/runtime/abci/types"
)

type BeaconKeeper interface {
	BeaconState(ctx context.Context) state.BeaconState
}

// BeaconPreBlockHandler is responsible for aggregating oracle data from each
// validator and writing the oracle data into the store before any transactions
// are executed/finalized for a given block.
type BeaconPreBlockHandler struct {
	// cfg is the configuration for block proposals and finalization.
	cfg *config.ABCI

	// logger is the logger used by the handler.
	logger log.Logger

	// chainService is the service that is responsible for interacting with
	// the beacon chain.
	chainService *blockchain.Service

	// syncService is the service that is responsible for syncing the beacon
	// chain.
	syncService *sync.Service
	// nextHandler is the next pre-block handler in the chain. This is always
	// nesting of the next pre-block handler into this handler.
	nextHandler sdk.PreBlocker
}

// NewBeaconPreBlockHandler returns a new BeaconPreBlockHandler. The handler
// is responsible for writing oracle data included in vote extensions to state.
func NewBeaconPreBlockHandler(
	cfg *config.ABCI,
	logger log.Logger,
	chainService *blockchain.Service,
	syncService *sync.Service,
	nextHandler sdk.PreBlocker,
) *BeaconPreBlockHandler {
	return &BeaconPreBlockHandler{
		cfg:          cfg,
		logger:       logger,
		chainService: chainService,
		syncService:  syncService,
		nextHandler:  nextHandler,
	}
}

// PreBlocker is called by the base app before the block is finalized. It
// is responsible for aggregating oracle data from each validator and writing
// the oracle data to the store.
func (h *BeaconPreBlockHandler) PreBlocker() sdk.PreBlocker {
	return func(
		ctx sdk.Context, req *cometabci.RequestFinalizeBlock,
	) (*sdk.ResponsePreBlock, error) {
		cometBlockHash := byteslib.ToBytes32(req.Hash)

		// Extract the beacon block from the ABCI request.
		//
		// TODO: Block factory struct?
		// TODO: Use protobuf and .(type)?
		buoy, err := abcitypes.ReadOnlyBeaconBlockFromABCIRequest(
			req,
			h.cfg.BeaconBlockPosition,
			h.chainService.ActiveForkVersionForSlot(
				primitives.Slot(req.Height),
			),
		)
		if err != nil {
			h.logger.Error(
				"failed to extract beacon block from request",
				"error",
				err,
			)

			// If we fail to extract the beacon block from the request, we
			// create an empty beacon block to continue processing.
			// TODO: This is a temporary solution to avoid panics, we should
			// handle this better.
			if buoy, err = beacontypes.EmptyBeaconBlock(
				primitives.Slot(req.Height),
				h.chainService.BeaconState(ctx).GetParentBlockRoot(),
				h.chainService.ActiveForkVersionForSlot(
					primitives.Slot(req.Height),
				),
				randaotypes.Reveal{},
			); err != nil {
				return nil, err
			}
		}

		// Receive the beacon block to validate whether it is good and submit
		// any required newPayload and/or forkchoice updates. If we have
		// already ran this for the current block in ProcessProposal, this
		// call will exit early.
		if err = h.chainService.ReceiveBeaconBlock(
			ctx,
			cometBlockHash,
			buoy,
		); err != nil {
			h.logger.Warn(
				"failed to receive beacon block",
				"error",
				err,
			)
		}

		// Process the finalization of the beacon block.
		if err = h.chainService.FinalizeBeaconBlock(
			ctx, buoy, cometBlockHash,
		); err != nil {
			h.chainService.Logger().
				Error("failed to finalize beacon block", "error", err)
		}

		// Call the nested child handler.
		return h.callNextHandler(ctx, req)
	}
}

// callNextHandler calls the next pre-block handler in the chain.
func (h *BeaconPreBlockHandler) callNextHandler(
	ctx sdk.Context, req *cometabci.RequestFinalizeBlock,
) (*sdk.ResponsePreBlock, error) {
	// If there is no child handler, we are done, this preblocker
	// does not modify any consensus params so we return an empty
	// response.
	if h.nextHandler == nil {
		return &sdk.ResponsePreBlock{}, nil
	}

	return h.nextHandler(ctx, req)
}
