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

package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/itsdevbear/bolaris/beacon/execution"
	"github.com/itsdevbear/bolaris/types/consensus/primitives"
	consensusv1 "github.com/itsdevbear/bolaris/types/consensus/v1"
	"github.com/itsdevbear/bolaris/types/engine"
	enginev1 "github.com/itsdevbear/bolaris/types/engine/v1"
)

type ExecutionService interface {
	// NotifyForkchoiceUpdate notifies the execution client of a forkchoice update.
	NotifyForkchoiceUpdate(
		ctx context.Context, fcuConfig *execution.FCUConfig,
	) (*enginev1.PayloadIDBytes, error)

	// NotifyNewPayload notifies the execution client of a new payload.
	NotifyNewPayload(ctx context.Context, preStateHeader engine.ExecutionPayload) (bool, error)

	// ProcessFinalizedLogs processes logs in the finalized execution block.
	ProcessFinalizedLogs(ctx context.Context, blkNum uint64) error
}

type BuilderService interface {
	BuildLocalPayload(
		ctx context.Context,
		parentEth1Hash common.Hash,
		slot primitives.Slot,
		timestamp uint64,
		parentBeaconBlockRoot []byte,
	) (*enginev1.PayloadIDBytes, error)
}

// StakingService is the interface for staking service.
type StakingService interface {
	AcceptDeposit(ctx context.Context, deposit *consensusv1.Deposit) error
	ApplyDeposits(ctx context.Context) error
}
