// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package examplerewardcalculator

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

// ExampleRewardCalculatorMetaData contains all meta data concerning the ExampleRewardCalculator contract.
var ExampleRewardCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"rewardBasisPoints_\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BIPS_CONVERSION_FACTOR\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPTIME_REWARDS_THRESHOLD_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"validatorStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stakingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stakingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uptimeSeconds\",\"type\":\"uint64\"}],\"name\":\"calculateReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardBasisPoints\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052348015600e575f80fd5b50604051610393380380610393833981016040819052602b91603b565b6001600160401b03166080526066565b5f60208284031215604a575f80fd5b81516001600160401b0381168114605f575f80fd5b9392505050565b60805161030f6100845f395f818160c50152610161015261030f5ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80634f22429f146100595780635dcc93911461007f578063a9778a7a1461008a578063afba878a146100a6578063bb65b242146100c0575b5f80fd5b61006c6100673660046101db565b610100565b6040519081526020015b60405180910390f35b61006c6301e1338081565b61009361271081565b60405161ffff9091168152602001610076565b6100ae605081565b60405160ff9091168152602001610076565b6100e77f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff9091168152602001610076565b5f605061010d8685610249565b6101179190610271565b67ffffffffffffffff1661012c836064610271565b67ffffffffffffffff16101561014357505f6101b6565b6127106301e133806101558686610249565b67ffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1689610195919061029d565b61019f919061029d565b6101a991906102ba565b6101b391906102ba565b90505b95945050505050565b803567ffffffffffffffff811681146101d6575f80fd5b919050565b5f805f805f60a086880312156101ef575f80fd5b853594506101ff602087016101bf565b935061020d604087016101bf565b925061021b606087016101bf565b9150610229608087016101bf565b90509295509295909350565b634e487b7160e01b5f52601160045260245ffd5b67ffffffffffffffff82811682821603908082111561026a5761026a610235565b5092915050565b67ffffffffffffffff81811683821602808216919082811461029557610295610235565b505092915050565b80820281158282048414176102b4576102b4610235565b92915050565b5f826102d457634e487b7160e01b5f52601260045260245ffd5b50049056fea264697066735822122072c44142fee22b25286a3bf642ed0244dded257d3a7e032adb158fd16ebbad0b64736f6c63430008190033",
}

// ExampleRewardCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleRewardCalculatorMetaData.ABI instead.
var ExampleRewardCalculatorABI = ExampleRewardCalculatorMetaData.ABI

// ExampleRewardCalculatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleRewardCalculatorMetaData.Bin instead.
var ExampleRewardCalculatorBin = ExampleRewardCalculatorMetaData.Bin

// DeployExampleRewardCalculator deploys a new Ethereum contract, binding an instance of ExampleRewardCalculator to it.
func DeployExampleRewardCalculator(auth *bind.TransactOpts, backend bind.ContractBackend, rewardBasisPoints_ uint64) (common.Address, *types.Transaction, *ExampleRewardCalculator, error) {
	parsed, err := ExampleRewardCalculatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleRewardCalculatorBin), backend, rewardBasisPoints_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleRewardCalculator{ExampleRewardCalculatorCaller: ExampleRewardCalculatorCaller{contract: contract}, ExampleRewardCalculatorTransactor: ExampleRewardCalculatorTransactor{contract: contract}, ExampleRewardCalculatorFilterer: ExampleRewardCalculatorFilterer{contract: contract}}, nil
}

// ExampleRewardCalculator is an auto generated Go binding around an Ethereum contract.
type ExampleRewardCalculator struct {
	ExampleRewardCalculatorCaller     // Read-only binding to the contract
	ExampleRewardCalculatorTransactor // Write-only binding to the contract
	ExampleRewardCalculatorFilterer   // Log filterer for contract events
}

// ExampleRewardCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleRewardCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleRewardCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleRewardCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleRewardCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleRewardCalculatorSession struct {
	Contract     *ExampleRewardCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ExampleRewardCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleRewardCalculatorCallerSession struct {
	Contract *ExampleRewardCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ExampleRewardCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleRewardCalculatorTransactorSession struct {
	Contract     *ExampleRewardCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ExampleRewardCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleRewardCalculatorRaw struct {
	Contract *ExampleRewardCalculator // Generic contract binding to access the raw methods on
}

// ExampleRewardCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorCallerRaw struct {
	Contract *ExampleRewardCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleRewardCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorTransactorRaw struct {
	Contract *ExampleRewardCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleRewardCalculator creates a new instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculator(address common.Address, backend bind.ContractBackend) (*ExampleRewardCalculator, error) {
	contract, err := bindExampleRewardCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculator{ExampleRewardCalculatorCaller: ExampleRewardCalculatorCaller{contract: contract}, ExampleRewardCalculatorTransactor: ExampleRewardCalculatorTransactor{contract: contract}, ExampleRewardCalculatorFilterer: ExampleRewardCalculatorFilterer{contract: contract}}, nil
}

// NewExampleRewardCalculatorCaller creates a new read-only instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculatorCaller(address common.Address, caller bind.ContractCaller) (*ExampleRewardCalculatorCaller, error) {
	contract, err := bindExampleRewardCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculatorCaller{contract: contract}, nil
}

// NewExampleRewardCalculatorTransactor creates a new write-only instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleRewardCalculatorTransactor, error) {
	contract, err := bindExampleRewardCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculatorTransactor{contract: contract}, nil
}

// NewExampleRewardCalculatorFilterer creates a new log filterer instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleRewardCalculatorFilterer, error) {
	contract, err := bindExampleRewardCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculatorFilterer{contract: contract}, nil
}

// bindExampleRewardCalculator binds a generic wrapper to an already deployed contract.
func bindExampleRewardCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleRewardCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleRewardCalculator *ExampleRewardCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleRewardCalculator.Contract.ExampleRewardCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleRewardCalculator *ExampleRewardCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.ExampleRewardCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleRewardCalculator *ExampleRewardCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.ExampleRewardCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleRewardCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleRewardCalculator *ExampleRewardCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleRewardCalculator *ExampleRewardCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.contract.Transact(opts, method, params...)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ExampleRewardCalculator.Contract.BIPSCONVERSIONFACTOR(&_ExampleRewardCalculator.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ExampleRewardCalculator.Contract.BIPSCONVERSIONFACTOR(&_ExampleRewardCalculator.CallOpts)
}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) SECONDSINYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "SECONDS_IN_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) SECONDSINYEAR() (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.SECONDSINYEAR(&_ExampleRewardCalculator.CallOpts)
}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) SECONDSINYEAR() (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.SECONDSINYEAR(&_ExampleRewardCalculator.CallOpts)
}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) UPTIMEREWARDSTHRESHOLDPERCENTAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "UPTIME_REWARDS_THRESHOLD_PERCENTAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _ExampleRewardCalculator.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_ExampleRewardCalculator.CallOpts)
}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _ExampleRewardCalculator.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_ExampleRewardCalculator.CallOpts)
}

// CalculateReward is a free data retrieval call binding the contract method 0x4f22429f.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 validatorStartTime, uint64 stakingStartTime, uint64 stakingEndTime, uint64 uptimeSeconds) view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) CalculateReward(opts *bind.CallOpts, stakeAmount *big.Int, validatorStartTime uint64, stakingStartTime uint64, stakingEndTime uint64, uptimeSeconds uint64) (*big.Int, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "calculateReward", stakeAmount, validatorStartTime, stakingStartTime, stakingEndTime, uptimeSeconds)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateReward is a free data retrieval call binding the contract method 0x4f22429f.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 validatorStartTime, uint64 stakingStartTime, uint64 stakingEndTime, uint64 uptimeSeconds) view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) CalculateReward(stakeAmount *big.Int, validatorStartTime uint64, stakingStartTime uint64, stakingEndTime uint64, uptimeSeconds uint64) (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.CalculateReward(&_ExampleRewardCalculator.CallOpts, stakeAmount, validatorStartTime, stakingStartTime, stakingEndTime, uptimeSeconds)
}

// CalculateReward is a free data retrieval call binding the contract method 0x4f22429f.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 validatorStartTime, uint64 stakingStartTime, uint64 stakingEndTime, uint64 uptimeSeconds) view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) CalculateReward(stakeAmount *big.Int, validatorStartTime uint64, stakingStartTime uint64, stakingEndTime uint64, uptimeSeconds uint64) (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.CalculateReward(&_ExampleRewardCalculator.CallOpts, stakeAmount, validatorStartTime, stakingStartTime, stakingEndTime, uptimeSeconds)
}

// RewardBasisPoints is a free data retrieval call binding the contract method 0xbb65b242.
//
// Solidity: function rewardBasisPoints() view returns(uint64)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) RewardBasisPoints(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "rewardBasisPoints")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RewardBasisPoints is a free data retrieval call binding the contract method 0xbb65b242.
//
// Solidity: function rewardBasisPoints() view returns(uint64)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) RewardBasisPoints() (uint64, error) {
	return _ExampleRewardCalculator.Contract.RewardBasisPoints(&_ExampleRewardCalculator.CallOpts)
}

// RewardBasisPoints is a free data retrieval call binding the contract method 0xbb65b242.
//
// Solidity: function rewardBasisPoints() view returns(uint64)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) RewardBasisPoints() (uint64, error) {
	return _ExampleRewardCalculator.Contract.RewardBasisPoints(&_ExampleRewardCalculator.CallOpts)
}
