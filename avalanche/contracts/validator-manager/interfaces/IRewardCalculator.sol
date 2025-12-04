// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

/**
 * @notice Interface for Validation and Delegation reward calculators
 */
interface IRewardCalculator {
    /**
     * @notice Calculate the reward for a staker (validator or delegator)
     * @param stakeAmount The amount of tokens staked
     * @param validatorStartTime The time the validator started validating
     * @param stakingStartTime The time the staker started staking
     * @param stakingEndTime The time the staker stopped staking
     * @param uptimeSeconds The total time the validator was validating
     */
    function calculateReward(
        uint256 stakeAmount,
        uint64 validatorStartTime,
        uint64 stakingStartTime,
        uint64 stakingEndTime,
        uint64 uptimeSeconds
    ) external view returns (uint256);
}
