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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"}],\"name\":\"initiateValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001805460ff60a01b19169055348015601b575f80fd5b50604051610d6c380380610d6c8339810160408190526038916081565b5f80546001600160a01b039384166001600160a01b0319918216179091556001805492909316911617905560ad565b80516001600160a01b0381168114607c575f80fd5b919050565b5f80604083850312156091575f80fd5b6098836067565b915060a4602084016067565b90509250929050565b610cb2806100ba5f395ff3fe608060405234801561000f575f80fd5b5060043610610090575f3560e01c80639681d940116100635780639681d9401461011f578063a3a65e4814610132578063a476f67514610145578063ae9483e014610158578063b6e6a2ca1461016b575f80fd5b806305af82561461009457806309fb2cf5146100a95780630f43a677146100cf57806336339388146100f5575b5f80fd5b6100a76100a2366004610704565b61017e565b005b6100bc6100b73660046108d9565b6101f3565b6040519081526020015b60405180910390f35b6001546100e390600160a01b900460ff1681565b60405160ff90911681526020016100c6565b5f54610107906001600160a01b031681565b6040516001600160a01b0390911681526020016100c6565b6100bc61012d366004610704565b610442565b6100bc610140366004610704565b6104ba565b6100a761015336600461098f565b6104f0565b600154610107906001600160a01b031681565b6100a761017936600461098f565b61054c565b60015460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af11580156101cb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101ef91906109a6565b5050565b6001545f90600a600160a01b90910460ff16106102575760405162461bcd60e51b815260206004820152601b60248201527f4d61782076616c696461746f72206c696d69742072656163686564000000000060448201526064015b60405180910390fd5b60018054600160a01b900460ff16906014610271836109d1565b825460ff9182166101009390930a9283029190920219909116179055505f546040516323b872dd60e01b8152336004820152306024820152600a60448201526001600160a01b03909116906323b872dd906064016020604051808303815f875af11580156102e1573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061030591906109ef565b15156001146103565760405162461bcd60e51b815260206004820152601960248201527f546f6b656e73206661696c656420746f207472616e7366657200000000000000604482015260640161024e565b600154604051634e5bb12760e11b81525f916001600160a01b031690639cb7624e9061038f908a908a908a908a90606490600401610ab7565b6020604051808303815f875af11580156103ab573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103cf91906109a6565b6040805180820182528981526001600160a01b0386166020808301919091525f84815260029091529190912081519293509091819061040e9082610ba1565b5060209190910151600190910180546001600160a01b0319166001600160a01b039092169190911790559695505050505050565b60015460405163025a076560e61b815263ffffffff831660048201525f916001600160a01b031690639681d940906024015b6020604051808303815f875af1158015610490573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104b491906109a6565b92915050565b600154604051631474cbc960e31b815263ffffffff831660048201525f916001600160a01b03169063a3a65e4890602401610474565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b158015610533575f80fd5b505af1158015610545573d5f803e3d5ffd5b5050505050565b600154600160a01b900460ff166105a55760405162461bcd60e51b815260206004820152601760248201527f43757272656e746c79206e6f2076616c696461746f7273000000000000000000604482015260640161024e565b60018054600160a01b900460ff169060146105bf83610c61565b825460ff9182166101009390930a928302919092021990911617905550600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b15801561061f575f80fd5b505af1158015610631573d5f803e3d5ffd5b50505f5460405163a9059cbb60e01b8152336004820152600a60248201526001600160a01b03909116925063a9059cbb91506044016020604051808303815f875af1158015610682573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106a691906109ef565b6106e95760405162461bcd60e51b8152602060048201526014602482015273119d5b991cc819985a5b1959081d1bc81cd95b9960621b604482015260640161024e565b50565b803563ffffffff811681146106ff575f80fd5b919050565b5f60208284031215610714575f80fd5b61071d826106ec565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff8111828210171561075b5761075b610724565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561078a5761078a610724565b604052919050565b5f82601f8301126107a1575f80fd5b813567ffffffffffffffff8111156107bb576107bb610724565b6107ce601f8201601f1916602001610761565b8181528460208386010111156107e2575f80fd5b816020850160208301375f918101602001919091529392505050565b80356001600160a01b03811681146106ff575f80fd5b5f60408284031215610824575f80fd5b61082c610738565b9050610837826106ec565b815260208083013567ffffffffffffffff80821115610854575f80fd5b818501915085601f830112610867575f80fd5b81358181111561087957610879610724565b8060051b915061088a848301610761565b81815291830184019184810190888411156108a3575f80fd5b938501935b838510156108c8576108b9856107fe565b825293850193908501906108a8565b808688015250505050505092915050565b5f805f805f60a086880312156108ed575f80fd5b853567ffffffffffffffff80821115610904575f80fd5b61091089838a01610792565b96506020880135915080821115610925575f80fd5b61093189838a01610792565b95506040880135915080821115610946575f80fd5b61095289838a01610814565b94506060880135915080821115610967575f80fd5b5061097488828901610814565b925050610983608087016107fe565b90509295509295909350565b5f6020828403121561099f575f80fd5b5035919050565b5f602082840312156109b6575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b5f60ff821660ff81036109e6576109e66109bd565b60010192915050565b5f602082840312156109ff575f80fd5b8151801515811461071d575f80fd5b5f81518084525f5b81811015610a3257602081850181015186830182015201610a16565b505f602082860101526020601f19601f83011685010191505092915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015610aac5784516001600160a01b03168252938301936001929092019190830190610a83565b509695505050505050565b60a081525f610ac960a0830188610a0e565b8281036020840152610adb8188610a0e565b90508281036040840152610aef8187610a51565b90508281036060840152610b038186610a51565b91505067ffffffffffffffff831660808301529695505050505050565b600181811c90821680610b3457607f821691505b602082108103610b5257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610b9c57805f5260205f20601f840160051c81016020851015610b7d5750805b601f840160051c820191505b81811015610545575f8155600101610b89565b505050565b815167ffffffffffffffff811115610bbb57610bbb610724565b610bcf81610bc98454610b20565b84610b58565b602080601f831160018114610c02575f8415610beb5750858301515b5f19600386901b1c1916600185901b178555610c59565b5f85815260208120601f198616915b82811015610c3057888601518255948401946001909101908401610c11565b5085821015610c4d57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60ff821680610c7357610c736109bd565b5f19019291505056fea26469706673582212202794e198ef90b3e153fcc7f783ab2f57d5b073357ba79857eeb05e7d3b049a0a64736f6c63430008190033",
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

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteValidatorRemoval(&_SlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteValidatorRemoval(&_SlotAuctionManager.TransactOpts, messageIndex)
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

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x09fb2cf5.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, address rewardAddress) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactor) InitiateValidatorRegistration(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, rewardAddress common.Address) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "initiateValidatorRegistration", nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, rewardAddress)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x09fb2cf5.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, address rewardAddress) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, rewardAddress common.Address) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateValidatorRegistration(&_SlotAuctionManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, rewardAddress)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x09fb2cf5.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, address rewardAddress) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, rewardAddress common.Address) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateValidatorRegistration(&_SlotAuctionManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, rewardAddress)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "initiateValidatorRemoval", validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateValidatorRemoval(&_SlotAuctionManager.TransactOpts, validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateValidatorRemoval(&_SlotAuctionManager.TransactOpts, validationID)
}
