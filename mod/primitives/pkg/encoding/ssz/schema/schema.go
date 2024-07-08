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

package schema

import (
	"github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/constants"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/types/types"
)

const chunkSize = constants.BytesPerChunk

type SSZType interface {
	ID() types.Type
	ItemLength() uint64
	ItemPosition(p string) (uint64, uint8, uint8, error)
	HashChunkCount() uint64
	child(p string) SSZType
}

// Container Type

// Enumerable Type (vectors and lists)

type Node struct {
	SSZType

	GIndex uint64
	Offset uint8
}

// GetTreeNode locates a node in the SSZ merkle tree by its path and a root
// schema node to begin traversal from with gindex 1.
//

func GetTreeNode(typ SSZType, path ObjectPath[[32]byte]) (Node, error) {
	gindex, offset, err := path.GetGeneralizedIndex(typ)
	return Node{SSZType: typ, GIndex: uint64(gindex), Offset: offset}, err
}

//nolint:mnd // binary math
func nextPowerOfTwo(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
