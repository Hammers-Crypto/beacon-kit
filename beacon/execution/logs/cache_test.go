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

package logs_test

import (
	"math/rand"
	"testing"

	"github.com/berachain/beacon-kit/beacon/execution/logs"
	logsmocks "github.com/berachain/beacon-kit/beacon/execution/logs/mocks"
	"github.com/stretchr/testify/require"
)

func TestLogCacheSimple(t *testing.T) {
	// Create a new log cache.
	cache := logs.NewCache()

	log := &logsmocks.LogContainer{}
	blockNumber := rand.Uint64() + 1
	logIndex := rand.Uint64() + 1
	log.EXPECT().BlockNumber().Return(blockNumber)
	log.EXPECT().LogIndex().Return(logIndex)
	require.Equal(t, blockNumber, log.BlockNumber())
	require.Equal(t, logIndex, log.LogIndex())
	cache.Insert(log)

	cachedLog, err := cache.RemoveFront()
	require.NoError(t, err)
	require.Equal(t, blockNumber, cachedLog.BlockNumber())
	require.Equal(t, logIndex, cachedLog.LogIndex())
}

func TestLogCacheMulti(t *testing.T) {
	// Create a new log cache.
	cache := logs.NewCache()

	numBlocks := 10
	numLogsPerBlock := 5
	totalLogs := numBlocks * numLogsPerBlock

	logList := make([]*logsmocks.LogContainer, 0, totalLogs)
	for i := 0; i < numBlocks; i++ {
		blockNumber := uint64(i)
		for j := 0; j < numLogsPerBlock; j++ {
			logIndex := uint64(j + 1)
			log := &logsmocks.LogContainer{}
			log.EXPECT().BlockNumber().Return(blockNumber)
			log.EXPECT().LogIndex().Return(logIndex)
			require.Equal(t, blockNumber, log.BlockNumber())
			require.Equal(t, logIndex, log.LogIndex())
			logList = append(logList, log)
		}
	}

	rand.Shuffle(
		len(logList), //#nosec:G703 // not a security-sensitive context.
		func(i, j int) {
			logList[i], logList[j] = logList[j], logList[i]
		})

	for _, log := range logList {
		cache.Insert(log)
	}

	require.Equal(t, totalLogs, cache.Len())
	prevLog := &logsmocks.LogContainer{}
	prevLog.EXPECT().BlockNumber().Return(0)
	prevLog.EXPECT().LogIndex().Return(0)
	c := logs.LogComparable{}
	for range totalLogs {
		log, err := cache.RemoveFront()
		require.NoError(t, err)
		require.Less(t, c.Compare(prevLog, log), 0)
	}

	require.Equal(t, 0, cache.Len())
}
