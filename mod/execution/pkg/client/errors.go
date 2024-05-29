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

package client

import (
	engineerrors "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/errors"
	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/net/http"
	jsonrpc "github.com/berachain/beacon-kit/mod/primitives/pkg/net/json-rpc"
	gethRPC "github.com/ethereum/go-ethereum/rpc"
)

// ErrUnauthenticatedConnection indicates that the connection is not
// authenticated.
//
//nolint:lll
const (
	UnauthenticatedConnectionErrorStr = `could not verify execution chain ID as your 
	connection is not authenticated. If connecting to your execution client via HTTP, you 
	will need to set up JWT authentication...`

	AuthErrMsg = "HTTP authentication to your execution client " +
		"is not working. Please ensure you are setting a correct " +
		"value for the JWT secret path" +
		"is set correctly, or use an IPC " +
		"connection if on the same machine."
)

var (
	// ErrNotStarted indicates that the execution client is not started.
	ErrNotStarted = errors.New("engine client is not started")
)

// Handles errors received from the RPC server according to the specification.
func (s *EngineClient[ExecutionPayloadDenebT]) handleRPCError(err error) error {
	// Exit early if there is no error.
	if err == nil {
		return nil
	}

	// Check for timeout errors.
	if http.IsTimeoutError(err) {
		return http.ErrTimeout
	}

	// Check for connection errors.
	//
	//nolint:errorlint // from prysm.
	e, ok := err.(jsonrpc.Error)
	if !ok {
		if jsonrpc.IsUnauthorizedError(e) {
			return http.ErrUnauthorized
		}
		return errors.Wrapf(
			err,
			"got an unexpected server error in JSON-RPC response "+
				"failed to convert from jsonrpc.Error",
		)
	}

	// Check to see if the error is one of the predefined errors
	// as per the JSON-RPC 2.0 specification.
	if err = jsonrpc.GetPredefinedError(e); err != nil {
		return err
	}

	// Otherwise check for our engine errors.
	switch e.ErrorCode() {
	case -38001:
		// telemetry.IncrCounter(1, MetricKeyUnknownPayloadErrorCount)
		return engineerrors.ErrUnknownPayload
	case -38002:
		// telemetry.IncrCounter(1, MetricKeyInvalidForkchoiceStateCount)
		return engineerrors.ErrInvalidForkchoiceState
	case -38003:
		// telemetry.IncrCounter(1, MetricKeyInvalidPayloadAttributesCount)
		return engineerrors.ErrInvalidPayloadAttributes
	case -38004:
		// telemetry.IncrCounter(1, MetricKeyRequestTooLargeCount)
		return engineerrors.ErrRequestTooLarge
	case -32000:
		// telemetry.IncrCounter(1, MetricKeyInternalServerErrorCount)
		// Only -32000 status codes are data errors in the RPC specification.
		var errWithData gethRPC.DataError
		errWithData, ok = err.(gethRPC.DataError) //nolint:errorlint // from prysm.
		if !ok {
			return errors.Wrapf(
				errors.Join(err.Error(), jsonrpc.ErrServer),
				"got an unexpected data error in JSON-RPC response",
			)
		}
		return errors.Wrapf(jsonrpc.ErrServer, "%v", errWithData.Error())
	default:
		return err
	}
}
