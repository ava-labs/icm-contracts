// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {Math} from "@openzeppelin/contracts@5.0.2/utils/math/Math.sol";
import {IValidatorManager} from "./interfaces/IValidatorManager.sol";
import {PChainOwner} from "./interfaces/IACP99Manager.sol";
import {ISlotAuctionManager, ValidatorBid, ValidatorInfo, AuctionState} from "./interfaces/ISlotAuctionManager.sol";
import {Heap} from "@openzeppelin/contracts@5.0.2/utils/structs/Heap.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts@5.0.2/utils/math/Math.sol";
import {Comparators} from "@openzeppelin/contracts@5.0.2/utils/Comparators.sol";
// import {ContextUpgradeable} from
//     "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";
// import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";


// TODO: Make AuctionStatus an ENUM and edit modifiers and checks

abstract contract SlotAuctionManager is ReentrancyGuardUpgradeable, ISlotAuctionManager {
    using Heap for Heap.Uint256Heap;

    IValidatorManager public VALIDATOR_MANAGER;
    AuctionState public auctionState;
    // auctionEndTime is the end of the auction (no bids)
    uint256 public auctionEndTime;
    // Total amount of validator slots in the set
    uint16 public totalValidatorSlots;
    // Amount of validator slots being auctioned right now
    uint16 public auctioningValidatorSlots;
    // Amount of validator slots unused
    uint16 public openValidatorSlots;
    uint64 public validatorWeight;
    uint256 public MinValidatorDuration;
    uint256 public MinAuctionDuration;
    uint256 public MinimumBid;
    // checks to see if the NodeID is currently in the heap
    mapping (bytes nodeID => bool isQualified) private _nodeIsQualified; 
    mapping (bytes nodeID => ValidatorInfo) private validatorsByNodeID;
    mapping (bytes32 validationID => ValidatorInfo) private validatorsByValidationID;
    mapping (uint256 bid => ValidatorBid) public bidderInfo;
    uint256 private _secondPrice;
    
    Heap.Uint256Heap internal _bids;

    error AuctionInProgress();
    error AuctionNotInProgress();
    error NodeIsValidator(bytes nodeID);
    error DuplicateNodeIDInContention(bytes nodeID);
    error DuplicateBidInContention(uint256 bid);
    error BidSmallerThanMinimum(uint256 minumumBid, uint256 userBid);
    error InsufficientBidToWinAuction(uint256 smallestAcceptableBid, uint256 userBid);
    error AuctionEndTimeNotReached(uint256 auctionEndTime);
    error ValidatorTimeLimitNotPassed(uint256 validationTimeLimit);
    error AuctionFinalizing();
    error ValidatorWeightTooHigh(uint64 validatorWeight);
    error NoOpenValidatorSlots(uint16 validatorSlots);
    error MoreActiveValidatorsThanTotalSlots(uint16 totalValidatorSlots, uint16 activeValidators);

    modifier AuctionOn {
        if (auctionState != AuctionState.AuctionInProgress) {
            revert AuctionNotInProgress();
        }
        _;
    }

    modifier AuctionOff {
        if (auctionState == AuctionState.AuctionInProgress) {
            revert AuctionInProgress();
        }
        if (auctionState == AuctionState.AuctionFinalizing) {
            revert AuctionFinalizing();
        }
        _;
    }

    function initiateAuction() AuctionOff external {
        auctionState = AuctionState.AuctionInProgress;
        auctionEndTime = block.timestamp + MinAuctionDuration; 
        _secondPrice = 0;
        if (openValidatorSlots == 0) {
            revert NoOpenValidatorSlots(openValidatorSlots);
        }
        // Gets maximum amount of validators that are able to be auctioned off without triggering churn, making sure its not more
        // than the current amount of available slots
        uint64 maxValidatorSlotsBeforeChurn = VALIDATOR_MANAGER.getMaximumChurnPercentage() * VALIDATOR_MANAGER.l1TotalWeight() / validatorWeight * 100;
        
        if (maxValidatorSlotsBeforeChurn == 0){
            revert ValidatorWeightTooHigh(validatorWeight);
        }
        if (maxValidatorSlotsBeforeChurn > openValidatorSlots) {
            auctioningValidatorSlots = openValidatorSlots;
        }
        else {
            auctioningValidatorSlots = uint16(maxValidatorSlotsBeforeChurn);
        }
        // Creates new heap for bids
        delete _bids.tree; 
        emit NewValidatorAuction(auctioningValidatorSlots, validatorWeight, MinValidatorDuration, auctionEndTime, MinimumBid);
    }

    function endAuction() nonReentrant AuctionOn external {

        auctionState = AuctionState.AuctionFinalizing;
        if (block.timestamp <= auctionEndTime) {
            revert AuctionEndTimeNotReached(auctionEndTime);
        }
        auctionEndTime = 0;

        // avoids array out of bounds for Heap.peek
        if (_secondPrice == 0 && Heap.length(_bids) != 0) { 
            _secondPrice = Heap.peek(_bids);
        }

        while (Heap.length(_bids) > 0) {
            uint256 currentBid = Heap.pop(_bids);
            ValidatorBid memory bidInfo = bidderInfo[currentBid];

            // sends back extra tokens due to second price
            _unlock(bidInfo.addr, currentBid - _secondPrice);
            // TOKEN_CONTRACT.transfer(bidInfo.addr, currentBid - _secondPrice);
            
            bytes32 validationID = VALIDATOR_MANAGER.initiateValidatorRegistration(
                bidInfo.nodeID, bidInfo.blsPublicKey, bidInfo.remainingBalanceOwner, bidInfo.disableOwner, validatorWeight
            );

            emit InitiatedAuctionValidatorRegistration(validationID, bidInfo.addr, MinValidatorDuration + auctionEndTime, validatorWeight);
            validatorsByNodeID[bidInfo.nodeID] = ValidatorInfo(bidInfo.addr, MinValidatorDuration + auctionEndTime, 
                bidInfo.nodeID, bidInfo.blsPublicKey, validationID, validatorWeight);

            validatorsByValidationID[validationID] = ValidatorInfo(bidInfo.addr, MinValidatorDuration + auctionEndTime, 
                bidInfo.nodeID, bidInfo.blsPublicKey, validationID, validatorWeight);

            _secondPrice = currentBid;
        }
        
        auctioningValidatorSlots = 0;
        auctionState = AuctionState.NoAuction;
    }


    function initiateValidatorRemoval(
        bytes32 validationID
    ) external {
        // if validationID doesnt exist then endtime will be 0, however it wont be logged in the Validator Manager either so this should be ok
        if (validatorsByValidationID[validationID].endTime > block.timestamp) {
            revert ValidatorTimeLimitNotPassed(validatorsByValidationID[validationID].endTime);
        }
        delete validatorsByNodeID[validatorsByValidationID[validationID].nodeID];
        delete validatorsByValidationID[validationID];
        openValidatorSlots++;
        VALIDATOR_MANAGER.initiateValidatorRemoval(validationID);
    }
    
    // working on removing this, along with making the contract upgradeable, removing it right now makes the e2e tests fail
    function initiateRemoveInitialValidator(
        bytes32 validationID
    ) external {
        VALIDATOR_MANAGER.initiateValidatorRemoval(validationID);
    }

    function completeRemoveInitialValidator(
        uint32 messageIndex
    ) external {
        VALIDATOR_MANAGER.completeValidatorRemoval(messageIndex);
    }

    function completeValidatorRegistration(
        uint32 messageIndex
    ) external returns (bytes32) {
        return VALIDATOR_MANAGER.completeValidatorRegistration(messageIndex);
    }

    function completeValidatorRemoval(
        uint32 messageIndex
    ) external returns (bytes32) {
        return VALIDATOR_MANAGER.completeValidatorRemoval(messageIndex);
    }
    
    function setSlotAuctionSettings(
        uint16 newValidatorSlots,
        uint64 newWeight,
        uint256 newMinAuctionDuration,
        uint256 newMinValidatorDuration,
        uint256 newMinimumBid
    ) external AuctionOff { // OnlyOwner
        // Make sure validator slots are not pushed below the current amount of slotted validators
        if (newValidatorSlots > totalValidatorSlots) {
            openValidatorSlots += newValidatorSlots - totalValidatorSlots;
            totalValidatorSlots = newValidatorSlots;
        }
        else {
            if (totalValidatorSlots - openValidatorSlots > newValidatorSlots) {
                revert MoreActiveValidatorsThanTotalSlots(newValidatorSlots, totalValidatorSlots - openValidatorSlots);
            }
        }
        MinAuctionDuration = newMinAuctionDuration;
        MinValidatorDuration = newMinValidatorDuration;
        MinimumBid = newMinimumBid;
        validatorWeight = newWeight;

    }

    function MinBidRequired() AuctionOn external view returns (uint256){
        if (Heap.length(_bids) < totalValidatorSlots) {
            return MinimumBid;
        }
        return Heap.peek(_bids) + 1;
    }

    function _setSlotAuctionSettings(
        uint16 validatorslots,
        uint64 weight,
        uint256 minAuctionDuration,
        uint256 minValidatorDuration,
        uint256 minimumBid
    ) internal {
        // sets up default values of auction
        totalValidatorSlots = validatorslots;
        openValidatorSlots = totalValidatorSlots;
        validatorWeight = weight;
        MinAuctionDuration = minAuctionDuration; 
        MinValidatorDuration = minValidatorDuration;
        MinimumBid = minimumBid;
    } 

    function _placeBid (
        uint256 bid,
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) AuctionOn internal {
        if (VALIDATOR_MANAGER.getNodeValidationID(nodeID) != 0) {
            revert NodeIsValidator(nodeID);
        }
        if (_nodeIsQualified[nodeID]) {
            revert DuplicateNodeIDInContention(nodeID);
        }
        if (bidderInfo[bid].addr != address(0)) {
            revert DuplicateBidInContention(bid);
        }
        if (MinimumBid > bid) {
            revert BidSmallerThanMinimum(MinimumBid, bid);
        }

        // If all slots aren't contended, then fill the heap with any bid
        if (Heap.length(_bids) < auctioningValidatorSlots) {
            _lock(bid);
            Heap.insert(_bids, bid);
        } 
        else if (Heap.peek(_bids) < bid) {
            _lock(bid);
        
            uint256 poppedBid = Heap.replace(_bids, bid);
            _secondPrice = poppedBid;
            emit BidEvicted(poppedBid, bidderInfo[poppedBid].nodeID);
            // send back held funds if lost auction
            _unlock(bidderInfo[poppedBid].addr, poppedBid);

            // deletes info of bidder no longer needed along with replacing it in the heap
            delete _nodeIsQualified[bidderInfo[poppedBid].nodeID];
            delete bidderInfo[poppedBid]; 
        }
        else {
            revert InsufficientBidToWinAuction(Heap.peek(_bids) + 1, bid);
        }
        bidderInfo[bid] = ValidatorBid(msg.sender, bid, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner);
        _nodeIsQualified[nodeID] = true;
        emit SuccessfulBidPlaced(bid, nodeID);
    }

    /**
     * @notice Locks tokens in this contract.
     * @param value Number of tokens to lock.
     */
    function _lock(
        uint256 value
    ) internal virtual returns (uint256);

    /**
     * @notice Unlocks token to a specific address.
     * @param to Address to send token to.
     * @param value Number of tokens to lock.
     */
    function _unlock(
        address to, 
        uint256 value
    ) internal virtual;
}