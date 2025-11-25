// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleportermessenger

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ICMMessage is an auto generated low-level Go binding around an user-defined struct.
type ICMMessage struct {
	UnsignedMessage      ICMUnsignedMessage
	UnsignedMessageBytes []byte
	Signature            ICMSignature
}

// ICMSignature is an auto generated low-level Go binding around an user-defined struct.
type ICMSignature struct {
	Signers   []byte
	Signature []byte
}

// ICMUnsignedMessage is an auto generated low-level Go binding around an user-defined struct.
type ICMUnsignedMessage struct {
	AvalancheNetworkID          uint32
	AvalancheSourceBlockchainID [32]byte
	Payload                     []byte
}

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// TeleporterMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessage struct {
	MessageNonce            *big.Int
	OriginSenderAddress     common.Address
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}

// TeleporterMessageInput is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageInput struct {
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	FeeInfo                 TeleporterFeeInfo
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Message                 []byte
}

// TeleporterMessageReceipt is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageReceipt struct {
	ReceivedMessageNonce *big.Int
	RelayerRewardAddress common.Address
}

// TeleporterMessengerMetaData contains all meta data concerning the TeleporterMessenger contract.
var TeleporterMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"updatedFeeInfo\",\"type\":\"tuple\"}],\"name\":\"AddFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"BlockchainIDInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"MessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"ReceiptReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRedeemer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ReceiveCrossChainMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RelayerRewardsRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"SendCrossChainMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalFeeAmount\",\"type\":\"uint256\"}],\"name\":\"addFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"calculateMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"checkRelayerRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getNextMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getReceiptAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getReceiptQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getRelayerRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"warpContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"messageReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"receiptQueues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"avalancheNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheSourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structICMUnsignedMessage\",\"name\":\"unsignedMessage\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"unsignedMessageBytes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signers\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structICMSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structICMMessage\",\"name\":\"icmMessage\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveInterChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"receivedFailedMessageHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"redeemRelayerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retryMessageExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retrySendCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageInput\",\"name\":\"messageInput\",\"type\":\"tuple\"}],\"name\":\"sendCrossChainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"messageIDs\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"sendSpecifiedReceipts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"sentMessageInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b5060015f81905580556135eb806100245f395ff3fe608060405234801561000f575f80fd5b506004361061016d575f3560e01c8063a8898181116100d9578063df20e8bc11610093578063ebc3b1ba1161006e578063ebc3b1ba14610423578063ecc7042814610446578063fc2d61971461044f578063ffa1ad7414610462575f80fd5b8063df20e8bc1461036e578063e69d606a14610381578063e6e67bd5146103e8575f80fd5b8063a8898181146102e1578063a9a85614146102f4578063c473eef814610307578063c4d66de81461033f578063ccb5f80914610352578063d127dc9b14610365575f80fd5b8063624488501161012a5780636244885014610256578063807c0be0146102695780638245a1b01461027c578063860a3b061461028f578063892bf412146102ae5780638ac0fd04146102ce575f80fd5b80630af5b4ff1461017157806322296c3a1461018c5780632bc8b0bf146101a15780632ca40f55146101b45780632e27c2231461020c578063399b77da14610237575b5f80fd5b610179610490565b6040519081526020015b60405180910390f35b61019f61019a366004612414565b61056b565b005b6101796101af36600461242f565b61065e565b6101fe6101c236600461242f565b600660209081525f9182526040918290208054835180850190945260018201546001600160a01b03168452600290910154918301919091529082565b604051610183929190612446565b61021f61021a36600461242f565b61067a565b6040516001600160a01b039091168152602001610183565b61017961024536600461242f565b5f9081526006602052604090205490565b61017961026436600461246d565b610701565b61019f6102773660046124a3565b61075a565b61019f61028a36600461250c565b610806565b61017961029d36600461242f565b60076020525f908152604090205481565b6102c16102bc36600461253d565b6109b0565b604051610183919061255d565b61019f6102dc36600461257d565b6109e1565b6101796102ef3660046125b2565b610c18565b610179610302366004612622565b610c69565b6101796103153660046126b3565b6001600160a01b039182165f908152600a6020908152604080832093909416825291909152205490565b61019f61034d366004612414565b610efc565b61019f6103603660046126e2565b61101a565b61017960035481565b61017961037c36600461242f565b6111a4565b6103c961038f36600461242f565b5f90815260066020908152604091829020825180840190935260018101546001600160a01b03168084526002909101549290910182905291565b604080516001600160a01b039093168352602083019190915201610183565b61040e6103f636600461242f565b60056020525f90815260409020805460019091015482565b60408051928352602083019190915201610183565b61043661043136600461242f565b6111eb565b6040519015158152602001610183565b61017960045481565b61019f61045d3660046126fc565b611200565b610483604051806040016040528060028152602001612b1960f11b81525081565b604051610183919061278c565b6003545f90806105665760025f9054906101000a90046001600160a01b03166001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104ea573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061050e919061279e565b9050806105365760405162461bcd60e51b815260040161052d906127b5565b60405180910390fd5b600381905560405181907f1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104905f90a25b919050565b335f908152600a602090815260408083206001600160a01b0385168452909152902054806105ec5760405162461bcd60e51b815260206004820152602860248201527f54656c65706f727465724d657373656e6765723a206e6f2072657761726420746044820152676f2072656465656d60c01b606482015260840161052d565b335f818152600a602090815260408083206001600160a01b03871680855290835281842093909355518481529192917f3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88910160405180910390a361065a6001600160a01b0383163383611455565b5050565b5f818152600560205260408120610674906114b9565b92915050565b5f818152600860205260408120546106e65760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f74604482015268081c9958d95a5d995960ba1b606482015260840161052d565b505f908152600960205260409020546001600160a01b031690565b5f60015f54146107235760405162461bcd60e51b815260040161052d906127fc565b60025f5561075061073383612a33565b83355f90815260056020526040902061074b906114cb565b6115c5565b60015f5592915050565b600180541461077b5760405162461bcd60e51b815260040161052d90612ad2565b6002600181905554604051631636e63960e31b81525f916001600160a01b03169063b1b731c8906107b0908690600401612bdd565b5f60405180830381865afa1580156107ca573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526107f19190810190612d66565b90506107fd818361170b565b50506001805550565b60015f54146108275760405162461bcd60e51b815260040161052d906127fc565b60025f9081556003546108409060408401358435610c18565b5f818152600660209081526040918290208251808401845281548152835180850190945260018201546001600160a01b0316845260029091015483830152908101919091528051919250906108a75760405162461bcd60e51b815260040161052d90612d97565b5f836040516020016108b99190612fad565b60408051601f19818403018152919052825181516020830120919250146108f25760405162461bcd60e51b815260040161052d90612fbf565b8360400135837f2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df886856020015160405161092d929190613008565b60405180910390a360025460405163ee5b48eb60e01b81526001600160a01b039091169063ee5b48eb9061096590849060040161278c565b6020604051808303815f875af1158015610981573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109a5919061279e565b505060015f55505050565b604080518082019091525f80825260208201525f8381526005602052604090206109da90836119af565b9392505050565b60015f5414610a025760405162461bcd60e51b815260040161052d906127fc565b60025f556001805414610a275760405162461bcd60e51b815260040161052d90612ad2565b600260015580610a915760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f727465724d657373656e6765723a207a65726f2061646469746960448201526e1bdb985b0819995948185b5bdd5b9d608a1b606482015260840161052d565b6001600160a01b038216610ab75760405162461bcd60e51b815260040161052d9061303c565b5f83815260066020526040902054610ae15760405162461bcd60e51b815260040161052d90612d97565b5f838152600660205260409020600101546001600160a01b03838116911614610b725760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465724d657373656e6765723a20696e76616c69642066656560448201527f20617373657420636f6e74726163742061646472657373000000000000000000606482015260840161052d565b5f610b7d8383611a70565b5f85815260066020526040812060020180549293508392909190610ba29084906130a4565b90915550505f8481526006602052604090819020905185917fc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e191610c03916001019081546001600160a01b0316815260019190910154602082015260400190565b60405180910390a2505060018080555f555050565b5f604051806040016040528060028152602001612b1960f11b815250848484604051602001610c4a94939291906130b7565b6040516020818303038152906040528051906020012090509392505050565b5f60015f5414610c8b5760405162461bcd60e51b815260040161052d906127fc565b60025f90815560035490866001600160401b03811115610cad57610cad61283f565b604051908082528060200260200182016040528015610cf157816020015b604080518082019091525f8082526020820152815260200190600190039081610ccb5790505b509050865f5b81811015610e69575f8a8a83818110610d1257610d126130e5565b9050602002013590505f60085f8381526020019081526020015f20549050805f03610d8e5760405162461bcd60e51b815260206004820152602660248201527f54656c65706f727465724d657373656e6765723a2072656365697074206e6f7460448201526508199bdd5b9960d21b606482015260840161052d565b610d998d8783610c18565b8214610e0d5760405162461bcd60e51b815260206004820152603a60248201527f54656c65706f727465724d657373656e6765723a206d6573736167652049442060448201527f6e6f742066726f6d20736f7572636520626c6f636b636861696e000000000000606482015260840161052d565b5f828152600960209081526040918290205482518084019093528383526001600160a01b03169082018190528651909190879086908110610e5057610e506130e5565b6020026020010181905250505050806001019050610cf7565b506040805160c0810182528b81525f6020820152610eea918101610e92368b90038b018b6130f9565b81526020015f81526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f92018290525093855250506040805192835260208084019091529092015250836115c5565b60015f559a9950505050505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015610f405750825b90505f826001600160401b03166001148015610f5b5750303b155b905081158015610f69575080155b15610f875760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610fb157845460ff60401b1916600160401b1785555b600280546001600160a01b0319166001600160a01b038816179055831561101257845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b600180541461103b5760405162461bcd60e51b815260040161052d90612ad2565b60026001819055546040516306f8253560e41b815263ffffffff841660048201525f9182916001600160a01b0390911690636f825350906024015f60405180830381865afa15801561108f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526110b69190810190613122565b91509150806111195760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a20696e76616c69642077617260448201526870206d65737361676560b81b606482015260840161052d565b60208201516001600160a01b031630146111905760405162461bcd60e51b815260206004820152603260248201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206f726960448201527167696e2073656e646572206164647265737360701b606482015260840161052d565b61119a828461170b565b5050600180555050565b6003545f90806111c65760405162461bcd60e51b815260040161052d906127b5565b5f60045460016111d691906130a4565b90506111e3828583610c18565b949350505050565b5f818152600860205260408120541515610674565b60018054146112215760405162461bcd60e51b815260040161052d90612ad2565b60026001556003545f906112389084908435610c18565b5f81815260076020526040902054909150806112665760405162461bcd60e51b815260040161052d90612d97565b80836040516020016112789190612fad565b60405160208183030381529060405280519060200120146112ab5760405162461bcd60e51b815260040161052d90612fbf565b5f6112bc6080850160608601612414565b6001600160a01b03163b116113305760405162461bcd60e51b815260206004820152603460248201527f54656c65706f727465724d657373656e6765723a2064657374696e6174696f6e604482015273206164647265737320686173206e6f20636f646560601b606482015260840161052d565b604051849083907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35f828152600760209081526040808320839055869161137f918701908701612414565b61138c60e087018761316c565b60405160240161139f94939291906131ae565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b17905290505f6113e66113df6080870160608801612414565b5a84611a7c565b9050806114495760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f727465724d657373656e6765723a20726574727920657865637560448201526a1d1a5bdb8819985a5b195960aa1b606482015260840161052d565b50506001805550505050565b6040516001600160a01b038381166024830152604482018390526114b491859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050611a93565b505050565b805460018201545f91610674916131d8565b60605f6114e160056114dc856114b9565b611af4565b9050805f0361152d57604080515f8082526020820190925290611525565b604080518082019091525f80825260208201528152602001906001900390816114ff5790505b509392505050565b5f816001600160401b038111156115465761154661283f565b60405190808252806020026020018201604052801561158a57816020015b604080518082019091525f80825260208201528152602001906001900390816115645790505b5090505f5b82811015611525576115a085611b09565b8282815181106115b2576115b26130e5565b602090810291909101015260010161158f565b5f805f806115d38686611bd3565b9250925092505f826040516020016115eb9190613303565b60408051808303601f190181528282018252805160208083019190912084528084018681525f89815260068352849020945185555180516001860180546001600160a01b0319166001600160a01b039092169190911790550151600290930192909255885190519192509085907f2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8906116879087908790613315565b60405180910390a360025460405163ee5b48eb60e01b81526001600160a01b039091169063ee5b48eb906116bf90849060040161278c565b6020604051808303815f875af11580156116db573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116ff919061279e565b50929695505050505050565b5f8260400151806020019051810190611724919061340f565b90505f61172f610490565b90508082604001511461179e5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206465736044820152701d1a5b985d1a5bdb8818da185a5b881251607a1b606482015260840161052d565b835182515f916117af918490610c18565b5f81815260086020526040902054909150156118235760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465724d657373656e6765723a206d65737361676520616c7260448201526c1958591e481c9958d95a5d9959609a1b606482015260840161052d565b611831338460a00151611d37565b61188f5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a20756e617574686f72697a6560448201526832103932b630bcb2b960b91b606482015260840161052d565b61189c81845f0151611da3565b6001600160a01b038416156118d2575f81815260096020526040902080546001600160a01b0319166001600160a01b0386161790555b60c0830151515f5b818110156119155761190d84885f01518760c001518481518110611900576119006130e5565b6020026020010151611e13565b6001016118da565b50604080518082018252855181526001600160a01b03871660208083019190915288515f90815260059091529190912061194e91611f37565b336001600160a01b0316865f0151837f292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34888860405161198e9291906134fd565b60405180910390a460e084015151156110125761101282875f015186611f91565b604080518082019091525f80825260208201526119cb836114b9565b8210611a235760405162461bcd60e51b815260206004820152602160248201527f5265636569707451756575653a20696e646578206f7574206f6620626f756e646044820152607360f81b606482015260840161052d565b826002015f83855f0154611a3791906130a4565b815260208082019290925260409081015f20815180830190925280548252600101546001600160a01b0316918101919091529392505050565b5f6109da8333846120c0565b5f805f808451602086015f8989f195945050505050565b5f611aa76001600160a01b03841683612223565b905080515f14158015611acb575080806020019051810190611ac99190613520565b155b156114b457604051635274afe760e01b81526001600160a01b038416600482015260240161052d565b5f818310611b0257816109da565b5090919050565b604080518082019091525f808252602082015281546001830154819003611b725760405162461bcd60e51b815260206004820152601960248201527f5265636569707451756575653a20656d70747920717565756500000000000000604482015260640161052d565b5f8181526002840160208181526040808420815180830190925280548252600180820180546001600160a01b03811685870152888852959094529490556001600160a01b0319909216905590611bc99083906130a4565b9093555090919050565b60408051610100810182525f8082526020820181905291810182905260608082018390526080820183905260a0820181905260c0820181905260e0820152604080518082019091525f80825260208201525f611c2d610490565b90505f60045f8154611c3e90613539565b9190508190559050611c5482885f015183610c18565b9450604051806101000160405280828152602001336001600160a01b03168152602001885f0151815260200188602001516001600160a01b0316815260200188606001518152602001886080015181526020018781526020018860a0015181525093505f808860400151602001511115611d0d576040880151516001600160a01b0316611cf35760405162461bcd60e51b815260040161052d9061303c565b60408801518051602090910151611d0a9190611a70565b90505b604080518082018252980151516001600160a01b0316885260208801525092959194935090915050565b5f81515f03611d4857506001610674565b81515f5b81811015611d9957846001600160a01b0316848281518110611d7057611d706130e5565b60200260200101516001600160a01b031603611d9157600192505050610674565b600101611d4c565b505f949350505050565b805f03611e025760405162461bcd60e51b815260206004820152602760248201527f54656c65706f727465724d657373656e6765723a207a65726f206d657373616760448201526665206e6f6e636560c81b606482015260840161052d565b5f9182526008602052604090912055565b5f611e228484845f0151610c18565b5f818152600660209081526040918290208251808401845281548152835180850190945260018201546001600160a01b031684526002909101548383015290810191909152805191925090611e78575050505050565b5f8281526006602090815260408083208381556001810180546001600160a01b03191690556002018390558382018051830151878401516001600160a01b039081168652600a855283862092515116855292528220805491929091611ede9084906130a4565b9250508190555082602001516001600160a01b031684837fd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf578460200151604051611f289190613551565b60405180910390a45050505050565b600182018054829160028501915f9182611f5083613539565b9091555081526020808201929092526040015f2082518155910151600190910180546001600160a01b0319166001600160a01b039092169190911790555050565b80608001515a1015611ff35760405162461bcd60e51b815260206004820152602560248201527f54656c65706f727465724d657373656e6765723a20696e73756666696369656e604482015264742067617360d81b606482015260840161052d565b80606001516001600160a01b03163b5f03612013576114b4838383612230565b602081015160e08201516040515f92612030928692602401613571565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b179052606083015160808401519192505f91612074919084611a7c565b90508061208d57612086858585612230565b5050505050565b604051849086907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35050505050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015612106573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061212a919061279e565b90506121416001600160a01b0386168530866122a4565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015612185573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121a9919061279e565b905081811161220f5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b606482015260840161052d565b61221982826131d8565b9695505050505050565b60606109da83835f6122e3565b806040516020016122419190613303565b60408051601f1981840301815282825280516020918201205f878152600790925291902055829084907f4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c98590612297908590613303565b60405180910390a3505050565b6040516001600160a01b0384811660248301528381166044830152606482018390526122dd9186918216906323b872dd90608401611482565b50505050565b6060814710156123085760405163cd78605960e01b815230600482015260240161052d565b5f80856001600160a01b03168486604051612323919061359a565b5f6040518083038185875af1925050503d805f811461235d576040519150601f19603f3d011682016040523d82523d5f602084013e612362565b606091505b50915091506122198683836060826123825761237d826123c9565b6109da565b815115801561239957506001600160a01b0384163b155b156123c257604051639996b31560e01b81526001600160a01b038516600482015260240161052d565b50806109da565b8051156123d95780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b50565b6001600160a01b03811681146123f2575f80fd5b8035610566816123f5565b5f60208284031215612424575f80fd5b81356109da816123f5565b5f6020828403121561243f575f80fd5b5035919050565b828152606081016109da602083018480516001600160a01b03168252602090810151910152565b5f6020828403121561247d575f80fd5b81356001600160401b03811115612492575f80fd5b820160e081850312156109da575f80fd5b5f80604083850312156124b4575f80fd5b82356001600160401b038111156124c9575f80fd5b8301606081860312156124da575f80fd5b915060208301356124ea816123f5565b809150509250929050565b5f6101008284031215612506575f80fd5b50919050565b5f6020828403121561251c575f80fd5b81356001600160401b03811115612531575f80fd5b6111e3848285016124f5565b5f806040838503121561254e575f80fd5b50508035926020909101359150565b815181526020808301516001600160a01b03169082015260408101610674565b5f805f6060848603121561258f575f80fd5b8335925060208401356125a1816123f5565b929592945050506040919091013590565b5f805f606084860312156125c4575f80fd5b505081359360208301359350604090920135919050565b5f8083601f8401126125eb575f80fd5b5081356001600160401b03811115612601575f80fd5b6020830191508360208260051b850101111561261b575f80fd5b9250929050565b5f805f805f8086880360a0811215612638575f80fd5b8735965060208801356001600160401b0380821115612655575f80fd5b6126618b838c016125db565b90985096508691506040603f198401121561267a575f80fd5b60408a01955060808a0135925080831115612693575f80fd5b50506126a189828a016125db565b979a9699509497509295939492505050565b5f80604083850312156126c4575f80fd5b82356124da816123f5565b803563ffffffff81168114610566575f80fd5b5f80604083850312156126f3575f80fd5b6124da836126cf565b5f806040838503121561270d575f80fd5b8235915060208301356001600160401b03811115612729575f80fd5b612735858286016124f5565b9150509250929050565b5f5b83811015612759578181015183820152602001612741565b50505f910152565b5f815180845261277881602086016020860161273f565b601f01601f19169290920160200192915050565b602081525f6109da6020830184612761565b5f602082840312156127ae575f80fd5b5051919050565b60208082526027908201527f54656c65706f727465724d657373656e6765723a207a65726f20626c6f636b636040820152661a185a5b88125160ca1b606082015260800190565b60208082526023908201527f5265656e7472616e63794775617264733a2073656e646572207265656e7472616040820152626e637960e81b606082015260800190565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156128755761287561283f565b60405290565b60405160c081016001600160401b03811182821017156128755761287561283f565b60405161010081016001600160401b03811182821017156128755761287561283f565b604051601f8201601f191681016001600160401b03811182821017156128e8576128e861283f565b604052919050565b5f60408284031215612900575f80fd5b612908612853565b90508135612915816123f5565b808252506020820135602082015292915050565b5f6001600160401b038211156129415761294161283f565b5060051b60200190565b5f82601f83011261295a575f80fd5b8135602061296f61296a83612929565b6128c0565b8083825260208201915060208460051b870101935086841115612990575f80fd5b602086015b848110156129b55780356129a8816123f5565b8352918301918301612995565b509695505050505050565b5f6001600160401b038211156129d8576129d861283f565b50601f01601f191660200190565b5f82601f8301126129f5575f80fd5b8135612a0361296a826129c0565b818152846020838601011115612a17575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60e08236031215612a43575f80fd5b612a4b61287b565b82358152612a5b60208401612409565b6020820152612a6d36604085016128f0565b60408201526080830135606082015260a08301356001600160401b0380821115612a95575f80fd5b612aa13683870161294b565b608084015260c0850135915080821115612ab9575f80fd5b50612ac6368286016129e6565b60a08301525092915050565b60208082526025908201527f5265656e7472616e63794775617264733a207265636569766572207265656e7460408201526472616e637960d81b606082015260800190565b5f808335601e19843603018112612b2c575f80fd5b83016020810192503590506001600160401b03811115612b4a575f80fd5b80360382131561261b575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8235603e19833603018112612b94575f80fd5b90910192915050565b5f612ba88283612b17565b60408552612bba604086018284612b58565b915050612bca6020840184612b17565b8583036020870152612219838284612b58565b602081525f8235605e19843603018112612bf5575f80fd5b60606020840152830163ffffffff612c0c826126cf565b166080840152602081013560a0840152612c296040820182612b17565b9150606060c0850152612c4060e085018383612b58565b915050612c506020850185612b17565b601f1980868503016040870152612c68848385612b58565b9350612c776040880188612b80565b9250808685030160608701525050612c8f8282612b9d565b95945050505050565b8051610566816123f5565b5f82601f830112612cb2575f80fd5b8151612cc061296a826129c0565b818152846020838601011115612cd4575f80fd5b6111e382602083016020870161273f565b5f60608284031215612cf5575f80fd5b604051606081016001600160401b038282108183111715612d1857612d1861283f565b816040528293508451835260208501519150612d33826123f5565b8160208401526040850151915080821115612d4c575f80fd5b50612d5985828601612ca3565b6040830152505092915050565b5f60208284031215612d76575f80fd5b81516001600160401b03811115612d8b575f80fd5b6111e384828501612ce5565b60208082526026908201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f7460408201526508199bdd5b9960d21b606082015260800190565b5f808335601e19843603018112612df2575f80fd5b83016020810192503590506001600160401b03811115612e10575f80fd5b8060051b360382131561261b575f80fd5b8183525f60208085019450825f5b85811015612e5d578135612e42816123f5565b6001600160a01b031687529582019590820190600101612e2f565b509495945050505050565b5f808335601e19843603018112612e7d575f80fd5b83016020810192503590506001600160401b03811115612e9b575f80fd5b8060061b360382131561261b575f80fd5b8183525f60208085019450825f5b85811015612e5d578135875282820135612ed3816123f5565b6001600160a01b0316878401526040968701969190910190600101612eba565b5f610100823584526020830135612f09816123f5565b6001600160a01b0316602085015260408381013590850152612f2d60608401612409565b6001600160a01b0316606085015260808381013590850152612f5260a0840184612ddd565b8260a0870152612f658387018284612e21565b92505050612f7660c0840184612e68565b85830360c0870152612f89838284612eac565b92505050612f9a60e0840184612b17565b85830360e0870152612219838284612b58565b602081525f6109da6020830184612ef3565b60208082526029908201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206d65736040820152680e6c2ceca40d0c2e6d60bb1b606082015260800190565b606081525f61301a6060830185612ef3565b90506109da602083018480516001600160a01b03168252602090810151910152565b60208082526034908201527f54656c65706f727465724d657373656e6765723a207a65726f2066656520617360408201527373657420636f6e7472616374206164647265737360601b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561067457610674613090565b608081525f6130c96080830187612761565b6020830195909552506040810192909252606090910152919050565b634e487b7160e01b5f52603260045260245ffd5b5f60408284031215613109575f80fd5b6109da83836128f0565b80518015158114610566575f80fd5b5f8060408385031215613133575f80fd5b82516001600160401b03811115613148575f80fd5b61315485828601612ce5565b92505061316360208401613113565b90509250929050565b5f808335601e19843603018112613181575f80fd5b8301803591506001600160401b0382111561319a575f80fd5b60200191503681900382131561261b575f80fd5b8481526001600160a01b03841660208201526060604082018190525f906122199083018486612b58565b8181038181111561067457610674613090565b5f815180845260208085019450602084015f5b83811015612e5d5781516001600160a01b0316875295820195908201906001016131fe565b5f815180845260208085019450602084015f5b83811015612e5d5761325c878351805182526020908101516001600160a01b0316910152565b6040969096019590820190600101613236565b5f6101008251845260018060a01b0360208401511660208501526040830151604085015260608301516132ad60608601826001600160a01b03169052565b506080830151608085015260a08301518160a08601526132cf828601826131eb565b91505060c083015184820360c08601526132e98282613223565b91505060e083015184820360e0860152612c8f8282612761565b602081525f6109da602083018461326f565b606081525f61301a606083018561326f565b5f82601f830112613336575f80fd5b8151602061334661296a83612929565b8083825260208201915060208460051b870101935086841115613367575f80fd5b602086015b848110156129b557805161337f816123f5565b835291830191830161336c565b5f82601f83011261339b575f80fd5b815160206133ab61296a83612929565b82815260069290921b840181019181810190868411156133c9575f80fd5b8286015b848110156129b557604081890312156133e4575f80fd5b6133ec612853565b81518152848201516133fd816123f5565b818601528352918301916040016133cd565b5f6020828403121561341f575f80fd5b81516001600160401b0380821115613435575f80fd5b908301906101008286031215613449575f80fd5b61345161289d565b8251815261346160208401612c98565b60208201526040830151604082015261347c60608401612c98565b60608201526080830151608082015260a08301518281111561349c575f80fd5b6134a887828601613327565b60a08301525060c0830151828111156134bf575f80fd5b6134cb8782860161338c565b60c08301525060e0830151828111156134e2575f80fd5b6134ee87828601612ca3565b60e08301525095945050505050565b6001600160a01b03831681526040602082018190525f906111e39083018461326f565b5f60208284031215613530575f80fd5b6109da82613113565b5f6001820161354a5761354a613090565b5060010190565b81516001600160a01b031681526020808301519082015260408101610674565b8381526001600160a01b03831660208201526060604082018190525f90612c8f90830184612761565b5f82516135ab81846020870161273f565b919091019291505056fea26469706673582212208629f6f30c98719aabaee78cff5b57c9a9667bace8fefc10a846374af4e07c8564736f6c63430008190033",
}

// TeleporterMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterMessengerMetaData.ABI instead.
var TeleporterMessengerABI = TeleporterMessengerMetaData.ABI

// TeleporterMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeleporterMessengerMetaData.Bin instead.
var TeleporterMessengerBin = TeleporterMessengerMetaData.Bin

// DeployTeleporterMessenger deploys a new Ethereum contract, binding an instance of TeleporterMessenger to it.
func DeployTeleporterMessenger(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TeleporterMessenger, error) {
	parsed, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeleporterMessengerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeleporterMessenger{TeleporterMessengerCaller: TeleporterMessengerCaller{contract: contract}, TeleporterMessengerTransactor: TeleporterMessengerTransactor{contract: contract}, TeleporterMessengerFilterer: TeleporterMessengerFilterer{contract: contract}}, nil
}

// TeleporterMessenger is an auto generated Go binding around an Ethereum contract.
type TeleporterMessenger struct {
	TeleporterMessengerCaller     // Read-only binding to the contract
	TeleporterMessengerTransactor // Write-only binding to the contract
	TeleporterMessengerFilterer   // Log filterer for contract events
}

// TeleporterMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterMessengerSession struct {
	Contract     *TeleporterMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TeleporterMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterMessengerCallerSession struct {
	Contract *TeleporterMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TeleporterMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterMessengerTransactorSession struct {
	Contract     *TeleporterMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TeleporterMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterMessengerRaw struct {
	Contract *TeleporterMessenger // Generic contract binding to access the raw methods on
}

// TeleporterMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterMessengerCallerRaw struct {
	Contract *TeleporterMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterMessengerTransactorRaw struct {
	Contract *TeleporterMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterMessenger creates a new instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessenger(address common.Address, backend bind.ContractBackend) (*TeleporterMessenger, error) {
	contract, err := bindTeleporterMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessenger{TeleporterMessengerCaller: TeleporterMessengerCaller{contract: contract}, TeleporterMessengerTransactor: TeleporterMessengerTransactor{contract: contract}, TeleporterMessengerFilterer: TeleporterMessengerFilterer{contract: contract}}, nil
}

// NewTeleporterMessengerCaller creates a new read-only instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerCaller(address common.Address, caller bind.ContractCaller) (*TeleporterMessengerCaller, error) {
	contract, err := bindTeleporterMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerCaller{contract: contract}, nil
}

// NewTeleporterMessengerTransactor creates a new write-only instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterMessengerTransactor, error) {
	contract, err := bindTeleporterMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerTransactor{contract: contract}, nil
}

// NewTeleporterMessengerFilterer creates a new log filterer instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterMessengerFilterer, error) {
	contract, err := bindTeleporterMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerFilterer{contract: contract}, nil
}

// bindTeleporterMessenger binds a generic wrapper to an already deployed contract.
func bindTeleporterMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterMessenger *TeleporterMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterMessenger.Contract.TeleporterMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterMessenger *TeleporterMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.TeleporterMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterMessenger *TeleporterMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.TeleporterMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterMessenger *TeleporterMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterMessenger *TeleporterMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterMessenger *TeleporterMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_TeleporterMessenger *TeleporterMessengerCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_TeleporterMessenger *TeleporterMessengerSession) VERSION() (string, error) {
	return _TeleporterMessenger.Contract.VERSION(&_TeleporterMessenger.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) VERSION() (string, error) {
	return _TeleporterMessenger.Contract.VERSION(&_TeleporterMessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterMessenger.Contract.BlockchainID(&_TeleporterMessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterMessenger.Contract.BlockchainID(&_TeleporterMessenger.CallOpts)
}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) pure returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) CalculateMessageID(opts *bind.CallOpts, sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "calculateMessageID", sourceBlockchainID, destinationBlockchainID, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) pure returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) CalculateMessageID(sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	return _TeleporterMessenger.Contract.CalculateMessageID(&_TeleporterMessenger.CallOpts, sourceBlockchainID, destinationBlockchainID, nonce)
}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) pure returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) CalculateMessageID(sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	return _TeleporterMessenger.Contract.CalculateMessageID(&_TeleporterMessenger.CallOpts, sourceBlockchainID, destinationBlockchainID, nonce)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) CheckRelayerRewardAmount(opts *bind.CallOpts, relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "checkRelayerRewardAmount", relayer, feeAsset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleporterMessenger.Contract.CheckRelayerRewardAmount(&_TeleporterMessenger.CallOpts, relayer, feeAsset)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleporterMessenger.Contract.CheckRelayerRewardAmount(&_TeleporterMessenger.CallOpts, relayer, feeAsset)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetFeeInfo(opts *bind.CallOpts, messageID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getFeeInfo", messageID)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _TeleporterMessenger.Contract.GetFeeInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _TeleporterMessenger.Contract.GetFeeInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetMessageHash(opts *bind.CallOpts, messageID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getMessageHash", messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetMessageHash(&_TeleporterMessenger.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetMessageHash(&_TeleporterMessenger.CallOpts, messageID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetNextMessageID(opts *bind.CallOpts, destinationBlockchainID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getNextMessageID", destinationBlockchainID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetNextMessageID(&_TeleporterMessenger.CallOpts, destinationBlockchainID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetNextMessageID(&_TeleporterMessenger.CallOpts, destinationBlockchainID)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerCaller) GetReceiptAtIndex(opts *bind.CallOpts, sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getReceiptAtIndex", sourceBlockchainID, index)

	if err != nil {
		return *new(TeleporterMessageReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(TeleporterMessageReceipt)).(*TeleporterMessageReceipt)

	return out0, err

}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerSession) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _TeleporterMessenger.Contract.GetReceiptAtIndex(&_TeleporterMessenger.CallOpts, sourceBlockchainID, index)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _TeleporterMessenger.Contract.GetReceiptAtIndex(&_TeleporterMessenger.CallOpts, sourceBlockchainID, index)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetReceiptQueueSize(opts *bind.CallOpts, sourceBlockchainID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getReceiptQueueSize", sourceBlockchainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.GetReceiptQueueSize(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.GetReceiptQueueSize(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetRelayerRewardAddress(opts *bind.CallOpts, messageID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getRelayerRewardAddress", messageID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerSession) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _TeleporterMessenger.Contract.GetRelayerRewardAddress(&_TeleporterMessenger.CallOpts, messageID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _TeleporterMessenger.Contract.GetRelayerRewardAddress(&_TeleporterMessenger.CallOpts, messageID)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) MessageNonce() (*big.Int, error) {
	return _TeleporterMessenger.Contract.MessageNonce(&_TeleporterMessenger.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) MessageNonce() (*big.Int, error) {
	return _TeleporterMessenger.Contract.MessageNonce(&_TeleporterMessenger.CallOpts)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerCaller) MessageReceived(opts *bind.CallOpts, messageID [32]byte) (bool, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "messageReceived", messageID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerSession) MessageReceived(messageID [32]byte) (bool, error) {
	return _TeleporterMessenger.Contract.MessageReceived(&_TeleporterMessenger.CallOpts, messageID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) MessageReceived(messageID [32]byte) (bool, error) {
	return _TeleporterMessenger.Contract.MessageReceived(&_TeleporterMessenger.CallOpts, messageID)
}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceBlockchainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessenger *TeleporterMessengerCaller) ReceiptQueues(opts *bind.CallOpts, sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "receiptQueues", sourceBlockchainID)

	outstruct := new(struct {
		First *big.Int
		Last  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.First = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Last = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceBlockchainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessenger *TeleporterMessengerSession) ReceiptQueues(sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _TeleporterMessenger.Contract.ReceiptQueues(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceBlockchainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) ReceiptQueues(sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _TeleporterMessenger.Contract.ReceiptQueues(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerCaller) ReceivedFailedMessageHashes(opts *bind.CallOpts, messageID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "receivedFailedMessageHashes", messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerSession) ReceivedFailedMessageHashes(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.ReceivedFailedMessageHashes(&_TeleporterMessenger.CallOpts, messageID)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) ReceivedFailedMessageHashes(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.ReceivedFailedMessageHashes(&_TeleporterMessenger.CallOpts, messageID)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x2ca40f55.
//
// Solidity: function sentMessageInfo(bytes32 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerCaller) SentMessageInfo(opts *bind.CallOpts, messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "sentMessageInfo", messageID)

	outstruct := new(struct {
		MessageHash [32]byte
		FeeInfo     TeleporterFeeInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MessageHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.FeeInfo = *abi.ConvertType(out[1], new(TeleporterFeeInfo)).(*TeleporterFeeInfo)

	return *outstruct, err

}

// SentMessageInfo is a free data retrieval call binding the contract method 0x2ca40f55.
//
// Solidity: function sentMessageInfo(bytes32 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerSession) SentMessageInfo(messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleporterMessenger.Contract.SentMessageInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x2ca40f55.
//
// Solidity: function sentMessageInfo(bytes32 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) SentMessageInfo(messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleporterMessenger.Contract.SentMessageInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) AddFeeAmount(opts *bind.TransactOpts, messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "addFeeAmount", messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.AddFeeAmount(&_TeleporterMessenger.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.AddFeeAmount(&_TeleporterMessenger.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address warpContract) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) Initialize(opts *bind.TransactOpts, warpContract common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "initialize", warpContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address warpContract) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) Initialize(warpContract common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.Initialize(&_TeleporterMessenger.TransactOpts, warpContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address warpContract) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) Initialize(warpContract common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.Initialize(&_TeleporterMessenger.TransactOpts, warpContract)
}

// InitializeBlockchainID is a paid mutator transaction binding the contract method 0x0af5b4ff.
//
// Solidity: function initializeBlockchainID() returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactor) InitializeBlockchainID(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "initializeBlockchainID")
}

// InitializeBlockchainID is a paid mutator transaction binding the contract method 0x0af5b4ff.
//
// Solidity: function initializeBlockchainID() returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) InitializeBlockchainID() (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.InitializeBlockchainID(&_TeleporterMessenger.TransactOpts)
}

// InitializeBlockchainID is a paid mutator transaction binding the contract method 0x0af5b4ff.
//
// Solidity: function initializeBlockchainID() returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) InitializeBlockchainID() (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.InitializeBlockchainID(&_TeleporterMessenger.TransactOpts)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) ReceiveCrossChainMessage(opts *bind.TransactOpts, messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "receiveCrossChainMessage", messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.ReceiveCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.ReceiveCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// ReceiveInterChainMessage is a paid mutator transaction binding the contract method 0x807c0be0.
//
// Solidity: function receiveInterChainMessage(((uint32,bytes32,bytes),bytes,(bytes,bytes)) icmMessage, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) ReceiveInterChainMessage(opts *bind.TransactOpts, icmMessage ICMMessage, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "receiveInterChainMessage", icmMessage, relayerRewardAddress)
}

// ReceiveInterChainMessage is a paid mutator transaction binding the contract method 0x807c0be0.
//
// Solidity: function receiveInterChainMessage(((uint32,bytes32,bytes),bytes,(bytes,bytes)) icmMessage, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) ReceiveInterChainMessage(icmMessage ICMMessage, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.ReceiveInterChainMessage(&_TeleporterMessenger.TransactOpts, icmMessage, relayerRewardAddress)
}

// ReceiveInterChainMessage is a paid mutator transaction binding the contract method 0x807c0be0.
//
// Solidity: function receiveInterChainMessage(((uint32,bytes32,bytes),bytes,(bytes,bytes)) icmMessage, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) ReceiveInterChainMessage(icmMessage ICMMessage, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.ReceiveInterChainMessage(&_TeleporterMessenger.TransactOpts, icmMessage, relayerRewardAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RedeemRelayerRewards(opts *bind.TransactOpts, feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "redeemRelayerRewards", feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RedeemRelayerRewards(&_TeleporterMessenger.TransactOpts, feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RedeemRelayerRewards(&_TeleporterMessenger.TransactOpts, feeAsset)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RetryMessageExecution(opts *bind.TransactOpts, sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "retryMessageExecution", sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetryMessageExecution(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetryMessageExecution(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RetrySendCrossChainMessage(opts *bind.TransactOpts, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "retrySendCrossChainMessage", message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RetrySendCrossChainMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetrySendCrossChainMessage(&_TeleporterMessenger.TransactOpts, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RetrySendCrossChainMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetrySendCrossChainMessage(&_TeleporterMessenger.TransactOpts, message)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactor) SendCrossChainMessage(opts *bind.TransactOpts, messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "sendCrossChainMessage", messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageInput)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactor) SendSpecifiedReceipts(opts *bind.TransactOpts, sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "sendSpecifiedReceipts", sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendSpecifiedReceipts(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendSpecifiedReceipts(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// TeleporterMessengerAddFeeAmountIterator is returned from FilterAddFeeAmount and is used to iterate over the raw logs and unpacked data for AddFeeAmount events raised by the TeleporterMessenger contract.
type TeleporterMessengerAddFeeAmountIterator struct {
	Event *TeleporterMessengerAddFeeAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerAddFeeAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerAddFeeAmount)
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
		it.Event = new(TeleporterMessengerAddFeeAmount)
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
func (it *TeleporterMessengerAddFeeAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerAddFeeAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerAddFeeAmount represents a AddFeeAmount event raised by the TeleporterMessenger contract.
type TeleporterMessengerAddFeeAmount struct {
	MessageID      [32]byte
	UpdatedFeeInfo TeleporterFeeInfo
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddFeeAmount is a free log retrieval operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterAddFeeAmount(opts *bind.FilterOpts, messageID [][32]byte) (*TeleporterMessengerAddFeeAmountIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerAddFeeAmountIterator{contract: _TeleporterMessenger.contract, event: "AddFeeAmount", logs: logs, sub: sub}, nil
}

// WatchAddFeeAmount is a free log subscription operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchAddFeeAmount(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerAddFeeAmount, messageID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerAddFeeAmount)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
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

// ParseAddFeeAmount is a log parse operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseAddFeeAmount(log types.Log) (*TeleporterMessengerAddFeeAmount, error) {
	event := new(TeleporterMessengerAddFeeAmount)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerBlockchainIDInitializedIterator is returned from FilterBlockchainIDInitialized and is used to iterate over the raw logs and unpacked data for BlockchainIDInitialized events raised by the TeleporterMessenger contract.
type TeleporterMessengerBlockchainIDInitializedIterator struct {
	Event *TeleporterMessengerBlockchainIDInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerBlockchainIDInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerBlockchainIDInitialized)
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
		it.Event = new(TeleporterMessengerBlockchainIDInitialized)
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
func (it *TeleporterMessengerBlockchainIDInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerBlockchainIDInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerBlockchainIDInitialized represents a BlockchainIDInitialized event raised by the TeleporterMessenger contract.
type TeleporterMessengerBlockchainIDInitialized struct {
	BlockchainID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBlockchainIDInitialized is a free log retrieval operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterBlockchainIDInitialized(opts *bind.FilterOpts, blockchainID [][32]byte) (*TeleporterMessengerBlockchainIDInitializedIterator, error) {

	var blockchainIDRule []interface{}
	for _, blockchainIDItem := range blockchainID {
		blockchainIDRule = append(blockchainIDRule, blockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "BlockchainIDInitialized", blockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerBlockchainIDInitializedIterator{contract: _TeleporterMessenger.contract, event: "BlockchainIDInitialized", logs: logs, sub: sub}, nil
}

// WatchBlockchainIDInitialized is a free log subscription operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchBlockchainIDInitialized(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerBlockchainIDInitialized, blockchainID [][32]byte) (event.Subscription, error) {

	var blockchainIDRule []interface{}
	for _, blockchainIDItem := range blockchainID {
		blockchainIDRule = append(blockchainIDRule, blockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "BlockchainIDInitialized", blockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerBlockchainIDInitialized)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "BlockchainIDInitialized", log); err != nil {
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

// ParseBlockchainIDInitialized is a log parse operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseBlockchainIDInitialized(log types.Log) (*TeleporterMessengerBlockchainIDInitialized, error) {
	event := new(TeleporterMessengerBlockchainIDInitialized)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "BlockchainIDInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TeleporterMessenger contract.
type TeleporterMessengerInitializedIterator struct {
	Event *TeleporterMessengerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerInitialized)
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
		it.Event = new(TeleporterMessengerInitialized)
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
func (it *TeleporterMessengerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerInitialized represents a Initialized event raised by the TeleporterMessenger contract.
type TeleporterMessengerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterInitialized(opts *bind.FilterOpts) (*TeleporterMessengerInitializedIterator, error) {

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerInitializedIterator{contract: _TeleporterMessenger.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerInitialized) (event.Subscription, error) {

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerInitialized)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseInitialized(log types.Log) (*TeleporterMessengerInitialized, error) {
	event := new(TeleporterMessengerInitialized)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerMessageExecutedIterator is returned from FilterMessageExecuted and is used to iterate over the raw logs and unpacked data for MessageExecuted events raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutedIterator struct {
	Event *TeleporterMessengerMessageExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerMessageExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerMessageExecuted)
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
		it.Event = new(TeleporterMessengerMessageExecuted)
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
func (it *TeleporterMessengerMessageExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerMessageExecuted represents a MessageExecuted event raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecuted struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecuted is a free log retrieval operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterMessageExecuted(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*TeleporterMessengerMessageExecutedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerMessageExecutedIterator{contract: _TeleporterMessenger.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

// WatchMessageExecuted is a free log subscription operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerMessageExecuted, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerMessageExecuted)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
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

// ParseMessageExecuted is a log parse operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseMessageExecuted(log types.Log) (*TeleporterMessengerMessageExecuted, error) {
	event := new(TeleporterMessengerMessageExecuted)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerMessageExecutionFailedIterator is returned from FilterMessageExecutionFailed and is used to iterate over the raw logs and unpacked data for MessageExecutionFailed events raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutionFailedIterator struct {
	Event *TeleporterMessengerMessageExecutionFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerMessageExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerMessageExecutionFailed)
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
		it.Event = new(TeleporterMessengerMessageExecutionFailed)
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
func (it *TeleporterMessengerMessageExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerMessageExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerMessageExecutionFailed represents a MessageExecutionFailed event raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutionFailed struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecutionFailed is a free log retrieval operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterMessageExecutionFailed(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*TeleporterMessengerMessageExecutionFailedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerMessageExecutionFailedIterator{contract: _TeleporterMessenger.contract, event: "MessageExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchMessageExecutionFailed is a free log subscription operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchMessageExecutionFailed(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerMessageExecutionFailed, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerMessageExecutionFailed)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
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

// ParseMessageExecutionFailed is a log parse operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseMessageExecutionFailed(log types.Log) (*TeleporterMessengerMessageExecutionFailed, error) {
	event := new(TeleporterMessengerMessageExecutionFailed)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerReceiptReceivedIterator is returned from FilterReceiptReceived and is used to iterate over the raw logs and unpacked data for ReceiptReceived events raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiptReceivedIterator struct {
	Event *TeleporterMessengerReceiptReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerReceiptReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerReceiptReceived)
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
		it.Event = new(TeleporterMessengerReceiptReceived)
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
func (it *TeleporterMessengerReceiptReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerReceiptReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerReceiptReceived represents a ReceiptReceived event raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiptReceived struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	RelayerRewardAddress    common.Address
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterReceiptReceived is a free log retrieval operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterReceiptReceived(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (*TeleporterMessengerReceiptReceivedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var relayerRewardAddressRule []interface{}
	for _, relayerRewardAddressItem := range relayerRewardAddress {
		relayerRewardAddressRule = append(relayerRewardAddressRule, relayerRewardAddressItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerReceiptReceivedIterator{contract: _TeleporterMessenger.contract, event: "ReceiptReceived", logs: logs, sub: sub}, nil
}

// WatchReceiptReceived is a free log subscription operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchReceiptReceived(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerReceiptReceived, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var relayerRewardAddressRule []interface{}
	for _, relayerRewardAddressItem := range relayerRewardAddress {
		relayerRewardAddressRule = append(relayerRewardAddressRule, relayerRewardAddressItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerReceiptReceived)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
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

// ParseReceiptReceived is a log parse operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseReceiptReceived(log types.Log) (*TeleporterMessengerReceiptReceived, error) {
	event := new(TeleporterMessengerReceiptReceived)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerReceiveCrossChainMessageIterator is returned from FilterReceiveCrossChainMessage and is used to iterate over the raw logs and unpacked data for ReceiveCrossChainMessage events raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiveCrossChainMessageIterator struct {
	Event *TeleporterMessengerReceiveCrossChainMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerReceiveCrossChainMessage)
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
		it.Event = new(TeleporterMessengerReceiveCrossChainMessage)
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
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerReceiveCrossChainMessage represents a ReceiveCrossChainMessage event raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiveCrossChainMessage struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Deliverer          common.Address
	RewardRedeemer     common.Address
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReceiveCrossChainMessage is a free log retrieval operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterReceiveCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (*TeleporterMessengerReceiveCrossChainMessageIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerReceiveCrossChainMessageIterator{contract: _TeleporterMessenger.contract, event: "ReceiveCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveCrossChainMessage is a free log subscription operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchReceiveCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerReceiveCrossChainMessage, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerReceiveCrossChainMessage)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
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

// ParseReceiveCrossChainMessage is a log parse operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseReceiveCrossChainMessage(log types.Log) (*TeleporterMessengerReceiveCrossChainMessage, error) {
	event := new(TeleporterMessengerReceiveCrossChainMessage)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerRelayerRewardsRedeemedIterator is returned from FilterRelayerRewardsRedeemed and is used to iterate over the raw logs and unpacked data for RelayerRewardsRedeemed events raised by the TeleporterMessenger contract.
type TeleporterMessengerRelayerRewardsRedeemedIterator struct {
	Event *TeleporterMessengerRelayerRewardsRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerRelayerRewardsRedeemed)
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
		it.Event = new(TeleporterMessengerRelayerRewardsRedeemed)
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
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerRelayerRewardsRedeemed represents a RelayerRewardsRedeemed event raised by the TeleporterMessenger contract.
type TeleporterMessengerRelayerRewardsRedeemed struct {
	Redeemer common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelayerRewardsRedeemed is a free log retrieval operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterRelayerRewardsRedeemed(opts *bind.FilterOpts, redeemer []common.Address, asset []common.Address) (*TeleporterMessengerRelayerRewardsRedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerRelayerRewardsRedeemedIterator{contract: _TeleporterMessenger.contract, event: "RelayerRewardsRedeemed", logs: logs, sub: sub}, nil
}

// WatchRelayerRewardsRedeemed is a free log subscription operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchRelayerRewardsRedeemed(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerRelayerRewardsRedeemed, redeemer []common.Address, asset []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerRelayerRewardsRedeemed)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
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

// ParseRelayerRewardsRedeemed is a log parse operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseRelayerRewardsRedeemed(log types.Log) (*TeleporterMessengerRelayerRewardsRedeemed, error) {
	event := new(TeleporterMessengerRelayerRewardsRedeemed)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerSendCrossChainMessageIterator is returned from FilterSendCrossChainMessage and is used to iterate over the raw logs and unpacked data for SendCrossChainMessage events raised by the TeleporterMessenger contract.
type TeleporterMessengerSendCrossChainMessageIterator struct {
	Event *TeleporterMessengerSendCrossChainMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerSendCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerSendCrossChainMessage)
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
		it.Event = new(TeleporterMessengerSendCrossChainMessage)
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
func (it *TeleporterMessengerSendCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerSendCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerSendCrossChainMessage represents a SendCrossChainMessage event raised by the TeleporterMessenger contract.
type TeleporterMessengerSendCrossChainMessage struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	Message                 TeleporterMessage
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSendCrossChainMessage is a free log retrieval operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterSendCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte) (*TeleporterMessengerSendCrossChainMessageIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerSendCrossChainMessageIterator{contract: _TeleporterMessenger.contract, event: "SendCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchSendCrossChainMessage is a free log subscription operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchSendCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerSendCrossChainMessage, messageID [][32]byte, destinationBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerSendCrossChainMessage)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
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

// ParseSendCrossChainMessage is a log parse operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseSendCrossChainMessage(log types.Log) (*TeleporterMessengerSendCrossChainMessage, error) {
	event := new(TeleporterMessengerSendCrossChainMessage)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
