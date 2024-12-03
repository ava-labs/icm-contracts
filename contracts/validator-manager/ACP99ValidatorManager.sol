// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

import {IACP99SecurityModule} from "./interfaces/IACP99SecurityModule.sol";
import {ValidatorManager} from "./ValidatorManager.sol";
import {IACP99ValidatorManager, ConversionData, ValidatorRegistrationInput} from "./interfaces/IACP99ValidatorManager.sol";

pragma solidity 0.8.25;

abstract contract ACP99ValidatorManager is IACP99ValidatorManager, ValidatorManager {
    IACP99SecurityModule public securityModule;

    // TODO: calling this should be restricted to...who?
    function setSecurityModule(IACP99SecurityModule _securityModule) external {
        securityModule = _securityModule;
    }

    function initializeValidatorSet(
        ConversionData calldata conversionData,
        uint32 messageIndex
    ) external {
        _initializeValidatorSet(conversionData, messageIndex);
    }

    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata input,
        uint64 weight,
        bytes calldata args
    ) external returns (bytes32){
        bytes32 validationID = _initializeValidatorRegistration(input, weight);
        securityModule.handleInitializeValidatorRegistration(validationID, weight, args);
        return validationID;
    }

    function completeValidatorRegistration(uint32 messageIndex) external{
        bytes32 validationID = _completeValidatorRegistration(messageIndex);
        securityModule.handleCompleteValidatorRegistration(validationID);
    }

    function initializeEndValidation(bytes32 validationID, bytes calldata args) external{
        _initializeEndValidation(validationID);
        securityModule.handleInitializeEndValidation(validationID, args);
    }

    function completeEndValidation(uint32 messageIndex) external{
        bytes32 validationID = _completeEndValidation(messageIndex);
        securityModule.handleCompleteEndValidation(validationID);
    }

    function initializeValidatorWeightChange(bytes32 validationID, uint64 weight, bytes calldata args) external{
        securityModule.handleInitializeValidatorWeightChange(validationID, weight, args);
    }

    function completeValidatorWeightChange(bytes32 validationID, bytes calldata args) external{
        securityModule.handleCompleteValidatorWeightChange(validationID, args);
    }
}