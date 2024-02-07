// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
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

package config

import (
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/itsdevbear/bolaris/config/flags"
	"github.com/itsdevbear/bolaris/config/parser"
	"github.com/spf13/cobra"
)

// BeaconKitConfig is the interface for a sub-config of the global BeaconKit configuration.
type BeaconKitConfig[T any] interface {
	Template() string
	Parse(parser parser.AppOptionsParser) (*T, error)
}

// Config is the main configuration struct for the BeaconKit chain.
type Config struct {
	// ExecutionClient is the configuration for the execution client.
	ExecutionClient ExecutionClient

	// Beacon is the configuration for the fork epochs.
	Beacon Beacon

	// ABCI is the configuration for ABCI related settings.
	ABCI ABCI
}

// Template returns the configuration template.
func (c Config) Template() string {
	return `
###############################################################################
###                                BeaconKit                                ###
###############################################################################
` + c.ExecutionClient.Template() + c.Beacon.Template() + c.ABCI.Template()
}

// DefaultConfig returns the default configuration for a BeaconKit chain.
func DefaultConfig() *Config {
	return &Config{
		ExecutionClient: DefaultExecutionClientConfig(),
		Beacon:          DefaultBeaconConfig(),
		ABCI:            DefaultABCIConfig(),
	}
}

// SetupCosmosConfig sets up the Cosmos SDK configuration to be compatible with the
// semantics of etheruem.
func SetupCosmosConfig() {
	// set the address prefixes
	config := sdk.GetConfig()

	// We use CoinType == 60 to match Ethereum.
	// This is not strictly necessary, though highly recommended.
	config.SetCoinType(60) //nolint:gomnd // its okay.
	config.SetPurpose(sdk.Purpose)
	config.Seal()
}

// MustReadConfigFromAppOpts reads the configuration options from the given
// application options. Panics if the configuration cannot be read.
func MustReadConfigFromAppOpts(opts servertypes.AppOptions) *Config {
	cfg, err := ReadConfigFromAppOpts(opts)
	if err != nil {
		panic(err)
	}
	return cfg
}

// ReadConfigFromAppOpts reads the configuration options from the given
// application options.
func ReadConfigFromAppOpts(opts servertypes.AppOptions) (*Config, error) {
	return readConfigFromAppOptsParser(parser.AppOptionsParser{AppOptions: opts})
}

// readConfigFromAppOptsParser reads the configuration options from the given.
func readConfigFromAppOptsParser(parser parser.AppOptionsParser) (*Config, error) {
	var (
		err                error
		conf               = &Config{}
		executionClientCfg *ExecutionClient
		beaconCfg          *Beacon
		abciCfg            *ABCI
	)
	// Read Execution Client Config
	executionClientCfg, err = ExecutionClient{}.Parse(parser)
	if err != nil {
		return nil, err
	}
	conf.ExecutionClient = *executionClientCfg

	// Read Beacon Config
	beaconCfg, err = Beacon{}.Parse(parser)
	if err != nil {
		return nil, err
	}
	conf.Beacon = *beaconCfg

	// Read ABCI Config
	abciCfg, err = ABCI{}.Parse(parser)
	if err != nil {
		return nil, err
	}
	conf.ABCI = *abciCfg

	return conf, nil
}

// AddBeaconKitFlags implements servertypes.ModuleInitFlags interface.
func AddBeaconKitFlags(startCmd *cobra.Command) {
	defaultCfg := DefaultConfig()
	startCmd.Flags().String(flags.JWTSecretPath, defaultCfg.ExecutionClient.JWTSecretPath,
		"path to the execution client secret")
	startCmd.Flags().String(flags.RPCDialURL, defaultCfg.ExecutionClient.RPCDialURL, "rpc dial url")
	startCmd.Flags().Uint64(flags.RPCRetries, defaultCfg.ExecutionClient.RPCRetries, "rpc retries")
	startCmd.Flags().Uint64(flags.RPCTimeout, defaultCfg.ExecutionClient.RPCTimeout, "rpc timeout")
	startCmd.Flags().Uint64(flags.RequiredChainID, defaultCfg.ExecutionClient.RequiredChainID,
		"required chain id")
	startCmd.Flags().String(flags.SuggestedFeeRecipient,
		defaultCfg.Beacon.Validator.SuggestedFeeRecipient.Hex(),
		"suggested fee recipient",
	)
}
