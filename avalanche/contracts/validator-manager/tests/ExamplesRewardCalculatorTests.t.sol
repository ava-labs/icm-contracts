// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {ExampleRewardCalculator} from "../ExampleRewardCalculator.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";

contract ExampleRewardCalculatorTest is Test {
    uint256 public constant DEFAULT_STAKE_AMOUNT = 1e12;
    uint64 public constant DEFAULT_START_TIME = 1000;
    uint64 public constant DEFAULT_END_TIME = 31537000; // a year + 1000 seonds
    uint64 public constant DEFAULT_UPTIME = (DEFAULT_END_TIME - DEFAULT_START_TIME) * 80 / 100;
    uint64 public constant DEFAULT_REWARD_BASIS_POINTS = 42;

    IRewardCalculator public exampleRewardCalculator;

    function setUp() public {
        exampleRewardCalculator = new ExampleRewardCalculator(DEFAULT_REWARD_BASIS_POINTS);
    }

    function testRewardCalculation() public view {
        uint256 output = exampleRewardCalculator.calculateReward({
            stakeAmount: DEFAULT_STAKE_AMOUNT,
            validatorStartTime: DEFAULT_START_TIME,
            stakingStartTime: DEFAULT_START_TIME,
            stakingEndTime: DEFAULT_END_TIME,
            uptimeSeconds: DEFAULT_UPTIME
        });
        assertEq(output, 42e8);
    }

    function testInsufficientUptime() public view {
        uint256 output = exampleRewardCalculator.calculateReward({
            stakeAmount: DEFAULT_STAKE_AMOUNT,
            validatorStartTime: DEFAULT_START_TIME,
            stakingStartTime: DEFAULT_START_TIME,
            stakingEndTime: DEFAULT_END_TIME,
            uptimeSeconds: DEFAULT_UPTIME - 1
        });
        assertEq(output, 0);
    }
}
