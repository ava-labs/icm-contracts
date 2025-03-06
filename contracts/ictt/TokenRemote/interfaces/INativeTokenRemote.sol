// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {INativeTokenTransferrer} from "../../interfaces/INativeTokenTransferrer.sol";
import {ITokenRemote} from "./ITokenRemote.sol";

/**
 * @dev Interface that describes functionalities for a contract that can mint native tokens when
 * paired with a {ITokenHome} contract that will lock tokens on another chain.
 */
interface INativeTokenRemote is ITokenRemote, INativeTokenTransferrer {
    /**
     * @notice Emitted when reporting burned tx fees to home chain.
     */
    event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned);

    /**
     * @notice Reports the increase in total burned transaction fees on this chain to its corresponding TokenHome instance.
     * @param requiredGasLimit The gas limit required to report the burned transaction fees to the home contract.
     */
    function reportBurnedTxFees(
        uint256 requiredGasLimit
    ) external;

    /**
     * @notice Returns a best-estimate (upper bound) of the supply of the native asset
     * in circulation on this chain. Does not account for other native asset burn mechanisms,
     * which can result in the value returned being greater than true circulating supply.
     */
    function totalNativeAssetSupply() external view returns (uint256);
}
