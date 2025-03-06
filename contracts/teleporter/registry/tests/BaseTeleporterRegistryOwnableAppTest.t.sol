// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {TeleporterRegistryOwnableAppUpgradeable} from
    "../TeleporterRegistryOwnableAppUpgradeable.sol";
import {TeleporterRegistryOwnableApp} from "../TeleporterRegistryOwnableApp.sol";
import {TeleporterRegistryApp} from "../TeleporterRegistryApp.sol";
import {BaseTeleporterRegistryAppTest} from "./BaseTeleporterRegistryAppTests.t.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {ITeleporterMessenger, TeleporterMessageInput} from "@teleporter/ITeleporterMessenger.sol";

contract ExampleRegistryOwnableAppUpgradeable is TeleporterRegistryOwnableAppUpgradeable {
    function initialize(
        address teleporterRegistryAddress,
        address initialOwner,
        uint256 minTeleporterVersion
    ) public initializer {
        __TeleporterRegistryOwnableApp_init(
            teleporterRegistryAddress, initialOwner, minTeleporterVersion
        );
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

    function checkTeleporterRegistryAppAccess() public view {
        _checkTeleporterRegistryAppAccess();
    }

    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message // solhint-disable-next-line no-empty-blocks
    ) internal override {}
}

contract ExampleRegistryOwnableApp is TeleporterRegistryOwnableApp {
    constructor(
        address teleporterRegistryAddress,
        address initialOwner,
        uint256 minTeleporterVersion
    ) TeleporterRegistryOwnableApp(teleporterRegistryAddress, initialOwner, minTeleporterVersion) {}

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

    function checkTeleporterRegistryAppAccess() public view {
        _checkTeleporterRegistryAppAccess();
    }

    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message // solhint-disable-next-line no-empty-blocks
    ) internal override {}
}

abstract contract BaseTeleporterRegistryOwnableAppTest is BaseTeleporterRegistryAppTest {
    ExampleRegistryOwnableApp public ownerApp;
    address public constant MOCK_INVALID_OWNER_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant DEFAULT_OWNER_ADDRESS = 0x1234512345123451234512345123451234512345;

    function setUp() public virtual override {
        BaseTeleporterRegistryAppTest.setUp();
    }

    function testOwnerUpdateMinTeleporterVersion() public {
        uint256 minTeleporterVersion = ownerApp.getMinTeleporterVersion();
        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        // Check that call to update minimum Teleporter version reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, MOCK_INVALID_OWNER_ADDRESS
            )
        );
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);

        // Check that minimum Teleporter version was not updated
        assertEq(ownerApp.getMinTeleporterVersion(), minTeleporterVersion);

        // Check that call to update minimum Teleporter version succeeds for owners
        _updateMinTeleporterVersionSuccess(ownerApp, minTeleporterVersion + 1);

        assertEq(ownerApp.getMinTeleporterVersion(), minTeleporterVersion + 1);
    }

    function testOwnerTransfer() public {
        uint256 minTeleporterVersion = ownerApp.getMinTeleporterVersion();
        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        // Check that call to transfer ownership reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, MOCK_INVALID_OWNER_ADDRESS
            )
        );
        ownerApp.transferOwnership(address(0));

        // Check that ownership was not transferred
        assertEq(ownerApp.owner(), DEFAULT_OWNER_ADDRESS);

        // Check that call for non owners reverts
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, MOCK_INVALID_OWNER_ADDRESS
            )
        );
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);

        // Check that after ownership transfer call succeeds
        address newOwner = address(0x123);
        vm.prank(DEFAULT_OWNER_ADDRESS);
        ownerApp.transferOwnership(newOwner);
        vm.prank(newOwner);
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);

        // Check that call with old owner reverts
        vm.prank(DEFAULT_OWNER_ADDRESS);
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, DEFAULT_OWNER_ADDRESS
            )
        );
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);
    }

    function testRenounceOwnership() public {
        // Check that call to renounce ownership reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, MOCK_INVALID_OWNER_ADDRESS
            )
        );
        ownerApp.renounceOwnership();

        // Check that ownership was not renounced
        assertEq(ownerApp.owner(), DEFAULT_OWNER_ADDRESS);
        vm.prank(DEFAULT_OWNER_ADDRESS);
        _addProtocolVersion(teleporterRegistry, teleporterAddress);
        uint256 latestVersion = teleporterRegistry.latestVersion();
        // Check that update Teleporter version call for owner succeeds
        _updateMinTeleporterVersionSuccess(ownerApp, latestVersion);

        // Check that after ownership renounce call reverts
        vm.prank(DEFAULT_OWNER_ADDRESS);
        ownerApp.renounceOwnership();

        vm.prank(DEFAULT_OWNER_ADDRESS);
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, DEFAULT_OWNER_ADDRESS
            )
        );
        ownerApp.updateMinTeleporterVersion(latestVersion);
    }

    function testPauseTeleporterAccess() public {
        // First pause the Teleporter address
        vm.prank(DEFAULT_OWNER_ADDRESS);
        _pauseTeleporterAddressSuccess(ownerApp, teleporterAddress);

        // Try to unpause the Teleporter address from non-owner account
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, MOCK_INVALID_OWNER_ADDRESS
            )
        );
        ownerApp.unpauseTeleporterAddress(teleporterAddress);

        // Check that the Teleporter address is still paused
        vm.prank(teleporterAddress);
        vm.expectRevert(_formatErrorMessage("Teleporter address paused"));
        ownerApp.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        // Unpause the Teleporter address from owner account
        vm.prank(DEFAULT_OWNER_ADDRESS);
        _unpauseTeleporterAddressSuccess(ownerApp, teleporterAddress);

        // Check that the Teleporter address can now deliver messages
        vm.prank(teleporterAddress);
        ownerApp.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testOwnerUpgradeAccess() public {
        // Check that call to check upgrade access reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, MOCK_INVALID_OWNER_ADDRESS
            )
        );
        ownerApp.checkTeleporterRegistryAppAccess();

        // Check that call to check upgrade access succeeds for owners
        vm.prank(DEFAULT_OWNER_ADDRESS);
        ownerApp.checkTeleporterRegistryAppAccess();
    }

    function testInitalOwner() public {
        // Create a new Ownable app with a passed in teleporterManager
        ExampleRegistryOwnableAppUpgradeable newOwnerApp =
            new ExampleRegistryOwnableAppUpgradeable();
        newOwnerApp.initialize(
            address(teleporterRegistry), DEFAULT_OWNER_ADDRESS, teleporterRegistry.latestVersion()
        );

        // Check that the teleporterManager is set correctly
        assertEq(newOwnerApp.owner(), DEFAULT_OWNER_ADDRESS);
        vm.prank(DEFAULT_OWNER_ADDRESS);
        newOwnerApp.checkTeleporterRegistryAppAccess();

        // Check that address(this) as the caller is not by default owner
        assertFalse(newOwnerApp.owner() == address(this));
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, address(this)
            )
        );
        newOwnerApp.checkTeleporterRegistryAppAccess();
    }

    function _updateMinTeleporterVersionSuccess(
        TeleporterRegistryApp app_,
        uint256 newMinTeleporterVersion
    ) internal virtual override {
        vm.expectEmit(true, true, true, true, address(app_));
        emit MinTeleporterVersionUpdated(app_.getMinTeleporterVersion(), newMinTeleporterVersion);
        vm.prank(DEFAULT_OWNER_ADDRESS);
        app_.updateMinTeleporterVersion(newMinTeleporterVersion);
    }
}
