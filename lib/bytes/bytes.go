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

package bytes

const (
	// Bytes4Size is the size of a 4 byte array.
	Bytes4Size = 4
)

// Bytes4 is a convenience type for a 4 byte array.
type Bytes4 [Bytes4Size]byte

// SafeCopy will copy and return a non-nil byte slice, otherwise it returns nil.
func SafeCopy(cp []byte) []byte {
	if cp != nil {
		if len(cp) == 32 { //nolint:gomnd // 32 is the size of a hash
			copied := [32]byte(cp)
			return copied[:]
		}
		copied := make([]byte, len(cp))
		copy(copied, cp)
		return copied
	}
	return nil
}

// CopyAndReverseEndianess will copy the input byte slice and return the
// flipped version of it.
func CopyAndReverseEndianess(input []byte) []byte {
	copied := make([]byte, len(input))
	copy(copied, input)
	for i, j := 0, len(copied)-1; i < j; i, j = i+1, j-1 {
		copied[i], copied[j] = copied[j], copied[i]
	}
	return copied
}

// ToBytes32 is a utility function that transforms a byte slice into a fixed
// 32-byte array. If the input exceeds 32 bytes, it gets truncated.
func ToBytes32(input []byte) [32]byte {
	return [32]byte(ExtendToSize(input, 32)) //nolint:gomnd // 32 bytes.
}

// ExtendToSize extends a byte slice to a specified length. It returns the
// original slice if it's already larger.
func ExtendToSize(slice []byte, length int) []byte {
	if len(slice) >= length {
		return slice
	}
	return append(slice, make([]byte, length-len(slice))...)
}

// ToBytes4 is a convenience method for converting
// a byte slice to a fix sized 4 byte array.
// This method will truncate the input if it is larger
// than 4 bytes.
func ToBytes4(x []byte) Bytes4 {
	return Bytes4(PadTo(x, Bytes4Size))
}

// PadTo pads a byte slice to the given size.
// If the byte slice is larger than the given size,
// the original slice is returned.
func PadTo(b []byte, size int) []byte {
	if len(b) >= size {
		return b
	}
	return append(b, make([]byte, size-len(b))...)
}
