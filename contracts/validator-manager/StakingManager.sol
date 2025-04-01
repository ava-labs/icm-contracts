// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorManager} from "./ValidatorManager.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";
import {
    Delegator,
    DelegatorStatus,
    IStakingManager,
    PoSValidatorInfo,
    StakingManagerSettings
} from "./interfaces/IStakingManager.sol";
import {Validator, ValidatorStatus, PChainOwner} from "./ACP99Manager.sol";
import {
    IWarpMessenger,
    WarpMessage
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ReentrancyGuardUpgradeable.sol";
import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";

/**
 * @dev Implementation of the {IStakingManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
abstract contract StakingManager is
    IStakingManager,
    ContextUpgradeable,
    ReentrancyGuardUpgradeable
{
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.StakingManager
    struct StakingManagerStorage {
        ValidatorManager _manager;
        /// @notice The minimum amount of stake required to be a validator.
        uint256 _minimumStakeAmount;
        /// @notice The maximum amount of stake allowed to be a validator.
        uint256 _maximumStakeAmount;
        /// @notice The minimum amount of time in seconds a validator must be staked for. Must be at least {_churnPeriodSeconds}.
        uint64 _minimumStakeDuration;
        /// @notice The maximum amount of staked NFTs allowed to be a validator.
        uint256 _maximumNFTAmount;
        /// @notice The minimum delegation amount
        uint256 _minimumDelegationAmount;
        /// @notice The minimum delegation fee percentage, in basis points, required to delegate to a validator.
        uint16 _minimumDelegationFeeBips;
        /// @notice The factor used to convert between weight and value.
        uint256 _weightToValueFactor;
        /// @notice The ID of the blockchain that submits uptime proofs. This must be a blockchain validated by the subnetID that this contract manages.
        bytes32 _uptimeBlockchainID;
        /// @notice admin address
        address _admin;
        /// @notice The duration of an epoch in seconds
        uint64 _epochDuration;
        /// @notice The duration of the unlock period in seconds
        uint64 _unlockDuration;
        /// @notice Maps the validation ID to its requirements.
        mapping(bytes32 validationID => PoSValidatorInfo) _posValidatorInfo;
        /// @notice Maps the delegation ID to the delegator information.
        mapping(bytes32 delegationID => Delegator) _delegatorStakes;

        mapping(bytes32 delegationID => uint256[]) _lockedNFTs;

        mapping(uint64 epoch => uint256) _totalRewardWeight;
        mapping(uint64 epoch => uint256) _totalRewardWeightNFT;

        mapping(uint64 epoch => mapping(address account => uint256)) _accountRewardWeight;
        mapping(uint64 epoch => mapping(address account => uint256)) _accountRewardWeightNFT;

        mapping(uint64 epoch => mapping(address account => mapping(address token => uint256))) _rewardWithdrawn;
        mapping(uint64 epoch => mapping(address account => mapping(address token => uint256))) _rewardWithdrawnNFT;

        mapping(uint64 epoch => mapping(address token => uint256)) _rewardPools; 
        mapping(uint64 epoch => mapping(address token => uint256)) _rewardPoolsNFT;

        uint64 _epochOffset;

        mapping(bytes32 ID => bool) _unlocked;
        mapping(uint64 epoch => mapping(bytes32 validationID => uint256)) _validationUptimes;

        address _uptimeKeeper;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.StakingManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant STAKING_MANAGER_STORAGE_LOCATION =
        0xafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab600;

    uint16 public constant MAXIMUM_DELEGATION_FEE_BIPS = 10000;

    uint16 public constant BIPS_CONVERSION_FACTOR = 10000;

    bytes32 public constant P_CHAIN_BLOCKCHAIN_ID = bytes32(0);

    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    error InvalidDelegationFee(uint16 delegationFeeBips);
    error InvalidDelegationID(bytes32 delegationID);
    error InvalidDelegatorStatus(DelegatorStatus status);
    error InvalidStakeAmount(uint256 stakeAmount);
    error InvalidMinStakeDuration(uint64 minStakeDuration);
    error MaxWeightExceeded(uint64 newValidatorWeight);
    error MinStakeDurationNotPassed(uint64 endTime);
    error UnauthorizedOwner(address sender);
    error ValidatorNotPoS(bytes32 validationID);
    error ZeroWeightToValueFactor();
    error InvalidUptimeBlockchainID(bytes32 uptimeBlockchainID);
    error UnlockDurationNotPassed(uint64 endTime);
    error InvalidWarpOriginSenderAddress(address senderAddress);
    error InvalidWarpSourceChainID(bytes32 sourceChainID);
    error UnexpectedValidationID(bytes32 validationID, bytes32 expectedValidationID);
    error InvalidValidatorStatus(ValidatorStatus status);
    error InvalidNonce(uint64 nonce);
    error InvalidWarpMessage();

    // solhint-disable ordering
    /**
     * @dev This storage is visible to child contracts for convenience.
     *      External getters would be better practice, but code size limitations are preventing this.
     *      Child contracts should probably never write to this storage.
     */
    function _getStakingManagerStorage() internal pure returns (StakingManagerStorage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := STAKING_MANAGER_STORAGE_LOCATION
        }
    }

    // solhint-disable-next-line func-name-mixedcase
    function __StakingManager_init(StakingManagerSettings calldata settings)
        internal
        onlyInitializing
    {
        __ReentrancyGuard_init();
        __StakingManager_init_unchained({
            manager: settings.manager,
            minimumStakeAmount: settings.minimumStakeAmount,
            maximumStakeAmount: settings.maximumStakeAmount,
            minimumStakeDuration: settings.minimumStakeDuration,
            minimumDelegationAmount: settings.minimumDelegationAmount,
            minimumDelegationFeeBips: settings.minimumDelegationFeeBips,
            admin: settings.admin,
            weightToValueFactor: settings.weightToValueFactor,
            uptimeBlockchainID: settings.uptimeBlockchainID,
            unlockDuration: settings.unlockDuration,
            epochDuration: settings.epochDuration,
            maximumNFTAmount: settings.maximumNFTAmount,
            uptimeKeeper: settings.uptimeKeeper,
            epochOffset: settings.epochOffset
        });
    }

    // solhint-disable-next-line func-name-mixedcase
    function __StakingManager_init_unchained(
        ValidatorManager manager,
        uint256 minimumStakeAmount,
        uint256 maximumStakeAmount,
        uint256 maximumNFTAmount,
        uint64 minimumStakeDuration,
        uint256 minimumDelegationAmount,
        uint16 minimumDelegationFeeBips,
        address admin,
        uint256 weightToValueFactor,
        bytes32 uptimeBlockchainID,
        uint64 unlockDuration,
        uint64 epochDuration,
        address uptimeKeeper,
        uint64 epochOffset
    ) internal onlyInitializing {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        if (minimumDelegationFeeBips == 0 || minimumDelegationFeeBips > MAXIMUM_DELEGATION_FEE_BIPS)
        {
            revert InvalidDelegationFee(minimumDelegationFeeBips);
        }
        if (minimumStakeAmount > maximumStakeAmount) {
            revert InvalidStakeAmount(minimumStakeAmount);
        }
        // Minimum stake duration should be at least one churn period in order to prevent churn tracker abuse.
        if (minimumStakeDuration < manager.getChurnPeriodSeconds()) {
            revert InvalidMinStakeDuration(minimumStakeDuration);
        }
        if (weightToValueFactor == 0) {
            revert ZeroWeightToValueFactor();
        }
        if (uptimeBlockchainID == bytes32(0)) {
            revert InvalidUptimeBlockchainID(uptimeBlockchainID);
        }

        $._manager = manager;
        $._minimumStakeAmount = minimumStakeAmount;
        $._maximumStakeAmount = maximumStakeAmount;
        $._maximumNFTAmount = maximumNFTAmount;
        $._minimumStakeDuration = minimumStakeDuration;
        $._minimumDelegationAmount = minimumDelegationAmount;
        $._minimumDelegationFeeBips = minimumDelegationFeeBips;
        $._admin = admin;
        $._weightToValueFactor = weightToValueFactor;
        $._uptimeBlockchainID = uptimeBlockchainID;
        $._unlockDuration = unlockDuration;
        $._epochDuration = epochDuration;
        $._uptimeKeeper = uptimeKeeper;
        $._epochOffset = epochOffset;
    }

    /**
     * @notice See {IStakingManager-submitUptimeProof}.
     */
    function submitUptimeProof(bytes32 validationID, uint32 messageIndex) external {
        if (!_isPoSValidator(validationID)) {
            revert ValidatorNotPoS(validationID);
        }

        // Uptime proofs include the absolute number of seconds the validator has been active.
        _updateUptime(validationID, messageIndex);
    }

    /**
     * @notice See {IStakingManager-initiateValidatorRemoval}.
     * Extends the functionality of {ACP99Manager-initiateValidatorRemoval} updating staker state.
     */
    function initiateValidatorRemoval(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        _initiatePoSValidatorRemoval(validationID);
    }

    /**
     * @dev Helper function that initiates the end of a PoS validation period.
     */
    function _initiatePoSValidatorRemoval(
        bytes32 validationID
    ) internal {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        $._manager.initiateValidatorRemoval(validationID);

        // The validator must be fetched after the removal has been initiated, since the above call modifies
        // the validator's state.
        Validator memory validator = $._manager.getValidator(validationID);

        // Non-PoS validators are required to boostrap the network, but are not eligible for rewards.
        if (!_isPoSValidator(validationID)) {
            // Initial Validators can only be removed by the removal admin
            if ($._admin != _msgSender()) {
                revert UnauthorizedOwner(_msgSender());
            }
            return;
        }

        // PoS validations can only be ended by their owners.
        if ($._posValidatorInfo[validationID].owner != _msgSender()) {
            revert UnauthorizedOwner(_msgSender());
        }

        // Check that minimum stake duration has passed.
        if (
            validator.endTime
                < validator.startTime + $._posValidatorInfo[validationID].minStakeDuration
        ) {
            revert MinStakeDurationNotPassed(validator.endTime);
        }

        return;
    }

    /**
     * @notice See {IStakingManager-completeValidatorRemoval}.
     * Extends the functionality of {ACP99Manager-completeValidatorRemoval} by unlocking staking rewards.
     */
    function completeValidatorRemoval(uint32 messageIndex)
        external virtual
        nonReentrant
        returns (bytes32)
    {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        // Check if the validator has been already been removed from the validator manager.
        bytes32 validationID = $._manager.completeValidatorRemoval(messageIndex);

        return validationID;
    }

    /**
     * @dev Helper function that extracts the uptime from a ValidationUptimeMessage Warp message
     * If the uptime is greater than the stored uptime, update the stored uptime.
     */
    function _updateUptime(bytes32 validationID, uint32 messageIndex) virtual internal returns (uint64) {
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        if (!valid) {
            revert InvalidWarpMessage();
        }

        StakingManagerStorage storage $ = _getStakingManagerStorage();
        // The uptime proof must be from the specifed uptime blockchain
        if (warpMessage.sourceChainID != $._uptimeBlockchainID) {
            revert InvalidWarpSourceChainID(warpMessage.sourceChainID);
        }

        // The sender is required to be the zero address so that we know the validator node
        // signed the proof directly, rather than as an arbitrary on-chain message
        if (warpMessage.originSenderAddress != address(0)) {
            revert InvalidWarpOriginSenderAddress(warpMessage.originSenderAddress);
        }
        if (warpMessage.originSenderAddress != address(0)) {
            revert InvalidWarpOriginSenderAddress(warpMessage.originSenderAddress);
        }

        (bytes32 uptimeValidationID, uint64 uptime) =
            ValidatorMessages.unpackValidationUptimeMessage(warpMessage.payload);
        if (validationID != uptimeValidationID) {
            revert UnexpectedValidationID(uptimeValidationID, validationID);
        }

        if (uptime > $._posValidatorInfo[validationID].uptimeSeconds) {
            $._posValidatorInfo[validationID].uptimeSeconds = uptime;
            emit UptimeUpdated(validationID, uptime, 0);
        } else {
            uptime = $._posValidatorInfo[validationID].uptimeSeconds;
        }

        return uptime;
    }

    /**
     * @notice Initiates validator registration. Extends the functionality of {ACP99Manager-_initiateValidatorRegistration}
     * by locking stake and setting staking and delegation parameters.
     * @param delegationFeeBips The delegation fee in basis points.
     * @param minStakeDuration The minimum stake duration in seconds.
     * @param stakeAmount The amount of stake to lock.
     */
    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        uint64 registrationExpiry,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount
    ) internal virtual returns (bytes32) {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        // Validate and save the validator requirements
        if (
            delegationFeeBips < $._minimumDelegationFeeBips
                || delegationFeeBips > MAXIMUM_DELEGATION_FEE_BIPS
        ) {
            revert InvalidDelegationFee(delegationFeeBips);
        }

        if (minStakeDuration < $._minimumStakeDuration) {
            revert InvalidMinStakeDuration(minStakeDuration);
        }

        // Ensure the weight is within the valid range.
        if (stakeAmount < $._minimumStakeAmount || stakeAmount > $._maximumStakeAmount) {
            revert InvalidStakeAmount(stakeAmount);
        }

        // Lock the stake in the contract.
        uint256 lockedValue = _lock(stakeAmount);

        uint64 weight = valueToWeight(lockedValue);
        bytes32 validationID = $._manager.initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            registrationExpiry: registrationExpiry,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            weight: weight
        });

        address owner = _msgSender();

        $._posValidatorInfo[validationID].owner = owner;
        $._posValidatorInfo[validationID].delegationFeeBips = delegationFeeBips;
        $._posValidatorInfo[validationID].minStakeDuration = minStakeDuration;
        $._posValidatorInfo[validationID].uptimeSeconds = 0;

        return validationID;
    }

    /**
     * @notice See {IStakingManager-completeValidatorRegistration}.
     */
    function completeValidatorRegistration(uint32 messageIndex) external returns (bytes32) {
        return _getStakingManagerStorage()._manager.completeValidatorRegistration(messageIndex);
    }

    /**
     * @notice Converts a token value to a weight.
     * @param value Token value to convert.
     */
    function valueToWeight(uint256 value) public view returns (uint64) {
        uint256 weight = value / _getStakingManagerStorage()._weightToValueFactor;
        if (weight == 0 || weight > type(uint64).max) {
            revert InvalidStakeAmount(value);
        }
        return uint64(weight);
    }

    /**
     * @notice Converts a weight to a token value.
     * @param weight weight to convert.
     */
    function weightToValue(uint64 weight) public view returns (uint256) {
        return uint256(weight) * _getStakingManagerStorage()._weightToValueFactor;
    }

    /**
     * @notice Locks tokens in this contract.
     * @param value Number of tokens to lock.
     */
    function _lock(uint256 value) internal virtual returns (uint256);

    /**
     * @notice Unlocks token to a specific address.
     * @param to Address to send token to.
     * @param value Number of tokens to lock.
     */
    function _unlock(address to, uint256 value) internal virtual;

    /**
     * @notice Initiates delegator registration by updating the validator's weight and storing the delegation information.
     * Extends the functionality of {ACP99Manager-initiateValidatorWeightUpdate} by locking delegation stake.
     * @param validationID The ID of the validator to delegate to.
     * @param delegatorAddress The address of the delegator.
     * @param delegationAmount The amount of stake to delegate.
     */
    function _initiateDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint256 delegationAmount
    ) internal returns (bytes32) {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        uint64 weight = valueToWeight(_lock(delegationAmount));

        // Ensure the validation period is active
        Validator memory validator = $._manager.getValidator(validationID);
        // Check that the validation ID is a PoS validator
        if (!_isPoSValidator(validationID)) {
            revert ValidatorNotPoS(validationID);
        }
        if (validator.status != ValidatorStatus.Active) {
            revert InvalidValidatorStatus(validator.status);
        }

        if (delegationAmount < $._minimumDelegationAmount) {
            revert InvalidStakeAmount(delegationAmount);
        }
        // Update the validator weight
        uint64 newValidatorWeight = validator.weight + weight;
        if (newValidatorWeight > valueToWeight($._maximumStakeAmount)) {
            revert MaxWeightExceeded(newValidatorWeight);
        }

        (uint64 nonce, bytes32 messageID) =
            $._manager.initiateValidatorWeightUpdate(validationID, newValidatorWeight);

        bytes32 delegationID = keccak256(abi.encodePacked(validationID, nonce));

        // Store the delegation information. Set the delegator status to pending added,
        // so that it can be properly started in the complete step, even if the delivered
        // nonce is greater than the nonce used to initiate registration.
        $._delegatorStakes[delegationID].status = DelegatorStatus.PendingAdded;
        $._delegatorStakes[delegationID].owner = delegatorAddress;
        $._delegatorStakes[delegationID].validationID = validationID;
        $._delegatorStakes[delegationID].weight = weight;
        $._delegatorStakes[delegationID].startTime = 0;
        $._delegatorStakes[delegationID].startingNonce = nonce;
        $._delegatorStakes[delegationID].endingNonce = 0;

        emit InitiatedDelegatorRegistration({
            delegationID: delegationID,
            validationID: validationID,
            delegatorAddress: delegatorAddress,
            nonce: nonce,
            validatorWeight: newValidatorWeight,
            delegatorWeight: weight,
            setWeightMessageID: messageID
        });
        return delegationID;
    }

    /**
     * @notice See {IStakingManager-completeDelegatorRegistration}.
     * Extends the functionality of {ACP99Manager-completeValidatorWeightUpdate} by updating the delegation status.
     */
    function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) external {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        Delegator memory delegator = $._delegatorStakes[delegationID];
        bytes32 validationID = delegator.validationID;
        Validator memory validator = $._manager.getValidator(validationID);

        // Ensure the delegator is pending added. Since anybody can call this function once
        // delegator registration has been initiated, we need to make sure that this function is only
        // callable after that has been done.
        if (delegator.status != DelegatorStatus.PendingAdded) {
            revert InvalidDelegatorStatus(delegator.status);
        }

        // In the case where the validator has completed its validation period, we can no
        // longer stake and should move our status directly to completed and return the stake.
        if (validator.status == ValidatorStatus.Completed) {
            return _completeDelegatorRemoval(delegationID);
        }

        // If we've already received a weight update with a nonce greater than the delegation's starting nonce,
        // then there's no requirement to include an ICM message in this function call.
        if (validator.receivedNonce < delegator.startingNonce) {
            (bytes32 messageValidationID, uint64 nonce) =
                $._manager.completeValidatorWeightUpdate(messageIndex);

            if (validationID != messageValidationID) {
                revert UnexpectedValidationID(messageValidationID, validationID);
            }
            if (nonce < delegator.startingNonce) {
                revert InvalidNonce(nonce);
            }
        }

        // Update the delegation status
        $._delegatorStakes[delegationID].status = DelegatorStatus.Active;
        $._delegatorStakes[delegationID].startTime = uint64(block.timestamp);

        emit CompletedDelegatorRegistration({
            delegationID: delegationID,
            validationID: validationID,
            startTime: uint64(block.timestamp)
        });
    }

    /**
     * @notice See {IStakingManager-initiateRedelegation}.
     */
    function initiateRedelegation(
        bytes32 delegationID,
        bytes32 validationID
    ) external returns (bytes32) {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        Delegator memory delegator = $._delegatorStakes[delegationID];

        // Ensure the delegator is removed and tokens are not unlocked yet
        if (delegator.status != DelegatorStatus.Removed || $._unlocked[delegationID]) {
            revert InvalidDelegatorStatus(delegator.status);
        }

        $._unlocked[delegationID] = true;
        emit UnlockedDelegation(delegationID);

        // Ensure the validation period is active
        Validator memory validator = $._manager.getValidator(validationID);
        // Check that the validation ID is a PoS validator
        if (!_isPoSValidator(validationID)) {
            revert ValidatorNotPoS(validationID);
        }
        if (validator.status != ValidatorStatus.Active) {
            revert InvalidValidatorStatus(validator.status);
        }

        // Update the validator weight
        uint64 newValidatorWeight = validator.weight + delegator.weight;
        if (newValidatorWeight > valueToWeight($._maximumStakeAmount)) {
            revert MaxWeightExceeded(newValidatorWeight);
        }

        (uint64 nonce, bytes32 messageID) =
            $._manager.initiateValidatorWeightUpdate(validationID, newValidatorWeight);

        delegationID = keccak256(abi.encodePacked(validationID, nonce));

        // Store the delegation information. Set the delegator status to pending added,
        // so that it can be properly started in the complete step, even if the delivered
        // nonce is greater than the nonce used to initiate registration.
        $._delegatorStakes[delegationID].status = DelegatorStatus.PendingAdded;
        $._delegatorStakes[delegationID].owner = delegator.owner;
        $._delegatorStakes[delegationID].validationID = validationID;
        $._delegatorStakes[delegationID].weight = delegator.weight;
        $._delegatorStakes[delegationID].startTime = 0;
        $._delegatorStakes[delegationID].startingNonce = nonce;
        $._delegatorStakes[delegationID].endingNonce = 0;

        emit InitiatedDelegatorRegistration({
            delegationID: delegationID,
            validationID: validationID,
            delegatorAddress: delegator.owner,
            nonce: nonce,
            validatorWeight: newValidatorWeight,
            delegatorWeight: delegator.weight,
            setWeightMessageID: messageID
        });
        return delegationID;
    }

    /**
     * @notice See {IStakingManager-initiateDelegatorRemoval}.
     */
    function initiateDelegatorRemoval(
        bytes32 delegationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        _initiateDelegatorRemoval(delegationID);
    }

    /**
    * @notice Initiates the process of ending an delegation for a given delegation ID.
    * @dev This function ensures that the delegation is active and validates that the caller is authorized to end it.
    *      If the validator status is valid, the delegation status is updated to `PendingRemoved`. If the validator
    *      is complete, then removal is completed directly. Status is updated to `Completed` and initate
    *      `InitiatedDelegatorRemoval` is not emitted. 
    * @param delegationID The unique identifier of the delegation to be ended.
    **/
    function _initiateDelegatorRemoval(
        bytes32 delegationID
    ) internal {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        Delegator memory delegator = $._delegatorStakes[delegationID];
        bytes32 validationID = delegator.validationID;
        Validator memory validator = $._manager.getValidator(validationID);

        // Ensure the delegator is active
        if (delegator.status != DelegatorStatus.Active) {
            revert InvalidDelegatorStatus(delegator.status);
        }

        if (delegator.owner != _msgSender()) {
            revert UnauthorizedOwner(_msgSender());
        }

        if (validator.status == ValidatorStatus.Active) {
            // Check that minimum stake duration has passed.
            if (block.timestamp < delegator.startTime + $._minimumStakeDuration) {
                revert MinStakeDurationNotPassed(uint64(block.timestamp));
            }

            // Set the delegator status to pending removed, so that it can be properly removed in
            // the complete step, even if the delivered nonce is greater than the nonce used to
            // initiate the removal.
            $._delegatorStakes[delegationID].status = DelegatorStatus.PendingRemoved;
            $._delegatorStakes[delegationID].endTime = uint64(block.timestamp);

            ($._delegatorStakes[delegationID].endingNonce,) = $
                ._manager
                .initiateValidatorWeightUpdate(validationID, validator.weight - delegator.weight);

            emit InitiatedDelegatorRemoval({delegationID: delegationID, validationID: validationID});
            return;
        } else if (validator.status == ValidatorStatus.Completed) {
            $._delegatorStakes[delegationID].endTime = validator.endTime; 
            _completeDelegatorRemoval(delegationID);
            // If the validator has completed, then no further uptimes may be submitted, so we always
            // end the delegation.
            return;
        } else {
            revert InvalidValidatorStatus(validator.status);
        }
    }

    /**
     * @notice See {IStakingManager-resendUpdateDelegator}.
     * @dev Resending the latest validator weight with the latest nonce is safe because all weight changes are
     * cumulative, so the latest weight change will always include the weight change for any added delegators.
     */
    function resendUpdateDelegator(bytes32 delegationID) external {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        Delegator memory delegator = $._delegatorStakes[delegationID];
        if (
            delegator.status != DelegatorStatus.PendingAdded
                && delegator.status != DelegatorStatus.PendingRemoved
        ) {
            revert InvalidDelegatorStatus(delegator.status);
        }

        Validator memory validator = $._manager.getValidator(delegator.validationID);
        if (validator.sentNonce == 0) {
            // Should be unreachable.
            revert InvalidDelegationID(delegationID);
        }

        // Submit the message to the Warp precompile.
        WARP_MESSENGER.sendWarpMessage(
            ValidatorMessages.packL1ValidatorWeightMessage(
                delegator.validationID, validator.sentNonce, validator.weight
            )
        );
    }

    /**
     * @notice See {IStakingManager-completeDelegatorRemoval}.
     * Extends the functionality of {ACP99Manager-completeValidatorWeightUpdate} by updating the delegation status and unlocking delegation rewards.
     */
    function completeDelegatorRemoval(
        bytes32 delegationID,
        uint32 messageIndex
    ) external nonReentrant {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        Delegator memory delegator = $._delegatorStakes[delegationID];

        // Ensure the delegator is pending removed. Since anybody can call this function once
        // end delegation has been initiated, we need to make sure that this function is only
        // callable after that has been done.
        if (delegator.status != DelegatorStatus.PendingRemoved) {
            revert InvalidDelegatorStatus(delegator.status);
        }
        Validator memory validator = $._manager.getValidator(delegator.validationID);

        // We only expect an ICM message if we haven't received a weight update with a nonce greater than the delegation's ending nonce
        if (
            $._manager.getValidator(delegator.validationID).status != ValidatorStatus.Completed
                && validator.receivedNonce < delegator.endingNonce
        ) {
            (bytes32 validationID, uint64 nonce) =
                $._manager.completeValidatorWeightUpdate(messageIndex);
            if (delegator.validationID != validationID) {
                revert UnexpectedValidationID(validationID, delegator.validationID);
            }

            // The received nonce should be at least as high as the delegation's ending nonce. This allows a weight
            // update using a higher nonce (which implicitly includes the delegation's weight update) to be used to
            // complete delisting for an earlier delegation. This is necessary because the P-Chain is only willing
            // to sign the latest weight update.
            if (delegator.endingNonce > nonce) {
                revert InvalidNonce(nonce);
            }
        }

        _completeDelegatorRemoval(delegationID);
    }

    /**
     * @notice unlocks the validator stake, to be called after removal and passing of unlock duration
     * @param validationID The unique identifier of the validator to unlock.
     */ 
    function unlockValidator(bytes32 validationID) external virtual nonReentrant {
        _unlockValidator(validationID);
    }
    
    /**
     * @notice unlocks the delegator stake, to be called after removal and passing of unlock duration
     * @param delegationID The unique identifier of the delegator to unlock.
     */ 
    function unlockDelegator(bytes32 delegationID) external nonReentrant {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        Delegator memory delegator = $._delegatorStakes[delegationID]; 

        if (delegator.status != DelegatorStatus.Removed || $._unlocked[delegationID]) {
            revert InvalidDelegatorStatus(delegator.status);
        }

        if(delegator.startTime != 0 && block.timestamp < delegator.endTime + $._unlockDuration) {
            revert UnlockDurationNotPassed(uint64(block.timestamp));
        }

        $._unlocked[delegationID] = true;

        emit UnlockedDelegation(delegationID);
        // Unlock the delegator's stake.
        _unlock(delegator.owner, weightToValue(delegator.weight));
    }

    function _unlockValidator(bytes32 validationID) internal {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        Validator memory validator = $._manager.getValidator(validationID);

        if ((validator.status != ValidatorStatus.Completed && validator.status != ValidatorStatus.Invalidated) 
                || $._unlocked[validationID]) {
            revert InvalidValidatorStatus(validator.status);
        }

        if (!_isPoSValidator(validationID)) {
            revert ValidatorNotPoS(validationID);
        }

        if(validator.startTime != 0 && block.timestamp < validator.endTime + $._unlockDuration) {
            revert UnlockDurationNotPassed(uint64(block.timestamp));
        }

        $._unlocked[validationID] = true;

        emit UnlockedValidation(validationID);
        // The stake is unlocked whether the validation period is completed or invalidated.
        _unlock($._posValidatorInfo[validationID].owner, weightToValue(validator.startingWeight)); 
    }

    function _completeDelegatorRemoval(bytes32 delegationID) internal {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        Delegator memory delegator = $._delegatorStakes[delegationID];
        bytes32 validationID = delegator.validationID;

        // To prevent churn tracker abuse, check that one full churn period has passed,
        // so a delegator may not stake twice in the same churn period.
        if (block.timestamp < delegator.startTime + $._manager.getChurnPeriodSeconds()) {
            revert MinStakeDurationNotPassed(uint64(block.timestamp));
        }

        $._delegatorStakes[delegationID].status = DelegatorStatus.Removed;

        emit CompletedDelegatorRemoval(delegationID, validationID, 0, 0);
    }

    /**
     * @dev This function must be implemented to mint rewards to validators and delegators.
     */
    function _reward(address account, uint256 amount) internal virtual;

    /**
     * @dev Return true if this is a PoS validator with locked stake. Returns false if this was originally a PoA
     * validator that was later migrated to this PoS manager, or the validator was part of the initial validator set.
     */
    function _isPoSValidator(bytes32 validationID) internal view returns (bool) {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        return $._posValidatorInfo[validationID].owner != address(0);
    }
}
