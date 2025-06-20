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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TemporaryID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"becomeValidator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040525f6002556003805460ff19169055348015601c575f80fd5b50604051610ab8380380610ab88339810160408190526039916082565b5f80546001600160a01b039384166001600160a01b0319918216179091556001805492909316911617905560ae565b80516001600160a01b0381168114607d575f80fd5b919050565b5f80604083850312156092575f80fd5b6099836068565b915060a5602084016068565b90509250929050565b6109fd806100bb5f395ff3fe608060405234801561000f575f80fd5b5060043610610085575f3560e01c80636fb191f0116100585780636fb191f01461010d578063a476f67514610116578063ae9483e014610129578063cd993ba51461013c575f80fd5b806305af8256146100895780630f43a6771461009e57806336339388146100c257806337748bdb146100ec575b5f80fd5b61009c6100973660046105a6565b61014f565b005b6003546100ab9060ff1681565b60405160ff90911681526020015b60405180910390f35b5f546100d4906001600160a01b031681565b6040516001600160a01b0390911681526020016100b9565b6100ff6100fa366004610775565b6101c4565b6040519081526020016100b9565b6100ff60025481565b61009c61012436600461081b565b6103a1565b6001546100d4906001600160a01b031681565b61009c61014a36600461081b565b6103fd565b60015460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af115801561019c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101c09190610832565b5050565b6003545f90600a60ff909116106102225760405162461bcd60e51b815260206004820152601b60248201527f4d61782076616c696461746f72206c696d69742072656163686564000000000060448201526064015b60405180910390fd5b6003805460ff16905f6102348361085d565b825460ff9182166101009390930a9283029190920219909116179055505f546040516323b872dd60e01b8152336004820152306024820152600a60448201526001600160a01b03909116906323b872dd906064016020604051808303815f875af11580156102a4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102c8919061087b565b15156001146103195760405162461bcd60e51b815260206004820152601960248201527f546f6b656e73206661696c656420746f207472616e73666572000000000000006044820152606401610219565b600154604051634e5bb12760e11b81525f916001600160a01b031690639cb7624e90610352908990899089908990606490600401610943565b6020604051808303815f875af115801561036e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103929190610832565b60028190559695505050505050565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b1580156103e4575f80fd5b505af11580156103f6573d5f803e3d5ffd5b5050505050565b60035460ff1661044f5760405162461bcd60e51b815260206004820152601760248201527f43757272656e746c79206e6f2076616c696461746f72730000000000000000006044820152606401610219565b6003805460ff16905f610461836109ac565b825460ff9182166101009390930a928302919092021990911617905550600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b1580156104c1575f80fd5b505af11580156104d3573d5f803e3d5ffd5b50505f5460405163a9059cbb60e01b8152336004820152600a60248201526001600160a01b03909116925063a9059cbb91506044016020604051808303815f875af1158015610524573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610548919061087b565b61058b5760405162461bcd60e51b8152602060048201526014602482015273119d5b991cc819985a5b1959081d1bc81cd95b9960621b6044820152606401610219565b50565b803563ffffffff811681146105a1575f80fd5b919050565b5f602082840312156105b6575f80fd5b6105bf8261058e565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff811182821017156105fd576105fd6105c6565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561062c5761062c6105c6565b604052919050565b5f82601f830112610643575f80fd5b813567ffffffffffffffff81111561065d5761065d6105c6565b610670601f8201601f1916602001610603565b818152846020838601011115610684575f80fd5b816020850160208301375f918101602001919091529392505050565b5f604082840312156106b0575f80fd5b6106b86105da565b90506106c38261058e565b815260208083013567ffffffffffffffff808211156106e0575f80fd5b818501915085601f8301126106f3575f80fd5b813581811115610705576107056105c6565b8060051b9150610716848301610603565b818152918301840191848101908884111561072f575f80fd5b938501935b8385101561076457843592506001600160a01b0383168314610754575f80fd5b8282529385019390850190610734565b808688015250505050505092915050565b5f805f8060808587031215610788575f80fd5b843567ffffffffffffffff8082111561079f575f80fd5b6107ab88838901610634565b955060208701359150808211156107c0575f80fd5b6107cc88838901610634565b945060408701359150808211156107e1575f80fd5b6107ed888389016106a0565b93506060870135915080821115610802575f80fd5b5061080f878288016106a0565b91505092959194509250565b5f6020828403121561082b575f80fd5b5035919050565b5f60208284031215610842575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b5f60ff821660ff810361087257610872610849565b60010192915050565b5f6020828403121561088b575f80fd5b815180151581146105bf575f80fd5b5f81518084525f5b818110156108be576020818501810151868301820152016108a2565b505f602082860101526020601f19601f83011685010191505092915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b808310156109385784516001600160a01b0316825293830193600192909201919083019061090f565b509695505050505050565b60a081525f61095560a083018861089a565b8281036020840152610967818861089a565b9050828103604084015261097b81876108dd565b9050828103606084015261098f81866108dd565b91505067ffffffffffffffff831660808301529695505050505050565b5f60ff8216806109be576109be610849565b5f19019291505056fea26469706673582212202bc2bee7aa2372cc43d2611e489baeec803b86b9fc29cd64b5a781b02ff4aa4164736f6c63430008190033",
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

// RemoveValidator is a paid mutator transaction binding the contract method 0xcd993ba5.
//
// Solidity: function removeValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) RemoveValidator(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "removeValidator", validationID)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0xcd993ba5.
//
// Solidity: function removeValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) RemoveValidator(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.RemoveValidator(&_SlotAuctionManager.TransactOpts, validationID)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0xcd993ba5.
//
// Solidity: function removeValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) RemoveValidator(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.RemoveValidator(&_SlotAuctionManager.TransactOpts, validationID)
}
