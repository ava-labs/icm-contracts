// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IPoAManager} from "./interfaces/IPoAManager.sol";
import {IValidatorManagerExternalOwnable} from "./interfaces/IValidatorManagerExternalOwnable.sol";
import {PChainOwner} from "./interfaces/IACP99Manager.sol";
import {Ownable} from "@openzeppelin/contracts@5.0.2/access/Ownable.sol";

/**
 * @dev Implementation of the {IPoAManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract PoAManager is IPoAManager, Ownable {
    IValidatorManagerExternalOwnable private immutable _manager;

    constructor(address owner, IValidatorManagerExternalOwnable validatorManager) Ownable(owner) {
        _manager = validatorManager;
    }

    /**
     * @notice See {IPoAManager-initiateValidatorRegistration}.
     */
    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) external onlyOwner returns (bytes32) {
        return _manager.initiateValidatorRegistration(
            nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, weight
        );
    }

    /**
     * @notice See {IPoAManager-initiateValidatorRemoval}.
     */
    function initiateValidatorRemoval(
        bytes32 validationID
    ) external onlyOwner {
        return _manager.initiateValidatorRemoval(validationID);
    }

    /**
     * @notice See {IPoAManager-initiateValidatorWeightUpdate}.
     */
    function initiateValidatorWeightUpdate(
        bytes32 validationID,
        uint64 newWeight
    ) external onlyOwner returns (uint64, bytes32) {
        return _manager.initiateValidatorWeightUpdate(validationID, newWeight);
    }

    /**
     * @notice See {IPoAManager-completeValidatorRegistration}.
     */
    function completeValidatorRegistration(
        uint32 messageIndex
    ) external returns (bytes32) {
        return _manager.completeValidatorRegistration(messageIndex);
    }

    /**
     * @notice See {IPoAManager-completeValidatorRemoval}.
     */
    function completeValidatorRemoval(
        uint32 messageIndex
    ) external returns (bytes32) {
        return _manager.completeValidatorRemoval(messageIndex);
    }

    /**
     * @notice See {IPoAManager-completeValidatorWeightUpdate}.
     */
    function completeValidatorWeightUpdate(
        uint32 messageIndex
    ) external returns (bytes32, uint64) {
        return _manager.completeValidatorWeightUpdate(messageIndex);
    }

    /**
     * @notice See {IPoAManager-transferValidatorManagerOwnership}.
     */
    function transferValidatorManagerOwnership(
        address newOwner
    ) external onlyOwner {
        _manager.transferOwnership(newOwner);
    }
}
