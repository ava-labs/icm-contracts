// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IValidatorManager} from "./interfaces/IValidatorManager.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {PChainOwner} from "./interfaces";

contract EmreValidator {
    address public constant EmTokenAddress; //once I deploy I should be able to put the address here
    uint8 public validatorCount = 0;
    mapping(address nodeID => bytes32 validationID) validators;
    address public constant validatorManagerAddress;

    function becomeValidator(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) external returns (bytes32) {
        require(validatorCount < 10); //check to see if there are slots open
        require(validators[nodeID] == 0); // check to see if node if addr is already validating
        validatorCount++;

        IERC20(EmTokenAddress).transferFrom(nodeID, block.number, 10);
        //potential problem, dont want to make them a validator before marking them as a validator
        //might be smart to have it point to a struct that says if its a validator and holds its validationID, but for now well try this out
        bytes32 _validationID = IValidatorManager(validatorManagerAddress)
            .initiateValidatorRegistration(
            nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, 100
        );
        validators[nodeID] = _validationID;
        return _validationID;
    }

    function stopValidating(
        bytes32 validationID
    ) external {
        require(validatorCount > 0);
        require(validators[msg.sender] == validationID);
        validators[msg.sender] = false;
        validatorCount--;
        IValidatorManager(validatorManagerAddress).initiateValidatorRemoval(validationID);
        IERC20(EmTokenAddress).transfer(msg.sender, 10);
        //call the functions for emtoken to send the money back and remove them as validator
    }
}
