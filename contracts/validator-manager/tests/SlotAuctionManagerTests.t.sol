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

abstract contract SlotAuctionManagerTest is Test {
    SlotAuctionManager public slotauctionmanager;

    ValidatorManager public validatorManager;

    PChainOwner public DEFAULT_P_CHAIN_OWNER;
    bytes32 public constant DEFAULT_SUBNET_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes public constant DEFAULT_NODE_ID = bytes(hex"1234123412341234123412341234123412341234");
    bytes public constant DEFAULT_INITIAL_VALIDATOR_NODE_ID_1 =
        bytes(hex"2341234123412341234123412341234123412341");
    bytes public constant DEFAULT_INITIAL_VALIDATOR_NODE_ID_2 =
        bytes(hex"3412341234123412341234123412341234123412");
    bytes public constant DEFAULT_BIDDING_VALIDATOR_NODE_ID_1 =
        bytes(hex"4123412341234123412341234123412341234123");
    bytes public constant DEFAULT_BIDDING_VALIDATOR_NODE_ID_2 =
        bytes(hex"2345234523452345234523452345234523452345");
    bytes public constant DEFAULT_BIDDING_VALIDATOR_NODE_ID_3 =
        bytes(hex"3452345234523452345234523452345234523452");
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
    // uint256 public constant DEFAULT_MINIMUM_STAKE_AMOUNT = 20e12;
    // uint256 public constant DEFAULT_MAXIMUM_STAKE_AMOUNT = 1e22;
    uint64 public constant DEFAULT_CHURN_PERIOD = 1 hours;
    uint8 public constant DEFAULT_MAXIMUM_CHURN_PERCENTAGE = 20;
    // uint8 public constant DEFAULT_MAXIMUM_HOURLY_CHURN = 0;
    uint64 public constant DEFAULT_REGISTRATION_TIMESTAMP = 1000;
    // uint256 public constant DEFAULT_STARTING_TOTAL_WEIGHT = 1e10 + DEFAULT_WEIGHT;
    uint64 public constant DEFAULT_MINIMUM_VALIDATION_DURATION = 24 hours;
    // uint64 public constant DEFAULT_COMPLETION_TIMESTAMP = 100_000;
    uint64 public constant REGISTRATION_EXPIRY_LENGTH = 1 days;
    uint16 public constant DEFAULT_VALIDATOR_SLOTS = 3;
    uint64 public constant DEFAULT_VALIDATOR_SLOT_WEIGHT = 10;
    uint256 public constant DEFAULT_MINIMUM_AUCTION_DURATION = 60 seconds;
    uint256 public constant DEFAULT_AUCTION_START_TIME = 1000 seconds;
    uint256 public constant DEFAULT_MINIMUM_BID = 10;

    event NewValidatorAuction(
        uint16 validatorSlots,
        uint64 validatorWeight, 
        uint256 minValidatorDuration, 
        uint256 auctionEndTime,
        uint256 minimumBid
    );

    event SuccessfulBidPlaced(
        uint256 indexed bid,
        bytes indexed nodeID
    );
    
    event BidEvicted(
        uint256 indexed bid, 
        bytes indexed nodeID
    );

    event InitiatedAuctionValidatorRegistration(
        bytes32 indexed validationID, 
        address indexed ownerAddress, 
        uint256 validatorEndTime,
        uint64 weight
    );

    function setUp() public virtual {
        // This stays 
        address[] memory addresses = new address[](1);
        addresses[0] = 0x1234567812345678123456781234567812345678;
        DEFAULT_P_CHAIN_OWNER = PChainOwner({threshold: 1, addresses: addresses});
        // no matter what
    }

    function testStartAuction() public {
        vm.expectEmit(true, true, true, true, address(slotauctionmanager));
        emit NewValidatorAuction(
            DEFAULT_VALIDATOR_SLOTS, 
            DEFAULT_VALIDATOR_SLOT_WEIGHT, 
            DEFAULT_MINIMUM_VALIDATION_DURATION, 
            DEFAULT_MINIMUM_AUCTION_DURATION + block.timestamp, 
            DEFAULT_MINIMUM_BID
        );
        slotauctionmanager.initiateAuction();
    }

    function testAuctionAlreadyRunning() public {
        _startAuctionWithCheck();

        vm.expectRevert(
            abi.encodeWithSelector(SlotAuctionManager.AuctionInProgress.selector)
        );
        slotauctionmanager.initiateAuction();
    }

    function testEndAuctionWithoutAuctionRunning() public {
        vm.expectRevert(
            abi.encodeWithSelector(SlotAuctionManager.AuctionNotInProgress.selector)
        );
        slotauctionmanager.endAuction();
    }

    function testBidBelowMinimum() public {
        _startAuctionWithCheck();

        uint256 belowMinimumBid = DEFAULT_MINIMUM_BID - 1;
        _beforeBid(belowMinimumBid, address(this));
        vm.expectRevert(
            abi.encodeWithSelector(
                SlotAuctionManager.BidSmallerThanMinimum.selector,
                DEFAULT_MINIMUM_BID,
                belowMinimumBid
            )
        );
        _placeBid(
            belowMinimumBid, 
            DEFAULT_BIDDING_VALIDATOR_NODE_ID_1, 
            DEFAULT_BLS_PUBLIC_KEY, 
            DEFAULT_P_CHAIN_OWNER, 
            DEFAULT_P_CHAIN_OWNER
        );
    }

    function testDuplicateBid() public {
        _startAuctionWithCheck();
        _FundAndPlaceBidWithCheck(
            DEFAULT_MINIMUM_BID,
            DEFAULT_BIDDING_VALIDATOR_NODE_ID_1,
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_P_CHAIN_OWNER,
            DEFAULT_P_CHAIN_OWNER
        );
        vm.expectRevert(
            abi.encodeWithSelector(
                SlotAuctionManager.DuplicateBidInContention.selector, 
                DEFAULT_MINIMUM_BID
            )
        );
        _placeBid(
            DEFAULT_MINIMUM_BID, 
            DEFAULT_BIDDING_VALIDATOR_NODE_ID_2, 
            DEFAULT_BLS_PUBLIC_KEY, 
            DEFAULT_P_CHAIN_OWNER, 
            DEFAULT_P_CHAIN_OWNER
        );
    }

    function testBidWithDuplicateNodeID() public {
        _startAuctionWithCheck();
        _FundAndPlaceBidWithCheck(
            DEFAULT_MINIMUM_BID,
            DEFAULT_BIDDING_VALIDATOR_NODE_ID_1,
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_P_CHAIN_OWNER,
            DEFAULT_P_CHAIN_OWNER
        );
        vm.expectRevert(
            abi.encodeWithSelector(
                SlotAuctionManager.DuplicateNodeIDInContention.selector, 
                DEFAULT_BIDDING_VALIDATOR_NODE_ID_1
            )
        );
        _placeBid(
            DEFAULT_MINIMUM_BID, 
            DEFAULT_BIDDING_VALIDATOR_NODE_ID_1, 
            DEFAULT_BLS_PUBLIC_KEY, 
            DEFAULT_P_CHAIN_OWNER, 
            DEFAULT_P_CHAIN_OWNER
        );
    }

    function testSetSlotAuctionSettings() public {
        // Making sure default settings are set properly
        vm.assertEq(slotauctionmanager.totalValidatorSlots(), DEFAULT_VALIDATOR_SLOTS);
        vm.assertEq(slotauctionmanager.validatorWeight(), DEFAULT_VALIDATOR_SLOT_WEIGHT);
        vm.assertEq(slotauctionmanager.MinAuctionDuration(), DEFAULT_MINIMUM_AUCTION_DURATION);
        vm.assertEq(slotauctionmanager.MinValidatorDuration(), DEFAULT_MINIMUM_VALIDATION_DURATION);
        vm.assertEq(slotauctionmanager.MinimumBid(), DEFAULT_MINIMUM_BID);

        slotauctionmanager.setSlotAuctionSettings(
            DEFAULT_VALIDATOR_SLOTS + 1, 
            DEFAULT_VALIDATOR_SLOT_WEIGHT + 1, 
            DEFAULT_MINIMUM_AUCTION_DURATION + 1, 
            DEFAULT_MINIMUM_VALIDATION_DURATION + 1, 
            DEFAULT_MINIMUM_BID + 1
        );
        vm.assertEq(slotauctionmanager.totalValidatorSlots(), DEFAULT_VALIDATOR_SLOTS + 1);
        vm.assertEq(slotauctionmanager.validatorWeight(), DEFAULT_VALIDATOR_SLOT_WEIGHT + 1);
        vm.assertEq(slotauctionmanager.MinAuctionDuration(), DEFAULT_MINIMUM_AUCTION_DURATION + 1);
        vm.assertEq(slotauctionmanager.MinValidatorDuration(), DEFAULT_MINIMUM_VALIDATION_DURATION + 1);
        vm.assertEq(slotauctionmanager.MinimumBid(), DEFAULT_MINIMUM_BID + 1);
    }

    function testSetSlotAuctionSettingsDuringAuction() public {
        _startAuctionWithCheck();
        vm.expectRevert(
            abi.encodeWithSelector(
                SlotAuctionManager.AuctionInProgress.selector
            )
        );
        slotauctionmanager.setSlotAuctionSettings(
            DEFAULT_VALIDATOR_SLOTS + 1, 
            DEFAULT_VALIDATOR_SLOT_WEIGHT + 1, 
            DEFAULT_MINIMUM_AUCTION_DURATION + 1, 
            DEFAULT_MINIMUM_VALIDATION_DURATION + 1, 
            DEFAULT_MINIMUM_BID + 1
        );    
    }

    function testBidEvicted() public { //TODO: NOT WORKING RN, IDK IF ITS MY TEST CASE OR THE CONTRACT 
        // Setting slots to 1
        slotauctionmanager.setSlotAuctionSettings(
            1, 
            DEFAULT_VALIDATOR_SLOT_WEIGHT, 
            DEFAULT_MINIMUM_AUCTION_DURATION, 
            DEFAULT_MINIMUM_VALIDATION_DURATION, 
            DEFAULT_MINIMUM_BID
        );
        _startAuctionWithCheck();
        _FundAndPlaceBidWithCheck(
            DEFAULT_MINIMUM_BID,
            DEFAULT_BIDDING_VALIDATOR_NODE_ID_1,
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_P_CHAIN_OWNER,
            DEFAULT_P_CHAIN_OWNER
        );
        _beforeBid(DEFAULT_MINIMUM_BID + 1, address(this));
        vm.expectEmit(address(slotauctionmanager));
        emit BidEvicted(DEFAULT_MINIMUM_BID, DEFAULT_BIDDING_VALIDATOR_NODE_ID_1);
        _placeBid(
            DEFAULT_MINIMUM_BID + 1,
            DEFAULT_BIDDING_VALIDATOR_NODE_ID_2,
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_P_CHAIN_OWNER,
            DEFAULT_P_CHAIN_OWNER
        );
    }

    function _FundAndPlaceBidWithCheck(
        uint256 bid, 
        bytes memory nodeID, 
        bytes memory blsPublicKey, 
        PChainOwner memory remainingBalanceOwner, 
        PChainOwner memory disableOwner 
    ) internal {
        _beforeBid(bid, address(this));
        vm.expectEmit();
        emit SuccessfulBidPlaced(bid, nodeID);
        _placeBid(
            bid, 
            nodeID, 
            blsPublicKey, 
            remainingBalanceOwner, 
            disableOwner
        );
    }

    function _startAuctionWithCheck() internal {
        vm.expectEmit(true, true, true, true, address(slotauctionmanager));
        emit NewValidatorAuction(
            DEFAULT_VALIDATOR_SLOTS, 
            DEFAULT_VALIDATOR_SLOT_WEIGHT, 
            DEFAULT_MINIMUM_VALIDATION_DURATION, 
            DEFAULT_MINIMUM_AUCTION_DURATION + block.timestamp, 
            DEFAULT_MINIMUM_BID
        );
        slotauctionmanager.initiateAuction();
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

    function _mockGetPChainWarpMessage(
        bytes memory expectedPayload, 
        bool valid
    ) internal {
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
    function _setUp() internal virtual;

    function _beforeBid(
        uint256 amount, 
        address spender
    ) internal virtual;

    function _placeBid(
        uint256 bid, 
        bytes memory nodeID, 
        bytes memory blsPublicKey, 
        PChainOwner memory remainingBalanceOwner, 
        PChainOwner memory disableOwner 
    ) internal virtual;
    
}