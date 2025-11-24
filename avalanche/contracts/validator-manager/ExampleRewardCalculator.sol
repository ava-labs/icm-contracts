// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IRewardCalculator} from "./interfaces/IRewardCalculator.sol";

contract ExampleRewardCalculator is IRewardCalculator {
    uint256 public constant SECONDS_IN_YEAR = 31536000;

    uint8 public constant UPTIME_REWARDS_THRESHOLD_PERCENTAGE = 80;

    uint16 public constant BIPS_CONVERSION_FACTOR = 10000;

    uint64 public immutable rewardBasisPoints;

    constructor(
        uint64 rewardBasisPoints_
    ) {
        rewardBasisPoints = rewardBasisPoints_;
    }

    /**
     * @notice A linear, non-compounding reward calculation that rewards a set percentage of tokens per year.
     * See {IRewardCalculator-calculateReward}
     */
    function calculateReward(
        uint256 stakeAmount,
        uint64 validatorStartTime,
        uint64 stakingStartTime,
        uint64 stakingEndTime,
        uint64 uptimeSeconds
    ) external view returns (uint256) {
        // Equivalent to uptimeSeconds/(validator.endTime - validator.startTime) < UPTIME_REWARDS_THRESHOLD_PERCENTAGE/100
        // Rearranged to prevent integer division truncation.
        if (
            uptimeSeconds * 100
                < (stakingEndTime - validatorStartTime) * UPTIME_REWARDS_THRESHOLD_PERCENTAGE
        ) {
            return 0;
        }

        return (stakeAmount * rewardBasisPoints * (stakingEndTime - stakingStartTime))
            / SECONDS_IN_YEAR / BIPS_CONVERSION_FACTOR;
    }
}
