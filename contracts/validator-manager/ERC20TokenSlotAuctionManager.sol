// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {PChainOwner} from "./ACP99Manager.sol";
import {IERC20Mintable} from "./interfaces/IERC20Mintable.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {SlotAuctionManager} from "./SlotAuctionManager.sol";
import {SafeERC20TransferFrom} from "@utilities/SafeERC20TransferFrom.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {IERC20TokenSlotAuctionManager} from "./interfaces/IERC20TokenSlotAuctionManager.sol";
import {SlotAuctionManagerSettings} from "./interfaces/ISlotAuctionManager.sol";

/**
 * @dev Implementation of the {IERC20TokenSlotAuctionManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract ERC20TokenSlotAuctionManager is SlotAuctionManager, IERC20TokenSlotAuctionManager {
    using SafeERC20 for IERC20Mintable;
    using SafeERC20TransferFrom for IERC20Mintable;

    struct ERC20TokenSlotAuctionManagerStorage {
        // solhint-disable-next-line private-vars-leading-underscore
        IERC20Mintable _token;
    }

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.ERC20TokenSlotAuctionManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant ERC20_TOKEN_SLOT_AUCTION_MANAGER_STORAGE_LOCATION =
        0x04264d6e045c48d92b64fb3ce155b1f7a2673239fb0c9b60c505be1c17a7e700;

    // solhint-disable ordering
    function _getERC20SlotAuctionManagerStorage()
        private
        pure
        returns (ERC20TokenSlotAuctionManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := ERC20_TOKEN_SLOT_AUCTION_MANAGER_STORAGE_LOCATION
        }
    }

    constructor(
        ICMInitializable init
    ) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    /**
     * @notice Initialize the ERC20 token slot auction manager
     * @param token The ERC20 token to be bid
     * @param settings Initial settings for the slot auction validator manager
     */
    function initialize(
        IERC20Mintable token,
        SlotAuctionManagerSettings calldata settings
    ) external initializer {
        __ERC20TokenSlotAuctionManager_init(token, settings);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ERC20TokenSlotAuctionManager_init(
        IERC20Mintable token,
        SlotAuctionManagerSettings calldata settings
    ) internal onlyInitializing {
        __SlotAuctionManager_init(settings);
        __ERC20TokenSlotAuctionManager_init_unchained(token);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ERC20TokenSlotAuctionManager_init_unchained(
        IERC20Mintable token
    ) internal onlyInitializing {
        ERC20TokenSlotAuctionManagerStorage storage $ = _getERC20SlotAuctionManagerStorage();
        if (address(token) == address(0)) {
            revert ZeroAddress();
        }
        $._token = token;
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
    function erc20() external view returns (IERC20Mintable) {
        ERC20TokenSlotAuctionManagerStorage storage $ = _getERC20SlotAuctionManagerStorage();
        return $._token;
    }

    /**
     * @notice See {SlotAuctionManager-_lock}
     * Note: Must be guarded with reentrancy guard for safe transfer from.
     */
    function _lock(
        uint256 value
    ) internal virtual override returns (uint256) {
        ERC20TokenSlotAuctionManagerStorage storage $ = _getERC20SlotAuctionManagerStorage();
        return $._token.safeTransferFrom(value);
    }

    /**
     * @notice See {SlotAuctionManager-_unlock}
     * Note: Must be guarded with reentrancy guard for safe transfer.
     */
    function _unlock(address to, uint256 value) internal virtual override {
        ERC20TokenSlotAuctionManagerStorage storage $ = _getERC20SlotAuctionManagerStorage();
        $._token.safeIncreaseAllowance(to, value);
    }
}
