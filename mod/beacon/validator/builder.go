package validator

import (
	"context"

	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/errors"
	engineprimitives "github.com/berachain/beacon-kit/mod/primitives-engine"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
)

// GetEmptyBlock creates a new empty block.
func (s *Service[
	BeaconStateT,
	BlobSidecarsT,
]) GetEmptyBeaconBlock(st BeaconStateT, slot math.Slot) (types.BeaconBlock, error) {
	// Create a new block.
	parentBlockRoot, err := st.GetBlockRootAtIndex(
		uint64(slot) % s.chainSpec.SlotsPerHistoricalRoot(),
	)
	if err != nil {
		return nil, errors.Newf(
			"failed to get block root at index: %w",
			err,
		)
	}

	// Get the proposer index for the slot.
	proposerIndex, err := st.ValidatorIndexByPubkey(
		s.signer.PublicKey(),
	)
	if err != nil {
		return nil, errors.Newf(
			"failed to get validator by pubkey: %w",
			err,
		)
	}

	// Create a new empty block from the current state.
	return types.EmptyBeaconBlock(
		slot,
		proposerIndex,
		parentBlockRoot,
		s.chainSpec.ActiveForkVersionForSlot(slot),
	)
}

// BuildBlock assembles a fully formed block.
func (s *Service[
	BeaconStateT,
	BlobSidecarsT,
]) SetBlockStateRoot(
	ctx context.Context, st BeaconStateT, blk types.BeaconBlock,
) error {
	// Compute the state root for the block.
	stateRoot, err := s.computeStateRoot(st, blk)
	if err != nil {
		return err
	}
	blk.SetStateRoot(stateRoot)
	return nil
}

// BuildBlockBody assembles a fully formed block body.
func (s *Service[
	BeaconStateT,
	BlobSidecarsT,
]) BuildSidecars(
	blk types.BeaconBlock, blobsBundle engineprimitives.BlobsBundle,
) (BlobSidecarsT, error) {
	return s.blobFactory.BuildSidecars(blk, blobsBundle)
}

func (s *Service[
	BeaconStateT,
	BlobSidecarsT,
]) RetrievePayload(
	ctx context.Context, st BeaconStateT, blk types.BeaconBlock,
) (engineprimitives.BuiltExecutionPayloadEnv, error) {
	// The latest execution payload header, will be from the previous block
	// during the block building phase.
	parentExecutionPayload, err := st.GetLatestExecutionPayloadHeader()
	if err != nil {
		return nil, err
	}

	// Get the payload for the block.
	envelope, err := s.localBuilder.RetrieveOrBuildPayload(
		ctx,
		st,
		blk.GetSlot(),
		blk.GetParentBlockRoot(),
		parentExecutionPayload.GetBlockHash(),
	)
	if err != nil {
		return nil, err
	} else if envelope == nil {
		return nil, ErrNilPayload
	}
	return envelope, nil
}
