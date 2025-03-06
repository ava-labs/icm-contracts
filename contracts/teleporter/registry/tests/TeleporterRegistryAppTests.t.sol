// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {GetTeleporterMessengerTest} from "./GetTeleporterMessengerTests.t.sol";
import {PauseTeleporterAddressTest} from "./PauseTeleporterAddressTests.t.sol";
import {UpdateMinTeleporterVersionTest} from "./UpdateMinTeleporterVersionTests.t.sol";
import {SendTeleporterMessageTest} from "./SendTeleporterMessageTests.t.sol";
import {UnpauseTeleporterAddressTest} from "./UnpauseTeleporterAddressTests.t.sol";
import {BaseTeleporterRegistryAppTest} from "./BaseTeleporterRegistryAppTests.t.sol";
import {
    ExampleRegistryApp,
    ExampleRegistryAppUpgradeable
} from "./BaseTeleporterRegistryAppTests.t.sol";
import {NonReentrantTest} from "./NonReentrantTests.t.sol";
import {TeleporterRegistry, ProtocolRegistryEntry} from "../TeleporterRegistry.sol";

contract TeleporterRegistryAppTest is
    GetTeleporterMessengerTest,
    PauseTeleporterAddressTest,
    UpdateMinTeleporterVersionTest,
    SendTeleporterMessageTest,
    UnpauseTeleporterAddressTest,
    NonReentrantTest
{
    function setUp() public virtual override (BaseTeleporterRegistryAppTest, NonReentrantTest) {
        BaseTeleporterRegistryAppTest.setUp();
        app =
            new ExampleRegistryApp(address(teleporterRegistry), teleporterRegistry.latestVersion());
        NonReentrantTest.setUp();
    }

    function testZeroRegistryAddress() public virtual {
        vm.expectRevert(_formatErrorMessage("zero Teleporter registry address"));
        app = new ExampleRegistryApp(address(0), 0);
    }

    function testInvalidRegistryAddress() public virtual {
        // Create a new Teleporter registry with no registered Teleporters
        TeleporterRegistry teleporterRegistry =
            new TeleporterRegistry(new ProtocolRegistryEntry[](0));
        vm.expectRevert(_formatErrorMessage("invalid Teleporter registry"));
        app = new ExampleRegistryApp(address(teleporterRegistry), 0);
    }

    function testGreaterThanLatestVersion() public virtual {
        uint256 minTeleporterVersion = teleporterRegistry.latestVersion() + 1;
        vm.expectRevert(_formatErrorMessage("invalid Teleporter version"));
        app = new ExampleRegistryApp(address(teleporterRegistry), minTeleporterVersion);
    }

    function testZeroMinTeleporterVersion() public virtual {
        vm.expectRevert(_formatErrorMessage("not greater than current minimum version"));
        app = new ExampleRegistryApp(address(teleporterRegistry), 0);
    }
}

contract TeleporterRegistryAppUpgradeableTest is TeleporterRegistryAppTest {
    function setUp() public virtual override {
        TeleporterRegistryAppTest.setUp();
        ExampleRegistryAppUpgradeable upgradeableApp = new ExampleRegistryAppUpgradeable();
        upgradeableApp.initialize(address(teleporterRegistry), teleporterRegistry.latestVersion());
        app = ExampleRegistryApp(address(upgradeableApp));
    }

    function testZeroRegistryAddress() public override {
        ExampleRegistryAppUpgradeable upgradeableApp = new ExampleRegistryAppUpgradeable();
        vm.expectRevert(_formatErrorMessage("zero Teleporter registry address"));
        upgradeableApp.initialize(address(0), 0);
    }

    function testInvalidRegistryAddress() public override {
        ExampleRegistryAppUpgradeable upgradeableApp = new ExampleRegistryAppUpgradeable();

        // Create a new Teleporter registry with no registered Teleporters
        TeleporterRegistry teleporterRegistry =
            new TeleporterRegistry(new ProtocolRegistryEntry[](0));
        vm.expectRevert(_formatErrorMessage("invalid Teleporter registry"));
        upgradeableApp.initialize(address(teleporterRegistry), 0);
    }

    function testGreaterThanLatestVersion() public override {
        ExampleRegistryAppUpgradeable upgradeableApp = new ExampleRegistryAppUpgradeable();

        uint256 minTeleporterVersion = teleporterRegistry.latestVersion() + 1;
        vm.expectRevert(_formatErrorMessage("invalid Teleporter version"));
        upgradeableApp.initialize(address(teleporterRegistry), minTeleporterVersion);
    }

    function testZeroMinTeleporterVersion() public override {
        ExampleRegistryAppUpgradeable upgradeableApp = new ExampleRegistryAppUpgradeable();

        vm.expectRevert(_formatErrorMessage("not greater than current minimum version"));
        upgradeableApp.initialize(address(teleporterRegistry), 0);
    }

    function testStorageSlot() public {
        assertEq(
            _erc7201StorageSlot("TeleporterRegistryApp"),
            new ExampleRegistryAppUpgradeable().TELEPORTER_REGISTRY_APP_STORAGE_LOCATION()
        );
    }

    function _erc7201StorageSlot(
        bytes memory storageName
    ) private pure returns (bytes32) {
        return keccak256(
            abi.encode(uint256(keccak256(abi.encodePacked("teleporter.storage.", storageName))) - 1)
        ) & ~bytes32(uint256(0xff));
    }
}
