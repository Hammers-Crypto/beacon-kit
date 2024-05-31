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

package client

import (
	"context"
	"time"

	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	engineerrors "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/errors"
	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/berachain/beacon-kit/mod/execution/pkg/client/ethclient"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/version"
)

// NewPayload calls the engine_newPayloadVX method via JSON-RPC.
func (s *EngineClient[ExecutionPayloadDenebT]) NewPayload(
	ctx context.Context,
	payload ExecutionPayload,
	versionedHashes []common.ExecutionHash,
	parentBeaconBlockRoot *primitives.Root,
) (*common.ExecutionHash, error) {
	startTime := time.Now()
	defer s.metrics.measureNewPayloadDuration(startTime)
	dctx, cancel := context.WithTimeoutCause(
		ctx, s.cfg.RPCTimeout, engineerrors.ErrEngineAPITimeout,
	)
	defer cancel()

	// Call the appropriate RPC method based on the payload version.
	result, err := s.callNewPayloadRPC(
		dctx,
		payload,
		versionedHashes,
		parentBeaconBlockRoot,
	)
	if err != nil {
		if errors.Is(err, engineerrors.ErrEngineAPITimeout) {
			s.metrics.incrementNewPayloadTimeout()
		}
		return nil, err
	} else if result == nil {
		return nil, engineerrors.ErrNilPayloadStatus
	}

	// This case is only true when the payload is invalid, so
	// `processPayloadStatusResult` below will return an error.
	if validationErr := result.ValidationError; validationErr != nil {
		s.logger.Error(
			"Got a validation error in newPayload",
			"err",
			errors.New(*validationErr),
		)
	}

	return processPayloadStatusResult(result)
}

// callNewPayloadRPC calls the engine_newPayloadVX method via JSON-RPC.
func (s *EngineClient[ExecutionPayloadDenebT]) callNewPayloadRPC(
	ctx context.Context,
	payload ExecutionPayload,
	versionedHashes []common.ExecutionHash,
	parentBeaconBlockRoot *primitives.Root,
) (*engineprimitives.PayloadStatusV1, error) {
	switch payload.Version() {
	case version.Deneb:
		return s.NewPayloadV3(
			ctx,
			payload,
			versionedHashes,
			parentBeaconBlockRoot,
		)
	case version.Electra:
		return nil, errors.New("TODO: implement Electra payload")
	default:
		return nil, engineerrors.ErrInvalidPayloadType
	}
}

// ForkchoiceUpdated calls the engine_forkchoiceUpdatedV1 method via JSON-RPC.
func (s *EngineClient[ExecutionPayloadDenebT]) ForkchoiceUpdated(
	ctx context.Context,
	state *engineprimitives.ForkchoiceStateV1,
	attrs engineprimitives.PayloadAttributer,
	forkVersion uint32,
) (*engineprimitives.PayloadID, *common.ExecutionHash, error) {
	startTime := time.Now()
	defer s.metrics.measureForkchoiceUpdateDuration(startTime)
	dctx, cancel := context.WithTimeoutCause(
		ctx, s.cfg.RPCTimeout, engineerrors.ErrEngineAPITimeout,
	)
	defer cancel()

	// If the suggested fee recipient is not set, log a warning.
	if attrs != nil && !attrs.IsNil() &&
		attrs.GetSuggestedFeeRecipient() == (common.ZeroAddress) {
		s.logger.Warn(
			"suggested fee recipient is not configured 🔆",
			"fee-recipent", common.DisplayBytes(
				common.ZeroAddress[:]).TerminalString(),
		)
	}

	result, err := s.callUpdatedForkchoiceRPC(dctx, state, attrs, forkVersion)

	if err != nil {
		if errors.Is(err, engineerrors.ErrEngineAPITimeout) {
			s.metrics.incrementForkchoiceUpdateTimeout()
		}
		return nil, nil, s.handleRPCError(err)
	} else if result == nil {
		return nil, nil, engineerrors.ErrNilForkchoiceResponse
	}

	latestValidHash, err := processPayloadStatusResult((&result.PayloadStatus))
	if err != nil {
		return nil, latestValidHash, err
	}
	return result.PayloadID, latestValidHash, nil
}

// updateForkChoiceByVersion calls the engine_forkchoiceUpdatedVX method via
// JSON-RPC.
func (s *EngineClient[ExecutionPayloadDenebT]) callUpdatedForkchoiceRPC(
	ctx context.Context,
	state *engineprimitives.ForkchoiceStateV1,
	attrs engineprimitives.PayloadAttributer,
	forkVersion uint32,
) (*engineprimitives.ForkchoiceResponseV1, error) {
	switch forkVersion {
	case version.Deneb:
		return s.ForkchoiceUpdatedV3(ctx, state, attrs)
	case version.Electra:
		return nil, errors.New("TODO: implement Electra forkchoice")
	default:
		return nil, engineerrors.ErrInvalidPayloadAttributes
	}
}

// GetPayload calls the engine_getPayloadVX method via JSON-RPC. It returns
// the execution data as well as the blobs bundle.
func (s *EngineClient[ExecutionPayloadDenebT]) GetPayload(
	ctx context.Context,
	payloadID engineprimitives.PayloadID,
	forkVersion uint32,
) (engineprimitives.BuiltExecutionPayloadEnv, error) {
	startTime := time.Now()
	defer s.metrics.measureGetPayloadDuration(startTime)
	dctx, cancel := context.WithTimeoutCause(
		ctx, s.cfg.RPCTimeout, engineerrors.ErrEngineAPITimeout,
	)
	defer cancel()

	// Determine what version we want to call.
	var fn func(
		context.Context, engineprimitives.PayloadID,
	) (engineprimitives.BuiltExecutionPayloadEnv, error)
	switch forkVersion {
	case version.Deneb:
		fn = s.GetPayloadV3
	case version.Electra:
		return nil, errors.New("TODO: implement Electra getPayload")
	default:
		return nil, engineerrors.ErrInvalidGetPayloadVersion
	}

	// Call and check for errors.
	result, err := fn(dctx, payloadID)
	switch {
	case err != nil:
		if errors.Is(err, engineerrors.ErrEngineAPITimeout) {
			s.metrics.incrementNewPayloadTimeout()
		}
		return result, s.handleRPCError(err)
	case result == nil:
		return result, engineerrors.ErrNilExecutionPayloadEnvelope
	case result.GetExecutionPayload() == nil:
		return result, engineerrors.ErrNilExecutionPayload
	case result.GetBlobsBundle() == nil && forkVersion >= version.Deneb:
		return result, engineerrors.ErrNilBlobsBundle
	}

	return result, nil
}

// ExchangeCapabilities calls the engine_exchangeCapabilities method via
// JSON-RPC.
func (s *EngineClient[ExecutionPayloadDenebT]) ExchangeCapabilities(
	ctx context.Context,
) ([]string, error) {
	result, err := s.Eth1Client.ExchangeCapabilities(
		ctx, ethclient.BeaconKitSupportedCapabilities(),
	)
	if err != nil {
		s.statusErrMu.Lock()
		defer s.statusErrMu.Unlock()
		//#nosec:G703 wtf is even this problem here.
		s.statusErr = s.handleRPCError(err)
		return nil, s.statusErr
	}

	// Capture and log the capabilities that the execution client has.
	for _, capability := range result {
		s.logger.Info("exchanged capability", "capability", capability)
		s.capabilities[capability] = struct{}{}
	}

	// Log the capabilities that the execution client does not have.
	for _, capability := range ethclient.BeaconKitSupportedCapabilities() {
		if _, exists := s.capabilities[capability]; !exists {
			s.logger.Warn(
				"your execution client may require an update 🚸",
				"unsupported_capability", capability,
			)
		}
	}

	s.statusErr = nil
	return result, nil
}
