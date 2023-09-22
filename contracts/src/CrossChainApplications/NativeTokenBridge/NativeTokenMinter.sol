//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@subnet-evm-contracts/AllowList.sol";
import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "./INativeTokenMinter.sol";
import "../../Teleporter/ITeleporterMessenger.sol";
import "../../Teleporter/ITeleporterReceiver.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

// Precompiled Native Minter Contract Address
address constant MINTER_ADDRESS = 0x0200000000000000000000000000000000000001;
// Designated Blackhole Address
address constant BLACKHOLE_ADDRESS = 0x0100000000000000000000000000000000000000;

contract NativeTokenMinter is ITeleporterReceiver, INativeTokenMinter, AllowList, ReentrancyGuard {
  INativeTokenMinter _nativeMinter = INativeTokenMinter(MINTER_ADDRESS);

  address public constant WARP_PRECOMPILE_ADDRESS =
      0x0200000000000000000000000000000000000005;

  uint256 public constant TRANSFER_NATIVE_TOKENS_REQUIRED_GAS = 300_000;
  bytes32 public immutable currentChainID;
  bytes32 public immutable partnerChainID;
  address public immutable partnerContractAddress;

  // Used for sending an receiving Teleporter messages.
  ITeleporterMessenger public immutable teleporterMessenger;

  error InvalidTeleporterMessengerAddress();
  error InvalidSourceChain();
  error InvalidPartnerContractAddress();
  error CannotBridgeTokenWithinSameChain();
  error Unauthorized();
  error InsufficientPayment();

  constructor(address teleporterMessengerAddress, bytes32 partnerChainID_, address partnerContractAddress_) AllowList(MINTER_ADDRESS) {
    if (teleporterMessengerAddress == address(0)) {
        revert InvalidTeleporterMessengerAddress();
    }
    teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);
    currentChainID = WarpMessenger(WARP_PRECOMPILE_ADDRESS)
        .getBlockchainID();

    if (partnerContractAddress_ == address(0)) {
        revert InvalidTeleporterMessengerAddress();
    }
    partnerContractAddress = partnerContractAddress_;

    if (partnerChainID_ == currentChainID) {
        revert CannotBridgeTokenWithinSameChain();
    }
    partnerChainID = partnerChainID_;
  }

  /**
    * @dev See {ITeleporterReceiver-receiveTeleporterMessage}.
    *
    * Receives a Teleporter message and routes to the appropriate internal function call.
    */
  function receiveTeleporterMessage(
    bytes32 nativeChainID,
    address nativeBridgeAddress,
    bytes calldata message
  ) external nonReentrant() {

    // Only allow the Teleporter messenger to deliver messages.
    if (msg.sender != address(teleporterMessenger)) {
      revert Unauthorized();
    }
    // Only allow messages from the partner chain.
    if (nativeChainID != partnerChainID) {
      revert InvalidSourceChain();
    }
    // Only allow the partner contract to send messages.
    if (nativeBridgeAddress != partnerContractAddress) {
      revert InvalidPartnerContractAddress();
    }

    (address recipient, uint256 amount) = abi.decode(message, (address, uint256));

    // Calls NativeMinter precompile through INativeMinter interface.
    _nativeMinter.mintNativeCoin(recipient, amount);
    emit MintNativeTokens(recipient, amount);
  }

    /**
    * @dev See {INativeTokenMinter-bridgeTokens}.
    *
    * Requirements:
    *
    * - `msg.value` must be greater than the fee amount.
    */
    function bridgeTokens(
        address recipient,
        address feeTokenContractAddress,
        uint256 feeAmount
    ) external payable nonReentrant {
        // The recipient cannot be the zero address.
        if (recipient == address(0)) {
            revert InvalidRecipientAddress();
        }

        // Lock tokens in this bridge instance. Supports "fee/burn on transfer" ERC20 token
        // implementations by only bridging the actual balance increase reflected by the call
        // to transferFrom.
        uint256 adjustedAmount = SafeERC20TransferFrom.safeTransferFrom(
            IERC20(feeTokenContractAddress),
            feeAmount
        );

        // Ensure that the adjusted amount is greater than the fee to be paid.
        if (adjustedAmount <= feeAmount) {
            revert InsufficientAdjustedAmount(adjustedAmount, feeAmount);
        }

        // Burn native token by sending to BLACKHOLE_ADDRESS
        payable(BLACKHOLE_ADDRESS).transfer(msg.value);

        // Send Teleporter message.
        bytes memory messageData = abi.encode(recipient, msg.value);

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: partnerChainID,
                destinationAddress: partnerContractAddress,
                feeInfo: TeleporterFeeInfo({
                    contractAddress: feeTokenContractAddress,
                    amount: feeAmount
                }),
                requiredGasLimit: TRANSFER_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );

        emit BridgeTokens({
            tokenContractAddress: feeTokenContractAddress,
            teleporterMessageID: messageID,
            destinationChainID: partnerChainID,
            destinationBridgeAddress: partnerBridgeAddress,
            recipient: recipient,
            transferAmount: msg.value,
            feeAmount: feeAmount
        });
    }
}
