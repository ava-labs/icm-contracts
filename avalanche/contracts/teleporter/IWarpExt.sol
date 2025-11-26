// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.25;

import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/IWarpMessenger.sol";
import {ICMMessage} from "./ITeleporterMessenger.sol";

/**
 * @dev Interface that allows adapting the Warp interface. This is necessary for external interoperability
 * since external chains do not receive Warp messages in their access lists.
 */
interface IWarpExt is IWarpMessenger {
    /**
     * @notice Depending on the chain this contract is deployed on, dispatch logic for
     * getting the actual verified Warp message.
     * @return message A verified Warp message.
     */
    function getVerifiedICMMessage(
        ICMMessage calldata icmMessage
    ) external view returns (WarpMessage memory message);
}