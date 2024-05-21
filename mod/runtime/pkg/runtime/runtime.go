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

package runtime

import (
	"context"

	"github.com/berachain/beacon-kit/mod/beacon/blockchain"
	"github.com/berachain/beacon-kit/mod/beacon/validator"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/runtime/pkg/abci"
	"github.com/berachain/beacon-kit/mod/runtime/pkg/service"
	"github.com/berachain/beacon-kit/mod/state-transition/pkg/core/state"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeaconKitRuntime is a struct that holds the
// service registry.
type BeaconKitRuntime[
	AvailabilityStoreT AvailabilityStore[types.BeaconBlockBody, BlobSidecarsT],
	BeaconBlockBodyT types.BeaconBlockBody,
	BeaconStateT state.BeaconState,
	BlobSidecarsT BlobSidecars,
	DepositStoreT DepositStore,
	StorageBackendT StorageBackend[
		AvailabilityStoreT,
		BeaconBlockBodyT,
		BeaconStateT,
		BlobSidecarsT,
		DepositStoreT,
	],
] struct {
	logger         log.Logger[any]
	services       *service.Registry
	storageBackend StorageBackendT
	chainSpec      primitives.ChainSpec
}

// NewBeaconKitRuntime creates a new BeaconKitRuntime
// and applies the provided options.
func NewBeaconKitRuntime[
	AvailabilityStoreT AvailabilityStore[types.BeaconBlockBody, BlobSidecarsT],
	BeaconBlockBodyT types.BeaconBlockBody,
	BeaconStateT state.BeaconState,
	BlobSidecarsT BlobSidecars,
	DepositStoreT DepositStore,
	StorageBackendT StorageBackend[
		AvailabilityStoreT,
		BeaconBlockBodyT,
		BeaconStateT,
		BlobSidecarsT,
		DepositStoreT,
	],
](
	chainSpec primitives.ChainSpec,
	logger log.Logger[any],
	services *service.Registry,
	storageBackend StorageBackendT,
) (*BeaconKitRuntime[
	AvailabilityStoreT,
	BeaconBlockBodyT,
	BeaconStateT,
	BlobSidecarsT,
	DepositStoreT,
	StorageBackendT,
], error) {
	return &BeaconKitRuntime[
		AvailabilityStoreT,
		BeaconBlockBodyT,
		BeaconStateT,
		BlobSidecarsT,
		DepositStoreT,
		StorageBackendT,
	]{
		// engineClient:   engineClient,
		chainSpec:      chainSpec,
		logger:         logger,
		services:       services,
		storageBackend: storageBackend,
	}, nil
}

// StartServices starts the services.
func (r *BeaconKitRuntime[
	AvailabilityStoreT,
	BeaconStateT,
	BlobSidecarsT,
	DepositStoreT,
	BeaconBlockBodyT,
	StorageBackendT,
]) StartServices(
	ctx context.Context,
) error {
	r.services.StartAll(ctx)
	return nil
}

// BuildABCIComponents returns the ABCI components for the beacon runtime.
func (r *BeaconKitRuntime[
	AvailabilityStoreT,
	BeaconBlockBodyT,
	BeaconStateT,
	BlobSidecarsT,
	DepositStoreT,
	StorageBackendT,
]) BuildABCIComponents() (
	sdk.PrepareProposalHandler, sdk.ProcessProposalHandler,
	sdk.PreBlocker,
) {
	var (
		chainService *blockchain.Service[
			AvailabilityStoreT,
			state.BeaconState,
			BlobSidecarsT,
			DepositStoreT,
		]
		builderService *validator.Service[BeaconStateT, BlobSidecarsT]
	)
	if err := r.services.FetchService(&chainService); err != nil {
		panic(err)
	}

	if err := r.services.FetchService(&builderService); err != nil {
		panic(err)
	}

	if chainService == nil || builderService == nil {
		panic("missing services")
	}

	handler := abci.NewHandler[BlobSidecarsT](
		r.chainSpec,
		builderService,
		chainService,
	)

	return handler.PrepareProposalHandler,
		handler.ProcessProposalHandler,
		handler.FinalizeBlock
}
