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

/**
 * @notice Script to upgrade the Native721TokenStakingManager implementation
 * 
 * @dev To run this script:
 * 1. Make sure the Native721TokenStakingManager contract has reinitializer(3) or higher
 *    (update from reinitializer(2) if needed)
 * 2. Update the proxy and proxy admin addresses below
 * 3. Run the script with forge:
 *    ```bash
 *    # Dry run (simulation)
 *    forge script contracts/validator-manager/scripts/UpgradeBEAMStakingManager.s.sol --slow --optimize --optimizer-runs 200 -vvv --rpc-url <your-rpc-url> --private-key <your-private-key>
 *
 *    # Live run
 *    forge script contracts/validator-manager/scripts/UpgradeBEAMStakingManager.s.sol --slow --optimize --optimizer-runs 200 -vvv --broadcast --verify --rpc-url <your-rpc-url> --private-key <your-private-key>
 *    ```
 */
contract UpgradeBEAMStakingManager is Script {
    // Update these addresses with your deployed contract addresses
    address constant _PROXY_ADDRESS = address(0xF4B5869AabE19a106C0df25E1537d855b54EEcBD);
    address constant _PROXY_ADMIN_ADDRESS = address(0x4CDd1785908756dc515aFc766E3e3A9630761fa1);

    // Add necessary constants
    address constant NFT_TOKEN_ADDRESS = address(0x732080D7aD6A9C50039d7Ad7F5BD0a79670f7654);
    address constant ADMIN_ADDRESS = address(0xd68F802fD0B6f56524F379805DD8FcC152DB9d5c);
    address constant VALIDATOR_MANAGER_ADDRESS = address(0x33B9785E20ec582d5009965FB3346F1716e8A423);
    uint64 constant MINIMUM_STAKE_DURATION = 1 hours;
    uint256 constant MINIMUM_STAKE_AMOUNT = 20_000e18;
    uint256 constant MAXIMUM_STAKE_AMOUNT = 200_000_000e18;
    uint64 constant UNLOCK_PERIOD = 1 hours;
    uint16 constant MINIMUM_DELEGATION_FEE = 100; // 0.1% in basis points
    uint64 constant EPOCH_DURATION = 2 days;
    uint256 constant MAXIMUM_NFT_AMOUNT = 1000;
    uint256 constant MINIMUM_DELEGATION_AMOUNT = 100e18;
    uint256 constant WEIGHT_TO_VALUE_FACTOR = 1e18;
    bytes32 constant UPTIME_BLOCKCHAIN_ID = bytes32(hex"7f78fe8ca06cefa186ef29c15231e45e1056cd8319ceca0695ca61099e610355");

    function run() external {
        vm.startBroadcast();

        // Deploy new implementation
        Native721TokenStakingManager newImplementation = new Native721TokenStakingManager(
            ICMInitializable.Disallowed
        );
        console.log("Deployed new implementation at:", address(newImplementation));

        // Add settings struct for initialization
        StakingManagerSettings memory settings = StakingManagerSettings({
            manager: ValidatorManager(VALIDATOR_MANAGER_ADDRESS),
            minimumStakeAmount: MINIMUM_STAKE_AMOUNT,
            maximumStakeAmount: MAXIMUM_STAKE_AMOUNT,
            maximumNFTAmount: MAXIMUM_NFT_AMOUNT,
            minimumStakeDuration: MINIMUM_STAKE_DURATION,
            minimumDelegationAmount: MINIMUM_DELEGATION_AMOUNT,
            minimumDelegationFeeBips: MINIMUM_DELEGATION_FEE,
            weightToValueFactor: WEIGHT_TO_VALUE_FACTOR,
            admin: ADMIN_ADDRESS,
            uptimeBlockchainID: UPTIME_BLOCKCHAIN_ID,
            epochDuration: EPOCH_DURATION,
            unlockDuration: UNLOCK_PERIOD,
            uptimeKeeper: ADMIN_ADDRESS,
            epochOffset: 0
        });

       // Get ProxyAdmin instance
        ProxyAdmin proxyAdmin = ProxyAdmin(_PROXY_ADMIN_ADDRESS);
        
        // bytes memory selector;
        // Upgrade proxy to new implementation
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(_PROXY_ADDRESS),
            address(newImplementation),
            // selector
            abi.encodeWithSelector(Native721TokenStakingManager.initialize.selector, settings, address(NFT_TOKEN_ADDRESS))
        );
        console.log("Upgraded proxy to new implementation");

        vm.stopBroadcast();

        console.log("Upgrade verified successfully");
    }
}
