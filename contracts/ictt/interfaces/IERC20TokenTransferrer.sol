// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {ITokenTransferrer, SendTokensInput, SendAndCallInput} from "./ITokenTransferrer.sol";

/**
 * @notice Interface for an Avalanche interchain token transferrer that sends ERC20 tokens to another chain.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
interface IERC20TokenTransferrer is ITokenTransferrer {
    /**
     * @notice Sends ERC20 tokens to the specified destination.
     * @param input Specifies information for delivery of the tokens
     * @param amount Amount of tokens to send
     */
    function send(SendTokensInput calldata input, uint256 amount) external;

    /**
     * @notice Sends ERC20 tokens to the specified destination to be used in a smart contract interaction.
     * @param input Specifies information for delivery of the tokens
     * @param amount Amount of tokens to send
     */
    function sendAndCall(SendAndCallInput calldata input, uint256 amount) external;
}
