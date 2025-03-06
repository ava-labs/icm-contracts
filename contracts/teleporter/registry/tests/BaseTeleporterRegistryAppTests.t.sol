// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {TeleporterRegistryAppUpgradeable} from "../TeleporterRegistryAppUpgradeable.sol";
import {TeleporterRegistryApp} from "../TeleporterRegistryApp.sol";
import {TeleporterRegistryTest} from "./TeleporterRegistryTests.t.sol";
import {ITeleporterMessenger, TeleporterMessageInput} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterMessenger} from "@teleporter/TeleporterMessenger.sol";
import {UnitTestMockERC20} from "@mocks/UnitTestMockERC20.sol";

contract ExampleRegistryAppUpgradeable is TeleporterRegistryAppUpgradeable {
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

    function sendTeleporterMessage(
        TeleporterMessageInput calldata messageInput
    ) public {
        _sendTeleporterMessage(messageInput);
    }

    function getTeleporterMessenger() public view returns (ITeleporterMessenger) {
        return _getTeleporterMessenger();
    }

    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message // solhint-disable-next-line no-empty-blocks
    ) internal override {}

    // solhint-disable-next-line no-empty-blocks
    function _checkTeleporterRegistryAppAccess() internal override {}
}

contract ExampleRegistryApp is TeleporterRegistryApp {
    constructor(
        address teleporterRegistryAddress,
        uint256 minTeleporterVersion
    ) TeleporterRegistryApp(teleporterRegistryAddress, minTeleporterVersion) {}

    function setMinTeleporterVersion(
        uint256 version
    ) public {
        _setMinTeleporterVersion(version);
    }

    function sendTeleporterMessage(
        TeleporterMessageInput calldata messageInput
    ) public {
        _sendTeleporterMessage(messageInput);
    }

    function getTeleporterMessenger() public view returns (ITeleporterMessenger) {
        return _getTeleporterMessenger();
    }

    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message // solhint-disable-next-line no-empty-blocks
    ) internal override {}

    // solhint-disable-next-line no-empty-blocks
    function _checkTeleporterRegistryAppAccess() internal virtual override {}
}

abstract contract BaseTeleporterRegistryAppTest is TeleporterRegistryTest {
    ExampleRegistryApp public app;
    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_DESTINATION_BLOCKCHAIN_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    address public constant DEFAULT_DESTINATION_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant DEFAULT_ORIGIN_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;

    UnitTestMockERC20 internal _mockFeeAsset;

    event MinTeleporterVersionUpdated(
        uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion
    );

    event TeleporterAddressPaused(address indexed teleporterAddress);

    event TeleporterAddressUnpaused(address indexed teleporterAddress);

    function setUp() public virtual override {
        TeleporterRegistryTest.setUp();
        _addProtocolVersion(teleporterRegistry, teleporterAddress);
        _mockFeeAsset = new UnitTestMockERC20();
        TeleporterMessenger(teleporterAddress).initializeBlockchainID();
    }

    function testOnlyAllowedTeleporter() public {
        assertEq(app.getMinTeleporterVersion(), 1);

        vm.expectRevert(_formatRegistryErrorMessage("protocol address not found"));
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testUpdateMinTeleporterVersion() public {
        // First check that calling with initial teleporter address works
        assertEq(app.getMinTeleporterVersion(), 1);
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        // Now add new protocol version to registry and update the app's min version
        address newTeleporterAddress = address(new TeleporterMessenger());
        _addProtocolVersion(teleporterRegistry, newTeleporterAddress);

        _updateMinTeleporterVersionSuccess(app, teleporterRegistry.latestVersion());
        assertEq(app.getMinTeleporterVersion(), 2);

        // Check that calling with the old teleporter address fails
        vm.expectRevert(_formatErrorMessage("invalid Teleporter sender"));
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        // Check that calling with the new teleporter address works
        vm.prank(newTeleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testSetMinTeleporterVersion() public {
        uint256 latestVersion = teleporterRegistry.latestVersion();

        // Check setting for a version > latest version fails
        vm.expectRevert(_formatErrorMessage("invalid Teleporter version"));
        app.setMinTeleporterVersion(latestVersion + 1);

        // Check setting for a version <= min version fails
        uint256 minVersion = app.getMinTeleporterVersion();
        assertEq(minVersion, teleporterRegistry.latestVersion());
        vm.expectRevert(_formatErrorMessage("not greater than current minimum version"));
        app.setMinTeleporterVersion(minVersion);

        // Add a new protocol version to the registry
        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        // Check setting a new valid minimum version
        vm.expectEmit(true, true, true, true, address(app));
        emit MinTeleporterVersionUpdated(latestVersion, latestVersion + 1);
        app.setMinTeleporterVersion(teleporterRegistry.latestVersion());
        assertEq(app.getMinTeleporterVersion(), latestVersion + 1);
    }

    function _updateMinTeleporterVersionSuccess(
        TeleporterRegistryApp app_,
        uint256 newMinTeleporterVersion
    ) internal virtual {
        vm.expectEmit(true, true, true, true, address(app_));
        emit MinTeleporterVersionUpdated(app_.getMinTeleporterVersion(), newMinTeleporterVersion);
        app_.updateMinTeleporterVersion(newMinTeleporterVersion);
    }

    function _pauseTeleporterAddressSuccess(
        TeleporterRegistryApp app_,
        address teleporterAddress_
    ) internal virtual {
        vm.expectEmit(true, true, true, true, address(app_));
        emit TeleporterAddressPaused(teleporterAddress_);
        app_.pauseTeleporterAddress(teleporterAddress_);
    }

    function _unpauseTeleporterAddressSuccess(
        TeleporterRegistryApp app_,
        address teleporterAddress_
    ) internal virtual {
        vm.expectEmit(true, true, true, true, address(app_));
        emit TeleporterAddressUnpaused(teleporterAddress_);
        app_.unpauseTeleporterAddress(teleporterAddress_);
    }

    function _formatErrorMessage(
        bytes memory errorMessage
    ) internal pure returns (bytes memory) {
        return abi.encodePacked("TeleporterRegistryApp: ", errorMessage);
    }
}
