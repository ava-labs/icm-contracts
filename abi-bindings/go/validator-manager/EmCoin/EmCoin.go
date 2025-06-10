// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package emcoin

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

// EmCoinMetaData contains all meta data concerning the EmCoin contract.
var EmCoinMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedAddr\",\"type\":\"bytes32\"}],\"name\":\"allowances\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originalAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowanceRemaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"buyCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052633b9aca006002553480156016575f80fd5b506108f1806100245f395ff3fe608060405234801561000f575f80fd5b50600436106100cb575f3560e01c8063521eb273116100885780638779257411610063578063877925741461021957806395d89b411461022e578063a9059cbb1461024d578063dd62ed3e14610260575f80fd5b8063521eb2731461017a5780635e7c9fe81461018357806370a08231146101f1575f80fd5b806306fdde03146100cf578063095ea7b31461010357806318160ddd1461012657806323b872dd1461013a57806327e235e31461014d578063313ce5671461016c575b5f80fd5b60408051808201909152600681526522b6a1b7b4b760d11b60208201525b6040516100fa9190610744565b60405180910390f35b6101166101113660046107ab565b610273565b60405190151581526020016100fa565b633b9aca005b6040519081526020016100fa565b6101166101483660046107d3565b6103cf565b61012c61015b36600461080c565b5f6020819052908152604090205481565b6040515f81526020016100fa565b61012c60025481565b6101c6610191366004610825565b600160208190525f918252604090912080549181015460028201546003909201546001600160a01b0393841693909116919084565b604080516001600160a01b0395861681529490931660208501529183015260608201526080016100fa565b61012c6101ff36600461080c565b6001600160a01b03165f9081526020819052604090205490565b61022c610227366004610825565b610598565b005b604080518082019091526003815262454d4360e81b60208201526100ed565b61011661025b3660046107ab565b6105d6565b61012c61026e36600461083c565b6105ff565b5f8061027f3385610624565b5f8181526001602052604090206002810154600390910154919250900361035157604080516080810182526001600160a01b038681168083523360208085018281528587018a8152606087018b81525f8a81526001808652908a9020985189549089166001600160a01b0319918216178a55935190890180549190981693169290921790955593516002860155925160039094019390935592518681527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a360019150506103c9565b604080516080810182526001600160a01b0380871682523360208084019182525f84860181815260608601828152978252600192839052958120945185549085166001600160a01b03199182161786559251918501805492909416919092161790915591516002820155915160039092019190915590505b92915050565b5f806103db8585610624565b5f818152600160205260408120600301549192500361043a5760405162461bcd60e51b81526020600482015260166024820152754e6f20616c6c6f77616e63652072656d61696e696e6760501b60448201526064015b60405180910390fd5b5f818152600160205260409020546001600160a01b038581169116146104945760405162461bcd60e51b815260206004820152600f60248201526e24b73b30b634b21039b832b73232b960891b6044820152606401610431565b5f81815260016020819052604090912001546001600160a01b038681169116146104f05760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b21037bbb732b960991b6044820152606401610431565b5f8181526001602052604090206003015483111561055c5760405162461bcd60e51b8152602060048201526024808201527f496e76616c696420616d6f756e742c206e6f7420656e6f75676820616c6c6f77604482015263616e636560e01b6064820152608401610431565b5f818152600160205260408120600301805485929061057c908490610881565b9091555061058d905085858561066d565b9150505b9392505050565b335f90815260208190526040812080548392906105b6908490610894565b925050819055508060025f8282546105ce9190610881565b909155505050565b335f908152602081905260408120548211156105f4576105f46108a7565b61059133848461066d565b5f8061060b8484610624565b5f90815260016020526040902060030154949350505050565b6040516bffffffffffffffffffffffff19606084811b8216602084015283901b1660348201525f9060480160405160208183030381529060405280519060200120905092915050565b6001600160a01b0383165f9081526020819052604081205482111561069357505f610591565b6001600160a01b0384165f90815260208190526040812080548492906106ba908490610881565b90915550506001600160a01b0383165f90815260208190526040812080548492906106e6908490610894565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161073291815260200190565b60405180910390a35060019392505050565b5f602080835283518060208501525f5b8181101561077057858101830151858201604001528201610754565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146107a6575f80fd5b919050565b5f80604083850312156107bc575f80fd5b6107c583610790565b946020939093013593505050565b5f805f606084860312156107e5575f80fd5b6107ee84610790565b92506107fc60208501610790565b9150604084013590509250925092565b5f6020828403121561081c575f80fd5b61059182610790565b5f60208284031215610835575f80fd5b5035919050565b5f806040838503121561084d575f80fd5b61085683610790565b915061086460208401610790565b90509250929050565b634e487b7160e01b5f52601160045260245ffd5b818103818111156103c9576103c961086d565b808201808211156103c9576103c961086d565b634e487b7160e01b5f52600160045260245ffdfea2646970667358221220e465e2356c1b5b10607f74d1250a2a71fe2db440d2679332399aff5371c0bab864736f6c63430008190033",
}

// EmCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use EmCoinMetaData.ABI instead.
var EmCoinABI = EmCoinMetaData.ABI

// EmCoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EmCoinMetaData.Bin instead.
var EmCoinBin = EmCoinMetaData.Bin

// DeployEmCoin deploys a new Ethereum contract, binding an instance of EmCoin to it.
func DeployEmCoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EmCoin, error) {
	parsed, err := EmCoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EmCoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EmCoin{EmCoinCaller: EmCoinCaller{contract: contract}, EmCoinTransactor: EmCoinTransactor{contract: contract}, EmCoinFilterer: EmCoinFilterer{contract: contract}}, nil
}

// EmCoin is an auto generated Go binding around an Ethereum contract.
type EmCoin struct {
	EmCoinCaller     // Read-only binding to the contract
	EmCoinTransactor // Write-only binding to the contract
	EmCoinFilterer   // Log filterer for contract events
}

// EmCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmCoinSession struct {
	Contract     *EmCoin           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EmCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmCoinCallerSession struct {
	Contract *EmCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EmCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmCoinTransactorSession struct {
	Contract     *EmCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EmCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmCoinRaw struct {
	Contract *EmCoin // Generic contract binding to access the raw methods on
}

// EmCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmCoinCallerRaw struct {
	Contract *EmCoinCaller // Generic read-only contract binding to access the raw methods on
}

// EmCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmCoinTransactorRaw struct {
	Contract *EmCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmCoin creates a new instance of EmCoin, bound to a specific deployed contract.
func NewEmCoin(address common.Address, backend bind.ContractBackend) (*EmCoin, error) {
	contract, err := bindEmCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EmCoin{EmCoinCaller: EmCoinCaller{contract: contract}, EmCoinTransactor: EmCoinTransactor{contract: contract}, EmCoinFilterer: EmCoinFilterer{contract: contract}}, nil
}

// NewEmCoinCaller creates a new read-only instance of EmCoin, bound to a specific deployed contract.
func NewEmCoinCaller(address common.Address, caller bind.ContractCaller) (*EmCoinCaller, error) {
	contract, err := bindEmCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmCoinCaller{contract: contract}, nil
}

// NewEmCoinTransactor creates a new write-only instance of EmCoin, bound to a specific deployed contract.
func NewEmCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*EmCoinTransactor, error) {
	contract, err := bindEmCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmCoinTransactor{contract: contract}, nil
}

// NewEmCoinFilterer creates a new log filterer instance of EmCoin, bound to a specific deployed contract.
func NewEmCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*EmCoinFilterer, error) {
	contract, err := bindEmCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmCoinFilterer{contract: contract}, nil
}

// bindEmCoin binds a generic wrapper to an already deployed contract.
func bindEmCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EmCoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmCoin *EmCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmCoin.Contract.EmCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmCoin *EmCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmCoin.Contract.EmCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmCoin *EmCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmCoin.Contract.EmCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmCoin *EmCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmCoin *EmCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmCoin *EmCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmCoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_EmCoin *EmCoinCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EmCoin.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_EmCoin *EmCoinSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EmCoin.Contract.Allowance(&_EmCoin.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_EmCoin *EmCoinCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EmCoin.Contract.Allowance(&_EmCoin.CallOpts, owner, spender)
}

// Allowances is a free data retrieval call binding the contract method 0x5e7c9fe8.
//
// Solidity: function allowances(bytes32 _hashedAddr) view returns(address spender, address owner, uint256 originalAllowance, uint256 allowanceRemaining)
func (_EmCoin *EmCoinCaller) Allowances(opts *bind.CallOpts, _hashedAddr [32]byte) (struct {
	Spender            common.Address
	Owner              common.Address
	OriginalAllowance  *big.Int
	AllowanceRemaining *big.Int
}, error) {
	var out []interface{}
	err := _EmCoin.contract.Call(opts, &out, "allowances", _hashedAddr)

	outstruct := new(struct {
		Spender            common.Address
		Owner              common.Address
		OriginalAllowance  *big.Int
		AllowanceRemaining *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Spender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.OriginalAllowance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AllowanceRemaining = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Allowances is a free data retrieval call binding the contract method 0x5e7c9fe8.
//
// Solidity: function allowances(bytes32 _hashedAddr) view returns(address spender, address owner, uint256 originalAllowance, uint256 allowanceRemaining)
func (_EmCoin *EmCoinSession) Allowances(_hashedAddr [32]byte) (struct {
	Spender            common.Address
	Owner              common.Address
	OriginalAllowance  *big.Int
	AllowanceRemaining *big.Int
}, error) {
	return _EmCoin.Contract.Allowances(&_EmCoin.CallOpts, _hashedAddr)
}

// Allowances is a free data retrieval call binding the contract method 0x5e7c9fe8.
//
// Solidity: function allowances(bytes32 _hashedAddr) view returns(address spender, address owner, uint256 originalAllowance, uint256 allowanceRemaining)
func (_EmCoin *EmCoinCallerSession) Allowances(_hashedAddr [32]byte) (struct {
	Spender            common.Address
	Owner              common.Address
	OriginalAllowance  *big.Int
	AllowanceRemaining *big.Int
}, error) {
	return _EmCoin.Contract.Allowances(&_EmCoin.CallOpts, _hashedAddr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_EmCoin *EmCoinCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EmCoin.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_EmCoin *EmCoinSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _EmCoin.Contract.BalanceOf(&_EmCoin.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_EmCoin *EmCoinCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _EmCoin.Contract.BalanceOf(&_EmCoin.CallOpts, owner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_EmCoin *EmCoinCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EmCoin.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_EmCoin *EmCoinSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _EmCoin.Contract.Balances(&_EmCoin.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_EmCoin *EmCoinCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _EmCoin.Contract.Balances(&_EmCoin.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_EmCoin *EmCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EmCoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_EmCoin *EmCoinSession) Decimals() (uint8, error) {
	return _EmCoin.Contract.Decimals(&_EmCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_EmCoin *EmCoinCallerSession) Decimals() (uint8, error) {
	return _EmCoin.Contract.Decimals(&_EmCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_EmCoin *EmCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EmCoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_EmCoin *EmCoinSession) Name() (string, error) {
	return _EmCoin.Contract.Name(&_EmCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_EmCoin *EmCoinCallerSession) Name() (string, error) {
	return _EmCoin.Contract.Name(&_EmCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_EmCoin *EmCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EmCoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_EmCoin *EmCoinSession) Symbol() (string, error) {
	return _EmCoin.Contract.Symbol(&_EmCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_EmCoin *EmCoinCallerSession) Symbol() (string, error) {
	return _EmCoin.Contract.Symbol(&_EmCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_EmCoin *EmCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmCoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_EmCoin *EmCoinSession) TotalSupply() (*big.Int, error) {
	return _EmCoin.Contract.TotalSupply(&_EmCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_EmCoin *EmCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _EmCoin.Contract.TotalSupply(&_EmCoin.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(uint256)
func (_EmCoin *EmCoinCaller) Wallet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmCoin.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(uint256)
func (_EmCoin *EmCoinSession) Wallet() (*big.Int, error) {
	return _EmCoin.Contract.Wallet(&_EmCoin.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(uint256)
func (_EmCoin *EmCoinCallerSession) Wallet() (*big.Int, error) {
	return _EmCoin.Contract.Wallet(&_EmCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_EmCoin *EmCoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _EmCoin.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_EmCoin *EmCoinSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _EmCoin.Contract.Approve(&_EmCoin.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_EmCoin *EmCoinTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _EmCoin.Contract.Approve(&_EmCoin.TransactOpts, spender, value)
}

// BuyCoin is a paid mutator transaction binding the contract method 0x87792574.
//
// Solidity: function buyCoin(uint256 val) returns()
func (_EmCoin *EmCoinTransactor) BuyCoin(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _EmCoin.contract.Transact(opts, "buyCoin", val)
}

// BuyCoin is a paid mutator transaction binding the contract method 0x87792574.
//
// Solidity: function buyCoin(uint256 val) returns()
func (_EmCoin *EmCoinSession) BuyCoin(val *big.Int) (*types.Transaction, error) {
	return _EmCoin.Contract.BuyCoin(&_EmCoin.TransactOpts, val)
}

// BuyCoin is a paid mutator transaction binding the contract method 0x87792574.
//
// Solidity: function buyCoin(uint256 val) returns()
func (_EmCoin *EmCoinTransactorSession) BuyCoin(val *big.Int) (*types.Transaction, error) {
	return _EmCoin.Contract.BuyCoin(&_EmCoin.TransactOpts, val)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_EmCoin *EmCoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EmCoin.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_EmCoin *EmCoinSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EmCoin.Contract.Transfer(&_EmCoin.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_EmCoin *EmCoinTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EmCoin.Contract.Transfer(&_EmCoin.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool success)
func (_EmCoin *EmCoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EmCoin.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool success)
func (_EmCoin *EmCoinSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EmCoin.Contract.TransferFrom(&_EmCoin.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool success)
func (_EmCoin *EmCoinTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EmCoin.Contract.TransferFrom(&_EmCoin.TransactOpts, from, to, value)
}

// EmCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EmCoin contract.
type EmCoinApprovalIterator struct {
	Event *EmCoinApproval // Event containing the contract specifics and raw log

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
func (it *EmCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmCoinApproval)
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
		it.Event = new(EmCoinApproval)
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
func (it *EmCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmCoinApproval represents a Approval event raised by the EmCoin contract.
type EmCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EmCoin *EmCoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EmCoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EmCoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EmCoinApprovalIterator{contract: _EmCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EmCoin *EmCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EmCoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EmCoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmCoinApproval)
				if err := _EmCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_EmCoin *EmCoinFilterer) ParseApproval(log types.Log) (*EmCoinApproval, error) {
	event := new(EmCoinApproval)
	if err := _EmCoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EmCoin contract.
type EmCoinTransferIterator struct {
	Event *EmCoinTransfer // Event containing the contract specifics and raw log

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
func (it *EmCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmCoinTransfer)
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
		it.Event = new(EmCoinTransfer)
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
func (it *EmCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmCoinTransfer represents a Transfer event raised by the EmCoin contract.
type EmCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EmCoin *EmCoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EmCoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EmCoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EmCoinTransferIterator{contract: _EmCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EmCoin *EmCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EmCoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EmCoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmCoinTransfer)
				if err := _EmCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_EmCoin *EmCoinFilterer) ParseTransfer(log types.Log) (*EmCoinTransfer, error) {
	event := new(EmCoinTransfer)
	if err := _EmCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
