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

package builder

import (
	consensustypes "github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	dastore "github.com/berachain/beacon-kit/mod/da/pkg/store"
	datypes "github.com/berachain/beacon-kit/mod/da/pkg/types"
	"github.com/berachain/beacon-kit/mod/node-core/pkg/components"
	"github.com/berachain/beacon-kit/mod/node-core/pkg/components/storage"
	"github.com/berachain/beacon-kit/mod/runtime/pkg/runtime"
	"github.com/berachain/beacon-kit/mod/runtime/pkg/runtime/middleware"
	depositdb "github.com/berachain/beacon-kit/mod/storage/pkg/deposit"
)

// dummy components for the builder. these are needed to support the
// direct dependencies needed to supply the Module.
//
//nolint:gochecknoglobals // we eventually shouldn't need to supply empty vars
var (
	// MIDDLEWARES.
	emptyFinalizeBlockMiddlware = &middleware.FinalizeBlockMiddleware[
		*consensustypes.BeaconBlock,
		runtime.BeaconState,
		*datypes.BlobSidecars,
	]{}
	emptyValidatorMiddleware = &middleware.ValidatorMiddleware[
		*dastore.Store[*consensustypes.BeaconBlockBody],
		*consensustypes.BeaconBlock,
		*consensustypes.BeaconBlockBody,
		runtime.BeaconState,
		*datypes.BlobSidecars,
		runtime.Backend,
	]{}
	// STORAGE BACKEND.
	emptyStorageBackend = &storage.Backend[
		*dastore.Store[*consensustypes.BeaconBlockBody],
		*consensustypes.BeaconBlock,
		*consensustypes.BeaconBlockBody,
		components.BeaconState,
		*depositdb.KVStore[*consensustypes.Deposit],
	]{}
)
