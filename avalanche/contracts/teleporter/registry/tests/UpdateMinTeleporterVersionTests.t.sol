// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {BaseTeleporterRegistryAppTest} from "./BaseTeleporterRegistryAppTests.t.sol";
import {TeleporterMessenger, WarpMessage} from "@teleporter/TeleporterMessenger.sol";

abstract contract UpdateMinTeleporterVersionTest is BaseTeleporterRegistryAppTest {
    function testMessageDeliveryFromOutdatedVersion() public {
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

    function testUpdateToNonRegisteredVersion() public {
        // The Teleporter registry allows skipping versions, so we don't have to update
        // in increments of 1. This test checks that the app can update to a version
        // that is not registered in the registry.

        // First add to the registry by skipping a version
        address newTeleporterAddress = address(new TeleporterMessenger());
        uint256 latestVersion = teleporterRegistry.latestVersion();
        uint32 messageIndex = 0;
        WarpMessage memory warpMessage = _createWarpOffChainMessage(
            latestVersion + 2, newTeleporterAddress, address(teleporterRegistry)
        );
        _mockGetVerifiedWarpMessage(messageIndex, warpMessage, true);

        vm.expectEmit(true, true, true, true, address(teleporterRegistry));
        emit AddProtocolVersion(latestVersion + 2, newTeleporterAddress);
        vm.expectEmit(true, true, true, true, address(teleporterRegistry));
        emit LatestVersionUpdated(latestVersion, latestVersion + 2);
        teleporterRegistry.addProtocolVersion(messageIndex);

        // Check that the latest version is updated
        assertEq(teleporterRegistry.latestVersion(), latestVersion + 2);

        // Make sure that a version was skipped and is not registered
        uint256 skippedVersion = latestVersion + 1;
        vm.expectRevert(_formatRegistryErrorMessage("version not found"));
        teleporterRegistry.getTeleporterFromVersion(skippedVersion);

        // Update to the skipped version
        _updateMinTeleporterVersionSuccess(app, skippedVersion);
        assertEq(app.getMinTeleporterVersion(), skippedVersion);

        // Make sure that the old minimum Teleporter version can not deliver messages
        vm.expectRevert(_formatErrorMessage("invalid Teleporter sender"));
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        // Make sure that the new minimum Teleporter version can still deliver messages
        vm.prank(newTeleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testUpdateWithCurrentVersion() public {
        // Get the current minimum Teleporter version
        uint256 minTeleporterVersion = app.getMinTeleporterVersion();

        // Check that current latest version is equal to minimum version
        assertEq(minTeleporterVersion, teleporterRegistry.latestVersion());

        // Try to update to current minimum version, should fail
        vm.expectRevert(_formatErrorMessage("not greater than current minimum version"));
        app.updateMinTeleporterVersion(minTeleporterVersion);

        // Check that minimum version is still the same
        assertEq(app.getMinTeleporterVersion(), minTeleporterVersion);

        // Check that calling with the teleporter address works
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testUpdateWithGreaterThanLatestVersion() public {
        // Get current latest version and minimum Teleporter version
        uint256 latestVersion = teleporterRegistry.latestVersion();
        uint256 minTeleporterVersion = app.getMinTeleporterVersion();

        // Try to update to a version greater than the latest version, should fail
        vm.expectRevert(_formatErrorMessage("invalid Teleporter version"));
        app.updateMinTeleporterVersion(latestVersion + 1);

        // Check that minimum version is still the same
        assertEq(app.getMinTeleporterVersion(), minTeleporterVersion);
    }
}
