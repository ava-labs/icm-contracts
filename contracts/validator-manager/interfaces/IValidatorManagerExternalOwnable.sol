// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity 0.8.25;

import {IValidatorManager} from "./IValidatorManager.sol";

/**
 * @notice Combines the {IValidatorManager} interface with the external ownership transfer function from {Ownable}.
 * @dev This interface is intended for contracts that need to interact with both validator management and ownership transfer
 *      without importing both {IValidatorManager} and {Ownable} separately. Note that {Ownable} is a contract, not an interface,
 *      so this interface only exposes the external ownership transfer method ({transferOwnership}) for convenience.
 */
interface IValidatorManagerExternalOwnable is IValidatorManager {
    /**
     * @notice Transfers ownership of the validator manager to a new owner.
     * @param newOwner The address of the new owner.
     * @dev This function can only be called by the current owner.
     */
    function transferOwnership(
        address newOwner
    ) external;
}
