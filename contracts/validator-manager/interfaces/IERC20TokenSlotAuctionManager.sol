// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IValidatorManager} from "../interfaces/IValidatorManager.sol";

import {PChainOwner} from "contracts/validator-manager/interfaces/IACP99Manager.sol";

import {IERC20Mintable} from "./IERC20Mintable.sol";

interface IERC20TokenSlotAuctionManager {

    /**
     * @notice Places a bid in the currently running auction
     * @param bid The amount of tokens to bid.
     * @param nodeID The ID of the node to add to the L1.
     * @param blsPublicKey The BLS public key of the validator.
     * @param remainingBalanceOwner The remaining balance owner of the validator.
     * @param disableOwner The disable owner of the validator.
     */
    function placeBid (
        uint256 bid,
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) external;

    function erc20() external view returns (IERC20Mintable);
}
