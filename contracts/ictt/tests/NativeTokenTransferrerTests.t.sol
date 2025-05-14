// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {TokenTransferrerTest} from "./TokenTransferrerTests.t.sol";
import {INativeTokenTransferrer} from "../interfaces/INativeTokenTransferrer.sol";
import {SendTokensInput, SendAndCallInput} from "../interfaces/ITokenTransferrer.sol";

abstract contract NativeTokenTransferrerTest is TokenTransferrerTest {
    INativeTokenTransferrer public nativeTokenTransferrer;

    event Deposit(address indexed sender, uint256 amount);
    event Withdrawal(address indexed sender, uint256 amount);

    function testZeroSendAmount() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        _setUpExpectedZeroAmountRevert();
        _send(input, 0);
    }

    function _send(SendTokensInput memory input, uint256 amount) internal virtual override {
        // solhint-disable-next-line check-send-result
        nativeTokenTransferrer.send{value: amount}(input);
    }

    function _sendAndCall(
        SendAndCallInput memory input,
        uint256 amount
    ) internal virtual override {
        nativeTokenTransferrer.sendAndCall{value: amount}(input);
    }
}
