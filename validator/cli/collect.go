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

package cli

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"

	bankexported "cosmossdk.io/x/bank/exported"
	cfg "github.com/cometbft/cometbft/config"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkruntime "github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"
)

// GenAppStateFromConfig gets the genesis app state from the config.
func GenAppStateFromConfig(
	cdc codec.JSONCodec,
	txEncodingConfig client.TxEncodingConfig,
	config *cfg.Config,
	initCfg types.InitConfig,
	genesis *types.AppGenesis,
	genBalIterator types.GenesisBalancesIterator,
	validator types.MessageValidator,
	valAddrCodec sdkruntime.ValidatorAddressCodec,
) (appState json.RawMessage, err error) {
	// process genesis transactions, else create default genesis.json
	appGenTxs, persistentPeers, err := CollectTxs(
		cdc,
		txEncodingConfig.TxJSONDecoder(),
		config.Moniker,
		initCfg.GenTxsDir,
		genesis,
		genBalIterator,
		validator,
		valAddrCodec,
	)
	if err != nil {
		return appState, err
	}

	config.P2P.PersistentPeers = persistentPeers
	cfg.WriteConfigFile(
		filepath.Join(config.RootDir, "config", "config.toml"),
		config,
	)

	// if there are no gen txs to be processed, return the default empty state
	if len(appGenTxs) == 0 {
		return appState, errors.New("there must be at least one genesis tx")
	}

	// create the app state
	appGenesisState, err := types.GenesisStateFromAppGenesis(genesis)
	if err != nil {
		return appState, err
	}

	appGenesisState, err = genutil.SetGenTxsInAppGenesisState(
		cdc,
		txEncodingConfig.TxJSONEncoder(),
		appGenesisState,
		appGenTxs,
	)
	if err != nil {
		return appState, err
	}

	appState, err = json.MarshalIndent(appGenesisState, "", "  ")
	if err != nil {
		return appState, err
	}

	genesis.AppState = appState
	err = genutil.ExportGenesisFile(genesis, config.GenesisFile())

	return appState, err
}

// CollectTxs processes and validates application's genesis Txs and returns
// the list of appGenTxs, and persistent peers required to generate
// genesis.json.
func CollectTxs(
	cdc codec.JSONCodec,
	txJSONDecoder sdk.TxDecoder,
	moniker, genTxsDir string,
	genesis *types.AppGenesis,
	genBalIterator types.GenesisBalancesIterator,
	validator types.MessageValidator,
	valAddrCodec sdkruntime.ValidatorAddressCodec,
) (appGenTxs []sdk.Tx, persistentPeers string, err error) {
	// prepare a map of all balances in genesis state to then validate
	// against the validators addresses
	var appState map[string]json.RawMessage
	if err = json.Unmarshal(genesis.AppState, &appState); err != nil {
		return appGenTxs, persistentPeers, err
	}

	fos, err := os.ReadDir(genTxsDir)
	if err != nil {
		return appGenTxs, persistentPeers, err
	}

	balancesMap := make(map[string]bankexported.GenesisBalance)

	genBalIterator.IterateGenesisBalances(
		cdc, appState,
		func(balance bankexported.GenesisBalance) bool {
			addr := balance.GetAddress()
			balancesMap[addr] = balance
			return false
		},
	)

	// addresses and IPs (and port) validator server info
	var addressesIPs []string

	// TODO (https://github.com/cosmos/cosmos-sdk/issues/17815):
	// Examine CPU and RAM profiles to see if we can parsing
	// and ValidateAndGetGenTx concurrent.
	for _, fo := range fos {
		if fo.IsDir() {
			continue
		}
		if !strings.HasSuffix(fo.Name(), ".json") {
			continue
		}

		// get the genTx
		var jsonRawTx []byte
		jsonRawTx, err = os.ReadFile(filepath.Join(genTxsDir, fo.Name()))
		if err != nil {
			return appGenTxs, persistentPeers, err
		}

		genTx, err := types.ValidateAndGetGenTx(
			jsonRawTx,
			txJSONDecoder,
			validator,
		)
		if err != nil {
			return appGenTxs, persistentPeers, err
		}

		appGenTxs = append(appGenTxs, genTx)

		// the memo flag is used to store
		// the ip and node-id, for example this may be:
		// "528fd3df22b31f4969b05652bfe8f0fe921321d5@192.168.2.37:26656"

		// memoTx, ok := genTx.(sdk.TxWithMemo)
		// if !ok {
		// 	return appGenTxs, persistentPeers, fmt.Errorf("expected TxWithMemo,
		// got %T", genTx)
		// }
		// nodeAddrIP := memoTx.GetMemo()

		// genesis transactions must be single-message
		// msgs := genTx.Get/Msgs()

		// TODO abstract out staking message validation back to staking
		// msg := msgs[0].(*beacontypes.MsgCreateValidatorX)

		// // validate validator addresses and funds against the accounts in the
		// state
		// valAddr, err := valAddrCodec.StringToBytes(msg.ValidatorAddress)
		// if err != nil {
		// 	return appGenTxs, persistentPeers, err
		// }

		// valAccAddr := sdk.AccAddress(valAddr).String()

		// // delBal, delOk := balancesMap[valAccAddr]
		// // if !delOk {
		// // 	_, file, no, ok := runtime.Caller(1)
		// // 	if ok {
		// // 		fmt.Printf("CollectTxs-1, called from %s#%d\n", file, no)
		// // 	}

		// // 	return appGenTxs, persistentPeers, fmt.Errorf("account %s balance
		// not in genesis state: %+v", valAccAddr, balancesMap)
		// // }

		// _, valOk := balancesMap[valAccAddr]
		// if !valOk {
		// 	_, file, no, ok := runtime.Caller(1)
		// 	if ok {
		// 		fmt.Printf("CollectTxs-2, called from %s#%d - %s\n", file, no,
		// sdk.AccAddress(msg.ValidatorAddress).String())
		// 	}
		// 	return appGenTxs, persistentPeers, fmt.Errorf("account %s balance
		// not in genesis state: %+v", valAddr, balancesMap)
		// }

		// if delBal.GetCoins().AmountOf(msg.Value.Denom).LT(msg.Value.Amount) {
		// 	return appGenTxs, persistentPeers, fmt.Errorf(
		// 		"insufficient fund for delegation %v: %v < %v",
		// 		delBal.GetAddress(), delBal.GetCoins().AmountOf(msg.Value.Denom),
		// msg.Value.Amount,
		// 	)
		// }

		// // exclude itself from persistent peers
		// if msg.Description.Moniker != moniker {
		// 	addressesIPs = append(addressesIPs, nodeAddrIP)
		// }
	}

	sort.Strings(addressesIPs)
	persistentPeers = strings.Join(addressesIPs, ",")

	return appGenTxs, persistentPeers, nil
}
