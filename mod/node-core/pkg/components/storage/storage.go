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

package storage

import (
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/state-transition/pkg/core"
	"github.com/berachain/beacon-kit/mod/state-transition/pkg/core/state"
)

// Backend is a struct that holds the storage backend. It provides a simple
// interface to access all types of storage required by the runtime.
type Backend[
	AvailabilityStoreT AvailabilityStore[
		BeaconBlockBodyT, BlobSidecarsT,
	],
	BeaconBlockBodyT types.RawBeaconBlockBody,
	BeaconBlockHeaderT core.BeaconBlockHeader[BeaconBlockHeaderT],
	BeaconStateT core.BeaconState[
		BeaconStateT, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT,
		ForkT, KVStoreT, ValidatorT, WithdrawalT,
	],
	BeaconStateMarshallableT state.BeaconStateMarshallable[
		BeaconStateMarshallableT, BeaconBlockHeaderT, Eth1DataT,
		ExecutionPayloadHeaderT, ForkT, ValidatorT,
	],
	BlobSidecarsT any,
	DepositT Deposit,
	DepositStoreT DepositStore[DepositT],
	Eth1DataT,
	ExecutionPayloadHeaderT,
	ForkT any,
	KVStoreT KVStore[
		KVStoreT, BeaconBlockHeaderT, Eth1DataT,
		ExecutionPayloadHeaderT, ForkT, ValidatorT,
	],
	ValidatorT Validator[WithdrawalCredentialsT],
	WithdrawalT Withdrawal[WithdrawalT],
	WithdrawalCredentialsT WithdrawalCredentials,
] struct {
	chainSpec         common.ChainSpec
	availabilityStore AvailabilityStoreT
	stateStore        KVStoreT
	depositStore      DepositStoreT
}

func NewBackend[
	AvailabilityStoreT AvailabilityStore[
		BeaconBlockBodyT, BlobSidecarsT,
	],
	BeaconBlockBodyT types.RawBeaconBlockBody,
	BeaconBlockHeaderT core.BeaconBlockHeader[BeaconBlockHeaderT],
	BeaconStateT core.BeaconState[
		BeaconStateT, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT,
		ForkT, KVStoreT, ValidatorT, WithdrawalT,
	],
	BeaconStateMarshallableT state.BeaconStateMarshallable[
		BeaconStateMarshallableT, BeaconBlockHeaderT, Eth1DataT,
		ExecutionPayloadHeaderT, ForkT, ValidatorT,
	],
	BlobSidecarsT any,
	DepositT Deposit,
	DepositStoreT DepositStore[DepositT],
	Eth1DataT,
	ExecutionPayloadHeaderT,
	ForkT any,
	KVStoreT KVStore[
		KVStoreT, BeaconBlockHeaderT, Eth1DataT,
		ExecutionPayloadHeaderT, ForkT, ValidatorT,
	],
	ValidatorT Validator[WithdrawalCredentialsT],
	WithdrawalT Withdrawal[WithdrawalT],
	WithdrawalCredentialsT WithdrawalCredentials,
](
	chainSpec common.ChainSpec,
	availabilityStore AvailabilityStoreT,
	stateStore KVStoreT,
	depositStore DepositStoreT,
) *Backend[
	AvailabilityStoreT, BeaconBlockBodyT, BeaconBlockHeaderT, BeaconStateT,
	BeaconStateMarshallableT, BlobSidecarsT, DepositT, DepositStoreT, Eth1DataT,
	ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT,
	WithdrawalT, WithdrawalCredentialsT,
] {
	return &Backend[
		AvailabilityStoreT, BeaconBlockBodyT, BeaconBlockHeaderT, BeaconStateT,
		BeaconStateMarshallableT, BlobSidecarsT, DepositT, DepositStoreT, Eth1DataT,
		ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT,
		WithdrawalT, WithdrawalCredentialsT,
	]{
		chainSpec:         chainSpec,
		availabilityStore: availabilityStore,
		stateStore:        stateStore,
		depositStore:      depositStore,
	}
}

// AvailabilityStore returns the availability store struct initialized with a
// given context.
func (k Backend[
	AvailabilityStoreT, _, _, _, _, _, _, _, _, _, _, _, _, _, _,
]) AvailabilityStore() AvailabilityStoreT {
	return k.availabilityStore
}

// BeaconState returns a new BeaconState initialized with the state store and
// chain spec.
func (k Backend[
	_, _, _, BeaconStateT, _, _, _, _, _, _, _, _, _, _, _,
]) BeaconState() BeaconStateT {
	var st BeaconStateT
	return st.NewFromDB(k.stateStore, k.chainSpec)
}

// DepositStore returns the deposit store struct initialized with a.
func (k Backend[
	_, _, _, _, _, _, _, DepositStoreT, _, _, _, _, _, _, _,
]) DepositStore() DepositStoreT {
	return k.depositStore
}
