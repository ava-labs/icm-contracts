// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package acp99manager

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

// ConversionData is an auto generated low-level Go binding around an user-defined struct.
type ConversionData struct {
	SubnetID                     [32]byte
	ValidatorManagerBlockchainID [32]byte
	ValidatorManagerAddress      common.Address
	InitialValidators            []InitialValidator
}

// InitialValidator is an auto generated low-level Go binding around an user-defined struct.
type InitialValidator struct {
	NodeID       []byte
	BlsPublicKey []byte
	Weight       uint64
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Status         uint8
	NodeID         []byte
	StartingWeight uint64
	SentNonce      uint64
	ReceivedNonce  uint64
	Weight         uint64
	StartTime      uint64
	EndTime        uint64
}

// ACP99ManagerMetaData contains all meta data concerning the ACP99Manager contract.
var ACP99ManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"CompletedValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"CompletedValidatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"CompletedValidatorWeightUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeID\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"registrationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitiatedValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"validatorWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"InitiatedValidatorRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"weightUpdateMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitiatedValidatorWeightUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeID\",\"type\":\"bytes20\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"RegisteredInitialValidator\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorWeightUpdate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startingWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sentNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receivedNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"internalType\":\"structValidator\",\"name\":\"validator\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structConversionData\",\"name\":\"conversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1TotalWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subnetID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ACP99ManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ACP99ManagerMetaData.ABI instead.
var ACP99ManagerABI = ACP99ManagerMetaData.ABI

// ACP99Manager is an auto generated Go binding around an Ethereum contract.
type ACP99Manager struct {
	ACP99ManagerCaller     // Read-only binding to the contract
	ACP99ManagerTransactor // Write-only binding to the contract
	ACP99ManagerFilterer   // Log filterer for contract events
}

// ACP99ManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ACP99ManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACP99ManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ACP99ManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACP99ManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ACP99ManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ACP99ManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ACP99ManagerSession struct {
	Contract     *ACP99Manager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ACP99ManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ACP99ManagerCallerSession struct {
	Contract *ACP99ManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ACP99ManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ACP99ManagerTransactorSession struct {
	Contract     *ACP99ManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ACP99ManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ACP99ManagerRaw struct {
	Contract *ACP99Manager // Generic contract binding to access the raw methods on
}

// ACP99ManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ACP99ManagerCallerRaw struct {
	Contract *ACP99ManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ACP99ManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ACP99ManagerTransactorRaw struct {
	Contract *ACP99ManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewACP99Manager creates a new instance of ACP99Manager, bound to a specific deployed contract.
func NewACP99Manager(address common.Address, backend bind.ContractBackend) (*ACP99Manager, error) {
	contract, err := bindACP99Manager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ACP99Manager{ACP99ManagerCaller: ACP99ManagerCaller{contract: contract}, ACP99ManagerTransactor: ACP99ManagerTransactor{contract: contract}, ACP99ManagerFilterer: ACP99ManagerFilterer{contract: contract}}, nil
}

// NewACP99ManagerCaller creates a new read-only instance of ACP99Manager, bound to a specific deployed contract.
func NewACP99ManagerCaller(address common.Address, caller bind.ContractCaller) (*ACP99ManagerCaller, error) {
	contract, err := bindACP99Manager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ACP99ManagerCaller{contract: contract}, nil
}

// NewACP99ManagerTransactor creates a new write-only instance of ACP99Manager, bound to a specific deployed contract.
func NewACP99ManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ACP99ManagerTransactor, error) {
	contract, err := bindACP99Manager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ACP99ManagerTransactor{contract: contract}, nil
}

// NewACP99ManagerFilterer creates a new log filterer instance of ACP99Manager, bound to a specific deployed contract.
func NewACP99ManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ACP99ManagerFilterer, error) {
	contract, err := bindACP99Manager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ACP99ManagerFilterer{contract: contract}, nil
}

// bindACP99Manager binds a generic wrapper to an already deployed contract.
func bindACP99Manager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ACP99ManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACP99Manager *ACP99ManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACP99Manager.Contract.ACP99ManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACP99Manager *ACP99ManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACP99Manager.Contract.ACP99ManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACP99Manager *ACP99ManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACP99Manager.Contract.ACP99ManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ACP99Manager *ACP99ManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ACP99Manager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ACP99Manager *ACP99ManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ACP99Manager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ACP99Manager *ACP99ManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ACP99Manager.Contract.contract.Transact(opts, method, params...)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64,uint64) validator)
func (_ACP99Manager *ACP99ManagerCaller) GetValidator(opts *bind.CallOpts, validationID [32]byte) (Validator, error) {
	var out []interface{}
	err := _ACP99Manager.contract.Call(opts, &out, "getValidator", validationID)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64,uint64) validator)
func (_ACP99Manager *ACP99ManagerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ACP99Manager.Contract.GetValidator(&_ACP99Manager.CallOpts, validationID)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64,uint64) validator)
func (_ACP99Manager *ACP99ManagerCallerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ACP99Manager.Contract.GetValidator(&_ACP99Manager.CallOpts, validationID)
}

// L1TotalWeight is a free data retrieval call binding the contract method 0xbb0b1938.
//
// Solidity: function l1TotalWeight() view returns(uint64 weight)
func (_ACP99Manager *ACP99ManagerCaller) L1TotalWeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ACP99Manager.contract.Call(opts, &out, "l1TotalWeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// L1TotalWeight is a free data retrieval call binding the contract method 0xbb0b1938.
//
// Solidity: function l1TotalWeight() view returns(uint64 weight)
func (_ACP99Manager *ACP99ManagerSession) L1TotalWeight() (uint64, error) {
	return _ACP99Manager.Contract.L1TotalWeight(&_ACP99Manager.CallOpts)
}

// L1TotalWeight is a free data retrieval call binding the contract method 0xbb0b1938.
//
// Solidity: function l1TotalWeight() view returns(uint64 weight)
func (_ACP99Manager *ACP99ManagerCallerSession) L1TotalWeight() (uint64, error) {
	return _ACP99Manager.Contract.L1TotalWeight(&_ACP99Manager.CallOpts)
}

// SubnetID is a free data retrieval call binding the contract method 0x5dc1f535.
//
// Solidity: function subnetID() view returns(bytes32 id)
func (_ACP99Manager *ACP99ManagerCaller) SubnetID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ACP99Manager.contract.Call(opts, &out, "subnetID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SubnetID is a free data retrieval call binding the contract method 0x5dc1f535.
//
// Solidity: function subnetID() view returns(bytes32 id)
func (_ACP99Manager *ACP99ManagerSession) SubnetID() ([32]byte, error) {
	return _ACP99Manager.Contract.SubnetID(&_ACP99Manager.CallOpts)
}

// SubnetID is a free data retrieval call binding the contract method 0x5dc1f535.
//
// Solidity: function subnetID() view returns(bytes32 id)
func (_ACP99Manager *ACP99ManagerCallerSession) SubnetID() ([32]byte, error) {
	return _ACP99Manager.Contract.SubnetID(&_ACP99Manager.CallOpts)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32 validationID)
func (_ACP99Manager *ACP99ManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32 validationID)
func (_ACP99Manager *ACP99ManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.Contract.CompleteValidatorRegistration(&_ACP99Manager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32 validationID)
func (_ACP99Manager *ACP99ManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.Contract.CompleteValidatorRegistration(&_ACP99Manager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32 validationID)
func (_ACP99Manager *ACP99ManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32 validationID)
func (_ACP99Manager *ACP99ManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.Contract.CompleteValidatorRemoval(&_ACP99Manager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32 validationID)
func (_ACP99Manager *ACP99ManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.Contract.CompleteValidatorRemoval(&_ACP99Manager.TransactOpts, messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32 validationID, uint64 nonce)
func (_ACP99Manager *ACP99ManagerTransactor) CompleteValidatorWeightUpdate(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.contract.Transact(opts, "completeValidatorWeightUpdate", messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32 validationID, uint64 nonce)
func (_ACP99Manager *ACP99ManagerSession) CompleteValidatorWeightUpdate(messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.Contract.CompleteValidatorWeightUpdate(&_ACP99Manager.TransactOpts, messageIndex)
}

// CompleteValidatorWeightUpdate is a paid mutator transaction binding the contract method 0xce161f14.
//
// Solidity: function completeValidatorWeightUpdate(uint32 messageIndex) returns(bytes32 validationID, uint64 nonce)
func (_ACP99Manager *ACP99ManagerTransactorSession) CompleteValidatorWeightUpdate(messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.Contract.CompleteValidatorWeightUpdate(&_ACP99Manager.TransactOpts, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_ACP99Manager *ACP99ManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.contract.Transact(opts, "initializeValidatorSet", conversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_ACP99Manager *ACP99ManagerSession) InitializeValidatorSet(conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.Contract.InitializeValidatorSet(&_ACP99Manager.TransactOpts, conversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_ACP99Manager *ACP99ManagerTransactorSession) InitializeValidatorSet(conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ACP99Manager.Contract.InitializeValidatorSet(&_ACP99Manager.TransactOpts, conversionData, messageIndex)
}

// ACP99ManagerCompletedValidatorRegistrationIterator is returned from FilterCompletedValidatorRegistration and is used to iterate over the raw logs and unpacked data for CompletedValidatorRegistration events raised by the ACP99Manager contract.
type ACP99ManagerCompletedValidatorRegistrationIterator struct {
	Event *ACP99ManagerCompletedValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *ACP99ManagerCompletedValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACP99ManagerCompletedValidatorRegistration)
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
		it.Event = new(ACP99ManagerCompletedValidatorRegistration)
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
func (it *ACP99ManagerCompletedValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACP99ManagerCompletedValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACP99ManagerCompletedValidatorRegistration represents a CompletedValidatorRegistration event raised by the ACP99Manager contract.
type ACP99ManagerCompletedValidatorRegistration struct {
	ValidationID [32]byte
	Weight       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedValidatorRegistration is a free log retrieval operation binding the contract event 0x967ae87813a3b5f201dd9bcba778d457176eafe6f41facee1c718091d3952d06.
//
// Solidity: event CompletedValidatorRegistration(bytes32 indexed validationID, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) FilterCompletedValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte) (*ACP99ManagerCompletedValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.FilterLogs(opts, "CompletedValidatorRegistration", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ACP99ManagerCompletedValidatorRegistrationIterator{contract: _ACP99Manager.contract, event: "CompletedValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchCompletedValidatorRegistration is a free log subscription operation binding the contract event 0x967ae87813a3b5f201dd9bcba778d457176eafe6f41facee1c718091d3952d06.
//
// Solidity: event CompletedValidatorRegistration(bytes32 indexed validationID, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) WatchCompletedValidatorRegistration(opts *bind.WatchOpts, sink chan<- *ACP99ManagerCompletedValidatorRegistration, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.WatchLogs(opts, "CompletedValidatorRegistration", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACP99ManagerCompletedValidatorRegistration)
				if err := _ACP99Manager.contract.UnpackLog(event, "CompletedValidatorRegistration", log); err != nil {
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

// ParseCompletedValidatorRegistration is a log parse operation binding the contract event 0x967ae87813a3b5f201dd9bcba778d457176eafe6f41facee1c718091d3952d06.
//
// Solidity: event CompletedValidatorRegistration(bytes32 indexed validationID, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) ParseCompletedValidatorRegistration(log types.Log) (*ACP99ManagerCompletedValidatorRegistration, error) {
	event := new(ACP99ManagerCompletedValidatorRegistration)
	if err := _ACP99Manager.contract.UnpackLog(event, "CompletedValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACP99ManagerCompletedValidatorRemovalIterator is returned from FilterCompletedValidatorRemoval and is used to iterate over the raw logs and unpacked data for CompletedValidatorRemoval events raised by the ACP99Manager contract.
type ACP99ManagerCompletedValidatorRemovalIterator struct {
	Event *ACP99ManagerCompletedValidatorRemoval // Event containing the contract specifics and raw log

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
func (it *ACP99ManagerCompletedValidatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACP99ManagerCompletedValidatorRemoval)
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
		it.Event = new(ACP99ManagerCompletedValidatorRemoval)
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
func (it *ACP99ManagerCompletedValidatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACP99ManagerCompletedValidatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACP99ManagerCompletedValidatorRemoval represents a CompletedValidatorRemoval event raised by the ACP99Manager contract.
type ACP99ManagerCompletedValidatorRemoval struct {
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedValidatorRemoval is a free log retrieval operation binding the contract event 0xafaccef7080649a725bc30a35359a257a4a27225be352875c80bdf6b5f04080c.
//
// Solidity: event CompletedValidatorRemoval(bytes32 indexed validationID)
func (_ACP99Manager *ACP99ManagerFilterer) FilterCompletedValidatorRemoval(opts *bind.FilterOpts, validationID [][32]byte) (*ACP99ManagerCompletedValidatorRemovalIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.FilterLogs(opts, "CompletedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ACP99ManagerCompletedValidatorRemovalIterator{contract: _ACP99Manager.contract, event: "CompletedValidatorRemoval", logs: logs, sub: sub}, nil
}

// WatchCompletedValidatorRemoval is a free log subscription operation binding the contract event 0xafaccef7080649a725bc30a35359a257a4a27225be352875c80bdf6b5f04080c.
//
// Solidity: event CompletedValidatorRemoval(bytes32 indexed validationID)
func (_ACP99Manager *ACP99ManagerFilterer) WatchCompletedValidatorRemoval(opts *bind.WatchOpts, sink chan<- *ACP99ManagerCompletedValidatorRemoval, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.WatchLogs(opts, "CompletedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACP99ManagerCompletedValidatorRemoval)
				if err := _ACP99Manager.contract.UnpackLog(event, "CompletedValidatorRemoval", log); err != nil {
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

// ParseCompletedValidatorRemoval is a log parse operation binding the contract event 0xafaccef7080649a725bc30a35359a257a4a27225be352875c80bdf6b5f04080c.
//
// Solidity: event CompletedValidatorRemoval(bytes32 indexed validationID)
func (_ACP99Manager *ACP99ManagerFilterer) ParseCompletedValidatorRemoval(log types.Log) (*ACP99ManagerCompletedValidatorRemoval, error) {
	event := new(ACP99ManagerCompletedValidatorRemoval)
	if err := _ACP99Manager.contract.UnpackLog(event, "CompletedValidatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACP99ManagerCompletedValidatorWeightUpdateIterator is returned from FilterCompletedValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for CompletedValidatorWeightUpdate events raised by the ACP99Manager contract.
type ACP99ManagerCompletedValidatorWeightUpdateIterator struct {
	Event *ACP99ManagerCompletedValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *ACP99ManagerCompletedValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACP99ManagerCompletedValidatorWeightUpdate)
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
		it.Event = new(ACP99ManagerCompletedValidatorWeightUpdate)
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
func (it *ACP99ManagerCompletedValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACP99ManagerCompletedValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACP99ManagerCompletedValidatorWeightUpdate represents a CompletedValidatorWeightUpdate event raised by the ACP99Manager contract.
type ACP99ManagerCompletedValidatorWeightUpdate struct {
	ValidationID [32]byte
	Nonce        uint64
	Weight       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedValidatorWeightUpdate is a free log retrieval operation binding the contract event 0xc917996591802ecedcfced71321d4bb5320f7dfbacf5477dffe1dbf8b8839ff9.
//
// Solidity: event CompletedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) FilterCompletedValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte) (*ACP99ManagerCompletedValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.FilterLogs(opts, "CompletedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ACP99ManagerCompletedValidatorWeightUpdateIterator{contract: _ACP99Manager.contract, event: "CompletedValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchCompletedValidatorWeightUpdate is a free log subscription operation binding the contract event 0xc917996591802ecedcfced71321d4bb5320f7dfbacf5477dffe1dbf8b8839ff9.
//
// Solidity: event CompletedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) WatchCompletedValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *ACP99ManagerCompletedValidatorWeightUpdate, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.WatchLogs(opts, "CompletedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACP99ManagerCompletedValidatorWeightUpdate)
				if err := _ACP99Manager.contract.UnpackLog(event, "CompletedValidatorWeightUpdate", log); err != nil {
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

// ParseCompletedValidatorWeightUpdate is a log parse operation binding the contract event 0xc917996591802ecedcfced71321d4bb5320f7dfbacf5477dffe1dbf8b8839ff9.
//
// Solidity: event CompletedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) ParseCompletedValidatorWeightUpdate(log types.Log) (*ACP99ManagerCompletedValidatorWeightUpdate, error) {
	event := new(ACP99ManagerCompletedValidatorWeightUpdate)
	if err := _ACP99Manager.contract.UnpackLog(event, "CompletedValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACP99ManagerInitiatedValidatorRegistrationIterator is returned from FilterInitiatedValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedValidatorRegistration events raised by the ACP99Manager contract.
type ACP99ManagerInitiatedValidatorRegistrationIterator struct {
	Event *ACP99ManagerInitiatedValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *ACP99ManagerInitiatedValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACP99ManagerInitiatedValidatorRegistration)
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
		it.Event = new(ACP99ManagerInitiatedValidatorRegistration)
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
func (it *ACP99ManagerInitiatedValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACP99ManagerInitiatedValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACP99ManagerInitiatedValidatorRegistration represents a InitiatedValidatorRegistration event raised by the ACP99Manager contract.
type ACP99ManagerInitiatedValidatorRegistration struct {
	ValidationID          [32]byte
	NodeID                [20]byte
	RegistrationMessageID [32]byte
	RegistrationExpiry    uint64
	Weight                uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterInitiatedValidatorRegistration is a free log retrieval operation binding the contract event 0x5881be437bdcb008bfa5f20e32d3e335ccf8ab90ef2818852a251625260af35d.
//
// Solidity: event InitiatedValidatorRegistration(bytes32 indexed validationID, bytes20 indexed nodeID, bytes32 registrationMessageID, uint64 registrationExpiry, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) FilterInitiatedValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][20]byte) (*ACP99ManagerInitiatedValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.FilterLogs(opts, "InitiatedValidatorRegistration", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ACP99ManagerInitiatedValidatorRegistrationIterator{contract: _ACP99Manager.contract, event: "InitiatedValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedValidatorRegistration is a free log subscription operation binding the contract event 0x5881be437bdcb008bfa5f20e32d3e335ccf8ab90ef2818852a251625260af35d.
//
// Solidity: event InitiatedValidatorRegistration(bytes32 indexed validationID, bytes20 indexed nodeID, bytes32 registrationMessageID, uint64 registrationExpiry, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) WatchInitiatedValidatorRegistration(opts *bind.WatchOpts, sink chan<- *ACP99ManagerInitiatedValidatorRegistration, validationID [][32]byte, nodeID [][20]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.WatchLogs(opts, "InitiatedValidatorRegistration", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACP99ManagerInitiatedValidatorRegistration)
				if err := _ACP99Manager.contract.UnpackLog(event, "InitiatedValidatorRegistration", log); err != nil {
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

// ParseInitiatedValidatorRegistration is a log parse operation binding the contract event 0x5881be437bdcb008bfa5f20e32d3e335ccf8ab90ef2818852a251625260af35d.
//
// Solidity: event InitiatedValidatorRegistration(bytes32 indexed validationID, bytes20 indexed nodeID, bytes32 registrationMessageID, uint64 registrationExpiry, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) ParseInitiatedValidatorRegistration(log types.Log) (*ACP99ManagerInitiatedValidatorRegistration, error) {
	event := new(ACP99ManagerInitiatedValidatorRegistration)
	if err := _ACP99Manager.contract.UnpackLog(event, "InitiatedValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACP99ManagerInitiatedValidatorRemovalIterator is returned from FilterInitiatedValidatorRemoval and is used to iterate over the raw logs and unpacked data for InitiatedValidatorRemoval events raised by the ACP99Manager contract.
type ACP99ManagerInitiatedValidatorRemovalIterator struct {
	Event *ACP99ManagerInitiatedValidatorRemoval // Event containing the contract specifics and raw log

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
func (it *ACP99ManagerInitiatedValidatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACP99ManagerInitiatedValidatorRemoval)
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
		it.Event = new(ACP99ManagerInitiatedValidatorRemoval)
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
func (it *ACP99ManagerInitiatedValidatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACP99ManagerInitiatedValidatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACP99ManagerInitiatedValidatorRemoval represents a InitiatedValidatorRemoval event raised by the ACP99Manager contract.
type ACP99ManagerInitiatedValidatorRemoval struct {
	ValidationID             [32]byte
	ValidatorWeightMessageID [32]byte
	Weight                   uint64
	EndTime                  uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterInitiatedValidatorRemoval is a free log retrieval operation binding the contract event 0xbae388a94e7f18411fe57098f12f418b8e1a8273e0532a90188a3a059b897273.
//
// Solidity: event InitiatedValidatorRemoval(bytes32 indexed validationID, bytes32 validatorWeightMessageID, uint64 weight, uint64 endTime)
func (_ACP99Manager *ACP99ManagerFilterer) FilterInitiatedValidatorRemoval(opts *bind.FilterOpts, validationID [][32]byte) (*ACP99ManagerInitiatedValidatorRemovalIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.FilterLogs(opts, "InitiatedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ACP99ManagerInitiatedValidatorRemovalIterator{contract: _ACP99Manager.contract, event: "InitiatedValidatorRemoval", logs: logs, sub: sub}, nil
}

// WatchInitiatedValidatorRemoval is a free log subscription operation binding the contract event 0xbae388a94e7f18411fe57098f12f418b8e1a8273e0532a90188a3a059b897273.
//
// Solidity: event InitiatedValidatorRemoval(bytes32 indexed validationID, bytes32 validatorWeightMessageID, uint64 weight, uint64 endTime)
func (_ACP99Manager *ACP99ManagerFilterer) WatchInitiatedValidatorRemoval(opts *bind.WatchOpts, sink chan<- *ACP99ManagerInitiatedValidatorRemoval, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.WatchLogs(opts, "InitiatedValidatorRemoval", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACP99ManagerInitiatedValidatorRemoval)
				if err := _ACP99Manager.contract.UnpackLog(event, "InitiatedValidatorRemoval", log); err != nil {
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

// ParseInitiatedValidatorRemoval is a log parse operation binding the contract event 0xbae388a94e7f18411fe57098f12f418b8e1a8273e0532a90188a3a059b897273.
//
// Solidity: event InitiatedValidatorRemoval(bytes32 indexed validationID, bytes32 validatorWeightMessageID, uint64 weight, uint64 endTime)
func (_ACP99Manager *ACP99ManagerFilterer) ParseInitiatedValidatorRemoval(log types.Log) (*ACP99ManagerInitiatedValidatorRemoval, error) {
	event := new(ACP99ManagerInitiatedValidatorRemoval)
	if err := _ACP99Manager.contract.UnpackLog(event, "InitiatedValidatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACP99ManagerInitiatedValidatorWeightUpdateIterator is returned from FilterInitiatedValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for InitiatedValidatorWeightUpdate events raised by the ACP99Manager contract.
type ACP99ManagerInitiatedValidatorWeightUpdateIterator struct {
	Event *ACP99ManagerInitiatedValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *ACP99ManagerInitiatedValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACP99ManagerInitiatedValidatorWeightUpdate)
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
		it.Event = new(ACP99ManagerInitiatedValidatorWeightUpdate)
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
func (it *ACP99ManagerInitiatedValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACP99ManagerInitiatedValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACP99ManagerInitiatedValidatorWeightUpdate represents a InitiatedValidatorWeightUpdate event raised by the ACP99Manager contract.
type ACP99ManagerInitiatedValidatorWeightUpdate struct {
	ValidationID          [32]byte
	Nonce                 uint64
	WeightUpdateMessageID [32]byte
	Weight                uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterInitiatedValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x6e350dd49b060d87f297206fd309234ed43156d890ced0f139ecf704310481d3.
//
// Solidity: event InitiatedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, bytes32 weightUpdateMessageID, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) FilterInitiatedValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte) (*ACP99ManagerInitiatedValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.FilterLogs(opts, "InitiatedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ACP99ManagerInitiatedValidatorWeightUpdateIterator{contract: _ACP99Manager.contract, event: "InitiatedValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchInitiatedValidatorWeightUpdate is a free log subscription operation binding the contract event 0x6e350dd49b060d87f297206fd309234ed43156d890ced0f139ecf704310481d3.
//
// Solidity: event InitiatedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, bytes32 weightUpdateMessageID, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) WatchInitiatedValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *ACP99ManagerInitiatedValidatorWeightUpdate, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.WatchLogs(opts, "InitiatedValidatorWeightUpdate", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACP99ManagerInitiatedValidatorWeightUpdate)
				if err := _ACP99Manager.contract.UnpackLog(event, "InitiatedValidatorWeightUpdate", log); err != nil {
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

// ParseInitiatedValidatorWeightUpdate is a log parse operation binding the contract event 0x6e350dd49b060d87f297206fd309234ed43156d890ced0f139ecf704310481d3.
//
// Solidity: event InitiatedValidatorWeightUpdate(bytes32 indexed validationID, uint64 nonce, bytes32 weightUpdateMessageID, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) ParseInitiatedValidatorWeightUpdate(log types.Log) (*ACP99ManagerInitiatedValidatorWeightUpdate, error) {
	event := new(ACP99ManagerInitiatedValidatorWeightUpdate)
	if err := _ACP99Manager.contract.UnpackLog(event, "InitiatedValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ACP99ManagerRegisteredInitialValidatorIterator is returned from FilterRegisteredInitialValidator and is used to iterate over the raw logs and unpacked data for RegisteredInitialValidator events raised by the ACP99Manager contract.
type ACP99ManagerRegisteredInitialValidatorIterator struct {
	Event *ACP99ManagerRegisteredInitialValidator // Event containing the contract specifics and raw log

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
func (it *ACP99ManagerRegisteredInitialValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ACP99ManagerRegisteredInitialValidator)
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
		it.Event = new(ACP99ManagerRegisteredInitialValidator)
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
func (it *ACP99ManagerRegisteredInitialValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ACP99ManagerRegisteredInitialValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ACP99ManagerRegisteredInitialValidator represents a RegisteredInitialValidator event raised by the ACP99Manager contract.
type ACP99ManagerRegisteredInitialValidator struct {
	ValidationID [32]byte
	NodeID       [20]byte
	SubnetID     [32]byte
	Weight       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisteredInitialValidator is a free log retrieval operation binding the contract event 0x070e0f6540afd612d613b71036a6d45f8c5fa97770a7e2d55663d75474ce0c0d.
//
// Solidity: event RegisteredInitialValidator(bytes32 indexed validationID, bytes20 indexed nodeID, bytes32 indexed subnetID, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) FilterRegisteredInitialValidator(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][20]byte, subnetID [][32]byte) (*ACP99ManagerRegisteredInitialValidatorIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var subnetIDRule []interface{}
	for _, subnetIDItem := range subnetID {
		subnetIDRule = append(subnetIDRule, subnetIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.FilterLogs(opts, "RegisteredInitialValidator", validationIDRule, nodeIDRule, subnetIDRule)
	if err != nil {
		return nil, err
	}
	return &ACP99ManagerRegisteredInitialValidatorIterator{contract: _ACP99Manager.contract, event: "RegisteredInitialValidator", logs: logs, sub: sub}, nil
}

// WatchRegisteredInitialValidator is a free log subscription operation binding the contract event 0x070e0f6540afd612d613b71036a6d45f8c5fa97770a7e2d55663d75474ce0c0d.
//
// Solidity: event RegisteredInitialValidator(bytes32 indexed validationID, bytes20 indexed nodeID, bytes32 indexed subnetID, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) WatchRegisteredInitialValidator(opts *bind.WatchOpts, sink chan<- *ACP99ManagerRegisteredInitialValidator, validationID [][32]byte, nodeID [][20]byte, subnetID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var subnetIDRule []interface{}
	for _, subnetIDItem := range subnetID {
		subnetIDRule = append(subnetIDRule, subnetIDItem)
	}

	logs, sub, err := _ACP99Manager.contract.WatchLogs(opts, "RegisteredInitialValidator", validationIDRule, nodeIDRule, subnetIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ACP99ManagerRegisteredInitialValidator)
				if err := _ACP99Manager.contract.UnpackLog(event, "RegisteredInitialValidator", log); err != nil {
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

// ParseRegisteredInitialValidator is a log parse operation binding the contract event 0x070e0f6540afd612d613b71036a6d45f8c5fa97770a7e2d55663d75474ce0c0d.
//
// Solidity: event RegisteredInitialValidator(bytes32 indexed validationID, bytes20 indexed nodeID, bytes32 indexed subnetID, uint64 weight)
func (_ACP99Manager *ACP99ManagerFilterer) ParseRegisteredInitialValidator(log types.Log) (*ACP99ManagerRegisteredInitialValidator, error) {
	event := new(ACP99ManagerRegisteredInitialValidator)
	if err := _ACP99Manager.contract.UnpackLog(event, "RegisteredInitialValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
