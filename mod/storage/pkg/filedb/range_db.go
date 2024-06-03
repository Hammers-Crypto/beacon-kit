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

package filedb

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/hex"
	db "github.com/berachain/beacon-kit/mod/storage/pkg/interfaces"
)

// two is a constant for the number 2.
const two = 2

// Compile-time assertion of prunable interface.
var _ db.Prunable = (*RangeDB)(nil)

// RangeDB is a database that stores versioned data.
// It prefixes keys with an index.
// Invariant: No index below firstNonNilIndex should be populated.
type RangeDB struct {
	db.DB
	firstNonNilIndex uint64
}

// NewRangeDB creates a new RangeDB.
func NewRangeDB(db db.DB) *RangeDB {
	return &RangeDB{
		DB:               db,
		firstNonNilIndex: 0,
	}
}

// Get retrieves the value associated with the given index and key.
// It prefixes the key with the index and a slash before querying the underlying
// database.
func (db *RangeDB) Get(index uint64, key []byte) ([]byte, error) {
	return db.DB.Get(db.prefix(index, key))
}

// Has checks if the given index and key exist in the database.
// It prefixes the key with the index and a slash before querying the underlying
// database.
func (db *RangeDB) Has(index uint64, key []byte) (bool, error) {
	return db.DB.Has(db.prefix(index, key))
}

// Set stores the value with the given index and key in the database.
// It prefixes the key with the index and a slash before storing it in the
// underlying database.
func (db *RangeDB) Set(index uint64, key []byte, value []byte) error {
	// enforce invariant
	if index < db.firstNonNilIndex {
		db.firstNonNilIndex = index
	}
	return db.DB.Set(db.prefix(index, key), value)
}

// Delete removes the value associated with the given index and key from the
// database. It prefixes the key with the index and a slash before deleting it
// from the underlying database.
func (db *RangeDB) Delete(index uint64, key []byte) error {
	return db.DB.Delete(db.prefix(index, key))
}

// DeleteRange removes all values associated with the given index from the
// filesystem. It is INCLUSIVE of the `from` index and EXCLUSIVE of
// the `to“ index.
func (db *RangeDB) DeleteRange(from, to uint64) error {
	f, ok := db.DB.(*DB)
	if !ok {
		return errors.New("rangedb: delete range not supported for this db")
	}
	for ; from < to; from++ {
		if err := f.fs.RemoveAll(fmt.Sprintf("%d/", from)); err != nil {
			return err
		}
	}
	return nil
}

func (db *RangeDB) PruneFromInclusive(startIndex uint64, _ uint64) error {
	if err := db.DeleteRange(db.firstNonNilIndex, startIndex); err != nil {
		// Resets last pruned index in case Delete somehow populates indices on
		// err. This will cause the next prune operation is O(n), but next
		// successful prune will set it to the correct value, so runtime is
		// ammortized
		db.firstNonNilIndex = 0
		return err
	}
	db.firstNonNilIndex = startIndex
	return nil
}

func (db *RangeDB) FirstNonNilIndex() uint64 {
	return db.firstNonNilIndex
}

// prefix prefixes the given key with the index and a slash.
func (db *RangeDB) prefix(index uint64, key []byte) []byte {
	return []byte(fmt.Sprintf("%d/%s", index, hex.FromBytes(key).Unwrap()))
}

// ExtractIndex extracts the index from a prefixed key.
func ExtractIndex(prefixedKey []byte) (uint64, error) {
	parts := bytes.SplitN(prefixedKey, []byte("/"), two)
	if len(parts) < two {
		return 0, errors.New("invalid key format")
	}

	indexStr := string(parts[0])
	index, err := strconv.ParseUint(indexStr, 10, 64)
	if err != nil {
		return 0, errors.Newf("invalid index: %w", err)
	}

	//#nosec:g
	return index, nil
}
