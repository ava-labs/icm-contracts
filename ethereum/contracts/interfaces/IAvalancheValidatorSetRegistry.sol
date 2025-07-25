// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Validator, ValidatorSet} from "../utils/ValidatorSets.sol";
import {ICMMessage} from "../utils/ICM.sol";

/**
 * @title IAvalancheValidatorSetRegistry
 * @notice Interface for an Avalanche Validator Registry contract
 * @dev This interface defines the events and functions for managing Avalanche validators and validating signed message
 * from a given validator set.
 */
interface IAvalancheValidatorSetRegistry {
    event ValidatorSetRegistered(
        uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID
    );
    event ValidatorSetUpdated(
        uint256 indexed validatorSetID, bytes32 indexed avalancheBlockchainID
    );

    /**
     * @notice Registers a new validator set
     * @dev A validator set can be registered by anyone, and its correctness should be verified with the actual validator set.
     * registered on the Avalanche P-Chain. Updates to any validator set after registration are always authenticated by a
     * signed ICM message from the current validator set, giving the party that registered the validator set no elevated
     * permission after registration. Because any given validator set cannot be confirmed to be correct at time of registration
     * on chain, a validator set for a given blockchain ID is allowed to be  be registered multiple times, and they are
     * identified by a unique ID assigned by the registry, rather than the blockchain ID they claim to be
     * representing.
     * @param message The ICM message containing the validator set to register. The message must be signed by validator set
     * that it claims to be representing.
     * @return The ID of the registered validator set
     */
    function registerValidatorSet(
        ICMMessage calldata message,
        bytes memory validatorBytes
    ) external returns (uint256);

    /**
     * @notice Updates a validator set
     * @dev Updates to any validator set after registration are always authenticated by a signed ICM message
     * from the current validator set, giving the party that registered the validator set no elevated
     * permission after registration.
     * @param validatorSetID The ID of the validator set to update
     * @param message The ICM message containing the update
     */
    function updateValidatorSet(
        uint256 validatorSetID,
        ICMMessage calldata message,
        bytes memory validatorBytes
    ) external;

    /**
     * @notice Gets a validator set by its ID
     * @param validatorSetID The ID of the validator set to get
     * @return The validator set
     */
    function getValidatorSet(
        uint256 validatorSetID
    ) external view returns (ValidatorSet memory);

    /**
     * @notice Gets the Avalanche network ID
     * @return The Avalanche network ID
     */
    function getAvalancheNetworkID() external view returns (uint32);

    /**
     * @notice Verifies an ICM message against a validator set
     * @dev This function validates that the message is properly signed by a sufficient quorum of validators
     * from the validator set identified by validatorSetID. The verification includes checking the network ID,
     * blockchain ID, and cryptographic signature verification.
     * @param validatorSetID The ID of the validator set to verify the message against
     * @param message The ICM message to verify
     * @return True if the message is valid, false otherwise
     */
    function verifyICMMessage(
        uint256 validatorSetID,
        ICMMessage calldata message
    ) external view returns (bool);
}
