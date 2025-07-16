// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenslotauctionmanager

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

// NativeTokenSlotAuctionManagerMetaData contains all meta data concerning the NativeTokenSlotAuctionManager contract.
var NativeTokenSlotAuctionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"}],\"name\":\"AuctionEndTimeNotReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionFinalizing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionNotInProgress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minumumBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBid\",\"type\":\"uint256\"}],\"name\":\"BidSmallerThanMinimum\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"DuplicateBidInContention\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"smallestAcceptableBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBid\",\"type\":\"uint256\"}],\"name\":\"InsufficientBidToWinAuction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeIsValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeIsWinning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationTimeLimit\",\"type\":\"uint256\"}],\"name\":\"ValidatorTimeLimitNotPassed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"BidEvicted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitiatedAuctionValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"validatorSlots\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minValidatorDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"}],\"name\":\"NewValidatorAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"SuccessfulBidPlaced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MinBidRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinValidatorDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinimumBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionState\",\"outputs\":[{\"internalType\":\"enumAuctionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"bidderInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"validatorslots\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minAuctionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValidatorDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"}],\"name\":\"initiateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSlots\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50604051612212380380612212833981016040819052602b91604e565b5f80546001600160a81b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f80fd5b81516001600160a01b03811681146072575f80fd5b9392505050565b61218c806100865f395ff3fe6080604052600436106100f9575f3560e01c8063913ed0b111610092578063ae9483e011610062578063ae9483e0146102b2578063b6e6a2ca146102e8578063c7d546f214610307578063f556f60a1461031a578063fe67a54b1461032f575f80fd5b8063913ed0b1146102415780639681d94014610255578063a3a65e4814610274578063a476f67514610293575f80fd5b80634b449cba116100cd5780634b449cba146101be5780636fe4d1e1146101e15780637a7df932146101f65780637fb4509914610215575f80fd5b8062d841f8146100fd57806305af82561461013f5780630d5daf3b146101605780631077830a14610191575b5f80fd5b348015610108575f80fd5b50600254610122906201000090046001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b34801561014a575f80fd5b5061015e610159366004611a63565b610343565b005b34801561016b575f80fd5b5061017f61017a366004611a7c565b6103b7565b60405161013696959493929190611b46565b34801561019c575f80fd5b506002546101ab9061ffff1681565b60405161ffff9091168152602001610136565b3480156101c9575f80fd5b506101d360015481565b604051908152602001610136565b3480156101ec575f80fd5b506101d360035481565b348015610201575f80fd5b5061015e610210366004611bb3565b6105ec565b348015610220575f80fd5b505f5461023490600160a01b900460ff1681565b6040516101369190611c25565b34801561024c575f80fd5b506101d361073b565b348015610260575f80fd5b506101d361026f366004611a63565b6107af565b34801561027f575f80fd5b506101d361028e366004611a63565b610827565b34801561029e575f80fd5b5061015e6102ad366004611a7c565b61085d565b3480156102bd575f80fd5b505f546102d0906001600160a01b031681565b6040516001600160a01b039091168152602001610136565b3480156102f3575f80fd5b5061015e610302366004611a7c565b6108b9565b61015e610315366004611df6565b610a0a565b348015610325575f80fd5b506101d360045481565b34801561033a575f80fd5b5061015e610a4e565b5f5460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af115801561038f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103b39190611e9b565b5050565b60086020525f90815260409020805460018201546002830180546001600160a01b039093169391926103e890611eb2565b80601f016020809104026020016040519081016040528092919081815260200182805461041490611eb2565b801561045f5780601f106104365761010080835404028352916020019161045f565b820191905f5260205f20905b81548152906001019060200180831161044257829003601f168201915b50505050509080600301805461047490611eb2565b80601f01602080910402602001604051908101604052809291908181526020018280546104a090611eb2565b80156104eb5780601f106104c2576101008083540402835291602001916104eb565b820191905f5260205f20905b8154815290600101906020018083116104ce57829003601f168201915b505060408051808201825260048701805463ffffffff16825260058801805484516020828102820181019096528181529899989397509195508387019450919283018282801561056257602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610544575b50505091909252505060408051808201825260068501805463ffffffff168252600786018054845160208281028201810190965281815296979693955091938086019392908301828280156105de57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116105c0575b505050505081525050905086565b60015f54600160a01b900460ff16600281111561060b5761060b611c11565b0361062957604051639bbb1f3b60e01b815260040160405180910390fd5b60025f54600160a01b900460ff16600281111561064857610648611c11565b0361066657604051631f86de2560e21b815260040160405180910390fd5b5f805460ff60a01b1916600160a01b1790556002805469ffffffffffffffff00001916620100006001600160401b038716021790556106a58342611efe565b6001556002805461ffff191661ffff8716179055600382905560048190555f60098190556106d590600a9061197f565b6002546004546040805161ffff90931683526001600160401b038716602084015282018490526060820185905260808201527f50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b549060a00160405180910390a15050505050565b5f60015f54600160a01b900460ff16600281111561075b5761075b611c11565b14610778576040516232f00d60e71b815260040160405180910390fd5b60025461ffff16610788600a5490565b1015610795575060045490565b61079f600a6110e2565b6107aa906001611efe565b905090565b5f805460405163025a076560e61b815263ffffffff841660048201526001600160a01b0390911690639681d940906024015b6020604051808303815f875af11580156107fd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108219190611e9b565b92915050565b5f8054604051631474cbc960e31b815263ffffffff841660048201526001600160a01b039091169063a3a65e48906024016107e1565b5f54604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015b5f604051808303815f87803b1580156108a0575f80fd5b505af11580156108b2573d5f803e3d5ffd5b5050505050565b5f81815260076020526040902060010154421015610905575f8181526007602052604090819020600101549051631a936b7b60e11b815260048101919091526024015b60405180910390fd5b600660075f8381526020019081526020015f206002016040516109289190611f11565b90815260405190819003602001902080546001600160a01b03191681555f6001820181905561095a600283018261199d565b610967600383015f61199d565b505f600482018190556005909101805467ffffffffffffffff1916905581815260076020526040812080546001600160a01b031916815560018101829055906109b3600283018261199d565b6109c0600383015f61199d565b505f60048281018290556005909201805467ffffffffffffffff1916905554604051635b73516560e11b81529182018390526001600160a01b03169063b6e6a2ca90602401610889565b610a12611107565b610a1f3485858585611151565b610a4860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050565b610a56611107565b60015f54600160a01b900460ff166002811115610a7557610a75611c11565b14610a92576040516232f00d60e71b815260040160405180910390fd5b5f805460ff60a01b1916600160a11b1790556001544211610acc57600154604051638230260760e01b81526004016108fc91815260200190565b5f600155600954158015610ae15750600a5415155b15610af457610af0600a6110e2565b6009555b5f610afe600a5490565b111561109b575f610b0f600a6115f4565b5f818152600860209081526040808320815160c08101835281546001600160a01b0316815260018201549381019390935260028101805495965093949293909291840191610b5c90611eb2565b80601f0160208091040260200160405190810160405280929190818152602001828054610b8890611eb2565b8015610bd35780601f10610baa57610100808354040283529160200191610bd3565b820191905f5260205f20905b815481529060010190602001808311610bb657829003601f168201915b50505050508152602001600382018054610bec90611eb2565b80601f0160208091040260200160405190810160405280929190818152602001828054610c1890611eb2565b8015610c635780601f10610c3a57610100808354040283529160200191610c63565b820191905f5260205f20905b815481529060010190602001808311610c4657829003601f168201915b505050918352505060408051808201825260048401805463ffffffff168252600585018054845160208281028201810190965281815295850195939492938584019390929190830182828015610ce057602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610cc2575b50505091909252505050815260408051808201825260068401805463ffffffff168252600785018054845160208281028201810190965281815295850195939492938584019390929190830182828015610d6157602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610d43575b505050919092525050509052508051600954919250610d8991610d849085611f83565b611601565b5f80546040838101516060850151608086015160a08701516002549451634e5bb12760e11b81526001600160a01b0390961695639cb7624e95610de39594939291620100009091046001600160401b031690600401611f96565b6020604051808303815f875af1158015610dff573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e239190611e9b565b9050815f01516001600160a01b0316817f032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285600154600354610e649190611efe565b60025460408051928352620100009091046001600160401b031660208301520160405180910390a36040518060c00160405280835f01516001600160a01b03168152602001600154600354610eb99190611efe565b8152602001836040015181526020018360600151815260200182815260200160028054906101000a90046001600160401b03166001600160401b031681525060068360400151604051610f0c9190611ffe565b9081526040805160209281900383019020835181546001600160a01b0319166001600160a01b039091161781559183015160018301558201516002820190610f54908261205d565b5060608201516003820190610f69908261205d565b506080820151600482015560a0909101516005909101805467ffffffffffffffff19166001600160401b039092169190911790556040805160c0810190915282516001600160a01b031681526001546003546020830191610fc991611efe565b8152604084810151602080840191909152606080870151838501528301859052600280546201000090046001600160401b03166080909401939093525f85815260078252829020845181546001600160a01b0319166001600160a01b03909116178155908401516001820155908301519091820190611048908261205d565b506060820151600382019061105d908261205d565b506080820151600482015560a0909101516005909101805467ffffffffffffffff19166001600160401b039092169190911790555050600955610af4565b6002805461ffff191690555f6003819055805460ff60a01b191690556110e060017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b565b5f815f015f815481106110f7576110f761211c565b905f5260205f2001549050919050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080546001190161114b57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60015f54600160a01b900460ff16600281111561117057611170611c11565b1461118d576040516232f00d60e71b815260040160405180910390fd5b5f5460405163d47a948b60e01b81526001600160a01b039091169063d47a948b906111bc908790600401612130565b602060405180830381865afa1580156111d7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111fb9190611e9b565b1561121b578360405163f3c815f760e01b81526004016108fc9190612130565b60058460405161122b9190611ffe565b9081526040519081900360200190205460ff161561125e578360405163f3c1d91160e01b81526004016108fc9190612130565b5f858152600860205260409020546001600160a01b03161561129657604051630517e2e760e21b8152600481018690526024016108fc565b8460045411156112c457600480546040516301a0013b60e61b815291820152602481018690526044016108fc565b60025461ffff166112d4600a5490565b10156112ea576112e5600a86611614565b611477565b846112f5600a6110e2565b101561143e575f611307600a87611621565b60098190555f8181526008602052604090819020905191925061132f91600290910190611f11565b6040519081900381209082907fdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc905f90a35f81815260086020526040902054611381906001600160a01b031682611601565b600560085f8381526020019081526020015f206002016040516113a49190611f11565b9081526040805160209281900383019020805460ff191690555f8381526008909252812080546001600160a01b031916815560018101829055906113eb600283018261199d565b6113f8600383015f61199d565b60048201805463ffffffff191681555f611415600585018261197f565b505060068201805463ffffffff191681555f611434600785018261197f565b5050505050611477565b611448600a6110e2565b611453906001611efe565b604051631a79674760e11b81526004810191909152602481018690526044016108fc565b6040805160c0810182523381526020808201888152828401888152606084018890526080840187905260a084018690525f8a8152600890935293909120825181546001600160a01b0319166001600160a01b0390911617815590516001820155915190919060028201906114eb908261205d565b5060608201516003820190611500908261205d565b506080820151805160048301805463ffffffff191663ffffffff909216919091178155602080830151805161153b92600587019201906119d4565b50505060a0820151805160068301805463ffffffff191663ffffffff909216919091178155602080830151805161157892600787019201906119d4565b50505090505060016005856040516115909190611ffe565b908152604051908190036020018120805492151560ff19909316929092179091556115bc908590611ffe565b6040519081900381209086907f864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029905f90a35050505050565b5f6108218261163661163a565b6103b36001600160a01b038316826116cb565b6103b3828261163661176a565b5f61162f8383611636611795565b9392505050565b1190565b5f80611644845490565b9050805f036116575761165760316117e2565b5f61166285826117f3565b5490505f611673865f1985016117f3565b54865490915086908061168857611688612142565b600190038181905f5260205f20015f90559055806116b15f885f016117f390919063ffffffff16565b556116c2865f1985015f848961181a565b50949350505050565b804710156116f55760405163cf47918160e01b8152476004820152602481018290526044016108fc565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f811461173e576040519150601f19603f3d011682016040523d82523d5f602084013e611743565b606091505b50509050806117655760405163d6bda27560e01b815260040160405180910390fd5b505050565b5f611773845490565b84546001810186555f86815260209020018490559050610a488482858561190f565b5f8061179f855490565b9050805f036117b2576117b260316117e2565b5f6117bd86826117f3565b549050846117cb875f6117f3565b556117d986835f888861181a565b95945050505050565b634e487b715f52806020526024601cfd5b5f8261181261180f8461180b845f9081526020902090565b0190565b90565b949350505050565b6001600160ff1b038310156108b2576002838102600181019101858110156118c3575f61184788846117f3565b5490505f61185589846117f3565b549050611866828763ffffffff8816565b8061187a575061187a81878763ffffffff16565b156118bc575f6118a0858561189386868b63ffffffff16565b1515918118919091021890565b90506118ad8a8983611957565b6118ba8a8a838a8a61181a565b505b5050611906565b85821015611906575f6118d688846117f3565b5490506118e7818663ffffffff8716565b15611904576118f7888785611957565b611904888885888861181a565b505b50505050505050565b8215610a485760025f198401045f61192786836117f3565b549050611938818563ffffffff8616565b15611944575050610a48565b61194f868684611957565b50925061190f565b5f61196284846117f3565b90505f61196f85846117f3565b8054835490915590915550505050565b5080545f8255905f5260205f209081019061199a9190611a37565b50565b5080546119a990611eb2565b5f825580601f106119b8575050565b601f0160209004905f5260205f209081019061199a9190611a37565b828054828255905f5260205f20908101928215611a27579160200282015b82811115611a2757825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906119f2565b50611a33929150611a37565b5090565b5b80821115611a33575f8155600101611a38565b803563ffffffff81168114611a5e575f80fd5b919050565b5f60208284031215611a73575f80fd5b61162f82611a4b565b5f60208284031215611a8c575f80fd5b5035919050565b5f5b83811015611aad578181015183820152602001611a95565b50505f910152565b5f8151808452611acc816020860160208601611a93565b601f01601f19169290920160200192915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611b3b5784516001600160a01b03168252938301936001929092019190830190611b12565b509695505050505050565b60018060a01b038716815285602082015260c060408201525f611b6c60c0830187611ab5565b8281036060840152611b7e8187611ab5565b90508281036080840152611b928186611ae0565b905082810360a0840152611ba68185611ae0565b9998505050505050505050565b5f805f805f60a08688031215611bc7575f80fd5b853561ffff81168114611bd8575f80fd5b945060208601356001600160401b0381168114611bf3575f80fd5b94979496505050506040830135926060810135926080909101359150565b634e487b7160e01b5f52602160045260245ffd5b6020810160038310611c4557634e487b7160e01b5f52602160045260245ffd5b91905290565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611c8157611c81611c4b565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611caf57611caf611c4b565b604052919050565b5f82601f830112611cc6575f80fd5b81356001600160401b03811115611cdf57611cdf611c4b565b611cf2601f8201601f1916602001611c87565b818152846020838601011115611d06575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60408284031215611d32575f80fd5b611d3a611c5f565b9050611d4582611a4b565b81526020808301356001600160401b0380821115611d61575f80fd5b818501915085601f830112611d74575f80fd5b813581811115611d8657611d86611c4b565b8060051b9150611d97848301611c87565b8181529183018401918481019088841115611db0575f80fd5b938501935b83851015611de557843592506001600160a01b0383168314611dd5575f80fd5b8282529385019390850190611db5565b808688015250505050505092915050565b5f805f8060808587031215611e09575f80fd5b84356001600160401b0380821115611e1f575f80fd5b611e2b88838901611cb7565b95506020870135915080821115611e40575f80fd5b611e4c88838901611cb7565b94506040870135915080821115611e61575f80fd5b611e6d88838901611d22565b93506060870135915080821115611e82575f80fd5b50611e8f87828801611d22565b91505092959194509250565b5f60208284031215611eab575f80fd5b5051919050565b600181811c90821680611ec657607f821691505b602082108103611ee457634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561082157610821611eea565b5f808354611f1e81611eb2565b60018281168015611f365760018114611f4b57611f77565b60ff1984168752821515830287019450611f77565b875f526020805f205f5b85811015611f6e5781548a820152908401908201611f55565b50505082870194505b50929695505050505050565b8181038181111561082157610821611eea565b60a081525f611fa860a0830188611ab5565b8281036020840152611fba8188611ab5565b90508281036040840152611fce8187611ae0565b90508281036060840152611fe28186611ae0565b9150506001600160401b03831660808301529695505050505050565b5f825161200f818460208701611a93565b9190910192915050565b601f82111561176557805f5260205f20601f840160051c8101602085101561203e5750805b601f840160051c820191505b818110156108b2575f815560010161204a565b81516001600160401b0381111561207657612076611c4b565b61208a816120848454611eb2565b84612019565b602080601f8311600181146120bd575f84156120a65750858301515b5f19600386901b1c1916600185901b178555612114565b5f85815260208120601f198616915b828110156120eb578886015182559484019460019091019084016120cc565b508582101561210857878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b634e487b7160e01b5f52603260045260245ffd5b602081525f61162f6020830184611ab5565b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220b501ca3d756690c35d8cd8bf8a3d4db14b3e8296df998bf02ada236ecd41426464736f6c63430008190033",
}

// NativeTokenSlotAuctionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenSlotAuctionManagerMetaData.ABI instead.
var NativeTokenSlotAuctionManagerABI = NativeTokenSlotAuctionManagerMetaData.ABI

// NativeTokenSlotAuctionManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenSlotAuctionManagerMetaData.Bin instead.
var NativeTokenSlotAuctionManagerBin = NativeTokenSlotAuctionManagerMetaData.Bin

// DeployNativeTokenSlotAuctionManager deploys a new Ethereum contract, binding an instance of NativeTokenSlotAuctionManager to it.
func DeployNativeTokenSlotAuctionManager(auth *bind.TransactOpts, backend bind.ContractBackend, vmAddress common.Address) (common.Address, *types.Transaction, *NativeTokenSlotAuctionManager, error) {
	parsed, err := NativeTokenSlotAuctionManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenSlotAuctionManagerBin), backend, vmAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenSlotAuctionManager{NativeTokenSlotAuctionManagerCaller: NativeTokenSlotAuctionManagerCaller{contract: contract}, NativeTokenSlotAuctionManagerTransactor: NativeTokenSlotAuctionManagerTransactor{contract: contract}, NativeTokenSlotAuctionManagerFilterer: NativeTokenSlotAuctionManagerFilterer{contract: contract}}, nil
}

// NativeTokenSlotAuctionManager is an auto generated Go binding around an Ethereum contract.
type NativeTokenSlotAuctionManager struct {
	NativeTokenSlotAuctionManagerCaller     // Read-only binding to the contract
	NativeTokenSlotAuctionManagerTransactor // Write-only binding to the contract
	NativeTokenSlotAuctionManagerFilterer   // Log filterer for contract events
}

// NativeTokenSlotAuctionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenSlotAuctionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSlotAuctionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenSlotAuctionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSlotAuctionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenSlotAuctionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSlotAuctionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenSlotAuctionManagerSession struct {
	Contract     *NativeTokenSlotAuctionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// NativeTokenSlotAuctionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenSlotAuctionManagerCallerSession struct {
	Contract *NativeTokenSlotAuctionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// NativeTokenSlotAuctionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenSlotAuctionManagerTransactorSession struct {
	Contract     *NativeTokenSlotAuctionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// NativeTokenSlotAuctionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenSlotAuctionManagerRaw struct {
	Contract *NativeTokenSlotAuctionManager // Generic contract binding to access the raw methods on
}

// NativeTokenSlotAuctionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenSlotAuctionManagerCallerRaw struct {
	Contract *NativeTokenSlotAuctionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenSlotAuctionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenSlotAuctionManagerTransactorRaw struct {
	Contract *NativeTokenSlotAuctionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenSlotAuctionManager creates a new instance of NativeTokenSlotAuctionManager, bound to a specific deployed contract.
func NewNativeTokenSlotAuctionManager(address common.Address, backend bind.ContractBackend) (*NativeTokenSlotAuctionManager, error) {
	contract, err := bindNativeTokenSlotAuctionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSlotAuctionManager{NativeTokenSlotAuctionManagerCaller: NativeTokenSlotAuctionManagerCaller{contract: contract}, NativeTokenSlotAuctionManagerTransactor: NativeTokenSlotAuctionManagerTransactor{contract: contract}, NativeTokenSlotAuctionManagerFilterer: NativeTokenSlotAuctionManagerFilterer{contract: contract}}, nil
}

// NewNativeTokenSlotAuctionManagerCaller creates a new read-only instance of NativeTokenSlotAuctionManager, bound to a specific deployed contract.
func NewNativeTokenSlotAuctionManagerCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenSlotAuctionManagerCaller, error) {
	contract, err := bindNativeTokenSlotAuctionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSlotAuctionManagerCaller{contract: contract}, nil
}

// NewNativeTokenSlotAuctionManagerTransactor creates a new write-only instance of NativeTokenSlotAuctionManager, bound to a specific deployed contract.
func NewNativeTokenSlotAuctionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenSlotAuctionManagerTransactor, error) {
	contract, err := bindNativeTokenSlotAuctionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSlotAuctionManagerTransactor{contract: contract}, nil
}

// NewNativeTokenSlotAuctionManagerFilterer creates a new log filterer instance of NativeTokenSlotAuctionManager, bound to a specific deployed contract.
func NewNativeTokenSlotAuctionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenSlotAuctionManagerFilterer, error) {
	contract, err := bindNativeTokenSlotAuctionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSlotAuctionManagerFilterer{contract: contract}, nil
}

// bindNativeTokenSlotAuctionManager binds a generic wrapper to an already deployed contract.
func bindNativeTokenSlotAuctionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenSlotAuctionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenSlotAuctionManager.Contract.NativeTokenSlotAuctionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.NativeTokenSlotAuctionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.NativeTokenSlotAuctionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenSlotAuctionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.contract.Transact(opts, method, params...)
}

// MinBidRequired is a free data retrieval call binding the contract method 0x913ed0b1.
//
// Solidity: function MinBidRequired() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCaller) MinBidRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSlotAuctionManager.contract.Call(opts, &out, "MinBidRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBidRequired is a free data retrieval call binding the contract method 0x913ed0b1.
//
// Solidity: function MinBidRequired() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) MinBidRequired() (*big.Int, error) {
	return _NativeTokenSlotAuctionManager.Contract.MinBidRequired(&_NativeTokenSlotAuctionManager.CallOpts)
}

// MinBidRequired is a free data retrieval call binding the contract method 0x913ed0b1.
//
// Solidity: function MinBidRequired() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCallerSession) MinBidRequired() (*big.Int, error) {
	return _NativeTokenSlotAuctionManager.Contract.MinBidRequired(&_NativeTokenSlotAuctionManager.CallOpts)
}

// MinValidatorDuration is a free data retrieval call binding the contract method 0x6fe4d1e1.
//
// Solidity: function MinValidatorDuration() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCaller) MinValidatorDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSlotAuctionManager.contract.Call(opts, &out, "MinValidatorDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValidatorDuration is a free data retrieval call binding the contract method 0x6fe4d1e1.
//
// Solidity: function MinValidatorDuration() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) MinValidatorDuration() (*big.Int, error) {
	return _NativeTokenSlotAuctionManager.Contract.MinValidatorDuration(&_NativeTokenSlotAuctionManager.CallOpts)
}

// MinValidatorDuration is a free data retrieval call binding the contract method 0x6fe4d1e1.
//
// Solidity: function MinValidatorDuration() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCallerSession) MinValidatorDuration() (*big.Int, error) {
	return _NativeTokenSlotAuctionManager.Contract.MinValidatorDuration(&_NativeTokenSlotAuctionManager.CallOpts)
}

// MinimumBid is a free data retrieval call binding the contract method 0xf556f60a.
//
// Solidity: function MinimumBid() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCaller) MinimumBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSlotAuctionManager.contract.Call(opts, &out, "MinimumBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumBid is a free data retrieval call binding the contract method 0xf556f60a.
//
// Solidity: function MinimumBid() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) MinimumBid() (*big.Int, error) {
	return _NativeTokenSlotAuctionManager.Contract.MinimumBid(&_NativeTokenSlotAuctionManager.CallOpts)
}

// MinimumBid is a free data retrieval call binding the contract method 0xf556f60a.
//
// Solidity: function MinimumBid() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCallerSession) MinimumBid() (*big.Int, error) {
	return _NativeTokenSlotAuctionManager.Contract.MinimumBid(&_NativeTokenSlotAuctionManager.CallOpts)
}

// VALIDATORMANAGER is a free data retrieval call binding the contract method 0xae9483e0.
//
// Solidity: function VALIDATOR_MANAGER() view returns(address)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCaller) VALIDATORMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSlotAuctionManager.contract.Call(opts, &out, "VALIDATOR_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORMANAGER is a free data retrieval call binding the contract method 0xae9483e0.
//
// Solidity: function VALIDATOR_MANAGER() view returns(address)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) VALIDATORMANAGER() (common.Address, error) {
	return _NativeTokenSlotAuctionManager.Contract.VALIDATORMANAGER(&_NativeTokenSlotAuctionManager.CallOpts)
}

// VALIDATORMANAGER is a free data retrieval call binding the contract method 0xae9483e0.
//
// Solidity: function VALIDATOR_MANAGER() view returns(address)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCallerSession) VALIDATORMANAGER() (common.Address, error) {
	return _NativeTokenSlotAuctionManager.Contract.VALIDATORMANAGER(&_NativeTokenSlotAuctionManager.CallOpts)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCaller) AuctionEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSlotAuctionManager.contract.Call(opts, &out, "auctionEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) AuctionEndTime() (*big.Int, error) {
	return _NativeTokenSlotAuctionManager.Contract.AuctionEndTime(&_NativeTokenSlotAuctionManager.CallOpts)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCallerSession) AuctionEndTime() (*big.Int, error) {
	return _NativeTokenSlotAuctionManager.Contract.AuctionEndTime(&_NativeTokenSlotAuctionManager.CallOpts)
}

// AuctionState is a free data retrieval call binding the contract method 0x7fb45099.
//
// Solidity: function auctionState() view returns(uint8)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCaller) AuctionState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenSlotAuctionManager.contract.Call(opts, &out, "auctionState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AuctionState is a free data retrieval call binding the contract method 0x7fb45099.
//
// Solidity: function auctionState() view returns(uint8)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) AuctionState() (uint8, error) {
	return _NativeTokenSlotAuctionManager.Contract.AuctionState(&_NativeTokenSlotAuctionManager.CallOpts)
}

// AuctionState is a free data retrieval call binding the contract method 0x7fb45099.
//
// Solidity: function auctionState() view returns(uint8)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCallerSession) AuctionState() (uint8, error) {
	return _NativeTokenSlotAuctionManager.Contract.AuctionState(&_NativeTokenSlotAuctionManager.CallOpts)
}

// BidderInfo is a free data retrieval call binding the contract method 0x0d5daf3b.
//
// Solidity: function bidderInfo(uint256 bid) view returns(address addr, uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCaller) BidderInfo(opts *bind.CallOpts, bid *big.Int) (struct {
	Addr                  common.Address
	Bid                   *big.Int
	NodeID                []byte
	BlsPublicKey          []byte
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}, error) {
	var out []interface{}
	err := _NativeTokenSlotAuctionManager.contract.Call(opts, &out, "bidderInfo", bid)

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
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) BidderInfo(bid *big.Int) (struct {
	Addr                  common.Address
	Bid                   *big.Int
	NodeID                []byte
	BlsPublicKey          []byte
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}, error) {
	return _NativeTokenSlotAuctionManager.Contract.BidderInfo(&_NativeTokenSlotAuctionManager.CallOpts, bid)
}

// BidderInfo is a free data retrieval call binding the contract method 0x0d5daf3b.
//
// Solidity: function bidderInfo(uint256 bid) view returns(address addr, uint256 bid, bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCallerSession) BidderInfo(bid *big.Int) (struct {
	Addr                  common.Address
	Bid                   *big.Int
	NodeID                []byte
	BlsPublicKey          []byte
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}, error) {
	return _NativeTokenSlotAuctionManager.Contract.BidderInfo(&_NativeTokenSlotAuctionManager.CallOpts, bid)
}

// ValidatorSlots is a free data retrieval call binding the contract method 0x1077830a.
//
// Solidity: function validatorSlots() view returns(uint16)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCaller) ValidatorSlots(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NativeTokenSlotAuctionManager.contract.Call(opts, &out, "validatorSlots")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ValidatorSlots is a free data retrieval call binding the contract method 0x1077830a.
//
// Solidity: function validatorSlots() view returns(uint16)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) ValidatorSlots() (uint16, error) {
	return _NativeTokenSlotAuctionManager.Contract.ValidatorSlots(&_NativeTokenSlotAuctionManager.CallOpts)
}

// ValidatorSlots is a free data retrieval call binding the contract method 0x1077830a.
//
// Solidity: function validatorSlots() view returns(uint16)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCallerSession) ValidatorSlots() (uint16, error) {
	return _NativeTokenSlotAuctionManager.Contract.ValidatorSlots(&_NativeTokenSlotAuctionManager.CallOpts)
}

// ValidatorWeight is a free data retrieval call binding the contract method 0x00d841f8.
//
// Solidity: function validatorWeight() view returns(uint64)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCaller) ValidatorWeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NativeTokenSlotAuctionManager.contract.Call(opts, &out, "validatorWeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValidatorWeight is a free data retrieval call binding the contract method 0x00d841f8.
//
// Solidity: function validatorWeight() view returns(uint64)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) ValidatorWeight() (uint64, error) {
	return _NativeTokenSlotAuctionManager.Contract.ValidatorWeight(&_NativeTokenSlotAuctionManager.CallOpts)
}

// ValidatorWeight is a free data retrieval call binding the contract method 0x00d841f8.
//
// Solidity: function validatorWeight() view returns(uint64)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCallerSession) ValidatorWeight() (uint64, error) {
	return _NativeTokenSlotAuctionManager.Contract.ValidatorWeight(&_NativeTokenSlotAuctionManager.CallOpts)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactor) CompleteRemoveInitialValidator(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.contract.Transact(opts, "completeRemoveInitialValidator", messageIndex)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) CompleteRemoveInitialValidator(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.CompleteRemoveInitialValidator(&_NativeTokenSlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteRemoveInitialValidator is a paid mutator transaction binding the contract method 0x05af8256.
//
// Solidity: function completeRemoveInitialValidator(uint32 messageIndex) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactorSession) CompleteRemoveInitialValidator(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.CompleteRemoveInitialValidator(&_NativeTokenSlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.CompleteValidatorRegistration(&_NativeTokenSlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.CompleteValidatorRegistration(&_NativeTokenSlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.CompleteValidatorRemoval(&_NativeTokenSlotAuctionManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.CompleteValidatorRemoval(&_NativeTokenSlotAuctionManager.TransactOpts, messageIndex)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) EndAuction() (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.EndAuction(&_NativeTokenSlotAuctionManager.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactorSession) EndAuction() (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.EndAuction(&_NativeTokenSlotAuctionManager.TransactOpts)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x7a7df932.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 minAuctionDuration, uint256 minValidatorDuration, uint256 minimumBid) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactor) InitiateAuction(opts *bind.TransactOpts, validatorslots uint16, weight uint64, minAuctionDuration *big.Int, minValidatorDuration *big.Int, minimumBid *big.Int) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.contract.Transact(opts, "initiateAuction", validatorslots, weight, minAuctionDuration, minValidatorDuration, minimumBid)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x7a7df932.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 minAuctionDuration, uint256 minValidatorDuration, uint256 minimumBid) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) InitiateAuction(validatorslots uint16, weight uint64, minAuctionDuration *big.Int, minValidatorDuration *big.Int, minimumBid *big.Int) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.InitiateAuction(&_NativeTokenSlotAuctionManager.TransactOpts, validatorslots, weight, minAuctionDuration, minValidatorDuration, minimumBid)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x7a7df932.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 minAuctionDuration, uint256 minValidatorDuration, uint256 minimumBid) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactorSession) InitiateAuction(validatorslots uint16, weight uint64, minAuctionDuration *big.Int, minValidatorDuration *big.Int, minimumBid *big.Int) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.InitiateAuction(&_NativeTokenSlotAuctionManager.TransactOpts, validatorslots, weight, minAuctionDuration, minValidatorDuration, minimumBid)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactor) InitiateRemoveInitialValidator(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.contract.Transact(opts, "initiateRemoveInitialValidator", validationID)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) InitiateRemoveInitialValidator(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.InitiateRemoveInitialValidator(&_NativeTokenSlotAuctionManager.TransactOpts, validationID)
}

// InitiateRemoveInitialValidator is a paid mutator transaction binding the contract method 0xa476f675.
//
// Solidity: function initiateRemoveInitialValidator(bytes32 validationID) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactorSession) InitiateRemoveInitialValidator(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.InitiateRemoveInitialValidator(&_NativeTokenSlotAuctionManager.TransactOpts, validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.contract.Transact(opts, "initiateValidatorRemoval", validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.InitiateValidatorRemoval(&_NativeTokenSlotAuctionManager.TransactOpts, validationID)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb6e6a2ca.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID) returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.InitiateValidatorRemoval(&_NativeTokenSlotAuctionManager.TransactOpts, validationID)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xc7d546f2.
//
// Solidity: function placeBid(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) payable returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactor) PlaceBid(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.contract.Transact(opts, "placeBid", nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xc7d546f2.
//
// Solidity: function placeBid(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) payable returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) PlaceBid(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.PlaceBid(&_NativeTokenSlotAuctionManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xc7d546f2.
//
// Solidity: function placeBid(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner) payable returns()
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerTransactorSession) PlaceBid(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner) (*types.Transaction, error) {
	return _NativeTokenSlotAuctionManager.Contract.PlaceBid(&_NativeTokenSlotAuctionManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner)
}

// NativeTokenSlotAuctionManagerBidEvictedIterator is returned from FilterBidEvicted and is used to iterate over the raw logs and unpacked data for BidEvicted events raised by the NativeTokenSlotAuctionManager contract.
type NativeTokenSlotAuctionManagerBidEvictedIterator struct {
	Event *NativeTokenSlotAuctionManagerBidEvicted // Event containing the contract specifics and raw log

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
func (it *NativeTokenSlotAuctionManagerBidEvictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSlotAuctionManagerBidEvicted)
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
		it.Event = new(NativeTokenSlotAuctionManagerBidEvicted)
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
func (it *NativeTokenSlotAuctionManagerBidEvictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSlotAuctionManagerBidEvictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSlotAuctionManagerBidEvicted represents a BidEvicted event raised by the NativeTokenSlotAuctionManager contract.
type NativeTokenSlotAuctionManagerBidEvicted struct {
	Bid    *big.Int
	NodeID common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidEvicted is a free log retrieval operation binding the contract event 0xdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc.
//
// Solidity: event BidEvicted(uint256 indexed bid, bytes indexed nodeID)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) FilterBidEvicted(opts *bind.FilterOpts, bid []*big.Int, nodeID [][]byte) (*NativeTokenSlotAuctionManagerBidEvictedIterator, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenSlotAuctionManager.contract.FilterLogs(opts, "BidEvicted", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSlotAuctionManagerBidEvictedIterator{contract: _NativeTokenSlotAuctionManager.contract, event: "BidEvicted", logs: logs, sub: sub}, nil
}

// WatchBidEvicted is a free log subscription operation binding the contract event 0xdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc.
//
// Solidity: event BidEvicted(uint256 indexed bid, bytes indexed nodeID)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) WatchBidEvicted(opts *bind.WatchOpts, sink chan<- *NativeTokenSlotAuctionManagerBidEvicted, bid []*big.Int, nodeID [][]byte) (event.Subscription, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenSlotAuctionManager.contract.WatchLogs(opts, "BidEvicted", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSlotAuctionManagerBidEvicted)
				if err := _NativeTokenSlotAuctionManager.contract.UnpackLog(event, "BidEvicted", log); err != nil {
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
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) ParseBidEvicted(log types.Log) (*NativeTokenSlotAuctionManagerBidEvicted, error) {
	event := new(NativeTokenSlotAuctionManagerBidEvicted)
	if err := _NativeTokenSlotAuctionManager.contract.UnpackLog(event, "BidEvicted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSlotAuctionManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenSlotAuctionManager contract.
type NativeTokenSlotAuctionManagerInitializedIterator struct {
	Event *NativeTokenSlotAuctionManagerInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenSlotAuctionManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSlotAuctionManagerInitialized)
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
		it.Event = new(NativeTokenSlotAuctionManagerInitialized)
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
func (it *NativeTokenSlotAuctionManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSlotAuctionManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSlotAuctionManagerInitialized represents a Initialized event raised by the NativeTokenSlotAuctionManager contract.
type NativeTokenSlotAuctionManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenSlotAuctionManagerInitializedIterator, error) {

	logs, sub, err := _NativeTokenSlotAuctionManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenSlotAuctionManagerInitializedIterator{contract: _NativeTokenSlotAuctionManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenSlotAuctionManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenSlotAuctionManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSlotAuctionManagerInitialized)
				if err := _NativeTokenSlotAuctionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) ParseInitialized(log types.Log) (*NativeTokenSlotAuctionManagerInitialized, error) {
	event := new(NativeTokenSlotAuctionManagerInitialized)
	if err := _NativeTokenSlotAuctionManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator is returned from FilterInitiatedAuctionValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedAuctionValidatorRegistration events raised by the NativeTokenSlotAuctionManager contract.
type NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator struct {
	Event *NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistration)
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
		it.Event = new(NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistration)
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
func (it *NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistration represents a InitiatedAuctionValidatorRegistration event raised by the NativeTokenSlotAuctionManager contract.
type NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistration struct {
	ValidationID     [32]byte
	OwnerAddress     common.Address
	ValidatorEndTime *big.Int
	Weight           uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInitiatedAuctionValidatorRegistration is a free log retrieval operation binding the contract event 0x032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 validatorEndTime, uint64 weight)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) FilterInitiatedAuctionValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, ownerAddress []common.Address) (*NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _NativeTokenSlotAuctionManager.contract.FilterLogs(opts, "InitiatedAuctionValidatorRegistration", validationIDRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistrationIterator{contract: _NativeTokenSlotAuctionManager.contract, event: "InitiatedAuctionValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedAuctionValidatorRegistration is a free log subscription operation binding the contract event 0x032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 validatorEndTime, uint64 weight)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) WatchInitiatedAuctionValidatorRegistration(opts *bind.WatchOpts, sink chan<- *NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistration, validationID [][32]byte, ownerAddress []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _NativeTokenSlotAuctionManager.contract.WatchLogs(opts, "InitiatedAuctionValidatorRegistration", validationIDRule, ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistration)
				if err := _NativeTokenSlotAuctionManager.contract.UnpackLog(event, "InitiatedAuctionValidatorRegistration", log); err != nil {
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
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) ParseInitiatedAuctionValidatorRegistration(log types.Log) (*NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistration, error) {
	event := new(NativeTokenSlotAuctionManagerInitiatedAuctionValidatorRegistration)
	if err := _NativeTokenSlotAuctionManager.contract.UnpackLog(event, "InitiatedAuctionValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSlotAuctionManagerNewValidatorAuctionIterator is returned from FilterNewValidatorAuction and is used to iterate over the raw logs and unpacked data for NewValidatorAuction events raised by the NativeTokenSlotAuctionManager contract.
type NativeTokenSlotAuctionManagerNewValidatorAuctionIterator struct {
	Event *NativeTokenSlotAuctionManagerNewValidatorAuction // Event containing the contract specifics and raw log

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
func (it *NativeTokenSlotAuctionManagerNewValidatorAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSlotAuctionManagerNewValidatorAuction)
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
		it.Event = new(NativeTokenSlotAuctionManagerNewValidatorAuction)
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
func (it *NativeTokenSlotAuctionManagerNewValidatorAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSlotAuctionManagerNewValidatorAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSlotAuctionManagerNewValidatorAuction represents a NewValidatorAuction event raised by the NativeTokenSlotAuctionManager contract.
type NativeTokenSlotAuctionManagerNewValidatorAuction struct {
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
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) FilterNewValidatorAuction(opts *bind.FilterOpts) (*NativeTokenSlotAuctionManagerNewValidatorAuctionIterator, error) {

	logs, sub, err := _NativeTokenSlotAuctionManager.contract.FilterLogs(opts, "NewValidatorAuction")
	if err != nil {
		return nil, err
	}
	return &NativeTokenSlotAuctionManagerNewValidatorAuctionIterator{contract: _NativeTokenSlotAuctionManager.contract, event: "NewValidatorAuction", logs: logs, sub: sub}, nil
}

// WatchNewValidatorAuction is a free log subscription operation binding the contract event 0x50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b54.
//
// Solidity: event NewValidatorAuction(uint16 validatorSlots, uint64 validatorWeight, uint256 minValidatorDuration, uint256 auctionEndTime, uint256 minimumBid)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) WatchNewValidatorAuction(opts *bind.WatchOpts, sink chan<- *NativeTokenSlotAuctionManagerNewValidatorAuction) (event.Subscription, error) {

	logs, sub, err := _NativeTokenSlotAuctionManager.contract.WatchLogs(opts, "NewValidatorAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSlotAuctionManagerNewValidatorAuction)
				if err := _NativeTokenSlotAuctionManager.contract.UnpackLog(event, "NewValidatorAuction", log); err != nil {
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
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) ParseNewValidatorAuction(log types.Log) (*NativeTokenSlotAuctionManagerNewValidatorAuction, error) {
	event := new(NativeTokenSlotAuctionManagerNewValidatorAuction)
	if err := _NativeTokenSlotAuctionManager.contract.UnpackLog(event, "NewValidatorAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSlotAuctionManagerSuccessfulBidPlacedIterator is returned from FilterSuccessfulBidPlaced and is used to iterate over the raw logs and unpacked data for SuccessfulBidPlaced events raised by the NativeTokenSlotAuctionManager contract.
type NativeTokenSlotAuctionManagerSuccessfulBidPlacedIterator struct {
	Event *NativeTokenSlotAuctionManagerSuccessfulBidPlaced // Event containing the contract specifics and raw log

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
func (it *NativeTokenSlotAuctionManagerSuccessfulBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSlotAuctionManagerSuccessfulBidPlaced)
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
		it.Event = new(NativeTokenSlotAuctionManagerSuccessfulBidPlaced)
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
func (it *NativeTokenSlotAuctionManagerSuccessfulBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSlotAuctionManagerSuccessfulBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSlotAuctionManagerSuccessfulBidPlaced represents a SuccessfulBidPlaced event raised by the NativeTokenSlotAuctionManager contract.
type NativeTokenSlotAuctionManagerSuccessfulBidPlaced struct {
	Bid    *big.Int
	NodeID common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSuccessfulBidPlaced is a free log retrieval operation binding the contract event 0x864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029.
//
// Solidity: event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) FilterSuccessfulBidPlaced(opts *bind.FilterOpts, bid []*big.Int, nodeID [][]byte) (*NativeTokenSlotAuctionManagerSuccessfulBidPlacedIterator, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenSlotAuctionManager.contract.FilterLogs(opts, "SuccessfulBidPlaced", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSlotAuctionManagerSuccessfulBidPlacedIterator{contract: _NativeTokenSlotAuctionManager.contract, event: "SuccessfulBidPlaced", logs: logs, sub: sub}, nil
}

// WatchSuccessfulBidPlaced is a free log subscription operation binding the contract event 0x864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029.
//
// Solidity: event SuccessfulBidPlaced(uint256 indexed bid, bytes indexed nodeID)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) WatchSuccessfulBidPlaced(opts *bind.WatchOpts, sink chan<- *NativeTokenSlotAuctionManagerSuccessfulBidPlaced, bid []*big.Int, nodeID [][]byte) (event.Subscription, error) {

	var bidRule []interface{}
	for _, bidItem := range bid {
		bidRule = append(bidRule, bidItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenSlotAuctionManager.contract.WatchLogs(opts, "SuccessfulBidPlaced", bidRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSlotAuctionManagerSuccessfulBidPlaced)
				if err := _NativeTokenSlotAuctionManager.contract.UnpackLog(event, "SuccessfulBidPlaced", log); err != nil {
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
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerFilterer) ParseSuccessfulBidPlaced(log types.Log) (*NativeTokenSlotAuctionManagerSuccessfulBidPlaced, error) {
	event := new(NativeTokenSlotAuctionManagerSuccessfulBidPlaced)
	if err := _NativeTokenSlotAuctionManager.contract.UnpackLog(event, "SuccessfulBidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
