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
	ABI: "[{\"inputs\":[],\"name\":\"EM_TOKEN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"becomeValidator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"stopValidating\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040525f805460ff191690553480156017575f80fd5b50610900806100255f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80630661ef7e146100645780630f43a6771461007957806337748bdb1461009c57806348d3938a146100bd578063734f05fb146100bd578063fa52c7d8146100dc575b5f80fd5b610077610072366004610494565b6100fb565b005b5f546100859060ff1681565b60405160ff90911681526020015b60405180910390f35b6100af6100aa36600461066f565b6102be565b604051908152602001610093565b6100c45f81565b6040516001600160a01b039091168152602001610093565b6100af6100ea366004610715565b60016020525f908152604090205481565b5f5460ff166101515760405162461bcd60e51b815260206004820152601760248201527f43757272656e746c79206e6f2076616c696461746f727300000000000000000060448201526064015b60405180910390fd5b335f9081526001602052604090205481146101c05760405162461bcd60e51b815260206004820152602960248201527f496e76616c69642073656e6465722c2073656e646572206973206e6f742061206044820152683b30b634b230ba37b960b91b6064820152608401610148565b335f908152600160205260408120819055805460ff1690806101e183610749565b91906101000a81548160ff021916908360ff160217905550505f6001600160a01b031663b6e6a2ca826040518263ffffffff1660e01b815260040161022891815260200190565b5f604051808303815f87803b15801561023f575f80fd5b505af1158015610251573d5f803e3d5ffd5b505060405163a9059cbb60e01b8152336004820152600a60248201525f925063a9059cbb91506044016020604051808303815f875af1158015610296573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102ba9190610764565b5050565b5f8054600a60ff909116106103155760405162461bcd60e51b815260206004820152601b60248201527f4d61782076616c696461746f72206c696d6974207265616368656400000000006044820152606401610148565b335f90815260016020526040902054156103715760405162461bcd60e51b815260206004820152601d60248201527f53656e64657220697320616c726561647920612076616c696461746f720000006044820152606401610148565b5f805460ff16908061038283610783565b825460ff9182166101009390930a9283029190920219909116179055506040516323b872dd60e01b8152336004820152306024820152600a60448201525f906323b872dd906064016020604051808303815f875af11580156103e6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061040a9190610764565b50604051634e5bb12760e11b81525f908190639cb7624e9061043990899089908990899060649060040161084a565b6020604051808303815f875af1158015610455573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061047991906108b3565b335f9081526001602052604090208190559695505050505050565b5f602082840312156104a4575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff811182821017156104e2576104e26104ab565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610511576105116104ab565b604052919050565b5f82601f830112610528575f80fd5b813567ffffffffffffffff811115610542576105426104ab565b610555601f8201601f19166020016104e8565b818152846020838601011115610569575f80fd5b816020850160208301375f918101602001919091529392505050565b80356001600160a01b038116811461059b575f80fd5b919050565b5f604082840312156105b0575f80fd5b6105b86104bf565b9050813563ffffffff811681146105cd575f80fd5b815260208281013567ffffffffffffffff808211156105ea575f80fd5b818501915085601f8301126105fd575f80fd5b81358181111561060f5761060f6104ab565b8060051b91506106208483016104e8565b8181529183018401918481019088841115610639575f80fd5b938501935b8385101561065e5761064f85610585565b8252938501939085019061063e565b808688015250505050505092915050565b5f805f8060808587031215610682575f80fd5b843567ffffffffffffffff80821115610699575f80fd5b6106a588838901610519565b955060208701359150808211156106ba575f80fd5b6106c688838901610519565b945060408701359150808211156106db575f80fd5b6106e7888389016105a0565b935060608701359150808211156106fc575f80fd5b50610709878288016105a0565b91505092959194509250565b5f60208284031215610725575f80fd5b61072e82610585565b9392505050565b634e487b7160e01b5f52601160045260245ffd5b5f60ff82168061075b5761075b610735565b5f190192915050565b5f60208284031215610774575f80fd5b8151801515811461072e575f80fd5b5f60ff821660ff810361079857610798610735565b60010192915050565b5f81518084525f5b818110156107c5576020818501810151868301820152016107a9565b505f602082860101526020601f19601f83011685010191505092915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b8083101561083f5784516001600160a01b03168252938301936001929092019190830190610816565b509695505050505050565b60a081525f61085c60a08301886107a1565b828103602084015261086e81886107a1565b9050828103604084015261088281876107e4565b9050828103606084015261089681866107e4565b91505067ffffffffffffffff831660808301529695505050505050565b5f602082840312156108c3575f80fd5b505191905056fea2646970667358221220150f2ce6695378297d1449c43b597c2c54fc34275862e5231dc816a567036a1264736f6c63430008190033",
}

// SlotAuctionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use SlotAuctionManagerMetaData.ABI instead.
var SlotAuctionManagerABI = SlotAuctionManagerMetaData.ABI

// SlotAuctionManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlotAuctionManagerMetaData.Bin instead.
var SlotAuctionManagerBin = SlotAuctionManagerMetaData.Bin

// DeploySlotAuctionManager deploys a new Ethereum contract, binding an instance of SlotAuctionManager to it.
func DeploySlotAuctionManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SlotAuctionManager, error) {
	parsed, err := SlotAuctionManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlotAuctionManagerBin), backend)
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

// EMTOKENADDRESS is a free data retrieval call binding the contract method 0x734f05fb.
//
// Solidity: function EM_TOKEN_ADDRESS() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerCaller) EMTOKENADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "EM_TOKEN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EMTOKENADDRESS is a free data retrieval call binding the contract method 0x734f05fb.
//
// Solidity: function EM_TOKEN_ADDRESS() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerSession) EMTOKENADDRESS() (common.Address, error) {
	return _SlotAuctionManager.Contract.EMTOKENADDRESS(&_SlotAuctionManager.CallOpts)
}

// EMTOKENADDRESS is a free data retrieval call binding the contract method 0x734f05fb.
//
// Solidity: function EM_TOKEN_ADDRESS() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) EMTOKENADDRESS() (common.Address, error) {
	return _SlotAuctionManager.Contract.EMTOKENADDRESS(&_SlotAuctionManager.CallOpts)
}

// VALIDATORMANAGERADDRESS is a free data retrieval call binding the contract method 0x48d3938a.
//
// Solidity: function VALIDATOR_MANAGER_ADDRESS() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerCaller) VALIDATORMANAGERADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "VALIDATOR_MANAGER_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORMANAGERADDRESS is a free data retrieval call binding the contract method 0x48d3938a.
//
// Solidity: function VALIDATOR_MANAGER_ADDRESS() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerSession) VALIDATORMANAGERADDRESS() (common.Address, error) {
	return _SlotAuctionManager.Contract.VALIDATORMANAGERADDRESS(&_SlotAuctionManager.CallOpts)
}

// VALIDATORMANAGERADDRESS is a free data retrieval call binding the contract method 0x48d3938a.
//
// Solidity: function VALIDATOR_MANAGER_ADDRESS() view returns(address)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) VALIDATORMANAGERADDRESS() (common.Address, error) {
	return _SlotAuctionManager.Contract.VALIDATORMANAGERADDRESS(&_SlotAuctionManager.CallOpts)
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
