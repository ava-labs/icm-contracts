// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {TeleporterMessengerTest, TeleporterMessageReceipt} from "./TeleporterMessengerTest.t.sol";

contract GetFeeInfoTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // First submit a message with a fee
        uint256 feeAmount = 1687435413;
        bytes32 messageID = _sendTestMessageWithFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, feeAmount);

        // Get the fee info to make sure it is correct.
        (address actualFeeAsset, uint256 actualFeeAmount) =
            teleporterMessenger.getFeeInfo(messageID);
        assertEq(actualFeeAsset, address(_mockFeeAsset));
        assertEq(actualFeeAmount, feeAmount);
    }

    function testFeeOnTransferTokenUsed() public {
        // First submit a message with a fee, but mock their being a "fee on token transfer"
        // during the call to transferFrom.
        uint256 feeAmount = 1687435413;
        uint256 tokenTransferFee = 35413;
        _mockFeeAsset.setFeeOnTransferSender(address(this), tokenTransferFee);
        bytes32 messageID = _sendTestMessageWithFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, feeAmount);

        // Get the fee info to make sure it is correct, including the fee amount being less than
        // the amount specified when submitting the message due to the "fee on token transfer".
        (address actualFeeAsset, uint256 actualFeeAmount) =
            teleporterMessenger.getFeeInfo(messageID);
        assertEq(actualFeeAsset, address(_mockFeeAsset));
        assertEq(actualFeeAmount, feeAmount - tokenTransferFee);
    }

    function testAfterReceipt() public {
        // First submit a message with a small fee
        uint256 feeAmount = 10;
        uint256 expectedNonce = _getNextMessageNonce();
        bytes32 messageID = _sendTestMessageWithFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, feeAmount);

        // Now mock receiving a message back from that L1 with a receipt of the above message.
        address relayerRewardAddress = 0xA66884fAdC0D4d7B7eedcF61Eb863Ff413bB6234;
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](1);
        receipts[0] = TeleporterMessageReceipt({
            receivedMessageNonce: expectedNonce,
            relayerRewardAddress: relayerRewardAddress
        });
        _receiveTestMessage(DEFAULT_DESTINATION_BLOCKCHAIN_ID, 6, relayerRewardAddress, receipts);

        // Now, if we get the fee info for the message it should be reported as zero since the receipt has already been received.
        (address actualFeeAsset, uint256 actualFeeAmount) =
            teleporterMessenger.getFeeInfo(messageID);
        assertEq(actualFeeAsset, address(0));
        assertEq(actualFeeAmount, 0);
    }

    function testInvalidMessage() public view {
        bytes32 fakeMessageID = bytes32(uint256(4646));

        // Get the fee info to make sure it is zero since the message doesn't exist.
        (address actualFeeAsset, uint256 actualFeeAmount) =
            teleporterMessenger.getFeeInfo(fakeMessageID);
        assertEq(actualFeeAsset, address(0));
        assertEq(actualFeeAmount, 0);
    }
}
