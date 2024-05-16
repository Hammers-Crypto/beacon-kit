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
	"context"
	"net"
	"net/rpc"
	"os"
)

//nolint:lll // long line length due to struct tags.
func (s *EngineClient[ExecutionPayloadDenebT]) startIPCServer(ctx context.Context) {
	// remove existing socket file if exists
	// alternatively we can use existing one by checking for os.IsNotExist(err)
	if _, err := os.Stat(s.cfg.IPCPath); err != nil {
		s.logger.Info("Removing existing IPC file", "path", s.cfg.IPCPath)
		os.Remove(s.cfg.IPCPath)
	}

	// use UDS for IPC
	listener, err := net.Listen("unix", s.cfg.IPCPath)
	if err != nil {
		s.logger.Error("failed to listen on IPC socket", "err", err)
		return
	}
	s.ipcListener = listener

	// register the RPC server
	server := rpc.NewServer()
	if err = server.Register(s); err != nil {
		s.logger.Error("failed to register RPC server", "err", err)
		return
	}
	s.logger.Info("IPC server started", "path", s.cfg.IPCPath)

	// start server in a goroutine
	go func() {
		for {
			// continuously accept incoming connections until context is cancelled
			select {
			case <-ctx.Done():
				s.logger.Info("shutting down IPC server")
				return
			default:
				var conn net.Conn
				conn, err = listener.Accept()
				if err != nil {
					s.logger.Error("failed to accept IPC connection", "err", err)
					continue
				}
				go server.ServeConn(conn)
			}
		}
	}()
}
