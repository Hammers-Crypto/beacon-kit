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
	"reflect"
	"testing"

	beacontypes "github.com/berachain/beacon-kit/beacon/core/types"
	loghandler "github.com/berachain/beacon-kit/beacon/execution/logs"
	"github.com/berachain/beacon-kit/beacon/staking/logs"
	"github.com/berachain/beacon-kit/contracts/abi"
	"github.com/berachain/beacon-kit/primitives"
	ethcommon "github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func createValidFactory(
	t *testing.T,
	contractAddress primitives.ExecutionAddress,
) *loghandler.Factory {
	stakingLogRequest, err := logs.NewStakingRequest(
		contractAddress,
	)
	require.NoError(t, err)
	factory, err := loghandler.NewFactory(
		loghandler.WithRequest(stakingLogRequest),
	)
	require.NoError(t, err)
	return factory
}

func TestLogFactory(t *testing.T) {
	// Test Setup
	contractAddress := ethcommon.HexToAddress("0x1234")
	depositContractAbi, err := abi.BeaconDepositContractMetaData.GetAbi()
	require.NoError(t, err)
	require.NotNil(t, depositContractAbi)
	factory := createValidFactory(t, contractAddress)

	// Deposit dummy data.
	event, ok := depositContractAbi.Events[logs.DepositName]
	require.True(t, ok)
	pubKey := []byte("pubkey")
	Credentials := []byte("12345678901234567890123456789012")
	amount := uint64(10000)
	signature := []byte("signature")

	// Create a log from the deposit.
	data, err := event.Inputs.Pack(
		pubKey,
		Credentials,
		amount,
		signature,
	)
	require.NoError(t, err)
	log := &coretypes.Log{
		Topics:  []primitives.ExecutionHash{event.ID},
		Data:    data,
		Address: contractAddress,
	}

	// Unmarshal the log.
	val, err := factory.UnmarshalEthLog(log)
	require.NoError(t, err)

	// Check the type of the unmarshaled value.
	valType := reflect.TypeOf(val.Interface())
	require.NotNil(t, valType)
	require.Equal(t, reflect.Ptr, valType.Kind())
	require.Equal(t, logs.DepositType, valType.Elem())

	// Check the values of the unmarshaled deposit.
	newDeposit, ok := val.Interface().(*beacontypes.Deposit)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, pubKey, newDeposit.ValidatorPubkey)
	require.Equal(t, Credentials, newDeposit.Credentials)
	require.Equal(t, amount, newDeposit.Amount)
	require.Equal(t, signature, newDeposit.Signature)
}

func TestLogFactoryIncorrectType(t *testing.T) {
	// Test Setup
	contractAddress := ethcommon.HexToAddress("0x1234")
	depositContractAbi, err := abi.BeaconDepositContractMetaData.GetAbi()
	require.NoError(t, err)
	require.NotNil(t, depositContractAbi)
	factory := createValidFactory(t, contractAddress)

	// Incorrect dummy data.
	event, ok := depositContractAbi.Events[logs.WithdrawalName]
	require.True(t, ok)
	pubKey := []byte{}
	Credentials := []byte{}
	signature := []byte{}
	amount := uint64(1000)

	// Create a log from the deposit.
	data, err := event.Inputs.Pack(
		pubKey,
		Credentials,
		signature,
		amount,
	)
	require.NoError(t, err)
	log := &coretypes.Log{
		Topics:  []primitives.ExecutionHash{event.ID},
		Data:    data,
		Address: contractAddress,
	}

	_, err = factory.UnmarshalEthLog(log)
	// An error is expected because the event type in ABI and
	// withdrawalType are mismatched,
	// (no validatorPubkey in withdrawalType currently).
	require.Error(t, err)
}
