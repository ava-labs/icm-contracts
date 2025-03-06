// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */
import {ExampleERC20} from "@mocks/ExampleERC20.sol";

contract ExampleERC20Decimals is ExampleERC20 {
    uint8 public immutable tokenDecimals;

    constructor(
        uint8 tokenDecimals_
    ) ExampleERC20() {
        tokenDecimals = tokenDecimals_;
    }

    function decimals() public view virtual override returns (uint8) {
        return tokenDecimals;
    }
}
