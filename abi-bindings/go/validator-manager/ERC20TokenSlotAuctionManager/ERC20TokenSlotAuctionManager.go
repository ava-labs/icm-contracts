// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20tokenslotauctionmanager

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

// ERC20TokenSlotAuctionManagerMetaData contains all meta data concerning the ERC20TokenSlotAuctionManager contract.
var ERC20TokenSlotAuctionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20Address\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"}],\"name\":\"AuctionEndTimeNotReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionFinalizing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionNotInProgress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minumumBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBid\",\"type\":\"uint256\"}],\"name\":\"BidSmallerThanMinimum\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"DuplicateBidInContention\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"smallestAcceptableBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBid\",\"type\":\"uint256\"}],\"name\":\"InsufficientBidToWinAuction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeIsValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeIsWinning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationTimeLimit\",\"type\":\"uint256\"}],\"name\":\"ValidatorTimeLimitNotPassed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"BidEvicted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitiatedAuctionValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"validatorSlots\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minValidatorDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"}],\"name\":\"NewValidatorAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"SuccessfulBidPlaced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MinBidRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinValidatorDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinimumBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"bidderInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20\",\"outputs\":[{\"internalType\":\"contractIERC20Mintable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"validatorslots\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minAuctionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValidatorDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"}],\"name\":\"initiateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSlots\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50604051612416380380612416833981016040819052602b91607d565b600c80546001600160a01b0319166001600160a01b03928316179055600180546001600160a81b0319169290911691909117905560a9565b80516001600160a01b03811681146078575f80fd5b919050565b5f8060408385031215608d575f80fd5b6094836063565b915060a0602084016063565b90509250929050565b612360806100b65f395ff3fe608060405234801561000f575f80fd5b5060043610610110575f3560e01c8063913ed0b11161009e578063ae9483e01161006e578063ae9483e014610254578063b6e6a2ca14610267578063da4312a41461027a578063f556f60a1461028d578063fe67a54b14610296575f80fd5b8063913ed0b1146102135780639681d9401461021b578063a3a65e481461022e578063a476f67514610241575f80fd5b806336339388116100e457806336339388146101a55780634b449cba146101cf5780636fe4d1e1146101e6578063785e9e86146101ef5780637a7df93214610200575f80fd5b8062d841f81461011457806305af82561461014a5780630d5daf3b1461015f5780631077830a14610184575b5f80fd5b60035461012d906201000090046001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b61015d610158366004611c15565b61029e565b005b61017261016d366004611c2e565b610313565b60405161014196959493929190611cf8565b6003546101929061ffff1681565b60405161ffff9091168152602001610141565b5f546101b7906001600160a01b031681565b6040516001600160a01b039091168152602001610141565b6101d860025481565b604051908152602001610141565b6101d860045481565b600c546001600160a01b03166101b7565b61015d61020e366004611d65565b610548565b6101d8610699565b6101d8610229366004611c15565b61070d565b6101d861023c366004611c15565b610785565b61015d61024f366004611c2e565b6107bb565b6001546101b7906001600160a01b031681565b61015d610275366004611c2e565b610818565b61015d610288366004611f6e565b61096c565b6101d860055481565b61015d610997565b60015460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af11580156102eb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061030f919061201c565b5050565b60096020525f90815260409020805460018201546002830180546001600160a01b0390931693919261034490612033565b80601f016020809104026020016040519081016040528092919081815260200182805461037090612033565b80156103bb5780601f10610392576101008083540402835291602001916103bb565b820191905f5260205f20905b81548152906001019060200180831161039e57829003601f168201915b5050505050908060030180546103d090612033565b80601f01602080910402602001604051908101604052809291908181526020018280546103fc90612033565b80156104475780601f1061041e57610100808354040283529160200191610447565b820191905f5260205f20905b81548152906001019060200180831161042a57829003601f168201915b505060408051808201825260048701805463ffffffff1682526005880180548451602082810282018101909652818152989998939750919550838701945091928301828280156104be57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116104a0575b50505091909252505060408051808201825260068501805463ffffffff1682526007860180548451602082810282018101909652818152969796939550919380860193929083018282801561053a57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161051c575b505050505081525050905086565b60018054600160a01b900460ff1660028111156105675761056761206b565b0361058557604051639bbb1f3b60e01b815260040160405180910390fd5b6002600154600160a01b900460ff1660028111156105a5576105a561206b565b036105c357604051631f86de2560e21b815260040160405180910390fd5b6001805460ff60a01b1916600160a01b1790556003805469ffffffffffffffff00001916620100006001600160401b038716021790556106038342612093565b6002556003805461ffff191661ffff8716179055600482905560058190555f600a81905561063390600b90611b31565b6003546005546040805161ffff90931683526001600160401b038716602084015282018490526060820185905260808201527f50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b549060a00160405180910390a15050505050565b5f60018054600160a01b900460ff1660028111156106b9576106b961206b565b146106d6576040516232f00d60e71b815260040160405180910390fd5b60035461ffff166106e6600b5490565b10156106f3575060055490565b6106fd600b61100f565b610708906001612093565b905090565b60015460405163025a076560e61b815263ffffffff831660048201525f916001600160a01b031690639681d940906024015b6020604051808303815f875af115801561075b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061077f919061201c565b92915050565b600154604051631474cbc960e31b815263ffffffff831660048201525f916001600160a01b03169063a3a65e489060240161073f565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015b5f604051808303815f87803b1580156107ff575f80fd5b505af1158015610811573d5f803e3d5ffd5b5050505050565b5f81815260086020526040902060010154421015610864575f8181526008602052604090819020600101549051631a936b7b60e11b815260048101919091526024015b60405180910390fd5b600760085f8381526020019081526020015f2060020160405161088791906120a6565b90815260405190819003602001902080546001600160a01b03191681555f600182018190556108b96002830182611b4f565b6108c6600383015f611b4f565b505f600482018190556005909101805467ffffffffffffffff1916905581815260086020526040812080546001600160a01b031916815560018101829055906109126002830182611b4f565b61091f600383015f611b4f565b505f6004828101919091556005909101805467ffffffffffffffff19169055600154604051635b73516560e11b81529182018390526001600160a01b03169063b6e6a2ca906024016107e8565b610974611034565b610981858585858561106b565b61081160015f8051602061230b83398151915255565b61099f611034565b60018054600160a01b900460ff1660028111156109be576109be61206b565b146109db576040516232f00d60e71b815260040160405180910390fd5b6001805460ff60a01b1916600160a11b1790556002544211610a1657600254604051638230260760e01b815260040161085b91815260200190565b5f600255600a54158015610a2b5750600b5415155b15610a3e57610a3a600b61100f565b600a555b5f610a48600b5490565b1115610fe2575f610a59600b6115ab565b5f818152600960209081526040808320815160c08101835281546001600160a01b0316815260018201549381019390935260028101805495965093949293909291840191610aa690612033565b80601f0160208091040260200160405190810160405280929190818152602001828054610ad290612033565b8015610b1d5780601f10610af457610100808354040283529160200191610b1d565b820191905f5260205f20905b815481529060010190602001808311610b0057829003601f168201915b50505050508152602001600382018054610b3690612033565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6290612033565b8015610bad5780601f10610b8457610100808354040283529160200191610bad565b820191905f5260205f20905b815481529060010190602001808311610b9057829003601f168201915b505050918352505060408051808201825260048401805463ffffffff168252600585018054845160208281028201810190965281815295850195939492938584019390929190830182828015610c2a57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610c0c575b50505091909252505050815260408051808201825260068401805463ffffffff168252600785018054845160208281028201810190965281815295850195939492938584019390929190830182828015610cab57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610c8d575b505050919092525050509052508051600a54919250610cd391610cce9085612118565b6115b8565b6001546040828101516060840151608085015160a08601516003549451634e5bb12760e11b81525f966001600160a01b031695639cb7624e95610d309590949093909290916201000090046001600160401b03169060040161212b565b6020604051808303815f875af1158015610d4c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d70919061201c565b9050815f01516001600160a01b0316817f032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285600254600454610db19190612093565b60035460408051928352620100009091046001600160401b031660208301520160405180910390a36040518060c00160405280835f01516001600160a01b03168152602001600254600454610e069190612093565b81526040848101805160208401526060808701518385015283018590526003546201000090046001600160401b031660809093019290925290519051600791610e4e91612193565b9081526040805160209281900383019020835181546001600160a01b0319166001600160a01b039091161781559183015160018301558201516002820190610e9690826121f2565b5060608201516003820190610eab90826121f2565b506080820151816004015560a0820151816005015f6101000a8154816001600160401b0302191690836001600160401b031602179055509050506040518060c00160405280835f01516001600160a01b03168152602001600254600454610f129190612093565b81526040848101516020808401919091526060808701518385015283018590526003546201000090046001600160401b03166080909301929092525f84815260088352819020835181546001600160a01b0319166001600160a01b039091161781559183015160018301558201516002820190610f8f90826121f2565b5060608201516003820190610fa490826121f2565b506080820151600482015560a0909101516005909101805467ffffffffffffffff19166001600160401b039092169190911790555050600a55610a3e565b6003805461ffff191690555f6004556001805460ff60a01b191681555f8051602061230b83398151915255565b5f815f015f81548110611024576110246122b1565b905f5260205f2001549050919050565b5f8051602061230b83398151915280546001190161106557604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b611073611034565b60018054600160a01b900460ff1660028111156110925761109261206b565b146110af576040516232f00d60e71b815260040160405180910390fd5b60015460405163d47a948b60e01b81526001600160a01b039091169063d47a948b906110df9087906004016122c5565b602060405180830381865afa1580156110fa573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061111e919061201c565b1561113e578360405163f3c815f760e01b815260040161085b91906122c5565b60068460405161114e9190612193565b9081526040519081900360200190205460ff1615611181578360405163f3c1d91160e01b815260040161085b91906122c5565b5f858152600960205260409020546001600160a01b0316156111b957604051630517e2e760e21b81526004810186905260240161085b565b8460055411156111ea576005546040516301a0013b60e61b815260048101919091526024810186905260440161085b565b60035461ffff166111fa600b5490565b1015611288575f546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015611253573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061127791906122d7565b50611283600b866115cf565b61141f565b84611293600b61100f565b10156113e6576112a2856115dc565b505f6112af600b876115f4565b600a8190555f818152600960205260409081902090519192506112d7916002909101906120a6565b6040519081900381209082907fdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc905f90a35f81815260096020526040902054611329906001600160a01b0316826115b8565b600660095f8381526020019081526020015f2060020160405161134c91906120a6565b9081526040805160209281900383019020805460ff191690555f8381526009909252812080546001600160a01b031916815560018101829055906113936002830182611b4f565b6113a0600383015f611b4f565b60048201805463ffffffff191681555f6113bd6005850182611b31565b505060068201805463ffffffff191681555f6113dc6007850182611b31565b505050505061141f565b6113f0600b61100f565b6113fb906001612093565b604051631a79674760e11b815260048101919091526024810186905260440161085b565b6040805160c0810182523381526020808201888152828401888152606084018890526080840187905260a084018690525f8a8152600990935293909120825181546001600160a01b0319166001600160a01b03909116178155905160018201559151909190600282019061149390826121f2565b50606082015160038201906114a890826121f2565b506080820151805160048301805463ffffffff191663ffffffff90921691909117815560208083015180516114e39260058701920190611b86565b50505060a0820151805160068301805463ffffffff191663ffffffff90921691909117815560208083015180516115209260078701920190611b86565b50505090505060016006856040516115389190612193565b908152604051908190036020018120805492151560ff1990931692909217909155611564908590612193565b6040519081900381209086907f864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029905f90a361081160015f8051602061230b83398151915255565b5f61077f8261160961160d565b600c5461030f906001600160a01b0316838361169e565b61030f8282611609611702565b600c545f9061077f906001600160a01b031683611733565b5f611602838361160961173f565b9392505050565b1190565b5f80611617845490565b9050805f0361162a5761162a603161178c565b5f611635858261179d565b5490505f611646865f19850161179d565b54865490915086908061165b5761165b6122f6565b600190038181905f5260205f20015f90559055806116845f885f0161179d90919063ffffffff16565b55611695865f1985015f84896117c4565b50949350505050565b6040516001600160a01b038381166024830152604482018390526116fd91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506118b9565b505050565b5f61170b845490565b84546001810186555f8681526020902001849055905061172d84828585611925565b50505050565b5f61160283338461196d565b5f80611749855490565b9050805f0361175c5761175c603161178c565b5f611767868261179d565b54905084611775875f61179d565b5561178386835f88886117c4565b95945050505050565b634e487b715f52806020526024601cfd5b5f826117bc6117b9846117b5845f9081526020902090565b0190565b90565b949350505050565b6001600160ff1b038310156108115760028381026001810191018581101561186d575f6117f1888461179d565b5490505f6117ff898461179d565b549050611810828763ffffffff8816565b80611824575061182481878763ffffffff16565b15611866575f61184a858561183d86868b63ffffffff16565b1515918118919091021890565b90506118578a8983611ad0565b6118648a8a838a8a6117c4565b505b50506118b0565b858210156118b0575f611880888461179d565b549050611891818663ffffffff8716565b156118ae576118a1888785611ad0565b6118ae88888588886117c4565b505b50505050505050565b5f8060205f8451602086015f885af1806118d8576040513d5f823e3d81fd5b50505f513d915081156118ef5780600114156118fc565b6001600160a01b0384163b155b1561172d57604051635274afe760e01b81526001600160a01b038516600482015260240161085b565b821561172d5760025f198401045f61193d868361179d565b54905061194e818563ffffffff8616565b1561195a57505061172d565b611965868684611ad0565b509250611925565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa1580156119b3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119d7919061201c565b90506119ee6001600160a01b038616853086611af8565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015611a32573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a56919061201c565b9050818111611abc5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b606482015260840161085b565b611ac68282612118565b9695505050505050565b5f611adb848461179d565b90505f611ae8858461179d565b8054835490915590915550505050565b6040516001600160a01b03848116602483015283811660448301526064820183905261172d9186918216906323b872dd906084016116cb565b5080545f8255905f5260205f2090810190611b4c9190611be9565b50565b508054611b5b90612033565b5f825580601f10611b6a575050565b601f0160209004905f5260205f2090810190611b4c9190611be9565b828054828255905f5260205f20908101928215611bd9579160200282015b82811115611bd957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611ba4565b50611be5929150611be9565b5090565b5b80821115611be5575f8155600101611bea565b803563ffffffff81168114611c10575f80fd5b919050565b5f60208284031215611c25575f80fd5b61160282611bfd565b5f60208284031215611c3e575f80fd5b5035919050565b5f5b83811015611c5f578181015183820152602001611c47565b50505f910152565b5f8151808452611c7e816020860160208601611c45565b601f01601f19169290920160200192915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611ced5784516001600160a01b03168252938301936001929092019190830190611cc4565b509695505050505050565b60018060a01b038716815285602082015260c060408201525f611d1e60c0830187611c67565b8281036060840152611d308187611c67565b90508281036080840152611d448186611c92565b905082810360a0840152611d588185611c92565b9998505050505050505050565b5f805f805f60a08688031215611d79575f80fd5b853561ffff81168114611d8a575f80fd5b945060208601356001600160401b0381168114611da5575f80fd5b94979496505050506040830135926060810135926080909101359150565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611df957611df9611dc3565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611e2757611e27611dc3565b604052919050565b5f82601f830112611e3e575f80fd5b81356001600160401b03811115611e5757611e57611dc3565b611e6a601f8201601f1916602001611dff565b818152846020838601011115611e7e575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60408284031215611eaa575f80fd5b611eb2611dd7565b9050611ebd82611bfd565b81526020808301356001600160401b0380821115611ed9575f80fd5b818501915085601f830112611eec575f80fd5b813581811115611efe57611efe611dc3565b8060051b9150611f0f848301611dff565b8181529183018401918481019088841115611f28575f80fd5b938501935b83851015611f5d57843592506001600160a01b0383168314611f4d575f80fd5b8282529385019390850190611f2d565b808688015250505050505092915050565b5f805f805f60a08688031215611f82575f80fd5b8535945060208601356001600160401b0380821115611f9f575f80fd5b611fab89838a01611e2f565b95506040880135915080821115611fc0575f80fd5b611fcc89838a01611e2f565b94506060880135915080821115611fe1575f80fd5b611fed89838a01611e9a565b93506080880135915080821115612002575f80fd5b5061200f88828901611e9a565b9150509295509295909350565b5f6020828403121561202c575f80fd5b5051919050565b600181811c9082168061204757607f821691505b60208210810361206557634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b8082018082111561077f5761077f61207f565b5f8083546120b381612033565b600182811680156120cb57600181146120e05761210c565b60ff198416875282151583028701945061210c565b875f526020805f205f5b858110156121035781548a8201529084019082016120ea565b50505082870194505b50929695505050505050565b8181038181111561077f5761077f61207f565b60a081525f61213d60a0830188611c67565b828103602084015261214f8188611c67565b905082810360408401526121638187611c92565b905082810360608401526121778186611c92565b9150506001600160401b03831660808301529695505050505050565b5f82516121a4818460208701611c45565b9190910192915050565b601f8211156116fd57805f5260205f20601f840160051c810160208510156121d35750805b601f840160051c820191505b81811015610811575f81556001016121df565b81516001600160401b0381111561220b5761220b611dc3565b61221f816122198454612033565b846121ae565b602080601f831160018114612252575f841561223b5750858301515b5f19600386901b1c1916600185901b1785556122a9565b5f85815260208120601f198616915b8281101561228057888601518255948401946001909101908401612261565b508582101561229d57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b634e487b7160e01b5f52603260045260245ffd5b602081525f6116026020830184611c67565b5f602082840312156122e7575f80fd5b81518015158114611602575f80fd5b634e487b7160e01b5f52603160045260245ffdfe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220bbec4bd28a7a6fc92b2a2b4fd1cc020ecb02e1382b791d27f55cc1b16a0177da64736f6c63430008190033",
}

// ERC20TokenSlotAuctionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenSlotAuctionManagerMetaData.ABI instead.
var ERC20TokenSlotAuctionManagerABI = ERC20TokenSlotAuctionManagerMetaData.ABI

// ERC20TokenSlotAuctionManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenSlotAuctionManagerMetaData.Bin instead.
var ERC20TokenSlotAuctionManagerBin = ERC20TokenSlotAuctionManagerMetaData.Bin

// DeployERC20TokenSlotAuctionManager deploys a new Ethereum contract, binding an instance of ERC20TokenSlotAuctionManager to it.
func DeployERC20TokenSlotAuctionManager(auth *bind.TransactOpts, backend bind.ContractBackend, vmAddress common.Address, erc20Address common.Address) (common.Address, *types.Transaction, *ERC20TokenSlotAuctionManager, error) {
	parsed, err := ERC20TokenSlotAuctionManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenSlotAuctionManagerBin), backend, vmAddress, erc20Address)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenSlotAuctionManager{ERC20TokenSlotAuctionManagerCaller: ERC20TokenSlotAuctionManagerCaller{contract: contract}, ERC20TokenSlotAuctionManagerTransactor: ERC20TokenSlotAuctionManagerTransactor{contract: contract}, ERC20TokenSlotAuctionManagerFilterer: ERC20TokenSlotAuctionManagerFilterer{contract: contract}}, nil
}

// ERC20TokenSlotAuctionManager is an auto generated Go binding around an Ethereum contract.
type ERC20TokenSlotAuctionManager struct {
	ERC20TokenSlotAuctionManagerCaller     // Read-only binding to the contract
	ERC20TokenSlotAuctionManagerTransactor // Write-only binding to the contract
	ERC20TokenSlotAuctionManagerFilterer   // Log filterer for contract events
}

// ERC20TokenSlotAuctionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenSlotAuctionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSlotAuctionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenSlotAuctionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSlotAuctionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenSlotAuctionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSlotAuctionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenSlotAuctionManagerSession struct {
	Contract     *ERC20TokenSlotAuctionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ERC20TokenSlotAuctionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenSlotAuctionManagerCallerSession struct {
	Contract *ERC20TokenSlotAuctionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// ERC20TokenSlotAuctionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenSlotAuctionManagerTransactorSession struct {
	Contract     *ERC20TokenSlotAuctionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// ERC20TokenSlotAuctionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenSlotAuctionManagerRaw struct {
	Contract *ERC20TokenSlotAuctionManager // Generic contract binding to access the raw methods on
}

// ERC20TokenSlotAuctionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenSlotAuctionManagerCallerRaw struct {
	Contract *ERC20TokenSlotAuctionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenSlotAuctionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenSlotAuctionManagerTransactorRaw struct {
	Contract *ERC20TokenSlotAuctionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenSlotAuctionManager creates a new instance of ERC20TokenSlotAuctionManager, bound to a specific deployed contract.
func NewERC20TokenSlotAuctionManager(address common.Address, backend bind.ContractBackend) (*ERC20TokenSlotAuctionManager, error) {
	contract, err := bindERC20TokenSlotAuctionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSlotAuctionManager{ERC20TokenSlotAuctionManagerCaller: ERC20TokenSlotAuctionManagerCaller{contract: contract}, ERC20TokenSlotAuctionManagerTransactor: ERC20TokenSlotAuctionManagerTransactor{contract: contract}, ERC20TokenSlotAuctionManagerFilterer: ERC20TokenSlotAuctionManagerFilterer{contract: contract}}, nil
}

// NewERC20TokenSlotAuctionManagerCaller creates a new read-only instance of ERC20TokenSlotAuctionManager, bound to a specific deployed contract.
func NewERC20TokenSlotAuctionManagerCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenSlotAuctionManagerCaller, error) {
	contract, err := bindERC20TokenSlotAuctionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSlotAuctionManagerCaller{contract: contract}, nil
}

// NewERC20TokenSlotAuctionManagerTransactor creates a new write-only instance of ERC20TokenSlotAuctionManager, bound to a specific deployed contract.
func NewERC20TokenSlotAuctionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenSlotAuctionManagerTransactor, error) {
	contract, err := bindERC20TokenSlotAuctionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSlotAuctionManagerTransactor{contract: contract}, nil
}

// NewERC20TokenSlotAuctionManagerFilterer creates a new log filterer instance of ERC20TokenSlotAuctionManager, bound to a specific deployed contract.
func NewERC20TokenSlotAuctionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenSlotAuctionManagerFilterer, error) {
	contract, err := bindERC20TokenSlotAuctionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSlotAuctionManagerFilterer{contract: contract}, nil
}

// bindERC20TokenSlotAuctionManager binds a generic wrapper to an already deployed contract.
func bindERC20TokenSlotAuctionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenSlotAuctionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenSlotAuctionManager.Contract.ERC20TokenSlotAuctionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.ERC20TokenSlotAuctionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.ERC20TokenSlotAuctionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenSlotAuctionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.contract.Transact(opts, method, params...)
}

// MinBidRequired is a free data retrieval call binding the contract method 0x913ed0b1.
//
// Solidity: function MinBidRequired() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) MinBidRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "MinBidRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBidRequired is a free data retrieval call binding the contract method 0x913ed0b1.
//
// Solidity: function MinBidRequired() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) MinBidRequired() (*big.Int, error) {
	return _ERC20TokenSlotAuctionManager.Contract.MinBidRequired(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// MinBidRequired is a free data retrieval call binding the contract method 0x913ed0b1.
//
// Solidity: function MinBidRequired() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) MinBidRequired() (*big.Int, error) {
	return _ERC20TokenSlotAuctionManager.Contract.MinBidRequired(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// MinValidatorDuration is a free data retrieval call binding the contract method 0x6fe4d1e1.
//
// Solidity: function MinValidatorDuration() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) MinValidatorDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "MinValidatorDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValidatorDuration is a free data retrieval call binding the contract method 0x6fe4d1e1.
//
// Solidity: function MinValidatorDuration() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) MinValidatorDuration() (*big.Int, error) {
	return _ERC20TokenSlotAuctionManager.Contract.MinValidatorDuration(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// MinValidatorDuration is a free data retrieval call binding the contract method 0x6fe4d1e1.
//
// Solidity: function MinValidatorDuration() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) MinValidatorDuration() (*big.Int, error) {
	return _ERC20TokenSlotAuctionManager.Contract.MinValidatorDuration(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// MinimumBid is a free data retrieval call binding the contract method 0xf556f60a.
//
// Solidity: function MinimumBid() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) MinimumBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "MinimumBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumBid is a free data retrieval call binding the contract method 0xf556f60a.
//
// Solidity: function MinimumBid() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) MinimumBid() (*big.Int, error) {
	return _ERC20TokenSlotAuctionManager.Contract.MinimumBid(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// MinimumBid is a free data retrieval call binding the contract method 0xf556f60a.
//
// Solidity: function MinimumBid() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) MinimumBid() (*big.Int, error) {
	return _ERC20TokenSlotAuctionManager.Contract.MinimumBid(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// TOKENCONTRACT is a free data retrieval call binding the contract method 0x36339388.
//
// Solidity: function TOKEN_CONTRACT() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) TOKENCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "TOKEN_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENCONTRACT is a free data retrieval call binding the contract method 0x36339388.
//
// Solidity: function TOKEN_CONTRACT() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) TOKENCONTRACT() (common.Address, error) {
	return _ERC20TokenSlotAuctionManager.Contract.TOKENCONTRACT(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// TOKENCONTRACT is a free data retrieval call binding the contract method 0x36339388.
//
// Solidity: function TOKEN_CONTRACT() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) TOKENCONTRACT() (common.Address, error) {
	return _ERC20TokenSlotAuctionManager.Contract.TOKENCONTRACT(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// VALIDATORMANAGER is a free data retrieval call binding the contract method 0xae9483e0.
//
// Solidity: function VALIDATOR_MANAGER() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) VALIDATORMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "VALIDATOR_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORMANAGER is a free data retrieval call binding the contract method 0xae9483e0.
//
// Solidity: function VALIDATOR_MANAGER() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) VALIDATORMANAGER() (common.Address, error) {
	return _ERC20TokenSlotAuctionManager.Contract.VALIDATORMANAGER(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// VALIDATORMANAGER is a free data retrieval call binding the contract method 0xae9483e0.
//
// Solidity: function VALIDATOR_MANAGER() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) VALIDATORMANAGER() (common.Address, error) {
	return _ERC20TokenSlotAuctionManager.Contract.VALIDATORMANAGER(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) AuctionEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "auctionEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) AuctionEndTime() (*big.Int, error) {
	return _ERC20TokenSlotAuctionManager.Contract.AuctionEndTime(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) AuctionEndTime() (*big.Int, error) {
	return _ERC20TokenSlotAuctionManager.Contract.AuctionEndTime(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// BidderInfo is a free data retrieval call binding the contract method 0x0d5daf3b.
//
// Solidity: function bidderInfo(uint256 bid) view returns(address addr, uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) BidderInfo(opts *bind.CallOpts, bid *big.Int) (struct {
	Addr                  common.Address
	Bid                   *big.Int
	NodeID                []byte
	BlsPublicKey          []byte
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "bidderInfo", bid)

	outstruct := new(struct {
		Addr                  common.Address
		Bid                   *big.Int
		NodeID                []byte
		BlsPublicKey          []byte
		RemainingBalanceOwner PChainOwner
		DisableOwner          PChainOwner
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Bid = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NodeID = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.BlsPublicKey = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.RemainingBalanceOwner = *abi.ConvertType(out[4], new(PChainOwner)).(*PChainOwner)
	outstruct.DisableOwner = *abi.ConvertType(out[5], new(PChainOwner)).(*PChainOwner)

	return *outstruct, err

}

// BidderInfo is a free data retrieval call binding the contract method 0x0d5daf3b.
//
// Solidity: function bidderInfo(uint256 bid) view returns(address addr, uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) BidderInfo(bid *big.Int) (struct {
	Addr                  common.Address
	Bid                   *big.Int
	NodeID                []byte
	BlsPublicKey          []byte
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}, error) {
	return _ERC20TokenSlotAuctionManager.Contract.BidderInfo(&_ERC20TokenSlotAuctionManager.CallOpts, bid)
}

// BidderInfo is a free data retrieval call binding the contract method 0x0d5daf3b.
//
// Solidity: function bidderInfo(uint256 bid) view returns(address addr, uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) BidderInfo(bid *big.Int) (struct {
	Addr                  common.Address
	Bid                   *big.Int
	NodeID                []byte
	BlsPublicKey          []byte
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}, error) {
	return _ERC20TokenSlotAuctionManager.Contract.BidderInfo(&_ERC20TokenSlotAuctionManager.CallOpts, bid)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) Erc20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "erc20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) Erc20() (common.Address, error) {
	return _ERC20TokenSlotAuctionManager.Contract.Erc20(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) Erc20() (common.Address, error) {
	return _ERC20TokenSlotAuctionManager.Contract.Erc20(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// ValidatorSlots is a free data retrieval call binding the contract method 0x1077830a.
//
// Solidity: function validatorSlots() view returns(uint16)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) ValidatorSlots(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "validatorSlots")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ValidatorSlots is a free data retrieval call binding the contract method 0x1077830a.
//
// Solidity: function validatorSlots() view returns(uint16)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) ValidatorSlots() (uint16, error) {
	return _ERC20TokenSlotAuctionManager.Contract.ValidatorSlots(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// ValidatorSlots is a free data retrieval call binding the contract method 0x1077830a.
//
// Solidity: function validatorSlots() view returns(uint16)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) ValidatorSlots() (uint16, error) {
	return _ERC20TokenSlotAuctionManager.Contract.ValidatorSlots(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// ValidatorWeight is a free data retrieval call binding the contract method 0x00d841f8.
//
// Solidity: function validatorWeight() view returns(uint64)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) ValidatorWeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "validatorWeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValidatorWeight is a free data retrieval call binding the contract method 0x00d841f8.
//
// Solidity: function validatorWeight() view returns(uint64)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) ValidatorWeight() (uint64, error) {
	return _ERC20TokenSlotAuctionManager.Contract.ValidatorWeight(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// ValidatorWeight is a free data retrieval call binding the contract method 0x00d841f8.
//
// Solidity: function validatorWeight() view returns(uint64)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) ValidatorWeight() (uint64, error) {
	return _ERC20TokenSlotAuctionManager.Contract.ValidatorWeight(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactor) CompleteRemoveInitialValidator(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.contract.Transact(opts, "completeRemoveInitialValidator", messageIndex)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) CompleteRemoveInitialValidator(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.CompleteRemoveInitialValidator(&_ERC20TokenSlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactorSession) CompleteRemoveInitialValidator(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.CompleteRemoveInitialValidator(&_ERC20TokenSlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.CompleteValidatorRegistration(&_ERC20TokenSlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.CompleteValidatorRegistration(&_ERC20TokenSlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.CompleteValidatorRemoval(&_ERC20TokenSlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.CompleteValidatorRemoval(&_ERC20TokenSlotAuctionManager.TransactOpts, messageIndex)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) EndAuction() (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.EndAuction(&_ERC20TokenSlotAuctionManager.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactorSession) EndAuction() (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.EndAuction(&_ERC20TokenSlotAuctionManager.TransactOpts)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x7a7df932.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 minAuctionDuration, uint256 minValidatorDuration, uint256 minimumBid) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactor) InitiateAuction(opts *bind.TransactOpts, validatorslots uint16, weight uint64, minAuctionDuration *big.Int, minValidatorDuration *big.Int, minimumBid *big.Int) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.contract.Transact(opts, "initiateAuction", validatorslots, weight, minAuctionDuration, minValidatorDuration, minimumBid)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x7a7df932.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 minAuctionDuration, uint256 minValidatorDuration, uint256 minimumBid) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) InitiateAuction(validatorslots uint16, weight uint64, minAuctionDuration *big.Int, minValidatorDuration *big.Int, minimumBid *big.Int) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.InitiateAuction(&_ERC20TokenSlotAuctionManager.TransactOpts, validatorslots, weight, minAuctionDuration, minValidatorDuration, minimumBid)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x7a7df932.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 minAuctionDuration, uint256 minValidatorDuration, uint256 minimumBid) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactorSession) InitiateAuction(validatorslots uint16, weight uint64, minAuctionDuration *big.Int, minValidatorDuration *big.Int, minimumBid *big.Int) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.InitiateAuction(&_ERC20TokenSlotAuctionManager.TransactOpts, validatorslots, weight, minAuctionDuration, minValidatorDuration, minimumBid)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactor) InitiateRemoveInitialValidator(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.contract.Transact(opts, "initiateRemoveInitialValidator", validationID)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) InitiateRemoveInitialValidator(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.InitiateRemoveInitialValidator(&_ERC20TokenSlotAuctionManager.TransactOpts, validationID)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactorSession) InitiateRemoveInitialValidator(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.InitiateRemoveInitialValidator(&_ERC20TokenSlotAuctionManager.TransactOpts, validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.contract.Transact(opts, "initiateValidatorRemoval", validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.InitiateValidatorRemoval(&_ERC20TokenSlotAuctionManager.TransactOpts, validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.InitiateValidatorRemoval(&_ERC20TokenSlotAuctionManager.TransactOpts, validationID)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xda4312a4.
//
// Solidity: function placeBid(uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactor) PlaceBid(opts *bind.TransactOpts, bid *big.Int, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.contract.Transact(opts, "placeBid", bid, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xda4312a4.
//
// Solidity: function placeBid(uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) PlaceBid(bid *big.Int, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.PlaceBid(&_ERC20TokenSlotAuctionManager.TransactOpts, bid, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xda4312a4.
//
// Solidity: function placeBid(uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns()
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerTransactorSession) PlaceBid(bid *big.Int, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _ERC20TokenSlotAuctionManager.Contract.PlaceBid(&_ERC20TokenSlotAuctionManager.TransactOpts, bid, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// ERC20TokenSlotAuctionManagerBidEvictedIterator is returned from FilterBidEvicted and is used to iterate over the raw logs and unpacked data for BidEvicted events raised by the ERC20TokenSlotAuctionManager contract.
type ERC20TokenSlotAuctionManagerBidEvictedIterator struct {
	Event *ERC20TokenSlotAuctionManagerBidEvicted // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSlotAuctionManagerBidEvictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSlotAuctionManagerBidEvicted)
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
		it.Event = new(ERC20TokenSlotAuctionManagerBidEvicted)
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
func (it *ERC20TokenSlotAuctionManagerBidEvictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSlotAuctionManagerBidEvictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSlotAuctionManagerBidEvicted represents a BidEvicted event raised by the ERC20TokenSlotAuctionManager contract.
type ERC20TokenSlotAuctionManagerBidEvicted struct {
	Bid    *big.Int
	NodeID common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidEvicted is a free log retrieval operation binding the contract event 0xdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc.
//
// Solidity: event BidEvicted(uint256 indexed bid, bytes indexed nodeID)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) FilterBidEvicted(opts *bind.FilterOpts, bid []*big.Int, nodeID [][]byte) (*ERC20TokenSlotAuctionManagerBidEvictedIterator, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenSlotAuctionManager.contract.FilterLogs(opts, "BidEvicted", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSlotAuctionManagerBidEvictedIterator{contract: _ERC20TokenSlotAuctionManager.contract, event: "BidEvicted", logs: logs, sub: sub}, nil
}

// WatchBidEvicted is a free log subscription operation binding the contract event 0xdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc.
//
// Solidity: event BidEvicted(uint256 indexed bid, bytes indexed nodeID)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) WatchBidEvicted(opts *bind.WatchOpts, sink chan<- *ERC20TokenSlotAuctionManagerBidEvicted, bid []*big.Int, nodeID [][]byte) (event.Subscription, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenSlotAuctionManager.contract.WatchLogs(opts, "BidEvicted", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSlotAuctionManagerBidEvicted)
				if err := _ERC20TokenSlotAuctionManager.contract.UnpackLog(event, "BidEvicted", log); err != nil {
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

// ParseBidEvicted is a log parse operation binding the contract event 0xdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc.
//
// Solidity: event BidEvicted(uint256 indexed bid, bytes indexed nodeID)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) ParseBidEvicted(log types.Log) (*ERC20TokenSlotAuctionManagerBidEvicted, error) {
	event := new(ERC20TokenSlotAuctionManagerBidEvicted)
	if err := _ERC20TokenSlotAuctionManager.contract.UnpackLog(event, "BidEvicted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenSlotAuctionManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20TokenSlotAuctionManager contract.
type ERC20TokenSlotAuctionManagerInitializedIterator struct {
	Event *ERC20TokenSlotAuctionManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSlotAuctionManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSlotAuctionManagerInitialized)
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
		it.Event = new(ERC20TokenSlotAuctionManagerInitialized)
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
func (it *ERC20TokenSlotAuctionManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSlotAuctionManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSlotAuctionManagerInitialized represents a Initialized event raised by the ERC20TokenSlotAuctionManager contract.
type ERC20TokenSlotAuctionManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20TokenSlotAuctionManagerInitializedIterator, error) {

	logs, sub, err := _ERC20TokenSlotAuctionManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSlotAuctionManagerInitializedIterator{contract: _ERC20TokenSlotAuctionManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenSlotAuctionManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenSlotAuctionManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSlotAuctionManagerInitialized)
				if err := _ERC20TokenSlotAuctionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) ParseInitialized(log types.Log) (*ERC20TokenSlotAuctionManagerInitialized, error) {
	event := new(ERC20TokenSlotAuctionManagerInitialized)
	if err := _ERC20TokenSlotAuctionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator is returned from FilterInitiatedAuctionValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedAuctionValidatorRegistration events raised by the ERC20TokenSlotAuctionManager contract.
type ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator struct {
	Event *ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistration)
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
		it.Event = new(ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistration)
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
func (it *ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistration represents a InitiatedAuctionValidatorRegistration event raised by the ERC20TokenSlotAuctionManager contract.
type ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistration struct {
	ValidationID     [32]byte
	OwnerAddress     common.Address
	ValidatorEndTime *big.Int
	Weight           uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInitiatedAuctionValidatorRegistration is a free log retrieval operation binding the contract event 0x032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 validatorEndTime, uint64 weight)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) FilterInitiatedAuctionValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, ownerAddress []common.Address) (*ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _ERC20TokenSlotAuctionManager.contract.FilterLogs(opts, "InitiatedAuctionValidatorRegistration", validationIDRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator{contract: _ERC20TokenSlotAuctionManager.contract, event: "InitiatedAuctionValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedAuctionValidatorRegistration is a free log subscription operation binding the contract event 0x032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 validatorEndTime, uint64 weight)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) WatchInitiatedAuctionValidatorRegistration(opts *bind.WatchOpts, sink chan<- *ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistration, validationID [][32]byte, ownerAddress []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _ERC20TokenSlotAuctionManager.contract.WatchLogs(opts, "InitiatedAuctionValidatorRegistration", validationIDRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistration)
				if err := _ERC20TokenSlotAuctionManager.contract.UnpackLog(event, "InitiatedAuctionValidatorRegistration", log); err != nil {
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

// ParseInitiatedAuctionValidatorRegistration is a log parse operation binding the contract event 0x032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 validatorEndTime, uint64 weight)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) ParseInitiatedAuctionValidatorRegistration(log types.Log) (*ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistration, error) {
	event := new(ERC20TokenSlotAuctionManagerInitiatedAuctionValidatorRegistration)
	if err := _ERC20TokenSlotAuctionManager.contract.UnpackLog(event, "InitiatedAuctionValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenSlotAuctionManagerNewValidatorAuctionIterator is returned from FilterNewValidatorAuction and is used to iterate over the raw logs and unpacked data for NewValidatorAuction events raised by the ERC20TokenSlotAuctionManager contract.
type ERC20TokenSlotAuctionManagerNewValidatorAuctionIterator struct {
	Event *ERC20TokenSlotAuctionManagerNewValidatorAuction // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSlotAuctionManagerNewValidatorAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSlotAuctionManagerNewValidatorAuction)
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
		it.Event = new(ERC20TokenSlotAuctionManagerNewValidatorAuction)
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
func (it *ERC20TokenSlotAuctionManagerNewValidatorAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSlotAuctionManagerNewValidatorAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSlotAuctionManagerNewValidatorAuction represents a NewValidatorAuction event raised by the ERC20TokenSlotAuctionManager contract.
type ERC20TokenSlotAuctionManagerNewValidatorAuction struct {
	ValidatorSlots       uint16
	ValidatorWeight      uint64
	MinValidatorDuration *big.Int
	AuctionEndTime       *big.Int
	MinimumBid           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewValidatorAuction is a free log retrieval operation binding the contract event 0x50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b54.
//
// Solidity: event NewValidatorAuction(uint16 validatorSlots, uint64 validatorWeight, uint256 minValidatorDuration, uint256 auctionEndTime, uint256 minimumBid)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) FilterNewValidatorAuction(opts *bind.FilterOpts) (*ERC20TokenSlotAuctionManagerNewValidatorAuctionIterator, error) {

	logs, sub, err := _ERC20TokenSlotAuctionManager.contract.FilterLogs(opts, "NewValidatorAuction")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSlotAuctionManagerNewValidatorAuctionIterator{contract: _ERC20TokenSlotAuctionManager.contract, event: "NewValidatorAuction", logs: logs, sub: sub}, nil
}

// WatchNewValidatorAuction is a free log subscription operation binding the contract event 0x50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b54.
//
// Solidity: event NewValidatorAuction(uint16 validatorSlots, uint64 validatorWeight, uint256 minValidatorDuration, uint256 auctionEndTime, uint256 minimumBid)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) WatchNewValidatorAuction(opts *bind.WatchOpts, sink chan<- *ERC20TokenSlotAuctionManagerNewValidatorAuction) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenSlotAuctionManager.contract.WatchLogs(opts, "NewValidatorAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSlotAuctionManagerNewValidatorAuction)
				if err := _ERC20TokenSlotAuctionManager.contract.UnpackLog(event, "NewValidatorAuction", log); err != nil {
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

// ParseNewValidatorAuction is a log parse operation binding the contract event 0x50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b54.
//
// Solidity: event NewValidatorAuction(uint16 validatorSlots, uint64 validatorWeight, uint256 minValidatorDuration, uint256 auctionEndTime, uint256 minimumBid)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) ParseNewValidatorAuction(log types.Log) (*ERC20TokenSlotAuctionManagerNewValidatorAuction, error) {
	event := new(ERC20TokenSlotAuctionManagerNewValidatorAuction)
	if err := _ERC20TokenSlotAuctionManager.contract.UnpackLog(event, "NewValidatorAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenSlotAuctionManagerSuccessfulBidPlacedIterator is returned from FilterSuccessfulBidPlaced and is used to iterate over the raw logs and unpacked data for SuccessfulBidPlaced events raised by the ERC20TokenSlotAuctionManager contract.
type ERC20TokenSlotAuctionManagerSuccessfulBidPlacedIterator struct {
	Event *ERC20TokenSlotAuctionManagerSuccessfulBidPlaced // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSlotAuctionManagerSuccessfulBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSlotAuctionManagerSuccessfulBidPlaced)
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
		it.Event = new(ERC20TokenSlotAuctionManagerSuccessfulBidPlaced)
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
func (it *ERC20TokenSlotAuctionManagerSuccessfulBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSlotAuctionManagerSuccessfulBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSlotAuctionManagerSuccessfulBidPlaced represents a SuccessfulBidPlaced event raised by the ERC20TokenSlotAuctionManager contract.
type ERC20TokenSlotAuctionManagerSuccessfulBidPlaced struct {
	Bid    *big.Int
	NodeID common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSuccessfulBidPlaced is a free log retrieval operation binding the contract event 0x864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029.
//
// Solidity: event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) FilterSuccessfulBidPlaced(opts *bind.FilterOpts, bid []*big.Int, nodeID [][]byte) (*ERC20TokenSlotAuctionManagerSuccessfulBidPlacedIterator, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenSlotAuctionManager.contract.FilterLogs(opts, "SuccessfulBidPlaced", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSlotAuctionManagerSuccessfulBidPlacedIterator{contract: _ERC20TokenSlotAuctionManager.contract, event: "SuccessfulBidPlaced", logs: logs, sub: sub}, nil
}

// WatchSuccessfulBidPlaced is a free log subscription operation binding the contract event 0x864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029.
//
// Solidity: event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) WatchSuccessfulBidPlaced(opts *bind.WatchOpts, sink chan<- *ERC20TokenSlotAuctionManagerSuccessfulBidPlaced, bid []*big.Int, nodeID [][]byte) (event.Subscription, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenSlotAuctionManager.contract.WatchLogs(opts, "SuccessfulBidPlaced", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSlotAuctionManagerSuccessfulBidPlaced)
				if err := _ERC20TokenSlotAuctionManager.contract.UnpackLog(event, "SuccessfulBidPlaced", log); err != nil {
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

// ParseSuccessfulBidPlaced is a log parse operation binding the contract event 0x864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029.
//
// Solidity: event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerFilterer) ParseSuccessfulBidPlaced(log types.Log) (*ERC20TokenSlotAuctionManagerSuccessfulBidPlaced, error) {
	event := new(ERC20TokenSlotAuctionManagerSuccessfulBidPlaced)
	if err := _ERC20TokenSlotAuctionManager.contract.UnpackLog(event, "SuccessfulBidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
