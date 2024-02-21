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
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/itsdevbear/bolaris/io/file"
	"github.com/itsdevbear/bolaris/io/jwt"
	"github.com/spf13/cobra"
)

const (
	DefaultSecretFileName = "jwt.hex"
	FlagOutputPath        = "output-path"
)

// NewGenerateJWTCommand creates a new command for generating a JWT secret.
func NewGenerateJWTCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-jwt-secret",
		Short: "Generates a new JWT authentication secret",
		Long: `This command generates a new JWT authentication secret and 
writes it to a file. If no output file path is specified, it uses the default 
file name "jwt.hex" in the current directory.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Get the file path from the command flags.
			outputPath, err := getFilePath(cmd)
			if err != nil {
				return err
			}

			return generateAuthSecretInFile(cmd, outputPath)
		},
	}
	cmd.Flags().StringP(FlagOutputPath, "o", "", "Optional output file path for the JWT secret")
	return cmd
}

// getFilePath retrieves the file path for the JWT secret from the command flags.
// If no path is specified, it returns the default secret file name.
func getFilePath(cmd *cobra.Command) (string, error) {
	specifiedFilePath, err := cmd.Flags().GetString(FlagOutputPath)
	if err != nil {
		return "", err
	}
	if specifiedFilePath != "" {
		return specifiedFilePath, nil
	}

	// If no path is specified, try to get the cosmos client context and use
	// the configured home directory to write the secret to the default file name.
	if specifiedFilePath == "" {
		clientCtx, ok := cmd.Context().Value(client.ClientContextKey).(*client.Context)
		if !ok {
			return "", ErrNoClientCtx
		}
		specifiedFilePath = filepath.Join(clientCtx.HomeDir+"/config/", DefaultSecretFileName)
	}

	return specifiedFilePath, nil // Use default secret file name if no path is specified
}

// generateAuthSecretInFile writes a newly generated JWT secret to a specified file.
func generateAuthSecretInFile(cmd *cobra.Command, fileName string) error {
	var err error
	fileDir := filepath.Dir(fileName)
	exists, err := file.HasDir(fileDir)
	if err != nil {
		return err
	}
	if !exists {
		if err = file.MkdirAll(fileDir); err != nil {
			return err
		}
	}
	secret, err := jwt.NewRandom()
	if err != nil {
		return err
	}
	if err = file.WriteFile(fileName, []byte(secret.Hex())); err != nil {
		return err
	}
	cmd.Printf("Successfully wrote new JSON-RPC authentication secret to: %s", fileName)
	return nil
}
