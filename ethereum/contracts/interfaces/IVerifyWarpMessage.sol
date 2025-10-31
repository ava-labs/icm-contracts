// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ICMMessage} from "../utils/ICM.sol";

/**
 * @notice Allows a contract to specify how a Warp message should be verified.
 */
interface IVerifyICMMessage {
    /*
     * @notice Verify an ICM message
     * @return True if the message is valid, false otherwise
     */
    function verifyICMMessage(
        ICMMessage calldata message
    ) external view returns (bool);
}