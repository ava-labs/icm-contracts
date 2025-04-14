// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";

/**
 * @dev Abstract contract that helps implement reentrancy guards for Avalanche interchain token transfer {_send} and {_sendAndCall}
 * functions.
 *
 * The send methods must not allow reentry given that can make calls to external contracts such as {safeTransferFrom}
 * and {safeDeposit}. However, the send methods should be allowed to be called from {receiveTeleporterMessage}, either
 * as a part of processing a multi-hop transfer, or as a part of an external call made to process a "sendAndCall"
 * message.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
abstract contract SendReentrancyGuardUpgradeable is Initializable {
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-ictt.storage.SendReentrancyGuard
    struct SendReentrancyGuardStorage {
        uint256 _sendEntered;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-ictt.storage.SendReentrancyGuard")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant _SEND_REENTRANCY_GUARD_STORAGE_LOCATION =
        0xd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c7500;

    // solhint-disable ordering
    function _getSendReentrancyGuardStorage()
        private
        pure
        returns (SendReentrancyGuardStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := _SEND_REENTRANCY_GUARD_STORAGE_LOCATION
        }
    }

    uint256 private constant _NOT_ENTERED = 1;
    uint256 private constant _ENTERED = 2;

    // sendNonReentrant modifier makes sure there is not reentry between {_send} or {_sendAndCall} calls.
    modifier sendNonReentrant() {
        SendReentrancyGuardStorage storage $ = _getSendReentrancyGuardStorage();
        require($._sendEntered == _NOT_ENTERED, "SendReentrancyGuard: send reentrancy");
        $._sendEntered = _ENTERED;
        _;
        $._sendEntered = _NOT_ENTERED;
    }

    //solhint-disable-next-line func-name-mixedcase
    function __SendReentrancyGuard_init() internal onlyInitializing {
        __SendReentrnacyGuard_init_unchained();
    }

    //solhint-disable-next-line func-name-mixedcase
    function __SendReentrnacyGuard_init_unchained() internal {
        SendReentrancyGuardStorage storage $ = _getSendReentrancyGuardStorage();
        $._sendEntered = _NOT_ENTERED;
    }
    // solhint-enable ordering
}
