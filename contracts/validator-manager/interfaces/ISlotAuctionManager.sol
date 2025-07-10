// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IValidatorManager} from "../interfaces/IValidatorManager.sol";
import {PChainOwner} from "contracts/validator-manager/interfaces/IACP99Manager.sol"; //not sure why it doesnt like relative path

/**
 * @dev Contains needed information to register node as validator.
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
    uint64 validatorWeight;
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
    event SuccessfulBidPlaced(
        uint256 indexed bid,
        bytes indexed nodeID
    );
    
    /**
     * @notice Event emitted when a bid has been removed from the running.
     * @param bid The value of the removed bid.
     * @param nodeID The nodeID that the bid was placed with.
     */
    event BidEvicted(
        uint256 indexed bid, 
        bytes indexed nodeID
    );

    /**
     * @notice Event emitted when a validator is registered.
     * @param validationID The validationID of the new validator.
     * @param ownerAddress The address of the owner of the nodeID.
     * @param validatorEndTime The end time of the validator.
     */
    event InitiatedAuctionValidatorRegistration(
        bytes32 indexed validationID, 
        address indexed ownerAddress, 
        uint256 validatorEndTime
    );

    /**
     * @notice Begins a new auction, only callable by owner of the contract.
     * @param validatorslots The number of slots being auctioned off.
     * @param weight The weight of each validator slot.
     * @param minAuctionDuration The total length of the auction.
     * @param minValidatorDuration The total length of the validation period.
     * @param minimumBid The smallest bid accepted in the auction.
     */
    function initiateAuction(
        uint16 validatorslots,
        uint64 weight,
        uint256 minAuctionDuration,
        uint256 minValidatorDuration,
        uint256 minimumBid
    ) external;

    /**
     * @notice Places a bid in the currently running auction
     * @param bid The amount of tokens to bid.
     * @param nodeID The ID of the node to add to the L1.
     * @param blsPublicKey The BLS public key of the validator.
     * @param remainingBalanceOwner The remaining balance owner of the validator.
     * @param disableOwner The disable owner of the validator.
     */
    function placeBid (
        uint256 bid,
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) external;

    /**
     * @notice Ends the current running auction, adding current winners to the L1
     */
    function endAuction() external;

    /**
     * @notice returns the minimum bid required to enter auction
     */
    function MinBidRequired() external returns (uint256);

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
    ) external returns(bytes32);
}
