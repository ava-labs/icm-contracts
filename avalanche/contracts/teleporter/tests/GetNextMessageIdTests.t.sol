// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {
    TeleporterMessengerTest,
    TeleporterMessageInput,
    TeleporterFeeInfo,
    IWarpMessenger
} from "./TeleporterMessengerTest.t.sol";

contract GetNextMessageIDTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testGetMessageID() public {
        // Generate the next expected message ID manually.
        bytes32 expectedMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            _getNextMessageNonce()
        );

        // Check the contract reports the same as expected.
        assertEq(
            teleporterMessenger.getNextMessageID(DEFAULT_DESTINATION_BLOCKCHAIN_ID),
            expectedMessageID
        );

        // Send a message to ensure it is assigned the expected ID.
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(bytes32(0))
        );
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: address(0),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: uint256(0)}),
            requiredGasLimit: 1e6,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        bytes32 messageID = teleporterMessenger.sendCrossChainMessage(messageInput);
        assertEq(messageID, expectedMessageID);

        // Generate the next expected message ID now that a message has been sent.
        bytes32 secondExpectedMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            _getNextMessageNonce()
        );

        // Check the contract reports the same as expected, and that is different than the first ID.
        assertEq(
            teleporterMessenger.getNextMessageID(DEFAULT_DESTINATION_BLOCKCHAIN_ID),
            secondExpectedMessageID
        );
        assertNotEq(expectedMessageID, secondExpectedMessageID);
    }
}
