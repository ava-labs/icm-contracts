// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IValidatorManager} from "./interfaces/IValidatorManager.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";
import {
    InitialValidator,
    PChainOwner,
    ConversionData,
    Validator,
    ValidatorStatus
} from "./interfaces/IACP99Manager.sol";
import {ACP99Manager} from "./ACP99Manager.sol";
import {
    IWarpMessenger,
    WarpMessage
} from "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/IWarpMessenger.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";

/**
 * @dev Describes the current churn period
 */
struct ValidatorChurnPeriod {
    uint256 startTime;
    uint64 initialWeight;
    uint64 totalWeight;
    uint64 churnAmount;
}

/**
 * @notice Validator Manager settings, used to initialize the Validator Manager
 * @param The subnetID is the ID of the L1 that the Validator Manager is managing
 * @param The churnPeriodSeconds is the duration of the churn period in seconds
 * @param The maximumChurnPercentage is the maximum percentage of the total weight that can be added or removed in a single churn period
 */
struct ValidatorManagerSettings {
    address admin;
    bytes32 subnetID;
    uint64 churnPeriodSeconds;
    uint8 maximumChurnPercentage;
}

/// @dev Legacy struct used to migrate from V1 contracts
struct ValidatorLegacy {
    ValidatorStatus status;
    bytes nodeID;
    uint64 startingWeight;
    uint64 messageNonce;
    uint64 weight;
    uint64 startedAt;
    uint64 endedAt;
}

/**
 * @dev Implementation of the {ACP99Manager} abstract contract.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract ValidatorManager is IValidatorManager, Initializable, OwnableUpgradeable, ACP99Manager {
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.ValidatorManager
    struct ValidatorManagerStorage {
        /// @notice The subnetID associated with this validator manager.
        bytes32 _subnetID;
        /// @notice The number of seconds after which to reset the churn tracker.
        uint64 _churnPeriodSeconds;
        /// @notice The maximum churn rate allowed per churn period.
        uint8 _maximumChurnPercentage;
        /// @notice The churn tracker used to track the amount of stake added or removed in the churn period.
        ValidatorChurnPeriod _churnTracker;
        /// @notice Maps the validationID to the registration message such that the message can be re-sent if needed.
        mapping(bytes32 => bytes) _pendingRegisterValidationMessages;
        /// @notice Legacy storage for V1 validators.
        mapping(bytes32 => ValidatorLegacy) _validationPeriodsLegacy;
        /// @notice Maps the nodeID to the validationID for validation periods that have not ended.
        mapping(bytes => bytes32) _registeredValidators;
        /// @notice Boolean that indicates if the initial validator set has been set.
        bool _initializedValidatorSet;
        /// @notice Maps the validationID to the validator information.
        mapping(bytes32 => Validator) _validationPeriods;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.ValidatorManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant VALIDATOR_MANAGER_STORAGE_LOCATION =
        0xe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb00;

    uint64 public constant REGISTRATION_EXPIRY_LENGTH = 1 days;
    // Limit churn period to the registration expiry length, so that a validation can only
    // be represented once per churn period
    uint64 public constant MAXIMUM_CHURN_PERIOD_LENGTH = REGISTRATION_EXPIRY_LENGTH;
    uint8 public constant MAXIMUM_CHURN_PERCENTAGE_LIMIT = 20;
    uint32 public constant NODE_ID_LENGTH = 20;
    uint8 public constant BLS_PUBLIC_KEY_LENGTH = 48;
    bytes32 public constant P_CHAIN_BLOCKCHAIN_ID = bytes32(0);

    // solhint-disable ordering
    /**
     * @dev This storage is visible to child contracts for convenience.
     *      External getters would be better practice, but code size limitations are preventing this.
     *      Child contracts should probably never write to this storage.
     */
    function _getValidatorManagerStorage()
        internal
        pure
        returns (ValidatorManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := VALIDATOR_MANAGER_STORAGE_LOCATION
        }
    }

    /**
     * @notice Warp precompile used for sending and receiving Warp messages.
     */
    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    constructor(
        ICMInitializable init
    ) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    /**
     * @notice Migrates a validator from the V1 contract to the V2 contract.
     * @param validationID The ID of the validation period to migrate.
     * @param receivedNonce The latest nonce received from the P-Chain.
     */
    function migrateFromV1(bytes32 validationID, uint32 receivedNonce) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        ValidatorLegacy storage legacy = $._validationPeriodsLegacy[validationID];
        if (legacy.status == ValidatorStatus.Unknown) {
            revert InvalidValidationID(validationID);
        }
        if (receivedNonce > legacy.messageNonce) {
            revert InvalidNonce(receivedNonce);
        }

        $._validationPeriods[validationID] = Validator({
            status: legacy.status,
            nodeID: legacy.nodeID,
            startingWeight: legacy.startingWeight,
            sentNonce: legacy.messageNonce,
            receivedNonce: receivedNonce,
            weight: legacy.weight,
            startTime: legacy.startedAt,
            endTime: legacy.endedAt
        });

        // Set the legacy status to unknown to disallow future migrations.
        $._validationPeriodsLegacy[validationID].status = ValidatorStatus.Unknown;
    }

    function initialize(
        ValidatorManagerSettings calldata settings
    ) external initializer {
        __ValidatorManager_init(settings);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ValidatorManager_init(
        ValidatorManagerSettings calldata settings
    ) internal onlyInitializing {
        __Ownable_init(settings.admin);
        __ValidatorManager_init_unchained(settings);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ValidatorManager_init_unchained(
        ValidatorManagerSettings calldata settings
    ) internal onlyInitializing {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        $._subnetID = settings.subnetID;

        if (
            settings.maximumChurnPercentage > MAXIMUM_CHURN_PERCENTAGE_LIMIT
                || settings.maximumChurnPercentage == 0
        ) {
            revert InvalidMaximumChurnPercentage(settings.maximumChurnPercentage);
        }
        if (settings.churnPeriodSeconds > MAXIMUM_CHURN_PERIOD_LENGTH) {
            revert InvalidChurnPeriodLength(settings.churnPeriodSeconds);
        }

        $._maximumChurnPercentage = settings.maximumChurnPercentage;
        $._churnPeriodSeconds = settings.churnPeriodSeconds;
    }

    modifier initializedValidatorSet() {
        if (!_getValidatorManagerStorage()._initializedValidatorSet) {
            revert InvalidInitializationStatus();
        }
        _;
    }

    /**
     * @notice See {IACP99Manager-initializeValidatorSet}.
     */
    function initializeValidatorSet(
        ConversionData calldata conversionData,
        uint32 messageIndex
    ) public virtual override {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        if ($._initializedValidatorSet) {
            revert InvalidInitializationStatus();
        }

        // Check that the blockchainID and validator manager address in the ConversionData correspond to this contract.
        // Other validation checks are done by the P-Chain when converting the L1, so are not required here.
        if (conversionData.validatorManagerBlockchainID != WARP_MESSENGER.getBlockchainID()) {
            revert InvalidValidatorManagerBlockchainID(conversionData.validatorManagerBlockchainID);
        }
        if (address(conversionData.validatorManagerAddress) != address(this)) {
            revert InvalidValidatorManagerAddress(address(conversionData.validatorManagerAddress));
        }

        // Verify that the sha256 hash of the L1 conversion data matches with the Warp message's conversionID.
        bytes32 conversionID = ValidatorMessages.unpackSubnetToL1ConversionMessage(
            _getPChainWarpMessage(messageIndex).payload
        );
        bytes memory encodedConversion = ValidatorMessages.packConversionData(conversionData);
        bytes32 encodedConversionID = sha256(encodedConversion);
        if (encodedConversionID != conversionID) {
            revert InvalidConversionID(encodedConversionID, conversionID);
        }

        uint256 numInitialValidators = conversionData.initialValidators.length;

        uint64 totalWeight;
        for (uint32 i; i < numInitialValidators; ++i) {
            InitialValidator memory initialValidator = conversionData.initialValidators[i];
            if ($._registeredValidators[initialValidator.nodeID] != bytes32(0)) {
                revert NodeAlreadyRegistered(initialValidator.nodeID);
            }
            if (initialValidator.nodeID.length != NODE_ID_LENGTH) {
                revert InvalidNodeID(initialValidator.nodeID);
            }

            // Validation ID of the initial validators is the sha256 hash of the
            // subnet ID and the index of the initial validator.
            bytes32 validationID = sha256(abi.encodePacked(conversionData.subnetID, i));

            // Save the initial validator as an active validator.
            $._registeredValidators[initialValidator.nodeID] = validationID;
            $._validationPeriods[validationID].status = ValidatorStatus.Active;
            $._validationPeriods[validationID].nodeID = initialValidator.nodeID;
            $._validationPeriods[validationID].startingWeight = initialValidator.weight;
            $._validationPeriods[validationID].sentNonce = 0;
            $._validationPeriods[validationID].weight = initialValidator.weight;
            $._validationPeriods[validationID].startTime = uint64(block.timestamp);
            $._validationPeriods[validationID].endTime = 0;
            totalWeight += initialValidator.weight;

            emit RegisteredInitialValidator(
                validationID,
                _fixedNodeID(initialValidator.nodeID),
                conversionData.subnetID,
                initialValidator.weight
            );
        }
        $._churnTracker.totalWeight = totalWeight;

        // Rearranged equation for totalWeight < (100 / $._maximumChurnPercentage)
        // Total weight must be above this value in order to not trigger churn limits with an added/removed weight of 1.
        if (totalWeight * $._maximumChurnPercentage < 100) {
            revert InvalidTotalWeight(totalWeight);
        }

        $._initializedValidatorSet = true;
    }

    function _validatePChainOwner(
        PChainOwner memory pChainOwner
    ) internal pure {
        // If threshold is 0, addresses must be empty.
        if (pChainOwner.threshold == 0 && pChainOwner.addresses.length != 0) {
            revert InvalidPChainOwnerThreshold(pChainOwner.threshold, pChainOwner.addresses.length);
        }
        // Threshold must be less than or equal to the number of addresses.
        if (pChainOwner.threshold > pChainOwner.addresses.length) {
            revert InvalidPChainOwnerThreshold(pChainOwner.threshold, pChainOwner.addresses.length);
        }
        // Zero address is invalid. Because we require addresses to be sorted, we only need to check if the first is 0
        if (pChainOwner.addresses.length > 0 && pChainOwner.addresses[0] == address(0)) {
            revert ZeroAddress();
        }
        // Addresses must be unique and sorted in ascending order
        for (uint256 i = 1; i < pChainOwner.addresses.length; i++) {
            // Compare current address with the previous one
            if (pChainOwner.addresses[i] <= pChainOwner.addresses[i - 1]) {
                revert InvalidPChainOwnerAddresses();
            }
        }
    }

    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) public onlyOwner returns (bytes32) {
        return _initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            weight: weight
        });
    }

    /**
     * @notice See {ACP99Manager-_initiateValidatorRegistration}.
     * @dev This function modifies the validator's state. Callers should ensure that any references are updated.
     */
    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) internal virtual override initializedValidatorSet returns (bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // Ensure the new validator doesn't overflow the total weight
        if (uint256(weight) + uint256($._churnTracker.totalWeight) > type(uint64).max) {
            revert InvalidTotalWeight(weight);
        }

        _validatePChainOwner(remainingBalanceOwner);
        _validatePChainOwner(disableOwner);

        // Ensure the nodeID is not the zero address, and is not already an active validator.

        if (blsPublicKey.length != BLS_PUBLIC_KEY_LENGTH) {
            revert InvalidBLSKeyLength(blsPublicKey.length);
        }
        if (nodeID.length != NODE_ID_LENGTH) {
            revert InvalidNodeID(nodeID);
        }
        if ($._registeredValidators[nodeID] != bytes32(0)) {
            revert NodeAlreadyRegistered(nodeID);
        }

        // Check that adding this validator would not exceed the maximum churn rate.
        _checkAndUpdateChurnTracker(weight, 0);

        uint64 registrationExpiry = uint64(block.timestamp) + REGISTRATION_EXPIRY_LENGTH;

        (bytes32 validationID, bytes memory registerL1ValidatorMessage) = ValidatorMessages
            .packRegisterL1ValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                subnetID: $._subnetID,
                nodeID: nodeID,
                blsPublicKey: blsPublicKey,
                remainingBalanceOwner: remainingBalanceOwner,
                disableOwner: disableOwner,
                registrationExpiry: registrationExpiry,
                weight: weight
            })
        );

        // Redundant check to ensure no collision or replay is possible, but with the expiry set as
        // the block timestamp + 1 day, this should not be possible.
        if ($._validationPeriods[validationID].status != ValidatorStatus.Unknown) {
            revert InvalidValidatorStatus($._validationPeriods[validationID].status);
        }

        $._pendingRegisterValidationMessages[validationID] = registerL1ValidatorMessage;
        $._registeredValidators[nodeID] = validationID;

        // Submit the message to the Warp precompile.
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(registerL1ValidatorMessage);
        $._validationPeriods[validationID].status = ValidatorStatus.PendingAdded;
        $._validationPeriods[validationID].nodeID = nodeID;
        $._validationPeriods[validationID].startingWeight = weight;
        $._validationPeriods[validationID].sentNonce = 0;
        $._validationPeriods[validationID].weight = weight;
        $._validationPeriods[validationID].startTime = 0; // The validation period only starts once the registration is acknowledged.
        $._validationPeriods[validationID].endTime = 0;

        emit InitiatedValidatorRegistration(
            validationID, _fixedNodeID(nodeID), messageID, registrationExpiry, weight
        );

        return validationID;
    }

    /**
     * @notice Resubmits a validator registration message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being registered.
     */
    function resendRegisterValidatorMessage(
        bytes32 validationID
    ) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        // The initial validator set must have been set already to have pending register validation messages.
        if ($._pendingRegisterValidationMessages[validationID].length == 0) {
            revert InvalidValidationID(validationID);
        }
        if ($._validationPeriods[validationID].status != ValidatorStatus.PendingAdded) {
            revert InvalidValidatorStatus($._validationPeriods[validationID].status);
        }

        // Submit the message to the Warp precompile.
        WARP_MESSENGER.sendWarpMessage($._pendingRegisterValidationMessages[validationID]);
    }

    /**
     * @notice See {IACP99Manager-completeValidatorRegistration}.
     */
    function completeValidatorRegistration(
        uint32 messageIndex
    ) public virtual override onlyOwner returns (bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        (bytes32 validationID, bool validRegistration) = ValidatorMessages
            .unpackL1ValidatorRegistrationMessage(_getPChainWarpMessage(messageIndex).payload);

        if (!validRegistration) {
            revert UnexpectedRegistrationStatus(validRegistration);
        }
        // The initial validator set must have been set already to have pending register validation messages.
        if ($._pendingRegisterValidationMessages[validationID].length == 0) {
            revert InvalidValidationID(validationID);
        }
        if ($._validationPeriods[validationID].status != ValidatorStatus.PendingAdded) {
            revert InvalidValidatorStatus($._validationPeriods[validationID].status);
        }

        delete $._pendingRegisterValidationMessages[validationID];
        $._validationPeriods[validationID].status = ValidatorStatus.Active;
        $._validationPeriods[validationID].startTime = uint64(block.timestamp);
        emit CompletedValidatorRegistration(validationID, $._validationPeriods[validationID].weight);

        return validationID;
    }

    /**
     * @notice See {IACP99Manager-getValidator}.
     */
    function getValidator(
        bytes32 validationID
    ) public view virtual override returns (Validator memory) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return $._validationPeriods[validationID];
    }

    /**
     * @notice See {IACP99Manager-l1TotalWeight}.
     */
    function l1TotalWeight() public view virtual override returns (uint64) {
        return _getValidatorManagerStorage()._churnTracker.totalWeight;
    }

    /**
     * @notice See {IACP99Manager-subnetID}.
     */
    function subnetID() public view virtual override returns (bytes32) {
        return _getValidatorManagerStorage()._subnetID;
    }

    /**
     * @notice Returns the current churn tracker and its configuration
     * @return The churn period duration in seconds
     * @return The maximum percentage of total L1 weight churn allowed per period
     * @return The current churn tracker
     */
    function getChurnTracker() public view returns (uint64, uint8, ValidatorChurnPeriod memory) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return ($._churnPeriodSeconds, $._maximumChurnPercentage, $._churnTracker);
    }

    /**
     * @notice Returns the validationID that the provided nodeID is registered under
     * @param nodeID ID of the node associated with the validation ID
     */
    function getNodeValidationID(
        bytes calldata nodeID
    ) public view returns (bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return $._registeredValidators[nodeID];
    }

    /**
     * @notice Returns true if the ValidatorManager has been initialized with the initial validator set
     */
    function isValidatorSetInitialized() public view returns (bool) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return $._initializedValidatorSet;
    }

    /**
     * @notice Initiates the removal of a validator from the active validator set.
     */
    function initiateValidatorRemoval(
        bytes32 validationID
    ) public onlyOwner {
        _initiateValidatorRemoval(validationID);
    }

    /**
     * @notice See {ACP99Manager-_initiateValidatorRemoval}.
     * @dev This function modifies the validator's state. Callers should ensure that any references are updated.
     */
    function _initiateValidatorRemoval(
        bytes32 validationID
    ) internal virtual override {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // Ensure the validation period is active.
        // The initial validator set must have been set already to have active validators.
        Validator memory validator = $._validationPeriods[validationID];
        if (validator.status != ValidatorStatus.Active) {
            revert InvalidValidatorStatus($._validationPeriods[validationID].status);
        }

        // Update the validator status to pending removal.
        // They are not removed from the active validators mapping until the P-Chain acknowledges the removal.
        validator.status = ValidatorStatus.PendingRemoved;

        // Set the end time of the validation period, since it is no longer known to be an active validator
        // on the P-Chain.
        validator.endTime = uint64(block.timestamp);

        // Save the validator updates.
        $._validationPeriods[validationID] = validator;

        (, bytes32 messageID) = _initiateValidatorWeightUpdate(validationID, 0);

        // Emit the event to signal the start of the validator removal process.
        emit InitiatedValidatorRemoval(
            validationID, messageID, validator.weight, uint64(block.timestamp)
        );
    }

    /**
     * @notice Resubmits a validator end message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being ended.
     */
    function resendValidatorRemovalMessage(
        bytes32 validationID
    ) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        Validator memory validator = $._validationPeriods[validationID];

        // The initial validator set must have been set already to have pending end validation messages.
        if (validator.status != ValidatorStatus.PendingRemoved) {
            revert InvalidValidatorStatus($._validationPeriods[validationID].status);
        }

        WARP_MESSENGER.sendWarpMessage(
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, validator.sentNonce, 0)
        );
    }

    /**
     * @notice See {IACP99Manager-completeValidatorRemoval}.
     */
    function completeValidatorRemoval(
        uint32 messageIndex
    ) public virtual override onlyOwner returns (bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // Get the Warp message.
        (bytes32 validationID, bool registered) = ValidatorMessages
            .unpackL1ValidatorRegistrationMessage(_getPChainWarpMessage(messageIndex).payload);
        if (registered) {
            revert UnexpectedRegistrationStatus(registered);
        }

        Validator memory validator = $._validationPeriods[validationID];

        // The validation status is PendingRemoved if validator removal was initiated with a call to {initiateValidatorRemoval}.
        // The validation status is PendingAdded if the validator was never registered on the P-Chain.
        // The initial validator set must have been set already to have pending validation messages.
        if (
            validator.status != ValidatorStatus.PendingRemoved
                && validator.status != ValidatorStatus.PendingAdded
        ) {
            revert InvalidValidatorStatus(validator.status);
        }

        if (validator.status == ValidatorStatus.PendingRemoved) {
            validator.status = ValidatorStatus.Completed;
        } else {
            // Remove the validator's weight from the total tracked weight, but don't track it as churn.
            $._churnTracker.totalWeight -= validator.weight;
            validator.status = ValidatorStatus.Invalidated;
        }
        // Remove the validator from the registered validators mapping.
        delete $._registeredValidators[validator.nodeID];

        // Update the validator.
        $._validationPeriods[validationID] = validator;

        // Emit event.
        emit CompletedValidatorRemoval(validationID);

        return validationID;
    }

    function _incrementSentNonce(
        bytes32 validationID
    ) internal returns (uint64) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return ++$._validationPeriods[validationID].sentNonce;
    }

    function _getPChainWarpMessage(
        uint32 messageIndex
    ) internal view returns (WarpMessage memory) {
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        if (!valid) {
            revert InvalidWarpMessage();
        }
        // Must match to P-Chain blockchain id, which is 0.
        if (warpMessage.sourceChainID != P_CHAIN_BLOCKCHAIN_ID) {
            revert InvalidWarpSourceChainID(warpMessage.sourceChainID);
        }
        if (warpMessage.originSenderAddress != address(0)) {
            revert InvalidWarpOriginSenderAddress(warpMessage.originSenderAddress);
        }

        return warpMessage;
    }

    function initiateValidatorWeightUpdate(
        bytes32 validationID,
        uint64 newWeight
    ) public onlyOwner returns (uint64, bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        if ($._validationPeriods[validationID].status != ValidatorStatus.Active) {
            revert InvalidValidatorStatus($._validationPeriods[validationID].status);
        }

        return _initiateValidatorWeightUpdate(validationID, newWeight);
    }

    /**
     * @notice See {ACP99Manager-_initiateValidatorWeightUpdate}.
     * @dev This function modifies the validator's state. Callers should ensure that any references are updated.
     */
    function _initiateValidatorWeightUpdate(
        bytes32 validationID,
        uint64 newWeight
    ) internal virtual override returns (uint64, bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        uint64 validatorWeight = $._validationPeriods[validationID].weight;

        // Check that changing the validator weight would not exceed the maximum churn rate.
        _checkAndUpdateChurnTracker(newWeight, validatorWeight);

        uint64 nonce = _incrementSentNonce(validationID);

        $._validationPeriods[validationID].weight = newWeight;

        // Submit the message to the Warp precompile.
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, nonce, newWeight)
        );

        emit InitiatedValidatorWeightUpdate({
            validationID: validationID,
            nonce: nonce,
            weightUpdateMessageID: messageID,
            weight: newWeight
        });

        return (nonce, messageID);
    }

    /**
     * @notice See {IACP99Manager-completeValidatorWeightUpdate}.
     */
    function completeValidatorWeightUpdate(
        uint32 messageIndex
    ) public virtual override onlyOwner returns (bytes32, uint64) {
        WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);
        (bytes32 validationID, uint64 nonce, uint64 weight) =
            ValidatorMessages.unpackL1ValidatorWeightMessage(warpMessage.payload);

        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // The received nonce should be no greater than the highest sent nonce to ensure
        // that weight changes are only initiated by this contract.
        if ($._validationPeriods[validationID].sentNonce < nonce) {
            revert InvalidNonce(nonce);
        }

        $._validationPeriods[validationID].receivedNonce = nonce;

        emit CompletedValidatorWeightUpdate(validationID, nonce, weight);

        return (validationID, nonce);
    }

    function getChurnPeriodSeconds() public view returns (uint64) {
        return _getValidatorManagerStorage()._churnPeriodSeconds;
    }

    /**
     * @dev Helper function to check if the stake weight to be added or removed would exceed the maximum stake churn
     * rate for the past churn period. If the churn rate is exceeded, the function will revert. If the churn rate is
     * not exceeded, the function will update the churn tracker with the new weight.
     */
    function _checkAndUpdateChurnTracker(
        uint64 newValidatorWeight,
        uint64 oldValidatorWeight
    ) private {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        uint64 weightChange;
        if (newValidatorWeight > oldValidatorWeight) {
            weightChange = newValidatorWeight - oldValidatorWeight;
        } else {
            weightChange = oldValidatorWeight - newValidatorWeight;
        }

        uint256 currentTime = block.timestamp;
        ValidatorChurnPeriod memory churnTracker = $._churnTracker;

        if (
            churnTracker.startTime == 0
                || currentTime >= churnTracker.startTime + $._churnPeriodSeconds
        ) {
            churnTracker.churnAmount = weightChange;
            churnTracker.startTime = currentTime;
            churnTracker.initialWeight = churnTracker.totalWeight;
        } else {
            // Churn is always additive whether the weight is being added or removed.
            churnTracker.churnAmount += weightChange;
        }

        // Rearranged equation of maximumChurnPercentage >= currentChurnPercentage to avoid integer division truncation.
        if ($._maximumChurnPercentage * churnTracker.initialWeight < churnTracker.churnAmount * 100)
        {
            revert MaxChurnRateExceeded(churnTracker.churnAmount);
        }

        // Two separate calculations because we're using uints and (newValidatorWeight - oldValidatorWeight) could underflow.
        churnTracker.totalWeight += newValidatorWeight;
        churnTracker.totalWeight -= oldValidatorWeight;

        // Rearranged equation for totalWeight < (100 / $._maximumChurnPercentage)
        // Total weight must be above this value in order to not trigger churn limits with an added/removed weight of 1.
        if (churnTracker.totalWeight * $._maximumChurnPercentage < 100) {
            revert InvalidTotalWeight(churnTracker.totalWeight);
        }

        $._churnTracker = churnTracker;
    }

    /**
     * @notice Converts a nodeID to a fixed length of 20 bytes.
     * @param nodeID The nodeID to convert.
     * @return The fixed length nodeID.
     */
    function _fixedNodeID(
        bytes memory nodeID
    ) private pure returns (bytes20) {
        bytes20 fixedID;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            fixedID := mload(add(nodeID, 32))
        }
        return fixedID;
    }
}
