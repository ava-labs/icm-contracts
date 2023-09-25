// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * @dev Abstract contract that helps implement reentrancy guards between functions for sending and receiving.
 *
 * Consecutive calls for sending functions should work together, same for receive functions, but recursive calls
 * should be detected as a reentrancy and revert.
 *
 * Calls between send and receive functions should also be allowed, but not in the case it ends up being a recursive
 * send or receive call. For example the following should fail: send -> receive -> send.
 */
abstract contract ReentrancyGuards {
    // Send and Receive reentrancy guards
    uint256 internal constant _NOT_ENTERED = 1;
    uint256 internal constant _ENTERED = 2;
    uint256 internal _sendEntered;
    uint256 internal _receiveEntered;

    // Errors
    error ReceiverReentrancy();
    error SenderReentrancy();

    // senderNonReentrant modifier makes sure we can not reenter between sender calls.
    // This modifier should be used for messenger sender functions that have external calls and do not want to allow
    // recursive calls with other sender functions.
    modifier senderNonReentrant() {
        if (_sendEntered != _NOT_ENTERED) {
            revert SenderReentrancy();
        }
        _sendEntered = _ENTERED;
        _;
        _sendEntered = _NOT_ENTERED;
    }

    // receiverNonReentrant modifier makes sure we can not reenter between receiver calls.
    // This modifier should be used for messenger receiver functions that have external calls and do not want to allow
    // recursive calls with other receiver functions.
    modifier receiverNonReentrant() {
        if (_receiveEntered != _NOT_ENTERED) {
            revert ReceiverReentrancy();
        }
        _receiveEntered = _ENTERED;
        _;
        _receiveEntered = _NOT_ENTERED;
    }

    constructor() {
        _sendEntered = _NOT_ENTERED;
        _receiveEntered = _NOT_ENTERED;
    }
}
