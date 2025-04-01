// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorManager} from "../ValidatorManager.sol";

/**
 * @dev Delegator status
 */
enum DelegatorStatus {
    Unknown,
    PendingAdded,
    Active,
    PendingRemoved,
    Removed
}

/**
 * @notice PoS Validator Manager settings, used to initialize the PoS Validator Manager
 * @notice baseSettings specified the base settings for the Validator Manager. See {ValidatorManager-ValidatorManagerSettings}
 * @notice minimumStakeAmount is the minimum amount of stake required to stake to a validator
 * @notice maximumStakeAmount is the maximum amount of stake that can be staked to a validator
 * @notice minimumStakeDuration is the minimum duration that validators must stake for
 * @notice minimumDelegationFeeBips is the minimum delegation fee in basis points that validators can charge
 * the maximum amount of stake a validator can have with delegations.
 * @notice weightToValueFactor is the factor used to convert validator weight to value
 * @notice rewardCalculator is the reward calculator used to calculate rewards for this validator manager
 * @notice uptimeBlockchainID is the ID of the blockchain that submits uptime proofs.
 * This must be a blockchain validated by the subnetID that this contract manages.
 */
struct StakingManagerSettings {
    ValidatorManager manager;
    uint256 minimumStakeAmount;
    uint256 maximumStakeAmount;
    uint256 maximumNFTAmount;
    uint64 minimumStakeDuration;
    uint256 minimumDelegationAmount;
    uint16 minimumDelegationFeeBips;
    uint256 weightToValueFactor;
    address admin;
    bytes32 uptimeBlockchainID;
    uint64 unlockDuration;
    uint64 epochDuration;
    address uptimeKeeper;
    uint64 epochOffset;
}

/**
 * @dev Contains the active state of a Delegator
 */
struct Delegator {
    DelegatorStatus status;
    address owner;
    bytes32 validationID;
    uint64 weight;
    uint64 startTime;
    uint64 endTime;
    uint64 startingNonce;
    uint64 endingNonce;
}

/**
 * @dev Describes the active state of a PoS Validator in addition the information in {ValidatorManager-Validator}
 */
struct PoSValidatorInfo {
    address owner;
    uint16 delegationFeeBips;
    uint64 minStakeDuration;
    uint64 uptimeSeconds;
    bytes32[] activeDelegations;
    uint256[] tokenIDs;
    uint64 tokenNonce;
    uint256 totalTokens;
}

/**
 * @notice Interface for Proof of Stake Validator Managers
 */
interface IStakingManager {
    /**
     * @notice Event emitted when a delegator registration is initiated
     * @param delegationID The ID of the delegation
     * @param validationID The ID of the validation period being delegated to
     * @param delegatorAddress The address of the delegator
     * @param nonce The message nonce used to update the validator weight
     * @param validatorWeight The updated validator weight that is sent to the P-Chain
     * @param delegatorWeight The weight of the delegator
     * @param setWeightMessageID The ID of the ICM message that updates the validator's weight on the P-Chain
     */
    event InitiatedDelegatorRegistration(
        bytes32 indexed delegationID,
        bytes32 indexed validationID,
        address indexed delegatorAddress,
        uint64 nonce,
        uint64 validatorWeight,
        uint64 delegatorWeight,
        bytes32 setWeightMessageID
    );

    /**
     * @notice Event emitted when a delegator registration is completed
     * @param delegationID The ID of the delegation
     * @param validationID The ID of the validation period
     * @param startTime The time at which the registration was completed
     */
    event CompletedDelegatorRegistration(
        bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime
    );

    /**
     * @notice Event emitted when delegator removal is initiated
     * @param delegationID The ID of the delegation
     * @param validationID The ID of the validation period the delegator was staked to
     */
    event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID);

    /**
     * @notice Event emitted when delegator removal is completed
     * @param delegationID The ID of the delegation
     * @param validationID The ID of the validator the delegator was staked to
     * @param rewards The rewards given to the delegator
     * @param fees The portion of the delegator's rewards paid to the validator
     */
    event CompletedDelegatorRemoval(
        bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees
    );

    /**
     * @notice Event emitted when the uptime of a validator is updated. Only emitted when the uptime is greater than the stored uptime.
     * @param validationID The ID of the validation period
     * @param uptime The updated uptime of the validator
     */
    event UptimeUpdated(bytes32 indexed validationID, uint64 uptime, uint64 epoch);

    /**
     * @notice Event emitted when the validator stake is unlocked
     * @param validationID The ID of the validation period
     */
    event UnlockedValidation(bytes32 indexed validationID);

    /**
     * @notice Event emitted when the delegator stake is unlocked
     * @param delegationID The ID of the delegation period
     */
    event UnlockedDelegation(bytes32 indexed delegationID);

    /**
     * @notice Event emitted when the delegator reward weight is resolved
     * @param delegationID The ID of the delegation period
     * @param epoch The current reward epoch
     */ 
    event RewardResolved(bytes32 indexed delegationID, uint64 epoch);

    /**
     * @notice Updates the uptime of the validationID if the submitted proof is greated than the stored uptime.
     * Anybody may call this function to ensure the stored uptime is accurate. Callable only when the validation period is active.
     * @param validationID The ID of the validation period
     * @param messageIndex The index of the ICM message to be received providing the uptime proof
     */
    function submitUptimeProof(bytes32 validationID, uint32 messageIndex) external;

    /**
     * @notice Completes validator registration by dispatching to the ValidatorManager to update the validator status,
     * and locking stake.
     *
     * @param messageIndex The index of the ICM message to be received providing the acknowledgement from the P-Chain.
     * This is forwarded to the ValidatorManager to be parsed.
     * @return The ID of the validator that was registered.
     */
    function completeValidatorRegistration(uint32 messageIndex) external returns (bytes32);

    /**
     * @notice Begins the process of ending an active validation period, and reverts if the validation period is not eligible
     * for uptime-based rewards. This function is used to exit the validator set when rewards are expected.
     * The validation period must have been previously started by a successful call to {completeValidatorRegistration} with the given validationID.
     * Any rewards for this validation period will stop accruing when this function is called.
     * Note: Reverts if the uptime is not eligible for rewards.
     * @param validationID The ID of the validation period being ended.
     * @param includeUptimeProof Whether or not an uptime proof is provided for the validation period. If no uptime proof is provided,
     * the latest known uptime will be used.
     * @param messageIndex The index of the ICM message to be received providing the uptime proof.
     */
    function initiateValidatorRemoval(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external;

    /**
     * @notice Completes validator removal by dispatching to the ValidatorManager to update the validator status,
     * and unlocking stake.
     *
     * @param messageIndex The index of the ICM message to be received providing the acknowledgement from the P-Chain.
     * This is forwarded to the ValidatorManager to be parsed.
     * @return The ID of the validator that was removed.
     */
    function completeValidatorRemoval(uint32 messageIndex) external returns (bytes32);

    /**
     * @notice Completes the delegator registration process by submitting an acknowledgement of the registration of a
     * validationID from the P-Chain.
     * Any P-Chain acknowledgement with a nonce greater than or equal to the nonce used to initiate registration of the
     * delegator is valid, as long as that nonce has been sent by the contract. For the purposes of computing delegation rewards,
     * the delegation is considered active after this function is completed.
     * Note: Only the specified delegation will be marked as registered, even if the validator weight update
     * message implicitly includes multiple weight changes.
     * @param delegationID The ID of the delegation being registered.
     * @param messageIndex The index of the ICM message to be received providing the acknowledgement.
     */
    function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) external;

    /**
     * @notice Begins the process of removing a delegator from a validation period, and reverts if the delegation is not eligible for rewards.
     * The delegator must have been previously registered with the given validationID. For the purposes of computing delegation rewards,
     * the delegation period is considered ended when this function is called. Uses the supplied uptime proof to calculate rewards.
     * If none is provided in the call, the latest known uptime will be used. Reverts if the uptime is not eligible for rewards.
     * Note: This function can only be called by the address that registered the delegation.
     * Note: Reverts if the uptime is not eligible for rewards.
     * @param delegationID The ID of the delegation being removed.
     * @param includeUptimeProof Whether or not an uptime proof is provided for the validation period.
     * If the validator has completed its validation period, it has already provided an uptime proof, so {includeUptimeProof}
     * will be ignored and can be set to false. If the validator has not completed its validation period and no uptime proof
     * is provided, the latest known uptime will be used.
     * @param messageIndex If {includeUptimeProof} is true, the index of the ICM message to be received providing the
     * uptime proof.
     */
    function initiateDelegatorRemoval(
        bytes32 delegationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external;

    /**
     * @notice Begins the process of redelegating a delegator's stake to a new validator. The delegator must have already initiatied the
     * removal of previous delegation. This skips the locking and unlocking of tokens.
     * Note: This function can only be called by the address that registered the delegation.
     * @param delegationID The ID of the delegation being redelegated.
     * @param validationID The ID of the validation period the delegation is being redelegated to.
     * @return The ID of the redelegation.
     */
    function initiateRedelegation(
        bytes32 delegationID,
        bytes32 validationID
    ) external returns (bytes32);

    /**
     * @notice Resubmits a delegator registration or delegator end message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param delegationID The ID of the delegation.
     */
    function resendUpdateDelegator(bytes32 delegationID) external;

    /**
     * @notice Completes the process of ending a delegation by receiving an acknowledgement from the P-Chain.
     * Any P-Chain acknowledgement with a nonce greater than or equal to the nonce used to initiate the end of the
     * delegator's delegation is valid, as long as that nonce has been sent by the contract. This is because the validator
     * weight change pertaining to the delegation ending is included in any subsequent validator weight update messages.
     * Note: Only the specified delegation will be marked as completed, even if the validator weight update
     * message implicitly includes multiple weight changes.
     * @param delegationID The ID of the delegation being removed.
     * @param messageIndex The index of the ICM message to be received providing the acknowledgement.
     */
    function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) external;
}
