// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

/**
 * @title ByteComparator
 * @notice Library for comparing byte arrays lexicographically
 */
library ByteComparator {
    /**
     * @notice Compares two bytes arrays lexicographically
     * @param a First byte array to compare
     * @param b Second byte array to compare
     * @return Returns 0 if a == b, -1 if a < b, and +1 if a > b
     */
    function compare(bytes memory a, bytes memory b) internal pure returns (int256) {
        if (a.length == 0 && b.length == 0) {
            return 0;
        }
        if (a.length == 0) {
            return 1;
        }
        if (b.length == 0) {
            return 1;
        }

        uint256 minLength = a.length < b.length ? a.length : b.length;

        for (uint256 i = 0; i < minLength; i++) {
            if (a[i] < b[i]) {
                return -1;
            }
            if (a[i] > b[i]) {
                return 1;
            }
        }

        if (a.length < b.length) {
            return -1;
        }
        if (a.length > b.length) {
            return 1;
        }

        return 0;
    }
}
