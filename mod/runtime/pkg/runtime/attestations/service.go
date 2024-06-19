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

package attestations

import (
	"github.com/berachain/beacon-kit/mod/log"
	lru "github.com/hashicorp/golang-lru"
)

// Service represenst the deposit service that processes deposit events.
type Service[
	AttestationDataT AttestationData,
] struct {

	// logger is used for logging information and errors.
	logger log.Logger[any]

	// // ds is the deposit store that stores deposits.
	// ds Store[DepositT]

	// cache is an in memory store for attestation data
	cache *lru.Cache
	// // feed is the block feed that provides block events.
	// feed BlockFeed[
	// 	DepositT,
	// 	BeaconBlockBodyT,
	// 	BeaconBlockT,
	// 	BlockEventT,
	// 	ExecutionPayloadT,
	// 	SubscriptionT,
	// ]
	// // metrics is the metrics for the deposit service.
	// metrics *depositMetrics
}

// NewService creates a new instance of the Service struct.
func NewService[
	AttestationDataT AttestationData,
](
	logger log.Logger[any],
) *Service[AttestationDataT] {
	// TODO: Move this to a config
	const cacheSize = 2048
	cache, err := lru.New(cacheSize)
	if err != nil {
		logger.Error("failed to create cache", "error", err)
	}
	return &Service[AttestationDataT]{
		logger: logger,
		cache:  cache,
	}
}

// // NewService creates a new instance of the Service struct.
// func NewService[
// 	// BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT],
// 	// BeaconBlockT BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT],
// 	// BlockEventT BlockEvent[
// 	// 	DepositT, BeaconBlockBodyT,
// 	// 	BeaconBlockT, ExecutionPayloadT,
// 	// ],
// 	DepositStoreT Store[DepositT],
// 	ExecutionPayloadT interface{ GetNumber() math.U64 },
// 	SubscriptionT interface {
// 		Unsubscribe()
// 	},
// 	WithdrawalCredentialsT any,
// 	DepositT Deposit[DepositT, WithdrawalCredentialsT],
// ](
// 	logger log.Logger[any],
// 	eth1FollowDistance math.U64,
// 	ethclient EthClient,
// 	telemetrySink TelemetrySink,
// 	ds Store[DepositT],
// 	dc Contract[DepositT],
// 	feed BlockFeed[
// 		DepositT, BeaconBlockBodyT, BeaconBlockT, BlockEventT,
// 		ExecutionPayloadT, SubscriptionT,
// 	],
// ) *Service[
// 	BeaconBlockT, BeaconBlockBodyT, BlockEventT, DepositT,
// 	ExecutionPayloadT, SubscriptionT,
// 	WithdrawalCredentialsT,
// ] {
// 	return &Service[
// 		BeaconBlockT, BeaconBlockBodyT, BlockEventT, DepositT,
// 		ExecutionPayloadT, SubscriptionT,
// 		WithdrawalCredentialsT,
// 	]{
// 		feed:               feed,
// 		logger:             logger,
// 		ethclient:          ethclient,
// 		eth1FollowDistance: eth1FollowDistance,
// 		metrics:            newDepositMetrics(telemetrySink),
// 		dc:                 dc,
// 		ds:                 ds,
// 		newBlock:           make(chan BeaconBlockT),
// 		failedBlocks:       make(map[math.U64]struct{}),
// 	}
// }

// // Start starts the service and begins processing block events.
// func (s *Service[
// 	BeaconBlockT, BeaconBlockBodyT, BlockEventT,
// 	ExecutionPayloadT, SubscriptionT,
// 	WithdrawalCredentialsT, DepositT,
// ]) Start(
// 	ctx context.Context,
// ) error {
// 	go s.blockFeedListener(ctx)
// 	go s.depositFetcher(ctx)
// 	go s.depositCatchupFetcher(ctx)
// 	return nil
// }

// func (s *Service[
// 	BeaconBlockT, BeaconBlockBodyT, BlockEventT,
// 	ExecutionPayloadT, SubscriptionT,
// 	WithdrawalCredentialsT, DepositT,
// ]) blockFeedListener(ctx context.Context) {
// 	ch := make(chan BlockEventT)
// 	sub := s.feed.Subscribe(ch)
// 	defer sub.Unsubscribe()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		case event := <-ch:
// 			if event.Is(events.BeaconBlockFinalized) {
// 				s.newBlock <- event.Data()
// 			}
// 		}
// 	}
// }

// // Name returns the name of the service.
// func (s *Service[
// 	BeaconBlockT, BeaconBlockBodyT, BlockEventT,
// 	ExecutionPayloadT, SubscriptionT,
// 	WithdrawalCredentialsT, DepositT,
// ]) Name() string {
// 	return "deposit-handler"
// }

// // Status returns the current status of the service.
// func (s *Service[
// 	BeaconBlockT, BeaconBlockBodyT, BlockEventT,
// 	ExecutionPayloadT, SubscriptionT,
// 	WithdrawalCredentialsT, DepositT,
// ]) Status() error {
// 	return nil
// }

// // WaitForHealthy waits for the service to become healthy.
// func (s *Service[
// 	BeaconBlockT, BeaconBlockBodyT, BlockEventT,
// 	ExecutionPayloadT, SubscriptionT,
// 	WithdrawalCredentialsT, DepositT,
// ]) WaitForHealthy(
// 	_ context.Context,
// ) {
// }