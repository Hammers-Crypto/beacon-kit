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

package beacon

import (
	"context"
	"encoding/json"

	appmodulev2 "cosmossdk.io/core/appmodule/v2"
	"cosmossdk.io/core/registry"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/genesis"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/node-core/pkg/components"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const (
	// ConsensusVersion defines the current x/beacon module consensus version.
	ConsensusVersion = 1
	// ModuleName is the module name constant used in many places.
	ModuleName = "beacon"
)

var (
	_ appmodulev2.AppModule  = AppModule{}
	_ module.HasABCIGenesis  = AppModule{}
	_ module.HasABCIEndBlock = AppModule{}
)

// AppModule implements an application module for the evm module.
// It is a wrapper around the FinalizeBlockMiddleware and ValidatorMiddleware.
type AppModule struct {
	ABCIFinalizeBlockMiddleware components.FinalizeBlockMiddleware
	ABCIValidatorMiddleware     components.ValidatorMiddleware
}

// NewAppModule creates a new AppModule object.
func NewAppModule(
	abciFinalizedBlockMiddleware components.FinalizeBlockMiddleware,
	abciValidatorMiddleware components.ValidatorMiddleware,
) AppModule {
	return AppModule{
		ABCIFinalizeBlockMiddleware: abciFinalizedBlockMiddleware,
		ABCIValidatorMiddleware:     abciValidatorMiddleware,
	}
}

// Name is the name of this module.
func (am AppModule) Name() string {
	return ModuleName
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 {
	return ConsensusVersion
}

// RegisterInterfaces registers the module's interface types.
func (am AppModule) RegisterInterfaces(registry.InterfaceRegistrar) {}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// DefaultGenesis returns default genesis state as raw bytes
// for the beacon module.
func (AppModule) DefaultGenesis() json.RawMessage {
	bz, err := json.Marshal(
		genesis.DefaultGenesisDeneb(),
	)
	if err != nil {
		panic(err)
	}
	return bz
}

// ValidateGenesis performs genesis state validation for the evm module.
func (AppModule) ValidateGenesis(
	_ json.RawMessage,
) error {
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the evm
// module.
func (am AppModule) ExportGenesis(
	_ context.Context,
) (json.RawMessage, error) {
	return json.Marshal(
		&genesis.Genesis[
			*types.Deposit, *types.ExecutionPayloadHeaderDeneb,
		]{},
	)
}

// InitGenesis
// TODO: InitGenesis should be calling into the StateProcessor.
func (am AppModule) InitGenesis(
	ctx context.Context,
	bz json.RawMessage,
) ([]appmodulev2.ValidatorUpdate, error) {
	return am.ABCIFinalizeBlockMiddleware.InitGenesis(ctx, bz)
}

// EndBlock returns the validator set updates from the beacon state.
func (am AppModule) EndBlock(
	ctx context.Context,
) ([]appmodulev2.ValidatorUpdate, error) {
	return am.ABCIFinalizeBlockMiddleware.EndBlock(ctx)
}
