// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.25;

import {Script} from "forge-std/Script.sol";
import {Native721TokenStakingManager} from "../Native721TokenStakingManager.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";

contract RegisterRewards is Script {
    function run() external {
        vm.startBroadcast();
        
        Native721TokenStakingManager stakingManager = Native721TokenStakingManager(0xF4B5869AabE19a106C0df25E1537d855b54EEcBD);

        // ExampleERC20 rewardToken = new ExampleERC20();
        ExampleERC20 rewardToken = ExampleERC20(0xFE42fC7Ac06ad13BD54aA3010E2f5e9e0DF6D752);
        rewardToken.approve(address(stakingManager), 10000e18);
        
        stakingManager.registerRewards(true, 10091, address(rewardToken), 1000e18);
        stakingManager.registerRewards(false, 10091, address(rewardToken), 1000e18);
        
        vm.stopBroadcast();
    }
}
