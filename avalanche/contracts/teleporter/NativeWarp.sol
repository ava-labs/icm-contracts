// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {
    WarpMessage,
    WarpBlockHash,
    IWarpMessenger
} from "@subnet-evm/IWarpMessenger.sol";
import {ICMMessage} from "./ITeleporterMessenger.sol";
import {IWarpExt} from "./IWarpExt.sol";

error UnsupportedMethod(ICMMessage icmMessage, string reason);

contract WarpNative is IWarpExt {
    /**
     * @notice Warp precompile used for sending and receiving Warp messages.
     */
    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    function sendWarpMessage(bytes calldata payload) external returns (bytes32 messageID) {
        return WARP_MESSENGER.sendWarpMessage(payload);
    }

    function getVerifiedICMMessage(
        ICMMessage calldata icmMessage
    ) external view returns (WarpMessage memory) {
        revert UnsupportedMethod({
            icmMessage: icmMessage,
            reason: "This method is not supported by this contract"
        });
    }

    function getVerifiedWarpMessage(uint32 index) external view returns (WarpMessage memory message, bool valid) {
        return WARP_MESSENGER.getVerifiedWarpMessage(index);
    }


    function getVerifiedWarpBlockHash(
        uint32 index
    ) external view returns (WarpBlockHash memory warpBlockHash, bool valid) {
        return WARP_MESSENGER.getVerifiedWarpBlockHash(index);
    }

    function getBlockchainID() external view returns (bytes32 blockchainID) {
        return WARP_MESSENGER.getBlockchainID();
    }

}