// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {TeleporterMessenger} from "../TeleporterMessenger.sol";
import {Test} from "@forge-std/Test.sol";

contract CheckIsAllowedRelayerTest is TeleporterMessenger, Test {
    function testIsSpecifiedAllowedRelayer() public pure {
        address relayerAddress = 0x6288dAdf62B57dd9A4ddcd02F88A98d0eb6c2598;
        address[] memory allowedRelayers = new address[](3);
        allowedRelayers[0] = relayerAddress;
        allowedRelayers[1] = 0xDeaDBeEf62B57DD9a4ddcD02F88a98d0eb6C2598;
        allowedRelayers[2] = 0xfFfFfFFffFB57Dd9A4ddcD02f88A98D0Eb6c2598;
        assertTrue(_checkIsAllowedRelayer(relayerAddress, allowedRelayers));
    }

    function testAnyRelayerIsAllowed() public pure {
        address[] memory relayerAddresses = new address[](3);
        relayerAddresses[0] = 0x6288dAdf62B57dd9A4ddcd02F88A98d0eb6c2598;
        relayerAddresses[1] = 0xDeaDBeEf62B57DD9a4ddcD02F88a98d0eb6C2598;
        relayerAddresses[2] = 0xfFfFfFFffFB57Dd9A4ddcD02f88A98D0Eb6c2598;
        address[] memory allowedRelayers = new address[](0);

        for (uint256 i; i < relayerAddresses.length; ++i) {
            assertTrue(_checkIsAllowedRelayer(relayerAddresses[i], allowedRelayers));
        }
    }

    function testUnauthorizedRelayer() public pure {
        address relayerAddress = 0x6288dAdf62B57dd9A4ddcd02F88A98d0eb6c2598;
        address[] memory allowedRelayers = new address[](3);
        allowedRelayers[0] = 0xCaFEbabE62B57DD9A4dDcD02F88A98d0Eb6c2598;
        allowedRelayers[1] = 0xDeaDBeEf62B57DD9a4ddcD02F88a98d0eb6C2598;
        allowedRelayers[2] = 0xfFfFfFFffFB57Dd9A4ddcD02f88A98D0Eb6c2598;

        assertFalse(_checkIsAllowedRelayer(relayerAddress, allowedRelayers));
    }
}
