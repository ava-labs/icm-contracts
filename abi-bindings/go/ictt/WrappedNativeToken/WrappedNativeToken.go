// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappednativetoken

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

// WrappedNativeTokenMetaData contains all meta data concerning the WrappedNativeToken contract.
var WrappedNativeTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610c5e380380610c5e83398101604081905261002e916100c7565b8060405160200161003f919061016e565b6040516020818303038152906040528160405160200161005f919061019d565b60408051601f19818403018152919052600361007b8382610249565b5060046100888282610249565b50505050610308565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156100bf5781810151838201526020016100a7565b50505f910152565b5f602082840312156100d7575f80fd5b81516001600160401b03808211156100ed575f80fd5b818401915084601f830112610100575f80fd5b81518181111561011257610112610091565b604051601f8201601f19908116603f0116810190838211818310171561013a5761013a610091565b81604052828152876020848701011115610152575f80fd5b6101638360208301602088016100a5565b979650505050505050565b6702bb930b83832b2160c51b81525f82516101908160088501602087016100a5565b9190910160080192915050565b605760f81b81525f82516101b88160018501602087016100a5565b9190910160010192915050565b600181811c908216806101d957607f821691505b6020821081036101f757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561024457805f5260205f20601f840160051c810160208510156102225750805b601f840160051c820191505b81811015610241575f815560010161022e565b50505b505050565b81516001600160401b0381111561026257610262610091565b6102768161027084546101c5565b846101fd565b602080601f8311600181146102a9575f84156102925750858301515b5f19600386901b1c1916600185901b178555610300565b5f85815260208120601f198616915b828110156102d7578886015182559484019460019091019084016102b8565b50858210156102f457878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b610949806103155f395ff3fe60806040526004361061009f575f3560e01c8063313ce56711610063578063313ce5671461016b57806370a082311461018657806395d89b41146101ba578063a9059cbb146101ce578063d0e30db0146100ae578063dd62ed3e146101ed576100ae565b806306fdde03146100b6578063095ea7b3146100e057806318160ddd1461010f57806323b872dd1461012d5780632e1a7d4d1461014c576100ae565b366100ae576100ac610231565b005b6100ac610231565b3480156100c1575f80fd5b506100ca610272565b6040516100d7919061078c565b60405180910390f35b3480156100eb575f80fd5b506100ff6100fa3660046107f3565b610302565b60405190151581526020016100d7565b34801561011a575f80fd5b506002545b6040519081526020016100d7565b348015610138575f80fd5b506100ff61014736600461081b565b61031b565b348015610157575f80fd5b506100ac610166366004610854565b61033e565b348015610176575f80fd5b50604051601281526020016100d7565b348015610191575f80fd5b5061011f6101a036600461086b565b6001600160a01b03165f9081526020819052604090205490565b3480156101c5575f80fd5b506100ca61038a565b3480156101d9575f80fd5b506100ff6101e83660046107f3565b610399565b3480156101f8575f80fd5b5061011f61020736600461088b565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b61023b33346103a6565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b606060038054610281906108bc565b80601f01602080910402602001604051908101604052809291908181526020018280546102ad906108bc565b80156102f85780601f106102cf576101008083540402835291602001916102f8565b820191905f5260205f20905b8154815290600101906020018083116102db57829003601f168201915b5050505050905090565b5f3361030f8185856103e3565b60019150505b92915050565b5f336103288582856103f5565b610333858585610470565b506001949350505050565b61034833826104cd565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a26103873382610501565b50565b606060048054610281906108bc565b5f3361030f818585610470565b6001600160a01b0382166103d45760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b6103df5f8383610594565b5050565b6103f083838360016106ba565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811461046a578181101561045c57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016103cb565b61046a84848484035f6106ba565b50505050565b6001600160a01b03831661049957604051634b637e8f60e11b81525f60048201526024016103cb565b6001600160a01b0382166104c25760405163ec442f0560e01b81525f60048201526024016103cb565b6103f0838383610594565b6001600160a01b0382166104f657604051634b637e8f60e11b81525f60048201526024016103cb565b6103df825f83610594565b804710156105245760405163cd78605960e01b81523060048201526024016103cb565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f811461056d576040519150601f19603f3d011682016040523d82523d5f602084013e610572565b606091505b50509050806103f057604051630a12f52160e11b815260040160405180910390fd5b6001600160a01b0383166105be578060025f8282546105b391906108f4565b9091555061062e9050565b6001600160a01b0383165f90815260208190526040902054818110156106105760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103cb565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661064a57600280548290039055610668565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516106ad91815260200190565b60405180910390a3505050565b6001600160a01b0384166106e35760405163e602df0560e01b81525f60048201526024016103cb565b6001600160a01b03831661070c57604051634a1406b160e11b81525f60048201526024016103cb565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561046a57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161077e91815260200190565b60405180910390a350505050565b5f602080835283518060208501525f5b818110156107b85785810183015185820160400152820161079c565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146107ee575f80fd5b919050565b5f8060408385031215610804575f80fd5b61080d836107d8565b946020939093013593505050565b5f805f6060848603121561082d575f80fd5b610836846107d8565b9250610844602085016107d8565b9150604084013590509250925092565b5f60208284031215610864575f80fd5b5035919050565b5f6020828403121561087b575f80fd5b610884826107d8565b9392505050565b5f806040838503121561089c575f80fd5b6108a5836107d8565b91506108b3602084016107d8565b90509250929050565b600181811c908216806108d057607f821691505b6020821081036108ee57634e487b7160e01b5f52602260045260245ffd5b50919050565b8082018082111561031557634e487b7160e01b5f52601160045260245ffdfea264697066735822122019feb013315592a9dc19984edd54707e5ab4913905e67cfcd77df649d0d7bfdb64736f6c63430008190033",
}

// WrappedNativeTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappedNativeTokenMetaData.ABI instead.
var WrappedNativeTokenABI = WrappedNativeTokenMetaData.ABI

// WrappedNativeTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WrappedNativeTokenMetaData.Bin instead.
var WrappedNativeTokenBin = WrappedNativeTokenMetaData.Bin

// DeployWrappedNativeToken deploys a new Ethereum contract, binding an instance of WrappedNativeToken to it.
func DeployWrappedNativeToken(auth *bind.TransactOpts, backend bind.ContractBackend, symbol string) (common.Address, *types.Transaction, *WrappedNativeToken, error) {
	parsed, err := WrappedNativeTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrappedNativeTokenBin), backend, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrappedNativeToken{WrappedNativeTokenCaller: WrappedNativeTokenCaller{contract: contract}, WrappedNativeTokenTransactor: WrappedNativeTokenTransactor{contract: contract}, WrappedNativeTokenFilterer: WrappedNativeTokenFilterer{contract: contract}}, nil
}

// WrappedNativeToken is an auto generated Go binding around an Ethereum contract.
type WrappedNativeToken struct {
	WrappedNativeTokenCaller     // Read-only binding to the contract
	WrappedNativeTokenTransactor // Write-only binding to the contract
	WrappedNativeTokenFilterer   // Log filterer for contract events
}

// WrappedNativeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappedNativeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappedNativeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrappedNativeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappedNativeTokenSession struct {
	Contract     *WrappedNativeToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WrappedNativeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappedNativeTokenCallerSession struct {
	Contract *WrappedNativeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// WrappedNativeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappedNativeTokenTransactorSession struct {
	Contract     *WrappedNativeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// WrappedNativeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappedNativeTokenRaw struct {
	Contract *WrappedNativeToken // Generic contract binding to access the raw methods on
}

// WrappedNativeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappedNativeTokenCallerRaw struct {
	Contract *WrappedNativeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// WrappedNativeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappedNativeTokenTransactorRaw struct {
	Contract *WrappedNativeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappedNativeToken creates a new instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeToken(address common.Address, backend bind.ContractBackend) (*WrappedNativeToken, error) {
	contract, err := bindWrappedNativeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeToken{WrappedNativeTokenCaller: WrappedNativeTokenCaller{contract: contract}, WrappedNativeTokenTransactor: WrappedNativeTokenTransactor{contract: contract}, WrappedNativeTokenFilterer: WrappedNativeTokenFilterer{contract: contract}}, nil
}

// NewWrappedNativeTokenCaller creates a new read-only instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeTokenCaller(address common.Address, caller bind.ContractCaller) (*WrappedNativeTokenCaller, error) {
	contract, err := bindWrappedNativeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenCaller{contract: contract}, nil
}

// NewWrappedNativeTokenTransactor creates a new write-only instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappedNativeTokenTransactor, error) {
	contract, err := bindWrappedNativeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenTransactor{contract: contract}, nil
}

// NewWrappedNativeTokenFilterer creates a new log filterer instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappedNativeTokenFilterer, error) {
	contract, err := bindWrappedNativeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenFilterer{contract: contract}, nil
}

// bindWrappedNativeToken binds a generic wrapper to an already deployed contract.
func bindWrappedNativeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WrappedNativeTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedNativeToken *WrappedNativeTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedNativeToken.Contract.WrappedNativeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedNativeToken *WrappedNativeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.WrappedNativeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedNativeToken *WrappedNativeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.WrappedNativeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedNativeToken *WrappedNativeTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedNativeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedNativeToken *WrappedNativeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedNativeToken *WrappedNativeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.Allowance(&_WrappedNativeToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.Allowance(&_WrappedNativeToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.BalanceOf(&_WrappedNativeToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.BalanceOf(&_WrappedNativeToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeToken *WrappedNativeTokenSession) Decimals() (uint8, error) {
	return _WrappedNativeToken.Contract.Decimals(&_WrappedNativeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Decimals() (uint8, error) {
	return _WrappedNativeToken.Contract.Decimals(&_WrappedNativeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenSession) Name() (string, error) {
	return _WrappedNativeToken.Contract.Name(&_WrappedNativeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Name() (string, error) {
	return _WrappedNativeToken.Contract.Name(&_WrappedNativeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenSession) Symbol() (string, error) {
	return _WrappedNativeToken.Contract.Symbol(&_WrappedNativeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Symbol() (string, error) {
	return _WrappedNativeToken.Contract.Symbol(&_WrappedNativeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenSession) TotalSupply() (*big.Int, error) {
	return _WrappedNativeToken.Contract.TotalSupply(&_WrappedNativeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _WrappedNativeToken.Contract.TotalSupply(&_WrappedNativeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Approve(&_WrappedNativeToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Approve(&_WrappedNativeToken.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Deposit() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Deposit(&_WrappedNativeToken.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Deposit() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Deposit(&_WrappedNativeToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Transfer(&_WrappedNativeToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Transfer(&_WrappedNativeToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.TransferFrom(&_WrappedNativeToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.TransferFrom(&_WrappedNativeToken.TransactOpts, from, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Withdraw(&_WrappedNativeToken.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Withdraw(&_WrappedNativeToken.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Fallback(&_WrappedNativeToken.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Fallback(&_WrappedNativeToken.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Receive() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Receive(&_WrappedNativeToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Receive() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Receive(&_WrappedNativeToken.TransactOpts)
}

// WrappedNativeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WrappedNativeToken contract.
type WrappedNativeTokenApprovalIterator struct {
	Event *WrappedNativeTokenApproval // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenApproval)
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
		it.Event = new(WrappedNativeTokenApproval)
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
func (it *WrappedNativeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenApproval represents a Approval event raised by the WrappedNativeToken contract.
type WrappedNativeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WrappedNativeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenApprovalIterator{contract: _WrappedNativeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenApproval)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseApproval(log types.Log) (*WrappedNativeTokenApproval, error) {
	event := new(WrappedNativeTokenApproval)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedNativeTokenDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WrappedNativeToken contract.
type WrappedNativeTokenDepositIterator struct {
	Event *WrappedNativeTokenDeposit // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenDeposit)
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
		it.Event = new(WrappedNativeTokenDeposit)
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
func (it *WrappedNativeTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenDeposit represents a Deposit event raised by the WrappedNativeToken contract.
type WrappedNativeTokenDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*WrappedNativeTokenDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenDepositIterator{contract: _WrappedNativeToken.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenDeposit)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseDeposit(log types.Log) (*WrappedNativeTokenDeposit, error) {
	event := new(WrappedNativeTokenDeposit)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedNativeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WrappedNativeToken contract.
type WrappedNativeTokenTransferIterator struct {
	Event *WrappedNativeTokenTransfer // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenTransfer)
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
		it.Event = new(WrappedNativeTokenTransfer)
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
func (it *WrappedNativeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenTransfer represents a Transfer event raised by the WrappedNativeToken contract.
type WrappedNativeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedNativeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenTransferIterator{contract: _WrappedNativeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenTransfer)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseTransfer(log types.Log) (*WrappedNativeTokenTransfer, error) {
	event := new(WrappedNativeTokenTransfer)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedNativeTokenWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WrappedNativeToken contract.
type WrappedNativeTokenWithdrawalIterator struct {
	Event *WrappedNativeTokenWithdrawal // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenWithdrawal)
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
		it.Event = new(WrappedNativeTokenWithdrawal)
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
func (it *WrappedNativeTokenWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenWithdrawal represents a Withdrawal event raised by the WrappedNativeToken contract.
type WrappedNativeTokenWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*WrappedNativeTokenWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenWithdrawalIterator{contract: _WrappedNativeToken.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenWithdrawal)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseWithdrawal(log types.Log) (*WrappedNativeTokenWithdrawal, error) {
	event := new(WrappedNativeTokenWithdrawal)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
