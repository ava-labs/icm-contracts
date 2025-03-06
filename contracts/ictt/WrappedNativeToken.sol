// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IWrappedNativeToken} from "./interfaces/IWrappedNativeToken.sol";
import {ERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/ERC20.sol";
import {Address} from "@openzeppelin/contracts@5.0.2/utils/Address.sol";

contract WrappedNativeToken is IWrappedNativeToken, ERC20 {
    using Address for address payable;

    constructor(
        string memory symbol
    ) ERC20(string.concat("Wrapped ", symbol), string.concat("W", symbol)) {}

    receive() external payable {
        deposit();
    }

    fallback() external payable {
        deposit();
    }

    function withdraw(
        uint256 amount
    ) external {
        _burn(_msgSender(), amount);
        emit Withdrawal(_msgSender(), amount);
        payable(_msgSender()).sendValue(amount);
    }

    function deposit() public payable {
        _mint(_msgSender(), msg.value);
        emit Deposit(_msgSender(), msg.value);
    }
}
