// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {
    TeleporterMessengerTest,
    TeleporterMessage,
    TeleporterMessageReceipt,
    WarpMessage
} from "./TeleporterMessengerTest.t.sol";
import {ITeleporterMessenger} from "../ITeleporterMessenger.sol";
import {ITeleporterReceiver} from "../ITeleporterReceiver.sol";

enum SampleMessageReceiverAction {
    Receive,
    ReceiveRecursive
}

contract SampleMessageReceiver is ITeleporterReceiver {
    address public immutable teleporterContract;
    string public latestMessage;
    bytes32 public latestMessageSenderSubnetID;
    address public latestMessageSenderAddress;

    constructor(
        address teleporterContractAddress
    ) {
        teleporterContract = teleporterContractAddress;
    }

    function receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes calldata message
    ) external {
        require(msg.sender == teleporterContract, "unauthorized");
        // Decode the payload to recover the action and corresponding function parameters
        (SampleMessageReceiverAction action, bytes memory actionData) =
            abi.decode(message, (SampleMessageReceiverAction, bytes));
        if (action == SampleMessageReceiverAction.Receive) {
            (string memory messageString, bool succeed) = abi.decode(actionData, (string, bool));
            _receiveMessage(sourceBlockchainID, originSenderAddress, messageString, succeed);
        } else if (action == SampleMessageReceiverAction.ReceiveRecursive) {
            string memory messageString = abi.decode(actionData, (string));
            _receiveMessageRecursive(sourceBlockchainID, originSenderAddress, messageString);
        } else {
            revert("invalid action");
        }
    }

    // Stores the message in this contract to be fetched by anyone.
    function _receiveMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        string memory message,
        bool succeed
    ) internal {
        require(msg.sender == teleporterContract, "unauthorized");
        require(succeed, "intended to fail");
        latestMessage = message;
        latestMessageSenderSubnetID = sourceBlockchainID;
        latestMessageSenderAddress = originSenderAddress;
    }

    // Tries to recursively call the teleporterContract to receive a message, which should always fail.
    function _receiveMessageRecursive(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        string memory message
    ) internal {
        require(msg.sender == teleporterContract, "unauthorized");
        ITeleporterMessenger messenger = ITeleporterMessenger(teleporterContract);
        messenger.receiveCrossChainMessage(0, address(42));
        latestMessage = message;
        latestMessageSenderSubnetID = sourceBlockchainID;
        latestMessageSenderAddress = originSenderAddress;
    }
}

contract HandleInitialMessageExecutionTest is TeleporterMessengerTest {
    SampleMessageReceiver public destinationContract;

    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
        destinationContract = new SampleMessageReceiver(address(teleporterMessenger));
    }

    function testSuccess() public {
        // Construct the mock message to be received.
        string memory messageString = "Testing successful message";
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageNonce: 42,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: address(destinationContract),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: abi.encode(SampleMessageReceiverAction.Receive, abi.encode(messageString, true))
        });
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the message and check that message execution was successful.
        bytes32 expectedMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            messageToReceive.messageNonce
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiveCrossChainMessage(
            expectedMessageID,
            warpMessage.sourceChainID,
            address(this),
            DEFAULT_RELAYER_REWARD_ADDRESS,
            messageToReceive
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit MessageExecuted(expectedMessageID, DEFAULT_SOURCE_BLOCKCHAIN_ID);
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);

        // Check that the message had the proper affect on the destination contract.
        assertEq(destinationContract.latestMessage(), messageString);
        assertEq(destinationContract.latestMessageSenderSubnetID(), DEFAULT_SOURCE_BLOCKCHAIN_ID);
        assertEq(destinationContract.latestMessageSenderAddress(), address(this));
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(expectedMessageID),
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
    }

    function testInsufficientGasProvided() public {
        // Construct the mock message to be received.
        string memory messageString = "Testing successful message";
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageNonce: 42,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: address(destinationContract),
            requiredGasLimit: uint256(
                bytes32(hex"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF") // UINT256_MAX
            ),
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: abi.encode(SampleMessageReceiverAction.Receive, abi.encode(messageString, true))
        });
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the message.
        vm.expectRevert(_formatTeleporterErrorMessage("insufficient gas"));
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);
    }

    function testCannotReceiveMessageRecursively() public {
        // Construct the mock message to be received.
        string memory messageString = "Testing successful message";
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageNonce: 42,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: address(destinationContract),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: abi.encode(SampleMessageReceiverAction.ReceiveRecursive, abi.encode(messageString))
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

        // Receive the message - this does not revert because the recursive call
        // is considered a failed message execution, but the message itself is
        // still successfully delivered.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiveCrossChainMessage(
            messageID,
            warpMessage.sourceChainID,
            address(this),
            DEFAULT_RELAYER_REWARD_ADDRESS,
            messageToReceive
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit MessageExecutionFailed(messageID, DEFAULT_SOURCE_BLOCKCHAIN_ID, messageToReceive);
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);

        // Check that the message hash was stored in state and the message did not have any affect on the destination.
        assertEq(destinationContract.latestMessage(), "");
        assertEq(destinationContract.latestMessageSenderSubnetID(), bytes32(0));
        assertEq(destinationContract.latestMessageSenderAddress(), address(0));
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(messageID), DEFAULT_RELAYER_REWARD_ADDRESS
        );
        vm.expectRevert(_formatTeleporterErrorMessage("retry execution failed"));
        teleporterMessenger.retryMessageExecution(DEFAULT_SOURCE_BLOCKCHAIN_ID, messageToReceive);
    }

    function testStoreHashOfFailedMessageExecution() public {
        // Construct the mock message to be received.
        string memory messageString = "Testing successful message";
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageNonce: 42,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: address(destinationContract),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: abi.encode(SampleMessageReceiverAction.Receive, abi.encode(messageString, false))
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

        // Receive the message.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiveCrossChainMessage(
            messageID,
            warpMessage.sourceChainID,
            address(this),
            DEFAULT_RELAYER_REWARD_ADDRESS,
            messageToReceive
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit MessageExecutionFailed(messageID, DEFAULT_SOURCE_BLOCKCHAIN_ID, messageToReceive);
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);

        // Check that the message hash was stored in state and the message did not have any affect on the destination.
        assertEq(destinationContract.latestMessage(), "");
        assertEq(destinationContract.latestMessageSenderSubnetID(), bytes32(0));
        assertEq(destinationContract.latestMessageSenderAddress(), address(0));
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(messageID), DEFAULT_RELAYER_REWARD_ADDRESS
        );
        vm.expectRevert(_formatTeleporterErrorMessage("retry execution failed"));
        teleporterMessenger.retryMessageExecution(DEFAULT_SOURCE_BLOCKCHAIN_ID, messageToReceive);
    }
}
