// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IStakingManager} from "./IStakingManager.sol";
import {PChainOwner} from "../ACP99Manager.sol";

/**
 * Proof of Stake Validator Manager that stakes the blockchain's native tokens.
 */
interface INativeTokenStakingManager is IStakingManager {
    /**
     * @notice Begins the validator registration process. Locks the provided native asset in the contract as the stake.
     * @param nodeID The ID of the node to add to the L1.
     * @param blsPublicKey The BLS public key of the validator.
     * @param remainingBalanceOwner The remaining balance owner of the validator.
     * @param disableOwner The disable owner of the validator.
     * @param delegationFeeBips The fee that delegators must pay to delegate to this validator.
     * @param minStakeDuration The minimum amount of time this validator must be staked for in seconds.
     * @param rewardRecipient The address of the reward recipient.
     * @return validationID The ID of the registered validator.
     */
    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        address rewardRecipient
    ) external payable returns (bytes32);

    /**
     * @notice Begins the delegator registration process. Locks the provided native asset in the contract as the stake.
     * @param validationID The ID of the validator to stake to.
     * @param rewardRecipient The address of the reward recipient.
     * @return delegationID The ID of the registered delegator.
     */
    function initiateDelegatorRegistration(
        bytes32 validationID,
        address rewardRecipient
    ) external payable returns (bytes32);
}
