// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {
    TeleporterMessengerTest,
    TeleporterMessage,
    TeleporterFeeInfo,
    TeleporterMessageInput,
    IWarpMessenger,
    IERC20
} from "./TeleporterMessengerTest.t.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

contract SendCrossChainMessageTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSendMessageNoFee() public {
        // Arrange
        uint256 expectedMessageNonce = _getNextMessageNonce();
        bytes32 expectedMessageID = teleporterMessenger.calculateMessageID(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            expectedMessageNonce
        );
        TeleporterMessage memory expectedMessage =
            _createMockTeleporterMessage(expectedMessageNonce, hex"deadbeef");
        TeleporterFeeInfo memory feeInfo = TeleporterFeeInfo(address(0), 0);
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: expectedMessage.destinationAddress,
            feeInfo: feeInfo,
            requiredGasLimit: expectedMessage.requiredGasLimit,
            allowedRelayerAddresses: expectedMessage.allowedRelayerAddresses,
            message: expectedMessage.message
        });

        // We have to mock the precompile call so that the test does not revert.
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(bytes32(0))
        );

        // Expect the exact message to be passed to the precompile.
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(expectedMessage)))
        );

        // Expect the SendCrossChainMessage event to be emitted.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            expectedMessageID, messageInput.destinationBlockchainID, expectedMessage, feeInfo
        );

        // Act
        teleporterMessenger.sendCrossChainMessage(messageInput);

        // Check receipt queue
        uint256 queueSize =
            teleporterMessenger.getReceiptQueueSize(messageInput.destinationBlockchainID);
        assertEq(queueSize, 0);
    }

    function testSendMessageWithFee() public {
        // Arrange
        // Construct the message to submit.
        uint256 expectedMessageNonce = _getNextMessageNonce();
        TeleporterMessage memory expectedMessage =
            _createMockTeleporterMessage(expectedMessageNonce, hex"deadbeef");
        TeleporterFeeInfo memory feeInfo =
            TeleporterFeeInfo(address(_mockFeeAsset), 13131313131313131313);
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: expectedMessage.destinationAddress,
            feeInfo: feeInfo,
            requiredGasLimit: expectedMessage.requiredGasLimit,
            allowedRelayerAddresses: expectedMessage.allowedRelayerAddresses,
            message: expectedMessage.message
        });

        // We have to mock the precompile call so that the test does not revert.
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(bytes32(0))
        );

        // Expect the exact message to be passed to the precompile.
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(expectedMessage)))
        );

        // Expect the SendCrossChainMessage event to be emitted.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            teleporterMessenger.calculateMessageID(
                DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                expectedMessageNonce
            ),
            messageInput.destinationBlockchainID,
            expectedMessage,
            feeInfo
        );

        // Expect the ERC20 contract transferFrom method to be called to transfer the fee.
        vm.expectCall(
            messageInput.feeInfo.feeTokenAddress,
            abi.encodeCall(
                IERC20.transferFrom,
                (address(this), address(teleporterMessenger), messageInput.feeInfo.amount)
            )
        );

        // Act
        teleporterMessenger.sendCrossChainMessage(messageInput);
    }

    function testFeeAssetDoesNotExist() public {
        address invalidFeeAsset = 0xb8be9140D8717f4a8fd7e8ae23C5668bc3A4B39c;
        uint256 feeAmount = 4567;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(invalidFeeAsset, feeAmount),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        // We expect a call to the fee asset address, but since we did not
        // mock it there is no code there to execute.
        vm.expectCall(
            invalidFeeAsset, abi.encodeCall(IERC20.balanceOf, (address(teleporterMessenger)))
        );
        vm.expectRevert();
        teleporterMessenger.sendCrossChainMessage(messageInput);
    }

    function testFeeTransferFailure() public {
        uint256 feeAmount = 4567;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(address(_mockFeeAsset), feeAmount),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        vm.mockCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom, (address(this), address(teleporterMessenger), feeAmount)
            ),
            abi.encode(false)
        );
        vm.expectCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom, (address(this), address(teleporterMessenger), feeAmount)
            )
        );
        vm.expectRevert(
            abi.encodeWithSelector(
                SafeERC20.SafeERC20FailedOperation.selector, address(_mockFeeAsset)
            )
        );
        teleporterMessenger.sendCrossChainMessage(messageInput);
    }

    function testInvalidFeeAssetFailure() public {
        address invalidFeeAsset = address(0);
        uint256 feeAmount = 4567;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(invalidFeeAsset, feeAmount),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        vm.expectRevert(_formatTeleporterErrorMessage("zero fee asset contract address"));
        teleporterMessenger.sendCrossChainMessage(messageInput);
    }
}
