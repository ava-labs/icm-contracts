// SPDX-License-Identifier: BSD-3-Clause-Clear
pragma solidity ^0.8.25;

/**
 * @title ValidatorSet
 * @dev Contract for managing validator information using struct storage in arrays
 */
contract ValidatorSet {
    // Struct to represent validator information
    struct Validator {
        uint256 weight;
        bytes publicKey;
    }

    // Array to store validators
    Validator[] public validators;

    /**
     * @dev Public function to save/register a new validator
     * @param weight The weight of the validator
     * @param publicKey The public key of the validator
     */
    function registerValidator(
        uint256 weight,
        bytes memory publicKey
    ) public returns (uint256 validatorId) {
        // Create and push the validator struct to the array
        validators.push(Validator({weight: weight, publicKey: publicKey}));

        // Return the index (validator ID) - arrays are 0-indexed
        validatorId = validators.length - 1;
        return validatorId;
    }

    /**
     * @dev Public function to update an existing validator's information
     * @param validatorId The ID of the validator to update
     * @param weight The new weight of the validator
     * @param publicKey The new public key of the validator
     */
    function updateValidator(uint256 validatorId, uint256 weight, bytes memory publicKey) public {
        require(validatorId < validators.length, "Invalid validator ID");

        // Update the validator struct in the array
        validators[validatorId].weight = weight;
        validators[validatorId].publicKey = publicKey;
    }

    /**
     * @dev Get validator information by ID
     * @param validatorId The ID of the validator
     * @return validator The validator struct
     */
    function getValidator(
        uint256 validatorId
    ) public view returns (Validator memory) {
        require(validatorId < validators.length, "Invalid validator ID");
        return validators[validatorId];
    }

    /**
     * @dev Get the total number of registered validators
     * @return count The total count of validators
     */
    function getTotalValidators() public view returns (uint256) {
        return validators.length;
    }
}
