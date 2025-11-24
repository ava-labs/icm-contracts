// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {TeleporterMessengerTest, TeleporterMessageReceipt} from "./TeleporterMessengerTest.t.sol";

contract MessageReceivedTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testReceivedMessage() public {
        // Mock receiving a message from another L1.
        address relayerRewardAddress = 0xA66884fAdC0D4d7B7eedcF61Eb863Ff413bB6234;
        uint256 messageNonce = 1;
        bytes32 messageID = teleporterMessenger.calculateMessageID(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, messageNonce
        );
        _receiveTestMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            messageNonce,
            relayerRewardAddress,
            new TeleporterMessageReceipt[](0)
        );

        assertEq(teleporterMessenger.messageReceived(messageID), true);
    }

    function testUnreceivedMessage() public view {
        assertEq(
            teleporterMessenger.messageReceived(
                teleporterMessenger.calculateMessageID(
                    DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, 1
                )
            ),
            false
        );
    }
}
