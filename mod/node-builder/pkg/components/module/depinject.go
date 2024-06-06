// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
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

package beacon

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"cosmossdk.io/depinject/appconfig"
	"github.com/berachain/beacon-kit/mod/beacon/blockchain"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/events"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/da/pkg/kzg"
	dastore "github.com/berachain/beacon-kit/mod/da/pkg/store"
	datypes "github.com/berachain/beacon-kit/mod/da/pkg/types"
	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	engineclient "github.com/berachain/beacon-kit/mod/execution/pkg/client"
	"github.com/berachain/beacon-kit/mod/execution/pkg/deposit"
	execution "github.com/berachain/beacon-kit/mod/execution/pkg/engine"
	"github.com/berachain/beacon-kit/mod/node-builder/pkg/components"
	"github.com/berachain/beacon-kit/mod/node-builder/pkg/components/metrics"
	modulev1alpha1 "github.com/berachain/beacon-kit/mod/node-builder/pkg/components/module/api/module/v1alpha1"
	"github.com/berachain/beacon-kit/mod/node-builder/pkg/components/storage"
	"github.com/berachain/beacon-kit/mod/node-builder/pkg/config"
	payloadbuilder "github.com/berachain/beacon-kit/mod/payload/pkg/builder"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/transition"
	"github.com/berachain/beacon-kit/mod/runtime/pkg/service"
	"github.com/berachain/beacon-kit/mod/state-transition/pkg/core"
	"github.com/berachain/beacon-kit/mod/storage/pkg/beacondb"
	"github.com/berachain/beacon-kit/mod/storage/pkg/beacondb/encoding"
	depositdb "github.com/berachain/beacon-kit/mod/storage/pkg/deposit"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/ethereum/go-ethereum/event"
)

// TODO: we don't allow generics here? Why? Is it fixable?
//
//nolint:gochecknoinits // required by sdk.
func init() {
	appconfig.RegisterModule(&modulev1alpha1.Module{},
		appconfig.Provide(
			ProvideModule,
		),
	)
}

// DepInjectInput is the input for the dep inject framework.
type DepInjectInput struct {
	depinject.In

	// Cosmos components
	AppOpts     servertypes.AppOptions
	Environment appmodule.Environment

	// BeaconKit components
	AvailabilityStore     *dastore.Store[types.BeaconBlockBody]
	BeaconConfig          *config.Config
	ChainSpec             primitives.ChainSpec
	DepositStore          *depositdb.KVStore[*types.Deposit]
	EngineClient          *engineclient.EngineClient[*types.ExecutionPayload]
	ExecutionEngine       *execution.Engine[*types.ExecutionPayload]
	BeaconDepositContract *deposit.WrappedBeaconDepositContract[
		*types.Deposit,
		types.WithdrawalCredentials,
	]
	BlobProofVerifier kzg.BlobProofVerifier
	Signer            crypto.BLSSigner
	TelemetrySink     *metrics.TelemetrySink
	LocalBuilder      *payloadbuilder.PayloadBuilder[
		components.BeaconState,
		*types.ExecutionPayload,
		*types.ExecutionPayloadHeader,
	]
	BlockFeed      *event.FeedOf[events.Block[*types.BeaconBlock]]
	StateProcessor blockchain.StateProcessor[
		*types.BeaconBlock,
		components.BeaconState,
		*datypes.BlobSidecars,
		*transition.Context,
		*types.Deposit,
	]
	ValidatorService service.Basic
}

// DepInjectOutput is the output for the dep inject framework.
type DepInjectOutput struct {
	depinject.Out
	Module appmodule.AppModule
}

// ProvideModule is a function that provides the module to the application.
func ProvideModule(in DepInjectInput) (DepInjectOutput, error) {
	payloadCodec := &encoding.
		SSZInterfaceCodec[*types.ExecutionPayloadHeader]{}
	storageBackend := storage.NewBackend[
		*dastore.Store[types.BeaconBlockBody],
		types.BeaconBlockBody,
		core.BeaconState[
			*types.BeaconBlockHeader, *types.ExecutionPayloadHeader, *types.Fork,
			*types.Validator, *engineprimitives.Withdrawal,
		],
	](
		in.ChainSpec,
		in.AvailabilityStore,
		beacondb.New[
			*types.Fork,
			*types.BeaconBlockHeader,
			*types.ExecutionPayloadHeader,
			*types.Eth1Data,
			*types.Validator,
		](in.Environment.KVStoreService, payloadCodec),
		in.DepositStore,
	)

	// TODO: this is hood as fuck.
	if in.BeaconConfig.KZG.Implementation == "" {
		in.BeaconConfig.KZG.Implementation = "crate-crypto/go-kzg-4844"
	}

	runtime, err := components.ProvideRuntime(
		in.BeaconConfig,
		in.BlobVerifier,
		in.ChainSpec,
		in.Signer,
		in.EngineClient,
		in.ExecutionEngine,
		in.BeaconDepositContract,
		in.LocalBuilder,
		in.BlobProofVerifier,
		in.BlockFeed,
		storageBackend,
		in.StateProcessor,
		in.TelemetrySink,
		in.Environment.Logger.With("module", "beacon-kit"),
		in.ValidatorService,
	)
	if err != nil {
		return DepInjectOutput{}, err
	}

	return DepInjectOutput{
		Module: NewAppModule(runtime),
	}, nil
}
