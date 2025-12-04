// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {
    TeleporterMessengerTest,
    TeleporterMessage,
    TeleporterMessageReceipt
} from "./TeleporterMessengerTest.t.sol";

contract GetMessageHashTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // Submit a message
        TeleporterMessage memory expectedMessage = TeleporterMessage({
            messageNonce: _getNextMessageNonce(),
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: new bytes(0)
        });
        bytes32 messageID = _sendTestMessageWithNoFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID);

        bytes memory expectedMessageBytes = abi.encode(expectedMessage);
        bytes32 expectedMessageHash = keccak256(expectedMessageBytes);

        // Get its stored hash
        bytes32 actualMessageHash = teleporterMessenger.getMessageHash(messageID);
        assertEq(actualMessageHash, expectedMessageHash);
    }

    function testMessageAlreadyReceived() public {
        // Submit a message
        uint256 expectedNonce = _getNextMessageNonce();
        bytes32 messageID = _sendTestMessageWithNoFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID);

        // Now mock receiving a message back from that L1 with a receipt of the above message.
        address relayerRewardAddress = 0xA66884fAdC0D4d7B7eedcF61Eb863Ff413bB6234;
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](1);
        receipts[0] = TeleporterMessageReceipt({
            receivedMessageNonce: expectedNonce,
            relayerRewardAddress: relayerRewardAddress
        });
        _receiveTestMessage(DEFAULT_DESTINATION_BLOCKCHAIN_ID, 43, relayerRewardAddress, receipts);

        // Now the message hash should be cleared.
        assertEq(teleporterMessenger.getMessageHash(messageID), bytes32(0));
    }

    function testMessageDoesNotExist() public view {
        assertEq(teleporterMessenger.getMessageHash(bytes32(uint256(42))), bytes32(0));
    }
}
