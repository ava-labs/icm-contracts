// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {WrappedNativeToken} from "../WrappedNativeToken.sol";

contract WrappedNativeTokenTest is Test {
    address public constant TEST_ACCOUNT = 0xd4E96eF8eee8678dBFf4d535E033Ed1a4F7605b7;
    WrappedNativeToken public wavax;

    function setUp() public virtual {
        wavax = new WrappedNativeToken("AVAX");
    }

    function testFallback() public {
        (bool success,) = address(wavax).call{value: 1}("1234567812345678");
        assertTrue(success);
        assertEq(wavax.balanceOf(address(this)), 1);
    }

    function testDepositWithdraw() public {
        uint256 depositAmount = 500;
        uint256 withdrawAmount = 100;
        vm.deal(TEST_ACCOUNT, depositAmount);
        vm.startPrank(TEST_ACCOUNT);
        wavax.deposit{value: depositAmount}();
        assertEq(wavax.balanceOf(TEST_ACCOUNT), depositAmount);
        wavax.withdraw(withdrawAmount);
        assertEq(wavax.balanceOf(TEST_ACCOUNT), depositAmount - withdrawAmount);
        assertEq(TEST_ACCOUNT.balance, withdrawAmount);
    }
}
