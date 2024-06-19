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

package phuslu

import "sync"

// byteBuffer is a byte buffer.
type byteBuffer struct {
	Bytes []byte
}

// Write writes to the byte buffer.
func (b *byteBuffer) Write(bytes []byte) (int, error) {
	b.Bytes = append(b.Bytes, bytes...)
	return len(bytes), nil
}

// byteBufferPool is a pool of byte buffers.
//
//nolint:gochecknoglobals // buffer pool
var byteBufferPool = sync.Pool{
	New: func() any {
		return new(byteBuffer)
	},
}

func resetBuffer(b *byteBuffer) {
	if b.Bytes != nil {
		b.Bytes = b.Bytes[:0]
	} else {
		b.Bytes = make([]byte, 0)
	}
}
