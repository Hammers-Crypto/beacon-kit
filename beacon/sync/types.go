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

package sync

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/itsdevbear/bolaris/beacon/execution"
	"github.com/itsdevbear/bolaris/beacon/state"
	enginev1 "github.com/itsdevbear/bolaris/types/engine/v1"
)

type (
	// Status represents the synchronization status of the initial sync service.
	Status int

	// BeaconSyncStatus represents the synchronization status of the beacon
	// chain.
	BeaconSyncProgress struct {
		Status      Status
		clFinalized common.Hash
		elFinalized common.Hash
	}
)

const (
	// StatusWaiting indicates the initial status of the service, waiting
	// for further actions.
	StatusWaiting = Status(iota)
	// StatusBeaconAhead indicates that the beacon chain is ahead of the
	//  execution chain.
	StatusBeaconAhead
	// StatusExecutionAhead indicates that the execution chain is ahead
	// of the beacon chain.
	StatusExecutionAhead
	// StatusSynced indicates that both the beacon and execution chains
	// are synchronized.
	StatusSynced
)

// BeaconStateProvider is an interface that wraps the basic BeaconState method.
type BeaconStateProvider interface {
	// BeaconState provides access to the underlying beacon state.
	BeaconState(ctx context.Context) state.BeaconState
}

// executionService defines an interface for interacting with the execution
// client.
type executionService interface {
	// NotifyForkchoiceUpdate notifies the execution client of a forkchoice
	// update.
	NotifyForkchoiceUpdate(
		ctx context.Context, fcuConfig *execution.FCUConfig,
	) (*enginev1.PayloadIDBytes, error)
}
