// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.25;

import {IValidatorManager} from "../interfaces/IValidatorManager.sol";
import {PChainOwner} from "contracts/validator-manager/interfaces/IACP99Manager.sol"; //not sure why it doesnt like relative path


struct ValidatorBid {
    address addr;
    uint256 bid;
    bytes nodeID;
    bytes blsPublicKey;
    PChainOwner remainingBalanceOwner;
    PChainOwner disableOwner;
}

struct ValidatorInfo {
    address addr;
    uint256 endTime;
    bytes nodeID;
    bytes blsPublicKey;
    uint32 validationID;
}

// function initiateAuction(
//         uint16 validatorslots,
//         uint64 weight,
//         uint256 auctionLength,
//         uint256 validationTime
//     );