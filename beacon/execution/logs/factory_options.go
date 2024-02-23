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

package logs

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/itsdevbear/bolaris/runtime/service"
)

// WithTypeAllocator returns an Option for
// registering the TypeAllocator under the given
// address with the Factory.
func WithTypeAllocator(
	contractAddress common.Address,
	allocator *TypeAllocator,
) service.Option[Factory] {
	return func(f *Factory) error {
		f.addressToAllocator[contractAddress] = allocator
		return nil
	}
}

// WithInitilizer returns an Option for initializing the Factory.
// With this function called at the beginning,
// WithTypeAllocator does not need to check
// if the map was initialized or not.
func WithInitilizer() service.Option[Factory] {
	return func(f *Factory) error {
		f.addressToAllocator = make(map[common.Address]*TypeAllocator)
		return nil
	}
}
