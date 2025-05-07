// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleporterregistry

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

// ProtocolRegistryEntry is an auto generated low-level Go binding around an user-defined struct.
type ProtocolRegistryEntry struct {
	Version         *big.Int
	ProtocolAddress common.Address
}

// TeleporterRegistryMetaData contains all meta data concerning the TeleporterRegistry contract.
var TeleporterRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"internalType\":\"structProtocolRegistryEntry[]\",\"name\":\"initialEntries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"AddProtocolVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"LatestVersionUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_VERSION_INCREMENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORS_SOURCE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"addProtocolVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getAddressFromVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestTeleporter\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getTeleporterFromVersion\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"getVersionFromAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610fe0380380610fe083398101604081905261002e916103e2565b7302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801561007e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100a291906104bc565b60805280515f5b818110156100e1576100d98382815181106100c6576100c66104d3565b60200260200101516100e960201b60201c565b6001016100a9565b50505061050c565b80515f0361013e5760405162461bcd60e51b815260206004820181905260248201527f54656c65706f7274657252656769737472793a207a65726f2076657273696f6e60448201526064015b60405180910390fd5b80515f908152600160205260409020546001600160a01b0316156101b75760405162461bcd60e51b815260206004820152602a60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20616c72656044820152696164792065786973747360b01b6064820152608401610135565b60208101516001600160a01b03166102235760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472793a207a65726f2070726f746f636f6044820152686c206164647265737360b81b6064820152608401610135565b5f546102316101f4826104e7565b825111156102985760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20696e637260448201526d0cadacadce840e8dede40d0d2ced60931b6064820152608401610135565b6020828101805184515f90815260018452604080822080546001600160a01b0319166001600160a01b03948516179055925190911681526002909252902054825111156102ff5781516020808401516001600160a01b03165f908152600290915260409020555b602082015182516040516001600160a01b03909216917fa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a905f90a381518110156103725781515f81815560405183917f30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b091a35b5050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156103ac576103ac610376565b60405290565b604051601f8201601f191681016001600160401b03811182821017156103da576103da610376565b604052919050565b5f60208083850312156103f3575f80fd5b82516001600160401b0380821115610409575f80fd5b818501915085601f83011261041c575f80fd5b81518181111561042e5761042e610376565b61043c848260051b016103b2565b818152848101925060069190911b83018401908782111561045b575f80fd5b928401925b818410156104b15760408489031215610477575f80fd5b61047f61038a565b84518152858501516001600160a01b038116811461049b575f80fd5b8187015283526040939093019291840191610460565b979650505050505050565b5f602082840312156104cc575f80fd5b5051919050565b634e487b7160e01b5f52603260045260245ffd5b8082018082111561050657634e487b7160e01b5f52601160045260245ffd5b92915050565b608051610ab561052b5f395f818161014301526102580152610ab55ff3fe608060405234801561000f575f80fd5b506004361061009b575f3560e01c8063ac473ac311610063578063ac473ac31461011f578063b771b3bc14610128578063c07f47d414610136578063d127dc9b1461013e578063d820e64f14610165575f80fd5b80630731775d1461009f578063215abce9146100c357806341f34ed9146100d657806346f9ef49146100eb5780634c1f08ce146100fe575b5f80fd5b6100a65f81565b6040516001600160a01b0390911681526020015b60405180910390f35b6100a66100d13660046107aa565b61016d565b6100e96100e43660046107c1565b61017d565b005b6100a66100f93660046107aa565b6103ec565b61011161010c366004610802565b6104ae565b6040519081526020016100ba565b6101116101f481565b6100a66005600160991b0181565b6101115f5481565b6101117f000000000000000000000000000000000000000000000000000000000000000081565b6100a6610554565b5f610177826103ec565b92915050565b6040516306f8253560e41b815263ffffffff821660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa1580156101c6573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526101ed919081019061089f565b91509150806102545760405162461bcd60e51b815260206004820152602860248201527f54656c65706f7274657252656769737472793a20696e76616c69642077617270604482015267206d65737361676560c01b60648201526084015b60405180910390fd5b81517f0000000000000000000000000000000000000000000000000000000000000000146102d85760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f7274657252656769737472793a20696e76616c696420736f757260448201526a18d94818da185a5b88125160aa1b606482015260840161024b565b60208201516001600160a01b03161561034d5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472793a20696e76616c6964206f726967604482015270696e2073656e646572206164647265737360781b606482015260840161024b565b5f808360400151806020019051810190610367919061099e565b90925090506001600160a01b03811630146103dc5760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f7274657252656769737472793a20696e76616c6964206465737460448201526e696e6174696f6e206164647265737360881b606482015260840161024b565b6103e582610564565b5050505050565b5f815f0361043c5760405162461bcd60e51b815260206004820181905260248201527f54656c65706f7274657252656769737472793a207a65726f2076657273696f6e604482015260640161024b565b5f828152600160205260409020546001600160a01b0316806101775760405162461bcd60e51b815260206004820152602560248201527f54656c65706f7274657252656769737472793a2076657273696f6e206e6f7420604482015264199bdd5b9960da1b606482015260840161024b565b5f6001600160a01b0382166104d55760405162461bcd60e51b815260040161024b90610a17565b6001600160a01b0382165f90815260026020526040812054908190036101775760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472793a2070726f746f636f6c2061646460448201526d1c995cdcc81b9bdd08199bdd5b9960921b606482015260840161024b565b5f61055f5f546103ec565b905090565b80515f036105b45760405162461bcd60e51b815260206004820181905260248201527f54656c65706f7274657252656769737472793a207a65726f2076657273696f6e604482015260640161024b565b80515f908152600160205260409020546001600160a01b03161561062d5760405162461bcd60e51b815260206004820152602a60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20616c72656044820152696164792065786973747360b01b606482015260840161024b565b60208101516001600160a01b03166106575760405162461bcd60e51b815260040161024b90610a17565b5f546106656101f482610a60565b825111156106cc5760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20696e637260448201526d0cadacadce840e8dede40d0d2ced60931b606482015260840161024b565b6020828101805184515f90815260018452604080822080546001600160a01b0319166001600160a01b03948516179055925190911681526002909252902054825111156107335781516020808401516001600160a01b03165f908152600290915260409020555b602082015182516040516001600160a01b03909216917fa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a905f90a381518110156107a65781515f81815560405183917f30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b091a35b5050565b5f602082840312156107ba575f80fd5b5035919050565b5f602082840312156107d1575f80fd5b813563ffffffff811681146107e4575f80fd5b9392505050565b6001600160a01b03811681146107ff575f80fd5b50565b5f60208284031215610812575f80fd5b81356107e4816107eb565b634e487b7160e01b5f52604160045260245ffd5b6040516060810167ffffffffffffffff811182821017156108545761085461081d565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156108835761088361081d565b604052919050565b8051801515811461089a575f80fd5b919050565b5f80604083850312156108b0575f80fd5b825167ffffffffffffffff808211156108c7575f80fd5b90840190606082870312156108da575f80fd5b6108e2610831565b825181526020808401516108f5816107eb565b8282015260408401518381111561090a575f80fd5b80850194505087601f85011261091e575f80fd5b8351838111156109305761093061081d565b610942601f8201601f1916830161085a565b93508084528882828701011115610957575f80fd5b5f5b81811015610974578581018301518582018401528201610959565b505f8282860101525082604083015281955061099181880161088b565b9450505050509250929050565b5f8082840360608112156109b0575f80fd5b60408112156109bd575f80fd5b506040516040810181811067ffffffffffffffff821117156109e1576109e161081d565b6040528351815260208401516109f6816107eb565b60208201526040840151909250610a0c816107eb565b809150509250929050565b60208082526029908201527f54656c65706f7274657252656769737472793a207a65726f2070726f746f636f6040820152686c206164647265737360b81b606082015260800190565b8082018082111561017757634e487b7160e01b5f52601160045260245ffdfea26469706673582212209162d7e828050fd09c74d7e4669100d88b34c03d1d6e0d9e7ad8152d615fc2e464736f6c63430008190033",
}

// TeleporterRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterRegistryMetaData.ABI instead.
var TeleporterRegistryABI = TeleporterRegistryMetaData.ABI

// TeleporterRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeleporterRegistryMetaData.Bin instead.
var TeleporterRegistryBin = TeleporterRegistryMetaData.Bin

// DeployTeleporterRegistry deploys a new Ethereum contract, binding an instance of TeleporterRegistry to it.
func DeployTeleporterRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, initialEntries []ProtocolRegistryEntry) (common.Address, *types.Transaction, *TeleporterRegistry, error) {
	parsed, err := TeleporterRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeleporterRegistryBin), backend, initialEntries)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeleporterRegistry{TeleporterRegistryCaller: TeleporterRegistryCaller{contract: contract}, TeleporterRegistryTransactor: TeleporterRegistryTransactor{contract: contract}, TeleporterRegistryFilterer: TeleporterRegistryFilterer{contract: contract}}, nil
}

// TeleporterRegistry is an auto generated Go binding around an Ethereum contract.
type TeleporterRegistry struct {
	TeleporterRegistryCaller     // Read-only binding to the contract
	TeleporterRegistryTransactor // Write-only binding to the contract
	TeleporterRegistryFilterer   // Log filterer for contract events
}

// TeleporterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterRegistrySession struct {
	Contract     *TeleporterRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TeleporterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterRegistryCallerSession struct {
	Contract *TeleporterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TeleporterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterRegistryTransactorSession struct {
	Contract     *TeleporterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TeleporterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterRegistryRaw struct {
	Contract *TeleporterRegistry // Generic contract binding to access the raw methods on
}

// TeleporterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterRegistryCallerRaw struct {
	Contract *TeleporterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterRegistryTransactorRaw struct {
	Contract *TeleporterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterRegistry creates a new instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistry(address common.Address, backend bind.ContractBackend) (*TeleporterRegistry, error) {
	contract, err := bindTeleporterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistry{TeleporterRegistryCaller: TeleporterRegistryCaller{contract: contract}, TeleporterRegistryTransactor: TeleporterRegistryTransactor{contract: contract}, TeleporterRegistryFilterer: TeleporterRegistryFilterer{contract: contract}}, nil
}

// NewTeleporterRegistryCaller creates a new read-only instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeleporterRegistryCaller, error) {
	contract, err := bindTeleporterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryCaller{contract: contract}, nil
}

// NewTeleporterRegistryTransactor creates a new write-only instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterRegistryTransactor, error) {
	contract, err := bindTeleporterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryTransactor{contract: contract}, nil
}

// NewTeleporterRegistryFilterer creates a new log filterer instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterRegistryFilterer, error) {
	contract, err := bindTeleporterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryFilterer{contract: contract}, nil
}

// bindTeleporterRegistry binds a generic wrapper to an already deployed contract.
func bindTeleporterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterRegistry *TeleporterRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterRegistry.Contract.TeleporterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterRegistry *TeleporterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.TeleporterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterRegistry *TeleporterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.TeleporterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterRegistry *TeleporterRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterRegistry *TeleporterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterRegistry *TeleporterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCaller) MAXVERSIONINCREMENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "MAX_VERSION_INCREMENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistrySession) MAXVERSIONINCREMENT() (*big.Int, error) {
	return _TeleporterRegistry.Contract.MAXVERSIONINCREMENT(&_TeleporterRegistry.CallOpts)
}

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) MAXVERSIONINCREMENT() (*big.Int, error) {
	return _TeleporterRegistry.Contract.MAXVERSIONINCREMENT(&_TeleporterRegistry.CallOpts)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) VALIDATORSSOURCEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "VALIDATORS_SOURCE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _TeleporterRegistry.Contract.VALIDATORSSOURCEADDRESS(&_TeleporterRegistry.CallOpts)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _TeleporterRegistry.Contract.VALIDATORSSOURCEADDRESS(&_TeleporterRegistry.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterRegistry.Contract.WARPMESSENGER(&_TeleporterRegistry.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterRegistry.Contract.WARPMESSENGER(&_TeleporterRegistry.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterRegistry *TeleporterRegistryCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterRegistry *TeleporterRegistrySession) BlockchainID() ([32]byte, error) {
	return _TeleporterRegistry.Contract.BlockchainID(&_TeleporterRegistry.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterRegistry.Contract.BlockchainID(&_TeleporterRegistry.CallOpts)
}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetAddressFromVersion(opts *bind.CallOpts, version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getAddressFromVersion", version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) GetAddressFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetAddressFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetAddressFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetAddressFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetLatestTeleporter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getLatestTeleporter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) GetLatestTeleporter() (common.Address, error) {
	return _TeleporterRegistry.Contract.GetLatestTeleporter(&_TeleporterRegistry.CallOpts)
}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetLatestTeleporter() (common.Address, error) {
	return _TeleporterRegistry.Contract.GetLatestTeleporter(&_TeleporterRegistry.CallOpts)
}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetTeleporterFromVersion(opts *bind.CallOpts, version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getTeleporterFromVersion", version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) GetTeleporterFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetTeleporterFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetTeleporterFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetTeleporterFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetVersionFromAddress(opts *bind.CallOpts, protocolAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getVersionFromAddress", protocolAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistrySession) GetVersionFromAddress(protocolAddress common.Address) (*big.Int, error) {
	return _TeleporterRegistry.Contract.GetVersionFromAddress(&_TeleporterRegistry.CallOpts, protocolAddress)
}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetVersionFromAddress(protocolAddress common.Address) (*big.Int, error) {
	return _TeleporterRegistry.Contract.GetVersionFromAddress(&_TeleporterRegistry.CallOpts, protocolAddress)
}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCaller) LatestVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "latestVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistrySession) LatestVersion() (*big.Int, error) {
	return _TeleporterRegistry.Contract.LatestVersion(&_TeleporterRegistry.CallOpts)
}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) LatestVersion() (*big.Int, error) {
	return _TeleporterRegistry.Contract.LatestVersion(&_TeleporterRegistry.CallOpts)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_TeleporterRegistry *TeleporterRegistryTransactor) AddProtocolVersion(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterRegistry.contract.Transact(opts, "addProtocolVersion", messageIndex)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_TeleporterRegistry *TeleporterRegistrySession) AddProtocolVersion(messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.AddProtocolVersion(&_TeleporterRegistry.TransactOpts, messageIndex)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_TeleporterRegistry *TeleporterRegistryTransactorSession) AddProtocolVersion(messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.AddProtocolVersion(&_TeleporterRegistry.TransactOpts, messageIndex)
}

// TeleporterRegistryAddProtocolVersionIterator is returned from FilterAddProtocolVersion and is used to iterate over the raw logs and unpacked data for AddProtocolVersion events raised by the TeleporterRegistry contract.
type TeleporterRegistryAddProtocolVersionIterator struct {
	Event *TeleporterRegistryAddProtocolVersion // Event containing the contract specifics and raw log

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
func (it *TeleporterRegistryAddProtocolVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterRegistryAddProtocolVersion)
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
		it.Event = new(TeleporterRegistryAddProtocolVersion)
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
func (it *TeleporterRegistryAddProtocolVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterRegistryAddProtocolVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterRegistryAddProtocolVersion represents a AddProtocolVersion event raised by the TeleporterRegistry contract.
type TeleporterRegistryAddProtocolVersion struct {
	Version         *big.Int
	ProtocolAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddProtocolVersion is a free log retrieval operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_TeleporterRegistry *TeleporterRegistryFilterer) FilterAddProtocolVersion(opts *bind.FilterOpts, version []*big.Int, protocolAddress []common.Address) (*TeleporterRegistryAddProtocolVersionIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.FilterLogs(opts, "AddProtocolVersion", versionRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryAddProtocolVersionIterator{contract: _TeleporterRegistry.contract, event: "AddProtocolVersion", logs: logs, sub: sub}, nil
}

// WatchAddProtocolVersion is a free log subscription operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_TeleporterRegistry *TeleporterRegistryFilterer) WatchAddProtocolVersion(opts *bind.WatchOpts, sink chan<- *TeleporterRegistryAddProtocolVersion, version []*big.Int, protocolAddress []common.Address) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.WatchLogs(opts, "AddProtocolVersion", versionRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterRegistryAddProtocolVersion)
				if err := _TeleporterRegistry.contract.UnpackLog(event, "AddProtocolVersion", log); err != nil {
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

// ParseAddProtocolVersion is a log parse operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_TeleporterRegistry *TeleporterRegistryFilterer) ParseAddProtocolVersion(log types.Log) (*TeleporterRegistryAddProtocolVersion, error) {
	event := new(TeleporterRegistryAddProtocolVersion)
	if err := _TeleporterRegistry.contract.UnpackLog(event, "AddProtocolVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterRegistryLatestVersionUpdatedIterator is returned from FilterLatestVersionUpdated and is used to iterate over the raw logs and unpacked data for LatestVersionUpdated events raised by the TeleporterRegistry contract.
type TeleporterRegistryLatestVersionUpdatedIterator struct {
	Event *TeleporterRegistryLatestVersionUpdated // Event containing the contract specifics and raw log

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
func (it *TeleporterRegistryLatestVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterRegistryLatestVersionUpdated)
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
		it.Event = new(TeleporterRegistryLatestVersionUpdated)
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
func (it *TeleporterRegistryLatestVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterRegistryLatestVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterRegistryLatestVersionUpdated represents a LatestVersionUpdated event raised by the TeleporterRegistry contract.
type TeleporterRegistryLatestVersionUpdated struct {
	OldVersion *big.Int
	NewVersion *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLatestVersionUpdated is a free log retrieval operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_TeleporterRegistry *TeleporterRegistryFilterer) FilterLatestVersionUpdated(opts *bind.FilterOpts, oldVersion []*big.Int, newVersion []*big.Int) (*TeleporterRegistryLatestVersionUpdatedIterator, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.FilterLogs(opts, "LatestVersionUpdated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryLatestVersionUpdatedIterator{contract: _TeleporterRegistry.contract, event: "LatestVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchLatestVersionUpdated is a free log subscription operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_TeleporterRegistry *TeleporterRegistryFilterer) WatchLatestVersionUpdated(opts *bind.WatchOpts, sink chan<- *TeleporterRegistryLatestVersionUpdated, oldVersion []*big.Int, newVersion []*big.Int) (event.Subscription, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.WatchLogs(opts, "LatestVersionUpdated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterRegistryLatestVersionUpdated)
				if err := _TeleporterRegistry.contract.UnpackLog(event, "LatestVersionUpdated", log); err != nil {
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

// ParseLatestVersionUpdated is a log parse operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_TeleporterRegistry *TeleporterRegistryFilterer) ParseLatestVersionUpdated(log types.Log) (*TeleporterRegistryLatestVersionUpdated, error) {
	event := new(TeleporterRegistryLatestVersionUpdated)
	if err := _TeleporterRegistry.contract.UnpackLog(event, "LatestVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
