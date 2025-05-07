// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {ERC20TokenRemoteUpgradeable} from "./ERC20TokenRemoteUpgradeable.sol";
import {TokenRemoteSettings} from "./interfaces/ITokenRemote.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";

/**
 * @title ERC20TokenRemote
 * @notice A non-upgradeable version of {ERC20TokenRemoteUpgradeable} that calls the parent upgradeable contract's initialize function.
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract ERC20TokenRemote is ERC20TokenRemoteUpgradeable {
    constructor(
        TokenRemoteSettings memory settings,
        string memory tokenName,
        string memory tokenSymbol,
        uint8 tokenDecimals
    ) ERC20TokenRemoteUpgradeable(ICMInitializable.Allowed) {
        initialize(settings, tokenName, tokenSymbol, tokenDecimals);
    }
}
