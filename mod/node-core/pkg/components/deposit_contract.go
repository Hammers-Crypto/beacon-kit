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

package components

import (
	"cosmossdk.io/depinject"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	engineclient "github.com/berachain/beacon-kit/mod/execution/pkg/client"
	"github.com/berachain/beacon-kit/mod/execution/pkg/deposit"
	"github.com/berachain/beacon-kit/mod/primitives"
)

// BeaconDepositContractInput is the input for the beacon deposit contract
// for the dep inject framework.
type BeaconDepositContractInput struct {
	depinject.In
	ChainSpec    primitives.ChainSpec
	EngineClient *engineclient.EngineClient[*types.ExecutionPayload]
}

// ProvideBeaconDepositContract provides a beacon deposit contract through the
// dep inject framework.
func ProvideBeaconDepositContract[
	DepositT deposit.Deposit[DepositT, WithdrawalCredentialsT],
	ExecutionPayloadT engineprimitives.ExecutionPayload[WithdrawalT],
	WithdrawalT any,
	WithdrawalCredentialsT ~[32]byte,
](
	in BeaconDepositContractInput,
) (*deposit.WrappedBeaconDepositContract[
	DepositT,
	WithdrawalCredentialsT,
], error) {
	// Build the deposit contract.
	return deposit.NewWrappedBeaconDepositContract[
		DepositT, WithdrawalCredentialsT,
	](
		in.ChainSpec.DepositContractAddress(),
		in.EngineClient,
	)
}
