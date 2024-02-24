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
	"errors"
	"sync/atomic"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	eth "github.com/itsdevbear/bolaris/engine/ethclient"
	"github.com/itsdevbear/bolaris/runtime/service"
	"golang.org/x/sync/errgroup"
)

// syncCheckInterval is the interval at which the sync status is checked
// during WaitForExecutionClientSync.
const syncCheckInterval = 3 * time.Second

// Service is responsible for tracking the synchornization status
// of both the beacon and execution chains.
type Service struct {
	service.BaseService
	ethClient        *eth.Eth1Client
	clientCtx        *client.Context
	notifySyncSignal chan struct{}
	synced           atomic.Bool
}

func (s *Service) Status() error {
	if !s.synced.Load() {
		return errors.New("fallen out of sync")
	}
	return nil
}

func (s *Service) SetClientContext(clientCtx client.Context) {
	s.clientCtx = &clientCtx
}

func (s *Service) Start(ctx context.Context) {
	s.notifySyncSignal = make(chan struct{})

	go func() {
		s.syncLoop(ctx)
		<-ctx.Done()
		close(s.notifySyncSignal)
	}()
}

func (s *Service) WaitForSync(ctx context.Context) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		for !s.synced.Load() {
		}
		ch <- struct{}{}
	}()
	return ch
}

// WaitForSync  checks whether the beacon node has sync to the latest head.
func (s *Service) syncLoop(ctx context.Context) error {
	var err error
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.synced.Store(false)
			return ctx.Err()
		case <-ticker.C:
			err = s.RequestSyncProgress(ctx)
		}

		if err == nil {
			s.synced.Store(true)
		} else {
			s.synced.Store(false)
		}
	}
}

func (s *Service) RequestSyncProgress(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// Query Consensus Client for sync progress.
	g.Go(func() error {
		return s.CheckCLSync(ctx)
	})

	// Query Execution Client for sync progress.
	g.Go(func() error {
		return s.CheckELSync(ctx)
	})

	return g.Wait()
}
