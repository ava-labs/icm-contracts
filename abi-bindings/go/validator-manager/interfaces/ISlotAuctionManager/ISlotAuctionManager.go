// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package islotauctionmanager

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

// ISlotAuctionManagerMetaData contains all meta data concerning the ISlotAuctionManager contract.
var ISlotAuctionManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"BidEvicted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitiatedAuctionValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"validatorSlots\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minValidatorDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"}],\"name\":\"NewValidatorAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"SuccessfulBidPlaced\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuctionCooldownDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuctioningValidatorWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinAuctionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinValidatorDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpenValidatorSlots\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalValidatorSlots\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initiateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBidRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISlotAuctionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ISlotAuctionManagerMetaData.ABI instead.
var ISlotAuctionManagerABI = ISlotAuctionManagerMetaData.ABI

// ISlotAuctionManager is an auto generated Go binding around an Ethereum contract.
type ISlotAuctionManager struct {
	ISlotAuctionManagerCaller     // Read-only binding to the contract
	ISlotAuctionManagerTransactor // Write-only binding to the contract
	ISlotAuctionManagerFilterer   // Log filterer for contract events
}

// ISlotAuctionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISlotAuctionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlotAuctionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISlotAuctionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlotAuctionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISlotAuctionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlotAuctionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISlotAuctionManagerSession struct {
	Contract     *ISlotAuctionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ISlotAuctionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISlotAuctionManagerCallerSession struct {
	Contract *ISlotAuctionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ISlotAuctionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISlotAuctionManagerTransactorSession struct {
	Contract     *ISlotAuctionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ISlotAuctionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISlotAuctionManagerRaw struct {
	Contract *ISlotAuctionManager // Generic contract binding to access the raw methods on
}

// ISlotAuctionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISlotAuctionManagerCallerRaw struct {
	Contract *ISlotAuctionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ISlotAuctionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISlotAuctionManagerTransactorRaw struct {
	Contract *ISlotAuctionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISlotAuctionManager creates a new instance of ISlotAuctionManager, bound to a specific deployed contract.
func NewISlotAuctionManager(address common.Address, backend bind.ContractBackend) (*ISlotAuctionManager, error) {
	contract, err := bindISlotAuctionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISlotAuctionManager{ISlotAuctionManagerCaller: ISlotAuctionManagerCaller{contract: contract}, ISlotAuctionManagerTransactor: ISlotAuctionManagerTransactor{contract: contract}, ISlotAuctionManagerFilterer: ISlotAuctionManagerFilterer{contract: contract}}, nil
}

// NewISlotAuctionManagerCaller creates a new read-only instance of ISlotAuctionManager, bound to a specific deployed contract.
func NewISlotAuctionManagerCaller(address common.Address, caller bind.ContractCaller) (*ISlotAuctionManagerCaller, error) {
	contract, err := bindISlotAuctionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISlotAuctionManagerCaller{contract: contract}, nil
}

// NewISlotAuctionManagerTransactor creates a new write-only instance of ISlotAuctionManager, bound to a specific deployed contract.
func NewISlotAuctionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ISlotAuctionManagerTransactor, error) {
	contract, err := bindISlotAuctionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISlotAuctionManagerTransactor{contract: contract}, nil
}

// NewISlotAuctionManagerFilterer creates a new log filterer instance of ISlotAuctionManager, bound to a specific deployed contract.
func NewISlotAuctionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ISlotAuctionManagerFilterer, error) {
	contract, err := bindISlotAuctionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISlotAuctionManagerFilterer{contract: contract}, nil
}

// bindISlotAuctionManager binds a generic wrapper to an already deployed contract.
func bindISlotAuctionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISlotAuctionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlotAuctionManager *ISlotAuctionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlotAuctionManager.Contract.ISlotAuctionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlotAuctionManager *ISlotAuctionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.ISlotAuctionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlotAuctionManager *ISlotAuctionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.ISlotAuctionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlotAuctionManager *ISlotAuctionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlotAuctionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlotAuctionManager *ISlotAuctionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlotAuctionManager *ISlotAuctionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.contract.Transact(opts, method, params...)
}

// GetAuctionCooldownDuration is a free data retrieval call binding the contract method 0xf714b6aa.
//
// Solidity: function getAuctionCooldownDuration() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerCaller) GetAuctionCooldownDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISlotAuctionManager.contract.Call(opts, &out, "getAuctionCooldownDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAuctionCooldownDuration is a free data retrieval call binding the contract method 0xf714b6aa.
//
// Solidity: function getAuctionCooldownDuration() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerSession) GetAuctionCooldownDuration() (*big.Int, error) {
	return _ISlotAuctionManager.Contract.GetAuctionCooldownDuration(&_ISlotAuctionManager.CallOpts)
}

// GetAuctionCooldownDuration is a free data retrieval call binding the contract method 0xf714b6aa.
//
// Solidity: function getAuctionCooldownDuration() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerCallerSession) GetAuctionCooldownDuration() (*big.Int, error) {
	return _ISlotAuctionManager.Contract.GetAuctionCooldownDuration(&_ISlotAuctionManager.CallOpts)
}

// GetAuctioningValidatorWeight is a free data retrieval call binding the contract method 0x671ef477.
//
// Solidity: function getAuctioningValidatorWeight() view returns(uint64)
func (_ISlotAuctionManager *ISlotAuctionManagerCaller) GetAuctioningValidatorWeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ISlotAuctionManager.contract.Call(opts, &out, "getAuctioningValidatorWeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAuctioningValidatorWeight is a free data retrieval call binding the contract method 0x671ef477.
//
// Solidity: function getAuctioningValidatorWeight() view returns(uint64)
func (_ISlotAuctionManager *ISlotAuctionManagerSession) GetAuctioningValidatorWeight() (uint64, error) {
	return _ISlotAuctionManager.Contract.GetAuctioningValidatorWeight(&_ISlotAuctionManager.CallOpts)
}

// GetAuctioningValidatorWeight is a free data retrieval call binding the contract method 0x671ef477.
//
// Solidity: function getAuctioningValidatorWeight() view returns(uint64)
func (_ISlotAuctionManager *ISlotAuctionManagerCallerSession) GetAuctioningValidatorWeight() (uint64, error) {
	return _ISlotAuctionManager.Contract.GetAuctioningValidatorWeight(&_ISlotAuctionManager.CallOpts)
}

// GetMinAuctionDuration is a free data retrieval call binding the contract method 0x9940d4c6.
//
// Solidity: function getMinAuctionDuration() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerCaller) GetMinAuctionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISlotAuctionManager.contract.Call(opts, &out, "getMinAuctionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinAuctionDuration is a free data retrieval call binding the contract method 0x9940d4c6.
//
// Solidity: function getMinAuctionDuration() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerSession) GetMinAuctionDuration() (*big.Int, error) {
	return _ISlotAuctionManager.Contract.GetMinAuctionDuration(&_ISlotAuctionManager.CallOpts)
}

// GetMinAuctionDuration is a free data retrieval call binding the contract method 0x9940d4c6.
//
// Solidity: function getMinAuctionDuration() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerCallerSession) GetMinAuctionDuration() (*big.Int, error) {
	return _ISlotAuctionManager.Contract.GetMinAuctionDuration(&_ISlotAuctionManager.CallOpts)
}

// GetMinValidatorDuration is a free data retrieval call binding the contract method 0x01f6cec8.
//
// Solidity: function getMinValidatorDuration() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerCaller) GetMinValidatorDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISlotAuctionManager.contract.Call(opts, &out, "getMinValidatorDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinValidatorDuration is a free data retrieval call binding the contract method 0x01f6cec8.
//
// Solidity: function getMinValidatorDuration() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerSession) GetMinValidatorDuration() (*big.Int, error) {
	return _ISlotAuctionManager.Contract.GetMinValidatorDuration(&_ISlotAuctionManager.CallOpts)
}

// GetMinValidatorDuration is a free data retrieval call binding the contract method 0x01f6cec8.
//
// Solidity: function getMinValidatorDuration() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerCallerSession) GetMinValidatorDuration() (*big.Int, error) {
	return _ISlotAuctionManager.Contract.GetMinValidatorDuration(&_ISlotAuctionManager.CallOpts)
}

// GetMinimumBid is a free data retrieval call binding the contract method 0xc5b63600.
//
// Solidity: function getMinimumBid() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerCaller) GetMinimumBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISlotAuctionManager.contract.Call(opts, &out, "getMinimumBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumBid is a free data retrieval call binding the contract method 0xc5b63600.
//
// Solidity: function getMinimumBid() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerSession) GetMinimumBid() (*big.Int, error) {
	return _ISlotAuctionManager.Contract.GetMinimumBid(&_ISlotAuctionManager.CallOpts)
}

// GetMinimumBid is a free data retrieval call binding the contract method 0xc5b63600.
//
// Solidity: function getMinimumBid() view returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerCallerSession) GetMinimumBid() (*big.Int, error) {
	return _ISlotAuctionManager.Contract.GetMinimumBid(&_ISlotAuctionManager.CallOpts)
}

// GetOpenValidatorSlots is a free data retrieval call binding the contract method 0x78eb9c10.
//
// Solidity: function getOpenValidatorSlots() view returns(uint16)
func (_ISlotAuctionManager *ISlotAuctionManagerCaller) GetOpenValidatorSlots(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ISlotAuctionManager.contract.Call(opts, &out, "getOpenValidatorSlots")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetOpenValidatorSlots is a free data retrieval call binding the contract method 0x78eb9c10.
//
// Solidity: function getOpenValidatorSlots() view returns(uint16)
func (_ISlotAuctionManager *ISlotAuctionManagerSession) GetOpenValidatorSlots() (uint16, error) {
	return _ISlotAuctionManager.Contract.GetOpenValidatorSlots(&_ISlotAuctionManager.CallOpts)
}

// GetOpenValidatorSlots is a free data retrieval call binding the contract method 0x78eb9c10.
//
// Solidity: function getOpenValidatorSlots() view returns(uint16)
func (_ISlotAuctionManager *ISlotAuctionManagerCallerSession) GetOpenValidatorSlots() (uint16, error) {
	return _ISlotAuctionManager.Contract.GetOpenValidatorSlots(&_ISlotAuctionManager.CallOpts)
}

// GetTotalValidatorSlots is a free data retrieval call binding the contract method 0x31b6822e.
//
// Solidity: function getTotalValidatorSlots() view returns(uint16)
func (_ISlotAuctionManager *ISlotAuctionManagerCaller) GetTotalValidatorSlots(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ISlotAuctionManager.contract.Call(opts, &out, "getTotalValidatorSlots")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetTotalValidatorSlots is a free data retrieval call binding the contract method 0x31b6822e.
//
// Solidity: function getTotalValidatorSlots() view returns(uint16)
func (_ISlotAuctionManager *ISlotAuctionManagerSession) GetTotalValidatorSlots() (uint16, error) {
	return _ISlotAuctionManager.Contract.GetTotalValidatorSlots(&_ISlotAuctionManager.CallOpts)
}

// GetTotalValidatorSlots is a free data retrieval call binding the contract method 0x31b6822e.
//
// Solidity: function getTotalValidatorSlots() view returns(uint16)
func (_ISlotAuctionManager *ISlotAuctionManagerCallerSession) GetTotalValidatorSlots() (uint16, error) {
	return _ISlotAuctionManager.Contract.GetTotalValidatorSlots(&_ISlotAuctionManager.CallOpts)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_ISlotAuctionManager *ISlotAuctionManagerTransactor) CompleteRemoveInitialValidator(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ISlotAuctionManager.contract.Transact(opts, "completeRemoveInitialValidator", messageIndex)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_ISlotAuctionManager *ISlotAuctionManagerSession) CompleteRemoveInitialValidator(messageIndex uint32) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.CompleteRemoveInitialValidator(&_ISlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_ISlotAuctionManager *ISlotAuctionManagerTransactorSession) CompleteRemoveInitialValidator(messageIndex uint32) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.CompleteRemoveInitialValidator(&_ISlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ISlotAuctionManager *ISlotAuctionManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ISlotAuctionManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ISlotAuctionManager *ISlotAuctionManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.CompleteValidatorRegistration(&_ISlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ISlotAuctionManager *ISlotAuctionManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.CompleteValidatorRegistration(&_ISlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ISlotAuctionManager *ISlotAuctionManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ISlotAuctionManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ISlotAuctionManager *ISlotAuctionManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.CompleteValidatorRemoval(&_ISlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ISlotAuctionManager *ISlotAuctionManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.CompleteValidatorRemoval(&_ISlotAuctionManager.TransactOpts, messageIndex)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_ISlotAuctionManager *ISlotAuctionManagerTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlotAuctionManager.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_ISlotAuctionManager *ISlotAuctionManagerSession) EndAuction() (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.EndAuction(&_ISlotAuctionManager.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_ISlotAuctionManager *ISlotAuctionManagerTransactorSession) EndAuction() (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.EndAuction(&_ISlotAuctionManager.TransactOpts)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x5561f9df.
//
// Solidity: function initiateAuction() returns()
func (_ISlotAuctionManager *ISlotAuctionManagerTransactor) InitiateAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlotAuctionManager.contract.Transact(opts, "initiateAuction")
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x5561f9df.
//
// Solidity: function initiateAuction() returns()
func (_ISlotAuctionManager *ISlotAuctionManagerSession) InitiateAuction() (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.InitiateAuction(&_ISlotAuctionManager.TransactOpts)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x5561f9df.
//
// Solidity: function initiateAuction() returns()
func (_ISlotAuctionManager *ISlotAuctionManagerTransactorSession) InitiateAuction() (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.InitiateAuction(&_ISlotAuctionManager.TransactOpts)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_ISlotAuctionManager *ISlotAuctionManagerTransactor) InitiateRemoveInitialValidator(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ISlotAuctionManager.contract.Transact(opts, "initiateRemoveInitialValidator", validationID)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_ISlotAuctionManager *ISlotAuctionManagerSession) InitiateRemoveInitialValidator(validationID [32]byte) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.InitiateRemoveInitialValidator(&_ISlotAuctionManager.TransactOpts, validationID)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_ISlotAuctionManager *ISlotAuctionManagerTransactorSession) InitiateRemoveInitialValidator(validationID [32]byte) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.InitiateRemoveInitialValidator(&_ISlotAuctionManager.TransactOpts, validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_ISlotAuctionManager *ISlotAuctionManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ISlotAuctionManager.contract.Transact(opts, "initiateValidatorRemoval", validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_ISlotAuctionManager *ISlotAuctionManagerSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.InitiateValidatorRemoval(&_ISlotAuctionManager.TransactOpts, validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_ISlotAuctionManager *ISlotAuctionManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.InitiateValidatorRemoval(&_ISlotAuctionManager.TransactOpts, validationID)
}

// MinBidRequired is a paid mutator transaction binding the contract method 0xf1c7b8d2.
//
// Solidity: function minBidRequired() returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerTransactor) MinBidRequired(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlotAuctionManager.contract.Transact(opts, "minBidRequired")
}

// MinBidRequired is a paid mutator transaction binding the contract method 0xf1c7b8d2.
//
// Solidity: function minBidRequired() returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerSession) MinBidRequired() (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.MinBidRequired(&_ISlotAuctionManager.TransactOpts)
}

// MinBidRequired is a paid mutator transaction binding the contract method 0xf1c7b8d2.
//
// Solidity: function minBidRequired() returns(uint256)
func (_ISlotAuctionManager *ISlotAuctionManagerTransactorSession) MinBidRequired() (*types.Transaction, error) {
	return _ISlotAuctionManager.Contract.MinBidRequired(&_ISlotAuctionManager.TransactOpts)
}

// ISlotAuctionManagerBidEvictedIterator is returned from FilterBidEvicted and is used to iterate over the raw logs and unpacked data for BidEvicted events raised by the ISlotAuctionManager contract.
type ISlotAuctionManagerBidEvictedIterator struct {
	Event *ISlotAuctionManagerBidEvicted // Event containing the contract specifics and raw log

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
func (it *ISlotAuctionManagerBidEvictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlotAuctionManagerBidEvicted)
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
		it.Event = new(ISlotAuctionManagerBidEvicted)
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
func (it *ISlotAuctionManagerBidEvictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlotAuctionManagerBidEvictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlotAuctionManagerBidEvicted represents a BidEvicted event raised by the ISlotAuctionManager contract.
type ISlotAuctionManagerBidEvicted struct {
	Bid    *big.Int
	NodeID common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidEvicted is a free log retrieval operation binding the contract event 0xdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc.
//
// Solidity: event BidEvicted(uint256 indexed bid, bytes indexed nodeID)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) FilterBidEvicted(opts *bind.FilterOpts, bid []*big.Int, nodeID [][]byte) (*ISlotAuctionManagerBidEvictedIterator, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ISlotAuctionManager.contract.FilterLogs(opts, "BidEvicted", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ISlotAuctionManagerBidEvictedIterator{contract: _ISlotAuctionManager.contract, event: "BidEvicted", logs: logs, sub: sub}, nil
}

// WatchBidEvicted is a free log subscription operation binding the contract event 0xdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc.
//
// Solidity: event BidEvicted(uint256 indexed bid, bytes indexed nodeID)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) WatchBidEvicted(opts *bind.WatchOpts, sink chan<- *ISlotAuctionManagerBidEvicted, bid []*big.Int, nodeID [][]byte) (event.Subscription, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ISlotAuctionManager.contract.WatchLogs(opts, "BidEvicted", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlotAuctionManagerBidEvicted)
				if err := _ISlotAuctionManager.contract.UnpackLog(event, "BidEvicted", log); err != nil {
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

// ParseBidEvicted is a log parse operation binding the contract event 0xdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc.
//
// Solidity: event BidEvicted(uint256 indexed bid, bytes indexed nodeID)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) ParseBidEvicted(log types.Log) (*ISlotAuctionManagerBidEvicted, error) {
	event := new(ISlotAuctionManagerBidEvicted)
	if err := _ISlotAuctionManager.contract.UnpackLog(event, "BidEvicted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator is returned from FilterInitiatedAuctionValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedAuctionValidatorRegistration events raised by the ISlotAuctionManager contract.
type ISlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator struct {
	Event *ISlotAuctionManagerInitiatedAuctionValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *ISlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlotAuctionManagerInitiatedAuctionValidatorRegistration)
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
		it.Event = new(ISlotAuctionManagerInitiatedAuctionValidatorRegistration)
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
func (it *ISlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlotAuctionManagerInitiatedAuctionValidatorRegistration represents a InitiatedAuctionValidatorRegistration event raised by the ISlotAuctionManager contract.
type ISlotAuctionManagerInitiatedAuctionValidatorRegistration struct {
	ValidationID     [32]byte
	OwnerAddress     common.Address
	ValidatorEndTime *big.Int
	Weight           uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInitiatedAuctionValidatorRegistration is a free log retrieval operation binding the contract event 0x032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 validatorEndTime, uint64 weight)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) FilterInitiatedAuctionValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, ownerAddress []common.Address) (*ISlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _ISlotAuctionManager.contract.FilterLogs(opts, "InitiatedAuctionValidatorRegistration", validationIDRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ISlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator{contract: _ISlotAuctionManager.contract, event: "InitiatedAuctionValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedAuctionValidatorRegistration is a free log subscription operation binding the contract event 0x032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 validatorEndTime, uint64 weight)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) WatchInitiatedAuctionValidatorRegistration(opts *bind.WatchOpts, sink chan<- *ISlotAuctionManagerInitiatedAuctionValidatorRegistration, validationID [][32]byte, ownerAddress []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _ISlotAuctionManager.contract.WatchLogs(opts, "InitiatedAuctionValidatorRegistration", validationIDRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlotAuctionManagerInitiatedAuctionValidatorRegistration)
				if err := _ISlotAuctionManager.contract.UnpackLog(event, "InitiatedAuctionValidatorRegistration", log); err != nil {
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

// ParseInitiatedAuctionValidatorRegistration is a log parse operation binding the contract event 0x032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 validatorEndTime, uint64 weight)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) ParseInitiatedAuctionValidatorRegistration(log types.Log) (*ISlotAuctionManagerInitiatedAuctionValidatorRegistration, error) {
	event := new(ISlotAuctionManagerInitiatedAuctionValidatorRegistration)
	if err := _ISlotAuctionManager.contract.UnpackLog(event, "InitiatedAuctionValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlotAuctionManagerNewValidatorAuctionIterator is returned from FilterNewValidatorAuction and is used to iterate over the raw logs and unpacked data for NewValidatorAuction events raised by the ISlotAuctionManager contract.
type ISlotAuctionManagerNewValidatorAuctionIterator struct {
	Event *ISlotAuctionManagerNewValidatorAuction // Event containing the contract specifics and raw log

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
func (it *ISlotAuctionManagerNewValidatorAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlotAuctionManagerNewValidatorAuction)
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
		it.Event = new(ISlotAuctionManagerNewValidatorAuction)
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
func (it *ISlotAuctionManagerNewValidatorAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlotAuctionManagerNewValidatorAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlotAuctionManagerNewValidatorAuction represents a NewValidatorAuction event raised by the ISlotAuctionManager contract.
type ISlotAuctionManagerNewValidatorAuction struct {
	ValidatorSlots       uint16
	ValidatorWeight      uint64
	MinValidatorDuration *big.Int
	AuctionEndTime       *big.Int
	MinimumBid           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewValidatorAuction is a free log retrieval operation binding the contract event 0x50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b54.
//
// Solidity: event NewValidatorAuction(uint16 validatorSlots, uint64 validatorWeight, uint256 minValidatorDuration, uint256 auctionEndTime, uint256 minimumBid)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) FilterNewValidatorAuction(opts *bind.FilterOpts) (*ISlotAuctionManagerNewValidatorAuctionIterator, error) {

	logs, sub, err := _ISlotAuctionManager.contract.FilterLogs(opts, "NewValidatorAuction")
	if err != nil {
		return nil, err
	}
	return &ISlotAuctionManagerNewValidatorAuctionIterator{contract: _ISlotAuctionManager.contract, event: "NewValidatorAuction", logs: logs, sub: sub}, nil
}

// WatchNewValidatorAuction is a free log subscription operation binding the contract event 0x50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b54.
//
// Solidity: event NewValidatorAuction(uint16 validatorSlots, uint64 validatorWeight, uint256 minValidatorDuration, uint256 auctionEndTime, uint256 minimumBid)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) WatchNewValidatorAuction(opts *bind.WatchOpts, sink chan<- *ISlotAuctionManagerNewValidatorAuction) (event.Subscription, error) {

	logs, sub, err := _ISlotAuctionManager.contract.WatchLogs(opts, "NewValidatorAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlotAuctionManagerNewValidatorAuction)
				if err := _ISlotAuctionManager.contract.UnpackLog(event, "NewValidatorAuction", log); err != nil {
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

// ParseNewValidatorAuction is a log parse operation binding the contract event 0x50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b54.
//
// Solidity: event NewValidatorAuction(uint16 validatorSlots, uint64 validatorWeight, uint256 minValidatorDuration, uint256 auctionEndTime, uint256 minimumBid)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) ParseNewValidatorAuction(log types.Log) (*ISlotAuctionManagerNewValidatorAuction, error) {
	event := new(ISlotAuctionManagerNewValidatorAuction)
	if err := _ISlotAuctionManager.contract.UnpackLog(event, "NewValidatorAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlotAuctionManagerSuccessfulBidPlacedIterator is returned from FilterSuccessfulBidPlaced and is used to iterate over the raw logs and unpacked data for SuccessfulBidPlaced events raised by the ISlotAuctionManager contract.
type ISlotAuctionManagerSuccessfulBidPlacedIterator struct {
	Event *ISlotAuctionManagerSuccessfulBidPlaced // Event containing the contract specifics and raw log

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
func (it *ISlotAuctionManagerSuccessfulBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlotAuctionManagerSuccessfulBidPlaced)
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
		it.Event = new(ISlotAuctionManagerSuccessfulBidPlaced)
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
func (it *ISlotAuctionManagerSuccessfulBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlotAuctionManagerSuccessfulBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlotAuctionManagerSuccessfulBidPlaced represents a SuccessfulBidPlaced event raised by the ISlotAuctionManager contract.
type ISlotAuctionManagerSuccessfulBidPlaced struct {
	Bid    *big.Int
	NodeID common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSuccessfulBidPlaced is a free log retrieval operation binding the contract event 0x864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029.
//
// Solidity: event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) FilterSuccessfulBidPlaced(opts *bind.FilterOpts, bid []*big.Int, nodeID [][]byte) (*ISlotAuctionManagerSuccessfulBidPlacedIterator, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ISlotAuctionManager.contract.FilterLogs(opts, "SuccessfulBidPlaced", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ISlotAuctionManagerSuccessfulBidPlacedIterator{contract: _ISlotAuctionManager.contract, event: "SuccessfulBidPlaced", logs: logs, sub: sub}, nil
}

// WatchSuccessfulBidPlaced is a free log subscription operation binding the contract event 0x864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029.
//
// Solidity: event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) WatchSuccessfulBidPlaced(opts *bind.WatchOpts, sink chan<- *ISlotAuctionManagerSuccessfulBidPlaced, bid []*big.Int, nodeID [][]byte) (event.Subscription, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ISlotAuctionManager.contract.WatchLogs(opts, "SuccessfulBidPlaced", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlotAuctionManagerSuccessfulBidPlaced)
				if err := _ISlotAuctionManager.contract.UnpackLog(event, "SuccessfulBidPlaced", log); err != nil {
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

// ParseSuccessfulBidPlaced is a log parse operation binding the contract event 0x864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029.
//
// Solidity: event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID)
func (_ISlotAuctionManager *ISlotAuctionManagerFilterer) ParseSuccessfulBidPlaced(log types.Log) (*ISlotAuctionManagerSuccessfulBidPlaced, error) {
	event := new(ISlotAuctionManagerSuccessfulBidPlaced)
	if err := _ISlotAuctionManager.contract.UnpackLog(event, "SuccessfulBidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
