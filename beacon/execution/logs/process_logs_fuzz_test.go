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
	"testing"

	beacontypes "github.com/berachain/beacon-kit/beacon/core/types"
	executionlogs "github.com/berachain/beacon-kit/beacon/execution/logs"
	logmocks "github.com/berachain/beacon-kit/beacon/execution/logs/mocks"
	stakinglogs "github.com/berachain/beacon-kit/beacon/staking/logs"
	"github.com/berachain/beacon-kit/primitives"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func FuzzProcessStakingLogs(f *testing.F) {
	contractAddress := ethcommon.HexToAddress("0x1234")

	stakingLogRequest, err := stakinglogs.NewStakingRequest(
		contractAddress,
	)
	require.NoError(f, err)
	logFactory, err := executionlogs.NewFactory(
		executionlogs.WithRequest(stakingLogRequest),
	)
	require.NoError(f, err)

	f.Add(uint64(100), 3, 10)
	f.Fuzz(
		func(t *testing.T, blkNum uint64, depositFactor, numDepositLogs int) {
			if depositFactor <= 0 || numDepositLogs <= 0 {
				t.Skip()
			}

			var logs []ethtypes.Log
			logs, err = logmocks.CreateDepositLogs(
				numDepositLogs,
				depositFactor,
				contractAddress,
				blkNum,
			)
			require.NoError(t, err)

			blockNumToHash := make(map[uint64]primitives.ExecutionHash)
			for _, log := range logs {
				blockNumToHash[log.BlockNumber] = log.BlockHash
			}

			var containers []executionlogs.LogContainer
			containers, err = logFactory.ProcessLogs(logs, blockNumToHash)
			require.NoError(t, err)
			require.Len(t, containers, numDepositLogs)

			// Check if the returned values are correct.
			for i, container := range containers {
				val := container.Value()
				processedDeposit, ok := val.Interface().(*beacontypes.Deposit)
				require.True(t, ok)
				require.Equal(
					t,
					uint64(i*depositFactor),
					processedDeposit.Amount,
				)
			}
		},
	)
}
