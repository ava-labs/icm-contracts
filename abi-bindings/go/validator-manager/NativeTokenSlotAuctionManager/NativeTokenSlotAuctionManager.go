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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"}],\"name\":\"AuctionEndTimeNotReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionFinalizing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionNotInProgress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minumumBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBid\",\"type\":\"uint256\"}],\"name\":\"BidSmallerThanMinimum\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"DuplicateBidInContention\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"smallestAcceptableBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBid\",\"type\":\"uint256\"}],\"name\":\"InsufficientBidToWinAuction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeIsValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeIsWinning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationTimeLimit\",\"type\":\"uint256\"}],\"name\":\"ValidatorTimeLimitNotPassed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"BidEvicted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitiatedAuctionValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"validatorSlots\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minValidatorDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"}],\"name\":\"NewValidatorAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"SuccessfulBidPlaced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MinBidRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinValidatorDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinimumBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"bidderInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"validatorslots\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minAuctionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValidatorDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"}],\"name\":\"initiateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSlots\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50604051612277380380612277833981016040819052602b91604f565b600180546001600160a81b0319166001600160a01b0392909216919091179055607a565b5f60208284031215605e575f80fd5b81516001600160a01b03811681146073575f80fd5b9392505050565b6121f0806100875f395ff3fe6080604052600436106100f9575f3560e01c8063913ed0b111610092578063ae9483e011610062578063ae9483e0146102bc578063b6e6a2ca146102db578063c7d546f2146102fa578063f556f60a1461030d578063fe67a54b14610322575f80fd5b8063913ed0b11461024b5780639681d9401461025f578063a3a65e481461027e578063a476f6751461029d575f80fd5b806336339388116100cd57806336339388146101be5780634b449cba146101f45780636fe4d1e1146102175780637a7df9321461022c575f80fd5b8062d841f8146100fd57806305af82561461013f5780630d5daf3b146101605780631077830a14610191575b5f80fd5b348015610108575f80fd5b50600354610122906201000090046001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b34801561014a575f80fd5b5061015e610159366004611aae565b610336565b005b34801561016b575f80fd5b5061017f61017a366004611ac7565b6103ab565b60405161013696959493929190611b91565b34801561019c575f80fd5b506003546101ab9061ffff1681565b60405161ffff9091168152602001610136565b3480156101c9575f80fd5b505f546101dc906001600160a01b031681565b6040516001600160a01b039091168152602001610136565b3480156101ff575f80fd5b5061020960025481565b604051908152602001610136565b348015610222575f80fd5b5061020960045481565b348015610237575f80fd5b5061015e610246366004611bfe565b6105e0565b348015610256575f80fd5b50610209610731565b34801561026a575f80fd5b50610209610279366004611aae565b6107a5565b348015610289575f80fd5b50610209610298366004611aae565b61081d565b3480156102a8575f80fd5b5061015e6102b7366004611ac7565b610853565b3480156102c7575f80fd5b506001546101dc906001600160a01b031681565b3480156102e6575f80fd5b5061015e6102f5366004611ac7565b6108b0565b61015e610308366004611e07565b610a04565b348015610318575f80fd5b5061020960055481565b34801561032d575f80fd5b5061015e610a35565b60015460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af1158015610383573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103a79190611eac565b5050565b60096020525f90815260409020805460018201546002830180546001600160a01b039093169391926103dc90611ec3565b80601f016020809104026020016040519081016040528092919081815260200182805461040890611ec3565b80156104535780601f1061042a57610100808354040283529160200191610453565b820191905f5260205f20905b81548152906001019060200180831161043657829003601f168201915b50505050509080600301805461046890611ec3565b80601f016020809104026020016040519081016040528092919081815260200182805461049490611ec3565b80156104df5780601f106104b6576101008083540402835291602001916104df565b820191905f5260205f20905b8154815290600101906020018083116104c257829003601f168201915b505060408051808201825260048701805463ffffffff16825260058801805484516020828102820181019096528181529899989397509195508387019450919283018282801561055657602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610538575b50505091909252505060408051808201825260068501805463ffffffff168252600786018054845160208281028201810190965281815296979693955091938086019392908301828280156105d257602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116105b4575b505050505081525050905086565b60018054600160a01b900460ff1660028111156105ff576105ff611efb565b0361061d57604051639bbb1f3b60e01b815260040160405180910390fd5b6002600154600160a01b900460ff16600281111561063d5761063d611efb565b0361065b57604051631f86de2560e21b815260040160405180910390fd5b6001805460ff60a01b1916600160a01b1790556003805469ffffffffffffffff00001916620100006001600160401b0387160217905561069b8342611f23565b6002556003805461ffff191661ffff8716179055600482905560058190555f600a8190556106cb90600b906119ca565b6003546005546040805161ffff90931683526001600160401b038716602084015282018490526060820185905260808201527f50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b549060a00160405180910390a15050505050565b5f60018054600160a01b900460ff16600281111561075157610751611efb565b1461076e576040516232f00d60e71b815260040160405180910390fd5b60035461ffff1661077e600b5490565b101561078b575060055490565b610795600b6110ad565b6107a0906001611f23565b905090565b60015460405163025a076560e61b815263ffffffff831660048201525f916001600160a01b031690639681d940906024015b6020604051808303815f875af11580156107f3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108179190611eac565b92915050565b600154604051631474cbc960e31b815263ffffffff831660048201525f916001600160a01b03169063a3a65e48906024016107d7565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015b5f604051808303815f87803b158015610897575f80fd5b505af11580156108a9573d5f803e3d5ffd5b5050505050565b5f818152600860205260409020600101544210156108fc575f8181526008602052604090819020600101549051631a936b7b60e11b815260048101919091526024015b60405180910390fd5b600760085f8381526020019081526020015f2060020160405161091f9190611f36565b90815260405190819003602001902080546001600160a01b03191681555f6001820181905561095160028301826119e8565b61095e600383015f6119e8565b505f600482018190556005909101805467ffffffffffffffff1916905581815260086020526040812080546001600160a01b031916815560018101829055906109aa60028301826119e8565b6109b7600383015f6119e8565b505f6004828101919091556005909101805467ffffffffffffffff19169055600154604051635b73516560e11b81529182018390526001600160a01b03169063b6e6a2ca90602401610880565b610a0c6110d2565b610a193485858585611109565b610a2f60015f8051602061219b83398151915255565b50505050565b610a3d6110d2565b60018054600160a01b900460ff166002811115610a5c57610a5c611efb565b14610a79576040516232f00d60e71b815260040160405180910390fd5b6001805460ff60a01b1916600160a11b1790556002544211610ab457600254604051638230260760e01b81526004016108f391815260200190565b5f600255600a54158015610ac95750600b5415155b15610adc57610ad8600b6110ad565b600a555b5f610ae6600b5490565b1115611080575f610af7600b61163f565b5f818152600960209081526040808320815160c08101835281546001600160a01b0316815260018201549381019390935260028101805495965093949293909291840191610b4490611ec3565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7090611ec3565b8015610bbb5780601f10610b9257610100808354040283529160200191610bbb565b820191905f5260205f20905b815481529060010190602001808311610b9e57829003601f168201915b50505050508152602001600382018054610bd490611ec3565b80601f0160208091040260200160405190810160405280929190818152602001828054610c0090611ec3565b8015610c4b5780601f10610c2257610100808354040283529160200191610c4b565b820191905f5260205f20905b815481529060010190602001808311610c2e57829003601f168201915b505050918352505060408051808201825260048401805463ffffffff168252600585018054845160208281028201810190965281815295850195939492938584019390929190830182828015610cc857602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610caa575b50505091909252505050815260408051808201825260068401805463ffffffff168252600785018054845160208281028201810190965281815295850195939492938584019390929190830182828015610d4957602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610d2b575b505050919092525050509052508051600a54919250610d7191610d6c9085611fa8565b61164c565b6001546040828101516060840151608085015160a08601516003549451634e5bb12760e11b81525f966001600160a01b031695639cb7624e95610dce9590949093909290916201000090046001600160401b031690600401611fbb565b6020604051808303815f875af1158015610dea573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e0e9190611eac565b9050815f01516001600160a01b0316817f032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285600254600454610e4f9190611f23565b60035460408051928352620100009091046001600160401b031660208301520160405180910390a36040518060c00160405280835f01516001600160a01b03168152602001600254600454610ea49190611f23565b81526040848101805160208401526060808701518385015283018590526003546201000090046001600160401b031660809093019290925290519051600791610eec91612023565b9081526040805160209281900383019020835181546001600160a01b0319166001600160a01b039091161781559183015160018301558201516002820190610f349082612082565b5060608201516003820190610f499082612082565b506080820151816004015560a0820151816005015f6101000a8154816001600160401b0302191690836001600160401b031602179055509050506040518060c00160405280835f01516001600160a01b03168152602001600254600454610fb09190611f23565b81526040848101516020808401919091526060808701518385015283018590526003546201000090046001600160401b03166080909301929092525f84815260088352819020835181546001600160a01b0319166001600160a01b03909116178155918301516001830155820151600282019061102d9082612082565b50606082015160038201906110429082612082565b506080820151600482015560a0909101516005909101805467ffffffffffffffff19166001600160401b039092169190911790555050600a55610adc565b6003805461ffff191690555f6004556001805460ff60a01b191681555f8051602061219b83398151915255565b5f815f015f815481106110c2576110c2612141565b905f5260205f2001549050919050565b5f8051602061219b83398151915280546001190161110357604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6111116110d2565b60018054600160a01b900460ff16600281111561113057611130611efb565b1461114d576040516232f00d60e71b815260040160405180910390fd5b60015460405163d47a948b60e01b81526001600160a01b039091169063d47a948b9061117d908790600401612155565b602060405180830381865afa158015611198573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111bc9190611eac565b156111dc578360405163f3c815f760e01b81526004016108f39190612155565b6006846040516111ec9190612023565b9081526040519081900360200190205460ff161561121f578360405163f3c1d91160e01b81526004016108f39190612155565b5f858152600960205260409020546001600160a01b03161561125757604051630517e2e760e21b8152600481018690526024016108f3565b846005541115611288576005546040516301a0013b60e61b81526004810191909152602481018690526044016108f3565b60035461ffff16611298600b5490565b1015611326575f546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af11580156112f1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113159190612167565b50611321600b8661165f565b6114b3565b84611331600b6110ad565b101561147a575f611343600b8761166c565b600a8190555f8181526009602052604090819020905191925061136b91600290910190611f36565b6040519081900381209082907fdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc905f90a35f818152600960205260409020546113bd906001600160a01b03168261164c565b600660095f8381526020019081526020015f206002016040516113e09190611f36565b9081526040805160209281900383019020805460ff191690555f8381526009909252812080546001600160a01b0319168155600181018290559061142760028301826119e8565b611434600383015f6119e8565b60048201805463ffffffff191681555f61145160058501826119ca565b505060068201805463ffffffff191681555f61147060078501826119ca565b50505050506114b3565b611484600b6110ad565b61148f906001611f23565b604051631a79674760e11b81526004810191909152602481018690526044016108f3565b6040805160c0810182523381526020808201888152828401888152606084018890526080840187905260a084018690525f8a8152600990935293909120825181546001600160a01b0319166001600160a01b0390911617815590516001820155915190919060028201906115279082612082565b506060820151600382019061153c9082612082565b506080820151805160048301805463ffffffff191663ffffffff90921691909117815560208083015180516115779260058701920190611a1f565b50505060a0820151805160068301805463ffffffff191663ffffffff90921691909117815560208083015180516115b49260078701920190611a1f565b50505090505060016006856040516115cc9190612023565b908152604051908190036020018120805492151560ff19909316929092179091556115f8908590612023565b6040519081900381209086907f864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029905f90a36108a960015f8051602061219b83398151915255565b5f61081782611681611685565b6103a76001600160a01b03831682611716565b6103a782826116816117b5565b5f61167a83836116816117e0565b9392505050565b1190565b5f8061168f845490565b9050805f036116a2576116a2603161182d565b5f6116ad858261183e565b5490505f6116be865f19850161183e565b5486549091508690806116d3576116d3612186565b600190038181905f5260205f20015f90559055806116fc5f885f0161183e90919063ffffffff16565b5561170d865f1985015f8489611865565b50949350505050565b804710156117405760405163cf47918160e01b8152476004820152602481018290526044016108f3565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114611789576040519150601f19603f3d011682016040523d82523d5f602084013e61178e565b606091505b50509050806117b05760405163d6bda27560e01b815260040160405180910390fd5b505050565b5f6117be845490565b84546001810186555f86815260209020018490559050610a2f8482858561195a565b5f806117ea855490565b9050805f036117fd576117fd603161182d565b5f611808868261183e565b54905084611816875f61183e565b5561182486835f8888611865565b95945050505050565b634e487b715f52806020526024601cfd5b5f8261185d61185a84611856845f9081526020902090565b0190565b90565b949350505050565b6001600160ff1b038310156108a95760028381026001810191018581101561190e575f611892888461183e565b5490505f6118a0898461183e565b5490506118b1828763ffffffff8816565b806118c557506118c581878763ffffffff16565b15611907575f6118eb85856118de86868b63ffffffff16565b1515918118919091021890565b90506118f88a89836119a2565b6119058a8a838a8a611865565b505b5050611951565b85821015611951575f611921888461183e565b549050611932818663ffffffff8716565b1561194f576119428887856119a2565b61194f8888858888611865565b505b50505050505050565b8215610a2f5760025f198401045f611972868361183e565b549050611983818563ffffffff8616565b1561198f575050610a2f565b61199a8686846119a2565b50925061195a565b5f6119ad848461183e565b90505f6119ba858461183e565b8054835490915590915550505050565b5080545f8255905f5260205f20908101906119e59190611a82565b50565b5080546119f490611ec3565b5f825580601f10611a03575050565b601f0160209004905f5260205f20908101906119e59190611a82565b828054828255905f5260205f20908101928215611a72579160200282015b82811115611a7257825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611a3d565b50611a7e929150611a82565b5090565b5b80821115611a7e575f8155600101611a83565b803563ffffffff81168114611aa9575f80fd5b919050565b5f60208284031215611abe575f80fd5b61167a82611a96565b5f60208284031215611ad7575f80fd5b5035919050565b5f5b83811015611af8578181015183820152602001611ae0565b50505f910152565b5f8151808452611b17816020860160208601611ade565b601f01601f19169290920160200192915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611b865784516001600160a01b03168252938301936001929092019190830190611b5d565b509695505050505050565b60018060a01b038716815285602082015260c060408201525f611bb760c0830187611b00565b8281036060840152611bc98187611b00565b90508281036080840152611bdd8186611b2b565b905082810360a0840152611bf18185611b2b565b9998505050505050505050565b5f805f805f60a08688031215611c12575f80fd5b853561ffff81168114611c23575f80fd5b945060208601356001600160401b0381168114611c3e575f80fd5b94979496505050506040830135926060810135926080909101359150565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611c9257611c92611c5c565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611cc057611cc0611c5c565b604052919050565b5f82601f830112611cd7575f80fd5b81356001600160401b03811115611cf057611cf0611c5c565b611d03601f8201601f1916602001611c98565b818152846020838601011115611d17575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60408284031215611d43575f80fd5b611d4b611c70565b9050611d5682611a96565b81526020808301356001600160401b0380821115611d72575f80fd5b818501915085601f830112611d85575f80fd5b813581811115611d9757611d97611c5c565b8060051b9150611da8848301611c98565b8181529183018401918481019088841115611dc1575f80fd5b938501935b83851015611df657843592506001600160a01b0383168314611de6575f80fd5b8282529385019390850190611dc6565b808688015250505050505092915050565b5f805f8060808587031215611e1a575f80fd5b84356001600160401b0380821115611e30575f80fd5b611e3c88838901611cc8565b95506020870135915080821115611e51575f80fd5b611e5d88838901611cc8565b94506040870135915080821115611e72575f80fd5b611e7e88838901611d33565b93506060870135915080821115611e93575f80fd5b50611ea087828801611d33565b91505092959194509250565b5f60208284031215611ebc575f80fd5b5051919050565b600181811c90821680611ed757607f821691505b602082108103611ef557634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b8082018082111561081757610817611f0f565b5f808354611f4381611ec3565b60018281168015611f5b5760018114611f7057611f9c565b60ff1984168752821515830287019450611f9c565b875f526020805f205f5b85811015611f935781548a820152908401908201611f7a565b50505082870194505b50929695505050505050565b8181038181111561081757610817611f0f565b60a081525f611fcd60a0830188611b00565b8281036020840152611fdf8188611b00565b90508281036040840152611ff38187611b2b565b905082810360608401526120078186611b2b565b9150506001600160401b03831660808301529695505050505050565b5f8251612034818460208701611ade565b9190910192915050565b601f8211156117b057805f5260205f20601f840160051c810160208510156120635750805b601f840160051c820191505b818110156108a9575f815560010161206f565b81516001600160401b0381111561209b5761209b611c5c565b6120af816120a98454611ec3565b8461203e565b602080601f8311600181146120e2575f84156120cb5750858301515b5f19600386901b1c1916600185901b178555612139565b5f85815260208120601f198616915b82811015612110578886015182559484019460019091019084016120f1565b508582101561212d57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b634e487b7160e01b5f52603260045260245ffd5b602081525f61167a6020830184611b00565b5f60208284031215612177575f80fd5b8151801515811461167a575f80fd5b634e487b7160e01b5f52603160045260245ffdfe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220d9dbae405a685a505bae7092ccaa4048aba3643a3121d28cbd3fe05ac300dcce64736f6c63430008190033",
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

// TOKENCONTRACT is a free data retrieval call binding the contract method 0x36339388.
//
// Solidity: function TOKEN_CONTRACT() view returns(address)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCaller) TOKENCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSlotAuctionManager.contract.Call(opts, &out, "TOKEN_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENCONTRACT is a free data retrieval call binding the contract method 0x36339388.
//
// Solidity: function TOKEN_CONTRACT() view returns(address)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerSession) TOKENCONTRACT() (common.Address, error) {
	return _NativeTokenSlotAuctionManager.Contract.TOKENCONTRACT(&_NativeTokenSlotAuctionManager.CallOpts)
}

// TOKENCONTRACT is a free data retrieval call binding the contract method 0x36339388.
//
// Solidity: function TOKEN_CONTRACT() view returns(address)
func (_NativeTokenSlotAuctionManager *NativeTokenSlotAuctionManagerCallerSession) TOKENCONTRACT() (common.Address, error) {
	return _NativeTokenSlotAuctionManager.Contract.TOKENCONTRACT(&_NativeTokenSlotAuctionManager.CallOpts)
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
