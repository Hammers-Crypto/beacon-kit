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

package signing

import "github.com/berachain/beacon-kit/primitives"

type Version [VersionLength]byte

// ForkData is the fork data used for signing.
type ForkData struct {
	CurrentVersion        Version             `ssz-size:"4"`
	GenesisValidatorsRoot primitives.HashRoot `ssz-size:"32"`
}

// ComputeForkDataRoot computes the root of the fork data.
// Spec:
// def compute_fork_data_root(current_version: Version, genesis_validators_root: Root) -> Root:
//
//	"""
//	Return the 32-byte fork data root for the current_version and genesis_validators_root.
//	This is used primarily in signature domains to avoid collisions across forks/chains.
//	"""
//	return hash_tree_root(ForkData(
//		current_version=current_version,
//		genesis_validators_root=genesis_validators_root,
//	))
func computeForkDataRoot(
	currentVersion Version,
	genesisValidatorsRoot primitives.HashRoot,
) (primitives.HashRoot, error) {
	forkData := ForkData{
		CurrentVersion:        currentVersion,
		GenesisValidatorsRoot: genesisValidatorsRoot,
	}
	return forkData.HashTreeRoot()
}
