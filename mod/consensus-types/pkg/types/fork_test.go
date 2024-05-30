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

package types_test

import (
	"testing"

	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/stretchr/testify/require"
)

func TestFork_Serialization(t *testing.T) {
	// Create a Fork
	original := &types.Fork{
		PreviousVersion: common.Version{1, 2, 3, 4},
		CurrentVersion:  common.Version{5, 6, 7, 8},
		Epoch:           math.Epoch(1000),
	}

	// Marshal the Fork to bytes
	data, err := original.MarshalSSZ()
	require.NoError(t, err)

	// Unmarshal the bytes into a new Fork
	var unmarshalled types.Fork
	err = unmarshalled.UnmarshalSSZ(data)
	require.NoError(t, err)

	// The original and unmarshalled Fork should be the same
	require.Equal(t, original, &unmarshalled)
}

func TestFork_SizeSSZ(t *testing.T) {
	// Create a Fork
	fork := &types.Fork{
		PreviousVersion: common.Version{1, 2, 3, 4},
		CurrentVersion:  common.Version{5, 6, 7, 8},
		Epoch:           math.Epoch(1000),
	}

	// Get the SSZ size of the Fork
	size := fork.SizeSSZ()

	// The size should be 16
	require.Equal(t, 16, size)
}

func TestFork_HashTreeRoot(t *testing.T) {
	// Create a Fork
	fork := &types.Fork{
		PreviousVersion: common.Version{1, 2, 3, 4},
		CurrentVersion:  common.Version{5, 6, 7, 8},
		Epoch:           math.Epoch(1000),
	}

	// Get the hash tree root of the Fork
	_, err := fork.HashTreeRoot()

	// There should be no error
	require.NoError(t, err)
}

func TestFork_GetTree(t *testing.T) {
	// Create a Fork
	fork := &types.Fork{
		PreviousVersion: common.Version{1, 2, 3, 4},
		CurrentVersion:  common.Version{5, 6, 7, 8},
		Epoch:           math.Epoch(1000),
	}

	// Get the SSZ tree of the Fork
	tree, err := fork.GetTree()

	// There should be no error
	require.NoError(t, err)

	// The tree should not be nil
	require.NotNil(t, tree)
}
