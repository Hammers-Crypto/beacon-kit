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

package e2e_test

import (
	"math/big"

	stakingabi "github.com/berachain/beacon-kit/contracts/abi"
	byteslib "github.com/berachain/beacon-kit/lib/bytes"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	// DepositContractAddress is the address of the deposit contract.
	DepositContractAddress = "0x00000000219ab540356cbb839cbe05303d7705fa"
)

// TestDepositContract tests the deposit contract to attempt staking and
// increasing a validator's consensus power.
func (s *BeaconKitE2ESuite) TestDepositContract() {
	// Get the consensus client.
	client := s.ConsensusClients()["cl-validator-beaconkit-0"]
	s.Require().NotNil(client)

	// Get the public key.
	pubkey, err := client.GetPubKey(s.Ctx())
	s.Require().NoError(err)
	s.Require().Len(pubkey, 48)

	// Get the consensus power.
	power, err := client.GetConsensusPower(s.Ctx())
	s.Require().NoError(err)

	// Bind the deposit contract.
	dc, err := stakingabi.NewBeaconDepositContract(
		common.HexToAddress(DepositContractAddress),
		s.JSONRPCBalancer(),
	)
	s.Require().NoError(err)

	// Generate the credentials.
	credentials := byteslib.PrependExtendToSize(
		s.GenesisAccount().Address().Bytes(),
		32,
	)
	credentials[0] = 0x01

	// Generate the signature.
	signature := [96]byte{}
	s.Require().Len(signature[:], 96)

	// Get the chain ID.
	chainID, err := s.JSONRPCBalancer().ChainID(s.Ctx())
	s.Require().NoError(err)

	// Create a deposit transaction.
	val, _ := big.NewFloat(32e18).Int(nil)
	tx, err := dc.Deposit(&bind.TransactOpts{
		From:   s.GenesisAccount().Address(),
		Value:  val,
		Signer: s.GenesisAccount().SignerFunc(chainID),
	}, pubkey, credentials, 0, signature[:])
	s.Require().NoError(err)

	// Wait for the transaction to be mined.
	var receipt *coretypes.Receipt
	receipt, err = bind.WaitMined(s.Ctx(), s.JSONRPCBalancer(), tx)
	s.Require().NoError(err)
	s.Require().Equal(uint64(1), receipt.Status)
	s.Require().True(s.CheckForSuccessfulTx(receipt.TxHash))

	// Check that the consensus power has increased.
	newPower, err := client.GetConsensusPower(s.Ctx())
	s.Require().NoError(err)
	s.Require().Greater(newPower, power)
}
