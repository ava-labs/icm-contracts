// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {BaseTeleporterRegistryAppTest} from "./BaseTeleporterRegistryAppTests.t.sol";
import {TeleporterMessenger} from "@teleporter/TeleporterMessenger.sol";

abstract contract PauseTeleporterAddressTest is BaseTeleporterRegistryAppTest {
    function testPauseTeleporterAddressBasic() public {
        // Check that the teleporterAddress is not paused initially
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
        // Check that teleporterAddress can not deliver messages once paused
        _pauseTeleporterAddressSuccess(app, teleporterAddress);
        vm.expectRevert(_formatErrorMessage("Teleporter address paused"));
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testAlreadyPausedTeleporterAddress() public {
        // Check that teleporterAddress can not deliver messages once paused
        _pauseTeleporterAddressSuccess(app, teleporterAddress);
        vm.expectRevert(_formatErrorMessage("Teleporter address paused"));
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
        // Check that teleporterAddress can not be paused again
        vm.expectRevert(_formatErrorMessage("address already paused"));
        app.pauseTeleporterAddress(teleporterAddress);
    }

    function testPauseMultipleVersions() public {
        // Since a Teleporter address can be registered with multiple versions,
        // when pausing a Teleporter address, we should pause all versions of that address.

        // Add the same Teleporter address to the registry with a different version
        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        // Check that the teleporterAddress is not paused initially
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        // Check that teleporterAddress can not deliver messages once paused
        _pauseTeleporterAddressSuccess(app, teleporterAddress);
        vm.expectRevert(_formatErrorMessage("Teleporter address paused"));
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        // Check that after updating mininum Teleporter version, the address is still paused
        _updateMinTeleporterVersionSuccess(app, teleporterRegistry.latestVersion());
        vm.prank(teleporterAddress);
        vm.expectRevert(_formatErrorMessage("Teleporter address paused"));
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testPauseBeforeRegister() public {
        // Check that a Teleporter address can be paused before it is registered with the registry

        // Create a new Teleporter address that is not registered with the registry
        address newTeleporterAddress = address(new TeleporterMessenger());

        // Pause the new Teleporter address before it is registered
        _pauseTeleporterAddressSuccess(app, newTeleporterAddress);

        // Register the new Teleporter address
        _addProtocolVersion(teleporterRegistry, newTeleporterAddress);

        // Check that the new Teleporter address is paused
        vm.prank(newTeleporterAddress);
        vm.expectRevert(_formatErrorMessage("Teleporter address paused"));
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testPauseZeroAddress() public {
        // Check that a zero address can not be paused
        vm.expectRevert(_formatErrorMessage("zero Teleporter address"));
        app.pauseTeleporterAddress(address(0));
    }
}
