// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {
    TeleporterMessengerTest,
    TeleporterMessage,
    TeleporterMessageReceipt,
    WarpMessage
} from "./TeleporterMessengerTest.t.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

contract RedeemRelayerRewardsTest is TeleporterMessengerTest {
    struct FeeRewardInfo {
        uint256 feeAmount;
        address relayerRewardAddress;
    }

    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testZeroRewardBalance() public {
        vm.expectRevert(_formatTeleporterErrorMessage("no reward to redeem"));
        teleporterMessenger.redeemRelayerRewards(address(_mockFeeAsset));
    }

    function testRedemptionFails() public {
        FeeRewardInfo memory feeRewardInfo =
            FeeRewardInfo(2222222222222222, 0xeF6ed43EB8Ff15E336D64d1468947cA1046824E6);
        _setUpRelayerRewards(feeRewardInfo);

        vm.mockCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transfer, (feeRewardInfo.relayerRewardAddress, feeRewardInfo.feeAmount)
            ),
            abi.encode(false)
        );
        vm.expectCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transfer, (feeRewardInfo.relayerRewardAddress, feeRewardInfo.feeAmount)
            )
        );
        vm.prank(feeRewardInfo.relayerRewardAddress);
        vm.expectRevert(
            abi.encodeWithSelector(
                SafeERC20.SafeERC20FailedOperation.selector, address(_mockFeeAsset)
            )
        );
        teleporterMessenger.redeemRelayerRewards(address(_mockFeeAsset));

        // Check that the relayer still has redeemable balance since the transfer failed.
        assertEq(
            teleporterMessenger.checkRelayerRewardAmount(
                feeRewardInfo.relayerRewardAddress, address(_mockFeeAsset)
            ),
            feeRewardInfo.feeAmount
        );
    }

    function testRedemptionSucceeds() public {
        FeeRewardInfo memory feeRewardInfo =
            FeeRewardInfo(2222222222222222, 0xeF6ed43EB8Ff15E336D64d1468947cA1046824E6);
        _setUpRelayerRewards(feeRewardInfo);

        vm.mockCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transfer, (feeRewardInfo.relayerRewardAddress, feeRewardInfo.feeAmount)
            ),
            abi.encode(true)
        );
        vm.expectCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transfer, (feeRewardInfo.relayerRewardAddress, feeRewardInfo.feeAmount)
            )
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit RelayerRewardsRedeemed(
            feeRewardInfo.relayerRewardAddress, address(_mockFeeAsset), feeRewardInfo.feeAmount
        );
        vm.prank(feeRewardInfo.relayerRewardAddress);
        teleporterMessenger.redeemRelayerRewards(address(_mockFeeAsset));

        // Check that the relayer redeemable balance is now 0.
        assertEq(
            teleporterMessenger.checkRelayerRewardAmount(
                feeRewardInfo.relayerRewardAddress, address(_mockFeeAsset)
            ),
            0
        );
    }

    // Mocks sending a message with the given fee info to another L1, and then
    // receiving back a message with receipt of that message such that the relayer
    // is able to redeem the reward.
    function _setUpRelayerRewards(
        FeeRewardInfo memory feeRewardInfo
    ) private {
        uint256 messageNonce = _getNextMessageNonce();
        _sendTestMessageWithFee(DEFAULT_SOURCE_BLOCKCHAIN_ID, feeRewardInfo.feeAmount);

        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](1);
        receipts[0] = TeleporterMessageReceipt({
            receivedMessageNonce: messageNonce,
            relayerRewardAddress: feeRewardInfo.relayerRewardAddress
        });
        TeleporterMessage memory messageToReceive = _createMockTeleporterMessage(1, new bytes(0));
        bytes32 receivedMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            messageToReceive.messageNonce
        );

        messageToReceive.receipts = receipts;
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));

        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the mock message.
        address expectedRelayerRewardAddress = 0x93753a9eA4C9D6eeed9f64eA92E97ce1f5FBAeDe;
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiveCrossChainMessage(
            receivedMessageID,
            warpMessage.sourceChainID,
            address(this),
            expectedRelayerRewardAddress,
            messageToReceive
        );
        teleporterMessenger.receiveCrossChainMessage(0, expectedRelayerRewardAddress);

        // Check that the relayer has redeemable balance
        assertEq(
            teleporterMessenger.checkRelayerRewardAmount(
                feeRewardInfo.relayerRewardAddress, address(_mockFeeAsset)
            ),
            feeRewardInfo.feeAmount
        );

        // Check that the message received is considered delivered, and that the relayer reward address is stored.
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(receivedMessageID),
            expectedRelayerRewardAddress
        );
        assertTrue(teleporterMessenger.messageReceived(receivedMessageID));
    }
}
