// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {IValidatorManager} from "./interfaces/IValidatorManager.sol";
import {PChainOwner} from "./interfaces/IACP99Manager.sol";
import {EmCoin} from "./EmCoin.sol";
//TODO: make these constant and also public, prob some way to store it in memory or something from here
contract SlotAuctionManager {
    IERC20 public TOKEN_CONTRACT;
    IValidatorManager public VALIDATOR_MANAGER;
    bytes32 public TemporaryID = 0;
    uint8 public validatorCount = 0;

    constructor(address tokenAddress, address vmAddress) {
        TOKEN_CONTRACT = IERC20(tokenAddress);
        VALIDATOR_MANAGER = IValidatorManager(vmAddress);
    }
    
    function becomeValidator(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) public returns (bytes32) {
        require(validatorCount < 10, "Max validator limit reached");
        validatorCount++;

        require(
            TOKEN_CONTRACT.transferFrom(msg.sender, address(this), 10) == true,
            "Tokens failed to transfer"
        );

        //potential problem, dont want to make them a validator before marking them as a validator
        //might be smart to have it point to a struct that says if its a validator and holds its validationID, but for now well try this out
        bytes32 validationID = VALIDATOR_MANAGER.initiateValidatorRegistration(
            nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, 100
        );
        TemporaryID = validationID;
        return validationID;
    }

    function removeValidator(
        bytes32 validationID
    ) public {
        require(validatorCount > 0, "Currently no validators");
        validatorCount--;
        VALIDATOR_MANAGER.initiateValidatorRemoval(validationID);
        require(TOKEN_CONTRACT.transfer(msg.sender, 10), "Funds failed to send");
    }

    function initiateRemoveInitialValidator(
        bytes32 validationID
    ) public {
        VALIDATOR_MANAGER.initiateValidatorRemoval(validationID);
    }

    function completeRemoveInitialValidator(
        uint32 messageIndex
    ) public {
        VALIDATOR_MANAGER.completeValidatorRemoval(messageIndex);
    }
}   
