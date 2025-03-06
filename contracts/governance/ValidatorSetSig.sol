// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts@5.0.2/utils/ReentrancyGuard.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Message format for the WarpMessage payload to be forwarded to the target contract
 *
 * targetBlockchainID: Blockchain ID of the chain the message is intended for
 * validatorSetSigAddress: Address of the ValidatorSetSig contract this message is intended for
 * targetContractAddress: Address of the contract that the payload should be forwarded to
 * nonce: Unique nonce for the target contract address to provide replay protection
 * value: Value to be sent with the call to the target contract
 * payload: Payload to be forwarded to the target contract. Usually ABI encoded function call with parameters.
 */
struct ValidatorSetSigMessage {
    bytes32 targetBlockchainID;
    address validatorSetSigAddress;
    address targetContractAddress;
    uint256 nonce;
    uint256 value;
    bytes payload;
}

/**
 * @dev Contract that verifies that a set threshold of validators from a given blockchainID
 * have signed an off-chain ICM message and forwards the payload to the target contract specified in the ValidatorSetSigMessage.
 * The threshold itself is set by the validator themselves in their ICM configs:
 * https://github.com/ava-labs/subnet-evm/blob/6c018f89339f3d381909e02013f002f234dc7ae3/precompile/contracts/warp/config.go#L50
 *
 * This is intended to be used for safe off-chain governance of enabled contracts. An example use case would be
 * to deploy an `Ownable` target contract that is owned by an instance of this contract and adding the
 * `onlyOwner` modifier to the functions that should be governed.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract ValidatorSetSig is ReentrancyGuard {
    /**
     * @notice A blockchain ID whose validators need to sign the message for it to be considered valid.
     */
    bytes32 public immutable validatorBlockchainID;

    /**
     * @dev Tracks nonces for messages sent to a specific contract address to provide replay protection.
     * The target contract address is used as the key instead of keeping a global nonce to prevent
     * one party from consuming a global nonce that another was intending to use
     * in case a single instance of ValidatorSetSig contract is used to manage multiple downstream contracts.
     */
    mapping(address targetContractAddress => uint256 nonce) public nonces;

    /**
     * @notice Address that the off-chain ICM message sets as the "source" address.
     * @dev The address is not owned by any EOA or smart contract account, so it
     * cannot possibly be the source address of any other ICM message emitted by the VM.
     */
    address public constant VALIDATORS_SOURCE_ADDRESS = address(0);

    /**
     * @notice The blockchain ID of the chain the contract is deployed on.
     */
    bytes32 public immutable blockchainID;

    /**
     * @notice ICM precompile used for sending and receiving ICM messages.
     */
    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    /**
     * @notice Emited when the payload is successfully delivered to the target contract.
     */
    event Delivered(address indexed targetContractAddress, uint256 indexed nonce);

    constructor(
        bytes32 validatorBlockchainID_
    ) {
        validatorBlockchainID = validatorBlockchainID_;
        blockchainID = WARP_MESSENGER.getBlockchainID();
    }

    receive() external payable {}

    function executeCall(
        uint32 messageIndex
    ) external payable nonReentrant {
        // Get the WarpMessage from the WarpMessenger precompile and verify that it is valid
        (WarpMessage memory message, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(valid, "ValidatorSetSig: invalid warp message");

        // Ensure that the sourceChainID of the validator set that signed the message is the authorized one
        require(
            message.sourceChainID == validatorBlockchainID, "ValidatorSetSig: invalid sourceChainID"
        );
        // Ensure that the originSenderAddress is the zero address to prevent on-chain messages from being forwarded
        require(
            message.originSenderAddress == VALIDATORS_SOURCE_ADDRESS,
            "ValidatorSetSig: non-zero originSenderAddress"
        );

        ValidatorSetSigMessage memory validatorSetSigMessage =
            abi.decode(message.payload, (ValidatorSetSigMessage));

        validateMessage(validatorSetSigMessage);

        nonces[validatorSetSigMessage.targetContractAddress] = validatorSetSigMessage.nonce + 1;

        // We don't need to protect against return bomb vectors below here since the caller is expected to have full control over the contract called.
        (bool success,) =
        // solhint-disable-next-line avoid-low-level-calls
        validatorSetSigMessage.targetContractAddress.call{value: validatorSetSigMessage.value}(
            validatorSetSigMessage.payload
        );

        // Use require to revert the transaction if the call fails. This is to prevent consuming the nonce if the call fails due to out of gas
        // and requiring re-signing of the message with a new nonce.
        require(success, "ValidatorSetSig: call failed");
        emit Delivered(validatorSetSigMessage.targetContractAddress, validatorSetSigMessage.nonce);
    }

    function validateMessage(
        ValidatorSetSigMessage memory message
    ) public view {
        require(
            message.targetBlockchainID == blockchainID,
            "ValidatorSetSig: invalid targetBlockchainID"
        );
        require(
            message.validatorSetSigAddress == address(this),
            "ValidatorSetSig: invalid validatorSetSigAddress"
        );
        require(
            nonces[message.targetContractAddress] == message.nonce, "ValidatorSetSig: invalid nonce"
        );
    }
}
