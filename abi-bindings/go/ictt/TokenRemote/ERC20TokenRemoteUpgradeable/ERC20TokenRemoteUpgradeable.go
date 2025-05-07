// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20tokenremoteupgradeable

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SendAndCallInput is an auto generated low-level Go binding around an user-defined struct.
type SendAndCallInput struct {
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	RecipientContract                  common.Address
	RecipientPayload                   []byte
	RequiredGasLimit                   *big.Int
	RecipientGasLimit                  *big.Int
	MultiHopFallback                   common.Address
	FallbackRecipient                  common.Address
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
}

// SendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type SendTokensInput struct {
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	Recipient                          common.Address
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
	RequiredGasLimit                   *big.Int
	MultiHopFallback                   common.Address
}

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// TokenRemoteSettings is an auto generated low-level Go binding around an user-defined struct.
type TokenRemoteSettings struct {
	TeleporterRegistryAddress common.Address
	TeleporterManager         common.Address
	MinTeleporterVersion      *big.Int
	TokenHomeBlockchainID     [32]byte
	TokenHomeAddress          common.Address
	TokenHomeDecimals         uint8
}

// ERC20TokenRemoteUpgradeableMetaData contains all meta data concerning the ERC20TokenRemoteUpgradeable contract.
var ERC20TokenRemoteUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumICMInitializable\",\"name\":\"init\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC20_TOKEN_REMOTE_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_SEND_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTER_REMOTE_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TELEPORTER_REGISTRY_APP_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_REMOTE_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"payloadSize\",\"type\":\"uint256\"}],\"name\":\"calculateNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMultiplyOnRemote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenHomeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenHomeBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTeleporterVersion\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tokenHomeBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenHomeAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenHomeDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenRemoteSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"registerWithHome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516148ce3803806148ce83398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b6147818061014d5f395ff3fe608060405234801561000f575f80fd5b506004361061021e575f3560e01c806370a082311161012a578063b8a46d02116100b4578063dd62ed3e11610079578063dd62ed3e146104bd578063e0fd9cb8146104d0578063ef793e2a146104d8578063f2fde38b146104e0578063f3f981d8146104f3575f80fd5b8063b8a46d0214610474578063c3cd692714610487578063c868efaa1461048f578063c9fe4ddf146104a2578063d2cc7a70146104b5575f80fd5b80638da5cb5b116100fa5780638da5cb5b146103ee578063909a6ac01461043257806395d89b4114610446578063973142971461044e578063a9059cbb14610461575f80fd5b806370a08231146103a0578063715018a6146103d457806371717c18146103dc5780637ee3779a146103e6575f80fd5b8063313ce567116101ab5780635507f3d11161017b5780635507f3d1146103365780635d16225d146103405780635eb995141461035357806362431a6514610366578063656900381461038d575f80fd5b8063313ce567146102c057806335cac159146102f45780634213cf781461031b5780634511243e14610323575f80fd5b806315beb59f116101f157806315beb59f1461027d57806318160ddd1461028657806323b872dd1461028e578063254ac160146102a15780632b0d8f18146102ab575f80fd5b806302a30c7d1461022257806306fdde031461023f5780630733c8c814610254578063095ea7b31461026a575b5f80fd5b61022a610506565b60405190151581526020015b60405180910390f35b61024761051d565b60405161023691906136ec565b61025c6105dd565b604051908152602001610236565b61022a610278366004613722565b6105f1565b61025c6105dc81565b61025c61060a565b61022a61029c36600461374c565b610625565b61025c6201fbd081565b6102be6102b936600461378a565b61064a565b005b7f9b9029a3537fcf0e984763da4ac33bbf592a3462819171bf424e91cf626223005460405160ff9091168152602001610236565b61025c7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0081565b61025c61074c565b6102be61033136600461378a565b61075d565b61025c6205302081565b6102be61034e3660046137a5565b61084c565b6102be6103613660046137d4565b61085a565b61025c7f9b9029a3537fcf0e984763da4ac33bbf592a3462819171bf424e91cf6262230081565b6102be61039b3660046137eb565b61086e565b61025c6103ae36600461378a565b6001600160a01b03165f9081525f8051602061470c833981519152602052604090205490565b6102be610878565b61025c6205573081565b61022a61088b565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03165b6040516001600160a01b039091168152602001610236565b61025c5f8051602061472c83398151915281565b6102476108a2565b61022a61045c36600461378a565b6108e0565b61022a61046f366004613722565b6108f9565b6102be610482366004613823565b610906565b61041a610acd565b6102be61049d366004613839565b610aea565b6102be6104b03660046139f3565b610ca7565b61025c610db9565b61025c6104cb366004613adc565b610dce565b61025c610e17565b61025c610e2b565b6102be6104ee36600461378a565b610e3f565b61025c6105013660046137d4565b610e79565b5f80610510610e8f565b6006015460ff1692915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060915f8051602061470c8339815191529161055b90613b13565b80601f016020809104026020016040519081016040528092919081815260200182805461058790613b13565b80156105d25780601f106105a9576101008083540402835291602001916105d2565b820191905f5260205f20905b8154815290600101906020018083116105b557829003601f168201915b505050505091505090565b5f806105e7610e8f565b6003015492915050565b5f336105fe818585610eb3565b60019150505b92915050565b5f805f8051602061470c8339815191525b6002015492915050565b5f33610632858285610ec5565b61063d858585610f22565b60019150505b9392505050565b5f8051602061472c833981519152610660610f7f565b6001600160a01b03821661068f5760405162461bcd60e51b815260040161068690613b45565b60405180910390fd5b6106998183610f87565b156106fc5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610686565b6001600160a01b0382165f81815260018381016020526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a25050565b5f80610756610e8f565b5492915050565b5f8051602061472c833981519152610773610f7f565b6001600160a01b0382166107995760405162461bcd60e51b815260040161068690613b45565b6107a38183610f87565b6108015760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472794170703a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610686565b6001600160a01b0382165f818152600183016020526040808220805460ff19169055517f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c39190a25050565b6108568282610fa8565b5050565b610862610f7f565b61086b81611030565b50565b61085682826111c8565b610880611250565b6108895f6112ab565b565b5f80610895610e8f565b6004015460ff1692915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060915f8051602061470c8339815191529161055b90613b13565b5f5f8051602061472c8339815191526106438184610f87565b5f336105fe818585610f22565b5f61090f610e8f565b6006810154909150610100900460ff161561096c5760405162461bcd60e51b815260206004820152601f60248201527f546f6b656e52656d6f74653a20616c72656164792072656769737465726564006044820152606401610686565b604080516060808201835260058401548252600284015460ff600160a01b820481166020808601918252600160a81b9093048216858701908152865180880188525f808252885188518188015293518516848a01529151909316828601528651808303909501855260809091019095528082019290925291929091610a01906109f79087018761378a565b866020013561131b565b6040805160c0810182526001870154815260028701546001600160a01b031660208083019190915282518084018452939450610ac5939192830191908190610a4b908b018b61378a565b6001600160a01b0316815260209081018690529082526201fbd0908201526040015f5b604051908082528060200260200182016040528015610a97578160200160208202803683370190505b50815260200184604051602001610aae9190613ba7565b604051602081830303815290604052815250611363565b505050505050565b5f80610ad7610e8f565b600201546001600160a01b031692915050565b610af261147e565b5f5f8051602061472c83398151915260028101548154919250906001600160a01b0316634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015610b5d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b819190613be9565b1015610be85760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610686565b610bf28133610f87565b15610c585760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610686565b610c98858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506114c892505050565b50610ca16116df565b50505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015610ceb5750825b90505f826001600160401b03166001148015610d065750303b155b905081158015610d14575080155b15610d325760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610d5c57845460ff60401b1916600160401b1785555b610d6889898989611709565b8315610dae57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f805f8051602061472c83398151915261061b565b6001600160a01b039182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b5f80610e21610e8f565b6001015492915050565b5f80610e35610e8f565b6005015492915050565b610e47611250565b6001600160a01b038116610e7057604051631e4fbdf760e01b81525f6004820152602401610686565b61086b816112ab565b5f6005610e8783601f613c14565b901c92915050565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0090565b610ec08383836001611759565b505050565b5f610ed08484610dce565b90505f198114610ca15781811015610f1457604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610686565b610ca184848484035f611759565b6001600160a01b038316610f4b57604051634b637e8f60e11b81525f6004820152602401610686565b6001600160a01b038216610f745760405163ec442f0560e01b81525f6004820152602401610686565b610ec083838361183c565b610889611250565b6001600160a01b03165f908152600191909101602052604090205460ff1690565b7fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c75008054600114610fea5760405162461bcd60e51b815260040161068690613c27565b600281555f610ff7610e8f565b905061100284611975565b600181015484350361101d576110188484611a60565b611027565b6110278484611be4565b50600190555050565b5f8051602061472c83398151915280546040805163301fd1f560e21b815290515f926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa158015611084573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110a89190613be9565b60028301549091508184111561111a5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610686565b80841161118f5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610686565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d905f90a350505050565b7fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c7500805460011461120a5760405162461bcd60e51b815260040161068690613c27565b600281555f611217610e8f565b905061122284611daf565b600181015484350361123e576112388484611fe9565b50611248565b61123884846121f8565b600190555050565b336112827f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146108895760405163118cdaa760e01b8152336004820152602401610686565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f815f0361132a57505f610604565b306001600160a01b0384160361135857611346335b3084610ec5565b611351333084610f22565b5080610604565b610643833384612494565b5f8061136d6125f7565b60408401516020015190915015611412576040830151516001600160a01b03166113ef5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610686565b604083015160208101519051611412916001600160a01b039091169083906126e7565b604051630624488560e41b81526001600160a01b0382169063624488509061143e908690600401613c6b565b6020604051808303815f875af115801561145a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106439190613be9565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016114c257604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f6114d1610e8f565b9050806001015484146115385760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20696e76616c696420736f7572636520626c6f636044820152681ad8da185a5b88125160ba1b6064820152608401610686565b60028101546001600160a01b038481169116146115aa5760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a20696e76616c6964206f726967696e2073656e646044820152696572206164647265737360b01b6064820152608401610686565b5f828060200190518101906115bf9190613d64565b6006830154909150610100900460ff1615806115e05750600682015460ff16155b156115f75760068201805461ffff19166101011790555b60018151600481111561160c5761160c613b93565b03611643575f816020015180602001905181019061162a9190613dec565b905061163d815f0151826020015161276e565b506116d8565b60028151600481111561165857611658613b93565b03611686575f81602001518060200190518101906116769190613e24565b905061163d8182608001516127bb565b60405162461bcd60e51b815260206004820152602160248201527f546f6b656e52656d6f74653a20696e76616c6964206d657373616765207479706044820152606560f81b6064820152608401610686565b5050505050565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b6001905550565b611711612914565b61171b838361295d565b611726845f8361296f565b7f9b9029a3537fcf0e984763da4ac33bbf592a3462819171bf424e91cf62622300805460ff191660ff8316179055610ca1565b5f8051602061470c8339815191526001600160a01b0385166117905760405163e602df0560e01b81525f6004820152602401610686565b6001600160a01b0384166117b957604051634a1406b160e11b81525f6004820152602401610686565b6001600160a01b038086165f908152600183016020908152604080832093881683529290522083905581156116d857836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161182d91815260200190565b60405180910390a35050505050565b5f8051602061470c8339815191526001600160a01b0384166118765781816002015f82825461186b9190613c14565b909155506118e69050565b6001600160a01b0384165f90815260208290526040902054828110156118c85760405163391434e360e21b81526001600160a01b03861660048201526024810182905260448101849052606401610686565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b038316611904576002810180548390039055611922565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161196791815260200190565b60405180910390a350505050565b5f611986606083016040840161378a565b6001600160a01b0316036119e85760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e74206164647260448201526265737360e81b6064820152608401610686565b5f8160c0013511611a0b5760405162461bcd60e51b815260040161068690613eee565b8035611a295760405162461bcd60e51b815260040161068690613f32565b5f611a3a604083016020840161378a565b6001600160a01b03160361086b5760405162461bcd60e51b815260040161068690613f7d565b5f611a69610e8f565b9050611a99611a7e604085016020860161378a565b60a0850135611a94610100870160e0880161378a565b6129a0565b5f611abd83611aae608087016060880161378a565b86608001358760a00135612a9d565b6040805180820190915291945091505f9080600181526020016040518060400160405280886040016020810190611af4919061378a565b6001600160a01b0316815260200187815250604051602001611b169190613fda565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f92611b969282019080611b6e60808c0160608d0161378a565b6001600160a01b03168152602090810188905290825260c08a0135908201526040015f610a6e565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb528888604051611bd4929190613ffa565b60405180910390a3505050505050565b5f611bed610e8f565b9050611c1a8335611c04604086016020870161378a565b611c15610100870160e0880161378a565b612b60565b5f611c2f83611aae608087016060880161378a565b60408051808201825260038152815160e081018352883581529396509193505f9260208084019282820191611c68918b01908b0161378a565b6001600160a01b03168152602001611c8660608a0160408b0161378a565b6001600160a01b031681526020810188905260a0890135604082015260c08901356060820152608001611cc06101008a0160e08b0161378a565b6001600160a01b03169052604051611d309190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f92611b969282019080611d8860808c0160608d0161378a565b6001600160a01b03168152602090810188905290825262053020908201526040015f610a6e565b8035611dcd5760405162461bcd60e51b815260040161068690613f32565b5f611dde604083016020840161378a565b6001600160a01b031603611e045760405162461bcd60e51b815260040161068690613f7d565b5f611e15606083016040840161378a565b6001600160a01b031603611e805760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420636f6e7460448201526b72616374206164647265737360a01b6064820152608401610686565b5f816080013511611ea35760405162461bcd60e51b815260040161068690613eee565b5f8160a0013511611f045760405162461bcd60e51b815260206004820152602560248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420676173206044820152641b1a5b5a5d60da1b6064820152608401610686565b80608001358160a0013510611f6c5760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a20696e76616c696420726563697069656e742067604482015267185cc81b1a5b5a5d60c21b6064820152608401610686565b5f611f7e610100830160e0840161378a565b6001600160a01b03160361086b5760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f2066616c6c6261636b20726563697060448201526b69656e74206164647265737360a01b6064820152608401610686565b5f611ff2610e8f565b905061201d612007604085016020860161378a565b610140850135611a9460e0870160c0880161378a565b5f612045836120346101208701610100880161378a565b866101200135876101400135612a9d565b6040805180820190915291945091505f908060028152602001604051806101000160405280865f01548152602001306001600160a01b031681526020016120893390565b6001600160a01b031681526020016120a760608a0160408b0161378a565b6001600160a01b03168152602081018890526040016120c960608a018a614099565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060a0890135602082015260400161211d6101008a0160e08b0161378a565b6001600160a01b0316905260405161213891906020016140db565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f926121ba92820190806121926101208c016101008d0161378a565b6001600160a01b03168152602090810188905290825260808a0135908201526040015f610a6e565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168888604051611bd49291906141e2565b5f612201610e8f565b90506122288335612218604086016020870161378a565b611c1560e0870160c0880161378a565b5f61223f836120346101208701610100880161378a565b6040805180820190915291945091505f90806004815260200160405180610160016040528061226b3390565b6001600160a01b03168152602001885f01358152602001886020016020810190612295919061378a565b6001600160a01b031681526020016122b360608a0160408b0161378a565b6001600160a01b03168152602081018890526040016122d560608a018a614099565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060a089013560208201526040016123296101008a0160e08b0161378a565b6001600160a01b031681526080890135602082015260400161235160e08a0160c08b0161378a565b6001600160a01b031681526101408901356020918201526040516123769291016142f0565b60408051601f19818403018152919052905290505f6105dc6123a561239e6060890189614099565b9050610e79565b6123af91906143cd565b6123bc9062055730613c14565b6040805160c0810182526001870154815260028701546001600160a01b03166020820152815180830183529293505f9261244592820190806124066101208d016101008e0161378a565b6001600160a01b031681526020908101899052908252818101869052604080515f815280830182528184015251606090920191610aae91889101613ba7565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b1689896040516124839291906141e2565b60405180910390a350505050505050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa1580156124da573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124fe9190613be9565b90506125156001600160a01b038616853086612bfe565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015612559573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061257d9190613be9565b90508181116125e35760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610686565b6125ed82826143e4565b9695505050505050565b5f8051602061472c83398151915280546040805163d820e64f60e01b815290515f939284926001600160a01b039091169163d820e64f916004808201926020929091908290030181865afa158015612651573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061267591906143f7565b90506126818282610f87565b156106045760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610686565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301525f919085169063dd62ed3e90604401602060405180830381865afa158015612734573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127589190613be9565b9050610ca184846127698585613c14565b612c65565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b826040516127a991815260200190565b60405180910390a26108568282612cf4565b6127c53082612cf4565b6127d430836060015183610eb3565b5f825f01518360200151846040015130858760a001516040516024016127ff96959493929190614412565b60408051601f198184030181529190526020810180516001600160e01b03166394395edd60e01b17905260c084015160608501519192505f91612843919084612d28565b90505f612854308660600151610dce565b90506128653086606001515f610eb3565b81156128b75784606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4856040516128aa91815260200190565b60405180910390a26128ff565b84606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0856040516128f691815260200190565b60405180910390a25b80156116d8576116d8308660e0015183610f22565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661088957604051631afcd79f60e31b815260040160405180910390fd5b612965612914565b6108568282612d3d565b612977612914565b61298d835f015184602001518560400151612d8d565b612995612da8565b610ec0838383612db8565b5f6129a9610e8f565b60028101549091506001600160a01b038581169116146129db5760405162461bcd60e51b815260040161068690614452565b8215612a355760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f207365636f6e646172792060448201526266656560e81b6064820152608401610686565b6001600160a01b03821615610ca15760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f206d756c74692d686f702060448201526766616c6c6261636b60c01b6064820152608401610686565b5f805f612aa8610e8f565b9050612ab3876130dc565b9650612abf868661131b565b60038201546004830154919650612ad99160ff16866130f4565b60038201546004830154612af1919060ff168a6130f4565b11612b535760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a20696e73756666696369656e7420746f6b656e7360448201526b103a37903a3930b739b332b960a11b6064820152608401610686565b5094959294509192505050565b5f612b69610e8f565b80549091508403612b9c57306001600160a01b03841603612b9c5760405162461bcd60e51b815260040161068690614452565b6001600160a01b038216610ca15760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f206d756c74692d686f702066616c6c6044820152636261636b60e01b6064820152608401610686565b6040516001600160a01b038481166024830152838116604483015260648201839052610ca19186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613101565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052612cb68482613162565b610ca1576040516001600160a01b0384811660248301525f6044830152612cea91869182169063095ea7b390606401612c33565b610ca18482613101565b6001600160a01b038216612d1d5760405163ec442f0560e01b81525f6004820152602401610686565b6108565f838361183c565b5f612d35845f8585613203565b949350505050565b612d45612914565b5f8051602061470c8339815191527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace03612d7e84826144f3565b5060048101610ca183826144f3565b612d95612914565b612d9f83826132d3565b610ec0826132f5565b612db0612914565b610889613306565b612dc0612914565b5f612dc9610e8f565b90506005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e0e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e329190613be9565b81556060840151612e985760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d6520626c6f60448201526918dad8da185a5b88125160b21b6064820152608401610686565b8054606085015103612f125760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e52656d6f74653a2063616e6e6f74206465706c6f7920746f20736160448201527f6d6520626c6f636b636861696e20617320746f6b656e20686f6d6500000000006064820152608401610686565b60808401516001600160a01b0316612f785760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d65206164646044820152637265737360e01b6064820152608401610686565b60128460a0015160ff161115612fe25760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20746f6b656e20686f6d6520646563696d616c73604482015268040e8dede40d0d2ced60bb1b6064820152608401610686565b60128260ff1611156130425760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a20746f6b656e20646563696d616c7320746f6f206044820152630d0d2ced60e31b6064820152608401610686565b60608401516001820155608084015160028201805460058401869055600684018054871560ff1990911617905560a08701516001600160a01b039093166001600160a81b031990911617600160a01b60ff808516919091029190911760ff60a81b1916600160a81b918616919091021790556130be908361332d565b60048301805460ff1916911515919091179055600390910155505050565b5f6130e63361133f565b6130f03383613377565b5090565b5f612d358484845f6133ab565b5f6131156001600160a01b038416836133d2565b905080515f1415801561313957508080602001905181019061313791906145ae565b155b15610ec057604051635274afe760e01b81526001600160a01b0384166004820152602401610686565b5f805f846001600160a01b03168460405161317d91906145cd565b5f604051808303815f865af19150503d805f81146131b6576040519150601f19603f3d011682016040523d82523d5f602084013e6131bb565b606091505b50915091508180156131e55750805115806131e55750808060200190518101906131e591906145ae565b80156131fa57505f856001600160a01b03163b115b95945050505050565b5f845a10156132545760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e742067617300000000006044820152606401610686565b834710156132a45760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c75650000006044820152606401610686565b826001600160a01b03163b5f036132bc57505f612d35565b5f805f84516020860188888bf19695505050505050565b6132db612914565b6132e36133df565b6132eb6133ef565b61085682826133f7565b6132fd612914565b61086b8161357b565b5f7fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c7500611702565b5f8060ff8085169084161181816133505761334885876145e8565b60ff1661335e565b61335a86866145e8565b60ff165b61336990600a6146e1565b9350909150505b9250929050565b6001600160a01b0382166133a057604051634b637e8f60e11b81525f6004820152602401610686565b610856825f8361183c565b5f811515841515036133c8576133c185846143cd565b9050612d35565b6131fa85846146ec565b606061064383835f613583565b6133e7612914565b610889613612565b610889612914565b6133ff612914565b6001600160a01b03821661347b5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201527f706f7274657220726567697374727920616464726573730000000000000000006064820152608401610686565b5f5f8051602061472c83398151915290505f8390505f816001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156134cd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134f19190613be9565b116135595760405162461bcd60e51b815260206004820152603260248201527f54656c65706f7274657252656769737472794170703a20696e76616c69642054604482015271656c65706f7274657220726567697374727960701b6064820152608401610686565b81546001600160a01b0319166001600160a01b038216178255610ca183611030565b610e47612914565b6060814710156135a85760405163cd78605960e01b8152306004820152602401610686565b5f80856001600160a01b031684866040516135c391906145cd565b5f6040518083038185875af1925050503d805f81146135fd576040519150601f19603f3d011682016040523d82523d5f602084013e613602565b606091505b50915091506125ed86838361361a565b6116df612914565b60608261362f5761362a82613676565b610643565b815115801561364657506001600160a01b0384163b155b1561366f57604051639996b31560e01b81526001600160a01b0385166004820152602401610686565b5080610643565b8051156136865780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f5b838110156136b95781810151838201526020016136a1565b50505f910152565b5f81518084526136d881602086016020860161369f565b601f01601f19169290920160200192915050565b602081525f61064360208301846136c1565b6001600160a01b038116811461086b575f80fd5b803561371d816136fe565b919050565b5f8060408385031215613733575f80fd5b823561373e816136fe565b946020939093013593505050565b5f805f6060848603121561375e575f80fd5b8335613769816136fe565b92506020840135613779816136fe565b929592945050506040919091013590565b5f6020828403121561379a575f80fd5b8135610643816136fe565b5f808284036101208112156137b8575f80fd5b610100808212156137c7575f80fd5b9395938601359450505050565b5f602082840312156137e4575f80fd5b5035919050565b5f80604083850312156137fc575f80fd5b82356001600160401b03811115613811575f80fd5b8301610160818603121561373e575f80fd5b5f60408284031215613833575f80fd5b50919050565b5f805f806060858703121561384c575f80fd5b84359350602085013561385e816136fe565b925060408501356001600160401b0380821115613879575f80fd5b818701915087601f83011261388c575f80fd5b81358181111561389a575f80fd5b8860208285010111156138ab575f80fd5b95989497505060200194505050565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b03811182821017156138f0576138f06138ba565b60405290565b604080519081016001600160401b03811182821017156138f0576138f06138ba565b60405161010081016001600160401b03811182821017156138f0576138f06138ba565b604051601f8201601f191681016001600160401b0381118282101715613963576139636138ba565b604052919050565b803560ff8116811461371d575f80fd5b5f6001600160401b03821115613993576139936138ba565b50601f01601f191660200190565b5f82601f8301126139b0575f80fd5b81356139c36139be8261397b565b61393b565b8181528460208386010111156139d7575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f80848603610120811215613a08575f80fd5b60c0811215613a15575f80fd5b50613a1e6138ce565b8535613a29816136fe565b81526020860135613a39816136fe565b8060208301525060408601356040820152606086013560608201526080860135613a62816136fe565b6080820152613a7360a0870161396b565b60a0820152935060c08501356001600160401b0380821115613a93575f80fd5b613a9f888389016139a1565b945060e0870135915080821115613ab4575f80fd5b50613ac1878288016139a1565b925050613ad1610100860161396b565b905092959194509250565b5f8060408385031215613aed575f80fd5b8235613af8816136fe565b91506020830135613b08816136fe565b809150509250929050565b600181811c90821680613b2757607f821691505b60208210810361383357634e487b7160e01b5f52602260045260245ffd5b6020808252602e908201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b634e487b7160e01b5f52602160045260245ffd5b602081525f825160058110613bca57634e487b7160e01b5f52602160045260245ffd5b806020840152506020830151604080840152612d3560608401826136c1565b5f60208284031215613bf9575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561060457610604613c00565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501525f929161010085019190606087015160a0870152608087015160e060c08801528051938490528401925f92506101208701905b80841015613cf857845183168252938501936001939093019290850190613cd6565b5060a0880151878203601f190160e08901529450613d1681866136c1565b98975050505050505050565b5f82601f830112613d31575f80fd5b8151613d3f6139be8261397b565b818152846020838601011115613d53575f80fd5b612d3582602083016020870161369f565b5f60208284031215613d74575f80fd5b81516001600160401b0380821115613d8a575f80fd5b9083019060408286031215613d9d575f80fd5b613da56138f6565b825160058110613db3575f80fd5b8152602083015182811115613dc6575f80fd5b613dd287828601613d22565b60208301525095945050505050565b805161371d816136fe565b5f60408284031215613dfc575f80fd5b613e046138f6565b8251613e0f816136fe565b81526020928301519281019290925250919050565b5f60208284031215613e34575f80fd5b81516001600160401b0380821115613e4a575f80fd5b908301906101008286031215613e5e575f80fd5b613e66613918565b82518152613e7660208401613de1565b6020820152613e8760408401613de1565b6040820152613e9860608401613de1565b60608201526080830151608082015260a083015182811115613eb8575f80fd5b613ec487828601613d22565b60a08301525060c083015160c0820152613ee060e08401613de1565b60e082015295945050505050565b60208082526024908201527f546f6b656e52656d6f74653a207a65726f20726571756972656420676173206c6040820152631a5b5a5d60e21b606082015260800190565b6020808252602b908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20626c60408201526a1bd8dad8da185a5b88125160aa1b606082015260800190565b60208082526037908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20746f60408201527f6b656e207472616e736665727265722061646472657373000000000000000000606082015260800190565b81516001600160a01b031681526020808301519082015260408101610604565b8235815261012081016020840135614011816136fe565b6001600160a01b039081166020840152604085013590614030826136fe565b16604083015261404260608501613712565b6001600160a01b0381166060840152506080840135608083015260a084013560a083015260c084013560c083015261407c60e08501613712565b6001600160a01b031660e083015261010090910191909152919050565b5f808335601e198436030181126140ae575f80fd5b8301803591506001600160401b038211156140c7575f80fd5b602001915036819003821315613370575f80fd5b60208152815160208201525f602083015160018060a01b0380821660408501528060408601511660608501525050606083015161412360808401826001600160a01b03169052565b50608083015160a083015260a08301516101008060c085015261414a6101208501836136c1565b915060c085015160e085015260e085015161416f828601826001600160a01b03169052565b5090949350505050565b5f808335601e1984360301811261418e575f80fd5b83016020810192503590506001600160401b038111156141ac575f80fd5b803603821315613370575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60408152823560408201525f6141fa60208501613712565b6001600160a01b0316606083015261421460408501613712565b6001600160a01b0316608083015261422f6060850185614179565b6101608060a08601526142476101a0860183856141ba565b9250608087013560c086015260a087013560e086015261426960c08801613712565b9150610100614282818701846001600160a01b03169052565b61428e60e08901613712565b92506101206142a7818801856001600160a01b03169052565b6142b2828a01613712565b935061014091506142cd828801856001600160a01b03169052565b880135918601919091529095013561018084015260209092019290925292915050565b6020815261430a6020820183516001600160a01b03169052565b602082015160408201525f604083015161432f60608401826001600160a01b03169052565b5060608301516001600160a01b038116608084015250608083015160a083015260a08301516101608060c085015261436b6101808501836136c1565b915060c085015160e085015260e0850151610100614393818701836001600160a01b03169052565b8601516101208681019190915286015190506101406143bc818701836001600160a01b03169052565b959095015193019290925250919050565b808202811582820484141761060457610604613c00565b8181038181111561060457610604613c00565b5f60208284031215614407575f80fd5b8151610643816136fe565b8681526001600160a01b0386811660208301528581166040830152841660608201526080810183905260c060a082018190525f90613d16908301846136c1565b6020808252603a908201527f546f6b656e52656d6f74653a20696e76616c69642064657374696e6174696f6e60408201527f20746f6b656e207472616e736665727265722061646472657373000000000000606082015260800190565b601f821115610ec057805f5260205f20601f840160051c810160208510156144d45750805b601f840160051c820191505b818110156116d8575f81556001016144e0565b81516001600160401b0381111561450c5761450c6138ba565b6145208161451a8454613b13565b846144af565b602080601f831160018114614553575f841561453c5750858301515b5f19600386901b1c1916600185901b178555610ac5565b5f85815260208120601f198616915b8281101561458157888601518255948401946001909101908401614562565b508582101561459e57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f602082840312156145be575f80fd5b81518015158114610643575f80fd5b5f82516145de81846020870161369f565b9190910192915050565b60ff828116828216039081111561060457610604613c00565b600181815b8085111561463b57815f190482111561462157614621613c00565b8085161561462e57918102915b93841c9390800290614606565b509250929050565b5f8261465157506001610604565b8161465d57505f610604565b8160018114614673576002811461467d57614699565b6001915050610604565b60ff84111561468e5761468e613c00565b50506001821b610604565b5060208310610133831016604e8410600b84101617156146bc575081810a610604565b6146c68383614601565b805f19048211156146d9576146d9613c00565b029392505050565b5f6106438383614643565b5f8261470657634e487b7160e01b5f52601260045260245ffd5b50049056fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00de77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00a2646970667358221220f0405fa52783f7b560e1866fa555998c3ab2490fa03d05c735b7f4aeb7dffd9264736f6c63430008190033",
}

// ERC20TokenRemoteUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenRemoteUpgradeableMetaData.ABI instead.
var ERC20TokenRemoteUpgradeableABI = ERC20TokenRemoteUpgradeableMetaData.ABI

// ERC20TokenRemoteUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenRemoteUpgradeableMetaData.Bin instead.
var ERC20TokenRemoteUpgradeableBin = ERC20TokenRemoteUpgradeableMetaData.Bin

// DeployERC20TokenRemoteUpgradeable deploys a new Ethereum contract, binding an instance of ERC20TokenRemoteUpgradeable to it.
func DeployERC20TokenRemoteUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *ERC20TokenRemoteUpgradeable, error) {
	parsed, err := ERC20TokenRemoteUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenRemoteUpgradeableBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenRemoteUpgradeable{ERC20TokenRemoteUpgradeableCaller: ERC20TokenRemoteUpgradeableCaller{contract: contract}, ERC20TokenRemoteUpgradeableTransactor: ERC20TokenRemoteUpgradeableTransactor{contract: contract}, ERC20TokenRemoteUpgradeableFilterer: ERC20TokenRemoteUpgradeableFilterer{contract: contract}}, nil
}

// ERC20TokenRemoteUpgradeable is an auto generated Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeable struct {
	ERC20TokenRemoteUpgradeableCaller     // Read-only binding to the contract
	ERC20TokenRemoteUpgradeableTransactor // Write-only binding to the contract
	ERC20TokenRemoteUpgradeableFilterer   // Log filterer for contract events
}

// ERC20TokenRemoteUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenRemoteUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenRemoteUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenRemoteUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenRemoteUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenRemoteUpgradeableSession struct {
	Contract     *ERC20TokenRemoteUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERC20TokenRemoteUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenRemoteUpgradeableCallerSession struct {
	Contract *ERC20TokenRemoteUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ERC20TokenRemoteUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenRemoteUpgradeableTransactorSession struct {
	Contract     *ERC20TokenRemoteUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ERC20TokenRemoteUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableRaw struct {
	Contract *ERC20TokenRemoteUpgradeable // Generic contract binding to access the raw methods on
}

// ERC20TokenRemoteUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableCallerRaw struct {
	Contract *ERC20TokenRemoteUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenRemoteUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableTransactorRaw struct {
	Contract *ERC20TokenRemoteUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenRemoteUpgradeable creates a new instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeable(address common.Address, backend bind.ContractBackend) (*ERC20TokenRemoteUpgradeable, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeable{ERC20TokenRemoteUpgradeableCaller: ERC20TokenRemoteUpgradeableCaller{contract: contract}, ERC20TokenRemoteUpgradeableTransactor: ERC20TokenRemoteUpgradeableTransactor{contract: contract}, ERC20TokenRemoteUpgradeableFilterer: ERC20TokenRemoteUpgradeableFilterer{contract: contract}}, nil
}

// NewERC20TokenRemoteUpgradeableCaller creates a new read-only instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenRemoteUpgradeableCaller, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableCaller{contract: contract}, nil
}

// NewERC20TokenRemoteUpgradeableTransactor creates a new write-only instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenRemoteUpgradeableTransactor, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTransactor{contract: contract}, nil
}

// NewERC20TokenRemoteUpgradeableFilterer creates a new log filterer instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenRemoteUpgradeableFilterer, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableFilterer{contract: contract}, nil
}

// bindERC20TokenRemoteUpgradeable binds a generic wrapper to an already deployed contract.
func bindERC20TokenRemoteUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenRemoteUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TokenRemoteUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TokenRemoteUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TokenRemoteUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenRemoteUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ERC20TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x62431a65.
//
// Solidity: function ERC20_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) ERC20TOKENREMOTESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "ERC20_TOKEN_REMOTE_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x62431a65.
//
// Solidity: function ERC20_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) ERC20TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// ERC20TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x62431a65.
//
// Solidity: function ERC20_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) ERC20TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) MULTIHOPCALLGASPERWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_CALL_GAS_PER_WORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLGASPERWORD(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLGASPERWORD(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) MULTIHOPCALLREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_CALL_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) MULTIHOPSENDREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_SEND_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPSENDREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPSENDREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) REGISTERREMOTEREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "REGISTER_REMOTE_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.REGISTERREMOTEREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.REGISTERREMOTEREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) TELEPORTERREGISTRYAPPSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "TELEPORTER_REGISTRY_APP_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) TOKENREMOTESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "TOKEN_REMOTE_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Allowance(&_ERC20TokenRemoteUpgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Allowance(&_ERC20TokenRemoteUpgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.BalanceOf(&_ERC20TokenRemoteUpgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.BalanceOf(&_ERC20TokenRemoteUpgradeable.CallOpts, account)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) CalculateNumWords(opts *bind.CallOpts, payloadSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "calculateNumWords", payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.CalculateNumWords(&_ERC20TokenRemoteUpgradeable.CallOpts, payloadSize)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.CalculateNumWords(&_ERC20TokenRemoteUpgradeable.CallOpts, payloadSize)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Decimals() (uint8, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Decimals(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Decimals() (uint8, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Decimals(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetInitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getInitialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetInitialReserveImbalance(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetInitialReserveImbalance(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetIsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getIsCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetIsCollateralized() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetIsCollateralized(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetIsCollateralized() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetIsCollateralized(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetMinTeleporterVersion(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetMinTeleporterVersion(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetMultiplyOnRemote(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getMultiplyOnRemote")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetMultiplyOnRemote() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetMultiplyOnRemote(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetMultiplyOnRemote() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetMultiplyOnRemote(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetTokenHomeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenHomeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetTokenHomeAddress() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeAddress(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetTokenHomeAddress() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeAddress(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetTokenHomeBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenHomeBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetTokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetTokenMultiplier() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenMultiplier(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetTokenMultiplier() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenMultiplier(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.IsTeleporterAddressPaused(&_ERC20TokenRemoteUpgradeable.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.IsTeleporterAddressPaused(&_ERC20TokenRemoteUpgradeable.CallOpts, teleporterAddress)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Name() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Name(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Name() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Name(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Owner() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Owner(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Owner() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Owner(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Symbol() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Symbol(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Symbol() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Symbol(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TotalSupply(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TotalSupply(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Approve(&_ERC20TokenRemoteUpgradeable.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Approve(&_ERC20TokenRemoteUpgradeable.TransactOpts, spender, value)
}

// Initialize is a paid mutator transaction binding the contract method 0xc9fe4ddf.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Initialize(opts *bind.TransactOpts, settings TokenRemoteSettings, tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "initialize", settings, tokenName, tokenSymbol, tokenDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0xc9fe4ddf.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Initialize(settings TokenRemoteSettings, tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Initialize(&_ERC20TokenRemoteUpgradeable.TransactOpts, settings, tokenName, tokenSymbol, tokenDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0xc9fe4ddf.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Initialize(settings TokenRemoteSettings, tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Initialize(&_ERC20TokenRemoteUpgradeable.TransactOpts, settings, tokenName, tokenSymbol, tokenDecimals)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.PauseTeleporterAddress(&_ERC20TokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.PauseTeleporterAddress(&_ERC20TokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ReceiveTeleporterMessage(&_ERC20TokenRemoteUpgradeable.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ReceiveTeleporterMessage(&_ERC20TokenRemoteUpgradeable.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) RegisterWithHome(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "registerWithHome", feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RegisterWithHome(&_ERC20TokenRemoteUpgradeable.TransactOpts, feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RegisterWithHome(&_ERC20TokenRemoteUpgradeable.TransactOpts, feeInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RenounceOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RenounceOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Send(opts *bind.TransactOpts, input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "send", input, amount)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Send(input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Send(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Send(input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Send(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "sendAndCall", input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) SendAndCall(input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.SendAndCall(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) SendAndCall(input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.SendAndCall(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Transfer(&_ERC20TokenRemoteUpgradeable.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Transfer(&_ERC20TokenRemoteUpgradeable.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferFrom(&_ERC20TokenRemoteUpgradeable.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferFrom(&_ERC20TokenRemoteUpgradeable.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.UnpauseTeleporterAddress(&_ERC20TokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.UnpauseTeleporterAddress(&_ERC20TokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.UpdateMinTeleporterVersion(&_ERC20TokenRemoteUpgradeable.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.UpdateMinTeleporterVersion(&_ERC20TokenRemoteUpgradeable.TransactOpts, version)
}

// ERC20TokenRemoteUpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableApprovalIterator struct {
	Event *ERC20TokenRemoteUpgradeableApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableApproval represents a Approval event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20TokenRemoteUpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableApprovalIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableApproval)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseApproval(log types.Log) (*ERC20TokenRemoteUpgradeableApproval, error) {
	event := new(ERC20TokenRemoteUpgradeableApproval)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallFailedIterator struct {
	Event *ERC20TokenRemoteUpgradeableCallFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableCallFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableCallFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableCallFailed represents a CallFailed event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*ERC20TokenRemoteUpgradeableCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableCallFailedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableCallFailed)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallFailed is a log parse operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseCallFailed(log types.Log) (*ERC20TokenRemoteUpgradeableCallFailed, error) {
	event := new(ERC20TokenRemoteUpgradeableCallFailed)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallSucceededIterator struct {
	Event *ERC20TokenRemoteUpgradeableCallSucceeded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableCallSucceeded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableCallSucceeded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableCallSucceeded represents a CallSucceeded event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*ERC20TokenRemoteUpgradeableCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableCallSucceededIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableCallSucceeded)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallSucceeded is a log parse operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseCallSucceeded(log types.Log) (*ERC20TokenRemoteUpgradeableCallSucceeded, error) {
	event := new(ERC20TokenRemoteUpgradeableCallSucceeded)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableInitializedIterator struct {
	Event *ERC20TokenRemoteUpgradeableInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableInitialized represents a Initialized event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20TokenRemoteUpgradeableInitializedIterator, error) {

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableInitializedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableInitialized)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseInitialized(log types.Log) (*ERC20TokenRemoteUpgradeableInitialized, error) {
	event := new(ERC20TokenRemoteUpgradeableInitialized)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator struct {
	Event *ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinTeleporterVersionUpdated is a log parse operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated, error) {
	event := new(ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableOwnershipTransferredIterator struct {
	Event *ERC20TokenRemoteUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20TokenRemoteUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableOwnershipTransferredIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20TokenRemoteUpgradeableOwnershipTransferred, error) {
	event := new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator struct {
	Event *ERC20TokenRemoteUpgradeableTeleporterAddressPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTeleporterAddressPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableTeleporterAddressPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTeleporterAddressPaused)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeleporterAddressPaused is a log parse operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTeleporterAddressPaused(log types.Log) (*ERC20TokenRemoteUpgradeableTeleporterAddressPaused, error) {
	event := new(ERC20TokenRemoteUpgradeableTeleporterAddressPaused)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator struct {
	Event *ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeleporterAddressUnpaused is a log parse operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused, error) {
	event := new(ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensAndCallSentIterator struct {
	Event *ERC20TokenRemoteUpgradeableTokensAndCallSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTokensAndCallSent represents a TokensAndCallSent event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*ERC20TokenRemoteUpgradeableTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTokensAndCallSentIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensAndCallSent is a log parse operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTokensAndCallSent(log types.Log) (*ERC20TokenRemoteUpgradeableTokensAndCallSent, error) {
	event := new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensSentIterator struct {
	Event *ERC20TokenRemoteUpgradeableTokensSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTokensSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableTokensSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTokensSent represents a TokensSent event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*ERC20TokenRemoteUpgradeableTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTokensSentIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTokensSent)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensSent is a log parse operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTokensSent(log types.Log) (*ERC20TokenRemoteUpgradeableTokensSent, error) {
	event := new(ERC20TokenRemoteUpgradeableTokensSent)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensWithdrawnIterator struct {
	Event *ERC20TokenRemoteUpgradeableTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTokensWithdrawn represents a TokensWithdrawn event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ERC20TokenRemoteUpgradeableTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTokensWithdrawnIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTokensWithdrawn(log types.Log) (*ERC20TokenRemoteUpgradeableTokensWithdrawn, error) {
	event := new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTransferIterator struct {
	Event *ERC20TokenRemoteUpgradeableTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTransfer represents a Transfer event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TokenRemoteUpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTransferIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTransfer)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTransfer(log types.Log) (*ERC20TokenRemoteUpgradeableTransfer, error) {
	event := new(ERC20TokenRemoteUpgradeableTransfer)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
