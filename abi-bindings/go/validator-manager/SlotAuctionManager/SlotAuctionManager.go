// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package slotauctionmanager

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

// PChainOwner is an auto generated low-level Go binding around an user-defined struct.
type PChainOwner struct {
	Threshold uint32
	Addresses []common.Address
}

// SlotAuctionManagerMetaData contains all meta data concerning the SlotAuctionManager contract.
var SlotAuctionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TemporaryID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"initiateValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040525f6002556003805460ff19169055348015601c575f80fd5b50604051610b4d380380610b4d8339810160408190526039916082565b5f80546001600160a01b039384166001600160a01b0319918216179091556001805492909316911617905560ae565b80516001600160a01b0381168114607d575f80fd5b919050565b5f80604083850312156092575f80fd5b6099836068565b915060a5602084016068565b90509250929050565b610a92806100bb5f395ff3fe608060405234801561000f575f80fd5b5060043610610090575f3560e01c80637d89c96d116100635780637d89c96d1461010e578063a3a65e4814610121578063a476f67514610134578063ae9483e014610147578063dc1113eb1461015a575f80fd5b806305af8256146100945780630f43a677146100a957806336339388146100cd5780636fb191f0146100f7575b5f80fd5b6100a76100a236600461063b565b61016d565b005b6003546100b69060ff1681565b60405160ff90911681526020015b60405180910390f35b5f546100df906001600160a01b031681565b6040516001600160a01b0390911681526020016100c4565b61010060025481565b6040519081526020016100c4565b6100a761011c36600461065b565b6101e2565b61010061012f36600461063b565b610378565b6100a761014236600461065b565b6103ef565b6001546100df906001600160a01b031681565b610100610168366004610821565b61044b565b60015460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af11580156101ba573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101de91906108c7565b5050565b60035460ff166102395760405162461bcd60e51b815260206004820152601760248201527f43757272656e746c79206e6f2076616c696461746f727300000000000000000060448201526064015b60405180910390fd5b6003805460ff16905f61024b836108f2565b825460ff9182166101009390930a928302919092021990911617905550600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b1580156102ab575f80fd5b505af11580156102bd573d5f803e3d5ffd5b50505f5460405163a9059cbb60e01b8152336004820152600a60248201526001600160a01b03909116925063a9059cbb91506044016020604051808303815f875af115801561030e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610332919061090d565b6103755760405162461bcd60e51b8152602060048201526014602482015273119d5b991cc819985a5b1959081d1bc81cd95b9960621b6044820152606401610230565b50565b600154604051631474cbc960e31b815263ffffffff831660048201525f916001600160a01b03169063a3a65e48906024016020604051808303815f875af11580156103c5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103e991906108c7565b92915050565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b158015610432575f80fd5b505af1158015610444573d5f803e3d5ffd5b5050505050565b6003545f90600a60ff909116106104a45760405162461bcd60e51b815260206004820152601b60248201527f4d61782076616c696461746f72206c696d6974207265616368656400000000006044820152606401610230565b6003805460ff16905f6104b68361092c565b825460ff9182166101009390930a9283029190920219909116179055505f546040516323b872dd60e01b8152336004820152306024820152600a60448201526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015610526573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061054a919061090d565b151560011461059b5760405162461bcd60e51b815260206004820152601960248201527f546f6b656e73206661696c656420746f207472616e73666572000000000000006044820152606401610230565b600154604051634e5bb12760e11b81525f916001600160a01b031690639cb7624e906105d49089908990899089906064906004016109f3565b6020604051808303815f875af11580156105f0573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061061491906108c7565b60028190559695505050505050565b803563ffffffff81168114610636575f80fd5b919050565b5f6020828403121561064b575f80fd5b61065482610623565b9392505050565b5f6020828403121561066b575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff811182821017156106a9576106a9610672565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156106d8576106d8610672565b604052919050565b5f82601f8301126106ef575f80fd5b813567ffffffffffffffff81111561070957610709610672565b61071c601f8201601f19166020016106af565b818152846020838601011115610730575f80fd5b816020850160208301375f918101602001919091529392505050565b5f6040828403121561075c575f80fd5b610764610686565b905061076f82610623565b815260208083013567ffffffffffffffff8082111561078c575f80fd5b818501915085601f83011261079f575f80fd5b8135818111156107b1576107b1610672565b8060051b91506107c28483016106af565b81815291830184019184810190888411156107db575f80fd5b938501935b8385101561081057843592506001600160a01b0383168314610800575f80fd5b82825293850193908501906107e0565b808688015250505050505092915050565b5f805f8060808587031215610834575f80fd5b843567ffffffffffffffff8082111561084b575f80fd5b610857888389016106e0565b9550602087013591508082111561086c575f80fd5b610878888389016106e0565b9450604087013591508082111561088d575f80fd5b6108998883890161074c565b935060608701359150808211156108ae575f80fd5b506108bb8782880161074c565b91505092959194509250565b5f602082840312156108d7575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b5f60ff821680610904576109046108de565b5f190192915050565b5f6020828403121561091d575f80fd5b81518015158114610654575f80fd5b5f60ff821660ff8103610941576109416108de565b60010192915050565b5f81518084525f5b8181101561096e57602081850181015186830182015201610952565b505f602082860101526020601f19601f83011685010191505092915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b808310156109e85784516001600160a01b031682529383019360019290920191908301906109bf565b509695505050505050565b60a081525f610a0560a083018861094a565b8281036020840152610a17818861094a565b90508281036040840152610a2b818761098d565b90508281036060840152610a3f818661098d565b91505067ffffffffffffffff83166080830152969550505050505056fea2646970667358221220cf9cda9991c45acfba3b17e07ca65cb2d416d7c8a6d527ea638534f41c13e17564736f6c63430008190033",
}

// SlotAuctionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use SlotAuctionManagerMetaData.ABI instead.
var SlotAuctionManagerABI = SlotAuctionManagerMetaData.ABI

// SlotAuctionManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlotAuctionManagerMetaData.Bin instead.
var SlotAuctionManagerBin = SlotAuctionManagerMetaData.Bin

// DeploySlotAuctionManager deploys a new Ethereum contract, binding an instance of SlotAuctionManager to it.
func DeploySlotAuctionManager(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress common.Address, vmAddress common.Address) (common.Address, *types.Transaction, *SlotAuctionManager, error) {
	parsed, err := SlotAuctionManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlotAuctionManagerBin), backend, tokenAddress, vmAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlotAuctionManager{SlotAuctionManagerCaller: SlotAuctionManagerCaller{contract: contract}, SlotAuctionManagerTransactor: SlotAuctionManagerTransactor{contract: contract}, SlotAuctionManagerFilterer: SlotAuctionManagerFilterer{contract: contract}}, nil
}

// SlotAuctionManager is an auto generated Go binding around an Ethereum contract.
type SlotAuctionManager struct {
	SlotAuctionManagerCaller     // Read-only binding to the contract
	SlotAuctionManagerTransactor // Write-only binding to the contract
	SlotAuctionManagerFilterer   // Log filterer for contract events
}

// SlotAuctionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlotAuctionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlotAuctionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlotAuctionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlotAuctionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlotAuctionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlotAuctionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlotAuctionManagerSession struct {
	Contract     *SlotAuctionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SlotAuctionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlotAuctionManagerCallerSession struct {
	Contract *SlotAuctionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SlotAuctionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlotAuctionManagerTransactorSession struct {
	Contract     *SlotAuctionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SlotAuctionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlotAuctionManagerRaw struct {
	Contract *SlotAuctionManager // Generic contract binding to access the raw methods on
}

// SlotAuctionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlotAuctionManagerCallerRaw struct {
	Contract *SlotAuctionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// SlotAuctionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlotAuctionManagerTransactorRaw struct {
	Contract *SlotAuctionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlotAuctionManager creates a new instance of SlotAuctionManager, bound to a specific deployed contract.
func NewSlotAuctionManager(address common.Address, backend bind.ContractBackend) (*SlotAuctionManager, error) {
	contract, err := bindSlotAuctionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlotAuctionManager{SlotAuctionManagerCaller: SlotAuctionManagerCaller{contract: contract}, SlotAuctionManagerTransactor: SlotAuctionManagerTransactor{contract: contract}, SlotAuctionManagerFilterer: SlotAuctionManagerFilterer{contract: contract}}, nil
}

// NewSlotAuctionManagerCaller creates a new read-only instance of SlotAuctionManager, bound to a specific deployed contract.
func NewSlotAuctionManagerCaller(address common.Address, caller bind.ContractCaller) (*SlotAuctionManagerCaller, error) {
	contract, err := bindSlotAuctionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlotAuctionManagerCaller{contract: contract}, nil
}

// NewSlotAuctionManagerTransactor creates a new write-only instance of SlotAuctionManager, bound to a specific deployed contract.
func NewSlotAuctionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*SlotAuctionManagerTransactor, error) {
	contract, err := bindSlotAuctionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlotAuctionManagerTransactor{contract: contract}, nil
}

// NewSlotAuctionManagerFilterer creates a new log filterer instance of SlotAuctionManager, bound to a specific deployed contract.
func NewSlotAuctionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*SlotAuctionManagerFilterer, error) {
	contract, err := bindSlotAuctionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlotAuctionManagerFilterer{contract: contract}, nil
}

// bindSlotAuctionManager binds a generic wrapper to an already deployed contract.
func bindSlotAuctionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlotAuctionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlotAuctionManager *SlotAuctionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlotAuctionManager.Contract.SlotAuctionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlotAuctionManager *SlotAuctionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.SlotAuctionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlotAuctionManager *SlotAuctionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.SlotAuctionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlotAuctionManager *SlotAuctionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlotAuctionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlotAuctionManager *SlotAuctionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlotAuctionManager *SlotAuctionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.contract.Transact(opts, method, params...)
}

// TOKENCONTRACT is a free data retrieval call binding the contract method 0x36339388.
//
// Solidity: function TOKEN_CONTRACT() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerCaller) TOKENCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "TOKEN_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENCONTRACT is a free data retrieval call binding the contract method 0x36339388.
//
// Solidity: function TOKEN_CONTRACT() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerSession) TOKENCONTRACT() (common.Address, error) {
	return _SlotAuctionManager.Contract.TOKENCONTRACT(&_SlotAuctionManager.CallOpts)
}

// TOKENCONTRACT is a free data retrieval call binding the contract method 0x36339388.
//
// Solidity: function TOKEN_CONTRACT() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) TOKENCONTRACT() (common.Address, error) {
	return _SlotAuctionManager.Contract.TOKENCONTRACT(&_SlotAuctionManager.CallOpts)
}

// TemporaryID is a free data retrieval call binding the contract method 0x6fb191f0.
//
// Solidity: function TemporaryID() view returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerCaller) TemporaryID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "TemporaryID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TemporaryID is a free data retrieval call binding the contract method 0x6fb191f0.
//
// Solidity: function TemporaryID() view returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerSession) TemporaryID() ([32]byte, error) {
	return _SlotAuctionManager.Contract.TemporaryID(&_SlotAuctionManager.CallOpts)
}

// TemporaryID is a free data retrieval call binding the contract method 0x6fb191f0.
//
// Solidity: function TemporaryID() view returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) TemporaryID() ([32]byte, error) {
	return _SlotAuctionManager.Contract.TemporaryID(&_SlotAuctionManager.CallOpts)
}

// VALIDATORMANAGER is a free data retrieval call binding the contract method 0xae9483e0.
//
// Solidity: function VALIDATOR_MANAGER() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerCaller) VALIDATORMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "VALIDATOR_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORMANAGER is a free data retrieval call binding the contract method 0xae9483e0.
//
// Solidity: function VALIDATOR_MANAGER() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerSession) VALIDATORMANAGER() (common.Address, error) {
	return _SlotAuctionManager.Contract.VALIDATORMANAGER(&_SlotAuctionManager.CallOpts)
}

// VALIDATORMANAGER is a free data retrieval call binding the contract method 0xae9483e0.
//
// Solidity: function VALIDATOR_MANAGER() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) VALIDATORMANAGER() (common.Address, error) {
	return _SlotAuctionManager.Contract.VALIDATORMANAGER(&_SlotAuctionManager.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint8)
func (_SlotAuctionManager *SlotAuctionManagerCaller) ValidatorCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint8)
func (_SlotAuctionManager *SlotAuctionManagerSession) ValidatorCount() (uint8, error) {
	return _SlotAuctionManager.Contract.ValidatorCount(&_SlotAuctionManager.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint8)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) ValidatorCount() (uint8, error) {
	return _SlotAuctionManager.Contract.ValidatorCount(&_SlotAuctionManager.CallOpts)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) CompleteRemoveInitialValidator(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "completeRemoveInitialValidator", messageIndex)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) CompleteRemoveInitialValidator(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteRemoveInitialValidator(&_SlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) CompleteRemoveInitialValidator(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteRemoveInitialValidator(&_SlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteValidatorRegistration(&_SlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteValidatorRegistration(&_SlotAuctionManager.TransactOpts, messageIndex)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) InitiateRemoveInitialValidator(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "initiateRemoveInitialValidator", validationID)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) InitiateRemoveInitialValidator(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateRemoveInitialValidator(&_SlotAuctionManager.TransactOpts, validationID)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) InitiateRemoveInitialValidator(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateRemoveInitialValidator(&_SlotAuctionManager.TransactOpts, validationID)
}

// InitiateRemoveValidator is a paid mutator transaction binding the contract method 0x7d89c96d.
//
// Solidity: function initiateRemoveValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) InitiateRemoveValidator(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "initiateRemoveValidator", validationID)
}

// InitiateRemoveValidator is a paid mutator transaction binding the contract method 0x7d89c96d.
//
// Solidity: function initiateRemoveValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) InitiateRemoveValidator(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateRemoveValidator(&_SlotAuctionManager.TransactOpts, validationID)
}

// InitiateRemoveValidator is a paid mutator transaction binding the contract method 0x7d89c96d.
//
// Solidity: function initiateRemoveValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) InitiateRemoveValidator(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateRemoveValidator(&_SlotAuctionManager.TransactOpts, validationID)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0xdc1113eb.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactor) InitiateValidatorRegistration(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "initiateValidatorRegistration", nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0xdc1113eb.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateValidatorRegistration(&_SlotAuctionManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0xdc1113eb.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateValidatorRegistration(&_SlotAuctionManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}
