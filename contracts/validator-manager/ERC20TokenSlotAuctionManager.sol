// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {PChainOwner} from "./ACP99Manager.sol";
import {IERC20Mintable} from "./interfaces/IERC20Mintable.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {SlotAuctionManager} from "./SlotAuctionManager.sol";
import {SafeERC20TransferFrom} from "@utilities/SafeERC20TransferFrom.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {IERC20TokenSlotAuctionManager} from "./interfaces/IERC20TokenSlotAuctionManager.sol";
import {IValidatorManager} from "./interfaces/IValidatorManager.sol";
import {AuctionState} from "./interfaces/ISlotAuctionManager.sol";


contract ERC20TokenSlotAuctionManager is SlotAuctionManager, IERC20TokenSlotAuctionManager {
    using SafeERC20 for IERC20Mintable;
    using SafeERC20TransferFrom for IERC20Mintable;

    IERC20Mintable _token;

    constructor(
        address vmAddress,
        address erc20Address,
        uint16 validatorslots,
        uint64 weight,
        uint256 minAuctionDuration,
        uint256 minValidatorDuration,
        uint256 minimumBid
    ) {
        _token = IERC20Mintable(erc20Address);
        VALIDATOR_MANAGER = IValidatorManager(vmAddress);
        auctionState = AuctionState.NoAuction;
        _setSlotAuctionSettings(validatorslots, weight, minAuctionDuration, minValidatorDuration, minimumBid);
    }

    function placeBid(
        uint256 bid, 
        bytes memory nodeID, 
        bytes memory blsPublicKey, 
        PChainOwner memory remainingBalanceOwner, 
        PChainOwner memory disableOwner 
    ) external nonReentrant {
        _placeBid(bid, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner);
    }
   
    /**
     * @notice Returns the ERC20 token being bid
     */
    function ERC20() external view returns (IERC20Mintable) {
        return _token;
    }

    /**
     * @notice See {SlotAuctionManager-_lock}
     * Note: Must be guarded with reentrancy guard for safe transfer from.
     */
    function _lock(
        uint256 value
    ) internal virtual override returns (uint256) {
        return _token.safeTransferFrom(value);
    }

    /**
     * @notice See {SlotAuctionManager-_unlock}
     * Note: Must be guarded with reentrancy guard for safe transfer.
     */
    function _unlock(address to, uint256 value) internal virtual override {
        _token.safeTransfer(to, value);
    }
}
