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
	"errors"
	"math/big"
	"time"

	"cosmossdk.io/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/itsdevbear/bolaris/config"
	eth "github.com/itsdevbear/bolaris/engine/ethclient"
	enginetypes "github.com/itsdevbear/bolaris/engine/types"
	enginev1 "github.com/itsdevbear/bolaris/engine/types/v1"
	"github.com/itsdevbear/bolaris/types/consensus/primitives"
	"github.com/itsdevbear/bolaris/types/consensus/version"
)

// Caller is implemented by engineClient.
var _ Caller = (*engineClient)(nil)

// engineClient is a struct that holds a pointer to an Eth1Client.
type engineClient struct {
	*eth.Eth1Client

	capabilities  map[string]struct{}
	engineTimeout time.Duration
	beaconCfg     *config.Beacon
	logger        log.Logger
}

// NewClient creates a new engine client engineClient.
// It takes an Eth1Client as an argument and returns a pointer to an
// engineClient.
func NewClient(opts ...Option) Caller {
	ec := &engineClient{
		capabilities: make(map[string]struct{}),
	}
	for _, opt := range opts {
		if err := opt(ec); err != nil {
			panic(err)
		}
	}

	return ec
}

// Start starts the engine client.
func (s *engineClient) Start(ctx context.Context) {
	s.Eth1Client.Start(ctx)
	if _, err := s.ExchangeCapabilities(ctx); err != nil {
		s.logger.Error("failed to exchange capabilities", "err", err)
	}
}

// NewPayload calls the engine_newPayloadVX method via JSON-RPC.
func (s *engineClient) NewPayload(
	ctx context.Context, payload enginetypes.ExecutionPayload,
	versionedHashes []common.Hash, parentBlockRoot *common.Hash,
) ([]byte, error) {
	dctx, cancel := context.WithTimeout(ctx, s.engineTimeout)
	defer cancel()

	// Call the appropriate RPC method based on the payload version.
	result, err := s.callNewPayloadRPC(
		dctx,
		payload,
		versionedHashes,
		parentBlockRoot,
	)
	if err != nil {
		return nil, err
	}

	// This case is only true when the payload is invalid, so
	// `processPayloadStatusResult` below will return an error.
	if validationErr := result.GetValidationError(); validationErr != "" {
		s.logger.Error(
			"Got a validation error in newPayload",
			"err",
			errors.New(validationErr),
		)
	}

	return processPayloadStatusResult(result)
}

// callNewPayloadRPC calls the engine_newPayloadVX method via JSON-RPC.
func (s *engineClient) callNewPayloadRPC(
	ctx context.Context, payload enginetypes.ExecutionPayload,
	versionedHashes []common.Hash, parentBlockRoot *common.Hash,
) (*enginev1.PayloadStatus, error) {
	switch payloadPb := payload.ToProto().(type) {
	case *enginev1.ExecutionPayloadDeneb:
		return s.NewPayloadV3(ctx, payloadPb, versionedHashes, parentBlockRoot)
	default:
		return nil, ErrInvalidPayloadType
	}
}

// ForkchoiceUpdated calls the engine_forkchoiceUpdatedV1 method via JSON-RPC.
func (s *engineClient) ForkchoiceUpdated(
	ctx context.Context,
	state *enginev1.ForkchoiceState,
	attrs enginetypes.PayloadAttributer,
) (*enginev1.PayloadIDBytes, []byte, error) {
	dctx, cancel := context.WithTimeout(ctx, s.engineTimeout)
	defer cancel()

	if attrs == nil {
		return nil, nil, ErrNilAttributesPassedToClient
	}

	result, err := s.callUpdatedForkchoiceRPC(dctx, state, attrs)
	if err != nil {
		return nil, nil, err
	}

	lastestValidHash, err := processPayloadStatusResult(result.Status)
	if err != nil {
		return nil, lastestValidHash, err
	}
	return result.PayloadID, lastestValidHash, nil
}

// updateForkChoiceByVersion calls the engine_forkchoiceUpdatedVX method via
// JSON-RPC.
func (s *engineClient) callUpdatedForkchoiceRPC(
	ctx context.Context,
	state *enginev1.ForkchoiceState,
	attrs enginetypes.PayloadAttributer,
) (*eth.ForkchoiceUpdatedResponse, error) {
	switch v := attrs.ToProto().(type) {
	case *enginev1.PayloadAttributesV3:
		return s.ForkchoiceUpdatedV3(ctx, state, v)
	default:
		return nil, ErrInvalidPayloadAttributeVersion
	}
}

// GetPayload calls the engine_getPayloadVX method via JSON-RPC. It returns
// the execution data as well as the blobs bundle.
func (s *engineClient) GetPayload(
	ctx context.Context, payloadID primitives.PayloadID, slot primitives.Slot,
) (enginetypes.ExecutionPayload, *enginev1.BlobsBundle, bool, error) {
	dctx, cancel := context.WithTimeout(ctx, s.engineTimeout)
	defer cancel()

	var fn func(
		context.Context, enginev1.PayloadIDBytes,
	) (*enginev1.ExecutionPayloadContainer, error)
	switch s.beaconCfg.ActiveForkVersion(primitives.Epoch(slot)) {
	case version.Deneb:
		fn = s.GetPayloadV3
	default:
		return nil, nil, false, ErrInvalidGetPayloadVersion
	}

	result, err := fn(dctx, enginev1.PayloadIDBytes(payloadID))
	if err != nil {
		return nil, nil, false, err
	}

	return result, result.GetBlobsBundle(), result.GetShouldOverrideBuilder(), nil
}

// GetLogs retrieves the logs from the Ethereum execution node.
// It calls the eth_getLogs method via JSON-RPC.
func (s *engineClient) GetLogs(
	ctx context.Context,
	fromBlock, toBlock uint64,
	addresses []common.Address,
) ([]coretypes.Log, error) {
	// Create a filter query for the block, to acquire all logs from contracts
	// that we care about.
	query := ethereum.FilterQuery{
		Addresses: addresses,
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
	}

	// Gather all the logs according to the query.
	return s.FilterLogs(ctx, query)
}

// ExchangeCapabilities calls the engine_exchangeCapabilities method via
// JSON-RPC.
func (s *engineClient) ExchangeCapabilities(
	ctx context.Context,
) ([]string, error) {
	result, err := s.Eth1Client.ExchangeCapabilities(
		ctx, eth.BeaconKitSupportedCapabilities(),
	)
	if err != nil {
		return nil, err
	}

	// Capture and log the capabilities that the execution client has.
	for _, capability := range result {
		s.logger.Info("exchanged capability", "capability", capability)
		s.capabilities[capability] = struct{}{}
	}

	// Log the capabilities that the execution client does not have.
	for _, capability := range eth.BeaconKitSupportedCapabilities() {
		if _, exists := s.capabilities[capability]; !exists {
			s.logger.Warn(
				"your execution client may require an update 🚸",
				"unsupported_capability", capability,
			)
		}
	}

	return result, nil
}
