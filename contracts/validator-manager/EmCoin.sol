// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
//very dumb mock coin im making, mainly for the fun of it but the plan is for it to be used with the other smart contract

struct AllowanceInfo {
    address spender;
    address owner;
    uint256 originalAllowance;
    uint256 allowanceRemaining;
}

contract EmCoin is IERC20 {
    mapping(address => uint256) public balances;
    //kinda scuffed but im not sure how else to do this in an efficient way
    mapping(bytes32 _hashedAddr => AllowanceInfo) public allowances;
    uint256 public wallet = 1e9;
    // uint256 receiptNum = 1;

    function transfer(address to, uint256 value) public returns (bool success) {
        assert(balances[msg.sender] >= value);
        return _transferFunds(msg.sender, to, value);
    }

    function transferFrom(address from, address to, uint256 value) public returns (bool success) {
        success = false;
        bytes32 hashedAddr = _hashAddresses(from, to);

        require(allowances[hashedAddr].allowanceRemaining != 0, "No allowance remaining");
        require(allowances[hashedAddr].spender == to, "Invalid spender");
        require(allowances[hashedAddr].owner == from, "Invalid owner");
        require(
            allowances[hashedAddr].allowanceRemaining >= value,
            "Invalid amount, not enough allowance"
        );
        allowances[hashedAddr].allowanceRemaining -= value;
        return _transferFunds(from, to, value);
    }

    function approve(address spender, uint256 value) public returns (bool success) {
        bytes32 hashedAddr = _hashAddresses(msg.sender, spender);
        if (allowances[hashedAddr].allowanceRemaining == allowances[hashedAddr].originalAllowance) {
            allowances[hashedAddr] = AllowanceInfo(spender, msg.sender, value, value);
            emit Approval(msg.sender, spender, value);
            return true;
        } else {
            //sets the balance back to zero if money was taken out, incase of the event where money was taken
            //prematurely on the chain, saves them from being taken again
            allowances[hashedAddr] = AllowanceInfo(spender, msg.sender, 0, 0);
            return false;
        }
    }

    //obviously would want something in return but for now this works just for the sake of getting tokens
    function buyCoin(
        uint256 val
    ) public {
        balances[msg.sender] += val;
        wallet -= val;
    }

    function allowance(address owner, address spender) public view returns (uint256 remaining) {
        bytes32 hashedAddr = _hashAddresses(owner, spender);
        return allowances[hashedAddr].allowanceRemaining;
    }

    function balanceOf(
        address owner
    ) public view returns (uint256 balance) {
        return balances[owner];
    }

    function name() public pure returns (string memory) {
        return "EmCoin";
    }

    function symbol() public pure returns (string memory) {
        return "EMC";
    }

    function decimals() public pure returns (uint8) {
        return 0;
    }

    function totalSupply() public pure returns (uint256) {
        return 1e9;
    }

    function _transferFunds(address from, address to, uint256 value) internal returns (bool) {
        if (balances[from] < value) return false;
        balances[from] -= value;
        balances[to] += value;
        emit Transfer(from, to, value);
        return true;
    }

    function _hashAddresses(address from, address to) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(from, to));
    }
}
