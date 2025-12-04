// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {
    TeleporterMessengerTest,
    TeleporterFeeInfo,
    TeleporterMessageReceipt,
    IERC20
} from "./TeleporterMessengerTest.t.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

contract AddFeeAmountTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        bytes32 messageID =
            _sendTestMessageWithFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, originalFeeAmount);

        // Add to the fee
        uint256 additionalFeeAmount = 131313;
        uint256 totalFeeAmount = originalFeeAmount + additionalFeeAmount;
        vm.expectCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom,
                (address(this), address(teleporterMessenger), additionalFeeAmount)
            )
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit AddFeeAmount(
            messageID,
            TeleporterFeeInfo({feeTokenAddress: address(_mockFeeAsset), amount: totalFeeAmount})
        );
        teleporterMessenger.addFeeAmount(messageID, address(_mockFeeAsset), additionalFeeAmount);

        // Get the fee info to make sure it is properly updated.
        (address actualFeeAsset, uint256 actualFeeAmount) =
            teleporterMessenger.getFeeInfo(messageID);
        assertEq(actualFeeAsset, address(_mockFeeAsset));
        assertEq(actualFeeAmount, totalFeeAmount);
    }

    function testInvalidMessage() public {
        // Add to the fee amount of a message that doesn't exist. Expect revert.
        uint256 additionalFeeAmount = 131313;
        bytes32 fakeMessageID = bytes32(uint256(13));
        vm.expectRevert(_formatTeleporterErrorMessage("message not found"));
        teleporterMessenger.addFeeAmount(fakeMessageID, address(_mockFeeAsset), additionalFeeAmount);
    }

    function testMessageAlreadyDelivered() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        uint256 expectedNonce = _getNextMessageNonce();
        bytes32 messageID =
            _sendTestMessageWithFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, originalFeeAmount);

        // Now mock receiving a message back from that L1 with a receipt of the above message.
        address relayerRewardAddress = 0xA66884fAdC0D4d7B7eedcF61Eb863Ff413bB6234;
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](1);
        receipts[0] = TeleporterMessageReceipt({
            receivedMessageNonce: expectedNonce,
            relayerRewardAddress: relayerRewardAddress
        });

        _receiveTestMessage(DEFAULT_DESTINATION_BLOCKCHAIN_ID, 1, relayerRewardAddress, receipts);

        // Now try to add to the fee of the message. Should revert since the message receipt was received already.
        uint256 additionalFeeAmount = 131313;
        vm.expectRevert(_formatTeleporterErrorMessage("message not found"));
        teleporterMessenger.addFeeAmount(messageID, address(_mockFeeAsset), additionalFeeAmount);
    }

    function testInvalidAmount() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        bytes32 messageID =
            _sendTestMessageWithFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, originalFeeAmount);

        // Expect revert when adding 0 additional amount.
        uint256 additionalFeeAmount = 0;
        vm.expectRevert(_formatTeleporterErrorMessage("zero additional fee amount"));
        teleporterMessenger.addFeeAmount(messageID, address(_mockFeeAsset), additionalFeeAmount);
    }

    function testMismatchFeeAsset() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        bytes32 messageID =
            _sendTestMessageWithFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, originalFeeAmount);

        // Expect revert when using a different fee asset than originally used.
        uint256 additionalFeeAmount = 131313;
        address differentFeeAsset = 0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664;
        vm.expectRevert(_formatTeleporterErrorMessage("invalid fee asset contract address"));
        teleporterMessenger.addFeeAmount(messageID, differentFeeAsset, additionalFeeAmount);
    }

    function testInvalidFeeAsset() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        bytes32 messageID =
            _sendTestMessageWithFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, originalFeeAmount);

        // Expect revert when using an invalid fee asset.
        uint256 additionalFeeAmount = 131313;
        address invalidFeeAsset = address(0);
        vm.expectRevert(_formatTeleporterErrorMessage("zero fee asset contract address"));
        teleporterMessenger.addFeeAmount(messageID, invalidFeeAsset, additionalFeeAmount);
    }

    function testInsufficientBalance() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        bytes32 messageID =
            _sendTestMessageWithFee(DEFAULT_DESTINATION_BLOCKCHAIN_ID, originalFeeAmount);

        // Add to the fee, but mock the ERC20 contract returning an error from transferFrom
        uint256 additionalFeeAmount = 131313;
        vm.expectCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom,
                (address(this), address(teleporterMessenger), additionalFeeAmount)
            )
        );
        vm.mockCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom,
                (address(this), address(teleporterMessenger), additionalFeeAmount)
            ),
            abi.encode(false)
        );
        vm.expectRevert(
            abi.encodeWithSelector(
                SafeERC20.SafeERC20FailedOperation.selector, address(_mockFeeAsset)
            )
        );

        teleporterMessenger.addFeeAmount(messageID, address(_mockFeeAsset), additionalFeeAmount);
    }
}
