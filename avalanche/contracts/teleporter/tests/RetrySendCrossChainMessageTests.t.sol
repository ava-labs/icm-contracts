// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {
    TeleporterMessengerTest,
    TeleporterMessage,
    TeleporterMessageReceipt
} from "./TeleporterMessengerTest.t.sol";

contract RetrySendCrossChainMessageTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // Send a message
        uint256 expectedNonce = _getNextMessageNonce();
        _sendTestMessageWithFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, 654456);
        TeleporterMessage memory expectedMessage = TeleporterMessage({
            messageNonce: expectedNonce,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: new bytes(0)
        });

        // Retry it
        teleporterMessenger.retrySendCrossChainMessage(expectedMessage);
    }

    function testMessageNotFound() public {
        TeleporterMessage memory fakeMessage = TeleporterMessage({
            messageNonce: 345,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: new bytes(0)
        });
        vm.expectRevert(_formatTeleporterErrorMessage("message not found"));
        teleporterMessenger.retrySendCrossChainMessage(fakeMessage);
    }

    function testInvalidMessageHash() public {
        // Send a message, then try to alter it's contents.
        uint256 expectedNonce = _getNextMessageNonce();
        _sendTestMessageWithFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, 654456);
        TeleporterMessage memory alteredMessage = TeleporterMessage({
            messageNonce: expectedNonce,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: hex"cafebabe" // The message would be expected to be empty
        });

        // Retry it - should fail.
        vm.expectRevert(_formatTeleporterErrorMessage("invalid message hash"));
        teleporterMessenger.retrySendCrossChainMessage(alteredMessage);
    }
}
