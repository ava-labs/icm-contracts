// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {
    IPoSValidatorManager, Delegator, DelegatorStatus
} from "./interfaces/IPoSValidatorManager.sol";
import {PoSValidatorManagerSettings} from "./interfaces/IPoSValidatorManager.sol";
import {Validator} from "./interfaces/IValidatorManager.sol";
import {ValidatorManager} from "./ValidatorManager.sol";
import {Validator, ValidatorStatus} from "./interfaces/IValidatorManager.sol";
import {WarpMessage} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";
import {IRewardCalculator} from "./interfaces/IRewardCalculator.sol";

abstract contract PoSValidatorManager is IPoSValidatorManager, ValidatorManager {
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.PoSValidatorManager
    struct PoSValidatorManagerStorage {
        /// @notice The minimum amount of stake required to be a validator.
        uint256 _minimumStakeAmount;
        /// @notice The maximum amount of stake allowed to be a validator.
        uint256 _maximumStakeAmount;
        /// @notice The minimum amount of time a validator must be staked for.
        uint64 _minimumStakeDuration;
        /// @notice The minimum fee rate in basis points charged to the delegators by the validator.
        uint256 _minimumDelegationFeeRate;
        /// @notice The reward calculator for this validator manager.
        IRewardCalculator _rewardCalculator;
        /// @notice Maps the validationID to the delegation fee rate.
        mapping(bytes32 validationID => uint256) _validatorDelegationFeeRates;
        /// @notice Maps the delegationID to the delegator information.
        mapping(bytes32 delegationID => Delegator) _delegatorStakes;
        /// @notice Maps the delegationID to pending register delegator messages.
        mapping(bytes32 delegationID => bytes) _pendingRegisterDelegatorMessages;
        /// @notice Maps the delegationID to pending end delegator messages.
        mapping(bytes32 delegationID => bytes) _pendingEndDelegatorMessages;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.PoSValidatorManager")) - 1)) & ~bytes32(uint256(0xff));
    // TODO: Unit test for storage slot and update slot
    bytes32 private constant _POS_VALIDATOR_MANAGER_STORAGE_LOCATION =
        0x4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d00;

    // The maximum delegation fee rate in basis points. This is 99.99%
    uint256 public constant MAXIMUM_DELEGATION_FEE_RATE = 1e4 - 1;

    // solhint-disable ordering
    function _getPoSValidatorManagerStorage()
        private
        pure
        returns (PoSValidatorManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := _POS_VALIDATOR_MANAGER_STORAGE_LOCATION
        }
    }

    // solhint-disable-next-line func-name-mixedcase
    function __POS_Validator_Manager_init(PoSValidatorManagerSettings calldata settings)
        internal
        onlyInitializing
    {
        __ValidatorManager_init(settings.baseSettings);
        __POS_Validator_Manager_init_unchained(
            settings.minimumStakeAmount,
            settings.maximumStakeAmount,
            settings.minimumStakeDuration,
            settings.minimumDelegationFeeRate,
            settings.rewardCalculator
        );
    }

    // solhint-disable-next-line func-name-mixedcase
    function __POS_Validator_Manager_init_unchained(
        uint256 minimumStakeAmount,
        uint256 maximumStakeAmount,
        uint64 minimumStakeDuration,
        uint256 minimumDelegationFeeRate,
        IRewardCalculator rewardCalculator
    ) internal onlyInitializing {
        require(
            minimumDelegationFeeRate > 0 && minimumDelegationFeeRate < MAXIMUM_DELEGATION_FEE_RATE,
            "PoSValidatorManager: invalid minimum delegation fee rate"
        );
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        $._minimumStakeAmount = minimumStakeAmount;
        $._maximumStakeAmount = maximumStakeAmount;
        $._minimumStakeDuration = minimumStakeDuration;
        $._rewardCalculator = rewardCalculator;
    }

    function initializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        if (includeUptimeProof) {
            _getUptime(validationID, messageIndex);
        }
        // TODO: Calculate the reward for the validator, but do not unlock it

        _initializeEndValidation(validationID);
    }

    // This should only be called from inside the initializeValidatorRegistration functions
    // of the child contracts.
    function _setDelegationFeeRate(bytes32 validationID, uint256 feeRate) internal {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        require(
            feeRate >= $._minimumDelegationFeeRate && feeRate <= MAXIMUM_DELEGATION_FEE_RATE,
            "PoSValidatorManager: invalid delegation fee rate"
        );
        $._validatorDelegationFeeRates[validationID] = feeRate;
        emit DelegationFeeRateSet(validationID, feeRate);
    }

    function completeEndValidation(uint32 messageIndex) external {
        Validator memory validator = _completeEndValidation(messageIndex);
        _unlock(validator.startingWeight, validator.owner);
    }

    function _getUptime(bytes32 validationID, uint32 messageIndex) internal view returns (uint64) {
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(valid, "PoSValidatorManager: invalid warp message");

        require(
            warpMessage.sourceChainID == WARP_MESSENGER.getBlockchainID(),
            "PoSValidatorManager: invalid source chain ID"
        );
        require(
            warpMessage.originSenderAddress == address(0),
            "PoSValidatorManager: invalid origin sender address"
        );

        (bytes32 uptimeValidationID, uint64 uptime) =
            ValidatorMessages.unpackValidationUptimeMessage(warpMessage.payload);
        require(
            validationID == uptimeValidationID, "PoSValidatorManager: invalid uptime validation ID"
        );

        return uptime;
    }

    function _processStake(uint256 stakeAmount) internal virtual returns (uint64) {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        // Lock the stake in the contract.
        uint256 lockedValue = _lock(stakeAmount);
        uint64 weight = valueToWeight(lockedValue);

        // Ensure the weight is within the valid range.
        require(
            weight >= $._minimumStakeAmount && weight <= $._maximumStakeAmount,
            "PoSValidatorManager: invalid stake amount"
        );
        return weight;
    }

    function valueToWeight(uint256 value) public pure returns (uint64) {
        return uint64(value / 1e12);
    }

    function weightToValue(uint64 weight) public pure returns (uint256) {
        return uint256(weight) * 1e12;
    }

    function _lock(uint256 value) internal virtual returns (uint256);
    function _unlock(uint256 value, address to) internal virtual;

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint256 delegationAmount
    ) internal nonReentrant returns (bytes32) {
        uint64 weight = valueToWeight(_lock(delegationAmount));
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        // Ensure the validation period is active
        Validator memory validator = _getValidator(validationID);
        require(
            validator.status == ValidatorStatus.Active, "PoSValidatorManager: validator not active"
        );
        // Update the validator weight
        uint64 newValidatorWeight = validator.weight + weight;
        _setValidatorWeight(validationID, newValidatorWeight);

        // Construct the delegation ID. This is guaranteed to be unique since it is
        // constructed using a new nonce.
        uint64 nonce = _getAndIncrementNonce(validationID);
        bytes32 delegationID = keccak256(abi.encodePacked(validationID, delegatorAddress, nonce));

        _checkAndUpdateChurnTracker(weight);

        // Submit the message to the Warp precompile.
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, nonce, newValidatorWeight);
        $._pendingRegisterDelegatorMessages[delegationID] = setValidatorWeightPayload;
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(setValidatorWeightPayload);

        // Store the delegation information. Set the delegator status to pending added,
        // so that it can be properly started in the complete step, even if the delivered
        // nonce is greater than the nonce used to initialize registration.
        $._delegatorStakes[delegationID] = Delegator({
            status: DelegatorStatus.PendingAdded,
            owner: delegatorAddress,
            validationID: validationID,
            weight: weight,
            startedAt: 0,
            endedAt: 0,
            startingNonce: nonce,
            endingNonce: 0
        });

        emit DelegatorAdded({
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

    function resendDelegatorRegistration(bytes32 delegationID) external {
        _checkPendingRegisterDelegatorMessages(delegationID);
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        // Submit the message to the Warp precompile.
        WARP_MESSENGER.sendWarpMessage($._pendingRegisterDelegatorMessages[delegationID]);
    }

    function completeDelegatorRegistration(uint32 messageIndex, bytes32 delegationID) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        // Unpack the Warp message
        WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);
        (bytes32 validationID, uint64 nonce,) =
            ValidatorMessages.unpackSubnetValidatorWeightUpdateMessage(warpMessage.payload);
        _checkPendingRegisterDelegatorMessages(delegationID);
        delete $._pendingRegisterDelegatorMessages[delegationID];

        Validator memory validator = _getValidator(validationID);

        // The received nonce should be no greater than the highest sent nonce. This should never
        // happen since the staking manager is the only entity that can trigger a weight update
        // on the P-Chain.
        require(validator.messageNonce >= nonce, "PoSValidatorManager: invalid nonce");

        // The nonce should also be greater than or equal to the delegationID's starting nonce. This allows
        // a weight update using a higher nonce (which implicitly includes the delegation's weight update)
        // to be used to complete registration for an earlier delegation. This is necessary because the P-Chain
        // is only willing to sign the latest weight update.
        require(
            $._delegatorStakes[delegationID].startingNonce <= nonce,
            "PoSValidatorManager: nonce does not match"
        );

        // Ensure the delegator is pending added. Since anybody can call this function once
        // delegator registration has been initialized, we need to make sure that this function is only
        // callable after that has been done.
        require(
            $._delegatorStakes[delegationID].status == DelegatorStatus.PendingAdded,
            "PoSValidatorManager: delegationID not pending added"
        );
        // Update the delegation status
        $._delegatorStakes[delegationID].status = DelegatorStatus.Active;
        $._delegatorStakes[delegationID].startedAt = uint64(block.timestamp);
        emit DelegatorRegistered({
            delegationID: delegationID,
            validationID: validationID,
            nonce: nonce,
            startTime: block.timestamp
        });
    }

    function initializeEndDelegation(
        bytes32 delegationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        bytes32 validationID = $._delegatorStakes[delegationID].validationID;

        uint64 uptime;
        if (includeUptimeProof) {
            uptime = _getUptime(validationID, messageIndex);
        }

        // TODO: Calculate the delegator's reward, but do not unlock it

        // Ensure the delegator is active
        Delegator memory delegator = $._delegatorStakes[delegationID];
        require(
            delegator.status == DelegatorStatus.Active, "PoSValidatorManager: delegation not active"
        );
        // Only the delegation owner can end the delegation.
        require(
            delegator.owner == _msgSender(), "PoSValidatorManager: delegation not owned by sender"
        );
        uint64 nonce = _getAndIncrementNonce(validationID);

        // Set the delegator status to pending removed, so that it can be properly removed in
        // the complete step, even if the delivered nonce is greater than the nonce used to
        // initialize the removal.
        delegator.status = DelegatorStatus.PendingRemoved;
        delegator.endedAt = uint64(block.timestamp);
        delegator.endingNonce = nonce;

        $._delegatorStakes[delegationID] = delegator;

        Validator memory validator = _getValidator(validationID);
        require(validator.weight > delegator.weight, "PoSValidatorManager: Invalid weight");
        uint64 newValidatorWeight = validator.weight - delegator.weight;
        _setValidatorWeight(validationID, newValidatorWeight);

        // Submit the message to the Warp precompile.
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, nonce, newValidatorWeight);
        $._pendingEndDelegatorMessages[delegationID] = setValidatorWeightPayload;
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(setValidatorWeightPayload);

        emit DelegatorRemovalInitialized({
            delegationID: delegationID,
            validationID: validationID,
            nonce: nonce,
            validatorWeight: newValidatorWeight,
            endTime: block.timestamp,
            setWeightMessageID: messageID
        });
    }

    function resendEndDelegation(bytes32 delegationID) external {
        _checkPendingEndDelegatorMessage(delegationID);
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        WARP_MESSENGER.sendWarpMessage($._pendingEndDelegatorMessages[delegationID]);
    }

    function completeEndDelegation(uint32 messageIndex, bytes32 delegationID) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        // Unpack the Warp message
        WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);
        (bytes32 validationID, uint64 nonce,) =
            ValidatorMessages.unpackSubnetValidatorWeightUpdateMessage(warpMessage.payload);
        _checkPendingEndDelegatorMessage(delegationID);
        delete $._pendingEndDelegatorMessages[delegationID];

        Validator memory validator = _getValidator(validationID);
        // The received nonce should be no greater than the highest sent nonce. This should never
        // happen since the staking manager is the only entity that can trigger a weight update
        // on the P-Chain.
        require(validator.messageNonce >= nonce, "PoSValidatorManager: invalid nonce");

        // The nonce should also be greater than or equal to the delegationID's ending nonce. This allows
        // a weight update using a higher nonce (which implicitly includes the delegation's weight update)
        // to be used to complete delisting for an earlier delegation. This is necessary because the P-Chain
        // is only willing to sign the latest weight update.
        require(
            $._delegatorStakes[delegationID].endingNonce <= nonce,
            "PoSValidatorManager: nonce does not match"
        );

        // Ensure the delegator is pending removed. Since anybody can call this function once
        // end delegation has been initialized, we need to make sure that this function is only
        // callable after that has been done.
        require(
            $._delegatorStakes[delegationID].status == DelegatorStatus.PendingRemoved,
            "PoSValidatorManager: delegation not pending added"
        );

        // Update the delegator status
        $._delegatorStakes[delegationID].status = DelegatorStatus.Completed;

        // TODO: Unlock the delegator's stake and their reward

        emit DelegationEnded(delegationID, validationID, nonce);
    }

    function _checkPendingEndDelegatorMessage(bytes32 delegationID) private view {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        require(
            $._pendingEndDelegatorMessages[delegationID].length > 0
                && $._delegatorStakes[delegationID].status == DelegatorStatus.PendingRemoved,
            "PoSValidatorManager: delegation removal not pending"
        );
    }

    function _checkPendingRegisterDelegatorMessages(bytes32 delegationID) private view {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        require(
            $._pendingRegisterDelegatorMessages[delegationID].length > 0
                && $._delegatorStakes[delegationID].status == DelegatorStatus.PendingAdded,
            "PoSValidatorManager: delegation registration not pending"
        );
    }
}
