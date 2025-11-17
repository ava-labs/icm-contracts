// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethwarp

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ICMMessage is an auto generated low-level Go binding around an user-defined struct.
type ICMMessage struct {
	UnsignedMessage      ICMUnsignedMessage
	UnsignedMessageBytes []byte
	Signature            ICMSignature
}

// ICMSignature is an auto generated low-level Go binding around an user-defined struct.
type ICMSignature struct {
	Signers   []byte
	Signature []byte
}

// ICMUnsignedMessage is an auto generated low-level Go binding around an user-defined struct.
type ICMUnsignedMessage struct {
	AvalancheNetworkID          uint32
	AvalancheSourceBlockchainID [32]byte
	Payload                     []byte
}

// WarpBlockHash is an auto generated low-level Go binding around an user-defined struct.
type WarpBlockHash struct {
	SourceChainID [32]byte
	BlockHash     [32]byte
}

// WarpMessage is an auto generated low-level Go binding around an user-defined struct.
type WarpMessage struct {
	SourceChainID       [32]byte
	OriginSenderAddress common.Address
	Payload             []byte
}

// EthWarpMetaData contains all meta data concerning the EthWarp contract.
var EthWarpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockChainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SendWarpMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"avalancheNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"avalancheSourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structICMUnsignedMessage\",\"name\":\"unsignedMessage\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"unsignedMessageBytes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signers\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structICMSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structICMMessage\",\"name\":\"icmMessage\",\"type\":\"tuple\"}],\"name\":\"getVerifiedICMMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structWarpMessage\",\"name\":\"warpMessage\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getVerifiedWarpBlockHash\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structWarpBlockHash\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getVerifiedWarpMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structWarpMessage\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avaBlockchainId\",\"type\":\"bytes32\"}],\"name\":\"isChainRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"avaBlockchainId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"verifyWarpMessage\",\"type\":\"address\"}],\"name\":\"registerChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sendWarpMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608034604d57601f610db238819003918201601f19168301916001600160401b03831184841017605157808492602094604052833981010312604d57515f55604051610d4c90816100668239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f905f3560e01c9081634213cf7814610b6e575080635eb13d5c14610a7a5780636f825350146109e057806385a88ca1146109a6578063b1b731c8146101c7578063ce7f592914610156578063d127dc9b146101395763ee5b48eb14610076575f80fd5b3461012e57602036600319011261012e5760043567ffffffffffffffff8111610135573660238201121561013557806004013567ffffffffffffffff8111610131573691016024011161012e5760405162461bcd60e51b815260206004820152603e60248201527f53656e64696e672057617270206d657373616765732066726f6d20457468657260448201527f65756d206973206e6f742063757272656e746c7920737570706f7274656400006064820152608490fd5b80fd5b8280fd5b5080fd5b503461012e578060031936011261012e5760209054604051908152f35b503461012e57602036600319011261012e57610170610b87565b5060405162461bcd60e51b815260206004820152602860248201527f54686973206d6574686f642063616e6e6f742062652063616c6c6564206f6e20604482015267457468657265756d60c01b6064820152608490fd5b3461089a57602036600319011261089a5760043567ffffffffffffffff811161089a57806004019080360390606060031983011261089a57610207610c6b565b5061023560206102178580610c8a565b01355f908152600160205260409020546001600160a01b0316151590565b156109295760206102468480610c8a565b01355f52600160205260018060a01b0360405f20541691604051926339f5645160e01b8452602060048501528435606219830181121561089a576102d06102bf856102ef930160446004820191606060248b015263ffffffff6102a884610b9a565b1660848b0152602481013560a48b01520190610bab565b606060c489015260e4880191610bdd565b6102dd6024860188610bab565b87830360231901604489015290610bdd565b6044840135926042190183121561089a57846103538194926020968394019061034660048301602319868403016064870152602461033e6103308380610bab565b604087526040870191610bdd565b940190610bab565b9189818503910152610bdd565b03915afa90811561091e575f916108e3575b501561089e578061037591610c8a565b60608136031261089a5760405161038b81610bfd565b61039482610b9a565b815260208101916020810135835260408101359067ffffffffffffffff821161089a570136601f8201121561089a578035916103cf83610c9f565b926103dd6040519485610c49565b808452366020828501011161089a576020815f92826040960183880137850101520190815261040a610c6b565b5051906060602060405161041d81610c2d565b82815201528151156108865760208201516001600160f81b0319161580610866575b1561082e5781519060068092106106bb5761045a6004610c9f565b6104676040519182610c49565b60048152601f196104786004610c9f565b013660208301375f5b600481106107f55750610495600191610ccc565b60e01c036107b057825191600a8093106106bb576104b36004610c9f565b906104c16040519283610c49565b60048252601f196104d26004610c9f565b013660208401375f5b600481106107795750506104ee90610ccc565b60e01c918351838201908183116107005781116106bb5761050e84610c9f565b9161051c6040519384610c49565b848352601f1961052b86610c9f565b013660208501375f5b85811061074257505063ffffffff81116107005763ffffffff168451600482019081831161070057106106bb5761056b6004610c9f565b906105796040519283610c49565b60048252601f1961058a6004610c9f565b013660208401375f5b600481106107145750506105a690610ccc565b60e01c92600e0163ffffffff81116107005763ffffffff169284516105cb8286610d09565b116106bb576105d981610c9f565b936105e76040519586610c49565b818552601f196105f683610c9f565b013660208701375f5b828110610687578560a086826020886040519061061b82610c2d565b81520152516040519061062d82610bfd565b81526020808201935f8552604083019081526040519485938385525183850152600180861b039051166040840152516060808401528051918291826080860152018484015e5f828201840152601f01601f19168101030190f35b6001906001600160f81b03196106a66106a08386610d09565b8a610cbb565b51165f1a6106b48289610cbb565b53016105ff565b60405162461bcd60e51b815260206004820152601960248201527f42797465536c696365723a206f7574206f6620626f756e6473000000000000006044820152606490fd5b634e487b7160e01b5f52601160045260245ffd5b6001906001600160f81b031961072d6106a08386610d09565b51165f1a61073b8286610cbb565b5301610593565b80820190818311610700576001916001600160f81b031990610764908a610cbb565b51165f1a6107728287610cbb565b5301610534565b80820190818311610700576001916001600160f81b03199061079b9089610cbb565b51165f1a6107a98286610cbb565b53016104db565b60405162461bcd60e51b815260206004820152601760248201527f496e76616c6964207061796c6f616420747970652049440000000000000000006044820152606490fd5b806002019081600211610700576001916001600160f81b0319906108199088610cbb565b51165f1a6108278285610cbb565b5301610481565b60405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a590818dbd91958c8125160821b6044820152606490fd5b508151600110156108865760218201516001600160f81b0319161561043f565b634e487b7160e01b5f52603260045260245ffd5b5f80fd5b60405162461bcd60e51b815260206004820152601f60248201527f526563656976656420616e20696e76616c69642049434d206d657373616765006044820152606490fd5b90506020813d602011610916575b816108fe60209383610c49565b8101031261089a5751801515810361089a5782610365565b3d91506108f1565b6040513d5f823e3d90fd5b60405162461bcd60e51b815260206004820152604960248201527f43616e6e6f74207265636569766520612057617270206d65737361676520667260448201527f6f6d206120636861696e2077686f73652076616c696461746f7220736574206960648201526839903ab735b737bbb760b91b608482015260a490fd5b3461089a57602036600319011261089a576004355f908152600160209081526040909120546001600160a01b031615156040519015158152f35b3461089a57602036600319011261089a576109f9610b87565b5060405162461bcd60e51b815260206004820152604c60248201527f54686973206d6574686f642063616e27742062652063616c6c6564206f6e204560448201527f7468657265756d2c207573652060676574566572696669656449434d4d65737360648201526b1859d958081a5b9cdd19585960a21b608482015260a490fd5b3461089a57604036600319011261089a576024356001600160a01b038116906004359082900361089a575f818152600160205260409020546001600160a01b0316610b2a578115610ae5575f90815260016020526040902080546001600160a01b0319169091179055005b60405162461bcd60e51b815260206004820152601f60248201527f50726f7669646564206164647265737320646f6573206e6f74206578697374006044820152606490fd5b606460405162461bcd60e51b815260206004820152602060248201527f5468697320636861696e20697320616c726561647920726567697374657265646044820152fd5b3461089a575f36600319011261089a576020905f548152f35b6004359063ffffffff8216820361089a57565b359063ffffffff8216820361089a57565b9035601e198236030181121561089a57016020813591019167ffffffffffffffff821161089a57813603831361089a57565b908060209392818452848401375f828201840152601f01601f1916010190565b6060810190811067ffffffffffffffff821117610c1957604052565b634e487b7160e01b5f52604160045260245ffd5b6040810190811067ffffffffffffffff821117610c1957604052565b90601f8019910116810190811067ffffffffffffffff821117610c1957604052565b60405190610c7882610bfd565b60606040835f81525f60208201520152565b903590605e198136030182121561089a570190565b67ffffffffffffffff8111610c1957601f01601f191660200190565b908151811015610886570160200190565b80516020909101516001600160e01b0319811692919060048210610cee575050565b6001600160e01b031960049290920360031b82901b16169150565b919082018092116107005756fea264697066735822122091dc6ea75b65535d1ae8d089c4ef749c513e835a65e63b3e4944c01a23ca112964736f6c634300081e0033",
}

// EthWarpABI is the input ABI used to generate the binding from.
// Deprecated: Use EthWarpMetaData.ABI instead.
var EthWarpABI = EthWarpMetaData.ABI

// EthWarpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthWarpMetaData.Bin instead.
var EthWarpBin = EthWarpMetaData.Bin

// DeployEthWarp deploys a new Ethereum contract, binding an instance of EthWarp to it.
func DeployEthWarp(auth *bind.TransactOpts, backend bind.ContractBackend, blockChainId *big.Int) (common.Address, *types.Transaction, *EthWarp, error) {
	parsed, err := EthWarpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthWarpBin), backend, blockChainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthWarp{EthWarpCaller: EthWarpCaller{contract: contract}, EthWarpTransactor: EthWarpTransactor{contract: contract}, EthWarpFilterer: EthWarpFilterer{contract: contract}}, nil
}

// EthWarp is an auto generated Go binding around an Ethereum contract.
type EthWarp struct {
	EthWarpCaller     // Read-only binding to the contract
	EthWarpTransactor // Write-only binding to the contract
	EthWarpFilterer   // Log filterer for contract events
}

// EthWarpCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthWarpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthWarpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthWarpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthWarpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthWarpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthWarpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthWarpSession struct {
	Contract     *EthWarp          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthWarpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthWarpCallerSession struct {
	Contract *EthWarpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EthWarpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthWarpTransactorSession struct {
	Contract     *EthWarpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EthWarpRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthWarpRaw struct {
	Contract *EthWarp // Generic contract binding to access the raw methods on
}

// EthWarpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthWarpCallerRaw struct {
	Contract *EthWarpCaller // Generic read-only contract binding to access the raw methods on
}

// EthWarpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthWarpTransactorRaw struct {
	Contract *EthWarpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthWarp creates a new instance of EthWarp, bound to a specific deployed contract.
func NewEthWarp(address common.Address, backend bind.ContractBackend) (*EthWarp, error) {
	contract, err := bindEthWarp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthWarp{EthWarpCaller: EthWarpCaller{contract: contract}, EthWarpTransactor: EthWarpTransactor{contract: contract}, EthWarpFilterer: EthWarpFilterer{contract: contract}}, nil
}

// NewEthWarpCaller creates a new read-only instance of EthWarp, bound to a specific deployed contract.
func NewEthWarpCaller(address common.Address, caller bind.ContractCaller) (*EthWarpCaller, error) {
	contract, err := bindEthWarp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthWarpCaller{contract: contract}, nil
}

// NewEthWarpTransactor creates a new write-only instance of EthWarp, bound to a specific deployed contract.
func NewEthWarpTransactor(address common.Address, transactor bind.ContractTransactor) (*EthWarpTransactor, error) {
	contract, err := bindEthWarp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthWarpTransactor{contract: contract}, nil
}

// NewEthWarpFilterer creates a new log filterer instance of EthWarp, bound to a specific deployed contract.
func NewEthWarpFilterer(address common.Address, filterer bind.ContractFilterer) (*EthWarpFilterer, error) {
	contract, err := bindEthWarp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthWarpFilterer{contract: contract}, nil
}

// bindEthWarp binds a generic wrapper to an already deployed contract.
func bindEthWarp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthWarpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthWarp *EthWarpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthWarp.Contract.EthWarpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthWarp *EthWarpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthWarp.Contract.EthWarpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthWarp *EthWarpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthWarp.Contract.EthWarpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthWarp *EthWarpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthWarp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthWarp *EthWarpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthWarp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthWarp *EthWarpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthWarp.Contract.contract.Transact(opts, method, params...)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_EthWarp *EthWarpCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthWarp.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_EthWarp *EthWarpSession) BlockchainID() ([32]byte, error) {
	return _EthWarp.Contract.BlockchainID(&_EthWarp.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_EthWarp *EthWarpCallerSession) BlockchainID() ([32]byte, error) {
	return _EthWarp.Contract.BlockchainID(&_EthWarp.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_EthWarp *EthWarpCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthWarp.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_EthWarp *EthWarpSession) GetBlockchainID() ([32]byte, error) {
	return _EthWarp.Contract.GetBlockchainID(&_EthWarp.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_EthWarp *EthWarpCallerSession) GetBlockchainID() ([32]byte, error) {
	return _EthWarp.Contract.GetBlockchainID(&_EthWarp.CallOpts)
}

// GetVerifiedICMMessage is a free data retrieval call binding the contract method 0xb1b731c8.
//
// Solidity: function getVerifiedICMMessage(((uint32,bytes32,bytes),bytes,(bytes,bytes)) icmMessage) view returns((bytes32,address,bytes) warpMessage)
func (_EthWarp *EthWarpCaller) GetVerifiedICMMessage(opts *bind.CallOpts, icmMessage ICMMessage) (WarpMessage, error) {
	var out []interface{}
	err := _EthWarp.contract.Call(opts, &out, "getVerifiedICMMessage", icmMessage)

	if err != nil {
		return *new(WarpMessage), err
	}

	out0 := *abi.ConvertType(out[0], new(WarpMessage)).(*WarpMessage)

	return out0, err

}

// GetVerifiedICMMessage is a free data retrieval call binding the contract method 0xb1b731c8.
//
// Solidity: function getVerifiedICMMessage(((uint32,bytes32,bytes),bytes,(bytes,bytes)) icmMessage) view returns((bytes32,address,bytes) warpMessage)
func (_EthWarp *EthWarpSession) GetVerifiedICMMessage(icmMessage ICMMessage) (WarpMessage, error) {
	return _EthWarp.Contract.GetVerifiedICMMessage(&_EthWarp.CallOpts, icmMessage)
}

// GetVerifiedICMMessage is a free data retrieval call binding the contract method 0xb1b731c8.
//
// Solidity: function getVerifiedICMMessage(((uint32,bytes32,bytes),bytes,(bytes,bytes)) icmMessage) view returns((bytes32,address,bytes) warpMessage)
func (_EthWarp *EthWarpCallerSession) GetVerifiedICMMessage(icmMessage ICMMessage) (WarpMessage, error) {
	return _EthWarp.Contract.GetVerifiedICMMessage(&_EthWarp.CallOpts, icmMessage)
}

// GetVerifiedWarpBlockHash is a free data retrieval call binding the contract method 0xce7f5929.
//
// Solidity: function getVerifiedWarpBlockHash(uint32 index) pure returns((bytes32,bytes32), bool)
func (_EthWarp *EthWarpCaller) GetVerifiedWarpBlockHash(opts *bind.CallOpts, index uint32) (WarpBlockHash, bool, error) {
	var out []interface{}
	err := _EthWarp.contract.Call(opts, &out, "getVerifiedWarpBlockHash", index)

	if err != nil {
		return *new(WarpBlockHash), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(WarpBlockHash)).(*WarpBlockHash)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetVerifiedWarpBlockHash is a free data retrieval call binding the contract method 0xce7f5929.
//
// Solidity: function getVerifiedWarpBlockHash(uint32 index) pure returns((bytes32,bytes32), bool)
func (_EthWarp *EthWarpSession) GetVerifiedWarpBlockHash(index uint32) (WarpBlockHash, bool, error) {
	return _EthWarp.Contract.GetVerifiedWarpBlockHash(&_EthWarp.CallOpts, index)
}

// GetVerifiedWarpBlockHash is a free data retrieval call binding the contract method 0xce7f5929.
//
// Solidity: function getVerifiedWarpBlockHash(uint32 index) pure returns((bytes32,bytes32), bool)
func (_EthWarp *EthWarpCallerSession) GetVerifiedWarpBlockHash(index uint32) (WarpBlockHash, bool, error) {
	return _EthWarp.Contract.GetVerifiedWarpBlockHash(&_EthWarp.CallOpts, index)
}

// GetVerifiedWarpMessage is a free data retrieval call binding the contract method 0x6f825350.
//
// Solidity: function getVerifiedWarpMessage(uint32 index) pure returns((bytes32,address,bytes), bool)
func (_EthWarp *EthWarpCaller) GetVerifiedWarpMessage(opts *bind.CallOpts, index uint32) (WarpMessage, bool, error) {
	var out []interface{}
	err := _EthWarp.contract.Call(opts, &out, "getVerifiedWarpMessage", index)

	if err != nil {
		return *new(WarpMessage), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(WarpMessage)).(*WarpMessage)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetVerifiedWarpMessage is a free data retrieval call binding the contract method 0x6f825350.
//
// Solidity: function getVerifiedWarpMessage(uint32 index) pure returns((bytes32,address,bytes), bool)
func (_EthWarp *EthWarpSession) GetVerifiedWarpMessage(index uint32) (WarpMessage, bool, error) {
	return _EthWarp.Contract.GetVerifiedWarpMessage(&_EthWarp.CallOpts, index)
}

// GetVerifiedWarpMessage is a free data retrieval call binding the contract method 0x6f825350.
//
// Solidity: function getVerifiedWarpMessage(uint32 index) pure returns((bytes32,address,bytes), bool)
func (_EthWarp *EthWarpCallerSession) GetVerifiedWarpMessage(index uint32) (WarpMessage, bool, error) {
	return _EthWarp.Contract.GetVerifiedWarpMessage(&_EthWarp.CallOpts, index)
}

// IsChainRegistered is a free data retrieval call binding the contract method 0x85a88ca1.
//
// Solidity: function isChainRegistered(bytes32 avaBlockchainId) view returns(bool)
func (_EthWarp *EthWarpCaller) IsChainRegistered(opts *bind.CallOpts, avaBlockchainId [32]byte) (bool, error) {
	var out []interface{}
	err := _EthWarp.contract.Call(opts, &out, "isChainRegistered", avaBlockchainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChainRegistered is a free data retrieval call binding the contract method 0x85a88ca1.
//
// Solidity: function isChainRegistered(bytes32 avaBlockchainId) view returns(bool)
func (_EthWarp *EthWarpSession) IsChainRegistered(avaBlockchainId [32]byte) (bool, error) {
	return _EthWarp.Contract.IsChainRegistered(&_EthWarp.CallOpts, avaBlockchainId)
}

// IsChainRegistered is a free data retrieval call binding the contract method 0x85a88ca1.
//
// Solidity: function isChainRegistered(bytes32 avaBlockchainId) view returns(bool)
func (_EthWarp *EthWarpCallerSession) IsChainRegistered(avaBlockchainId [32]byte) (bool, error) {
	return _EthWarp.Contract.IsChainRegistered(&_EthWarp.CallOpts, avaBlockchainId)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x5eb13d5c.
//
// Solidity: function registerChain(bytes32 avaBlockchainId, address verifyWarpMessage) returns()
func (_EthWarp *EthWarpTransactor) RegisterChain(opts *bind.TransactOpts, avaBlockchainId [32]byte, verifyWarpMessage common.Address) (*types.Transaction, error) {
	return _EthWarp.contract.Transact(opts, "registerChain", avaBlockchainId, verifyWarpMessage)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x5eb13d5c.
//
// Solidity: function registerChain(bytes32 avaBlockchainId, address verifyWarpMessage) returns()
func (_EthWarp *EthWarpSession) RegisterChain(avaBlockchainId [32]byte, verifyWarpMessage common.Address) (*types.Transaction, error) {
	return _EthWarp.Contract.RegisterChain(&_EthWarp.TransactOpts, avaBlockchainId, verifyWarpMessage)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x5eb13d5c.
//
// Solidity: function registerChain(bytes32 avaBlockchainId, address verifyWarpMessage) returns()
func (_EthWarp *EthWarpTransactorSession) RegisterChain(avaBlockchainId [32]byte, verifyWarpMessage common.Address) (*types.Transaction, error) {
	return _EthWarp.Contract.RegisterChain(&_EthWarp.TransactOpts, avaBlockchainId, verifyWarpMessage)
}

// SendWarpMessage is a paid mutator transaction binding the contract method 0xee5b48eb.
//
// Solidity: function sendWarpMessage(bytes payload) returns(bytes32)
func (_EthWarp *EthWarpTransactor) SendWarpMessage(opts *bind.TransactOpts, payload []byte) (*types.Transaction, error) {
	return _EthWarp.contract.Transact(opts, "sendWarpMessage", payload)
}

// SendWarpMessage is a paid mutator transaction binding the contract method 0xee5b48eb.
//
// Solidity: function sendWarpMessage(bytes payload) returns(bytes32)
func (_EthWarp *EthWarpSession) SendWarpMessage(payload []byte) (*types.Transaction, error) {
	return _EthWarp.Contract.SendWarpMessage(&_EthWarp.TransactOpts, payload)
}

// SendWarpMessage is a paid mutator transaction binding the contract method 0xee5b48eb.
//
// Solidity: function sendWarpMessage(bytes payload) returns(bytes32)
func (_EthWarp *EthWarpTransactorSession) SendWarpMessage(payload []byte) (*types.Transaction, error) {
	return _EthWarp.Contract.SendWarpMessage(&_EthWarp.TransactOpts, payload)
}

// EthWarpSendWarpMessageIterator is returned from FilterSendWarpMessage and is used to iterate over the raw logs and unpacked data for SendWarpMessage events raised by the EthWarp contract.
type EthWarpSendWarpMessageIterator struct {
	Event *EthWarpSendWarpMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthWarpSendWarpMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthWarpSendWarpMessage)
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
		it.Event = new(EthWarpSendWarpMessage)
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
func (it *EthWarpSendWarpMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthWarpSendWarpMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthWarpSendWarpMessage represents a SendWarpMessage event raised by the EthWarp contract.
type EthWarpSendWarpMessage struct {
	Sender    common.Address
	MessageID [32]byte
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSendWarpMessage is a free log retrieval operation binding the contract event 0x56600c567728a800c0aa927500f831cb451df66a7af570eb4df4dfbf4674887d.
//
// Solidity: event SendWarpMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_EthWarp *EthWarpFilterer) FilterSendWarpMessage(opts *bind.FilterOpts, sender []common.Address, messageID [][32]byte) (*EthWarpSendWarpMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _EthWarp.contract.FilterLogs(opts, "SendWarpMessage", senderRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &EthWarpSendWarpMessageIterator{contract: _EthWarp.contract, event: "SendWarpMessage", logs: logs, sub: sub}, nil
}

// WatchSendWarpMessage is a free log subscription operation binding the contract event 0x56600c567728a800c0aa927500f831cb451df66a7af570eb4df4dfbf4674887d.
//
// Solidity: event SendWarpMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_EthWarp *EthWarpFilterer) WatchSendWarpMessage(opts *bind.WatchOpts, sink chan<- *EthWarpSendWarpMessage, sender []common.Address, messageID [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _EthWarp.contract.WatchLogs(opts, "SendWarpMessage", senderRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthWarpSendWarpMessage)
				if err := _EthWarp.contract.UnpackLog(event, "SendWarpMessage", log); err != nil {
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

// ParseSendWarpMessage is a log parse operation binding the contract event 0x56600c567728a800c0aa927500f831cb451df66a7af570eb4df4dfbf4674887d.
//
// Solidity: event SendWarpMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_EthWarp *EthWarpFilterer) ParseSendWarpMessage(log types.Log) (*EthWarpSendWarpMessage, error) {
	event := new(EthWarpSendWarpMessage)
	if err := _EthWarp.contract.UnpackLog(event, "SendWarpMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
