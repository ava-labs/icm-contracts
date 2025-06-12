// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

// import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {IValidatorManager} from "./interfaces/IValidatorManager.sol";
import {PChainOwner} from "./interfaces/IACP99Manager.sol";
import {EmCoin} from "./EmCoin.sol";

contract SlotAuctionManager {
    address public constant EM_TOKEN_ADDRESS = address(0x0); //once I deploy I should be able to put the address here
    address public constant VALIDATOR_MANAGER_ADDRESS = address(0x0);

    uint8 public validatorCount = 0;
    mapping(address validator => bytes32 validationID) public validators;

    function becomeValidator(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) public returns (bytes32) {
        require(validatorCount < 10, "Max validator limit reached");
        require(validators[msg.sender] == 0, "Sender is already a validator");
        validatorCount++;

        EmCoin(EM_TOKEN_ADDRESS).transferFrom(msg.sender, address(this), 10);
        //potential problem, dont want to make them a validator before marking them as a validator
        //might be smart to have it point to a struct that says if its a validator and holds its validationID, but for now well try this out
        bytes32 validationID = IValidatorManager(VALIDATOR_MANAGER_ADDRESS)
        .initiateValidatorRegistration(
            nodeID,
            blsPublicKey,
            remainingBalanceOwner, 
            disableOwner, 
            100
        );
        validators[msg.sender] = validationID;
        return validationID;
    }

    function stopValidating(
        bytes32 validationID
    ) public {
        require(validatorCount > 0, "Currently no validators");
        require(validators[msg.sender] == validationID, "Invalid sender, sender is not a validator");
        validators[msg.sender] = 0;
        validatorCount--;
        IValidatorManager(VALIDATOR_MANAGER_ADDRESS).initiateValidatorRemoval(validationID);
        EmCoin(EM_TOKEN_ADDRESS).transfer(msg.sender, 10);
    }
}
