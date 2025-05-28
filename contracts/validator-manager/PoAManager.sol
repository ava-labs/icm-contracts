// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IPoAManager} from "./interfaces/IPoAManager.sol";
import {IValidatorManager} from "./interfaces/IValidatorManager.sol";
import {PChainOwner} from "./interfaces/IACP99Manager.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";

/**
 * @dev Implementation of the {IPoAManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract PoAManager is IPoAManager, Initializable, OwnableUpgradeable {
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.PoAManager
    struct PoAManagerStorage {
        IValidatorManager _manager;
    }

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.PoAManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant POA_MANAGER_STORAGE_LOCATION =
        0xafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab600;

    // solhint-disable ordering
    /**
     * @dev This storage is visible to child contracts for convenience.
     *      External getters would be better practice, but code size limitations are preventing this.
     *      Child contracts should probably never write to this storage.
     */
    function _getPoAManagerStorage() internal pure returns (PoAManagerStorage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := POA_MANAGER_STORAGE_LOCATION
        }
    }

    function initialize(address owner, address validatorManager) external initializer {
        __PoAManager_init(owner, validatorManager);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __PoAManager_init(address owner, address validatorManager) internal onlyInitializing {
        __Ownable_init(owner);
        __PoAManager_init_unchained(validatorManager);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __PoAManager_init_unchained(
        address validatorManager
    ) internal onlyInitializing {
        PoAManagerStorage storage $ = _getPoAManagerStorage();
        $._manager = IValidatorManager(validatorManager);
    }

    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) external onlyOwner returns (bytes32) {
        PoAManagerStorage storage $ = _getPoAManagerStorage();
        return $._manager.initiateValidatorRegistration(
            nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, weight
        );
    }

    function initiateValidatorRemoval(
        bytes32 validationID
    ) external onlyOwner {
        PoAManagerStorage storage $ = _getPoAManagerStorage();
        return $._manager.initiateValidatorRemoval(validationID);
    }

    function initiateValidatorWeightUpdate(
        bytes32 validationID,
        uint64 newWeight
    ) external onlyOwner returns (uint64, bytes32) {
        PoAManagerStorage storage $ = _getPoAManagerStorage();
        return $._manager.initiateValidatorWeightUpdate(validationID, newWeight);
    }

    function completeValidatorRegistration(
        uint32 messageIndex
    ) external returns (bytes32) {
        PoAManagerStorage storage $ = _getPoAManagerStorage();
        return $._manager.completeValidatorRegistration(messageIndex);
    }

    function completeValidatorRemoval(
        uint32 messageIndex
    ) external returns (bytes32) {
        PoAManagerStorage storage $ = _getPoAManagerStorage();
        return $._manager.completeValidatorRemoval(messageIndex);
    }

    function completeValidatorWeightUpdate(
        uint32 messageIndex
    ) external returns (bytes32, uint64) {
        PoAManagerStorage storage $ = _getPoAManagerStorage();
        return $._manager.completeValidatorWeightUpdate(messageIndex);
    }
}
