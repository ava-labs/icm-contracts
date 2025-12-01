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


    /**
     * @notice Add the source bytes to the target bytes at the indicated position
     * @param target the array be extended
     * @param source the array be copied to target
     * @param cursor the position in target to start copying source to. Typically the end
     * @return the new ending position of target
     */
    function extendFromSlice(
        bytes memory target,
        bytes memory source,
        uint256 cursor
    ) public pure returns (bytes memory, uint256){
        uint256 length = source.length;
        for (uint256 i = 0; i < length; i++) {
            target[cursor + i] = source[i];
        }
        return (target, cursor + length);
    }
}
