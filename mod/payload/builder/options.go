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

package builder

import (
	"cosmossdk.io/log"
	"github.com/berachain/beacon-kit/mod/payload/cache"
	"github.com/berachain/beacon-kit/mod/primitives"
	engineprimitves "github.com/berachain/beacon-kit/mod/primitives-engine"
	"github.com/berachain/beacon-kit/mod/primitives/math"
)

// Option is a functional option for the builder.
type Option = func(*Service) error

// WithChainSpec sets the chain spec.
func WithChainSpec(chainSpec primitives.ChainSpec) Option {
	return func(s *Service) error {
		s.chainSpec = chainSpec
		return nil
	}
}

// WithConfig sets the builder config.
func WithConfig(cfg *Config) Option {
	return func(s *Service) error {
		s.cfg = cfg
		return nil
	}
}

// WithLogger sets the logger.
func WithLogger(logger log.Logger) Option {
	return func(s *Service) error {
		s.logger = logger
		return nil
	}
}

// WithExecutionEngine sets the execution engine.
func WithExecutionEngine(ee ExecutionEngine) Option {
	return func(s *Service) error {
		s.ee = ee
		return nil
	}
}

// WithPayloadCache sets the payload cache.
func WithPayloadCache(
	pc *cache.PayloadIDCache[
		engineprimitves.PayloadID, [32]byte, math.Slot,
	],
) Option {
	return func(s *Service) error {
		s.pc = pc
		return nil
	}
}
