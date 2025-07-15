// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";

// import {Test} from "@forge-std/Test.sol";

import {IACP99Manager, PChainOwner, ConversionData} from "../interfaces/IACP99Manager.sol";

import {IValidatorManager, ValidatorManager, ValidatorStatus} from "../ValidatorManager.sol";

import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

import {ValidatorMessages} from "../ValidatorMessages.sol";

import {ICMInitializable} from "../../utilities/ICMInitializable.sol";

import {SlotAuctionManager} from "../SlotAuctionManager.sol";

import {ISlotAuctionManager, ValidatorBid, ValidatorInfo} from "../interfaces/ISlotAuctionManager.sol";

import {ExampleERC20} from "@mocks/ExampleERC20.sol";

import {IERC20Mintable} from "../interfaces/IERC20Mintable.sol";

import { WarpMessage, IWarpMessenger } from "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/IWarpMessenger.sol";

contract SlotAuctionManagerTest { // is ValidatorManagerTest
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