// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {BaseTeleporterRegistryAppTest} from "./BaseTeleporterRegistryAppTests.t.sol";
import {TeleporterMessenger} from "@teleporter/TeleporterMessenger.sol";

abstract contract UnpauseTeleporterAddressTest is BaseTeleporterRegistryAppTest {
    function testUnpauseLessThanMinimumVersion() public {
        // Check the case where a Teleporter address was previously paused, and now updates the minimum version.
        // If the dApp unpauses the previous Teleporter address, it still can not receive messages from it,
        // since receiving messages still checks against minimum Teleporter version.

        // Pause the Teleporter address
        _pauseTeleporterAddressSuccess(app, teleporterAddress);

        // The Teleporter address paused no longer can deliver messages to the app
        vm.expectRevert(_formatErrorMessage("Teleporter address paused"));
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        // Add a new Teleporter address to the registry and update minimum version
        address newTeleporterAddress = address(new TeleporterMessenger());
        _addProtocolVersion(teleporterRegistry, newTeleporterAddress);
        _updateMinTeleporterVersionSuccess(app, teleporterRegistry.latestVersion());

        // Unpause the previously paused Teleporter address
        _unpauseTeleporterAddressSuccess(app, teleporterAddress);

        // The previously paused Teleporter address still can not deliver messages to the app
        // because the minimum version is still greater than the Teleporter address version
        vm.expectRevert(_formatErrorMessage("invalid Teleporter sender"));
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        // Check that the new Teleporter address delivers messages fine
        vm.prank(newTeleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testUnpauseTeleporterAddressBasic() public {
        // Check that the teleporterAddress is not paused initially
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
        // Check that teleporterAddress can not deliver messages once paused
        _pauseTeleporterAddressSuccess(app, teleporterAddress);
        vm.expectRevert(_formatErrorMessage("Teleporter address paused"));
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        // Unpause the teleporterAddress and check that it can deliver messages again
        _unpauseTeleporterAddressSuccess(app, teleporterAddress);
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testAlreadyUnpausedTeleporterAddress() public {
        // Check unpausing for an address that was never paused
        vm.expectRevert(_formatErrorMessage("address not paused"));
        app.unpauseTeleporterAddress(teleporterAddress);

        // Pause the Teleporter address
        _pauseTeleporterAddressSuccess(app, teleporterAddress);

        // Try to unpause the Teleporter address first time should succeed
        _unpauseTeleporterAddressSuccess(app, teleporterAddress);

        // Try to unpause the Teleporter address second time should fail
        vm.expectRevert(_formatErrorMessage("address not paused"));
        app.unpauseTeleporterAddress(teleporterAddress);
    }

    function testUnpauseBeforeRegister() public {
        // Check that a Teleporter address can be paused before it is registered with the registry

        // Create a new Teleporter address that is not registered with the registry
        address newTeleporterAddress = address(new TeleporterMessenger());

        // Pause the new Teleporter address before it is registered
        _pauseTeleporterAddressSuccess(app, newTeleporterAddress);

        // Unpause before the Teleporter address is registered
        _unpauseTeleporterAddressSuccess(app, newTeleporterAddress);

        // Register the new Teleporter address
        _addProtocolVersion(teleporterRegistry, newTeleporterAddress);

        // Check that the new Teleporter address is unpaused
        vm.prank(newTeleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testUnpauseAfterRegister() public {
        // Check that a Teleporter address can be paused before it is registered with the registry

        // Create a new Teleporter address that is not registered with the registry
        address newTeleporterAddress = address(new TeleporterMessenger());

        // Pause the new Teleporter address before it is registered
        _pauseTeleporterAddressSuccess(app, newTeleporterAddress);

        // Register the new Teleporter address
        _addProtocolVersion(teleporterRegistry, newTeleporterAddress);

        // Unpause before the Teleporter address is registered
        _unpauseTeleporterAddressSuccess(app, newTeleporterAddress);

        // Check that the new Teleporter address is unpaused
        vm.prank(newTeleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testUnpauseZeroAddress() public {
        // Check that a zero address can not be paused
        vm.expectRevert(_formatErrorMessage("zero Teleporter address"));
        app.unpauseTeleporterAddress(address(0));
    }
}
