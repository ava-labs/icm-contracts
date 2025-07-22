// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * @title ByteSlicer
 * @notice Utility library for slicing bytes arrays
 * @dev This library provides helper functions for working with bytes arrays
 */
library ByteSlicer {
    /**
     * @notice Slices a bytes array from a given start index for a given length.
     * @param data The bytes array to slice.
     * @param start The starting index (inclusive).
     * @param length The number of bytes to slice.
     * @return result The sliced bytes array.
     */
    function slice(bytes memory data, uint256 start, uint256 length) internal pure returns (bytes memory result) {
        require(data.length >= start + length, "ByteSlicer: out of bounds");
        result = new bytes(length);
        for (uint256 i = 0; i < length; i++) {
            result[i] = data[start + i];
        }
    }
}
