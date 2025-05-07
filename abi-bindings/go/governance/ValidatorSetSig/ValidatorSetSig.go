// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorsetsig

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

// ValidatorSetSigMessage is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSetSigMessage struct {
	TargetBlockchainID     [32]byte
	ValidatorSetSigAddress common.Address
	TargetContractAddress  common.Address
	Nonce                  *big.Int
	Value                  *big.Int
	Payload                []byte
}

// ValidatorSetSigMetaData contains all meta data concerning the ValidatorSetSig contract.
var ValidatorSetSigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validatorBlockchainID_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Delivered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VALIDATORS_SOURCE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"executeCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContractAddress\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"targetBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorSetSigAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structValidatorSetSigMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"validateMessage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c060405234801561000f575f80fd5b50604051610b76380380610b7683398101604081905261002e916100b2565b60015f5560808190526040805163084279ef60e31b8152905173020000000000000000000000000000000000000591634213cf789160048083019260209291908290030181865afa158015610085573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100a991906100b2565b60a052506100c9565b5f602082840312156100c2575f80fd5b5051919050565b60805160a051610a7e6100f85f395f818161017701526104ba01525f818161012a01526102790152610a7e5ff3fe608060405260043610610071575f3560e01c80637ecebe001161004c5780637ecebe00146100e05780638d6e579d14610119578063b771b3bc1461014c578063d127dc9b14610166575f80fd5b80630731775d1461007c5780635433da42146100ac5780637d969c34146100c1575f80fd5b3661007857005b5f80fd5b348015610087575f80fd5b5061008f5f81565b6040516001600160a01b0390911681526020015b60405180910390f35b6100bf6100ba366004610642565b610199565b005b3480156100cc575f80fd5b506100bf6100db36600461079a565b6104b6565b3480156100eb575f80fd5b5061010b6100fa366004610842565b60016020525f908152604090205481565b6040519081526020016100a3565b348015610124575f80fd5b5061010b7f000000000000000000000000000000000000000000000000000000000000000081565b348015610157575f80fd5b5061008f6005600160991b0181565b348015610171575f80fd5b5061010b7f000000000000000000000000000000000000000000000000000000000000000081565b6101a161061a565b6040516306f8253560e41b815263ffffffff821660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa1580156101ea573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261021191908101906108c9565b91509150806102755760405162461bcd60e51b815260206004820152602560248201527f56616c696461746f725365745369673a20696e76616c69642077617270206d65604482015264737361676560d81b60648201526084015b60405180910390fd5b81517f0000000000000000000000000000000000000000000000000000000000000000146102f45760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f725365745369673a20696e76616c696420736f75726365436044820152651a185a5b925160d21b606482015260840161026c565b60208201516001600160a01b0316156103655760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f725365745369673a206e6f6e2d7a65726f206f726967696e60448201526c53656e6465724164647265737360981b606482015260840161026c565b5f826040015180602001905181019061037e9190610969565b9050610389816104b6565b6060810151610399906001610a08565b604080830180516001600160a01b039081165f9081526001602052838120949094559051608085015160a0860151935191909216926103d791610a2d565b5f6040518083038185875af1925050503d805f8114610411576040519150601f19603f3d011682016040523d82523d5f602084013e610416565b606091505b50509050806104675760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f725365745369673a2063616c6c206661696c656400000000604482015260640161026c565b816060015182604001516001600160a01b03167f5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed960405160405180910390a3505050506104b360015f55565b50565b80517f00000000000000000000000000000000000000000000000000000000000000001461053a5760405162461bcd60e51b815260206004820152602b60248201527f56616c696461746f725365745369673a20696e76616c6964207461726765744260448201526a1b1bd8dad8da185a5b925160aa1b606482015260840161026c565b60208101516001600160a01b031630146105ae5760405162461bcd60e51b815260206004820152602f60248201527f56616c696461746f725365745369673a20696e76616c69642076616c6964617460448201526e6f725365745369674164647265737360881b606482015260840161026c565b60608101516040808301516001600160a01b03165f90815260016020522054146104b35760405162461bcd60e51b815260206004820152601e60248201527f56616c696461746f725365745369673a20696e76616c6964206e6f6e63650000604482015260640161026c565b60025f540361063c57604051633ee5aeb560e01b815260040160405180910390fd5b60025f55565b5f60208284031215610652575f80fd5b813563ffffffff81168114610665575f80fd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b60405160c0810167ffffffffffffffff811182821017156106a3576106a361066c565b60405290565b6040516060810167ffffffffffffffff811182821017156106a3576106a361066c565b604051601f8201601f1916810167ffffffffffffffff811182821017156106f5576106f561066c565b604052919050565b6001600160a01b03811681146104b3575f80fd5b803561071c816106fd565b919050565b5f67ffffffffffffffff82111561073a5761073a61066c565b50601f01601f191660200190565b5f82601f830112610757575f80fd5b813561076a61076582610721565b6106cc565b81815284602083860101111561077e575f80fd5b816020850160208301375f918101602001919091529392505050565b5f602082840312156107aa575f80fd5b813567ffffffffffffffff808211156107c1575f80fd5b9083019060c082860312156107d4575f80fd5b6107dc610680565b823581526107ec60208401610711565b60208201526107fd60408401610711565b6040820152606083013560608201526080830135608082015260a083013582811115610827575f80fd5b61083387828601610748565b60a08301525095945050505050565b5f60208284031215610852575f80fd5b8135610665816106fd565b5f5b8381101561087757818101518382015260200161085f565b50505f910152565b5f82601f83011261088e575f80fd5b815161089c61076582610721565b8181528460208386010111156108b0575f80fd5b6108c182602083016020870161085d565b949350505050565b5f80604083850312156108da575f80fd5b825167ffffffffffffffff808211156108f1575f80fd5b9084019060608287031215610904575f80fd5b61090c6106a9565b82518152602083015161091e816106fd565b6020820152604083015182811115610934575f80fd5b6109408882860161087f565b6040830152508094505050506020830151801515811461095e575f80fd5b809150509250929050565b5f60208284031215610979575f80fd5b815167ffffffffffffffff80821115610990575f80fd5b9083019060c082860312156109a3575f80fd5b6109ab610680565b8251815260208301516109bd816106fd565b602082015260408301516109d0816106fd565b80604083015250606083015160608201526080830151608082015260a0830151828111156109fc575f80fd5b6108338782860161087f565b80820180821115610a2757634e487b7160e01b5f52601160045260245ffd5b92915050565b5f8251610a3e81846020870161085d565b919091019291505056fea2646970667358221220fb0992b7159afb847b6729087aa8a977058e0da20139e74ca5815e76314647d864736f6c63430008190033",
}

// ValidatorSetSigABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorSetSigMetaData.ABI instead.
var ValidatorSetSigABI = ValidatorSetSigMetaData.ABI

// ValidatorSetSigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorSetSigMetaData.Bin instead.
var ValidatorSetSigBin = ValidatorSetSigMetaData.Bin

// DeployValidatorSetSig deploys a new Ethereum contract, binding an instance of ValidatorSetSig to it.
func DeployValidatorSetSig(auth *bind.TransactOpts, backend bind.ContractBackend, validatorBlockchainID_ [32]byte) (common.Address, *types.Transaction, *ValidatorSetSig, error) {
	parsed, err := ValidatorSetSigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorSetSigBin), backend, validatorBlockchainID_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorSetSig{ValidatorSetSigCaller: ValidatorSetSigCaller{contract: contract}, ValidatorSetSigTransactor: ValidatorSetSigTransactor{contract: contract}, ValidatorSetSigFilterer: ValidatorSetSigFilterer{contract: contract}}, nil
}

// ValidatorSetSig is an auto generated Go binding around an Ethereum contract.
type ValidatorSetSig struct {
	ValidatorSetSigCaller     // Read-only binding to the contract
	ValidatorSetSigTransactor // Write-only binding to the contract
	ValidatorSetSigFilterer   // Log filterer for contract events
}

// ValidatorSetSigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorSetSigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorSetSigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorSetSigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSetSigSession struct {
	Contract     *ValidatorSetSig  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorSetSigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorSetSigCallerSession struct {
	Contract *ValidatorSetSigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ValidatorSetSigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorSetSigTransactorSession struct {
	Contract     *ValidatorSetSigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ValidatorSetSigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorSetSigRaw struct {
	Contract *ValidatorSetSig // Generic contract binding to access the raw methods on
}

// ValidatorSetSigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorSetSigCallerRaw struct {
	Contract *ValidatorSetSigCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorSetSigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorSetSigTransactorRaw struct {
	Contract *ValidatorSetSigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorSetSig creates a new instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSig(address common.Address, backend bind.ContractBackend) (*ValidatorSetSig, error) {
	contract, err := bindValidatorSetSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSig{ValidatorSetSigCaller: ValidatorSetSigCaller{contract: contract}, ValidatorSetSigTransactor: ValidatorSetSigTransactor{contract: contract}, ValidatorSetSigFilterer: ValidatorSetSigFilterer{contract: contract}}, nil
}

// NewValidatorSetSigCaller creates a new read-only instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSigCaller(address common.Address, caller bind.ContractCaller) (*ValidatorSetSigCaller, error) {
	contract, err := bindValidatorSetSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigCaller{contract: contract}, nil
}

// NewValidatorSetSigTransactor creates a new write-only instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSigTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorSetSigTransactor, error) {
	contract, err := bindValidatorSetSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigTransactor{contract: contract}, nil
}

// NewValidatorSetSigFilterer creates a new log filterer instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSigFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorSetSigFilterer, error) {
	contract, err := bindValidatorSetSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigFilterer{contract: contract}, nil
}

// bindValidatorSetSig binds a generic wrapper to an already deployed contract.
func bindValidatorSetSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorSetSigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSetSig *ValidatorSetSigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSetSig.Contract.ValidatorSetSigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSetSig *ValidatorSetSigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ValidatorSetSigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSetSig *ValidatorSetSigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ValidatorSetSigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSetSig *ValidatorSetSigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSetSig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSetSig *ValidatorSetSigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSetSig *ValidatorSetSigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.contract.Transact(opts, method, params...)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCaller) VALIDATORSSOURCEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "VALIDATORS_SOURCE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigSession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _ValidatorSetSig.Contract.VALIDATORSSOURCEADDRESS(&_ValidatorSetSig.CallOpts)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _ValidatorSetSig.Contract.VALIDATORSSOURCEADDRESS(&_ValidatorSetSig.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigSession) WARPMESSENGER() (common.Address, error) {
	return _ValidatorSetSig.Contract.WARPMESSENGER(&_ValidatorSetSig.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) WARPMESSENGER() (common.Address, error) {
	return _ValidatorSetSig.Contract.WARPMESSENGER(&_ValidatorSetSig.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigSession) BlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.BlockchainID(&_ValidatorSetSig.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) BlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.BlockchainID(&_ValidatorSetSig.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address targetContractAddress) view returns(uint256 nonce)
func (_ValidatorSetSig *ValidatorSetSigCaller) Nonces(opts *bind.CallOpts, targetContractAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "nonces", targetContractAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address targetContractAddress) view returns(uint256 nonce)
func (_ValidatorSetSig *ValidatorSetSigSession) Nonces(targetContractAddress common.Address) (*big.Int, error) {
	return _ValidatorSetSig.Contract.Nonces(&_ValidatorSetSig.CallOpts, targetContractAddress)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address targetContractAddress) view returns(uint256 nonce)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) Nonces(targetContractAddress common.Address) (*big.Int, error) {
	return _ValidatorSetSig.Contract.Nonces(&_ValidatorSetSig.CallOpts, targetContractAddress)
}

// ValidateMessage is a free data retrieval call binding the contract method 0x7d969c34.
//
// Solidity: function validateMessage((bytes32,address,address,uint256,uint256,bytes) message) view returns()
func (_ValidatorSetSig *ValidatorSetSigCaller) ValidateMessage(opts *bind.CallOpts, message ValidatorSetSigMessage) error {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "validateMessage", message)

	if err != nil {
		return err
	}

	return err

}

// ValidateMessage is a free data retrieval call binding the contract method 0x7d969c34.
//
// Solidity: function validateMessage((bytes32,address,address,uint256,uint256,bytes) message) view returns()
func (_ValidatorSetSig *ValidatorSetSigSession) ValidateMessage(message ValidatorSetSigMessage) error {
	return _ValidatorSetSig.Contract.ValidateMessage(&_ValidatorSetSig.CallOpts, message)
}

// ValidateMessage is a free data retrieval call binding the contract method 0x7d969c34.
//
// Solidity: function validateMessage((bytes32,address,address,uint256,uint256,bytes) message) view returns()
func (_ValidatorSetSig *ValidatorSetSigCallerSession) ValidateMessage(message ValidatorSetSigMessage) error {
	return _ValidatorSetSig.Contract.ValidateMessage(&_ValidatorSetSig.CallOpts, message)
}

// ValidatorBlockchainID is a free data retrieval call binding the contract method 0x8d6e579d.
//
// Solidity: function validatorBlockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCaller) ValidatorBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "validatorBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ValidatorBlockchainID is a free data retrieval call binding the contract method 0x8d6e579d.
//
// Solidity: function validatorBlockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigSession) ValidatorBlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.ValidatorBlockchainID(&_ValidatorSetSig.CallOpts)
}

// ValidatorBlockchainID is a free data retrieval call binding the contract method 0x8d6e579d.
//
// Solidity: function validatorBlockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) ValidatorBlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.ValidatorBlockchainID(&_ValidatorSetSig.CallOpts)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x5433da42.
//
// Solidity: function executeCall(uint32 messageIndex) payable returns()
func (_ValidatorSetSig *ValidatorSetSigTransactor) ExecuteCall(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorSetSig.contract.Transact(opts, "executeCall", messageIndex)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x5433da42.
//
// Solidity: function executeCall(uint32 messageIndex) payable returns()
func (_ValidatorSetSig *ValidatorSetSigSession) ExecuteCall(messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ExecuteCall(&_ValidatorSetSig.TransactOpts, messageIndex)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x5433da42.
//
// Solidity: function executeCall(uint32 messageIndex) payable returns()
func (_ValidatorSetSig *ValidatorSetSigTransactorSession) ExecuteCall(messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ExecuteCall(&_ValidatorSetSig.TransactOpts, messageIndex)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorSetSig *ValidatorSetSigTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSetSig.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorSetSig *ValidatorSetSigSession) Receive() (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.Receive(&_ValidatorSetSig.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorSetSig *ValidatorSetSigTransactorSession) Receive() (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.Receive(&_ValidatorSetSig.TransactOpts)
}

// ValidatorSetSigDeliveredIterator is returned from FilterDelivered and is used to iterate over the raw logs and unpacked data for Delivered events raised by the ValidatorSetSig contract.
type ValidatorSetSigDeliveredIterator struct {
	Event *ValidatorSetSigDelivered // Event containing the contract specifics and raw log

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
func (it *ValidatorSetSigDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetSigDelivered)
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
		it.Event = new(ValidatorSetSigDelivered)
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
func (it *ValidatorSetSigDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetSigDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetSigDelivered represents a Delivered event raised by the ValidatorSetSig contract.
type ValidatorSetSigDelivered struct {
	TargetContractAddress common.Address
	Nonce                 *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDelivered is a free log retrieval operation binding the contract event 0x5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed9.
//
// Solidity: event Delivered(address indexed targetContractAddress, uint256 indexed nonce)
func (_ValidatorSetSig *ValidatorSetSigFilterer) FilterDelivered(opts *bind.FilterOpts, targetContractAddress []common.Address, nonce []*big.Int) (*ValidatorSetSigDeliveredIterator, error) {

	var targetContractAddressRule []interface{}
	for _, targetContractAddressItem := range targetContractAddress {
		targetContractAddressRule = append(targetContractAddressRule, targetContractAddressItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ValidatorSetSig.contract.FilterLogs(opts, "Delivered", targetContractAddressRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigDeliveredIterator{contract: _ValidatorSetSig.contract, event: "Delivered", logs: logs, sub: sub}, nil
}

// WatchDelivered is a free log subscription operation binding the contract event 0x5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed9.
//
// Solidity: event Delivered(address indexed targetContractAddress, uint256 indexed nonce)
func (_ValidatorSetSig *ValidatorSetSigFilterer) WatchDelivered(opts *bind.WatchOpts, sink chan<- *ValidatorSetSigDelivered, targetContractAddress []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var targetContractAddressRule []interface{}
	for _, targetContractAddressItem := range targetContractAddress {
		targetContractAddressRule = append(targetContractAddressRule, targetContractAddressItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ValidatorSetSig.contract.WatchLogs(opts, "Delivered", targetContractAddressRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetSigDelivered)
				if err := _ValidatorSetSig.contract.UnpackLog(event, "Delivered", log); err != nil {
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

// ParseDelivered is a log parse operation binding the contract event 0x5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed9.
//
// Solidity: event Delivered(address indexed targetContractAddress, uint256 indexed nonce)
func (_ValidatorSetSig *ValidatorSetSigFilterer) ParseDelivered(log types.Log) (*ValidatorSetSigDelivered, error) {
	event := new(ValidatorSetSigDelivered)
	if err := _ValidatorSetSig.contract.UnpackLog(event, "Delivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
