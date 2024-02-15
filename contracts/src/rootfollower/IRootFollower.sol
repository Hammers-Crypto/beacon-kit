// SPDX-License-Identifier: MIT

pragma solidity >=0.8.10;

/**
 * @title IRootFollower
 * @dev The interface for an abstract follower of the beacon root contract.
 * @author Berachain
 */
interface IRootFollower {
    /**
     * @dev Emitted when the block is advanced.
     * @param blockNum The block number of the block just actioned upon. */
    event AdvancedBlock(uint256 blockNum);

    /**
     * @dev Emitted when the block count is skipped.
     * @param start The start block number of the block just actioned upon.
     * @param end The end block number of the block just actioned upon. */
    event SkipBlocks(uint256 start, uint256 end);

    /**
     * @dev Gets the address of the coinbase for the given block number. The size of
     * the BeaconRootsContract stores the coinbase for the last 256 blocks. Querying
     * a block number greater than the last 256 blocks will return an error. This also
     * implies that actions must be invoked within 256 blocks of being proposed.
     * @param blockNum The address performing the mint.
     * @return coinbase The address of the coinbase for the given block number. */
    function getCoinbase(
        uint256 blockNum
    ) external view returns (address coinbase);

    /**
     * @dev Gets the next block to be rewarded. This returns the greater of current
     * previously invoked block + 1, or current block number - 256 as that is the
     * limitation on number of blocks that can be queried, and actioned upon.
     * @return blockNum The block number of the next block to be invoked. */
    function getNextActionableBlock() external view returns (uint256 blockNum);

    /**
     * @dev Gets the last block that was actioned upon.
     */
    function getLastActionedBlock() external view returns (uint256 blockNum);
}
