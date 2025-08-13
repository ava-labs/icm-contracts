// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {PChainOwner, ConversionData, InitialValidator} from "../../interfaces/IACP99Manager.sol";
import {ValidatorManagerSettings} from "../../ValidatorManager.sol";
import {ValidatorMessages} from "../../ValidatorMessages.sol";
import {SlotAuctionManager} from "../SlotAuctionManager.sol";
import {AuctionSettings} from "../interfaces/ISlotAuctionManager.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/IWarpMessenger.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {ValidatorManagerFunctionality} from "../../tests/ValidatorManagerFunctionality.t.sol";

abstract contract SlotAuctionManagerTest is ValidatorManagerFunctionality {
    // solhint-disable var-name-mixedcase

    address public constant DEFAULT_AUCTION_OWNER_ADRESS =
        0x9900000000000000000000000000099999999999;
    uint16 public constant DEFAULT_VALIDATOR_SLOTS = 3;
    uint64 public constant DEFAULT_VALIDATOR_SLOT_WEIGHT = 10;
    uint256 public constant DEFAULT_MINIMUM_AUCTION_DURATION = 60 seconds;
    uint256 public constant DEFAULT_MINIMUM_BID = 10;
    uint256 public constant DEFUALT_AUCTION_COOLDOWN_DURATION = 1 days;
    AuctionSettings public DEFAULT_AUCTION_SETTINGS = AuctionSettings({
        totalValidatorSlots: DEFAULT_VALIDATOR_SLOTS,
        weight: DEFAULT_VALIDATOR_SLOT_WEIGHT,
        minValidatorDuration: DEFAULT_MINIMUM_VALIDATION_DURATION,
        minAuctionDuration: DEFAULT_MINIMUM_AUCTION_DURATION,
        minimumBid: DEFAULT_MINIMUM_BID,
        auctionCooldownDuration: DEFUALT_AUCTION_COOLDOWN_DURATION
    });
    address[4] public DEFAULT_SENDER_ADDRESSES = [
        0x4321432143214321432143214321432143214321,
        0x3214321432143214321432143214321432143214,
        0x2143214321432143214321432143214321432143,
        0x1432143214321432143214321432143214321432
    ];
    bytes public constant DEFAULT_CHURN_UPDATE_NODEID = 
        bytes(hex"5634563456345634563456345634563456345634");
    bytes[4] public DEFAULT_BIDDING_VALIDATOR_NODEIDS = [
        bytes(hex"4123412341234123412341234123412341234123"),
        bytes(hex"2345234523452345234523452345234523452345"),
        bytes(hex"3452345234523452345234523452345234523452"),
        bytes(hex"4523452345234523452345234523452345234523")
    ];

    SlotAuctionManager public slotauctionmanager;
    // solhint-enable var-name-mixedcase

    event NewValidatorAuction(
        uint16 validatorSlots,
        uint64 validatorWeight,
        uint256 minValidatorDuration,
        uint256 auctionEndTime,
        uint256 minimumBid
    );

    event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID);

    event BidEvicted(uint256 indexed bid, bytes indexed nodeID);

    event InitiatedAuctionValidatorRegistration(
        bytes32 indexed validationID,
        address indexed ownerAddress,
        uint256 validatorEndTime,
        uint64 weight
    );

    event AuctionVoucherCreated(
        bytes indexed nodeID,
        address indexed ownerAddress,
        uint256 validatorEndTime,
        uint64 weight
    );

    function setUp() public virtual override {
        address[] memory addresses = new address[](1);
        addresses[0] = 0x1234567812345678123456781234567812345678;
        DEFAULT_P_CHAIN_OWNER = PChainOwner({threshold: 1, addresses: addresses});
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

        vm.expectRevert(abi.encodeWithSelector(SlotAuctionManager.AuctionInProgress.selector));
        slotauctionmanager.initiateAuction();
    }

    function testEndAuctionWithoutAuctionRunning() public {
        vm.expectRevert(abi.encodeWithSelector(SlotAuctionManager.AuctionNotInProgress.selector));
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
            DEFAULT_BIDDING_VALIDATOR_NODEIDS[0],
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_P_CHAIN_OWNER,
            DEFAULT_P_CHAIN_OWNER
        );
    }

    function testDuplicateBid() public {
        _startAuctionWithCheck();
        _fundAndPlaceBidWithCheck({
            sender: DEFAULT_SENDER_ADDRESSES[0],
            bid: DEFAULT_MINIMUM_BID,
            nodeID: DEFAULT_BIDDING_VALIDATOR_NODEIDS[0],
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER
        });
        vm.expectRevert(
            abi.encodeWithSelector(
                SlotAuctionManager.DuplicateBidInContention.selector, DEFAULT_MINIMUM_BID
            )
        );
        _placeBid(
            DEFAULT_MINIMUM_BID,
            DEFAULT_BIDDING_VALIDATOR_NODEIDS[1],
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_P_CHAIN_OWNER,
            DEFAULT_P_CHAIN_OWNER
        );
    }

    function testBidWithDuplicateNodeID() public {
        _startAuctionWithCheck();
        _fundAndPlaceBidWithCheck({
            sender: DEFAULT_SENDER_ADDRESSES[0],
            bid: DEFAULT_MINIMUM_BID,
            nodeID: DEFAULT_BIDDING_VALIDATOR_NODEIDS[0],
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER
        });
        vm.expectRevert(
            abi.encodeWithSelector(
                SlotAuctionManager.DuplicateNodeIDInContention.selector,
                DEFAULT_BIDDING_VALIDATOR_NODEIDS[0]
            )
        );
        _placeBid(
            DEFAULT_MINIMUM_BID,
            DEFAULT_BIDDING_VALIDATOR_NODEIDS[0],
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_P_CHAIN_OWNER,
            DEFAULT_P_CHAIN_OWNER
        );
    }

    // TODO: change magic number
    function testSingleAuctionValidatorRegistration() public {
        bytes32 validationID = _initiateRegisterValidator(
            DEFAULT_SENDER_ADDRESSES[0],
            DEFAULT_BIDDING_VALIDATOR_NODEIDS[0],
            DEFAULT_BLS_PUBLIC_KEY,
            2 days
        );

        bytes memory l1ValidatorRegistrationMessage =
            ValidatorMessages.packL1ValidatorRegistrationMessage(validationID, true);

        _mockGetPChainWarpMessage(l1ValidatorRegistrationMessage, true);

        slotauctionmanager.completeValidatorRegistration(0);
    }

    function testSetSlotAuctionSettings() public {
        AuctionSettings memory testAuctionSettings = AuctionSettings({
            totalValidatorSlots: DEFAULT_VALIDATOR_SLOTS - 2,
            weight: DEFAULT_VALIDATOR_SLOT_WEIGHT + 1,
            minValidatorDuration: DEFAULT_MINIMUM_VALIDATION_DURATION + 1,
            minAuctionDuration: DEFAULT_MINIMUM_AUCTION_DURATION + 1,
            minimumBid: DEFAULT_MINIMUM_BID + 1,
            auctionCooldownDuration: DEFUALT_AUCTION_COOLDOWN_DURATION + 1 days
        });

        vm.prank(DEFAULT_AUCTION_OWNER_ADRESS);
        slotauctionmanager.setSlotAuctionSettings(testAuctionSettings);
        vm.assertEq(
            slotauctionmanager.getTotalValidatorSlots(), testAuctionSettings.totalValidatorSlots
        );
        vm.assertEq(slotauctionmanager.getAuctioningValidatorWeight(), testAuctionSettings.weight);
        vm.assertEq(
            slotauctionmanager.getMinAuctionDuration(), testAuctionSettings.minAuctionDuration
        );
        vm.assertEq(
            slotauctionmanager.getMinValidatorDuration(), testAuctionSettings.minValidatorDuration
        );
        vm.assertEq(slotauctionmanager.getMinimumBid(), testAuctionSettings.minimumBid);
        vm.assertEq(
            slotauctionmanager.getOpenValidatorSlots(), testAuctionSettings.totalValidatorSlots
        );
    }

    function testSetSlotAuctionSettingsDuringAuction() public {
        _startAuctionWithCheck();
        vm.expectRevert(abi.encodeWithSelector(SlotAuctionManager.AuctionInProgress.selector));
        vm.prank(DEFAULT_AUCTION_OWNER_ADRESS);
        slotauctionmanager.setSlotAuctionSettings(DEFAULT_AUCTION_SETTINGS);
    }

    function testBidEvicted() public {
        // Setting slots to 1
        vm.prank(DEFAULT_AUCTION_OWNER_ADRESS);
        slotauctionmanager.setSlotAuctionSettings(
            AuctionSettings({
                totalValidatorSlots: 1,
                weight: DEFAULT_VALIDATOR_SLOT_WEIGHT,
                minValidatorDuration: DEFAULT_MINIMUM_VALIDATION_DURATION,
                minAuctionDuration: DEFAULT_MINIMUM_AUCTION_DURATION,
                minimumBid: DEFAULT_MINIMUM_BID,
                auctionCooldownDuration: DEFUALT_AUCTION_COOLDOWN_DURATION
            })
        );

        vm.expectEmit(address(slotauctionmanager));
        emit NewValidatorAuction(
            1,
            DEFAULT_VALIDATOR_SLOT_WEIGHT,
            DEFAULT_MINIMUM_VALIDATION_DURATION,
            DEFAULT_MINIMUM_AUCTION_DURATION + block.timestamp,
            DEFAULT_MINIMUM_BID
        );
        slotauctionmanager.initiateAuction();

        _fundAndPlaceBidWithCheck({
            sender: DEFAULT_SENDER_ADDRESSES[0],
            bid: DEFAULT_MINIMUM_BID,
            nodeID: DEFAULT_BIDDING_VALIDATOR_NODEIDS[0],
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER
        });

        _beforeBid(DEFAULT_MINIMUM_BID + 1, address(this));
        vm.expectEmit(address(slotauctionmanager));
        emit BidEvicted(DEFAULT_MINIMUM_BID, DEFAULT_BIDDING_VALIDATOR_NODEIDS[0]);
        _placeBid(
            DEFAULT_MINIMUM_BID + 1,
            DEFAULT_BIDDING_VALIDATOR_NODEIDS[1],
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_P_CHAIN_OWNER,
            DEFAULT_P_CHAIN_OWNER
        );
    }

    function testEndAuctionBeforeAuctionEndtime() public {
        vm.warp(0);
        _startAuctionWithCheck();
        vm.expectRevert(
            abi.encodeWithSelector(
                SlotAuctionManager.AuctionEndTimeNotReached.selector,
                DEFAULT_MINIMUM_AUCTION_DURATION
            )
        );
        slotauctionmanager.endAuction();
    }

    function testStartAuctionBeforeCooldownEnded() public {
        _startAuctionWithCheck();
        vm.warp(block.timestamp + DEFAULT_MINIMUM_AUCTION_DURATION);
        slotauctionmanager.endAuction();
        vm.expectRevert(
            abi.encodeWithSelector(
                SlotAuctionManager.AuctionInCooldown.selector, block.timestamp + 1 days
            )
        );
        slotauctionmanager.initiateAuction();
    }

    function testEndAuction() public {
        _startAuctionWithCheck();
        vm.warp(2000);
        slotauctionmanager.endAuction();
    }

    // TODO change maguc number
    function testAuctionSlotsLowerThanCurrentActiveValidators() public {
        _initiateAndCompleteValidatorRegistration(
            DEFAULT_SENDER_ADDRESSES[0],
            DEFAULT_BIDDING_VALIDATOR_NODEIDS[0],
            DEFAULT_BLS_PUBLIC_KEY,
            2 days
        );
        // _initiateAndCompleteValidatorRegistration(
        //     DEFAULT_SENDER_ADDRESSES[0],
        //     DEFAULT_BIDDING_VALIDATOR_NODEIDS[1],
        //     DEFAULT_BLS_PUBLIC_KEY,
        //     DEFAULT_AUCTION_END_TIME + 2 days
        // );

        // vm.assertEq(slotauctionmanager.getOpenValidatorSlots(), DEFAULT_VALIDATOR_SLOTS - 2);

        // vm.prank(DEFAULT_AUCTION_OWNER_ADRESS);
        // slotauctionmanager.setSlotAuctionSettings(
        //     AuctionSettings({
        //         totalValidatorSlots: DEFAULT_VALIDATOR_SLOTS - 2,
        //         weight: DEFAULT_VALIDATOR_SLOT_WEIGHT,
        //         minValidatorDuration: DEFAULT_MINIMUM_VALIDATION_DURATION,
        //         minAuctionDuration: DEFAULT_MINIMUM_AUCTION_DURATION,
        //         minimumBid: DEFAULT_MINIMUM_BID,
        //         auctionCooldownDuration: DEFUALT_AUCTION_COOLDOWN_DURATION
        //     })
        // );
    }

    function testMultipleWinners() public {
        _startAuctionWithCheck();
        uint256[4] memory bids = [
            DEFAULT_MINIMUM_BID,
            DEFAULT_MINIMUM_BID + 5,
            DEFAULT_MINIMUM_BID + 2,
            DEFAULT_MINIMUM_BID + 10
        ];
        /**
         * Skips the last index because _fundAndPlaceBidWithCheck will throw an error since
         * it checks for successful bid placed, and if a bid is evicted then the BidEvicted event
         * is emitted first causing a false negative
         */
        for (uint8 i = 0; i < bids.length - 1; i++) {
            _fundAndPlaceBidWithCheck({
                sender: DEFAULT_SENDER_ADDRESSES[i],
                bid: bids[i],
                nodeID: DEFAULT_BIDDING_VALIDATOR_NODEIDS[i],
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER
            });
        }
        _beforeBid(DEFAULT_MINIMUM_BID + 10, DEFAULT_SENDER_ADDRESSES[3]);
        vm.deal(DEFAULT_SENDER_ADDRESSES[3], DEFAULT_MINIMUM_BID + 10);

        /**
         * Sends the final bid with appropriate checks here
         */
        vm.expectEmit(address(slotauctionmanager));
        emit BidEvicted(DEFAULT_MINIMUM_BID, DEFAULT_BIDDING_VALIDATOR_NODEIDS[0]);
        vm.expectEmit(address(slotauctionmanager));
        emit SuccessfulBidPlaced(DEFAULT_MINIMUM_BID + 10, DEFAULT_BIDDING_VALIDATOR_NODEIDS[3]);

        vm.prank(DEFAULT_SENDER_ADDRESSES[3]);
        _placeBid(
            DEFAULT_MINIMUM_BID + 10,
            DEFAULT_BIDDING_VALIDATOR_NODEIDS[3],
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_P_CHAIN_OWNER,
            DEFAULT_P_CHAIN_OWNER
        );

        vm.warp(block.timestamp + DEFAULT_MINIMUM_AUCTION_DURATION);

        bytes32[3] memory validationIDs;
        bytes[3] memory registerL1ValidatorMessages;

        for (uint8 i = 0; i < validationIDs.length; i++) {
            (validationIDs[i], registerL1ValidatorMessages[i]) = _packRegisterL1ValidatorMessage({
                auctionEndTime: block.timestamp,
                subnetID: DEFAULT_SUBNET_ID,
                nodeID: DEFAULT_BIDDING_VALIDATOR_NODEIDS[i + 1],
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                weight: DEFAULT_VALIDATOR_SLOT_WEIGHT
            });
            _mockSendWarpMessage(registerL1ValidatorMessages[i], bytes32(0));
        }

        slotauctionmanager.endAuction();

        for (uint8 i = 0; i < validationIDs.length; i++) {
            bytes memory l1ValidatorRegistrationMessage =
                ValidatorMessages.packL1ValidatorRegistrationMessage(validationIDs[i], true);

            _mockGetPChainWarpMessage(l1ValidatorRegistrationMessage, true);
            slotauctionmanager.completeValidatorRegistration(0);
        }
    }

    function testChangeAuctionSettingsNonOwner() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, address(this)
            )
        );
        slotauctionmanager.setSlotAuctionSettings(DEFAULT_AUCTION_SETTINGS);
    }

    // TODO change magic number
    function testOpenValidatorSlotDecrease() public {
        vm.assertEq(slotauctionmanager.getOpenValidatorSlots(), DEFAULT_VALIDATOR_SLOTS);
        _initiateAndCompleteValidatorRegistration(
            DEFAULT_SENDER_ADDRESSES[0],
            DEFAULT_BIDDING_VALIDATOR_NODEIDS[0],
            DEFAULT_BLS_PUBLIC_KEY,
            2 days
        );
        vm.assertEq(slotauctionmanager.getOpenValidatorSlots(), DEFAULT_VALIDATOR_SLOTS - 1);
    }

    // Simulates initiateRegisterValidator by bidding on an auction and ending it, causing the one person to win
    function _initiateRegisterValidator(
        address sender,
        bytes memory nodeID,
        bytes memory blsPublicKey,
        uint256 auctionEndTime
    ) internal returns (bytes32) {
        vm.warp(auctionEndTime - DEFAULT_MINIMUM_AUCTION_DURATION);
        _startAuctionWithCheck();
        _fundAndPlaceBidWithCheck({
            sender: sender,
            bid: DEFAULT_MINIMUM_BID,
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER
        });

        uint64 registrationExpiry = uint64(auctionEndTime) + 1 days;

        (bytes32 validationID, bytes memory registerL1ValidatorMessage) = ValidatorMessages
            .packRegisterL1ValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                nodeID: nodeID,
                subnetID: DEFAULT_SUBNET_ID,
                blsPublicKey: blsPublicKey,
                registrationExpiry: registrationExpiry,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                weight: DEFAULT_VALIDATOR_SLOT_WEIGHT
            })
        );

        _mockSendWarpMessage(registerL1ValidatorMessage, bytes32(0));

        vm.warp(auctionEndTime);

        vm.expectEmit(address(slotauctionmanager));
        emit InitiatedAuctionValidatorRegistration(
            validationID,
            sender,
            DEFAULT_MINIMUM_VALIDATION_DURATION + block.timestamp,
            DEFAULT_VALIDATOR_SLOT_WEIGHT
        );
        slotauctionmanager.endAuction();
        return validationID;
    }

    function _fundAndPlaceBidWithCheck(
        address sender,
        uint256 bid,
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) internal {
        _beforeBid(bid, sender);
        vm.expectEmit();
        emit SuccessfulBidPlaced(bid, nodeID);
        vm.deal(sender, bid);
        vm.prank(sender);
        _placeBid(bid, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner);
    }

    function _startAuctionWithCheck() internal {
        vm.expectEmit(address(slotauctionmanager));
        emit NewValidatorAuction(
            slotauctionmanager.getOpenValidatorSlots(),
            DEFAULT_VALIDATOR_SLOT_WEIGHT,
            DEFAULT_MINIMUM_VALIDATION_DURATION,
            DEFAULT_MINIMUM_AUCTION_DURATION + block.timestamp,
            DEFAULT_MINIMUM_BID
        );
        slotauctionmanager.initiateAuction();
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

    function _initiateAndCompleteValidatorRegistration(
        address sender,
        bytes memory nodeID,
        bytes memory blsPublicKey,
        uint256 auctionEndTime
    ) internal {
        bytes32 validationID =
            _initiateRegisterValidator(sender, nodeID, blsPublicKey, auctionEndTime);

        bytes memory l1ValidatorRegistrationMessage =
            ValidatorMessages.packL1ValidatorRegistrationMessage(validationID, true);

        _mockGetPChainWarpMessage(l1ValidatorRegistrationMessage, true);

        slotauctionmanager.completeValidatorRegistration(0);
    }

    function _mockSendWarpMessage(bytes memory payload, bytes32 expectedMessageID) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(expectedMessageID)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.sendWarpMessage, payload)
        );
    }

    
    
    function _setUp() internal virtual;

    function _beforeBid(uint256 amount, address spender) internal virtual;

    function _placeBid(
        uint256 bid,
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) internal virtual;

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

    function _packRegisterL1ValidatorMessage(
        uint256 auctionEndTime,
        bytes32 subnetID,
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) internal pure returns (bytes32, bytes memory) {
        uint64 registrationExpiry = uint64(auctionEndTime) + 1 days;

        (bytes32 validationID, bytes memory registerL1ValidatorMessage) = ValidatorMessages
            .packRegisterL1ValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                nodeID: nodeID,
                subnetID: subnetID,
                blsPublicKey: blsPublicKey,
                registrationExpiry: registrationExpiry,
                remainingBalanceOwner: remainingBalanceOwner,
                disableOwner: disableOwner,
                weight: weight
            })
        );
        return (validationID, registerL1ValidatorMessage);
    }
}
