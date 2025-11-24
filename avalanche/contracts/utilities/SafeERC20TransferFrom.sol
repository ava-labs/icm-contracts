// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

/**
 * @dev Provides a wrapper used for calling an ERC20 transferFrom method
 * to receive tokens to a contract from msg.sender.
 *
 * Checks the balance of the contract using the library before and after the call to safeTransferFrom, and
 * returns balance increase. Designed for safely handling ERC20 "fee on transfer" and "burn on transfer" implementations.
 *
 * Note: A reentrancy guard must always be used when calling token.safeTransferFrom in order to
 * prevent against possible "before-after" pattern vulnerabilities.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
library SafeERC20TransferFrom {
    using SafeERC20 for IERC20;

    /**
     * @dev Checks the balance of the contract before and after the call to safeTransferFrom, and returns the balance
     * increase. Designed for safely handling ERC20 "fee on transfer" and "burn on transfer" implementations.
     */
    // solhint-disable private-vars-leading-underscore
    function safeTransferFrom(IERC20 erc20, uint256 amount) internal returns (uint256) {
        return safeTransferFrom(erc20, msg.sender, amount);
    }
    // solhint-enable private-vars-leading-underscore

    /**
     * @dev Checks the balance of the contract before and after the call to safeTransferFrom, and returns the balance
     * increase. Designed for safely handling ERC20 "fee on transfer" and "burn on transfer" implementations.
     *
     * Supports passing arbitrary sender address values, allowing its use in ERC-2771 compliant meta-transactions.
     * Note: Contracts that use this function must ensure that users cannot pass arbitrary addresses as
     * the {from} address for the {transferFrom} call. Proper authorization (such as msg.sender) must
     * be required to ensure no one can improperly transfer tokens from any address.
     */
    // solhint-disable private-vars-leading-underscore
    function safeTransferFrom(
        IERC20 erc20,
        address from,
        uint256 amount
    ) internal returns (uint256) {
        uint256 balanceBefore = erc20.balanceOf(address(this));
        erc20.safeTransferFrom(from, address(this), amount);
        uint256 balanceAfter = erc20.balanceOf(address(this));

        require(balanceAfter > balanceBefore, "SafeERC20TransferFrom: balance not increased");

        return balanceAfter - balanceBefore;
    }
    // solhint-enable private-vars-leading-underscore
}
