// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {TeleporterRegistryAppUpgradeable} from "../TeleporterRegistryAppUpgradeable.sol";
import {BaseTeleporterRegistryAppTest} from "./BaseTeleporterRegistryAppTests.t.sol";
import {
    ITeleporterMessenger,
    TeleporterMessage,
    TeleporterMessageReceipt
} from "@teleporter/ITeleporterMessenger.sol";
import {WarpMessage} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {TeleporterMessenger} from "@teleporter/TeleporterMessenger.sol";

uint32 constant warpMessageIndex = 2;

contract NonReentrantUpgradeableApp is TeleporterRegistryAppUpgradeable {
    function initialize(
        address teleporterRegistryAddress,
        uint256 minTeleporterVersion
    ) public initializer {
        __TeleporterRegistryApp_init(teleporterRegistryAddress, minTeleporterVersion);
    }

    function setMinTeleporterVersion(
        uint256 version
    ) public {
        _setMinTeleporterVersion(version);
    }

    function getTeleporterMessenger() public view returns (ITeleporterMessenger) {
        return _getTeleporterMessenger();
    }

    // Calls receiveCrossChainMessage on the latest TeleporterMessenger Contract.
    // The Warp Precompile is mocked to return a message that will call
    // TeleporterRegistryAppUpgradeable.receiveTeleporterMessage which should revert because it is
    // non-reentrant.
    function _receiveTeleporterMessage(bytes32, address, bytes memory) internal override {
        // Call `receiveCrossChainMessage` of the latest version of Teleporter
        getTeleporterMessenger().receiveCrossChainMessage(warpMessageIndex, address(this));
    }

    // solhint-disable-next-line no-empty-blocks
    function _checkTeleporterRegistryAppAccess() internal override {}
}

// The flow for the tests below is as follows:
// NonreentrantUpgradeableApp::receiveTeleporterMessage ->
// NonreentrantUpgradeableApp::_receiveTeleporterMessage ->
// TeleporterMessenger::receiveCrossChainMessage ->
// NonreentrantUpgradeableApp::receiveTeleporterMessage
// The last step should revert because receiveTeleporterMessage (contained in
// TeleporterRegistryAppUpgradeable) is non-reentrant.
abstract contract NonReentrantTest is BaseTeleporterRegistryAppTest {
    bytes public constant DEFAULT_MESSAGE = bytes(hex"1234");
    uint256 public constant DEFAULT_REQUIRED_GAS_LIMIT = 1e6;

    NonReentrantUpgradeableApp public nonReentrantApp;

    event MessageExecutionFailed(
        bytes32 indexed messageID, bytes32 indexed originBlockchainID, TeleporterMessage message
    );

    event MessageExecuted(bytes32 indexed messageID, bytes32 indexed originBlockchainID);

    function setUp() public virtual override {
        nonReentrantApp = new NonReentrantUpgradeableApp();
        nonReentrantApp.initialize(address(teleporterRegistry), teleporterRegistry.latestVersion());
    }

    function testNonReentrantSameTeleporter() public {
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageNonce: 987,
            originSenderAddress: DEFAULT_ORIGIN_ADDRESS,
            destinationBlockchainID: MOCK_BLOCK_CHAIN_ID,
            destinationAddress: address(nonReentrantApp),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: DEFAULT_MESSAGE
        });
        WarpMessage memory warpMessage = WarpMessage({
            sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            originSenderAddress: teleporterAddress,
            payload: abi.encode(messageToReceive)
        });

        // Same index as in NonreentrantUpgradeableApp._receiveTeleporterMessage()
        _mockGetVerifiedWarpMessage(warpMessageIndex, warpMessage, true);

        bytes32 messageID = TeleporterMessenger(teleporterAddress).calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, MOCK_BLOCK_CHAIN_ID, messageToReceive.messageNonce
        );

        vm.expectEmit(true, true, true, true, address(teleporterAddress));
        emit MessageExecutionFailed(messageID, DEFAULT_SOURCE_BLOCKCHAIN_ID, messageToReceive);

        vm.prank(teleporterAddress);
        nonReentrantApp.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, DEFAULT_MESSAGE
        );
    }

    function testNonReentrantDifferentTeleporter() public {
        TeleporterMessenger teleporterV2 = new TeleporterMessenger();
        teleporterV2.initializeBlockchainID();
        _addProtocolVersion(teleporterRegistry, address(teleporterV2));

        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageNonce: 987,
            originSenderAddress: DEFAULT_ORIGIN_ADDRESS,
            destinationBlockchainID: MOCK_BLOCK_CHAIN_ID,
            destinationAddress: address(nonReentrantApp),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: DEFAULT_MESSAGE
        });
        WarpMessage memory warpMessage = WarpMessage({
            sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            originSenderAddress: address(teleporterV2),
            payload: abi.encode(messageToReceive)
        });

        // Same index as in NonreentrantUpgradeableApp._receiveTeleporterMessage()
        _mockGetVerifiedWarpMessage(warpMessageIndex, warpMessage, true);

        vm.expectCall(
            address(teleporterV2),
            abi.encodeCall(
                ITeleporterMessenger.receiveCrossChainMessage,
                (warpMessageIndex, address(nonReentrantApp))
            )
        );

        bytes32 messageID = TeleporterMessenger(teleporterV2).calculateMessageID(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, MOCK_BLOCK_CHAIN_ID, messageToReceive.messageNonce
        );

        vm.expectEmit(true, true, true, true, address(teleporterV2));
        emit MessageExecutionFailed(messageID, DEFAULT_SOURCE_BLOCKCHAIN_ID, messageToReceive);

        vm.prank(teleporterAddress);
        nonReentrantApp.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, DEFAULT_MESSAGE
        );
    }
}
