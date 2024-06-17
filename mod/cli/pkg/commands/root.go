// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package commands

import (
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/spf13/cobra"
)

// Root is a wrapper around cobra.Command.
type Root struct {
	cmd *cobra.Command
}

// New returns a new root command with the provided configuration.
func New(name string,
	description string,
	runHandler runHandler,
	clientCtx sdkclient.Context,
) *Root {
	cmd := &cobra.Command{
		Use:   name,
		Short: description,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			var err error
			clientCtx, err = sdkclient.ReadPersistentCommandFlags(
				clientCtx,
				cmd.Flags(),
			)
			if err != nil {
				return err
			}

			customClientTemplate, customClientConfig := InitClientConfig()
			clientCtx, err = config.CreateClientConfig(
				clientCtx,
				customClientTemplate,
				customClientConfig,
			)
			if err != nil {
				return err
			}

			if err = sdkclient.SetCmdClientContextHandler(
				clientCtx, cmd,
			); err != nil {
				return err
			}

			return runHandler(cmd)
		},
	}
	return &Root{
		cmd: cmd,
	}
}

// Run executes the root command.
func (root *Root) Run(defaultNodeHome string) error {
	return svrcmd.Execute(
		root.cmd, "", defaultNodeHome,
	)
}

// Enhance applies the given enhancer to the root command.
func (root *Root) Enhance(enhance enhancer) error {
	return enhance(root.cmd)
}
