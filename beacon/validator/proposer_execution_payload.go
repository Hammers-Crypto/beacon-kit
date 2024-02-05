// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
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

package validator

import (
	"context"
	"fmt"
	"time"

	"github.com/itsdevbear/bolaris/beacon/core"
	"github.com/itsdevbear/bolaris/types/consensus/v1/interfaces"
	"github.com/itsdevbear/bolaris/types/state"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	enginev1 "github.com/prysmaticlabs/prysm/v4/proto/engine/v1"
	"github.com/prysmaticlabs/prysm/v4/runtime/version"
)

func (s *Service) getLocalPayload(
	ctx context.Context,
	blk interfaces.ReadOnlyBeaconKitBlock, st state.BeaconState,
) (interfaces.ExecutionData, bool, error) {
	if blk.Version() < version.Bellatrix {
		return nil, false, nil
	}

	slot := blk.GetSlot()
	// vIdx := blk.ProposerIndex()
	// headRoot := blk.ParentRoot()
	// logFields := logrus.Fields{
	// 	"validatorIndex": vIdx,
	// 	"slot":           slot,
	// 	"headRoot":       fmt.Sprintf("%#x", headRoot),
	// }
	// p, err := blk.Execution()
	// if err != nil {
	// 	return nil, false, err
	// }

	// val, tracked := s.TrackedValidatorsCache.Validator(vIdx)
	// if !tracked {
	// 	logrus.WithFields(logFields).Warn("could not find tracked proposer index")
	// }

	// Otherwise we did not have a payload in the cache and we must build a new payload.

	// log.WithFields(logFields).Debug("payload ID cache miss")vs

	parentEth1Hash, err := s.getParentBlockHash(ctx)
	if err != nil {
		return nil, false, err
	}

	// If we have a payload ID in the cache, we can return the payload from the cache.
	payloadID, ok := s.payloadCache.PayloadID(slot, [32]byte(parentEth1Hash))
	if ok && (payloadID != primitives.PayloadID{}) {
		// Payload ID is cache hit. Return the cached payload ID.
		var pid primitives.PayloadID
		copy(pid[:], payloadID[:])
		// payloadIDCacheHit.Inc()
		var payload interfaces.ExecutionData
		var overrideBuilder bool
		payload, _, overrideBuilder, err = s.en.GetPayload(ctx, pid, slot)
		switch {
		case err == nil:
			// bundleCache.add(slot, bundle)
			// warnIfFeeRecipientDiffers(payload, val.FeeRecipient)
			return payload, overrideBuilder, nil
		case errors.Is(err, context.DeadlineExceeded):
		default:
			return nil, false, errors.Wrap(err, "could not get cached payload from execution client")
		}
	}

	// payloadIDCacheMiss.Inc()

	// random, err := helpers.RandaoMix(st, time.CurrentEpoch(st))
	// if err != nil {
	// 	return nil, false, err
	// }

	justifiedBlockHash := s.BeaconState(ctx).GetSafeEth1BlockHash()
	finalizedBlockHash := s.BeaconState(ctx).GetFinalizedEth1BlockHash()

	f := &enginev1.ForkchoiceState{
		HeadBlockHash:      parentEth1Hash,
		SafeBlockHash:      justifiedBlockHash[:],
		FinalizedBlockHash: finalizedBlockHash[:],
	}

	t := time.Now()              // todo: the proper mathematics for time must be done.
	random := make([]byte, 32)   //nolint:gomnd // todo: randao
	headRoot := make([]byte, 32) //nolint:gomnd // todo: cancaun
	attr := core.BuildPayloadAttributes(
		s.BeaconCfg(),
		st,
		s.Logger(),
		random,
		headRoot,
		//#nosec:G701 // won't overflow.
		uint64(t.Unix()),
	)

	var payloadIDBytes *enginev1.PayloadIDBytes
	payloadIDBytes, _, err = s.en.ForkchoiceUpdated(ctx, f, attr)
	if err != nil {
		return nil, false, errors.Wrap(err, "could not prepare payload")
	}
	if payloadIDBytes == nil {
		return nil, false, fmt.Errorf("nil payload with block hash: %#x", parentEth1Hash)
	}

	payload, _, overrideBuilder, err := s.en.GetPayload(ctx, *payloadIDBytes, slot)
	if err != nil {
		return nil, false, err
	}
	// bundleCache.add(slot, bundle)
	// warnIfFeeRecipientDiffers(payload, val.FeeRecipient)
	localValueGwei, err := payload.ValueInGwei()
	if err == nil {
		s.Logger().Debug("received execution payload from local engine", "value", localValueGwei)
	}
	return payload, overrideBuilder, nil
}

// getParentBlockHash retrieves the parent block hash for the given slot.
//
//nolint:unparam // todo: review this later.
func (s *Service) getParentBlockHash(ctx context.Context) ([]byte, error) {
	// The first slot should be proposed with the genesis block as parent.
	st := s.BeaconState(ctx)
	if st.Slot() == 1 {
		return st.GenesisEth1Hash().Bytes(), nil
	}

	// We always want the parent block to be the last finalized block.
	return st.GetFinalizedEth1BlockHash().Bytes(), nil
}
