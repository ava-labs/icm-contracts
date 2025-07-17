// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {
    TeleporterMessenger,
    TeleporterMessage,
    TeleporterMessageInput,
    TeleporterMessageReceipt,
    TeleporterFeeInfo,
    IWarpMessenger,
    WarpMessage
} from "../TeleporterMessenger.sol";

// TeleporterMessengerTest and its child contracts all initialize the blockchain ID in the set up
// step for ease of use. This test contract is separate to allow for testing the initialization
// step itself.
contract InitializeBlockchainIDTest is Test {
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_DESTINATION_BLOCKCHAIN_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");

    TeleporterMessenger public teleporterMessenger;

    // Set up the Teleporter contract but leave the blockchain ID uninitialized.
    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(DEFAULT_DESTINATION_BLOCKCHAIN_ID)
        );

        teleporterMessenger = new TeleporterMessenger();
    }

    function testUninitialized() public {
        assertEq(teleporterMessenger.blockchainID(), bytes32(0));
        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        teleporterMessenger.getNextMessageID(DEFAULT_DESTINATION_BLOCKCHAIN_ID);
    }

    function testSuccess() public {
        bytes32 res = teleporterMessenger.initializeBlockchainID();
        assertEq(res, DEFAULT_DESTINATION_BLOCKCHAIN_ID);
        assertEq(teleporterMessenger.blockchainID(), DEFAULT_DESTINATION_BLOCKCHAIN_ID);

        // Should be able to call multiple times and get same result.
        res = teleporterMessenger.initializeBlockchainID();
        assertEq(res, DEFAULT_DESTINATION_BLOCKCHAIN_ID);
    }

    function testInitializedBySending() public {
        // Send a message.
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(bytes32(0))
        );
        teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                destinationAddress: address(1),
                feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
                requiredGasLimit: 1000,
                allowedRelayerAddresses: new address[](0),
                message: new bytes(0)
            })
        );

        // Check the blockchain ID is initialized.
        assertEq(teleporterMessenger.blockchainID(), DEFAULT_DESTINATION_BLOCKCHAIN_ID);
    }

    function testInitializedByReceiving() public {
        TeleporterMessage memory teleporterMessage = TeleporterMessage({
            messageNonce: 1,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: address(1),
            requiredGasLimit: 100,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: new bytes(0)
        });
        WarpMessage memory warpMessage = WarpMessage({
            sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            originSenderAddress: address(teleporterMessenger),
            payload: abi.encode(teleporterMessage)
        });
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, (0)),
            abi.encode(warpMessage, true)
        );
        teleporterMessenger.receiveCrossChainMessage(0, address(0));

        // Check the blockchain ID is initialized.
        assertEq(teleporterMessenger.blockchainID(), DEFAULT_DESTINATION_BLOCKCHAIN_ID);
    }

    function testCannotInitializeToZero() public {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(bytes32(0))
        );

        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        teleporterMessenger.initializeBlockchainID();
    }
}
