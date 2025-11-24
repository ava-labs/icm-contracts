// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {
    TeleporterMessengerTest,
    TeleporterMessage,
    TeleporterMessageInput,
    TeleporterFeeInfo,
    TeleporterMessageReceipt,
    IWarpMessenger
} from "./TeleporterMessengerTest.t.sol";

contract GetOutstandingReceiptsToSendTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    // getOutstandingReceipts to send is private to the teleporter contract, but we can test
    // its correctness by looking at the receipts included in messages that are submitted to be
    // sent. First, mock receiving 3 messages such that we have 3 outstanding receipts. Then, send
    // a message that should contain the receipts. Finally, send another message that should not
    // have any more receipts since they were included in the previous message.
    function testSuccess() public {
        bytes32 blockchainID =
            bytes32(hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff");

        // Assemble mock receipt information that we will expect to be
        // included in a subsequent message sent to another chain.
        TeleporterMessageReceipt[] memory expectedReceipts = new TeleporterMessageReceipt[](3);
        expectedReceipts[0] =
            TeleporterMessageReceipt(13, 0xF1DFE63909C027Ed814Dd92C5a3644590abf4850);
        expectedReceipts[1] =
            TeleporterMessageReceipt(42, 0xF1DFE63909C027Ed814Dd92C5a3644590abf4850);
        expectedReceipts[2] =
            TeleporterMessageReceipt(94, 0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542);

        // Mock receiving each of the messages corresponding to the receipts.
        for (uint256 i; i < expectedReceipts.length; ++i) {
            _receiveTestMessage(
                blockchainID,
                expectedReceipts[i].receivedMessageNonce,
                expectedReceipts[i].relayerRewardAddress,
                new TeleporterMessageReceipt[](0)
            );
        }

        // Now that we have "received" 3 mock messages, when we send a message back to the
        // other chain, we should expect to see the 3 receipts included in the message metadata.
        TeleporterMessage memory expectedMessage =
            _createMockTeleporterMessage(_getNextMessageNonce(), hex"deadbeef");
        expectedMessage.receipts = expectedReceipts;
        expectedMessage.destinationBlockchainID = blockchainID;
        bytes32 expectedMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            expectedMessage.destinationBlockchainID,
            expectedMessage.messageNonce
        );
        TeleporterFeeInfo memory feeInfo = TeleporterFeeInfo(address(0), 0);
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: expectedMessage.destinationBlockchainID,
            destinationAddress: expectedMessage.destinationAddress,
            feeInfo: feeInfo,
            requiredGasLimit: expectedMessage.requiredGasLimit,
            allowedRelayerAddresses: expectedMessage.allowedRelayerAddresses,
            message: expectedMessage.message
        });

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(expectedMessageID)
        );

        // Expect the exact message to be passed to the precompile.
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(expectedMessage)))
        );

        // Expect the SendCrossChainMessage event to be emitted.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            expectedMessageID, messageInput.destinationBlockchainID, expectedMessage, feeInfo
        );

        // Submit the message.
        teleporterMessenger.sendCrossChainMessage(messageInput);

        // Submit another message to be sent to check that it does not contain any more receipts.
        TeleporterMessage memory nextExpectedMessage =
            _createMockTeleporterMessage(_getNextMessageNonce(), hex"deadbeef");
        nextExpectedMessage.destinationBlockchainID = blockchainID;
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(nextExpectedMessage)))
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            teleporterMessenger.calculateMessageID(
                DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                nextExpectedMessage.destinationBlockchainID,
                nextExpectedMessage.messageNonce
            ),
            messageInput.destinationBlockchainID,
            nextExpectedMessage,
            feeInfo
        );

        // Submit the new message.
        teleporterMessenger.sendCrossChainMessage(messageInput);
    }

    // Test that when there are more than the maximum limit of receipts to include in a message (5),
    // the first batch of receipts on the first outbound message has the maximum batch size, and
    // next contains the remaining receipts.
    function testExceedsLimit() public {
        bytes32 blockchainID =
            bytes32(hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff");

        // Assemble mock receipt information that we will expect to be
        // included in a subsequent message sent to another chain.
        TeleporterMessageReceipt[] memory expectedReceiptsBatch1 = new TeleporterMessageReceipt[](5); // the limit of receipts per message is 5.
        expectedReceiptsBatch1[0] =
            TeleporterMessageReceipt(13, 0xF1DFE63909C027Ed814Dd92C5a3644590abf4850);
        expectedReceiptsBatch1[1] =
            TeleporterMessageReceipt(42, 0x52A258ED593C793251a89bfd36caE158EE9fC4F8);
        expectedReceiptsBatch1[2] =
            TeleporterMessageReceipt(94, 0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542);
        expectedReceiptsBatch1[3] =
            TeleporterMessageReceipt(3, 0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542);
        expectedReceiptsBatch1[4] =
            TeleporterMessageReceipt(53, 0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542);

        TeleporterMessageReceipt[] memory expectedReceiptsBatch2 = new TeleporterMessageReceipt[](2); // the limit of receipts per message is 5.
        expectedReceiptsBatch2[0] =
            TeleporterMessageReceipt(75, 0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542);
        expectedReceiptsBatch2[1] =
            TeleporterMessageReceipt(80, 0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542);

        // Mock receiving each of the messages corresponding to the receipts.
        for (uint256 i; i < expectedReceiptsBatch1.length; ++i) {
            _receiveTestMessage(
                blockchainID,
                expectedReceiptsBatch1[i].receivedMessageNonce,
                expectedReceiptsBatch1[i].relayerRewardAddress,
                new TeleporterMessageReceipt[](0)
            );
        }
        for (uint256 i; i < expectedReceiptsBatch2.length; ++i) {
            _receiveTestMessage(
                blockchainID,
                expectedReceiptsBatch2[i].receivedMessageNonce,
                expectedReceiptsBatch2[i].relayerRewardAddress,
                new TeleporterMessageReceipt[](0)
            );
        }

        // Now that we have "received" 7 mock messages, when we send a message back to the
        // other chain, we should expect to see the 5 receipts included in the message metadata because
        // that is the max receipt batch size limit.
        TeleporterMessage memory expectedMessage =
            _createMockTeleporterMessage(_getNextMessageNonce(), hex"deadbeef");
        expectedMessage.receipts = expectedReceiptsBatch1;
        expectedMessage.destinationBlockchainID = blockchainID;
        TeleporterFeeInfo memory feeInfo = TeleporterFeeInfo(address(0), 0);
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: expectedMessage.destinationBlockchainID,
            destinationAddress: expectedMessage.destinationAddress,
            feeInfo: feeInfo,
            requiredGasLimit: expectedMessage.requiredGasLimit,
            allowedRelayerAddresses: expectedMessage.allowedRelayerAddresses,
            message: expectedMessage.message
        });

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(bytes32(0))
        );

        // Expect the exact message to be passed to the precompile.
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(expectedMessage)))
        );

        // Expect the SendCrossChainMessage event to be emitted.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            teleporterMessenger.calculateMessageID(
                DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                expectedMessage.destinationBlockchainID,
                expectedMessage.messageNonce
            ),
            messageInput.destinationBlockchainID,
            expectedMessage,
            feeInfo
        );

        // Submit the message.
        teleporterMessenger.sendCrossChainMessage(messageInput);

        // Submit another message to be sent to check that it contains the remaining 2 receipts to be sent.
        TeleporterMessage memory nextExpectedMessage =
            _createMockTeleporterMessage(_getNextMessageNonce(), hex"deadbeef");
        nextExpectedMessage.receipts = expectedReceiptsBatch2;
        nextExpectedMessage.destinationBlockchainID = blockchainID;
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(nextExpectedMessage)))
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            teleporterMessenger.calculateMessageID(
                DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                nextExpectedMessage.destinationBlockchainID,
                nextExpectedMessage.messageNonce
            ),
            messageInput.destinationBlockchainID,
            nextExpectedMessage,
            feeInfo
        );

        // Submit the new message.
        teleporterMessenger.sendCrossChainMessage(messageInput);
    }
}
