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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"becomeValidator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"stopValidating\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zilch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001805460ff60a01b19169055348015601b575f80fd5b50604051610ac3380380610ac38339810160408190526038916081565b5f80546001600160a01b039384166001600160a01b0319918216179091556001805492909316911617905560ad565b80516001600160a01b0381168114607c575f80fd5b919050565b5f80604083850312156091575f80fd5b6098836067565b915060a4602084016067565b90509250929050565b610a09806100ba5f395ff3fe608060405234801561000f575f80fd5b506004361061007a575f3560e01c806337748bdb1161005857806337748bdb146100e8578063ae9483e014610109578063bebf2eee1461011c578063fa52c7d814610125575f80fd5b80630661ef7e1461007e5780630f43a6771461009357806336339388146100be575b5f80fd5b61009161008c36600461059d565b610144565b005b6001546100a790600160a01b900460ff1681565b60405160ff90911681526020015b60405180910390f35b5f546100d0906001600160a01b031681565b6040516001600160a01b0390911681526020016100b5565b6100fb6100f6366004610778565b610360565b6040519081526020016100b5565b6001546100d0906001600160a01b031681565b6100fb60035481565b6100fb61013336600461081e565b60026020525f908152604090205481565b600154600160a01b900460ff166101a25760405162461bcd60e51b815260206004820152601760248201527f43757272656e746c79206e6f2076616c696461746f727300000000000000000060448201526064015b60405180910390fd5b335f9081526002602052604090205481146102115760405162461bcd60e51b815260206004820152602960248201527f496e76616c69642073656e6465722c2073656e646572206973206e6f742061206044820152683b30b634b230ba37b960b91b6064820152608401610199565b335f52600260205260018054600160a01b900460ff1690601461023383610852565b825460ff9182166101009390930a928302919092021990911617905550600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b158015610293575f80fd5b505af11580156102a5573d5f803e3d5ffd5b50505f5460405163a9059cbb60e01b8152336004820152600a60248201526001600160a01b03909116925063a9059cbb91506044016020604051808303815f875af11580156102f6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061031a919061086d565b61035d5760405162461bcd60e51b8152602060048201526014602482015273119d5b991cc819985a5b1959081d1bc81cd95b9960621b6044820152606401610199565b50565b6001545f90600a600160a01b90910460ff16106103bf5760405162461bcd60e51b815260206004820152601b60248201527f4d61782076616c696461746f72206c696d6974207265616368656400000000006044820152606401610199565b335f908152600260205260409020541561041b5760405162461bcd60e51b815260206004820152601d60248201527f53656e64657220697320616c726561647920612076616c696461746f720000006044820152606401610199565b60018054600160a01b900460ff169060146104358361088c565b825460ff9182166101009390930a9283029190920219909116179055505f546040516323b872dd60e01b8152336004820152306024820152600a60448201526001600160a01b03909116906323b872dd906064016020604051808303815f875af11580156104a5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104c9919061086d565b151560011461051a5760405162461bcd60e51b815260206004820152601960248201527f546f6b656e73206661696c656420746f207472616e73666572000000000000006044820152606401610199565b600154604051634e5bb12760e11b81525f916001600160a01b031690639cb7624e90610553908990899089908990606490600401610953565b6020604051808303815f875af115801561056f573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061059391906109bc565b9695505050505050565b5f602082840312156105ad575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff811182821017156105eb576105eb6105b4565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561061a5761061a6105b4565b604052919050565b5f82601f830112610631575f80fd5b813567ffffffffffffffff81111561064b5761064b6105b4565b61065e601f8201601f19166020016105f1565b818152846020838601011115610672575f80fd5b816020850160208301375f918101602001919091529392505050565b80356001600160a01b03811681146106a4575f80fd5b919050565b5f604082840312156106b9575f80fd5b6106c16105c8565b9050813563ffffffff811681146106d6575f80fd5b815260208281013567ffffffffffffffff808211156106f3575f80fd5b818501915085601f830112610706575f80fd5b813581811115610718576107186105b4565b8060051b91506107298483016105f1565b8181529183018401918481019088841115610742575f80fd5b938501935b83851015610767576107588561068e565b82529385019390850190610747565b808688015250505050505092915050565b5f805f806080858703121561078b575f80fd5b843567ffffffffffffffff808211156107a2575f80fd5b6107ae88838901610622565b955060208701359150808211156107c3575f80fd5b6107cf88838901610622565b945060408701359150808211156107e4575f80fd5b6107f0888389016106a9565b93506060870135915080821115610805575f80fd5b50610812878288016106a9565b91505092959194509250565b5f6020828403121561082e575f80fd5b6108378261068e565b9392505050565b634e487b7160e01b5f52601160045260245ffd5b5f60ff8216806108645761086461083e565b5f190192915050565b5f6020828403121561087d575f80fd5b81518015158114610837575f80fd5b5f60ff821660ff81036108a1576108a161083e565b60010192915050565b5f81518084525f5b818110156108ce576020818501810151868301820152016108b2565b505f602082860101526020601f19601f83011685010191505092915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b808310156109485784516001600160a01b0316825293830193600192909201919083019061091f565b509695505050505050565b60a081525f61096560a08301886108aa565b828103602084015261097781886108aa565b9050828103604084015261098b81876108ed565b9050828103606084015261099f81866108ed565b91505067ffffffffffffffff831660808301529695505050505050565b5f602082840312156109cc575f80fd5b505191905056fea2646970667358221220100675b96d0086be6ea0d88c855a736669760178475b004e341f3ec25fc9b6ae64736f6c63430008190033",
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
