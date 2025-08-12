// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IValidatorManager, ValidatorChurnPeriod} from "./interfaces/IValidatorManager.sol";
import {PChainOwner} from "./interfaces/IACP99Manager.sol";
import {
    ISlotAuctionManager,
    ValidatorBid,
    ValidatorInfo,
    AuctionState,
    SlotAuctionManagerSettings,
    AuctionSettings,
    ValidatorVoucher
} from "./interfaces/ISlotAuctionManager.sol";
import {Heap} from "@openzeppelin/contracts@5.0.2/utils/structs/Heap.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ReentrancyGuardUpgradeable.sol";
import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";

/**
 * @dev Implementation of the {ISlotAuctionManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
abstract contract SlotAuctionManager is
    ISlotAuctionManager,
    ContextUpgradeable,
    ReentrancyGuardUpgradeable,
    OwnableUpgradeable
{
    using Heap for Heap.Uint256Heap;

    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.SlotAuctionManager
    struct SlotAuctionManagerStorage {
        IValidatorManager _manager;
        /// @notice The current state of the auction
        AuctionState _auctionState;
        /// @notice The total amount of validator slots used or unused
        uint16 _totalValidatorSlots;
        /// @notice The total amount of validator slots being auctioned off
        uint16 _auctioningValidatorSlots;
        /// @notice The total amount of validator slots ready to be auctioned off
        uint16 _occupiedValidatorSlots;
        /// @notice The weight of validator slots being auctioned off
        uint64 _auctioningValidatorWeight;
        /// @notice The minimum amount of time a validator is allowed to validate
        uint256 _minValidatorDuration;
        /// @notice The minimum amount of time the auction must run
        uint256 _minAuctionDuration;
        /// @notice The minimum bid to be placed into the auction
        uint256 _minimumBid;
        /// @notice The length of time for the auction cooldown
        uint256 _auctionCooldownDuration;
        /// @notice The timestamp of when the next auction can be initiated
        uint256 _auctionCooldownEndtime;
        /// @notice The endtime of the current running auction (No bids are allowed on auctionEndTime)
        uint256 _auctionEndTime;
        /// @notice Maps nodeIDs to boolean to check if nodeID is already in the running and capable of winning
        uint256 _secondPrice;
        /// @notice The data structure holding auction bid information
        mapping(bytes nodeID => bool isQualified) _nodeIsQualified;
        /// @notice Maps ValidationIDs to a struct holding validator information
        mapping(bytes32 validationID => ValidatorInfo) _validatorsByValidationID;
        /// @notice Maps bid amount to a struct containing the bidders info
        mapping(uint256 bid => ValidatorBid) _bidderInfo;
        /// @notice Maps nodeID to struct ValidatorVoucher
        mapping(bytes nodeID => ValidatorVoucher) _vouchers;
        /// @notice The current second price for the lowest slot winner, is 0 if no second price
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
    error ValidatorWeightTooHigh(uint64 validatorWeight);
    error NoOpenValidatorSlots();
    error AuctionInCooldown(uint256 auctionCooldownEndtime);
    error ZeroWeight();
    error InvalidMinValidatorDuration(uint256 minimumValidationDuration);
    error ZeroAddress();
    error ZeroMinBid();
    error VaidatorRegistrationTimePeriodOver(uint256 endTime, uint256 currentTime, bytes nodeID); 

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

        if ($._auctionState != AuctionState.NoAuction) {
            revert AuctionInProgress();
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
        __SlotAuctionManager_init_unchained(settings.manager, settings.auctionSettings);
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
        _validateSettings(auctionSettings);

        $._manager = manager;
        $._totalValidatorSlots = auctionSettings.totalValidatorSlots;
        $._occupiedValidatorSlots = 0;
        $._auctioningValidatorWeight = auctionSettings.weight;
        $._minValidatorDuration = auctionSettings.minValidatorDuration;
        $._minAuctionDuration = auctionSettings.minAuctionDuration;
        $._minimumBid = auctionSettings.minimumBid;
        $._auctionCooldownDuration = auctionSettings.auctionCooldownDuration;
    }

    function _getSlotAuctionManagerStorage()
        internal
        pure
        returns (SlotAuctionManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := SLOT_AUCTION_MANAGER_STORAGE_LOCATION
        }
    }

    function initiateAuction() external AuctionOff AuctionCooldown {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        $._auctionState = AuctionState.AuctionInProgress;
        $._auctionEndTime = block.timestamp + $._minAuctionDuration;
        $._secondPrice = $._minimumBid;
        if ($._occupiedValidatorSlots == $._totalValidatorSlots) {
            revert NoOpenValidatorSlots();
        }
        // Gets maximum amount of validators that are able to be auctioned off without triggering churn.
        // There is still the possibility churn is triggered due to the ending of an auction potentially happening
        // in a different churn period.
        ( , uint8 maxChurnPercentage, ValidatorChurnPeriod memory churnPeriod)= $._manager.getChurnTracker();
        uint64 maxValidatorSlotsBeforeChurn = (
            (maxChurnPercentage / 2) * churnPeriod.totalWeight
        ) / ($._auctioningValidatorWeight * 100);

        if (maxValidatorSlotsBeforeChurn == 0) {
            revert ValidatorWeightTooHigh($._auctioningValidatorWeight);
        }
        if (maxValidatorSlotsBeforeChurn > $._totalValidatorSlots - $._occupiedValidatorSlots) {
            $._auctioningValidatorSlots = $._totalValidatorSlots - $._occupiedValidatorSlots;
        } else {
            $._auctioningValidatorSlots = uint16(maxValidatorSlotsBeforeChurn);
        }
        emit NewValidatorAuction(
            $._auctioningValidatorSlots,
            $._auctioningValidatorWeight,
            $._minValidatorDuration,
            $._auctionEndTime,
            $._minimumBid
        );
    }

    function endAuction() external nonReentrant AuctionOn {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        if (block.timestamp < $._auctionEndTime) {
            revert AuctionEndTimeNotReached($._auctionEndTime);
        }

        $._auctionState = AuctionState.AuctionFinalizing;

        // Checks to see if the current amount of validators to be registerd will trigger churn. If so,
        // will end the auction by distributing vouchers, which allows anyone to register the nodeID
        // as a validator between the end of the auction and the validators end time.
        bool safeMode = false;
        ( , uint8 maxChurnPercentage, ValidatorChurnPeriod memory churnPeriod)= $._manager.getChurnTracker();
        if (maxChurnPercentage * churnPeriod.initialWeight < 
            (churnPeriod.churnAmount + (Heap.length($._bids) * $._auctioningValidatorWeight)) * 100) {
                safeMode = true;
            }

        $._auctionCooldownEndtime = block.timestamp + $._auctionCooldownDuration;
        while (Heap.length($._bids) > 0) {
            uint256 currentBid = Heap.pop($._bids);
            ValidatorBid memory bidInfo = $._bidderInfo[currentBid];
            delete $._bidderInfo[currentBid];
            delete $._nodeIsQualified[bidInfo.nodeID];
            // sends back extra tokens due to second price
            if (currentBid - $._secondPrice != 0) {
                _unlock(bidInfo.addr, currentBid - $._secondPrice);
            }
            if (safeMode) {
                $._vouchers[node] = ValidatorVoucher({
                    addr: bidInfo.addr,
                    endTime: $._minValidatorDuration + $._auctionEndTime,
                    registrationEndTime: $._auctionCooldownEndtime,
                    nodeID: bidInfo.nodeID,
                    blsPublicKey: bidInfo.blsPublicKey,
                    weight: $._auctioningValidatorWeight,
                    remainingBalanceOwner: bidInfo.remainingBalanceOwner,
                    disableOwner: bidInfo.disableOwner            
                });
                emit AuctionVoucherCreated(
                    bidInfo.nodeID, 
                    bidInfo.addr, 
                    $._minValidatorDuration + $._auctionEndTime, 
                    $._auctioningValidatorWeight
                );
            }
            else {
                ++$._occupiedValidatorSlots;
                bytes32 validationID = $._manager.initiateValidatorRegistration(
                    bidInfo.nodeID,
                    bidInfo.blsPublicKey,
                    bidInfo.remainingBalanceOwner,
                    bidInfo.disableOwner,
                    $._auctioningValidatorWeight
                );
                emit InitiatedAuctionValidatorRegistration(
                    validationID,
                    bidInfo.addr,
                    $._minValidatorDuration + $._auctionEndTime,
                    $._auctioningValidatorWeight
                );
                $._validatorsByValidationID[validationID] = ValidatorInfo({
                    addr: bidInfo.addr,
                    endTime: $._minValidatorDuration + $._auctionEndTime,
                    nodeID: bidInfo.nodeID,
                    blsPublicKey: bidInfo.blsPublicKey,
                    validationID: validationID,
                    weight: $._auctioningValidatorWeight
                });
            }
            $._secondPrice = currentBid;
        }

        $._auctionEndTime = 0;
        $._auctioningValidatorSlots = 0;
        $._auctionState = AuctionState.NoAuction;
    }

    function initiateValidatorRegistration (
        bytes memory nodeID
    ) external returns (bytes32){
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        ValidatorVoucher memory voucher = $._vouchers[nodeID];
        delete $._vouchers[nodeID];

        if (block.timestamp > voucher.registrationEndTime) {
            revert VaidatorRegistrationTimePeriodOver(voucher.registrationEndTime, block.timestamp, voucher.nodeID);
        }
        ++$._occupiedValidatorSlots;
        bytes32 validationID = $._manager.initiateValidatorRegistration(
            voucher.nodeID, 
            voucher.blsPublicKey, 
            voucher.remainingBalanceOwner, 
            voucher.disableOwner, 
            voucher.weight
        );

        emit InitiatedAuctionVoucherValidatorRegistration(
            validationID,
            voucher.addr,
            voucher.endTime,
            voucher.weight
        );

        $._validatorsByValidationID[validationID] = ValidatorInfo({
            addr: voucher.addr,
            endTime: voucher.endTime,
            nodeID: voucher.nodeID,
            blsPublicKey: voucher.blsPublicKey,
            validationID: validationID,
            weight: voucher.weight
        });

        return validationID;
    }

    function initiateValidatorRemoval(
        bytes32 validationID
    ) external {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        // Removes validator if current block timestamp is after endTime, also works with initialValidators that
        // Were not registered using the SAVM since their validationID will not be logged making the time 0
        if ($._validatorsByValidationID[validationID].endTime < block.timestamp) {
            $._manager.initiateValidatorRemoval(validationID);
            return;
        }
        revert ValidatorTimeLimitNotPassed($._validatorsByValidationID[validationID].endTime);
    }

    function completeValidatorRegistration(
        uint32 messageIndex
    ) external returns (bytes32) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        return $._manager.completeValidatorRegistration(messageIndex);
    }

    function completeValidatorRemoval(
        uint32 messageIndex
    ) external returns (bytes32) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        bytes32 validationID = $._manager.completeValidatorRemoval(messageIndex);
        // Checks to see if initialValidator, if not, decrement occupiedValidatorSlots
        if ($._validatorsByValidationID[validationID].endTime != 0) {
            --$._occupiedValidatorSlots;
        }
        delete $._validatorsByValidationID[validationID];
        return validationID;
    }

    function setSlotAuctionSettings(
        AuctionSettings calldata auctionSettings
    ) external AuctionOff onlyOwner {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        _validateSettings(auctionSettings);

        $._totalValidatorSlots = auctionSettings.totalValidatorSlots;
        $._minAuctionDuration = auctionSettings.minAuctionDuration;
        $._minValidatorDuration = auctionSettings.minValidatorDuration;
        $._minimumBid = auctionSettings.minimumBid;
        $._auctioningValidatorWeight = auctionSettings.weight;
    }

    function minBidRequired() external view AuctionOn returns (uint256) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();

        if (Heap.length($._bids) < $._totalValidatorSlots) {
            return $._minimumBid;
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
        if ($._minimumBid > bid) {
            revert BidSmallerThanMinimum($._minimumBid, bid);
        }

        // If all slots aren't contended, then fill the heap with any bid
        if (Heap.length($._bids) < $._auctioningValidatorSlots) {
            _lock(bid);
            Heap.insert($._bids, bid);
        } else if (Heap.peek($._bids) < bid) {
            _lock(bid);

            uint256 evictedBid = Heap.replace($._bids, bid);
            emit BidEvicted(evictedBid, $._bidderInfo[evictedBid].nodeID);
            // send back held funds if lost auction
            _unlock($._bidderInfo[evictedBid].addr, evictedBid);

            // deletes info of bidder no longer needed along with replacing it in the heap
            delete $._nodeIsQualified[$._bidderInfo[evictedBid].nodeID];
            delete $._bidderInfo[evictedBid];
            $._secondPrice = evictedBid;
        } else {
            revert InsufficientBidToWinAuction(Heap.peek($._bids) + 1, bid);
        }
        $._bidderInfo[bid] = ValidatorBid({
            addr: msg.sender,
            bid: bid,
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner
        });
        $._nodeIsQualified[nodeID] = true;
        emit SuccessfulBidPlaced(bid, nodeID);
    }

    function _validateSettings(
        AuctionSettings calldata auctionSettings
    ) internal pure {
        if (auctionSettings.weight == 0) {
            revert ZeroWeight();
        }
        if (auctionSettings.minimumBid == 0) {
            revert ZeroMinBid();
        }
    }

    function getTotalValidatorSlots() external view returns (uint16) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._totalValidatorSlots;
    }

    function getAuctioningValidatorWeight() external view returns (uint64) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._auctioningValidatorWeight;
    }

    function getMinAuctionDuration() external view returns (uint256) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._minAuctionDuration;
    }

    function getMinValidatorDuration() external view returns (uint256) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._minValidatorDuration;
    }

    function getMinimumBid() external view returns (uint256) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._minimumBid;
    }

    function getOpenValidatorSlots() external view returns (uint16) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._totalValidatorSlots - $._occupiedValidatorSlots;
    }

    function getAuctionCooldownDuration() external view returns (uint256) {
        SlotAuctionManagerStorage storage $ = _getSlotAuctionManagerStorage();
        return $._auctionCooldownDuration;
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
