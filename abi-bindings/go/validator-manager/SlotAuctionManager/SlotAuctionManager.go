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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"becomeValidator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"stopValidating\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"validationIDs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zilch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001805460ff60a01b19169055348015601b575f80fd5b50604051610cb8380380610cb88339810160408190526038916081565b5f80546001600160a01b039384166001600160a01b0319918216179091556001805492909316911617905560ad565b80516001600160a01b0381168114607c575f80fd5b919050565b5f80604083850312156091575f80fd5b6098836067565b915060a4602084016067565b90509250929050565b610bfe806100ba5f395ff3fe608060405234801561000f575f80fd5b506004361061009b575f3560e01c8063a476f67511610063578063a476f6751461013d578063ae9483e014610150578063bebf2eee14610163578063f76236c11461016c578063fa52c7d814610194575f80fd5b806305af82561461009f5780630661ef7e146100b45780630f43a677146100c757806336339388146100f257806337748bdb1461011c575b5f80fd5b6100b26100ad366004610788565b6101b3565b005b6100b26100c23660046107a8565b610228565b6001546100db90600160a01b900460ff1681565b60405160ff90911681526020015b60405180910390f35b5f54610104906001600160a01b031681565b6040516001600160a01b0390911681526020016100e9565b61012f61012a366004610974565b610444565b6040519081526020016100e9565b6100b261014b3660046107a8565b6106b0565b600154610104906001600160a01b031681565b61012f60045481565b61010461017a3660046107a8565b60036020525f90815260409020546001600160a01b031681565b61012f6101a2366004610a1a565b60026020525f908152604090205481565b60015460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af1158015610200573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102249190610a33565b5050565b600154600160a01b900460ff166102865760405162461bcd60e51b815260206004820152601760248201527f43757272656e746c79206e6f2076616c696461746f727300000000000000000060448201526064015b60405180910390fd5b335f9081526002602052604090205481146102f55760405162461bcd60e51b815260206004820152602960248201527f496e76616c69642073656e6465722c2073656e646572206973206e6f742061206044820152683b30b634b230ba37b960b91b606482015260840161027d565b335f52600260205260018054600160a01b900460ff1690601461031783610a5e565b825460ff9182166101009390930a928302919092021990911617905550600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b158015610377575f80fd5b505af1158015610389573d5f803e3d5ffd5b50505f5460405163a9059cbb60e01b8152336004820152600a60248201526001600160a01b03909116925063a9059cbb91506044016020604051808303815f875af11580156103da573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103fe9190610a79565b6104415760405162461bcd60e51b8152602060048201526014602482015273119d5b991cc819985a5b1959081d1bc81cd95b9960621b604482015260640161027d565b50565b6001545f90600a600160a01b90910460ff16106104a35760405162461bcd60e51b815260206004820152601b60248201527f4d61782076616c696461746f72206c696d697420726561636865640000000000604482015260640161027d565b335f90815260026020526040902054156104ff5760405162461bcd60e51b815260206004820152601d60248201527f53656e64657220697320616c726561647920612076616c696461746f72000000604482015260640161027d565b60018054600160a01b900460ff1690601461051983610a98565b825460ff9182166101009390930a9283029190920219909116179055505f546040516323b872dd60e01b8152336004820152306024820152600a60448201526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015610589573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105ad9190610a79565b15156001146105fe5760405162461bcd60e51b815260206004820152601960248201527f546f6b656e73206661696c656420746f207472616e7366657200000000000000604482015260640161027d565b600154604051634e5bb12760e11b81525f916001600160a01b031690639cb7624e90610637908990899089908990606490600401610b5f565b6020604051808303815f875af1158015610653573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106779190610a33565b335f8181526002602090815260408083208590558483526003909152902080546001600160a01b03191690911790559695505050505050565b5f818152600360205260409020546001600160a01b0316156107145760405162461bcd60e51b815260206004820152601f60248201527f56616c696461746f72206e6f7420696e697469616c2056616c696461746f7200604482015260640161027d565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b158015610757575f80fd5b505af1158015610769573d5f803e3d5ffd5b5050505050565b803563ffffffff81168114610783575f80fd5b919050565b5f60208284031215610798575f80fd5b6107a182610770565b9392505050565b5f602082840312156107b8575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff811182821017156107f6576107f66107bf565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610825576108256107bf565b604052919050565b5f82601f83011261083c575f80fd5b813567ffffffffffffffff811115610856576108566107bf565b610869601f8201601f19166020016107fc565b81815284602083860101111561087d575f80fd5b816020850160208301375f918101602001919091529392505050565b80356001600160a01b0381168114610783575f80fd5b5f604082840312156108bf575f80fd5b6108c76107d3565b90506108d282610770565b815260208083013567ffffffffffffffff808211156108ef575f80fd5b818501915085601f830112610902575f80fd5b813581811115610914576109146107bf565b8060051b91506109258483016107fc565b818152918301840191848101908884111561093e575f80fd5b938501935b838510156109635761095485610899565b82529385019390850190610943565b808688015250505050505092915050565b5f805f8060808587031215610987575f80fd5b843567ffffffffffffffff8082111561099e575f80fd5b6109aa8883890161082d565b955060208701359150808211156109bf575f80fd5b6109cb8883890161082d565b945060408701359150808211156109e0575f80fd5b6109ec888389016108af565b93506060870135915080821115610a01575f80fd5b50610a0e878288016108af565b91505092959194509250565b5f60208284031215610a2a575f80fd5b6107a182610899565b5f60208284031215610a43575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b5f60ff821680610a7057610a70610a4a565b5f190192915050565b5f60208284031215610a89575f80fd5b815180151581146107a1575f80fd5b5f60ff821660ff8103610aad57610aad610a4a565b60010192915050565b5f81518084525f5b81811015610ada57602081850181015186830182015201610abe565b505f602082860101526020601f19601f83011685010191505092915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015610b545784516001600160a01b03168252938301936001929092019190830190610b2b565b509695505050505050565b60a081525f610b7160a0830188610ab6565b8281036020840152610b838188610ab6565b90508281036040840152610b978187610af9565b90508281036060840152610bab8186610af9565b91505067ffffffffffffffff83166080830152969550505050505056fea26469706673582212208d190e73c4d40d339d2b792b43e5c012fb2af39442ee5edb333b0b42a75d352064736f6c63430008190033",
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

// ValidationIDs is a free data retrieval call binding the contract method 0xf76236c1.
//
// Solidity: function validationIDs(bytes32 validationID) view returns(address validator)
func (_SlotAuctionManager *SlotAuctionManagerCaller) ValidationIDs(opts *bind.CallOpts, validationID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "validationIDs", validationID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidationIDs is a free data retrieval call binding the contract method 0xf76236c1.
//
// Solidity: function validationIDs(bytes32 validationID) view returns(address validator)
func (_SlotAuctionManager *SlotAuctionManagerSession) ValidationIDs(validationID [32]byte) (common.Address, error) {
	return _SlotAuctionManager.Contract.ValidationIDs(&_SlotAuctionManager.CallOpts, validationID)
}

// ValidationIDs is a free data retrieval call binding the contract method 0xf76236c1.
//
// Solidity: function validationIDs(bytes32 validationID) view returns(address validator)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) ValidationIDs(validationID [32]byte) (common.Address, error) {
	return _SlotAuctionManager.Contract.ValidationIDs(&_SlotAuctionManager.CallOpts, validationID)
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

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address validator) view returns(bytes32 validationID)
func (_SlotAuctionManager *SlotAuctionManagerCaller) Validators(opts *bind.CallOpts, validator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "validators", validator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address validator) view returns(bytes32 validationID)
func (_SlotAuctionManager *SlotAuctionManagerSession) Validators(validator common.Address) ([32]byte, error) {
	return _SlotAuctionManager.Contract.Validators(&_SlotAuctionManager.CallOpts, validator)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address validator) view returns(bytes32 validationID)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) Validators(validator common.Address) ([32]byte, error) {
	return _SlotAuctionManager.Contract.Validators(&_SlotAuctionManager.CallOpts, validator)
}

// Zilch is a free data retrieval call binding the contract method 0xbebf2eee.
//
// Solidity: function zilch() view returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerCaller) Zilch(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "zilch")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Zilch is a free data retrieval call binding the contract method 0xbebf2eee.
//
// Solidity: function zilch() view returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerSession) Zilch() ([32]byte, error) {
	return _SlotAuctionManager.Contract.Zilch(&_SlotAuctionManager.CallOpts)
}

// Zilch is a free data retrieval call binding the contract method 0xbebf2eee.
//
// Solidity: function zilch() view returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) Zilch() ([32]byte, error) {
	return _SlotAuctionManager.Contract.Zilch(&_SlotAuctionManager.CallOpts)
}

// BecomeValidator is a paid mutator transaction binding the contract method 0x37748bdb.
//
// Solidity: function becomeValidator(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactor) BecomeValidator(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "becomeValidator", nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// BecomeValidator is a paid mutator transaction binding the contract method 0x37748bdb.
//
// Solidity: function becomeValidator(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerSession) BecomeValidator(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.BecomeValidator(&_SlotAuctionManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// BecomeValidator is a paid mutator transaction binding the contract method 0x37748bdb.
//
// Solidity: function becomeValidator(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) BecomeValidator(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.BecomeValidator(&_SlotAuctionManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
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

// StopValidating is a paid mutator transaction binding the contract method 0x0661ef7e.
//
// Solidity: function stopValidating(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) StopValidating(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "stopValidating", validationID)
}

// StopValidating is a paid mutator transaction binding the contract method 0x0661ef7e.
//
// Solidity: function stopValidating(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) StopValidating(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.StopValidating(&_SlotAuctionManager.TransactOpts, validationID)
}

// StopValidating is a paid mutator transaction binding the contract method 0x0661ef7e.
//
// Solidity: function stopValidating(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) StopValidating(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.StopValidating(&_SlotAuctionManager.TransactOpts, validationID)
}
