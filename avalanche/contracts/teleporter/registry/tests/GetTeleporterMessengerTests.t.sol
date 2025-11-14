// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {BaseTeleporterRegistryAppTest} from "./BaseTeleporterRegistryAppTests.t.sol";
import {ITeleporterMessenger} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterMessenger} from "@teleporter/TeleporterMessenger.sol";

abstract contract GetTeleporterMessengerTest is BaseTeleporterRegistryAppTest {
    function testGetPausedTeleporterMessenger() public {
        _pauseTeleporterAddressSuccess(app, teleporterAddress);
        vm.expectRevert(_formatErrorMessage("Teleporter sending paused"));
        app.getTeleporterMessenger();
    }

    function testGetUnpausedTeleporterMessenger() public {
        _pauseTeleporterAddressSuccess(app, teleporterAddress);
        vm.expectRevert(_formatErrorMessage("Teleporter sending paused"));
        app.getTeleporterMessenger();

        // Unpause the Teleporter address, and now we should getTeleporterMessenger successfully
        _unpauseTeleporterAddressSuccess(app, teleporterAddress);
        app.getTeleporterMessenger();
    }

    function testGetNewTeleporterMessenger() public {
        // Pause the current latest version of Teleporter

        _pauseTeleporterAddressSuccess(app, teleporterAddress);
        vm.expectRevert(_formatErrorMessage("Teleporter sending paused"));
        app.getTeleporterMessenger();

        // Add a new version of Teleporter, and make sure we can get
        // the new Teleporter successfully.
        address newTeleporterAddress = address(new TeleporterMessenger());
        _addProtocolVersion(teleporterRegistry, newTeleporterAddress);
        ITeleporterMessenger messenger = app.getTeleporterMessenger();
        assertEq(address(messenger), newTeleporterAddress);
        assertEq(
            teleporterRegistry.getVersionFromAddress(address(messenger)),
            teleporterRegistry.latestVersion()
        );
    }

    function testPauseNonLatestTeleporter() public {
        address newTeleporterAddress = address(new TeleporterMessenger());
        _addProtocolVersion(teleporterRegistry, newTeleporterAddress);

        // Pause a non-latest version of Teleporter
        _pauseTeleporterAddressSuccess(app, teleporterAddress);

        // Make sure we can still get the latest version of Teleporter for sending
        ITeleporterMessenger messenger = app.getTeleporterMessenger();
        assertEq(address(messenger), newTeleporterAddress);
        assertEq(
            teleporterRegistry.getVersionFromAddress(address(messenger)),
            teleporterRegistry.latestVersion()
        );
    }

    function testGetTeleporterMessengerBasic() public view {
        ITeleporterMessenger messenger = app.getTeleporterMessenger();
        assertEq(
            teleporterRegistry.getVersionFromAddress(address(messenger)),
            teleporterRegistry.latestVersion()
        );
    }
}
