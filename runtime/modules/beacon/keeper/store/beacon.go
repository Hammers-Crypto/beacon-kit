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

package store

import (
	"context"

	sdkcollections "cosmossdk.io/collections"
	corestore "cosmossdk.io/core/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/itsdevbear/bolaris/config"
	"github.com/itsdevbear/bolaris/lib/store/collections"
	"github.com/itsdevbear/bolaris/lib/store/collections/encoding"
	consensusv1 "github.com/itsdevbear/bolaris/types/consensus/v1"
)

// BeaconStore is a wrapper around a KVStore sdk.Context
// that provides access to all beacon related data.
type BeaconStore struct {
	// sdkCtx is the context of the store.
	sdkCtx sdk.Context

	// cfg is the beacon configuration.
	cfg *config.Beacon

	// depositQueue is a list of depositQueue that are queued to be processed.
	depositQueue *collections.Queue[*consensusv1.Deposit]

	// fcSafeEth1BlockHash is the safe block hash.
	fcSafeEth1BlockHash sdkcollections.Item[common.Hash]

	// fcFinalizedEth1BlockHash is the finalized block hash.
	fcFinalizedEth1BlockHash sdkcollections.Item[common.Hash]

	// eth1GenesisHash is the Eth1 genesis hash.
	eth1GenesisHash sdkcollections.Item[common.Hash]

	// lastValidHash is the last valid head in the store.
	// TODO: we need to handle this in a better way.
	lastValidHash *common.Hash
}

// NewBeaconStore creates a new instance of BeaconStore.
func NewBeaconStore(
	ctx context.Context,
	kvs corestore.KVStoreService,
	// TODO: should this be stored in on-chain params?
	cfg *config.Beacon,
) *BeaconStore {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	schemaBuilder := sdkcollections.NewSchemaBuilder(kvs)
	hashValueCodec := encoding.Hash{}
	depositQueue := collections.NewQueue[*consensusv1.Deposit](
		schemaBuilder,
		depositQueuePrefix,
		&consensusv1.Deposit{})
	fcSafeEth1BlockHash := sdkcollections.NewItem[common.Hash](
		schemaBuilder,
		sdkcollections.NewPrefix(fcSafeEth1BlockHashPrefix),
		fcSafeEth1BlockHashPrefix,
		hashValueCodec)
	fcFinalizedEth1BlockHash := sdkcollections.NewItem[common.Hash](
		schemaBuilder,
		sdkcollections.NewPrefix(fcFinalizedEth1BlockHashPrefix),
		fcFinalizedEth1BlockHashPrefix,
		hashValueCodec)
	eth1GenesisHash := sdkcollections.NewItem[common.Hash](
		schemaBuilder,
		sdkcollections.NewPrefix(eth1GenesisHashPrefix),
		eth1GenesisHashPrefix,
		hashValueCodec)
	return &BeaconStore{
		sdkCtx:                   sdkCtx,
		depositQueue:             depositQueue,
		fcSafeEth1BlockHash:      fcSafeEth1BlockHash,
		fcFinalizedEth1BlockHash: fcFinalizedEth1BlockHash,
		eth1GenesisHash:          eth1GenesisHash,
		cfg:                      cfg,
	}
}
