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

package beacon_test

import (
	"testing"

	sdklog "cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	sdkruntime "github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil/integration"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	beaconstore "github.com/itsdevbear/bolaris/store/beacon"
	consensusv1 "github.com/itsdevbear/bolaris/types/consensus/v1"
	"github.com/stretchr/testify/require"
)

func TestBeaconStore(t *testing.T) {
	testName := "test"
	logger := sdklog.NewNopLogger()
	keys := storetypes.NewKVStoreKeys(testName)
	cms := integration.CreateMultiStore(keys, logger)
	ctx := sdk.NewContext(cms, true, logger)
	storeKey := keys[testName]
	kvs := sdkruntime.NewKVStoreService(storeKey)
	kv := ctx.KVStore(storeKey)

	beaconStore := beaconstore.NewStore(kvs)
	beaconStore = beaconStore.WithContext(ctx)
	t.Run("should return correct hashes", func(t *testing.T) {
		safeHash := common.HexToHash("0x123")
		beaconStore.SetSafeEth1BlockHash(safeHash)
		hash := beaconStore.GetSafeEth1BlockHash()
		require.Equal(t, safeHash, hash)
		hash.SetBytes([]byte("0x789"))
		require.Equal(t, safeHash, beaconStore.GetSafeEth1BlockHash())
		newSafeHash := common.HexToHash("0x456")
		beaconStore.SetSafeEth1BlockHash(newSafeHash)
		require.Equal(t, newSafeHash, beaconStore.GetSafeEth1BlockHash())

		finalHash := common.HexToHash("0x456")
		beaconStore.SetFinalizedEth1BlockHash(finalHash)
		require.Equal(t, finalHash, beaconStore.GetFinalizedEth1BlockHash())
		// Recheck to make sure there is no collision.
		require.Equal(t, newSafeHash, beaconStore.GetSafeEth1BlockHash())

		genesisHash := common.HexToHash("0x789")
		beaconStore.SetGenesisEth1Hash(genesisHash)
		require.Equal(t, genesisHash, beaconStore.GenesisEth1Hash())
		require.Equal(t, finalHash, beaconStore.GetFinalizedEth1BlockHash())
		require.Equal(t, newSafeHash, beaconStore.GetSafeEth1BlockHash())
	})

	t.Run("should not have state breaking", func(t *testing.T) {
		safeHash := common.HexToHash("0x123")
		kv.Set([]byte("fc_safe"), safeHash.Bytes())
		require.Equal(t, safeHash, beaconStore.GetSafeEth1BlockHash())

		finalHash := common.HexToHash("0x456")
		kv.Set([]byte("fc_finalized"), finalHash.Bytes())
		require.Equal(t, finalHash, beaconStore.GetFinalizedEth1BlockHash())

		genesisHash := common.HexToHash("0x789")
		kv.Set([]byte("eth1_genesis_hash"), genesisHash.Bytes())
		require.Equal(t, genesisHash, beaconStore.GenesisEth1Hash())

		require.Equal(t, safeHash, beaconStore.GetSafeEth1BlockHash())
		require.Equal(t, finalHash, beaconStore.GetFinalizedEth1BlockHash())
	})

	t.Run("should work with deposit", func(t *testing.T) {
		deposit := &consensusv1.Deposit{
			ValidatorPubkey:       []byte("pubkey"),
			WithdrawalCredentials: []byte("12345678901234567890"),
			Amount:                100,
		}
		err := beaconStore.EnqueueDeposits([]*consensusv1.Deposit{deposit})
		require.NoError(t, err)
		deposits, err := beaconStore.DequeueDeposits(1)
		require.NoError(t, err)
		require.Equal(t, deposit, deposits[0])
	})

	t.Run("should return correct log info", func(t *testing.T) {
		logSig := common.HexToHash("0x123")
		blkNum := uint64(100)
		index := uint64(10)
		err := beaconStore.SetLastProcessedBlock(logSig, blkNum)
		require.NoError(t, err)
		err = beaconStore.SetLastProcessedIndex(logSig, index)
		require.NoError(t, err)

		lastProcessedBlock, err := beaconStore.GetLastProcessedBlock(logSig)
		require.NoError(t, err)
		require.Equal(t, blkNum, lastProcessedBlock)

		lastProcessedIndex, err := beaconStore.GetLastProcessedIndex(logSig)
		require.NoError(t, err)
		require.Equal(t, index, lastProcessedIndex)

		invalidLogSig := common.HexToHash("0x456")
		_, err = beaconStore.GetLastProcessedBlock(invalidLogSig)
		require.Error(t, err)
		_, err = beaconStore.GetLastProcessedIndex(invalidLogSig)
		require.Error(t, err)
	})
}
