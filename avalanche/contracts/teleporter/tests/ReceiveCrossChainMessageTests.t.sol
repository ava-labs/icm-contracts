// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {
    TeleporterMessengerTest,
    TeleporterMessage,
    TeleporterMessageReceipt,
    WarpMessage,
    IWarpMessenger
} from "./TeleporterMessengerTest.t.sol";

// Tests of the logic in receiveCrossChainMessage.
// Tests of the execution and receipts helper methods are split out into independent
// test files in HandleInitialMessageExecutionTests and MarkReceiptAndPayRelayerTests.
contract ReceiveCrossChainMessagedTest is TeleporterMessengerTest {
    bytes public constant DEFAULT_MESSAGE_PAYLOAD =
        hex"cafebabe11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff11223344556677889900aabbccddeeffdeadbeef";

    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // This test contract must be an allowed relayer since it is what
        // will call receiveCrossChainMessage.
        address[] memory allowedRelayers = new address[](2);
        allowedRelayers[0] = address(this);
        allowedRelayers[1] = DEFAULT_RELAYER_REWARD_ADDRESS;

        // Construct the test message to be received.
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageNonce: 987,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: allowedRelayers,
            receipts: new TeleporterMessageReceipt[](0),
            message: DEFAULT_MESSAGE_PAYLOAD
        });
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));
        bytes32 messageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            messageToReceive.messageNonce
        );

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Check receipt queue size
        assertEq(teleporterMessenger.getReceiptQueueSize(DEFAULT_SOURCE_BLOCKCHAIN_ID), 0);

        // Receive the message.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiveCrossChainMessage(
            messageID,
            warpMessage.sourceChainID,
            address(this),
            DEFAULT_RELAYER_REWARD_ADDRESS,
            messageToReceive
        );
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);

        // Check receipt queue size
        assertEq(teleporterMessenger.getReceiptQueueSize(DEFAULT_SOURCE_BLOCKCHAIN_ID), 1);

        // Check receipt queue contents
        TeleporterMessageReceipt memory receipt =
            teleporterMessenger.getReceiptAtIndex(DEFAULT_SOURCE_BLOCKCHAIN_ID, 0);
        assertEq(receipt.receivedMessageNonce, messageToReceive.messageNonce);
        assertEq(receipt.relayerRewardAddress, DEFAULT_RELAYER_REWARD_ADDRESS);

        // Receive at a different index
        messageToReceive.messageNonce = 654;
        messageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            messageToReceive.messageNonce
        );
        warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));
        _setUpSuccessGetVerifiedWarpMessageMock(3, warpMessage);

        // Receive the message.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiveCrossChainMessage(
            messageID,
            warpMessage.sourceChainID,
            address(this),
            DEFAULT_RELAYER_REWARD_ADDRESS,
            messageToReceive
        );
        teleporterMessenger.receiveCrossChainMessage(3, DEFAULT_RELAYER_REWARD_ADDRESS);

        // Check receipt queue size
        assertEq(teleporterMessenger.getReceiptQueueSize(DEFAULT_SOURCE_BLOCKCHAIN_ID), 2);

        // Check receipt queue contents
        receipt = teleporterMessenger.getReceiptAtIndex(DEFAULT_SOURCE_BLOCKCHAIN_ID, 1);
        assertEq(receipt.receivedMessageNonce, messageToReceive.messageNonce);
        assertEq(receipt.relayerRewardAddress, DEFAULT_RELAYER_REWARD_ADDRESS);
    }

    function testNoValidMessage() public {
        // Mock the call to the warp precompile to get the message failing.
        WarpMessage memory emptyMessage = _createDefaultWarpMessage(bytes32(0), new bytes(0));

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, (0)),
            abi.encode(emptyMessage, false)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, (0))
        );

        vm.expectRevert(_formatTeleporterErrorMessage("invalid warp message"));
        teleporterMessenger.receiveCrossChainMessage(0, address(1));

        // Receive invalid message at index 3
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, (3)),
            abi.encode(emptyMessage, false)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, (3))
        );

        vm.expectRevert(_formatTeleporterErrorMessage("invalid warp message"));
        teleporterMessenger.receiveCrossChainMessage(3, address(1));
    }

    function testInvalidOriginSenderAddress() public {
        // Construct the test message to be received.
        TeleporterMessage memory messageToReceive = _createMockTeleporterMessage(1, new bytes(0));

        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));
        address invalidSenderAddress = 0xb73aD7e0FF026a805D1f1186EAB89E41bf01835D;
        warpMessage.originSenderAddress = invalidSenderAddress;

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        vm.expectRevert(_formatTeleporterErrorMessage("invalid origin sender address"));
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);
    }

    function testInvalidDestinationBlockchainID() public {
        // Construct the test message to be received.
        TeleporterMessage memory messageToReceive = _createMockTeleporterMessage(1, new bytes(0));
        bytes32 invalidDestinationBlockchainID =
            bytes32(hex"deadbeefcafebabedeadbeefcafebabedeadbeefcafebabedeadbeefcafebabe");
        messageToReceive.destinationBlockchainID = invalidDestinationBlockchainID;

        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        vm.expectRevert(_formatTeleporterErrorMessage("invalid destination chain ID"));
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);
    }

    function testZeroMessageNonce() public {
        // Construct the test message to be received.
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageNonce: 0,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: DEFAULT_MESSAGE_PAYLOAD
        });
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the message.
        vm.expectRevert(_formatTeleporterErrorMessage("zero message nonce"));
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);
    }

    function testReceiveSameNonceFromMultipleChains() public {
        bytes32 sourceBlockchainID1 =
            bytes32(hex"1111111111111111111111111111111111111111111111111111111111111111");
        bytes32 sourceBlockchainID2 =
            bytes32(hex"2222222222222222222222222222222222222222222222222222222222222222");
        uint256 mockMessageNonce = 5454;
        bytes32 expectedMessageID1 = teleporterMessenger.calculateMessageID(
            sourceBlockchainID1, DEFAULT_DESTINATION_BLOCKCHAIN_ID, mockMessageNonce
        );
        bytes32 expectedMessageID2 = teleporterMessenger.calculateMessageID(
            sourceBlockchainID2, DEFAULT_DESTINATION_BLOCKCHAIN_ID, mockMessageNonce
        );

        // Receive two messages with the same nonce from different source chains.
        _receiveTestMessage(
            sourceBlockchainID1, mockMessageNonce, address(0), new TeleporterMessageReceipt[](0)
        );
        _receiveTestMessage(
            sourceBlockchainID2, mockMessageNonce, address(0), new TeleporterMessageReceipt[](0)
        );

        // Check that both messages were received.
        assertTrue(teleporterMessenger.messageReceived(expectedMessageID1));
        assertTrue(teleporterMessenger.messageReceived(expectedMessageID2));
    }

    function testMessageAlreadyReceived() public {
        // Receive a message successfully.
        testSuccess();

        // Check you can't deliver it again.
        vm.expectRevert(_formatTeleporterErrorMessage("message already received"));
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);
    }

    function testUnauthorizedRelayer() public {
        // It doesn't matter if the reward address is an allowed relayer.
        // It is based on who calls receiveCrossChainMessage.
        address[] memory allowedRelayers = new address[](1);
        allowedRelayers[0] = DEFAULT_RELAYER_REWARD_ADDRESS;

        // Construct the test message to be received.
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageNonce: 42,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: allowedRelayers,
            receipts: new TeleporterMessageReceipt[](0),
            message: DEFAULT_MESSAGE_PAYLOAD
        });
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the message.
        vm.expectRevert(_formatTeleporterErrorMessage("unauthorized relayer"));
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);
    }

    function testZeroRelayerAddress() public {
        _receiveTestMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, 12, address(0), new TeleporterMessageReceipt[](0)
        );
    }

    function testMessageSentToEOADoesNotExecute() public {
        _receiveMessageSentToEOA();
    }
}
