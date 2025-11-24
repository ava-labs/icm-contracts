// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity 0.8.25;

library TokenScalingUtils {
    uint256 public constant MAX_TOKEN_DECIMALS = 18;

    /**
     * @notice Scales the {amount} of home tokens to a TokenRemote instance's token scale.
     * @param tokenMultiplier The token multiplier of the TokenRemote instance.
     * @param multiplyOnRemote Whether the amount of home tokens will be multiplied on the remote, or divided.
     * @param homeTokenAmount The amount of home tokens to scale.
     */
    function applyTokenScale(
        uint256 tokenMultiplier,
        bool multiplyOnRemote,
        uint256 homeTokenAmount
    ) internal pure returns (uint256) {
        return _scaleTokens(tokenMultiplier, multiplyOnRemote, homeTokenAmount, true);
    }

    /**
     * @notice Removes the TokenRemote instance's token scaling, and returns the corresponding
     * amount of home tokens.
     * @param tokenMultiplier The token multiplier of the TokenRemote instance.
     * @param multiplyOnRemote Whether the amount of home tokens will be multiplied on the remote, or divided.
     * @param remoteTokenAmount The amount of remote tokens to remove scaling from.
     */
    function removeTokenScale(
        uint256 tokenMultiplier,
        bool multiplyOnRemote,
        uint256 remoteTokenAmount
    ) internal pure returns (uint256) {
        return _scaleTokens(tokenMultiplier, multiplyOnRemote, remoteTokenAmount, false);
    }

    /**
     * @notice Takes both the home and remote token denominations and uses
     * them to derive the token transferrer scaling multiplier values
     * @param homeTokenDecimals The number of decimals of the home token.
     * @param remoteTokenDecimals The number of decimals of the remote token.
     */
    function deriveTokenMultiplierValues(
        uint8 homeTokenDecimals,
        uint8 remoteTokenDecimals
    ) internal pure returns (uint256, bool) {
        bool multiplyOnRemote = remoteTokenDecimals > homeTokenDecimals;
        uint256 tokenMultiplier = 10
            ** (
                multiplyOnRemote
                    ? uint256(remoteTokenDecimals - homeTokenDecimals)
                    : uint256(homeTokenDecimals - remoteTokenDecimals)
            );
        return (tokenMultiplier, multiplyOnRemote);
    }

    /**
     * @dev Scales {value} based on {tokenMultiplier} and if the amount is applying or
     * removing the TokenRemote instance's token scale.
     * Should be used for all tokens and fees being transferred to/from other L1s.
     * @param tokenMultiplier The token multiplier of the TokenRemote instance.
     * @param multiplyOnRemote Whether the amount of home tokens will be multiplied on the remote, or divided.
     * @param amount The amount of tokens to scale.
     * @param isSendToRemote If true, indicates the amount is being sent to the
     * TokenRemote instance, so applies token scale. If false, indicates the amount is being
     * sent back to the TokenHome instance, so removes token scale.
     */
    function _scaleTokens(
        uint256 tokenMultiplier,
        bool multiplyOnRemote,
        uint256 amount,
        bool isSendToRemote
    ) private pure returns (uint256) {
        // Multiply when multiplyOnRemote and isSendToRemote are
        // both true or both false.
        if (multiplyOnRemote == isSendToRemote) {
            return amount * tokenMultiplier;
        }
        // Otherwise divide.
        return amount / tokenMultiplier;
    }
}
