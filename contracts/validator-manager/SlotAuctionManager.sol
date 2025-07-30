// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {Math} from "@openzeppelin/contracts@5.0.2/utils/math/Math.sol";
import {IValidatorManager} from "./interfaces/IValidatorManager.sol";
import {PChainOwner} from "./interfaces/IACP99Manager.sol";
import {ISlotAuctionManager, 
    ValidatorBid, 
    ValidatorInfo, 
    AuctionState, 
    SlotAuctionManagerSettings,
    AuctionSettings} from 
    "./interfaces/ISlotAuctionManager.sol";
import {Heap} from "@openzeppelin/contracts@5.0.2/utils/structs/Heap.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts@5.0.2/utils/math/Math.sol";
import {Comparators} from "@openzeppelin/contracts@5.0.2/utils/Comparators.sol";
import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";

abstract contract SlotAuctionManager is 
    ISlotAuctionManager, 
    ContextUpgradeable, 
    ReentrancyGuardUpgradeable, 
    OwnableUpgradeable {

    using Heap for Heap.Uint256Heap;

    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.SlotAuctionManager
    struct SlotAuctionManagerStorage {
        IValidatorManager _manager;
        /// @notice The current state of the auction
        AuctionState _auctionState;
        /// @notice The endtime of the current running auction (No bids are allowed on auctionEndTime)
        uint256 _auctionEndTime;
        /// @notice The total amount of validator slots used or unused
        uint16 _totalValidatorSlots;
        /// @notice The total amount of validator slots being auctioned off
        uint16 _auctioningValidatorSlots;
        /// @notice The total amount of validator slots ready to be auctioned off
        uint16 _openValidatorSlots;
        /// @notice The weight of validator slots being auctioned off
        uint64 _validatorWeight;
        /// @notice The minimum amount of time a validator is allowed to validate
        uint256 _MinValidatorDuration;
        /// @notice The minimum amount of time the auction must run
        uint256 _MinAuctionDuration;
        /// @notice The minimum bid to be placed into the auction
        uint256 _MinimumBid;
        /// @notice The amount of time until the next auciton can be initiated
        uint256 _auctionCooldownEndtime;
        /// @notice Maps nodeIDs to boolean to check if nodeID is already in the running and capable of winning
        mapping(bytes nodeID => bool isQualified) _nodeIsQualified;
        /// @notice Maps nodeIDs to a struct holding validator information
        mapping(bytes nodeID => ValidatorInfo) _validatorsByNodeID;
        /// @notice Maps ValidationIDs to a struct holding validator information
        mapping(bytes32 validationID => ValidatorInfo) _validatorsByValidationID;
        /// @notice Maps bid amount to a struct containing the bidders info
        mapping(uint256 bid => ValidatorBid) _bidderInfo;
        /// @notice The current second price for the lowest slot winner, is 0 if no second price
        uint256 _secondPrice;
        /// @notice The data structure holding auction bid information
        Heap.Uint256Heap _bids;
    }
    // solhint-enable private-vars-leading-underscore
    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.SlotAuctionManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant SLOT_AUCTION_MANAGER_STORAGE_LOCATION = 
        0x6494e1e6ac2d5f44a06b929c1549cec0c499347f244a852c37aef6b6707be600;

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
    error AuctionInCooldown(uint256 auctionCooldownEndtime);
    error InvalidWeight(uint64 weight);
    error InvalidMinValidatorDuration(uint256 minimumValidationDuration);
    error ZeroAddress();
    error ZeroMinBid();

    // solhint-disable ordering

    modifier AuctionOn() {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        if ($._auctionState != AuctionState.AuctionInProgress) {
            revert AuctionNotInProgress();
        }
        _;
    }

    modifier AuctionOff() {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        if ($._auctionState == AuctionState.AuctionInProgress) {
            revert AuctionInProgress();
        }
        if ($._auctionState == AuctionState.AuctionFinalizing) {
            revert AuctionFinalizing();
        }
        _;
    }

    modifier AuctionCooldown() {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        if (block.timestamp < $._auctionCooldownEndtime) {
            revert AuctionInCooldown($._auctionCooldownEndtime);
        }
        _;
    }


    // solhint-disable-next-line func-name-mixedcase
    function __SlotAuctionManager_init(
        SlotAuctionManagerSettings calldata settings
    ) internal onlyInitializing {
        __ReentrancyGuard_init();
        __Ownable_init(settings.admin);
        __SlotAuctionManager_init_unchained(
            settings.manager,
            settings.auctionSettings
        );
    }

    // solhint-disable-next-line func-name-mixedcase
    function __SlotAuctionManager_init_unchained(
        IValidatorManager manager,
        AuctionSettings calldata auctionSettings
    ) internal onlyInitializing {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        if (address(manager) == address(0)) {
            revert ZeroAddress();
        }
        // Minimum stake duration should be at least one churn period in order to prevent churn tracker abuse.
        if (auctionSettings.MinValidatorDuration < manager.getChurnPeriodSeconds()) {
            revert InvalidMinValidatorDuration(auctionSettings.MinValidatorDuration);
        }
        if (auctionSettings.Weight == 0) {
            revert InvalidWeight(auctionSettings.Weight);
        }
        if (auctionSettings.MinimumBid == 0) {
            revert ZeroMinBid();
        }

        $._manager = manager;
        $._totalValidatorSlots = auctionSettings.totalValidatorSlots;
        $._openValidatorSlots = auctionSettings.totalValidatorSlots;
        $._validatorWeight = auctionSettings.Weight;
        $._MinValidatorDuration = auctionSettings.MinValidatorDuration;
        $._MinAuctionDuration = auctionSettings.MinAuctionDuration;
        $._MinimumBid = auctionSettings.MinimumBid;
    }


    function _getSlotAuctionManagerStorage() internal pure returns (SlotAuctionManagerStorage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := SLOT_AUCTION_MANAGER_STORAGE_LOCATION
        }
    }

    function initiateAuction() external AuctionOff AuctionCooldown {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        $._auctionState = AuctionState.AuctionInProgress;
        $._auctionEndTime = block.timestamp + $._MinAuctionDuration;
        $._secondPrice = 0;
        if ($._openValidatorSlots == 0) {
            revert NoOpenValidatorSlots($._openValidatorSlots);
        }
        // Gets maximum amount of validators that are able to be auctioned off without triggering churn, making sure its not more
        // than the current amount of available slots
        uint64 maxValidatorSlotsBeforeChurn = $._manager.getMaximumChurnPercentage()
            * $._manager.l1TotalWeight() / $._validatorWeight * 100;

        if (maxValidatorSlotsBeforeChurn == 0) {
            revert ValidatorWeightTooHigh($._validatorWeight);
        }
        if (maxValidatorSlotsBeforeChurn > $._openValidatorSlots) {
            $._auctioningValidatorSlots = $._openValidatorSlots;
        } else {
            $._auctioningValidatorSlots = uint16(maxValidatorSlotsBeforeChurn);
        }
        emit NewValidatorAuction(
            $._auctioningValidatorSlots,
            $._validatorWeight,
            $._MinValidatorDuration,
            $._auctionEndTime,
            $._MinimumBid
        );
    }

    function endAuction() external nonReentrant AuctionOn {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        $._auctionCooldownEndtime = block.timestamp + 1 days;
        $._auctionState = AuctionState.AuctionFinalizing;
        if (block.timestamp < $._auctionEndTime) {
            revert AuctionEndTimeNotReached($._auctionEndTime);
        }

        // avoids array out of bounds for Heap.peek
        if ($._secondPrice == 0 && Heap.length($._bids) != 0) {
            $._secondPrice = Heap.peek($._bids);
        }

        while (Heap.length($._bids) > 0) {
            uint256 currentBid = Heap.pop($._bids);
            ValidatorBid memory bidInfo = $._bidderInfo[currentBid];
            delete $._bidderInfo[currentBid];
            delete $._nodeIsQualified[bidInfo.nodeID];
            // sends back extra tokens due to second price
            if (currentBid - $._secondPrice != 0) {
                _unlock(bidInfo.addr, currentBid - $._secondPrice);
            }
            bytes32 validationID = $._manager.initiateValidatorRegistration(
                bidInfo.nodeID,
                bidInfo.blsPublicKey,
                bidInfo.remainingBalanceOwner,
                bidInfo.disableOwner,
                $._validatorWeight
            );
            emit InitiatedAuctionValidatorRegistration(
                validationID, bidInfo.addr, $._MinValidatorDuration + $._auctionEndTime, $._validatorWeight
            );
            $._validatorsByNodeID[bidInfo.nodeID] = ValidatorInfo(
                bidInfo.addr,
                $._MinValidatorDuration + $._auctionEndTime,
                bidInfo.nodeID,
                bidInfo.blsPublicKey,
                validationID,
                $._validatorWeight
            );

            $._validatorsByValidationID[validationID] = ValidatorInfo(
                bidInfo.addr,
                $._MinValidatorDuration + $._auctionEndTime,
                bidInfo.nodeID,
                bidInfo.blsPublicKey,
                validationID,
                $._validatorWeight
            );

            $._secondPrice = currentBid;
        }

        $._auctionEndTime = 0;
        $._auctioningValidatorSlots = 0;
        $._auctionState = AuctionState.NoAuction;
    }

    function initiateValidatorRemoval(
        bytes32 validationID
    ) external {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        // if validationID doesnt exist then endtime will be 0, however it wont be logged in the Validator Manager either so this should be ok
        if ($._validatorsByValidationID[validationID].endTime > block.timestamp) {
            revert ValidatorTimeLimitNotPassed($._validatorsByValidationID[validationID].endTime);
        }
        delete $._validatorsByNodeID[$._validatorsByValidationID[validationID].nodeID];
        delete $._validatorsByValidationID[validationID];
        $._openValidatorSlots++;
        $._manager.initiateValidatorRemoval(validationID);
    }

    // working on removing this, along with making the contract upgradeable, removing it right now makes the e2e tests fail
    function initiateRemoveInitialValidator(
        bytes32 validationID
    ) external {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        $._manager.initiateValidatorRemoval(validationID);
    }

    function completeRemoveInitialValidator(
        uint32 messageIndex
    ) external {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        $._manager.completeValidatorRemoval(messageIndex);
    }

    function completeValidatorRegistration(
        uint32 messageIndex
    ) external returns (bytes32) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        $._openValidatorSlots--;
        return $._manager.completeValidatorRegistration(messageIndex);
    }

    function completeValidatorRemoval(
        uint32 messageIndex
    ) external returns (bytes32) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        return $._manager.completeValidatorRemoval(messageIndex);
    }

    function setSlotAuctionSettings(
        AuctionSettings calldata auctionSettings
    ) external AuctionOff onlyOwner {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        
        // OnlyOwner
        // Make sure validator slots are not pushed below the current amount of slotted validators
        if (auctionSettings.totalValidatorSlots > $._totalValidatorSlots) {
            $._openValidatorSlots += auctionSettings.totalValidatorSlots - $._totalValidatorSlots;
            $._totalValidatorSlots = auctionSettings.totalValidatorSlots;
        } else {
            if ($._totalValidatorSlots - $._openValidatorSlots > auctionSettings.totalValidatorSlots) {
                revert MoreActiveValidatorsThanTotalSlots(
                    auctionSettings.totalValidatorSlots, $._totalValidatorSlots - $._openValidatorSlots
                );
            }
            $._openValidatorSlots -= $._totalValidatorSlots - auctionSettings.totalValidatorSlots;
            $._totalValidatorSlots = auctionSettings.totalValidatorSlots;
        }
        $._MinAuctionDuration = auctionSettings.MinAuctionDuration;
        $._MinValidatorDuration = auctionSettings.MinValidatorDuration;
        $._MinimumBid = auctionSettings.MinimumBid;
        $._validatorWeight = auctionSettings.Weight;
    }

    function MinBidRequired() external view AuctionOn returns (uint256) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        if (Heap.length($._bids) < $._totalValidatorSlots) {
            return $._MinimumBid;
        }
        return Heap.peek($._bids) + 1;
    }

    function _placeBid(
        uint256 bid,
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) internal AuctionOn {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        if ($._manager.getNodeValidationID(nodeID) != 0) {
            revert NodeIsValidator(nodeID);
        }
        if ($._nodeIsQualified[nodeID]) {
            revert DuplicateNodeIDInContention(nodeID);
        }
        if ($._bidderInfo[bid].addr != address(0)) {
            revert DuplicateBidInContention(bid);
        }
        if ($._MinimumBid > bid) {
            revert BidSmallerThanMinimum($._MinimumBid, bid);
        }

        // If all slots aren't contended, then fill the heap with any bid
        if (Heap.length($._bids) < $._auctioningValidatorSlots) {
            _lock(bid);
            Heap.insert($._bids, bid);
        } else if (Heap.peek($._bids) < bid) {
            _lock(bid);

            $._secondPrice = Heap.replace($._bids, bid); 
            emit BidEvicted($._secondPrice, $._bidderInfo[$._secondPrice].nodeID);
            // send back held funds if lost auction
            _unlock($._bidderInfo[$._secondPrice].addr, $._secondPrice);

            // deletes info of bidder no longer needed along with replacing it in the heap
            delete $._nodeIsQualified[$._bidderInfo[$._secondPrice].nodeID];
            delete $._bidderInfo[$._secondPrice];
        } else {
            revert InsufficientBidToWinAuction(Heap.peek($._bids) + 1, bid);
        }
        $._bidderInfo[bid] =
            ValidatorBid(msg.sender, bid, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner);
        $._nodeIsQualified[nodeID] = true;
        emit SuccessfulBidPlaced(bid, nodeID);
    }

    function getTotalValidatorSlots() external view returns (uint16){
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._totalValidatorSlots;
    }

    function getValidatorWeight() external view returns (uint64){
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._validatorWeight;
    }

    function getMinAuctionDuration() external view returns (uint256){
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._MinAuctionDuration;
    }

    function getMinValidatorDuration() external view returns (uint256){
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._MinValidatorDuration;
    }

    function getMinimumBid() external view returns (uint256){
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._MinimumBid;
    }

    function getOpenValidatorSlots() external view returns (uint16){
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._openValidatorSlots;
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
    function _unlock(address to, uint256 value) internal virtual;
}
