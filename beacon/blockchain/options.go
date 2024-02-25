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
	"github.com/itsdevbear/bolaris/beacon/builder/local/cache"
	"github.com/itsdevbear/bolaris/runtime/service"
)

// WithBaseService returns an Option that sets the BaseService for the Service.
func WithBaseService(base service.BaseService) service.Option[Service] {
	return func(s *Service) error {
		s.BaseService = base
		return nil
	}
}

// WithBuilderService is a function that returns an Option.
// It sets the BuilderService of the Service to the provided Service.
func WithBuilderService(bs BuilderService) service.Option[Service] {
	return func(s *Service) error {
		s.bs = bs
		return nil
	}
}

// WithStakingService is a function that returns an Option.
// It sets the StakingService of the Service to the provided Service.
func WithStakingService(ss StakingService) service.Option[Service] {
	return func(s *Service) error {
		s.ss = ss
		return nil
	}
}

// WithExecutionService is a function that returns an Option.
// It sets the ExecutionService of the Service to the provided Service.
func WithExecutionService(es ExecutionService) service.Option[Service] {
	return func(s *Service) error {
		s.es = es
		return nil
	}
}

func WithPayloadCache(
	payloadCache *cache.PayloadIDCache,
) service.Option[Service] {
	return func(s *Service) error {
		s.payloadCache = payloadCache
		return nil
	}
}
