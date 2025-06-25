// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {IValidatorManager} from "./interfaces/IValidatorManager.sol";
import {PChainOwner} from "./interfaces/IACP99Manager.sol";
import {EmCoin} from "./EmCoin.sol";
import {Heap} from "@openzeppelin/contracts@5.0.2/utils/structs/Heap.sol";
//TODO: make these constant and also public, prob some way to store it in memory or something from here

struct ValidatorBid {
    address addr;
    uint256 bid;
    bytes nodeID;
    bytes blsPublicKey;
    PChainOwner remainingBalanceOwner;
    PChainOwner disableOwner;
}

contract SlotAuctionManager {
    IERC20 public TOKEN_CONTRACT;
    IValidatorManager public VALIDATOR_MANAGER;
    bool public auctionInProgress = false;
    uint256 public auctionEndTime = 0;
    uint16 public ValidatorSlots = 0;
    mapping (bytes nodeID => bool hasBid) public nodeHasBid;
    mapping (uint256 bid => bool duplicate) public duplicateBid;
    uint64 public validatorWeight = 0;
    constructor(address tokenAddress, address vmAddress) {
        TOKEN_CONTRACT = IERC20(tokenAddress);
        VALIDATOR_MANAGER = IValidatorManager(vmAddress);
    }
    
    function initiateAuction(
        uint16 validatorSlots,
        uint64 weight,
        uint256 auctionLength
    ) external {
        require(!auctionInProgress, "Auction already running");
        validatorWeight = weight;
        //require(only owner can call this contract)
        //require(churn percentage will not be overlooked)
        //Things to add:
        //no re entry
        //probably a lot more
        auctionEndTime = block.timestamp + auctionLength; //sets the minimum end time of the auction, can go over
        //create a heap here
        ValidatorSlots = validatorSlots;
    }

    function endAuction() external returns (bool) {
        //only owner
        require(auctionInProgress, "Auction not in progress");
        require(block.timestamp > auctionEndTime, "Before set end of auction time"); //gotta word this better
        auctionInProgress = false; //set auction to false so no more bids can come in

        //while loop through the heap, labeling them as a validator, the first guy gets his money back
        //check for lots of edge cases where 

    }

    function placeBid(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint256 bid
    ) external {
        require(auctionInProgress, "Auction is not running");
        nodeHasBid[nodeID] = true; //immediately label them as bid, this makes it so they cant spam bid even if they already bid a smaller amount
        if (true) { // change this to peek the current root and check to see if the newest bid is large enough to take it over
            return; 
        }
        require (TOKEN_CONTRACT.transferFrom(msg.sender, address(this), bid), "Insufficient funds based on bid");
    //once funda are recieved, assuming the p

    }

    function initiateValidatorRemoval(
        bytes32 validationID
    ) public {
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

    function comparison (
        uint256 parent, 
        uint256 child
    ) internal view returns (bool) {
        return parent < child; //should be a min queue
    }

    function _completeValidatorRegistration(
        uint32 messageIndex
    ) internal returns (bytes32) {
        return VALIDATOR_MANAGER.completeValidatorRegistration(messageIndex);
    }

    function _completeValidatorRemoval(
        uint32 messageIndex
    ) internal returns (bytes32) {
        return VALIDATOR_MANAGER.completeValidatorRemoval(messageIndex);
    }

    function _initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint64 weight
    ) internal returns (bytes32) {
        require(
            TOKEN_CONTRACT.transferFrom(msg.sender, address(this), 10) == true,
            "Tokens failed to transfer"
        );

        //potential problem, dont want to make them a validator before marking them as a validator
        //might be smart to have it point to a struct that says if its a validator and holds its validationID, but for now well try this out
        bytes32 validationID = VALIDATOR_MANAGER.initiateValidatorRegistration(
            nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, weight
        );
        return validationID;
    }
}   
