// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Math} from "@openzeppelin/contracts@5.0.2/utils/math/Math.sol";
import {TeleporterMessageReceipt} from "./ITeleporterMessenger.sol";

/**
 * @dev ReceiptQueue is a convenience library that creates a queue-like interface of
 * TeleporterMessageReceipt structs. It provides FIFO properties.
 * Note: All functions in this library are internal so that the library is not deployed as a contract.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
library ReceiptQueue {
    // A receipt queue represents a FIFO list of receipts.
    struct TeleporterMessageReceiptQueue {
        // Tracks the head of the queue within the {data} mapping.
        uint256 first;
        // Tracks the tail of the queue within the {data} mapping.
        uint256 last;
        // Represents the elements in the queue. Each newly enqueue item is
        // assigned an index one greater than the previous tail.
        mapping(uint256 index => TeleporterMessageReceipt receipt) data;
    }

    // The maximum number of receipts to include in a single message.
    uint256 private constant _MAXIMUM_RECEIPT_COUNT = 5;

    // solhint-disable private-vars-leading-underscore
    /**
     * @dev Adds a receipt to the queue.
     */
    function enqueue(
        TeleporterMessageReceiptQueue storage queue,
        TeleporterMessageReceipt memory receipt
    ) internal {
        queue.data[queue.last++] = receipt;
    }

    /**
     * @dev Removes the oldest outstanding receipt from the queue.
     *
     * Requirements:
     * - The queue must be non-empty.
     */
    function dequeue(
        TeleporterMessageReceiptQueue storage queue
    ) internal returns (TeleporterMessageReceipt memory) {
        uint256 first_ = queue.first;
        require(queue.last != first_, "ReceiptQueue: empty queue");
        TeleporterMessageReceipt memory receipt = queue.data[first_];
        delete queue.data[first_];
        queue.first = first_ + 1;
        return receipt;
    }

    /**
     * @dev Returns the outstanding receipts for the given chain ID that should be included in the next message sent.
     */
    function getOutstandingReceiptsToSend(
        TeleporterMessageReceiptQueue storage queue
    ) internal returns (TeleporterMessageReceipt[] memory) {
        // Calculate the result size as the minimum of the number of receipts and maximum batch size.
        uint256 resultSize = Math.min(_MAXIMUM_RECEIPT_COUNT, size(queue));
        if (resultSize == 0) {
            return new TeleporterMessageReceipt[](0);
        }

        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](resultSize);
        for (uint256 i; i < resultSize; ++i) {
            receipts[i] = dequeue(queue);
        }
        return receipts;
    }

    /**
     * @dev Returns the number of outstanding receipts in the queue.
     */
    function size(
        TeleporterMessageReceiptQueue storage queue
    ) internal view returns (uint256) {
        return queue.last - queue.first;
    }

    /**
     * @dev Returns the receipt at the given index in the queue.
     */
    function getReceiptAtIndex(
        TeleporterMessageReceiptQueue storage queue,
        uint256 index
    ) internal view returns (TeleporterMessageReceipt memory) {
        require(index < size(queue), "ReceiptQueue: index out of bounds");
        return queue.data[queue.first + index];
    }
    // solhint-enable private-vars-leading-underscore
}
