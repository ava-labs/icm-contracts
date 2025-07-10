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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"}],\"name\":\"AuctionEndTimeNotReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionFinalizing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionNotInProgress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minumumBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBid\",\"type\":\"uint256\"}],\"name\":\"BidSmallerThanMinimum\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"DuplicateBidInContention\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"smallestAcceptableBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBid\",\"type\":\"uint256\"}],\"name\":\"InsufficientBidToWinAuction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeIsValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeIsWinning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationTimeLimit\",\"type\":\"uint256\"}],\"name\":\"ValidatorTimeLimitNotPassed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"BidEvicted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorEndTime\",\"type\":\"uint256\"}],\"name\":\"InitiatedAuctionValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"validatorSlots\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minValidatorDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"}],\"name\":\"NewValidatorAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"SuccessfulBidPlaced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MinBidRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinValidatorDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinimumBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionFinalizing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"bidderInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"validatorslots\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minAuctionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValidatorDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"}],\"name\":\"initiateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSlots\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"validatorsByNodeID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"validatorsByValidationID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506040516124c03803806124c0833981016040819052602b91607b565b5f80546001600160a01b039384166001600160a01b0319909116179055600180546001600160b01b0319169190921617905560a7565b80516001600160a01b03811681146076575f80fd5b919050565b5f8060408385031215608b575f80fd5b6092836061565b9150609e602084016061565b90509250929050565b61240c806100b45f395ff3fe608060405234801561000f575f80fd5b5060043610610131575f3560e01c80639681d940116100b4578063b6e6a2ca11610079578063b6e6a2ca146102af578063cb3f66fb146102c2578063da4312a4146102e6578063df393075146102f9578063f556f60a1461030d578063fe67a54b14610316575f80fd5b80639681d940146102505780639c91ddbb14610263578063a3a65e4814610276578063a476f67514610289578063ae9483e01461029c575f80fd5b80634b449cba116100fa5780634b449cba146101f05780636fe4d1e11461020757806372d49319146102105780637a7df93214610235578063913ed0b114610248575f80fd5b8062d841f81461013557806305af82561461016b5780630d5daf3b146101805780631077830a146101a557806336339388146101c6575b5f80fd5b60035461014e906201000090046001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b61017e610179366004611c64565b61031e565b005b61019361018e366004611c7d565b610393565b60405161016296959493929190611d47565b6003546101b39061ffff1681565b60405161ffff9091168152602001610162565b5f546101d8906001600160a01b031681565b6040516001600160a01b039091168152602001610162565b6101f960025481565b604051908152602001610162565b6101f960045481565b61022361021e366004611c7d565b6105c8565b60405161016296959493929190611db4565b61017e610243366004611e0f565b61071c565b6101f961084a565b6101f961025e366004611c64565b6108ac565b610223610271366004611f44565b610924565b6101f9610284366004611c64565b610961565b61017e610297366004611c7d565b610997565b6001546101d8906001600160a01b031681565b61017e6102bd366004611c7d565b6109f4565b6001546102d690600160a81b900460ff1681565b6040519015158152602001610162565b61017e6102f4366004612049565b610b48565b6001546102d690600160a01b900460ff1681565b6101f960055481565b61017e611157565b60015460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af115801561036b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061038f91906120f7565b5050565b60096020525f90815260409020805460018201546002830180546001600160a01b039093169391926103c49061210e565b80601f01602080910402602001604051908101604052809291908181526020018280546103f09061210e565b801561043b5780601f106104125761010080835404028352916020019161043b565b820191905f5260205f20905b81548152906001019060200180831161041e57829003601f168201915b5050505050908060030180546104509061210e565b80601f016020809104026020016040519081016040528092919081815260200182805461047c9061210e565b80156104c75780601f1061049e576101008083540402835291602001916104c7565b820191905f5260205f20905b8154815290600101906020018083116104aa57829003601f168201915b505060408051808201825260048701805463ffffffff16825260058801805484516020828102820181019096528181529899989397509195508387019450919283018282801561053e57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610520575b50505091909252505060408051808201825260068501805463ffffffff168252600786018054845160208281028201810190965281815296979693955091938086019392908301828280156105ba57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161059c575b505050505081525050905086565b60086020525f90815260409020805460018201546002830180546001600160a01b039093169391926105f99061210e565b80601f01602080910402602001604051908101604052809291908181526020018280546106259061210e565b80156106705780601f1061064757610100808354040283529160200191610670565b820191905f5260205f20905b81548152906001019060200180831161065357829003601f168201915b5050505050908060030180546106859061210e565b80601f01602080910402602001604051908101604052809291908181526020018280546106b19061210e565b80156106fc5780601f106106d3576101008083540402835291602001916106fc565b820191905f5260205f20905b8154815290600101906020018083116106df57829003601f168201915b5050505060048301546005909301549192916001600160401b0316905086565b600154600160a01b900460ff161561074757604051639bbb1f3b60e01b815260040160405180910390fd5b600154600160a81b900460ff161561077257604051631f86de2560e21b815260040160405180910390fd5b6001805460ff60a01b1916600160a01b179055600380546001600160401b038616620100000269ffffffffffffffff0000199091161790556107b4834261215a565b6002556003805461ffff191661ffff8716179055600482905560058190555f600a8190556107e490600b90611b80565b6003546005546040805161ffff90931683526001600160401b038716602084015282018490526060820185905260808201527f50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b549060a00160405180910390a15050505050565b6001545f90600160a01b900460ff16610875576040516232f00d60e71b815260040160405180910390fd5b60035461ffff16610885600b5490565b1015610892575060055490565b61089c600b611832565b6108a790600161215a565b905090565b60015460405163025a076560e61b815263ffffffff831660048201525f916001600160a01b031690639681d940906024015b6020604051808303815f875af11580156108fa573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061091e91906120f7565b92915050565b80516020818301810180516007825292820191909301209152805460018201546002830180546001600160a01b039093169391926105f99061210e565b600154604051631474cbc960e31b815263ffffffff831660048201525f916001600160a01b03169063a3a65e48906024016108de565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015b5f604051808303815f87803b1580156109db575f80fd5b505af11580156109ed573d5f803e3d5ffd5b5050505050565b5f81815260086020526040902060010154421015610a40575f8181526008602052604090819020600101549051631a936b7b60e11b815260048101919091526024015b60405180910390fd5b600760085f8381526020019081526020015f20600201604051610a63919061216d565b90815260405190819003602001902080546001600160a01b03191681555f60018201819055610a956002830182611b9e565b610aa2600383015f611b9e565b505f600482018190556005909101805467ffffffffffffffff1916905581815260086020526040812080546001600160a01b03191681556001810182905590610aee6002830182611b9e565b610afb600383015f611b9e565b505f6004828101919091556005909101805467ffffffffffffffff19169055600154604051635b73516560e11b81529182018390526001600160a01b03169063b6e6a2ca906024016109c4565b610b50611857565b600154600160a01b900460ff16610b79576040516232f00d60e71b815260040160405180910390fd5b60015460405163d47a948b60e01b81526001600160a01b039091169063d47a948b90610ba99087906004016121df565b602060405180830381865afa158015610bc4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610be891906120f7565b15610c08578360405163f3c815f760e01b8152600401610a3791906121df565b600684604051610c1891906121f1565b9081526040519081900360200190205460ff1615610c4b578360405163f3c1d91160e01b8152600401610a3791906121df565b5f858152600960205260409020546001600160a01b031615610c8357604051630517e2e760e21b815260048101869052602401610a37565b846005541115610cb4576005546040516301a0013b60e61b8152600481019190915260248101869052604401610a37565b60035461ffff16610cc4600b5490565b1015610d52575f546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015610d1d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d41919061220c565b50610d4d600b866118a1565b610fb8565b84610d5d600b611832565b1015610f7f575f546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015610db6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dda919061220c565b505f610de7600b876118ae565b600a8190555f81815260096020526040908190209051919250610e0f9160029091019061216d565b6040519081900381209082907fdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc905f90a35f8054828252600960205260409182902054915163a9059cbb60e01b81526001600160a01b0392831660048201526024810184905291169063a9059cbb906044016020604051808303815f875af1158015610e9d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ec1919061220c565b50600660095f8381526020019081526020015f20600201604051610ee5919061216d565b9081526040805160209281900383019020805460ff191690555f8381526009909252812080546001600160a01b03191681556001810182905590610f2c6002830182611b9e565b610f39600383015f611b9e565b60048201805463ffffffff191681555f610f566005850182611b80565b505060068201805463ffffffff191681555f610f756007850182611b80565b5050505050610fb8565b610f89600b611832565b610f9490600161215a565b604051631a79674760e11b8152600481019190915260248101869052604401610a37565b6040805160c0810182523381526020808201888152828401888152606084018890526080840187905260a084018690525f8a8152600990935293909120825181546001600160a01b0319166001600160a01b03909116178155905160018201559151909190600282019061102c9082612274565b50606082015160038201906110419082612274565b506080820151805160048301805463ffffffff191663ffffffff909216919091178155602080830151805161107c9260058701920190611bd5565b50505060a0820151805160068301805463ffffffff191663ffffffff90921691909117815560208083015180516110b99260078701920190611bd5565b50505090505060016006856040516110d191906121f1565b908152604051908190036020018120805492151560ff19909316929092179091556110fd9085906121f1565b6040519081900381209086907f864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029905f90a36109ed60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b61115f611857565b600154600160a01b900460ff16611188576040516232f00d60e71b815260040160405180910390fd5b6001805461ffff60a01b1916600160a81b17905560025442116111c457600254604051638230260760e01b8152600401610a3791815260200190565b5f600255600a541580156111d95750600b5415155b156111ec576111e8600b611832565b600a555b5f6111f6600b5490565b11156117f2575f611207600b6118c3565b5f818152600960209081526040808320815160c08101835281546001600160a01b03168152600182015493810193909352600281018054959650939492939092918401916112549061210e565b80601f01602080910402602001604051908101604052809291908181526020018280546112809061210e565b80156112cb5780601f106112a2576101008083540402835291602001916112cb565b820191905f5260205f20905b8154815290600101906020018083116112ae57829003601f168201915b505050505081526020016003820180546112e49061210e565b80601f01602080910402602001604051908101604052809291908181526020018280546113109061210e565b801561135b5780601f106113325761010080835404028352916020019161135b565b820191905f5260205f20905b81548152906001019060200180831161133e57829003601f168201915b505050918352505060408051808201825260048401805463ffffffff1682526005850180548451602082810282018101909652818152958501959394929385840193909291908301828280156113d857602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116113ba575b50505091909252505050815260408051808201825260068401805463ffffffff16825260078501805484516020828102820181019096528181529585019593949293858401939092919083018282801561145957602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161143b575b505050919092525050509052505f548151600a549293506001600160a01b039091169163a9059cbb919061148d9086612333565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156114d5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114f9919061220c565b506001546040828101516060840151608085015160a08601516003549451634e5bb12760e11b81525f966001600160a01b031695639cb7624e956115579590949093909290916201000090046001600160401b031690600401612346565b6020604051808303815f875af1158015611573573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061159791906120f7565b9050815f01516001600160a01b0316817fccb9a8514540a67c432230e6261f90893e22e8e31f65c40e7629b5a47374dbbe6002546004546115d8919061215a565b60405190815260200160405180910390a36040518060c00160405280835f01516001600160a01b03168152602001600254600454611616919061215a565b81526040848101805160208401526060808701518385015283018590526003546201000090046001600160401b03166080909301929092529051905160079161165e916121f1565b9081526040805160209281900383019020835181546001600160a01b0319166001600160a01b0390911617815591830151600183015582015160028201906116a69082612274565b50606082015160038201906116bb9082612274565b506080820151816004015560a0820151816005015f6101000a8154816001600160401b0302191690836001600160401b031602179055509050506040518060c00160405280835f01516001600160a01b03168152602001600254600454611722919061215a565b81526040848101516020808401919091526060808701518385015283018590526003546201000090046001600160401b03166080909301929092525f84815260088352819020835181546001600160a01b0319166001600160a01b03909116178155918301516001830155820151600282019061179f9082612274565b50606082015160038201906117b49082612274565b506080820151600482015560a0909101516005909101805467ffffffffffffffff19166001600160401b039092169190911790555050600a556111ec565b6003805461ffff191690555f6004556001805460ff60a81b191681557f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f815f015f81548110611847576118476123ae565b905f5260205f2001549050919050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080546001190161189b57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b61038f82826118d06118d4565b5f6118bc83836118d0611905565b9392505050565b5f61091e826118d0611952565b1190565b5f6118dd845490565b84546001810186555f868152602090200184905590506118ff848285856119e3565b50505050565b5f8061190f855490565b9050805f03611922576119226031611a2b565b5f61192d8682611a3c565b5490508461193b875f611a3c565b5561194986835f8888611a63565b95945050505050565b5f8061195c845490565b9050805f0361196f5761196f6031611a2b565b5f61197a8582611a3c565b5490505f61198b865f198501611a3c565b5486549091508690806119a0576119a06123c2565b600190038181905f5260205f20015f90559055806119c95f885f01611a3c90919063ffffffff16565b556119da865f1985015f8489611a63565b50949350505050565b82156118ff5760025f198401045f6119fb8683611a3c565b549050611a0c818563ffffffff8616565b15611a185750506118ff565b611a23868684611b58565b5092506119e3565b634e487b715f52806020526024601cfd5b5f82611a5b611a5884611a54845f9081526020902090565b0190565b90565b949350505050565b6001600160ff1b038310156109ed57600283810260018101910185811015611b0c575f611a908884611a3c565b5490505f611a9e8984611a3c565b549050611aaf828763ffffffff8816565b80611ac35750611ac381878763ffffffff16565b15611b05575f611ae98585611adc86868b63ffffffff16565b1515918118919091021890565b9050611af68a8983611b58565b611b038a8a838a8a611a63565b505b5050611b4f565b85821015611b4f575f611b1f8884611a3c565b549050611b30818663ffffffff8716565b15611b4d57611b40888785611b58565b611b4d8888858888611a63565b505b50505050505050565b5f611b638484611a3c565b90505f611b708584611a3c565b8054835490915590915550505050565b5080545f8255905f5260205f2090810190611b9b9190611c38565b50565b508054611baa9061210e565b5f825580601f10611bb9575050565b601f0160209004905f5260205f2090810190611b9b9190611c38565b828054828255905f5260205f20908101928215611c28579160200282015b82811115611c2857825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611bf3565b50611c34929150611c38565b5090565b5b80821115611c34575f8155600101611c39565b803563ffffffff81168114611c5f575f80fd5b919050565b5f60208284031215611c74575f80fd5b6118bc82611c4c565b5f60208284031215611c8d575f80fd5b5035919050565b5f5b83811015611cae578181015183820152602001611c96565b50505f910152565b5f8151808452611ccd816020860160208601611c94565b601f01601f19169290920160200192915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611d3c5784516001600160a01b03168252938301936001929092019190830190611d13565b509695505050505050565b60018060a01b038716815285602082015260c060408201525f611d6d60c0830187611cb6565b8281036060840152611d7f8187611cb6565b90508281036080840152611d938186611ce1565b905082810360a0840152611da78185611ce1565b9998505050505050505050565b60018060a01b038716815285602082015260c060408201525f611dda60c0830187611cb6565b8281036060840152611dec8187611cb6565b9150508360808301526001600160401b03831660a0830152979650505050505050565b5f805f805f60a08688031215611e23575f80fd5b853561ffff81168114611e34575f80fd5b945060208601356001600160401b0381168114611e4f575f80fd5b94979496505050506040830135926060810135926080909101359150565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611ea357611ea3611e6d565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611ed157611ed1611e6d565b604052919050565b5f82601f830112611ee8575f80fd5b81356001600160401b03811115611f0157611f01611e6d565b611f14601f8201601f1916602001611ea9565b818152846020838601011115611f28575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215611f54575f80fd5b81356001600160401b03811115611f69575f80fd5b611a5b84828501611ed9565b5f60408284031215611f85575f80fd5b611f8d611e81565b9050611f9882611c4c565b81526020808301356001600160401b0380821115611fb4575f80fd5b818501915085601f830112611fc7575f80fd5b813581811115611fd957611fd9611e6d565b8060051b9150611fea848301611ea9565b8181529183018401918481019088841115612003575f80fd5b938501935b8385101561203857843592506001600160a01b0383168314612028575f80fd5b8282529385019390850190612008565b808688015250505050505092915050565b5f805f805f60a0868803121561205d575f80fd5b8535945060208601356001600160401b038082111561207a575f80fd5b61208689838a01611ed9565b9550604088013591508082111561209b575f80fd5b6120a789838a01611ed9565b945060608801359150808211156120bc575f80fd5b6120c889838a01611f75565b935060808801359150808211156120dd575f80fd5b506120ea88828901611f75565b9150509295509295909350565b5f60208284031215612107575f80fd5b5051919050565b600181811c9082168061212257607f821691505b60208210810361214057634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561091e5761091e612146565b5f80835461217a8161210e565b6001828116801561219257600181146121a7576121d3565b60ff19841687528215158302870194506121d3565b875f526020805f205f5b858110156121ca5781548a8201529084019082016121b1565b50505082870194505b50929695505050505050565b602081525f6118bc6020830184611cb6565b5f8251612202818460208701611c94565b9190910192915050565b5f6020828403121561221c575f80fd5b815180151581146118bc575f80fd5b601f82111561226f57805f5260205f20601f840160051c810160208510156122505750805b601f840160051c820191505b818110156109ed575f815560010161225c565b505050565b81516001600160401b0381111561228d5761228d611e6d565b6122a18161229b845461210e565b8461222b565b602080601f8311600181146122d4575f84156122bd5750858301515b5f19600386901b1c1916600185901b17855561232b565b5f85815260208120601f198616915b82811015612302578886015182559484019460019091019084016122e3565b508582101561231f57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b8181038181111561091e5761091e612146565b60a081525f61235860a0830188611cb6565b828103602084015261236a8188611cb6565b9050828103604084015261237e8187611ce1565b905082810360608401526123928186611ce1565b9150506001600160401b03831660808301529695505050505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52603160045260245ffdfea26469706673582212208854e27f48eb8aa4bd5f3062fafe19ee12f91c673a2e0a2cf7e9a0e8894416ea64736f6c63430008190033",
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

// MinBidRequired is a free data retrieval call binding the contract method 0x913ed0b1.
//
// Solidity: function MinBidRequired() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCaller) MinBidRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "MinBidRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBidRequired is a free data retrieval call binding the contract method 0x913ed0b1.
//
// Solidity: function MinBidRequired() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerSession) MinBidRequired() (*big.Int, error) {
	return _SlotAuctionManager.Contract.MinBidRequired(&_SlotAuctionManager.CallOpts)
}

// MinBidRequired is a free data retrieval call binding the contract method 0x913ed0b1.
//
// Solidity: function MinBidRequired() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) MinBidRequired() (*big.Int, error) {
	return _SlotAuctionManager.Contract.MinBidRequired(&_SlotAuctionManager.CallOpts)
}

// MinValidatorDuration is a free data retrieval call binding the contract method 0x6fe4d1e1.
//
// Solidity: function MinValidatorDuration() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCaller) MinValidatorDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "MinValidatorDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValidatorDuration is a free data retrieval call binding the contract method 0x6fe4d1e1.
//
// Solidity: function MinValidatorDuration() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerSession) MinValidatorDuration() (*big.Int, error) {
	return _SlotAuctionManager.Contract.MinValidatorDuration(&_SlotAuctionManager.CallOpts)
}

// MinValidatorDuration is a free data retrieval call binding the contract method 0x6fe4d1e1.
//
// Solidity: function MinValidatorDuration() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) MinValidatorDuration() (*big.Int, error) {
	return _SlotAuctionManager.Contract.MinValidatorDuration(&_SlotAuctionManager.CallOpts)
}

// MinimumBid is a free data retrieval call binding the contract method 0xf556f60a.
//
// Solidity: function MinimumBid() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCaller) MinimumBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "MinimumBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumBid is a free data retrieval call binding the contract method 0xf556f60a.
//
// Solidity: function MinimumBid() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerSession) MinimumBid() (*big.Int, error) {
	return _SlotAuctionManager.Contract.MinimumBid(&_SlotAuctionManager.CallOpts)
}

// MinimumBid is a free data retrieval call binding the contract method 0xf556f60a.
//
// Solidity: function MinimumBid() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) MinimumBid() (*big.Int, error) {
	return _SlotAuctionManager.Contract.MinimumBid(&_SlotAuctionManager.CallOpts)
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

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCaller) AuctionEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "auctionEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerSession) AuctionEndTime() (*big.Int, error) {
	return _SlotAuctionManager.Contract.AuctionEndTime(&_SlotAuctionManager.CallOpts)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) AuctionEndTime() (*big.Int, error) {
	return _SlotAuctionManager.Contract.AuctionEndTime(&_SlotAuctionManager.CallOpts)
}

// AuctionFinalizing is a free data retrieval call binding the contract method 0xcb3f66fb.
//
// Solidity: function auctionFinalizing() view returns(bool)
func (_SlotAuctionManager *SlotAuctionManagerCaller) AuctionFinalizing(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "auctionFinalizing")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuctionFinalizing is a free data retrieval call binding the contract method 0xcb3f66fb.
//
// Solidity: function auctionFinalizing() view returns(bool)
func (_SlotAuctionManager *SlotAuctionManagerSession) AuctionFinalizing() (bool, error) {
	return _SlotAuctionManager.Contract.AuctionFinalizing(&_SlotAuctionManager.CallOpts)
}

// AuctionFinalizing is a free data retrieval call binding the contract method 0xcb3f66fb.
//
// Solidity: function auctionFinalizing() view returns(bool)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) AuctionFinalizing() (bool, error) {
	return _SlotAuctionManager.Contract.AuctionFinalizing(&_SlotAuctionManager.CallOpts)
}

// AuctionInProgress is a free data retrieval call binding the contract method 0xdf393075.
//
// Solidity: function auctionInProgress() view returns(bool)
func (_SlotAuctionManager *SlotAuctionManagerCaller) AuctionInProgress(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "auctionInProgress")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuctionInProgress is a free data retrieval call binding the contract method 0xdf393075.
//
// Solidity: function auctionInProgress() view returns(bool)
func (_SlotAuctionManager *SlotAuctionManagerSession) AuctionInProgress() (bool, error) {
	return _SlotAuctionManager.Contract.AuctionInProgress(&_SlotAuctionManager.CallOpts)
}

// AuctionInProgress is a free data retrieval call binding the contract method 0xdf393075.
//
// Solidity: function auctionInProgress() view returns(bool)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) AuctionInProgress() (bool, error) {
	return _SlotAuctionManager.Contract.AuctionInProgress(&_SlotAuctionManager.CallOpts)
}

// BidderInfo is a free data retrieval call binding the contract method 0x0d5daf3b.
//
// Solidity: function bidderInfo(uint256 bid) view returns(address addr, uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner)
func (_SlotAuctionManager *SlotAuctionManagerCaller) BidderInfo(opts *bind.CallOpts, bid *big.Int) (struct {
	Addr                  common.Address
	Bid                   *big.Int
	NodeID                []byte
	BlsPublicKey          []byte
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "bidderInfo", bid)

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
func (_SlotAuctionManager *SlotAuctionManagerSession) BidderInfo(bid *big.Int) (struct {
	Addr                  common.Address
	Bid                   *big.Int
	NodeID                []byte
	BlsPublicKey          []byte
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}, error) {
	return _SlotAuctionManager.Contract.BidderInfo(&_SlotAuctionManager.CallOpts, bid)
}

// BidderInfo is a free data retrieval call binding the contract method 0x0d5daf3b.
//
// Solidity: function bidderInfo(uint256 bid) view returns(address addr, uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) BidderInfo(bid *big.Int) (struct {
	Addr                  common.Address
	Bid                   *big.Int
	NodeID                []byte
	BlsPublicKey          []byte
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}, error) {
	return _SlotAuctionManager.Contract.BidderInfo(&_SlotAuctionManager.CallOpts, bid)
}

// ValidatorSlots is a free data retrieval call binding the contract method 0x1077830a.
//
// Solidity: function validatorSlots() view returns(uint16)
func (_SlotAuctionManager *SlotAuctionManagerCaller) ValidatorSlots(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "validatorSlots")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ValidatorSlots is a free data retrieval call binding the contract method 0x1077830a.
//
// Solidity: function validatorSlots() view returns(uint16)
func (_SlotAuctionManager *SlotAuctionManagerSession) ValidatorSlots() (uint16, error) {
	return _SlotAuctionManager.Contract.ValidatorSlots(&_SlotAuctionManager.CallOpts)
}

// ValidatorSlots is a free data retrieval call binding the contract method 0x1077830a.
//
// Solidity: function validatorSlots() view returns(uint16)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) ValidatorSlots() (uint16, error) {
	return _SlotAuctionManager.Contract.ValidatorSlots(&_SlotAuctionManager.CallOpts)
}

// ValidatorWeight is a free data retrieval call binding the contract method 0x00d841f8.
//
// Solidity: function validatorWeight() view returns(uint64)
func (_SlotAuctionManager *SlotAuctionManagerCaller) ValidatorWeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "validatorWeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValidatorWeight is a free data retrieval call binding the contract method 0x00d841f8.
//
// Solidity: function validatorWeight() view returns(uint64)
func (_SlotAuctionManager *SlotAuctionManagerSession) ValidatorWeight() (uint64, error) {
	return _SlotAuctionManager.Contract.ValidatorWeight(&_SlotAuctionManager.CallOpts)
}

// ValidatorWeight is a free data retrieval call binding the contract method 0x00d841f8.
//
// Solidity: function validatorWeight() view returns(uint64)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) ValidatorWeight() (uint64, error) {
	return _SlotAuctionManager.Contract.ValidatorWeight(&_SlotAuctionManager.CallOpts)
}

// ValidatorsByNodeID is a free data retrieval call binding the contract method 0x9c91ddbb.
//
// Solidity: function validatorsByNodeID(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 validatorWeight)
func (_SlotAuctionManager *SlotAuctionManagerCaller) ValidatorsByNodeID(opts *bind.CallOpts, nodeID []byte) (struct {
	Addr            common.Address
	EndTime         *big.Int
	NodeID          []byte
	BlsPublicKey    []byte
	ValidationID    [32]byte
	ValidatorWeight uint64
}, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "validatorsByNodeID", nodeID)

	outstruct := new(struct {
		Addr            common.Address
		EndTime         *big.Int
		NodeID          []byte
		BlsPublicKey    []byte
		ValidationID    [32]byte
		ValidatorWeight uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NodeID = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.BlsPublicKey = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.ValidationID = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.ValidatorWeight = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// ValidatorsByNodeID is a free data retrieval call binding the contract method 0x9c91ddbb.
//
// Solidity: function validatorsByNodeID(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 validatorWeight)
func (_SlotAuctionManager *SlotAuctionManagerSession) ValidatorsByNodeID(nodeID []byte) (struct {
	Addr            common.Address
	EndTime         *big.Int
	NodeID          []byte
	BlsPublicKey    []byte
	ValidationID    [32]byte
	ValidatorWeight uint64
}, error) {
	return _SlotAuctionManager.Contract.ValidatorsByNodeID(&_SlotAuctionManager.CallOpts, nodeID)
}

// ValidatorsByNodeID is a free data retrieval call binding the contract method 0x9c91ddbb.
//
// Solidity: function validatorsByNodeID(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 validatorWeight)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) ValidatorsByNodeID(nodeID []byte) (struct {
	Addr            common.Address
	EndTime         *big.Int
	NodeID          []byte
	BlsPublicKey    []byte
	ValidationID    [32]byte
	ValidatorWeight uint64
}, error) {
	return _SlotAuctionManager.Contract.ValidatorsByNodeID(&_SlotAuctionManager.CallOpts, nodeID)
}

// ValidatorsByValidationID is a free data retrieval call binding the contract method 0x72d49319.
//
// Solidity: function validatorsByValidationID(bytes32 validationID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 validatorWeight)
func (_SlotAuctionManager *SlotAuctionManagerCaller) ValidatorsByValidationID(opts *bind.CallOpts, validationID [32]byte) (struct {
	Addr            common.Address
	EndTime         *big.Int
	NodeID          []byte
	BlsPublicKey    []byte
	ValidationID    [32]byte
	ValidatorWeight uint64
}, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "validatorsByValidationID", validationID)

	outstruct := new(struct {
		Addr            common.Address
		EndTime         *big.Int
		NodeID          []byte
		BlsPublicKey    []byte
		ValidationID    [32]byte
		ValidatorWeight uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NodeID = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.BlsPublicKey = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.ValidationID = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.ValidatorWeight = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// ValidatorsByValidationID is a free data retrieval call binding the contract method 0x72d49319.
//
// Solidity: function validatorsByValidationID(bytes32 validationID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 validatorWeight)
func (_SlotAuctionManager *SlotAuctionManagerSession) ValidatorsByValidationID(validationID [32]byte) (struct {
	Addr            common.Address
	EndTime         *big.Int
	NodeID          []byte
	BlsPublicKey    []byte
	ValidationID    [32]byte
	ValidatorWeight uint64
}, error) {
	return _SlotAuctionManager.Contract.ValidatorsByValidationID(&_SlotAuctionManager.CallOpts, validationID)
}

// ValidatorsByValidationID is a free data retrieval call binding the contract method 0x72d49319.
//
// Solidity: function validatorsByValidationID(bytes32 validationID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 validatorWeight)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) ValidatorsByValidationID(validationID [32]byte) (struct {
	Addr            common.Address
	EndTime         *big.Int
	NodeID          []byte
	BlsPublicKey    []byte
	ValidationID    [32]byte
	ValidatorWeight uint64
}, error) {
	return _SlotAuctionManager.Contract.ValidatorsByValidationID(&_SlotAuctionManager.CallOpts, validationID)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) CompleteRemoveInitialValidator(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "completeRemoveInitialValidator", messageIndex)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) CompleteRemoveInitialValidator(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteRemoveInitialValidator(&_SlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) CompleteRemoveInitialValidator(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteRemoveInitialValidator(&_SlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteValidatorRegistration(&_SlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteValidatorRegistration(&_SlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteValidatorRemoval(&_SlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.CompleteValidatorRemoval(&_SlotAuctionManager.TransactOpts, messageIndex)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) EndAuction() (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.EndAuction(&_SlotAuctionManager.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) EndAuction() (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.EndAuction(&_SlotAuctionManager.TransactOpts)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x7a7df932.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 minAuctionDuration, uint256 minValidatorDuration, uint256 minimumBid) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) InitiateAuction(opts *bind.TransactOpts, validatorslots uint16, weight uint64, minAuctionDuration *big.Int, minValidatorDuration *big.Int, minimumBid *big.Int) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "initiateAuction", validatorslots, weight, minAuctionDuration, minValidatorDuration, minimumBid)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x7a7df932.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 minAuctionDuration, uint256 minValidatorDuration, uint256 minimumBid) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) InitiateAuction(validatorslots uint16, weight uint64, minAuctionDuration *big.Int, minValidatorDuration *big.Int, minimumBid *big.Int) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateAuction(&_SlotAuctionManager.TransactOpts, validatorslots, weight, minAuctionDuration, minValidatorDuration, minimumBid)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x7a7df932.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 minAuctionDuration, uint256 minValidatorDuration, uint256 minimumBid) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) InitiateAuction(validatorslots uint16, weight uint64, minAuctionDuration *big.Int, minValidatorDuration *big.Int, minimumBid *big.Int) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateAuction(&_SlotAuctionManager.TransactOpts, validatorslots, weight, minAuctionDuration, minValidatorDuration, minimumBid)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) InitiateRemoveInitialValidator(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "initiateRemoveInitialValidator", validationID)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) InitiateRemoveInitialValidator(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateRemoveInitialValidator(&_SlotAuctionManager.TransactOpts, validationID)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) InitiateRemoveInitialValidator(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateRemoveInitialValidator(&_SlotAuctionManager.TransactOpts, validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "initiateValidatorRemoval", validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateValidatorRemoval(&_SlotAuctionManager.TransactOpts, validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateValidatorRemoval(&_SlotAuctionManager.TransactOpts, validationID)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xda4312a4.
//
// Solidity: function placeBid(uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) PlaceBid(opts *bind.TransactOpts, bid *big.Int, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "placeBid", bid, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xda4312a4.
//
// Solidity: function placeBid(uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) PlaceBid(bid *big.Int, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.PlaceBid(&_SlotAuctionManager.TransactOpts, bid, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xda4312a4.
//
// Solidity: function placeBid(uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) PlaceBid(bid *big.Int, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.PlaceBid(&_SlotAuctionManager.TransactOpts, bid, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// SlotAuctionManagerBidEvictedIterator is returned from FilterBidEvicted and is used to iterate over the raw logs and unpacked data for BidEvicted events raised by the SlotAuctionManager contract.
type SlotAuctionManagerBidEvictedIterator struct {
	Event *SlotAuctionManagerBidEvicted // Event containing the contract specifics and raw log

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
func (it *SlotAuctionManagerBidEvictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlotAuctionManagerBidEvicted)
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
		it.Event = new(SlotAuctionManagerBidEvicted)
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
func (it *SlotAuctionManagerBidEvictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlotAuctionManagerBidEvictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlotAuctionManagerBidEvicted represents a BidEvicted event raised by the SlotAuctionManager contract.
type SlotAuctionManagerBidEvicted struct {
	Bid    *big.Int
	NodeID common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidEvicted is a free log retrieval operation binding the contract event 0xdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc.
//
// Solidity: event BidEvicted(uint256 indexed bid, bytes indexed nodeID)
func (_SlotAuctionManager *SlotAuctionManagerFilterer) FilterBidEvicted(opts *bind.FilterOpts, bid []*big.Int, nodeID [][]byte) (*SlotAuctionManagerBidEvictedIterator, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _SlotAuctionManager.contract.FilterLogs(opts, "BidEvicted", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &SlotAuctionManagerBidEvictedIterator{contract: _SlotAuctionManager.contract, event: "BidEvicted", logs: logs, sub: sub}, nil
}

// WatchBidEvicted is a free log subscription operation binding the contract event 0xdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc.
//
// Solidity: event BidEvicted(uint256 indexed bid, bytes indexed nodeID)
func (_SlotAuctionManager *SlotAuctionManagerFilterer) WatchBidEvicted(opts *bind.WatchOpts, sink chan<- *SlotAuctionManagerBidEvicted, bid []*big.Int, nodeID [][]byte) (event.Subscription, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _SlotAuctionManager.contract.WatchLogs(opts, "BidEvicted", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlotAuctionManagerBidEvicted)
				if err := _SlotAuctionManager.contract.UnpackLog(event, "BidEvicted", log); err != nil {
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
func (_SlotAuctionManager *SlotAuctionManagerFilterer) ParseBidEvicted(log types.Log) (*SlotAuctionManagerBidEvicted, error) {
	event := new(SlotAuctionManagerBidEvicted)
	if err := _SlotAuctionManager.contract.UnpackLog(event, "BidEvicted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlotAuctionManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SlotAuctionManager contract.
type SlotAuctionManagerInitializedIterator struct {
	Event *SlotAuctionManagerInitialized // Event containing the contract specifics and raw log

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
func (it *SlotAuctionManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlotAuctionManagerInitialized)
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
		it.Event = new(SlotAuctionManagerInitialized)
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
func (it *SlotAuctionManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlotAuctionManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlotAuctionManagerInitialized represents a Initialized event raised by the SlotAuctionManager contract.
type SlotAuctionManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SlotAuctionManager *SlotAuctionManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*SlotAuctionManagerInitializedIterator, error) {

	logs, sub, err := _SlotAuctionManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SlotAuctionManagerInitializedIterator{contract: _SlotAuctionManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SlotAuctionManager *SlotAuctionManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SlotAuctionManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _SlotAuctionManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlotAuctionManagerInitialized)
				if err := _SlotAuctionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SlotAuctionManager *SlotAuctionManagerFilterer) ParseInitialized(log types.Log) (*SlotAuctionManagerInitialized, error) {
	event := new(SlotAuctionManagerInitialized)
	if err := _SlotAuctionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator is returned from FilterInitiatedAuctionValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedAuctionValidatorRegistration events raised by the SlotAuctionManager contract.
type SlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator struct {
	Event *SlotAuctionManagerInitiatedAuctionValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *SlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlotAuctionManagerInitiatedAuctionValidatorRegistration)
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
		it.Event = new(SlotAuctionManagerInitiatedAuctionValidatorRegistration)
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
func (it *SlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlotAuctionManagerInitiatedAuctionValidatorRegistration represents a InitiatedAuctionValidatorRegistration event raised by the SlotAuctionManager contract.
type SlotAuctionManagerInitiatedAuctionValidatorRegistration struct {
	ValidationID     [32]byte
	OwnerAddress     common.Address
	ValidatorEndTime *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInitiatedAuctionValidatorRegistration is a free log retrieval operation binding the contract event 0xccb9a8514540a67c432230e6261f90893e22e8e31f65c40e7629b5a47374dbbe.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 validatorEndTime)
func (_SlotAuctionManager *SlotAuctionManagerFilterer) FilterInitiatedAuctionValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, ownerAddress []common.Address) (*SlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _SlotAuctionManager.contract.FilterLogs(opts, "InitiatedAuctionValidatorRegistration", validationIDRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &SlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator{contract: _SlotAuctionManager.contract, event: "InitiatedAuctionValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedAuctionValidatorRegistration is a free log subscription operation binding the contract event 0xccb9a8514540a67c432230e6261f90893e22e8e31f65c40e7629b5a47374dbbe.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 validatorEndTime)
func (_SlotAuctionManager *SlotAuctionManagerFilterer) WatchInitiatedAuctionValidatorRegistration(opts *bind.WatchOpts, sink chan<- *SlotAuctionManagerInitiatedAuctionValidatorRegistration, validationID [][32]byte, ownerAddress []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _SlotAuctionManager.contract.WatchLogs(opts, "InitiatedAuctionValidatorRegistration", validationIDRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlotAuctionManagerInitiatedAuctionValidatorRegistration)
				if err := _SlotAuctionManager.contract.UnpackLog(event, "InitiatedAuctionValidatorRegistration", log); err != nil {
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

// ParseInitiatedAuctionValidatorRegistration is a log parse operation binding the contract event 0xccb9a8514540a67c432230e6261f90893e22e8e31f65c40e7629b5a47374dbbe.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 validatorEndTime)
func (_SlotAuctionManager *SlotAuctionManagerFilterer) ParseInitiatedAuctionValidatorRegistration(log types.Log) (*SlotAuctionManagerInitiatedAuctionValidatorRegistration, error) {
	event := new(SlotAuctionManagerInitiatedAuctionValidatorRegistration)
	if err := _SlotAuctionManager.contract.UnpackLog(event, "InitiatedAuctionValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlotAuctionManagerNewValidatorAuctionIterator is returned from FilterNewValidatorAuction and is used to iterate over the raw logs and unpacked data for NewValidatorAuction events raised by the SlotAuctionManager contract.
type SlotAuctionManagerNewValidatorAuctionIterator struct {
	Event *SlotAuctionManagerNewValidatorAuction // Event containing the contract specifics and raw log

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
func (it *SlotAuctionManagerNewValidatorAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlotAuctionManagerNewValidatorAuction)
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
		it.Event = new(SlotAuctionManagerNewValidatorAuction)
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
func (it *SlotAuctionManagerNewValidatorAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlotAuctionManagerNewValidatorAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlotAuctionManagerNewValidatorAuction represents a NewValidatorAuction event raised by the SlotAuctionManager contract.
type SlotAuctionManagerNewValidatorAuction struct {
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
func (_SlotAuctionManager *SlotAuctionManagerFilterer) FilterNewValidatorAuction(opts *bind.FilterOpts) (*SlotAuctionManagerNewValidatorAuctionIterator, error) {

	logs, sub, err := _SlotAuctionManager.contract.FilterLogs(opts, "NewValidatorAuction")
	if err != nil {
		return nil, err
	}
	return &SlotAuctionManagerNewValidatorAuctionIterator{contract: _SlotAuctionManager.contract, event: "NewValidatorAuction", logs: logs, sub: sub}, nil
}

// WatchNewValidatorAuction is a free log subscription operation binding the contract event 0x50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b54.
//
// Solidity: event NewValidatorAuction(uint16 validatorSlots, uint64 validatorWeight, uint256 minValidatorDuration, uint256 auctionEndTime, uint256 minimumBid)
func (_SlotAuctionManager *SlotAuctionManagerFilterer) WatchNewValidatorAuction(opts *bind.WatchOpts, sink chan<- *SlotAuctionManagerNewValidatorAuction) (event.Subscription, error) {

	logs, sub, err := _SlotAuctionManager.contract.WatchLogs(opts, "NewValidatorAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlotAuctionManagerNewValidatorAuction)
				if err := _SlotAuctionManager.contract.UnpackLog(event, "NewValidatorAuction", log); err != nil {
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
func (_SlotAuctionManager *SlotAuctionManagerFilterer) ParseNewValidatorAuction(log types.Log) (*SlotAuctionManagerNewValidatorAuction, error) {
	event := new(SlotAuctionManagerNewValidatorAuction)
	if err := _SlotAuctionManager.contract.UnpackLog(event, "NewValidatorAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlotAuctionManagerSuccessfulBidPlacedIterator is returned from FilterSuccessfulBidPlaced and is used to iterate over the raw logs and unpacked data for SuccessfulBidPlaced events raised by the SlotAuctionManager contract.
type SlotAuctionManagerSuccessfulBidPlacedIterator struct {
	Event *SlotAuctionManagerSuccessfulBidPlaced // Event containing the contract specifics and raw log

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
func (it *SlotAuctionManagerSuccessfulBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlotAuctionManagerSuccessfulBidPlaced)
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
		it.Event = new(SlotAuctionManagerSuccessfulBidPlaced)
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
func (it *SlotAuctionManagerSuccessfulBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlotAuctionManagerSuccessfulBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlotAuctionManagerSuccessfulBidPlaced represents a SuccessfulBidPlaced event raised by the SlotAuctionManager contract.
type SlotAuctionManagerSuccessfulBidPlaced struct {
	Bid    *big.Int
	NodeID common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSuccessfulBidPlaced is a free log retrieval operation binding the contract event 0x864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029.
//
// Solidity: event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID)
func (_SlotAuctionManager *SlotAuctionManagerFilterer) FilterSuccessfulBidPlaced(opts *bind.FilterOpts, bid []*big.Int, nodeID [][]byte) (*SlotAuctionManagerSuccessfulBidPlacedIterator, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _SlotAuctionManager.contract.FilterLogs(opts, "SuccessfulBidPlaced", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &SlotAuctionManagerSuccessfulBidPlacedIterator{contract: _SlotAuctionManager.contract, event: "SuccessfulBidPlaced", logs: logs, sub: sub}, nil
}

// WatchSuccessfulBidPlaced is a free log subscription operation binding the contract event 0x864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029.
//
// Solidity: event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID)
func (_SlotAuctionManager *SlotAuctionManagerFilterer) WatchSuccessfulBidPlaced(opts *bind.WatchOpts, sink chan<- *SlotAuctionManagerSuccessfulBidPlaced, bid []*big.Int, nodeID [][]byte) (event.Subscription, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _SlotAuctionManager.contract.WatchLogs(opts, "SuccessfulBidPlaced", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlotAuctionManagerSuccessfulBidPlaced)
				if err := _SlotAuctionManager.contract.UnpackLog(event, "SuccessfulBidPlaced", log); err != nil {
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
func (_SlotAuctionManager *SlotAuctionManagerFilterer) ParseSuccessfulBidPlaced(log types.Log) (*SlotAuctionManagerSuccessfulBidPlaced, error) {
	event := new(SlotAuctionManagerSuccessfulBidPlaced)
	if err := _SlotAuctionManager.contract.UnpackLog(event, "SuccessfulBidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
