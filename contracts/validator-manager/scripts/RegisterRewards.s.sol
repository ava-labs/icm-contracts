// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.25;

import {Script} from "forge-std/Script.sol";
import {ProxyAdmin} from "@openzeppelin/contracts@5.0.2/proxy/transparent/ProxyAdmin.sol";
import {Native721TokenStakingManager} from "../Native721TokenStakingManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {console} from "forge-std/console.sol";
import {StakingManagerSettings} from "../Native721TokenStakingManager.sol";
import {ValidatorManager} from "../ValidatorManager.sol";
import {IERC721} from "@openzeppelin/contracts@5.0.2/token/ERC721/IERC721.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts@5.0.2/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";

contract UpgradeBEAMStakingManager is Script {
    function run() external {
        vm.startBroadcast();
        
        Native721TokenStakingManager stakingManager = Native721TokenStakingManager(0xF4B5869AabE19a106C0df25E1537d855b54EEcBD);

        // ExampleERC20 rewardToken = new ExampleERC20();
        ExampleERC20 rewardToken = ExampleERC20(0xFE42fC7Ac06ad13BD54aA3010E2f5e9e0DF6D752);
        rewardToken.approve(address(stakingManager), 4000e18);
        
        stakingManager.registerRewards(true, 10084, address(rewardToken), 2000e18);
        stakingManager.registerRewards(false, 10084, address(rewardToken), 2000e18);
        
        vm.stopBroadcast();
    }
}
