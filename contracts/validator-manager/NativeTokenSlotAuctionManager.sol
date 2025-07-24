// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {PChainOwner} from "./ACP99Manager.sol";
import {SlotAuctionManager} from "./SlotAuctionManager.sol";
import {INativeTokenSlotAuctionManager} from "./interfaces/INativeTokenSlotAuctionManager.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/INativeMinter.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {Address} from "@openzeppelin/contracts@5.0.2/utils/Address.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {AuctionState} from "./interfaces/ISlotAuctionManager.sol";
import {IValidatorManager} from "./interfaces/IValidatorManager.sol";

contract NativeTokenSlotAuctionManager is SlotAuctionManager, INativeTokenSlotAuctionManager {
    using Address for address payable;

    constructor(
        address vmAddress,
        uint16 validatorslots,
        uint64 weight,
        uint256 minAuctionDuration,
        uint256 minValidatorDuration,
        uint256 minimumBid
    ) {
        VALIDATOR_MANAGER = IValidatorManager(vmAddress);
        auctionState = AuctionState.NoAuction;
        _setSlotAuctionSettings(
            validatorslots, weight, minAuctionDuration, minValidatorDuration, minimumBid
        );
    }

    function placeBid(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
    ) external payable nonReentrant {
        _placeBid(msg.value, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner);
    }

    /**
     * @notice See {SlotAuctionManager-_lock}
     */
    function _lock(
        uint256 value
    ) internal virtual override returns (uint256) {
        return value;
    }

    /**
     * @notice See {SlotAuctionManager-_unlock}
     */
    function _unlock(address to, uint256 value) internal virtual override {
        payable(to).sendValue(value);
    }
}
