// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {TokenRemote} from "./TokenRemote.sol";
import {TokenRemoteSettings} from "./interfaces/ITokenRemote.sol";
import {INativeTokenRemote} from "./interfaces/INativeTokenRemote.sol";
import {INativeSendAndCallReceiver} from "../interfaces/INativeSendAndCallReceiver.sol";
import {IWrappedNativeToken} from "../interfaces/IWrappedNativeToken.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    TransferrerMessageType,
    TransferrerMessage,
    SingleHopSendMessage,
    SingleHopCallMessage
} from "../interfaces/ITokenTransferrer.sol";
import {TeleporterFeeInfo, TeleporterMessageInput} from "@teleporter/ITeleporterMessenger.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/ERC20.sol";
import {ERC20Upgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/token/ERC20/ERC20Upgradeable.sol";
import {Address} from "@openzeppelin/contracts@5.0.2/utils/Address.sol";
import {CallUtils} from "@utilities/CallUtils.sol";
import {TokenScalingUtils} from "@utilities/TokenScalingUtils.sol";
import {SafeERC20TransferFrom} from "@utilities/SafeERC20TransferFrom.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";

/**
 * @title NativeTokenRemoteUpgradeable
 * @notice This contract is an {INativeTokenRemote} that receives tokens from its specifed {TokenHome} instance,
 * and represents the received tokens as the native token on this chain.
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract NativeTokenRemoteUpgradeable is
    INativeTokenRemote,
    IWrappedNativeToken,
    ERC20Upgradeable,
    TokenRemote
{
    using Address for address payable;

    // solhint-disable private-vars-leading-underscore
    /**
     * @dev Namespace storage slots following the ERC-7201 standard to prevent
     * storage collisions between upgradeable contracts.
     *
     * @custom:storage-location erc7201:avalanche-ictt.storage.NativeTokenRemote
     */
    struct NativeTokenRemoteStorage {
        /**
         * @notice Percentage of burned transaction fees that will be rewarded to a relayer delivering
         * the message created by calling calling reportBurnedTxFees().
         */
        uint256 _burnedFeesReportingRewardPercentage;
        /**
         * @notice Total number of tokens minted by this contract through the native minter precompile.
         */
        uint256 _totalMinted;
        /**
         * @notice The balance of BURNED_TX_FEES_ADDRESS the last time burned fees were reported to the TokenHome instance.
         */
        uint256 _lastestBurnedFeesReported;
    }
    // solhint-enable private-vars-leading-underscore

    /**
     * @dev Storage slot computed based off ERC-7201 formula
     * keccak256(abi.encode(uint256(keccak256("avalanche-ictt.storage.NativeTokenRemote")) - 1)) & ~bytes32(uint256(0xff));
     */
    bytes32 public constant NATIVE_TOKEN_REMOTE_STORAGE_LOCATION =
        0x69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf500;

    // solhint-disable ordering
    function _getNativeTokenRemoteStorage()
        private
        pure
        returns (NativeTokenRemoteStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := NATIVE_TOKEN_REMOTE_STORAGE_LOCATION
        }
    }

    /**
     * @notice The address where the burned transaction fees are credited.
     *
     * @dev Defined as BLACKHOLE_ADDRESS at
     * https://github.com/ava-labs/subnet-evm/blob/v0.6.0/constants/constants.go
     * C-Chain value found at https://github.com/ava-labs/coreth/blob/v0.13.2/constants/constants.go
     * It is a system-reserved address by default in subnet-evm and coreth, and transfers cannot be sent here manually.
     */
    address public constant BURNED_TX_FEES_ADDRESS = 0x0100000000000000000000000000000000000000;

    /**
     * @notice The address where native tokens are sent in order to be burned to transfer to other chains.
     *
     * @dev This address is distinct from {BURNED_TX_FEES_ADDRESS} so that the amount of burned transaction
     * fees and burned transferred amounts can be tracked separately.
     * This address was chosen arbitrarily.
     */
    address public constant BURNED_FOR_TRANSFER_ADDRESS = 0x0100000000000000000000000000000000010203;

    /**
     * @notice Address used to blackhole funds on the home chain, effectively burning them.
     *
     * @dev When reporting burned transaction fee amounts, this address is used as the recipient
     * address for the funds to be sent to be burned on the home chain.
     * This address was chosen arbitrarily.
     */
    address public constant HOME_CHAIN_BURN_ADDRESS = 0x0100000000000000000000000000000000010203;

    /**
     * @notice The native minter precompile.
     */
    INativeMinter public constant NATIVE_MINTER =
        INativeMinter(0x0200000000000000000000000000000000000001);

    /**
     * @dev When modifier is used, the function can only be called after the contract is fully collelateralized,
     * accounting for the initialReserveImbalance.
     */
    modifier onlyWhenCollateralized() {
        require(getIsCollateralized(), "NativeTokenRemote: contract undercollateralized");
        _;
    }

    constructor(
        ICMInitializable init
    ) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    /**
     * @notice Initializes this token TokenRemote instance to receive tokens from the specified TokenHome instance,
     * and represents the received tokens with the native token on this chain.
     * @param settings Constructor settings for this token TokenRemote instance.
     * @param nativeAssetSymbol The symbol of the native asset.
     * @param initialReserveImbalance The initial reserve imbalance that must be collateralized before minting.
     * @param burnedFeesReportingRewardPercentage The percentage of burned transaction fees
     * that will be rewarded to sender of the report.
     */
    function initialize(
        TokenRemoteSettings memory settings,
        string memory nativeAssetSymbol,
        uint256 initialReserveImbalance,
        uint256 burnedFeesReportingRewardPercentage
    ) public initializer {
        __NativeTokenRemote_init(
            settings,
            nativeAssetSymbol,
            initialReserveImbalance,
            burnedFeesReportingRewardPercentage
        );
    }

    // solhint-disable-next-line func-name-mixedcase
    function __NativeTokenRemote_init(
        TokenRemoteSettings memory settings,
        string memory nativeAssetSymbol,
        uint256 initialReserveImbalance,
        uint256 burnedFeesReportingRewardPercentage
    ) internal onlyInitializing {
        require(initialReserveImbalance != 0, "NativeTokenRemote: zero initial reserve imbalance");
        __ERC20_init(nativeAssetSymbol, nativeAssetSymbol);
        __TokenRemote_init(settings, initialReserveImbalance, 18);
        __NativeTokenRemote_init_unchained(burnedFeesReportingRewardPercentage);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __NativeTokenRemote_init_unchained(
        uint256 burnedFeesReportingRewardPercentage
    ) internal onlyInitializing {
        require(burnedFeesReportingRewardPercentage < 100, "NativeTokenRemote: invalid percentage");
        _getNativeTokenRemoteStorage()._burnedFeesReportingRewardPercentage =
            burnedFeesReportingRewardPercentage;
    }
    // solhint-enable ordering

    /**
     * @dev Receives native token with no calldata provided. The tokens are credited to the sender's
     * wrapped native token balance.
     */
    receive() external payable {
        deposit();
    }

    /**
     * @dev Fallback function for receiving native tokens. The tokens are credited to the sender's
     * wrapped native token balance.
     */
    fallback() external payable {
        deposit();
    }

    /**
     * @dev See {INativeTokenTransferrer-send}.
     */
    function send(
        SendTokensInput calldata input
    ) external payable onlyWhenCollateralized {
        _send(input, msg.value);
    }

    /**
     * @dev See {INativeTokenTransferrer-sendAndCall}
     */
    function sendAndCall(
        SendAndCallInput calldata input
    ) external payable onlyWhenCollateralized {
        _sendAndCall(input, msg.value);
    }

    /**
     * @dev See {INativeTokenRemote-reportBurnedTxFees}.
     */
    function reportBurnedTxFees(
        uint256 requiredGasLimit
    ) external sendNonReentrant {
        NativeTokenRemoteStorage storage $ = _getNativeTokenRemoteStorage();
        uint256 burnAddressBalance = BURNED_TX_FEES_ADDRESS.balance;
        require(
            burnAddressBalance > $._lastestBurnedFeesReported,
            "NativeTokenRemote: burn address balance not greater than last report"
        );

        uint256 burnedDifference = burnAddressBalance - $._lastestBurnedFeesReported;
        uint256 reward = (burnedDifference * $._burnedFeesReportingRewardPercentage) / 100;
        uint256 burnedTxFees = burnedDifference - reward;
        $._lastestBurnedFeesReported = burnAddressBalance;

        if (reward > 0) {
            // Re-mint the native tokens to this contract, and then deposit them to be the wrapped
            // native token (ERC20) representation, such that they can be used as a Teleporter
            // message fee.
            _mintNativeCoin(address(this), reward);
            _mint(address(this), reward);
        }

        // Check that the scaled amount on the TokenHome instance will be non-zero.
        require(
            TokenScalingUtils.removeTokenScale(
                getTokenMultiplier(), getMultiplyOnRemote(), burnedTxFees
            ) > 0,
            "NativeTokenRemote: zero scaled amount to report burn"
        );

        // Report the burned amount to the TokenHome instance.
        TransferrerMessage memory message = TransferrerMessage({
            messageType: TransferrerMessageType.SINGLE_HOP_SEND,
            payload: abi.encode(
                SingleHopSendMessage({recipient: HOME_CHAIN_BURN_ADDRESS, amount: burnedTxFees})
            )
        });

        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: getTokenHomeBlockchainID(),
                destinationAddress: getTokenHomeAddress(),
                feeInfo: TeleporterFeeInfo({feeTokenAddress: address(this), amount: reward}),
                requiredGasLimit: requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit ReportBurnedTxFees({teleporterMessageID: messageID, feesBurned: burnedTxFees});
    }

    /**
     * @dev See {IWrappedNativeToken-withdraw}.
     *
     * Note: {IWrappedNativeToken-withdraw} should not be confused with {TokenRemote-_withdraw}.
     * {IWrappedNativeToken-withdraw} is the external method to redeem a wrapped native token (ERC20) balance
     * for the native token itself. {TokenRemote-_withdraw} is the internal method used when
     * processing token transfers.
     */
    function withdraw(
        uint256 amount
    ) external {
        emit Withdrawal(_msgSender(), amount);
        _burn(_msgSender(), amount);
        payable(_msgSender()).sendValue(amount);
    }

    /**
     * @dev See {IWrappedNativeToken-deposit}.
     */
    function deposit() public payable {
        emit Deposit(_msgSender(), msg.value);
        _mint(_msgSender(), msg.value);
    }

    /**
     * @dev See {INativeTokenRemote-totalNativeAssetSupply}.
     *
     * Note: {INativeTokenRemote-totalNativeAssetSupply} should not be confused with {IERC20-totalSupply}
     * {INativeTokenRemote-totalNativeAssetSupply} returns the supply of the native asset of the chain,
     * accounting for the amounts that have been transferred in and out of the chain as well as burnt transaction
     * fees. The {initialReserveBalance} is included in this supply since it is in circulation on this
     * chain even prior to it being backed by collateral on the TokenHome instance.
     * {IERC20-totalSupply} returns the supply of the native asset held by this contract
     * that is represented as an ERC20.
     */
    function totalNativeAssetSupply() public view returns (uint256) {
        NativeTokenRemoteStorage storage $ = _getNativeTokenRemoteStorage();
        uint256 burned = BURNED_TX_FEES_ADDRESS.balance + BURNED_FOR_TRANSFER_ADDRESS.balance;
        uint256 created = $._totalMinted + getInitialReserveImbalance();
        return created - burned;
    }

    function getTotalMinted() external view returns (uint256) {
        return _getNativeTokenRemoteStorage()._totalMinted;
    }

    /**
     * @dev See {TokenRemote-_withdraw}
     */
    function _withdraw(address recipient, uint256 amount) internal virtual override {
        emit TokensWithdrawn(recipient, amount);
        _mintNativeCoin(recipient, amount);
    }

    /**
     * @dev See {TokenRemote-_burn}
     *
     * This is the internal {_burn} method called when transferring tokens to another chain.
     * The tokens to be burnt are already held by this contract. To burn the tokens, send the
     * native token amount to the BURNED_FOR_TRANSFER_ADDRESS.
     */
    function _burn(
        uint256 amount
    ) internal virtual override returns (uint256) {
        payable(BURNED_FOR_TRANSFER_ADDRESS).sendValue(amount);
        return amount;
    }

    /**
     * @dev See {TokenRemote-_handleSendAndCall}
     *
     * Mints the tokens to this contract, and send them to the recipient contract as a
     * part of the call to {INativeSendAndCallReceiver-receiveTokens} on the recipient contract.
     * If the call fails, the amount is sent to the fallback recipient.
     *
     * Note: If the recipient contract does not properly handle the full msg.value sent,
     * the balance can be locked in the recipient contract. Receiving contracts must make
     * sure to properly handle the balance to ensure it does not get locked improperly.
     */
    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal virtual override {
        // Mint the tokens to this contract address.
        _mintNativeCoin(address(this), amount);

        // Encode the call to {INativeSendAndCallReceiver-receiveTokens}
        bytes memory payload = abi.encodeCall(
            INativeSendAndCallReceiver.receiveTokens,
            (
                message.sourceBlockchainID,
                message.originTokenTransferrerAddress,
                message.originSenderAddress,
                message.recipientPayload
            )
        );

        // Call the recipient contract with the given payload, gas amount, and value.
        bool success = CallUtils._callWithExactGasAndValue(
            message.recipientGasLimit, amount, message.recipientContract, payload
        );

        // If the call failed, send the funds to the fallback recipient.
        if (success) {
            emit CallSucceeded(message.recipientContract, amount);
        } else {
            emit CallFailed(message.recipientContract, amount);
            payable(message.fallbackRecipient).sendValue(amount);
        }
    }

    /**
     * @notice See {TokenRemote-_handleFees}
     *
     * If the {feeTokenAddress} is this contract, use the internal ERC20 calls
     * to transfer the tokens directly. Otherwise, use the {SafeERC20TransferFrom} library
     * to transfer the tokens.
     */
    function _handleFees(
        address feeTokenAddress,
        uint256 feeAmount
    ) internal virtual override returns (uint256) {
        if (feeAmount == 0) {
            return 0;
        }
        // If the {feeTokenAddress} is this contract, then just deposit the tokens directly.
        if (feeTokenAddress == address(this)) {
            _spendAllowance(_msgSender(), address(this), feeAmount);
            _transfer(_msgSender(), address(this), feeAmount);
            return feeAmount;
        }
        return
            SafeERC20TransferFrom.safeTransferFrom(IERC20(feeTokenAddress), _msgSender(), feeAmount);
    }

    /**
     * @dev Mints coins to the recipient through the NativeMinter precompile.
     */
    function _mintNativeCoin(address recipient, uint256 amount) private {
        NativeTokenRemoteStorage storage $ = _getNativeTokenRemoteStorage();
        $._totalMinted += amount;
        // Calls NativeMinter precompile through INativeMinter interface.
        NATIVE_MINTER.mintNativeCoin(recipient, amount);
    }
}
