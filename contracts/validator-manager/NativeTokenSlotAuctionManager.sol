// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {PChainOwner} from "./ACP99Manager.sol";
import {SlotAuctionManager} from "./SlotAuctionManager.sol";
import {INativeTokenSlotAuctionManager} from "./interfaces/INativeTokenSlotAuctionManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {Address} from "@openzeppelin/contracts@5.0.2/utils/Address.sol";
import {SlotAuctionManagerSettings} from "./interfaces/ISlotAuctionManager.sol";

/**
 * @dev Implementation of the {INativeTokenSlotAuctionManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract NativeTokenSlotAuctionManager is SlotAuctionManager, INativeTokenSlotAuctionManager {
    using Address for address payable;

    constructor(
        ICMInitializable init
    ) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    // solhint-disable ordering

    /**
     * @notice Initialize the native token slot auction manager
     * @param settings Initial settings for the slot auction validator manager
     */
    function initialize(
        SlotAuctionManagerSettings calldata settings
    ) external initializer {
        __NativeTokenSlotAuctionManager_init(settings);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __NativeTokenSlotAuctionManager_init(
        SlotAuctionManagerSettings calldata settings
    ) internal onlyInitializing {
        __SlotAuctionManager_init(settings);
    }

    // solhint-disable-next-line func-name-mixedcase, no-empty-blocks
    function __NativeTokenStakingManager_init_unchained() internal onlyInitializing {}

    /**
     * @notice See {INativeTokenSlotAuctionManager-placeBid}
     */
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
