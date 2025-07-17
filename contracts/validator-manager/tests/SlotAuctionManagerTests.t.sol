// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {IACP99Manager, PChainOwner, ConversionData, InitialValidator} from "../interfaces/IACP99Manager.sol";
import {IValidatorManager, ValidatorManager, ValidatorStatus, ValidatorManagerSettings} from "../ValidatorManager.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";
import {ERC20TokenSlotAuctionManager} from "../ERC20TokenSlotAuctionManager.sol";
import {SlotAuctionManager} from "../SlotAuctionManager.sol";
import {ISlotAuctionManager, ValidatorBid, ValidatorInfo} from "../interfaces/ISlotAuctionManager.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {IERC20Mintable} from "../interfaces/IERC20Mintable.sol";
import { WarpMessage, IWarpMessenger } from "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/IWarpMessenger.sol";
import {Test} from "@forge-std/Test.sol";

contract SlotAuctionManagerTest is Test {
    using SafeERC20 for IERC20Mintable;

    ERC20TokenSlotAuctionManager public app;
    IERC20Mintable public token;
    ValidatorManager public validatorManager;

    PChainOwner public DEFAULT_P_CHAIN_OWNER;
    bytes32 public constant DEFAULT_SUBNET_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes public constant DEFAULT_NODE_ID = bytes(hex"1234123412341234123412341234123412341234");
    bytes public constant DEFAULT_INITIAL_VALIDATOR_NODE_ID_1 =
        bytes(hex"2341234123412341234123412341234123412341");
    bytes public constant DEFAULT_INITIAL_VALIDATOR_NODE_ID_2 =
        bytes(hex"3412341234123412341234123412341234123412");
    bytes public constant DEFAULT_BLS_PUBLIC_KEY = bytes(
        hex"123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
    );
    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_SUBNET_CONVERSION_ID =
        bytes32(hex"67e8531265d8e97bd5c23534a37f4ea42d41934ddf8fe2c77c27fac9ef89f973");
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;

    uint64 public constant DEFAULT_WEIGHT = 1e6;
    // Set the default weight to 1e10 to avoid churn issues
    uint64 public constant DEFAULT_INITIAL_VALIDATOR_WEIGHT = DEFAULT_WEIGHT * 1e4;
    uint64 public constant DEFAULT_INITIAL_TOTAL_WEIGHT =
        DEFAULT_INITIAL_VALIDATOR_WEIGHT + DEFAULT_WEIGHT;
    uint256 public constant DEFAULT_MINIMUM_STAKE_AMOUNT = 20e12;
    uint256 public constant DEFAULT_MAXIMUM_STAKE_AMOUNT = 1e22;
    uint64 public constant DEFAULT_CHURN_PERIOD = 1 hours;
    uint8 public constant DEFAULT_MAXIMUM_CHURN_PERCENTAGE = 20;
    uint8 public constant DEFAULT_MAXIMUM_HOURLY_CHURN = 0;
    uint64 public constant DEFAULT_REGISTRATION_TIMESTAMP = 1000;
    uint256 public constant DEFAULT_STARTING_TOTAL_WEIGHT = 1e10 + DEFAULT_WEIGHT;
    uint64 public constant DEFAULT_MINIMUM_VALIDATION_DURATION = 24 hours;
    uint64 public constant DEFAULT_COMPLETION_TIMESTAMP = 100_000;
    uint64 public constant REGISTRATION_EXPIRY_LENGTH = 1 days;

    function testSilly() public {
         app.initiateAuction(1, 1, 1, 1, 1);
    }

    function setUp() public {
        address[] memory addresses = new address[](1);
        addresses[0] = 0x1234567812345678123456781234567812345678;
        DEFAULT_P_CHAIN_OWNER = PChainOwner({threshold: 1, addresses: addresses});

        token = new ExampleERC20();
        validatorManager = new ValidatorManager(ICMInitializable.Allowed);
        app = new ERC20TokenSlotAuctionManager(address(validatorManager), address(token));

        validatorManager.initialize(_defaultSettings(address(app)));

        _mockGetBlockchainID();

        ConversionData memory conversion = _defaultConversionData();
        bytes32 conversionID = sha256(ValidatorMessages.packConversionData(conversion));
        _mockInitializeValidatorSet(conversionID);
        validatorManager.initializeValidatorSet(conversion, 0);
    }

    function _defaultSettings(
        address admin
    ) internal pure returns (ValidatorManagerSettings memory) {
        return ValidatorManagerSettings({
            admin: admin,
            subnetID: DEFAULT_SUBNET_ID,
            churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
            maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE
        });
    }

    function _mockGetPChainWarpMessage(bytes memory expectedPayload, bool valid) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: validatorManager.P_CHAIN_BLOCKCHAIN_ID(),
                    originSenderAddress: address(0),
                    payload: expectedPayload
                }),
                valid
            )
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, 0)
        );
    }

    function _mockInitializeValidatorSet(
        bytes32 conversionID
    ) internal {
        _mockGetPChainWarpMessage(
            ValidatorMessages.packSubnetToL1ConversionMessage(conversionID), true
        );
    }


    function _mockGetBlockchainID() internal {
        _mockGetBlockchainID(DEFAULT_SOURCE_BLOCKCHAIN_ID);
    }

    function _mockGetBlockchainID(
        bytes32 blockchainID
    ) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(blockchainID)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );
    }

    function _defaultConversionData() internal view returns (ConversionData memory) {
        InitialValidator[] memory initialValidators = new InitialValidator[](2);
        // The first initial validator has a high weight relative to the default PoS validator weight
        // to avoid churn issues
        initialValidators[0] = InitialValidator({
            nodeID: DEFAULT_INITIAL_VALIDATOR_NODE_ID_1,
            weight: DEFAULT_INITIAL_VALIDATOR_WEIGHT,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
        });
        // The second initial validator has a low weight so that it can be safely removed in tests
        initialValidators[1] = InitialValidator({
            nodeID: DEFAULT_INITIAL_VALIDATOR_NODE_ID_2,
            weight: DEFAULT_WEIGHT,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
        });

        // Confirm the total initial weight
        uint64 initialWeight;
        for (uint256 i = 0; i < initialValidators.length; i++) {
            initialWeight += initialValidators[i].weight;
        }
        assertEq(initialWeight, DEFAULT_INITIAL_TOTAL_WEIGHT);

        return ConversionData({
            subnetID: DEFAULT_SUBNET_ID,
            validatorManagerBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            validatorManagerAddress: address(validatorManager),
            initialValidators: initialValidators
        });
    }


    // using SafeERC20 for IERC20Mintable;

    // SlotAuctionManager public slotAuctionManager;
    // IERC20Mintable public token;
    
    // uint16 public constant DEFAULT_VALIDATOR_SLOTS = uint16(2);
    // uint256 public constant DEFAULT_MINIMUM_AUCTION_DURATION = uint256(100);
    // uint256 public constant DEFAULT_MINIMUM_BID = uint256(10); 
    // uint256 public constant DEFAULT_USER_BID = uint256(10); 
    // uint256 public constant DEFAULT_FUND_AMOUNT = uint256(50);
    // address public constant DEFAULT_FIRST_BIDDER_ADDRESS = 0x1847184718471847184718471847184718471847;
    // address public constant DEFAULT_SECOND_BIDDER_ADDRESS = 0x1963196319631963196319631963196319631963;
    // address public constant DEFAULT_THIRD_BIDDER_ADDRESS = 0x2014201420142014201420142014201420142014;

    // event InitiatedAuctionValidatorRegistration(
    //     bytes32 indexed validationID,
    //     address indexed owner,
    //     uint64 minStakeDuration,
    //     uint64 validatorWeight
    // );

    // function setUp() public override {
    //     ValidatorManagerTest.setUp();

    //     _setUp();
    //     _mockGetBlockchainID();

    //     ConversionData memory conversion = _defaultConversionData();
    //     bytes32 conversionID = sha256(ValidatorMessages.packConversionData(conversion));
    //     _mockInitializeValidatorSet(conversionID);
    //     validatorManager.initializeValidatorSet(conversion, 0);
    // }
    
    // function testBidWithoutAuction() public {
    //     vm.expectRevert(abi.encodeWithSelector(SlotAuctionManager.AuctionNotInProgress.selector));
    //     PChainOwner memory PCO;
    //     slotAuctionManager.placeBid(100, DEFAULT_NODE_ID, DEFAULT_BLS_PUBLIC_KEY, PCO, PCO);
    // }

    // function _beforeRegisterValidator(
    //     bytes32 validationID,
    //     address rewardRecipient
    // ) internal virtual override {
    //     // vm.expectEmit(true, true, true, true, address(slotAuctionManager));
    //     // emit InitiatedAuctionValidatorRegistration(
    //     //     validationID,
    //     //     address(this),
    //     //     DEFAULT_MINIMUM_VALIDATION_DURATION,
    //     //     DEFAULT_INITIAL_VALIDATOR_WEIGHT
    //     // );
    // }

    // function _beforeSend(uint256 amount, address spender) internal override {
    //     token.safeIncreaseAllowance(spender, amount);
    //     token.safeTransfer(spender, amount);

    //     // ERC20 tokens need to be pre-approved
    //     vm.startPrank(spender);
    //     token.safeIncreaseAllowance(address(slotAuctionManager), amount);
    //     vm.stopPrank();
    // }

    // function _completeValidatorRegistration(
    //     uint32 messageIndex
    // ) internal virtual override returns (bytes32) {
    //     return slotAuctionManager.completeValidatorRegistration(messageIndex);
    // }

    // function _completeValidatorRemoval(
    //     uint32 messageIndex
    // ) internal virtual override returns (bytes32) {
    //     return slotAuctionManager.completeValidatorRemoval(messageIndex);
    // }

    // function _initiateValidatorRemoval(
    //     bytes32 validationID,
    //     bool includeUptime
    // ) internal virtual override {
    //     slotAuctionManager.initiateValidatorRemoval(validationID);
    // }

    // function _initiateValidatorRegistration(
    //     bytes memory nodeID,
    //     bytes memory blsPublicKey,
    //     PChainOwner memory remainingBalanceOwner,
    //     PChainOwner memory disableOwner,
    //     uint64 weight,
    //     address rewardRecipient
    // ) override internal virtual returns (bytes32) {
    //     vm.warp(0);
    //     slotAuctionManager.initiateAuction(
    //         DEFAULT_VALIDATOR_SLOTS - 1, 
    //         DEFAULT_WEIGHT, 
    //         DEFAULT_MINIMUM_AUCTION_DURATION, 
    //         DEFAULT_MINIMUM_VALIDATION_DURATION, 
    //         DEFAULT_MINIMUM_BID
    //     );
    //     token.mint(address(this), DEFAULT_FUND_AMOUNT);
    //     slotAuctionManager.placeBid(DEFAULT_MINIMUM_BID, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner);
    //     vm.warp(DEFAULT_REGISTRATION_TIMESTAMP);
    //     slotAuctionManager.endAuction();

    //     ValidatorInfo memory newValidator = slotAuctionManager.getValidatorInfoByNodeID(nodeID);
    //     return newValidator.validationID;
    // }

    // function _initiateValidatorRegistration(
    //     bytes memory nodeID,
    //     bytes memory blsPublicKey,
    //     PChainOwner memory remainingBalanceOwner,
    //     PChainOwner memory disableOwner,
    //     uint64 weight
    // ) override internal virtual returns (bytes32) {
    //     slotAuctionManager.initiateAuction(
    //         DEFAULT_VALIDATOR_SLOTS - 1, 
    //         DEFAULT_WEIGHT, 
    //         DEFAULT_MINIMUM_AUCTION_DURATION, 
    //         DEFAULT_MINIMUM_VALIDATION_DURATION, 
    //         DEFAULT_MINIMUM_BID
    //     );
    //     token.mint(address(msg.sender), DEFAULT_FUND_AMOUNT);
    //     slotAuctionManager.placeBid(5, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner);
    //     slotAuctionManager.endAuction();

    //     ValidatorInfo memory newValidator = slotAuctionManager.getValidatorInfoByNodeID(nodeID);
    //     return newValidator.validationID;
    // }

    // function _forceInitiateValidatorRemoval(
    //     bytes32 validationID,
    //     bool includeUptime
    // ) internal virtual override {
    //     return slotAuctionManager.initiateValidatorRemoval(validationID);
    // }

    // function _setUp() internal override returns (IACP99Manager) {
    //     // Construct the object under test
    //     token = new ExampleERC20();
    //     validatorManager = new ValidatorManager(ICMInitializable.Allowed);
    //     slotAuctionManager = new SlotAuctionManager(address(token), address(validatorManager));

    //     validatorManager.initialize(_defaultSettings(address(slotAuctionManager)));

    //     return validatorManager;
    // }
    
}