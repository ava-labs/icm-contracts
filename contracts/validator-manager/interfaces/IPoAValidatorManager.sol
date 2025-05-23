// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {PChainOwner, ValidatorStatus, IACP99Manager} from "./IACP99Manager.sol";

/**
 * @dev Proof-of-authority validator manager interface meant for to be used as the owner of an
 * {IValidatorManager} contract instance that only allows modifications to the validator set
 * to be initiated by a fixed address. Once initiated, modifications are meant to be able to be
 * completed by any address to allow for imrpoved UX flows.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
interface IPoAValidatorManager {
    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) external returns (bytes32);

    function initiateValidatorRemoval(bytes32 validationID) external;

    function initiateValidatorWeightUpdate(
        bytes32 validationID,
        uint64 newWeight
    ) external returns (uint64, bytes32);

    function completeValidatorRegistration(uint32 messageIndex) external returns (bytes32);

    function completeValidatorRemoval(uint32 messageIndex) external returns (bytes32);

    function completeValidatorWeightUpdate(uint32 messageIndex)
        external
        returns (bytes32, uint64);
}
