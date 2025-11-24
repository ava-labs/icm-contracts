// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {TeleporterMessengerTest, TeleporterMessageReceipt} from "./TeleporterMessengerTest.t.sol";

contract GetRelayerRewardAddressTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // Receive a message
        uint256 mockNonce = 8;
        bytes32 mockMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, mockNonce
        );
        address relayerRewardAddress = 0xCAFebAbeDc0D4D7B7EEdCf61eb863fF413BB6234;
        _receiveTestMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            mockNonce,
            relayerRewardAddress,
            new TeleporterMessageReceipt[](0)
        );

        // Now it has the relayer reward address.
        assertEq(teleporterMessenger.getRelayerRewardAddress(mockMessageID), relayerRewardAddress);
    }

    function testZeroRelayerRewardAddress() public {
        // Receive a message with a zero relayer reward address
        uint256 mockNonce = 4343;
        bytes32 mockMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, mockNonce
        );
        _receiveTestMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, mockNonce, address(0), new TeleporterMessageReceipt[](0)
        );

        // Check that the zero address is returned as the reward address.
        assertEq(teleporterMessenger.getRelayerRewardAddress(mockMessageID), address(0));
    }

    function testMessageNotReceived() public {
        // Before receiving the given message, getRelayerRewardAddress should revert.
        bytes32 mockMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, 4242
        );
        vm.expectRevert(_formatTeleporterErrorMessage("message not received"));
        teleporterMessenger.getRelayerRewardAddress(mockMessageID);
    }
}
