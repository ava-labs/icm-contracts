// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {
    TeleporterMessengerTest,
    TeleporterMessage,
    TeleporterMessageReceipt,
    TeleporterFeeInfo,
    IWarpMessenger,
    IERC20
} from "./TeleporterMessengerTest.t.sol";

contract SendSpecifiedReceiptsTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // Mock receiving 3 messages from another chain
        address[3] memory relayerRewardAddresses = [
            0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d,
            0x22d4002028f537599bE9f666d1c4Fa138522f9c8,
            0xb85e0651385876C3E3A1fF81ABf410ecf12f52f4
        ];

        TeleporterMessageReceipt[] memory expectedReceipts = new TeleporterMessageReceipt[](3);
        for (uint256 i; i < relayerRewardAddresses.length; ++i) {
            _receiveTestMessage(
                DEFAULT_SOURCE_BLOCKCHAIN_ID,
                i + 1,
                relayerRewardAddresses[i],
                new TeleporterMessageReceipt[](0)
            );
            expectedReceipts[i] = TeleporterMessageReceipt({
                receivedMessageNonce: i + 1,
                relayerRewardAddress: relayerRewardAddresses[i]
            });
        }

        // Mock sending a message back to that chain, which should include the 3 receipts.
        uint256 expectedMessageNonce = _getNextMessageNonce();
        TeleporterMessage memory expectedMessage = TeleporterMessage({
            messageNonce: expectedMessageNonce,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: expectedReceipts,
            message: new bytes(0)
        });
        bytes32 expectedMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_SOURCE_BLOCKCHAIN_ID, expectedMessageNonce
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        TeleporterFeeInfo memory feeInfo = TeleporterFeeInfo(address(0), 0);
        emit SendCrossChainMessage(
            expectedMessageID, DEFAULT_SOURCE_BLOCKCHAIN_ID, expectedMessage, feeInfo
        );
        bytes32 outboundMessageID = _sendTestMessageWithNoFee(DEFAULT_SOURCE_BLOCKCHAIN_ID);
        assertEq(outboundMessageID, expectedMessageID);
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(
                teleporterMessenger.calculateMessageID(
                    DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, 1
                )
            ),
            relayerRewardAddresses[0]
        );

        TeleporterMessageReceipt[] memory newExpectedReceipts = new TeleporterMessageReceipt[](2);
        newExpectedReceipts[0] = TeleporterMessageReceipt({
            receivedMessageNonce: 3,
            relayerRewardAddress: relayerRewardAddresses[2]
        });
        newExpectedReceipts[1] = TeleporterMessageReceipt({
            receivedMessageNonce: 1,
            relayerRewardAddress: relayerRewardAddresses[0]
        });

        uint256 newExpectedMessageNonce = _getNextMessageNonce();
        bytes32 newExpectedMessageID =
            teleporterMessenger.getNextMessageID(DEFAULT_SOURCE_BLOCKCHAIN_ID);
        TeleporterMessage memory newExpectedMessage = TeleporterMessage({
            messageNonce: newExpectedMessageNonce,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationAddress: address(0),
            requiredGasLimit: uint256(0),
            allowedRelayerAddresses: new address[](0),
            receipts: newExpectedReceipts,
            message: new bytes(0)
        });

        // Retry sending two of the receipts with {sendSpecifiedReceipts}.
        bytes32[] memory receiptIDs = new bytes32[](2);
        receiptIDs[0] = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, 3
        );
        receiptIDs[1] = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, 1
        );

        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            newExpectedMessageID, DEFAULT_SOURCE_BLOCKCHAIN_ID, newExpectedMessage, feeInfo
        );

        outboundMessageID =
            _sendSpecifiedReceiptsWithNoFee(DEFAULT_SOURCE_BLOCKCHAIN_ID, receiptIDs);

        assertEq(outboundMessageID, newExpectedMessageID);
    }

    function testMessageIDNotFromSourceBlockchain() public {
        // Mock receiving a message from another chain
        uint256 receivedMessageNonce = 987;
        _receiveTestMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            receivedMessageNonce,
            address(0),
            new TeleporterMessageReceipt[](0)
        );
        bytes32[] memory receiptIDs = new bytes32[](1);
        receiptIDs[0] = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, receivedMessageNonce
        );

        // Call sendSpecifiedReceipts to try to send the receipts for those message IDs to a
        // different chain than they were delivered by.
        bytes32 otherBlockchainID =
            bytes32(hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff");
        vm.expectRevert(_formatTeleporterErrorMessage("message ID not from source blockchain"));
        _sendSpecifiedReceiptsWithNoFee(otherBlockchainID, receiptIDs);
    }

    function testDuplicateAllowed() public {
        // Mock receiving a message from another chain
        _receiveTestMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            1,
            DEFAULT_RELAYER_REWARD_ADDRESS,
            new TeleporterMessageReceipt[](0)
        );

        // Mock sending a message to that chain, which should include the 2 copies of the same receipt.
        TeleporterMessageReceipt memory expectedReceipt = TeleporterMessageReceipt({
            receivedMessageNonce: 1,
            relayerRewardAddress: DEFAULT_RELAYER_REWARD_ADDRESS
        });
        TeleporterMessageReceipt[] memory expectedReceipts = new TeleporterMessageReceipt[](2);
        expectedReceipts[0] = expectedReceipt;
        expectedReceipts[1] = expectedReceipt;
        uint256 expectedMessageNonce = _getNextMessageNonce();
        bytes32 expectedMessageID =
            teleporterMessenger.getNextMessageID(DEFAULT_DESTINATION_BLOCKCHAIN_ID);
        TeleporterMessage memory expectedMessage = TeleporterMessage({
            messageNonce: expectedMessageNonce,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: address(0),
            requiredGasLimit: uint256(0),
            allowedRelayerAddresses: new address[](0),
            receipts: expectedReceipts,
            message: new bytes(0)
        });

        bytes32[] memory messageIDs = new bytes32[](2);
        messageIDs[0] = teleporterMessenger.calculateMessageID(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, 1
        );
        messageIDs[1] = teleporterMessenger.calculateMessageID(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, 1
        );

        // Test sendSpecifiedReceipts when there are duplicate message IDs in the input.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            expectedMessageID,
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            expectedMessage,
            TeleporterFeeInfo(address(0), 0)
        );

        bytes32 outboundMessageID =
            _sendSpecifiedReceiptsWithNoFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, messageIDs);
        assertEq(outboundMessageID, expectedMessageID);
    }

    function testMissingMessage() public {
        // Mock receiving a message from another chain
        _receiveTestMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            1,
            DEFAULT_RELAYER_REWARD_ADDRESS,
            new TeleporterMessageReceipt[](0)
        );

        bytes32[] memory missingIDs = new bytes32[](1);
        missingIDs[0] = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, 21
        );

        // Try to send a receipt for an unreceived message from that chain - should fail.
        vm.expectRevert(_formatTeleporterErrorMessage("receipt not found"));
        _sendSpecifiedReceiptsWithNoFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, missingIDs);
    }

    function _sendSpecifiedReceiptsWithFee(
        bytes32 blockchainID,
        bytes32[] memory messageIDs,
        address feeAddress,
        uint256 feeAmount
    ) private returns (bytes32) {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(bytes32(0))
        );
        if (feeAmount > 0) {
            vm.mockCall(
                feeAddress,
                abi.encodeCall(
                    IERC20.transferFrom, (address(this), address(teleporterMessenger), feeAmount)
                ),
                abi.encode(true)
            );
        }

        TeleporterFeeInfo memory feeInfo =
            TeleporterFeeInfo({feeTokenAddress: feeAddress, amount: feeAmount});

        if (feeAmount > 0) {
            // Expect the ERC20 contract transferFrom method to be called to transfer the fee.
            vm.expectCall(
                feeAddress,
                abi.encodeCall(
                    IERC20.transferFrom, (address(this), address(teleporterMessenger), feeAmount)
                )
            );
        }

        return teleporterMessenger.sendSpecifiedReceipts(
            blockchainID, messageIDs, feeInfo, new address[](0)
        );
    }

    function _sendSpecifiedReceiptsWithNoFee(
        bytes32 blockchainID,
        bytes32[] memory messageIDs
    ) private returns (bytes32) {
        return _sendSpecifiedReceiptsWithFee(blockchainID, messageIDs, address(0), 0);
    }
}
