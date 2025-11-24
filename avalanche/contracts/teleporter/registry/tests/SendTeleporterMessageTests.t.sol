// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {BaseTeleporterRegistryAppTest} from "./BaseTeleporterRegistryAppTests.t.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {IERC20} from "@teleporter/TeleporterMessenger.sol";

abstract contract SendTeleporterMessageTest is BaseTeleporterRegistryAppTest {
    function testSendTeleporterPaused() public {
        // Pause the Teleporter address returned from _getTeleporterMessenger()
        address teleporter = address(app.getTeleporterMessenger());
        _pauseTeleporterAddressSuccess(app, teleporter);

        // Check that the app reverts when trying to send a message with paused Teleporter
        vm.expectRevert(_formatErrorMessage("Teleporter sending paused"));
        app.sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                destinationAddress: DEFAULT_DESTINATION_ADDRESS,
                feeInfo: TeleporterFeeInfo(address(0), 0),
                requiredGasLimit: 0,
                allowedRelayerAddresses: new address[](0),
                message: hex"deadbeef"
            })
        );
    }

    function testSendNoFee() public {
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(address(0), 0),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: hex"deadbeef"
        });
        vm.mockCall(
            address(app.getTeleporterMessenger()),
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (messageInput)),
            abi.encode(bytes32(0))
        );
        vm.expectCall(
            address(app.getTeleporterMessenger()),
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (messageInput))
        );

        app.sendTeleporterMessage(messageInput);
    }

    function testSendWithFee() public {
        uint256 feeAmount = 1;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(address(_mockFeeAsset), feeAmount),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        vm.mockCall(
            address(app.getTeleporterMessenger()),
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (messageInput)),
            abi.encode(bytes32(0))
        );
        vm.expectCall(
            address(_mockFeeAsset),
            abi.encodeCall(IERC20.allowance, (address(app), address(teleporterAddress)))
        );
        vm.expectCall(
            address(app.getTeleporterMessenger()),
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (messageInput))
        );

        app.sendTeleporterMessage(messageInput);
    }

    function testInvalidFeeAsset() public {
        address invalidFeeAsset = 0xb8be9140D8717f4a8fd7e8ae23C5668bc3A4B39c;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(invalidFeeAsset, 1),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        // Expect a call to check for allowance, but because we provide an invalid
        // fee asset address, the call will revert.
        vm.expectCall(
            invalidFeeAsset,
            abi.encodeCall(IERC20.allowance, (address(app), address(teleporterAddress)))
        );
        vm.expectRevert();
        app.sendTeleporterMessage(messageInput);
    }

    function testZeroFeeAsset() public {
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(address(0), 1),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        vm.expectRevert(_formatErrorMessage("zero fee token address"));
        app.sendTeleporterMessage(messageInput);
    }
}
