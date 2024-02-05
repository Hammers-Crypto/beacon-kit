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

package blockchain

import (
	"context"
	"time"

	payloadattribute "github.com/prysmaticlabs/prysm/v4/consensus-types/payload-attribute"
	enginev1 "github.com/prysmaticlabs/prysm/v4/proto/engine/v1"
	"github.com/prysmaticlabs/prysm/v4/runtime/version"
)

// TOOD: fix lint.
var _ = (&Service{}).getPayloadAttribute

// getPayloadAttributes returns the payload attributes for the given state and slot.
// The attribute is required to initiate a payload build process in the
// context of an `engine_forkchoiceUpdated` call.
func (s *Service) getPayloadAttribute(
	ctx context.Context,
) payloadattribute.Attributer {
	var (
		// NOTE: We have to use time.Now() and not the time on the block header coming from
		// Comet or else we attempt to build a block at an equivalent timestamp to the last.
		// TODO: figure out how to fix this, in ethereum I think we need to use the slot to timestamp
		// for the slot math thingy to calculate what the correct timestamp would be for the block we
		// are building.
		t          = uint64(time.Now().Unix()) //#nosec:G701 // won't overflow, time cannot be negative.
		st         = s.BeaconState(ctx)
		emptyAttri = payloadattribute.EmptyWithVersion(st.Version())
		// TODO: RANDAO
		prevRando = make([]byte, 32) //nolint:gomnd // TODO: later
		// TODO: Cancun
		headRoot = make([]byte, 32) //nolint:gomnd // TODO: Cancun
	)

	// TODO: RANDAO
	// // Get previous randao.
	// prevRando, err := helpers.RandaoMix(st, time.CurrentEpoch(st))
	// if err != nil {
	// 	log.WithError(err).Error("Could not get randao mix to get payload attribute")
	// 	return emptyAttri
	// }

	var attr payloadattribute.Attributer
	switch st.Version() {
	case version.Deneb:
		withdrawals, err := st.ExpectedWithdrawals()
		if err != nil {
			s.Logger().Error(
				"Could not get expected withdrawals to get payload attribute", "error", err)
			return emptyAttri
		}
		attr, err = payloadattribute.New(&enginev1.PayloadAttributesV3{
			Timestamp:             t,
			PrevRandao:            prevRando,
			SuggestedFeeRecipient: s.BeaconCfg().Validator.SuggestedFeeRecipient[:],
			Withdrawals:           withdrawals,
			ParentBeaconBlockRoot: headRoot,
		})
		if err != nil {
			s.Logger().Error("Could not get payload attribute", "error", err)
			return emptyAttri
		}
	case version.Capella:
		withdrawals, err := st.ExpectedWithdrawals()
		if err != nil {
			s.Logger().Error(
				"Could not get expected withdrawals to get payload attribute", "error", err)
			return emptyAttri
		}
		attr, err = payloadattribute.New(&enginev1.PayloadAttributesV2{
			Timestamp:             t,
			PrevRandao:            prevRando,
			SuggestedFeeRecipient: s.BeaconCfg().Validator.SuggestedFeeRecipient[:],
			Withdrawals:           withdrawals,
		})
		if err != nil {
			s.Logger().Error(
				"Could not get payload attribute", "error", err)
			return emptyAttri
		}
	case version.Bellatrix:
		var err error
		attr, err = payloadattribute.New(&enginev1.PayloadAttributes{
			Timestamp:             t,
			PrevRandao:            prevRando,
			SuggestedFeeRecipient: s.BeaconCfg().Validator.SuggestedFeeRecipient[:],
		})
		if err != nil {
			s.Logger().Error("Could not get payload attribute", "error", err)
			return emptyAttri
		}
	default:
		s.Logger().Error(
			"Could not get payload attribute due to unknown state version", "version", st.Version())
		return emptyAttri
	}

	return attr
}
