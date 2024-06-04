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

package state

import (
	"fmt"

	deneb "github.com/berachain/beacon-kit/mod/consensus-types/pkg/state/deneb"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/version"
)

// BeaconState is the interface for the beacon state.
type BeaconState struct {
	// TODO: decouple from deneb.BeaconState
	*deneb.BeaconState
}

// New creates a new BeaconState.
func (st *BeaconState) New(
	forkVersion uint32,
	genesisValidatorsRoot primitives.Root,
	slot math.Slot,
	fork *types.Fork,
	latestBlockHeader *types.BeaconBlockHeader,
	blockRoots []primitives.Root,
	stateRoots []primitives.Root,
	eth1Data *types.Eth1Data,
	eth1DepositIndex uint64,
	latestExecutionPayloadHeader *types.ExecutionPayloadHeader,
	validators []*types.Validator,
	balances []uint64,
	randaoMixes []primitives.Bytes32,
	nextWithdrawalIndex uint64,
	nextWithdrawalValidatorIndex math.ValidatorIndex,
	slashings []uint64,
	totalSlashing math.Gwei,
) (*BeaconState, error) {
	switch forkVersion {
	case version.Deneb:
		return &BeaconState{
			BeaconState: &deneb.BeaconState{
				Slot:                  slot,
				GenesisValidatorsRoot: genesisValidatorsRoot,
				Fork:                  fork,
				LatestBlockHeader:     latestBlockHeader,
				BlockRoots:            blockRoots,
				StateRoots:            stateRoots,
				
				LatestExecutionPayloadHeader: latestExecutionPayloadHeader.
					ExecutionPayloadHeader.(*types.ExecutionPayloadHeaderDeneb),
				Eth1Data:                     eth1Data,
				Eth1DepositIndex:             eth1DepositIndex,
				Validators:                   validators,
				Balances:                     balances,
				RandaoMixes:                  randaoMixes,
				NextWithdrawalIndex:          nextWithdrawalIndex,
				NextWithdrawalValidatorIndex: nextWithdrawalValidatorIndex,
				Slashings:                    slashings,
				TotalSlashing:                totalSlashing,
			},
		}, nil
	default:
		return nil, fmt.Errorf("unsupported version %d", forkVersion)
	}
}
