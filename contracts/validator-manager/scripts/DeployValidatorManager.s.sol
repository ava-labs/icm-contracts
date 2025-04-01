// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.25;

import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {Script} from "forge-std/Script.sol";
import {
    StakingManagerSettings,
    Native721TokenStakingManager
} from "../Native721TokenStakingManager.sol";
import {ValidatorManager} from "../ValidatorManager.sol";
import {IERC721} from "@openzeppelin/contracts@5.0.2/token/ERC721/IERC721.sol";
import {console} from "forge-std/console.sol";
import {ProxyAdmin} from "@openzeppelin/contracts@5.0.2/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts@5.0.2/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts@5.0.2/proxy/transparent/TransparentUpgradeableProxy.sol";

import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {Script} from "forge-std/Script.sol";
import {
    ValidatorManagerSettings, ValidatorManager
} from "../ValidatorManager.sol";
import {console} from "forge-std/console.sol";


/**
 * @notice Script to deploy and initialize ValidatorManager with a new proxy
 * 
 * @dev To run this script:
 * 1. Update the initialization parameters below
 * 2. Run the script with forge:
 *    ```bash
 *    # Dry run (simulation)
 *    forge script contracts/validator-manager/scripts/DeployValidatorManager.s.sol --slow --optimize --optimizer-runs 200 -vvv --rpc-url <your-rpc-url> --private-key <your-private-key>
 *
 *    # Live run
 *    forge script contracts/validator-manager/scripts/DeployValidatorManager.s.sol --slow --optimize --optimizer-runs 200 -vvv --broadcast --verify --rpc-url <your-rpc-url>  --private-key <your-private-key>
 *    ```
 */
contract DeployValidatorManager is Script {
    // Example initialization parameters - adjust as needed
    // address private constant _ADMIN_ADDRESS = address(0xd68F802fD0B6f56524F379805DD8FcC152DB9d5c); // Replace with admin address
    // bytes32 private constant _SUBNET_ID = bytes32(hex"55437b606820ab92499f5f44a72d9a17c777eec4e103fec17d48bb0283197cf2"); // mainnet
    // uint64 private constant _CHURN_PERIOD = 1 hours;
    // uint8 private constant _MAX_CHURN_PERCENTAGE = 20;

    address private constant _ADMIN_ADDRESS = address(0xd68F802fD0B6f56524F379805DD8FcC152DB9d5c); // Replace with admin address
    bytes32 private constant _SUBNET_ID = bytes32(hex"5e8b6e2e8155e93739f2fa6a7f8a32c6bb2e1dce2e471b56dcc60aac49bf3435"); // mainnet
    uint64 private constant _CHURN_PERIOD = 1 hours;
    uint8 private constant _MAX_CHURN_PERCENTAGE = 20;

    function run() external {
        // Start broadcasting transactions
        vm.startBroadcast();

        // Deploy new implementation
        ValidatorManager implementation = new ValidatorManager(ICMInitializable.Disallowed);
        console.log("ValidatorManager implementation deployed at:", address(implementation));
        
        // Prepare initialization data
        ValidatorManagerSettings memory settings = ValidatorManagerSettings({
            admin: _ADMIN_ADDRESS,
            subnetID: _SUBNET_ID,
            churnPeriodSeconds: _CHURN_PERIOD,
            maximumChurnPercentage: _MAX_CHURN_PERCENTAGE
        });

        // Encode the initialization call
        bytes memory initData = abi.encodeWithSelector(
            ValidatorManager.initialize.selector,
            settings
        );

        // Deploy proxy with initialization
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(implementation),
            msg.sender,
            initData
        );
        console.log("Deployed and initialized proxy at:", address(proxy));

        bytes32 ADMIN_SLOT = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        address proxyAdmin = address(uint160(uint256(vm.load(address(proxy), ADMIN_SLOT))));
        console.log("ProxyAdmin contract deployed at:", proxyAdmin);

        vm.stopBroadcast();
    }
}
