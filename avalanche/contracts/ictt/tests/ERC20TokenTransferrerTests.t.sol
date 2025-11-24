// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {TokenTransferrerTest} from "./TokenTransferrerTests.t.sol";
import {IERC20TokenTransferrer} from "../interfaces/IERC20TokenTransferrer.sol";
import {SendTokensInput, SendAndCallInput} from "../interfaces/ITokenTransferrer.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {IERC20Errors} from "@openzeppelin/contracts@5.0.2/interfaces/draft-IERC6093.sol";

abstract contract ERC20TokenTransferrerTest is TokenTransferrerTest {
    using SafeERC20 for IERC20;

    IERC20TokenTransferrer public erc20TokenTransferrer;

    function testZeroSendAmount() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        _setUpExpectedZeroAmountRevert();
        _send(input, 0);
    }

    function testSendNoAllowance() public {
        uint256 amount = 2e15;

        SendTokensInput memory input = _createDefaultSendTokensInput();
        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        vm.expectRevert(
            abi.encodeWithSelector(
                IERC20Errors.ERC20InsufficientAllowance.selector, erc20TokenTransferrer, 0, amount
            )
        );
        _send(input, amount);
    }

    function testSendAndCallNoAllowance() public {
        uint256 amount = 2e15;

        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        vm.expectRevert(
            abi.encodeWithSelector(
                IERC20Errors.ERC20InsufficientAllowance.selector, erc20TokenTransferrer, 0, amount
            )
        );
        _sendAndCall(input, amount);
    }

    function _send(SendTokensInput memory input, uint256 amount) internal virtual override {
        // solhint-disable-next-line check-send-result
        erc20TokenTransferrer.send(input, amount);
    }

    function _sendAndCall(
        SendAndCallInput memory input,
        uint256 amount
    ) internal virtual override {
        erc20TokenTransferrer.sendAndCall(input, amount);
    }
}
