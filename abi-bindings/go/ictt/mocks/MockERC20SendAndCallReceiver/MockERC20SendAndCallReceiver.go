// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockerc20sendandcallreceiver

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

// MockERC20SendAndCallReceiverMetaData contains all meta data concerning the MockERC20SendAndCallReceiver contract.
var MockERC20SendAndCallReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originTokenTransferrerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"TokensReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"blockSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"blockedSenders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"blocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"receiveTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506107608061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c8063487bd69e146100435780637f450d8d1461008157806394395edd146100be575b5f80fd5b61006d610051366004610589565b5f60208181529281526040808220909352908152205460ff1681565b604051901515815260200160405180910390f35b6100bc61008f366004610589565b5f918252602082815260408084206001600160a01b0390931684529190529020805460ff19166001179055565b005b6100bc6100cc3660046105b3565b5f878152602081815260408083206001600160a01b038916845290915290205460ff16156101565760405162461bcd60e51b815260206004820152602c60248201527f4d6f636b455243323053656e64416e6443616c6c52656365697665723a20736560448201526b1b99195c88189b1bd8dad95960a21b60648201526084015b60405180910390fd5b846001600160a01b0316866001600160a01b0316887f7149eb81ada224b86bfda05a4cf439b25596d3c0d8015aebd1cdc11ab17fe190878787876040516101a09493929190610662565b60405180910390a4806102095760405162461bcd60e51b815260206004820152602b60248201527f4d6f636b455243323053656e64416e6443616c6c52656365697665723a20656d60448201526a1c1d1e481c185e5b1bd85960aa1b606482015260840161014d565b61021484338561021e565b5050505050505050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015610264573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061028891906106a9565b905061029f6001600160a01b038616853086610383565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa1580156102e3573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061030791906106a9565b905081811161036d5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b606482015260840161014d565b61037782826106c0565b925050505b9392505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526103dd9085906103e3565b50505050565b5f6103f76001600160a01b03841683610449565b905080515f1415801561041b57508080602001905181019061041991906106df565b155b1561044457604051635274afe760e01b81526001600160a01b038416600482015260240161014d565b505050565b606061045683835f61045f565b90505b92915050565b6060814710156104845760405163cd78605960e01b815230600482015260240161014d565b5f80856001600160a01b0316848660405161049f91906106fe565b5f6040518083038185875af1925050503d805f81146104d9576040519150601f19603f3d011682016040523d82523d5f602084013e6104de565b606091505b50915091506103778683836060826104fe576104f982610545565b61037c565b815115801561051557506001600160a01b0384163b155b1561053e57604051639996b31560e01b81526001600160a01b038516600482015260240161014d565b508061037c565b8051156105555780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80356001600160a01b0381168114610584575f80fd5b919050565b5f806040838503121561059a575f80fd5b823591506105aa6020840161056e565b90509250929050565b5f805f805f805f60c0888a0312156105c9575f80fd5b873596506105d96020890161056e565b95506105e76040890161056e565b94506105f56060890161056e565b93506080880135925060a088013567ffffffffffffffff80821115610618575f80fd5b818a0191508a601f83011261062b575f80fd5b813581811115610639575f80fd5b8b602082850101111561064a575f80fd5b60208301945080935050505092959891949750929550565b6001600160a01b0385168152602081018490526060604082018190528101829052818360808301375f818301608090810191909152601f909201601f191601019392505050565b5f602082840312156106b9575f80fd5b5051919050565b8181038181111561045957634e487b7160e01b5f52601160045260245ffd5b5f602082840312156106ef575f80fd5b8151801515811461037c575f80fd5b5f82515f5b8181101561071d5760208186018101518583015201610703565b505f92019182525091905056fea264697066735822122095a8e87b0abdbb42c20fdabfcec6036e8858b2d142c5497bc8ccd73f2fcb6e8b64736f6c63430008190033",
}

// MockERC20SendAndCallReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MockERC20SendAndCallReceiverMetaData.ABI instead.
var MockERC20SendAndCallReceiverABI = MockERC20SendAndCallReceiverMetaData.ABI

// MockERC20SendAndCallReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockERC20SendAndCallReceiverMetaData.Bin instead.
var MockERC20SendAndCallReceiverBin = MockERC20SendAndCallReceiverMetaData.Bin

// DeployMockERC20SendAndCallReceiver deploys a new Ethereum contract, binding an instance of MockERC20SendAndCallReceiver to it.
func DeployMockERC20SendAndCallReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockERC20SendAndCallReceiver, error) {
	parsed, err := MockERC20SendAndCallReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockERC20SendAndCallReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockERC20SendAndCallReceiver{MockERC20SendAndCallReceiverCaller: MockERC20SendAndCallReceiverCaller{contract: contract}, MockERC20SendAndCallReceiverTransactor: MockERC20SendAndCallReceiverTransactor{contract: contract}, MockERC20SendAndCallReceiverFilterer: MockERC20SendAndCallReceiverFilterer{contract: contract}}, nil
}

// MockERC20SendAndCallReceiver is an auto generated Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiver struct {
	MockERC20SendAndCallReceiverCaller     // Read-only binding to the contract
	MockERC20SendAndCallReceiverTransactor // Write-only binding to the contract
	MockERC20SendAndCallReceiverFilterer   // Log filterer for contract events
}

// MockERC20SendAndCallReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC20SendAndCallReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC20SendAndCallReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockERC20SendAndCallReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC20SendAndCallReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockERC20SendAndCallReceiverSession struct {
	Contract     *MockERC20SendAndCallReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MockERC20SendAndCallReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockERC20SendAndCallReceiverCallerSession struct {
	Contract *MockERC20SendAndCallReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// MockERC20SendAndCallReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockERC20SendAndCallReceiverTransactorSession struct {
	Contract     *MockERC20SendAndCallReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// MockERC20SendAndCallReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiverRaw struct {
	Contract *MockERC20SendAndCallReceiver // Generic contract binding to access the raw methods on
}

// MockERC20SendAndCallReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiverCallerRaw struct {
	Contract *MockERC20SendAndCallReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MockERC20SendAndCallReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiverTransactorRaw struct {
	Contract *MockERC20SendAndCallReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockERC20SendAndCallReceiver creates a new instance of MockERC20SendAndCallReceiver, bound to a specific deployed contract.
func NewMockERC20SendAndCallReceiver(address common.Address, backend bind.ContractBackend) (*MockERC20SendAndCallReceiver, error) {
	contract, err := bindMockERC20SendAndCallReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiver{MockERC20SendAndCallReceiverCaller: MockERC20SendAndCallReceiverCaller{contract: contract}, MockERC20SendAndCallReceiverTransactor: MockERC20SendAndCallReceiverTransactor{contract: contract}, MockERC20SendAndCallReceiverFilterer: MockERC20SendAndCallReceiverFilterer{contract: contract}}, nil
}

// NewMockERC20SendAndCallReceiverCaller creates a new read-only instance of MockERC20SendAndCallReceiver, bound to a specific deployed contract.
func NewMockERC20SendAndCallReceiverCaller(address common.Address, caller bind.ContractCaller) (*MockERC20SendAndCallReceiverCaller, error) {
	contract, err := bindMockERC20SendAndCallReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiverCaller{contract: contract}, nil
}

// NewMockERC20SendAndCallReceiverTransactor creates a new write-only instance of MockERC20SendAndCallReceiver, bound to a specific deployed contract.
func NewMockERC20SendAndCallReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MockERC20SendAndCallReceiverTransactor, error) {
	contract, err := bindMockERC20SendAndCallReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiverTransactor{contract: contract}, nil
}

// NewMockERC20SendAndCallReceiverFilterer creates a new log filterer instance of MockERC20SendAndCallReceiver, bound to a specific deployed contract.
func NewMockERC20SendAndCallReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MockERC20SendAndCallReceiverFilterer, error) {
	contract, err := bindMockERC20SendAndCallReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiverFilterer{contract: contract}, nil
}

// bindMockERC20SendAndCallReceiver binds a generic wrapper to an already deployed contract.
func bindMockERC20SendAndCallReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockERC20SendAndCallReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockERC20SendAndCallReceiver.Contract.MockERC20SendAndCallReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.MockERC20SendAndCallReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.MockERC20SendAndCallReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockERC20SendAndCallReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.contract.Transact(opts, method, params...)
}

// BlockedSenders is a free data retrieval call binding the contract method 0x487bd69e.
//
// Solidity: function blockedSenders(bytes32 blockchainID, address senderAddress) view returns(bool blocked)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverCaller) BlockedSenders(opts *bind.CallOpts, blockchainID [32]byte, senderAddress common.Address) (bool, error) {
	var out []interface{}
	err := _MockERC20SendAndCallReceiver.contract.Call(opts, &out, "blockedSenders", blockchainID, senderAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlockedSenders is a free data retrieval call binding the contract method 0x487bd69e.
//
// Solidity: function blockedSenders(bytes32 blockchainID, address senderAddress) view returns(bool blocked)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverSession) BlockedSenders(blockchainID [32]byte, senderAddress common.Address) (bool, error) {
	return _MockERC20SendAndCallReceiver.Contract.BlockedSenders(&_MockERC20SendAndCallReceiver.CallOpts, blockchainID, senderAddress)
}

// BlockedSenders is a free data retrieval call binding the contract method 0x487bd69e.
//
// Solidity: function blockedSenders(bytes32 blockchainID, address senderAddress) view returns(bool blocked)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverCallerSession) BlockedSenders(blockchainID [32]byte, senderAddress common.Address) (bool, error) {
	return _MockERC20SendAndCallReceiver.Contract.BlockedSenders(&_MockERC20SendAndCallReceiver.CallOpts, blockchainID, senderAddress)
}

// BlockSender is a paid mutator transaction binding the contract method 0x7f450d8d.
//
// Solidity: function blockSender(bytes32 blockchainID, address senderAddress) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactor) BlockSender(opts *bind.TransactOpts, blockchainID [32]byte, senderAddress common.Address) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.contract.Transact(opts, "blockSender", blockchainID, senderAddress)
}

// BlockSender is a paid mutator transaction binding the contract method 0x7f450d8d.
//
// Solidity: function blockSender(bytes32 blockchainID, address senderAddress) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverSession) BlockSender(blockchainID [32]byte, senderAddress common.Address) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.BlockSender(&_MockERC20SendAndCallReceiver.TransactOpts, blockchainID, senderAddress)
}

// BlockSender is a paid mutator transaction binding the contract method 0x7f450d8d.
//
// Solidity: function blockSender(bytes32 blockchainID, address senderAddress) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactorSession) BlockSender(blockchainID [32]byte, senderAddress common.Address) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.BlockSender(&_MockERC20SendAndCallReceiver.TransactOpts, blockchainID, senderAddress)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x94395edd.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originTokenTransferrerAddress, address originSenderAddress, address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactor) ReceiveTokens(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originTokenTransferrerAddress common.Address, originSenderAddress common.Address, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.contract.Transact(opts, "receiveTokens", sourceBlockchainID, originTokenTransferrerAddress, originSenderAddress, token, amount, payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x94395edd.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originTokenTransferrerAddress, address originSenderAddress, address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverSession) ReceiveTokens(sourceBlockchainID [32]byte, originTokenTransferrerAddress common.Address, originSenderAddress common.Address, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.ReceiveTokens(&_MockERC20SendAndCallReceiver.TransactOpts, sourceBlockchainID, originTokenTransferrerAddress, originSenderAddress, token, amount, payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x94395edd.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originTokenTransferrerAddress, address originSenderAddress, address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactorSession) ReceiveTokens(sourceBlockchainID [32]byte, originTokenTransferrerAddress common.Address, originSenderAddress common.Address, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.ReceiveTokens(&_MockERC20SendAndCallReceiver.TransactOpts, sourceBlockchainID, originTokenTransferrerAddress, originSenderAddress, token, amount, payload)
}

// MockERC20SendAndCallReceiverTokensReceivedIterator is returned from FilterTokensReceived and is used to iterate over the raw logs and unpacked data for TokensReceived events raised by the MockERC20SendAndCallReceiver contract.
type MockERC20SendAndCallReceiverTokensReceivedIterator struct {
	Event *MockERC20SendAndCallReceiverTokensReceived // Event containing the contract specifics and raw log

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
func (it *MockERC20SendAndCallReceiverTokensReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockERC20SendAndCallReceiverTokensReceived)
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
		it.Event = new(MockERC20SendAndCallReceiverTokensReceived)
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
func (it *MockERC20SendAndCallReceiverTokensReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockERC20SendAndCallReceiverTokensReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockERC20SendAndCallReceiverTokensReceived represents a TokensReceived event raised by the MockERC20SendAndCallReceiver contract.
type MockERC20SendAndCallReceiverTokensReceived struct {
	SourceBlockchainID            [32]byte
	OriginTokenTransferrerAddress common.Address
	OriginSenderAddress           common.Address
	Token                         common.Address
	Amount                        *big.Int
	Payload                       []byte
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterTokensReceived is a free log retrieval operation binding the contract event 0x7149eb81ada224b86bfda05a4cf439b25596d3c0d8015aebd1cdc11ab17fe190.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originTokenTransferrerAddress, address indexed originSenderAddress, address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) FilterTokensReceived(opts *bind.FilterOpts, sourceBlockchainID [][32]byte, originTokenTransferrerAddress []common.Address, originSenderAddress []common.Address) (*MockERC20SendAndCallReceiverTokensReceivedIterator, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originTokenTransferrerAddressRule []interface{}
	for _, originTokenTransferrerAddressItem := range originTokenTransferrerAddress {
		originTokenTransferrerAddressRule = append(originTokenTransferrerAddressRule, originTokenTransferrerAddressItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _MockERC20SendAndCallReceiver.contract.FilterLogs(opts, "TokensReceived", sourceBlockchainIDRule, originTokenTransferrerAddressRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiverTokensReceivedIterator{contract: _MockERC20SendAndCallReceiver.contract, event: "TokensReceived", logs: logs, sub: sub}, nil
}

// WatchTokensReceived is a free log subscription operation binding the contract event 0x7149eb81ada224b86bfda05a4cf439b25596d3c0d8015aebd1cdc11ab17fe190.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originTokenTransferrerAddress, address indexed originSenderAddress, address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) WatchTokensReceived(opts *bind.WatchOpts, sink chan<- *MockERC20SendAndCallReceiverTokensReceived, sourceBlockchainID [][32]byte, originTokenTransferrerAddress []common.Address, originSenderAddress []common.Address) (event.Subscription, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originTokenTransferrerAddressRule []interface{}
	for _, originTokenTransferrerAddressItem := range originTokenTransferrerAddress {
		originTokenTransferrerAddressRule = append(originTokenTransferrerAddressRule, originTokenTransferrerAddressItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _MockERC20SendAndCallReceiver.contract.WatchLogs(opts, "TokensReceived", sourceBlockchainIDRule, originTokenTransferrerAddressRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockERC20SendAndCallReceiverTokensReceived)
				if err := _MockERC20SendAndCallReceiver.contract.UnpackLog(event, "TokensReceived", log); err != nil {
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

// ParseTokensReceived is a log parse operation binding the contract event 0x7149eb81ada224b86bfda05a4cf439b25596d3c0d8015aebd1cdc11ab17fe190.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originTokenTransferrerAddress, address indexed originSenderAddress, address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) ParseTokensReceived(log types.Log) (*MockERC20SendAndCallReceiverTokensReceived, error) {
	event := new(MockERC20SendAndCallReceiverTokensReceived)
	if err := _MockERC20SendAndCallReceiver.contract.UnpackLog(event, "TokensReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
