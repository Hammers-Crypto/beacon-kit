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

//go:build !ckzg

package ckzg

import (
	prooftypes "github.com/berachain/beacon-kit/mod/da/pkg/kzg/types"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844"
)

// VerifyBlobProof will error since cgo is not enabled.
func (v Verifier) VerifyBlobProof(
	*eip4844.Blob,
	eip4844.KZGProof,
	eip4844.KZGCommitment,
) error {
	return ErrCGONotEnabled
}

// VerifyBlobProofBatch will error since cgo is not enabled.
func (v Verifier) VerifyBlobProofBatch(
	*prooftypes.BlobProofArgs,
) error {
	return ErrCGONotEnabled
}
