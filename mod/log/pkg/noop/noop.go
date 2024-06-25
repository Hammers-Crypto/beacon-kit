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

package noop

// Logger is a logger that performs no operations. It can be used in
// environments where logging should be disabled. It implements the Logger
// interface with no-op methods.
type Logger[KeyValT any] struct{}

// NewLogger creates a blank no-op logger.
func NewLogger() *Logger[any] {
	return &Logger[any]{}
}

// Info logs an informational message with associated key-value pairs. This
// method does nothing.
func (n *Logger[KeyValT]) Info(string, ...KeyValT) {
	// No operation
}

// Warn logs a warning message with associated key-value pairs. This method does
// nothing.
func (n *Logger[KeyValT]) Warn(string, ...KeyValT) {
	// No operation
}

// Error logs an error message with associated key-value pairs. This method does
// nothing.
func (n *Logger[KeyValT]) Error(string, ...KeyValT) {
	// No operation
}

// Debug logs a debug message with associated key-value pairs. This method does
// nothing.
func (n *Logger[KeyValT]) Debug(string, ...KeyValT) {
	// No operation
}

// AdvancedLogger is a Logger that performs no operations.
type AdvancedLogger[KeyValT any, ImplT any] struct{}

// NewAdvancedLogger creates a blank no-op AdvancedLogger.
func NewAdvancedLogger[ImplT any]() *AdvancedLogger[any, ImplT] {
	return &AdvancedLogger[any, ImplT]{}
}

// Info logs an informational message with associated key-value pairs. This
// method does nothing.
func (n *AdvancedLogger[KeyValT, ImplT]) Info(string, ...KeyValT) {
	// No operation
}

// Warn logs a warning message with associated key-value pairs. This method does
// nothing.
func (n *AdvancedLogger[KeyValT, ImplT]) Warn(string, ...KeyValT) {
	// No operation
}

// Error logs an error message with associated key-value pairs. This method does
// nothing.
func (n *AdvancedLogger[KeyValT, ImplT]) Error(string, ...KeyValT) {
	// No operation
}

// Debug logs a debug message with associated key-value pairs. This method does
// nothing.
func (n *AdvancedLogger[KeyValT, ImplT]) Debug(string, ...KeyValT) {
	// No operation
}

// With returns a new AdvancedLogger with the provided key-value pairs. This
// method does nothing.
func (n *AdvancedLogger[KeyValT, ImplT]) With(...KeyValT) ImplT {
	return any(n).(ImplT)
}

func (n *AdvancedLogger[KeyValT, ImplT]) Impl() any {
	return nil
}
