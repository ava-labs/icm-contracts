// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {
    TeleporterMessenger,
    TeleporterMessage,
    TeleporterMessageReceipt,
    TeleporterMessageInput,
    TeleporterFeeInfo,
    IWarpMessenger,
    WarpMessage
} from "../TeleporterMessenger.sol";
import {UnitTestMockERC20} from "@mocks/UnitTestMockERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";

// Parent contract for TeleporterMessenger tests. Deploys a TeleporterMessenger
// instance in the test setup, and provides helper methods for sending and
// receiving empty mock messages.
contract TeleporterMessengerTest is Test {
    TeleporterMessenger public teleporterMessenger;

    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_DESTINATION_BLOCKCHAIN_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    address public constant DEFAULT_DESTINATION_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    uint256 public constant DEFAULT_REQUIRED_GAS_LIMIT = 1e6;
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
    address public constant DEFAULT_RELAYER_REWARD_ADDRESS =
        0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d;

    UnitTestMockERC20 internal _mockFeeAsset;

    event BlockchainIDInitialized(bytes32 indexed blockchainID);

    event SendCrossChainMessage(
        bytes32 indexed messageID,
        bytes32 indexed destinationBlockchainID,
        TeleporterMessage message,
        TeleporterFeeInfo feeInfo
    );

    event AddFeeAmount(bytes32 indexed messageID, TeleporterFeeInfo updatedFeeInfo);

    event ReceiveCrossChainMessage(
        bytes32 indexed messageID,
        bytes32 indexed sourceBlockchainID,
        address indexed deliverer,
        address rewardRedeemer,
        TeleporterMessage message
    );

    event MessageExecutionFailed(
        bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, TeleporterMessage message
    );

    event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID);

    event ReceiptReceived(
        bytes32 indexed messageID,
        bytes32 indexed destinationBlockchainID,
        address indexed relayerRewardAddress,
        TeleporterFeeInfo feeInfo
    );

    event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount);

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(DEFAULT_DESTINATION_BLOCKCHAIN_ID)
        );

        teleporterMessenger = new TeleporterMessenger();

        // Blockchain ID should be 0 before it is initialized.
        assertEq(teleporterMessenger.blockchainID(), bytes32(0));

        // Initialize the blockchain ID.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit BlockchainIDInitialized(DEFAULT_DESTINATION_BLOCKCHAIN_ID);
        teleporterMessenger.initializeBlockchainID();
        assertEq(teleporterMessenger.blockchainID(), DEFAULT_DESTINATION_BLOCKCHAIN_ID);

        _mockFeeAsset = new UnitTestMockERC20();
    }

    function testEmptyReceiptQueue() public {
        assertEq(teleporterMessenger.getReceiptQueueSize(DEFAULT_DESTINATION_BLOCKCHAIN_ID), 0);

        vm.expectRevert("ReceiptQueue: index out of bounds");
        TeleporterMessageReceipt memory receipt =
            teleporterMessenger.getReceiptAtIndex(DEFAULT_DESTINATION_BLOCKCHAIN_ID, 0);
        assertEq(receipt.receivedMessageNonce, 0);
        assertEq(receipt.relayerRewardAddress, address(0));
    }

    /*
     * TEST UTILS / HELPERS
     */

    function _sendTestMessageWithFee(
        bytes32 blockchainID,
        uint256 feeAmount
    ) internal returns (bytes32) {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(bytes32(0))
        );

        address feeAsset = address(0);
        if (feeAmount > 0) {
            feeAsset = address(_mockFeeAsset);
        }
        TeleporterFeeInfo memory feeInfo =
            TeleporterFeeInfo({feeTokenAddress: feeAsset, amount: feeAmount});

        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: blockchainID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: feeInfo,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        if (feeAmount > 0) {
            // Expect the ERC20 contract transferFrom method to be called to transfer the fee.
            vm.expectCall(
                address(_mockFeeAsset),
                abi.encodeCall(
                    IERC20.transferFrom, (address(this), address(teleporterMessenger), feeAmount)
                )
            );
        }

        return teleporterMessenger.sendCrossChainMessage(messageInput);
    }

    function _sendTestMessageWithNoFee(
        bytes32 blockchainID
    ) internal returns (bytes32) {
        return _sendTestMessageWithFee(blockchainID, 0);
    }

    function _setUpSuccessGetVerifiedWarpMessageMock(
        uint32 index,
        WarpMessage memory warpMessage
    ) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, (index)),
            abi.encode(warpMessage, true)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, (index))
        );
    }

    function _receiveTestMessage(
        bytes32 sourceBlockchainID,
        uint256 messageNonce,
        address relayerRewardAddress,
        TeleporterMessageReceipt[] memory receipts
    ) internal {
        // Construct the test message to be received.
        TeleporterMessage memory messageToReceive =
            _createMockTeleporterMessage(messageNonce, new bytes(0));
        messageToReceive.receipts = receipts;

        // Both the sender and destination address should be the teleporter contract address,
        // mocking as if the universal deployer pattern was used.
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(sourceBlockchainID, abi.encode(messageToReceive));

        // We have to mock the precompile call so that it doesn't revert in the tests.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the message.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiveCrossChainMessage(
            teleporterMessenger.calculateMessageID(
                sourceBlockchainID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, messageNonce
            ),
            warpMessage.sourceChainID,
            address(this),
            relayerRewardAddress,
            messageToReceive
        );
        teleporterMessenger.receiveCrossChainMessage(0, relayerRewardAddress);
    }

    function _receiveMessageSentToEOA()
        internal
        returns (bytes32, address, TeleporterMessage memory)
    {
        // Create valid mock message data, we will mock out the execution of the
        // message itself so it doesn't matter what it is.
        bytes memory messageData = new bytes(1);

        // Construct the test message to be received. By default, the destination
        // address will be the DEFAULT_DESTINATION_ADDRESS.
        uint256 messageNonce = 9;
        TeleporterMessage memory messageToReceive =
            _createMockTeleporterMessage(messageNonce, messageData);
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));
        bytes32 messageID = teleporterMessenger.calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_BLOCKCHAIN_ID, messageNonce
        );

        // We have to mock the precompile call so that it doesn't revert in the tests.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Mock there being no code at the destination address which will be called
        // when the message is executed. This mimics the default destination address being an EOA.
        vm.etch(DEFAULT_DESTINATION_ADDRESS, new bytes(0));

        // Receive the message - which should fail execution.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiveCrossChainMessage(
            messageID,
            warpMessage.sourceChainID,
            address(this),
            DEFAULT_RELAYER_REWARD_ADDRESS,
            messageToReceive
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit MessageExecutionFailed(messageID, DEFAULT_SOURCE_BLOCKCHAIN_ID, messageToReceive);
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);

        return (DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, messageToReceive);
    }

    // Create a mock message to be used in tests. It should include no receipts
    // because receipts are only valid if they refer to messages previously send
    // to the origin chain.
    function _createMockTeleporterMessage(
        uint256 messageNonce,
        bytes memory message
    ) internal view returns (TeleporterMessage memory) {
        return TeleporterMessage({
            messageNonce: messageNonce,
            originSenderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: message
        });
    }

    function _createDefaultWarpMessage(
        bytes32 sourceBlockchainID,
        bytes memory payload
    ) internal view returns (WarpMessage memory) {
        return WarpMessage({
            sourceChainID: sourceBlockchainID,
            originSenderAddress: address(teleporterMessenger),
            payload: payload
        });
    }

    function _getNextMessageNonce() internal view returns (uint256) {
        return teleporterMessenger.messageNonce() + 1;
    }

    function _formatTeleporterErrorMessage(
        bytes memory errorMessage
    ) internal pure returns (bytes memory) {
        return abi.encodePacked("TeleporterMessenger: ", errorMessage);
    }
}
