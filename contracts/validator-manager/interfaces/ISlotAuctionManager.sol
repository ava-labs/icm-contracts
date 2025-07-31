// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IValidatorManager} from "../interfaces/IValidatorManager.sol";
import {PChainOwner} from "contracts/validator-manager/interfaces/IACP99Manager.sol"; //not sure why it doesnt like relative path

/**
 * @dev Contains needed information to register node as validator.
 */

/**
 * @dev AuctionFinalizing is necessary (potentially) in the case where
 * a new auction is initiated while the auction is finalizing.
 */
enum AuctionState {
    NoAuction,
    AuctionInProgress,
    AuctionFinalizing
}

/**
 * @dev Contains info about bid in the running to win auction.
 */
struct ValidatorBid {
    address addr;
    uint256 bid;
    bytes nodeID;
    bytes blsPublicKey;
    PChainOwner remainingBalanceOwner;
    PChainOwner disableOwner;
}

/**
 * @dev Contains info about current active validator.
 */
struct ValidatorInfo {
    address addr;
    uint256 endTime;
    bytes nodeID;
    bytes blsPublicKey;
    bytes32 validationID;
    uint64 weight;
}

/**
 * @dev Initialization struct for slot auction manager
 */
struct SlotAuctionManagerSettings {
    address admin;
    IValidatorManager manager;
    AuctionSettings auctionSettings;
}
/**
 * @dev Struct for auction settings, is a part of SlotAuctionManagerSettings but due
 * to the nature of being able to change auction settings a separate struct is provided
 */
struct AuctionSettings {
    uint16 totalValidatorSlots;
    uint64 weight;
    uint256 minValidatorDuration;
    uint256 minAuctionDuration;
    uint256 minimumBid;
    uint256 auctionCooldownDuration;
}

/**
 * Proof of Purchase Validator Manager that takes ERC20 tokens.
 */
interface ISlotAuctionManager {
    /**
     * @notice Event emitted when a new auction has started.
     * @param validatorSlots The number of validators being auctioned off.
     * @param validatorWeight The weight of each validator.
     * @param minValidatorDuration The length of validation time.
     * @param auctionEndTime The final second of the auction.
     * @param minimumBid The minimum bid of the auction.
     */
    event NewValidatorAuction(
        uint16 validatorSlots,
        uint64 validatorWeight,
        uint256 minValidatorDuration,
        uint256 auctionEndTime,
        uint256 minimumBid
    );

    /**
     * @notice Event emitted when a bid capable of winning is placed.
     * @param bid The value of the successful bid.
     * @param nodeID The nodeID that the bid was placed with.
     */
    event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID);

    /**
     * @notice Event emitted when a bid has been removed from the running.
     * @param bid The value of the removed bid.
     * @param nodeID The nodeID that the bid was placed with.
     */
    event BidEvicted(uint256 indexed bid, bytes indexed nodeID);

    /**
     * @notice Event emitted when a validator is registered.
     * @param validationID The validationID of the new validator.
     * @param ownerAddress The address of the owner of the nodeID.
     * @param validatorEndTime The end time of the validator.
     */
    event InitiatedAuctionValidatorRegistration(
        bytes32 indexed validationID,
        address indexed ownerAddress,
        uint256 validatorEndTime,
        uint64 weight
    );

    /**
     * @notice Begins a new auction, callable by anyone, will automatically auction the maximum
     * amount of validator slots with default auction settings
     */
    function initiateAuction() external;

    /**
     * @notice Ends the current running auction, adding current winners to the L1
     */
    function endAuction() external;

    /**
     * @notice returns the minimum bid required to enter auction
     */
    function minBidRequired() external returns (uint256);

    function initiateValidatorRemoval(
        bytes32 validationID
    ) external;

    function initiateRemoveInitialValidator(
        bytes32 validationID
    ) external;

    function completeRemoveInitialValidator(
        uint32 messageIndex
    ) external;

    function completeValidatorRegistration(
        uint32 messageIndex
    ) external returns (bytes32);

    function completeValidatorRemoval(
        uint32 messageIndex
    ) external returns (bytes32);

    function getTotalValidatorSlots() external view returns (uint16);

    function getAuctioningValidatorWeight() external view returns (uint64);

    function getMinAuctionDuration() external view returns (uint256);

    function getMinValidatorDuration() external view returns (uint256);

    function getMinimumBid() external view returns (uint256);

    function getOpenValidatorSlots() external view returns (uint16);

    function getAuctionCooldownDuration() external view returns (uint256);
    // function getValidatorInfoByNodeID(
    //     bytes memory nodeID
    // ) external returns (ValidatorInfo memory);
}
