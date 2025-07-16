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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20Address\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"}],\"name\":\"AuctionEndTimeNotReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionFinalizing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionNotInProgress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minumumBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBid\",\"type\":\"uint256\"}],\"name\":\"BidSmallerThanMinimum\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"DuplicateBidInContention\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"smallestAcceptableBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBid\",\"type\":\"uint256\"}],\"name\":\"InsufficientBidToWinAuction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeIsValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeIsWinning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validationTimeLimit\",\"type\":\"uint256\"}],\"name\":\"ValidatorTimeLimitNotPassed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"BidEvicted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"InitiatedAuctionValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"validatorSlots\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minValidatorDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"}],\"name\":\"NewValidatorAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"SuccessfulBidPlaced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC20\",\"outputs\":[{\"internalType\":\"contractIERC20Mintable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinBidRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinValidatorDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinimumBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionState\",\"outputs\":[{\"internalType\":\"enumAuctionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"bidderInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"validatorslots\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minAuctionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValidatorDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBid\",\"type\":\"uint256\"}],\"name\":\"initiateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSlots\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506040516123bb3803806123bb833981016040819052602b91607c565b600b80546001600160a01b0319166001600160a01b039283161790555f80546001600160a81b0319169290911691909117905560a8565b80516001600160a01b03811681146077575f80fd5b919050565b5f8060408385031215608c575f80fd5b6093836062565b9150609f602084016062565b90509250929050565b612306806100b55f395ff3fe608060405234801561000f575f80fd5b5060043610610110575f3560e01c80639681d9401161009e578063b6e6a2ca1161006e578063b6e6a2ca14610263578063cc4aa20414610276578063da4312a414610287578063f556f60a1461029a578063fe67a54b146102a3575f80fd5b80639681d94014610200578063a3a65e4814610213578063a476f67514610226578063ae9483e014610239575f80fd5b80634b449cba116100e45780634b449cba146101a55780636fe4d1e1146101bc5780637a7df932146101c55780637fb45099146101d8578063913ed0b1146101f8575f80fd5b8062d841f81461011457806305af82561461014a5780630d5daf3b1461015f5780631077830a14610184575b5f80fd5b60025461012d906201000090046001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b61015d610158366004611bd4565b6102ab565b005b61017261016d366004611bed565b61031f565b60405161014196959493929190611cb7565b6002546101929061ffff1681565b60405161ffff9091168152602001610141565b6101ae60015481565b604051908152602001610141565b6101ae60035481565b61015d6101d3366004611d24565b610554565b5f546101eb90600160a01b900460ff1681565b6040516101419190611d96565b6101ae6106a3565b6101ae61020e366004611bd4565b610717565b6101ae610221366004611bd4565b61078f565b61015d610234366004611bed565b6107c5565b5f5461024b906001600160a01b031681565b6040516001600160a01b039091168152602001610141565b61015d610271366004611bed565b610821565b600b546001600160a01b031661024b565b61015d610295366004611f67565b610972565b6101ae60045481565b61015d6109b0565b5f5460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af11580156102f7573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061031b9190612015565b5050565b60086020525f90815260409020805460018201546002830180546001600160a01b039093169391926103509061202c565b80601f016020809104026020016040519081016040528092919081815260200182805461037c9061202c565b80156103c75780601f1061039e576101008083540402835291602001916103c7565b820191905f5260205f20905b8154815290600101906020018083116103aa57829003601f168201915b5050505050908060030180546103dc9061202c565b80601f01602080910402602001604051908101604052809291908181526020018280546104089061202c565b80156104535780601f1061042a57610100808354040283529160200191610453565b820191905f5260205f20905b81548152906001019060200180831161043657829003601f168201915b505060408051808201825260048701805463ffffffff1682526005880180548451602082810282018101909652818152989998939750919550838701945091928301828280156104ca57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116104ac575b50505091909252505060408051808201825260068501805463ffffffff1682526007860180548451602082810282018101909652818152969796939550919380860193929083018282801561054657602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610528575b505050505081525050905086565b60015f54600160a01b900460ff16600281111561057357610573611d82565b0361059157604051639bbb1f3b60e01b815260040160405180910390fd5b60025f54600160a01b900460ff1660028111156105b0576105b0611d82565b036105ce57604051631f86de2560e21b815260040160405180910390fd5b5f805460ff60a01b1916600160a01b1790556002805469ffffffffffffffff00001916620100006001600160401b0387160217905561060d8342612078565b6001556002805461ffff191661ffff8716179055600382905560048190555f600981905561063d90600a90611af0565b6002546004546040805161ffff90931683526001600160401b038716602084015282018490526060820185905260808201527f50e40b07b3cd91a269f83830c35b4c090f74ebd0cc06dce4db251ccc23627b549060a00160405180910390a15050505050565b5f60015f54600160a01b900460ff1660028111156106c3576106c3611d82565b146106e0576040516232f00d60e71b815260040160405180910390fd5b60025461ffff166106f0600a5490565b10156106fd575060045490565b610707600a611044565b610712906001612078565b905090565b5f805460405163025a076560e61b815263ffffffff841660048201526001600160a01b0390911690639681d940906024015b6020604051808303815f875af1158015610765573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107899190612015565b92915050565b5f8054604051631474cbc960e31b815263ffffffff841660048201526001600160a01b039091169063a3a65e4890602401610749565b5f54604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015b5f604051808303815f87803b158015610808575f80fd5b505af115801561081a573d5f803e3d5ffd5b5050505050565b5f8181526007602052604090206001015442101561086d575f8181526007602052604090819020600101549051631a936b7b60e11b815260048101919091526024015b60405180910390fd5b600660075f8381526020019081526020015f20600201604051610890919061208b565b90815260405190819003602001902080546001600160a01b03191681555f600182018190556108c26002830182611b0e565b6108cf600383015f611b0e565b505f600482018190556005909101805467ffffffffffffffff1916905581815260076020526040812080546001600160a01b0319168155600181018290559061091b6002830182611b0e565b610928600383015f611b0e565b505f60048281018290556005909201805467ffffffffffffffff1916905554604051635b73516560e11b81529182018390526001600160a01b03169063b6e6a2ca906024016107f1565b61097a611069565b61098785858585856110b3565b61081a60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6109b8611069565b60015f54600160a01b900460ff1660028111156109d7576109d7611d82565b146109f4576040516232f00d60e71b815260040160405180910390fd5b5f805460ff60a01b1916600160a11b1790556001544211610a2e57600154604051638230260760e01b815260040161086491815260200190565b5f600155600954158015610a435750600a5415155b15610a5657610a52600a611044565b6009555b5f610a60600a5490565b1115610ffd575f610a71600a61156a565b5f818152600860209081526040808320815160c08101835281546001600160a01b0316815260018201549381019390935260028101805495965093949293909291840191610abe9061202c565b80601f0160208091040260200160405190810160405280929190818152602001828054610aea9061202c565b8015610b355780601f10610b0c57610100808354040283529160200191610b35565b820191905f5260205f20905b815481529060010190602001808311610b1857829003601f168201915b50505050508152602001600382018054610b4e9061202c565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7a9061202c565b8015610bc55780601f10610b9c57610100808354040283529160200191610bc5565b820191905f5260205f20905b815481529060010190602001808311610ba857829003601f168201915b505050918352505060408051808201825260048401805463ffffffff168252600585018054845160208281028201810190965281815295850195939492938584019390929190830182828015610c4257602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610c24575b50505091909252505050815260408051808201825260068401805463ffffffff168252600785018054845160208281028201810190965281815295850195939492938584019390929190830182828015610cc357602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610ca5575b505050919092525050509052508051600954919250610ceb91610ce690856120fd565b611577565b5f80546040838101516060850151608086015160a08701516002549451634e5bb12760e11b81526001600160a01b0390961695639cb7624e95610d459594939291620100009091046001600160401b031690600401612110565b6020604051808303815f875af1158015610d61573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d859190612015565b9050815f01516001600160a01b0316817f032100a60b84cf31291641706872970585d53fa5be412562041fbf66a99ce285600154600354610dc69190612078565b60025460408051928352620100009091046001600160401b031660208301520160405180910390a36040518060c00160405280835f01516001600160a01b03168152602001600154600354610e1b9190612078565b8152602001836040015181526020018360600151815260200182815260200160028054906101000a90046001600160401b03166001600160401b031681525060068360400151604051610e6e9190612178565b9081526040805160209281900383019020835181546001600160a01b0319166001600160a01b039091161781559183015160018301558201516002820190610eb690826121d7565b5060608201516003820190610ecb90826121d7565b506080820151600482015560a0909101516005909101805467ffffffffffffffff19166001600160401b039092169190911790556040805160c0810190915282516001600160a01b031681526001546003546020830191610f2b91612078565b8152604084810151602080840191909152606080870151838501528301859052600280546201000090046001600160401b03166080909401939093525f85815260078252829020845181546001600160a01b0319166001600160a01b03909116178155908401516001820155908301519091820190610faa90826121d7565b5060608201516003820190610fbf90826121d7565b506080820151600482015560a0909101516005909101805467ffffffffffffffff19166001600160401b039092169190911790555050600955610a56565b6002805461ffff191690555f6003819055805460ff60a01b1916905561104260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b565b5f815f015f8154811061105957611059612296565b905f5260205f2001549050919050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016110ad57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60015f54600160a01b900460ff1660028111156110d2576110d2611d82565b146110ef576040516232f00d60e71b815260040160405180910390fd5b5f5460405163d47a948b60e01b81526001600160a01b039091169063d47a948b9061111e9087906004016122aa565b602060405180830381865afa158015611139573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061115d9190612015565b1561117d578360405163f3c815f760e01b815260040161086491906122aa565b60058460405161118d9190612178565b9081526040519081900360200190205460ff16156111c0578360405163f3c1d91160e01b815260040161086491906122aa565b5f858152600860205260409020546001600160a01b0316156111f857604051630517e2e760e21b815260048101869052602401610864565b84600454111561122657600480546040516301a0013b60e61b81529182015260248101869052604401610864565b60025461ffff16611236600a5490565b1015611256576112458561158e565b50611251600a866115a6565b6113ed565b84611261600a611044565b10156113b4576112708561158e565b505f61127d600a876115b3565b60098190555f818152600860205260409081902090519192506112a59160029091019061208b565b6040519081900381209082907fdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc905f90a35f818152600860205260409020546112f7906001600160a01b031682611577565b600560085f8381526020019081526020015f2060020160405161131a919061208b565b9081526040805160209281900383019020805460ff191690555f8381526008909252812080546001600160a01b031916815560018101829055906113616002830182611b0e565b61136e600383015f611b0e565b60048201805463ffffffff191681555f61138b6005850182611af0565b505060068201805463ffffffff191681555f6113aa6007850182611af0565b50505050506113ed565b6113be600a611044565b6113c9906001612078565b604051631a79674760e11b8152600481019190915260248101869052604401610864565b6040805160c0810182523381526020808201888152828401888152606084018890526080840187905260a084018690525f8a8152600890935293909120825181546001600160a01b0319166001600160a01b03909116178155905160018201559151909190600282019061146190826121d7565b506060820151600382019061147690826121d7565b506080820151805160048301805463ffffffff191663ffffffff90921691909117815560208083015180516114b19260058701920190611b45565b50505060a0820151805160068301805463ffffffff191663ffffffff90921691909117815560208083015180516114ee9260078701920190611b45565b50505090505060016005856040516115069190612178565b908152604051908190036020018120805492151560ff1990931692909217909155611532908590612178565b6040519081900381209086907f864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029905f90a35050505050565b5f610789826115c86115cc565b600b5461031b906001600160a01b0316838361165d565b600b545f90610789906001600160a01b0316836116c1565b61031b82826115c86116cd565b5f6115c183836115c86116fe565b9392505050565b1190565b5f806115d6845490565b9050805f036115e9576115e9603161174b565b5f6115f4858261175c565b5490505f611605865f19850161175c565b54865490915086908061161a5761161a6122bc565b600190038181905f5260205f20015f90559055806116435f885f0161175c90919063ffffffff16565b55611654865f1985015f8489611783565b50949350505050565b6040516001600160a01b038381166024830152604482018390526116bc91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050611878565b505050565b5f6115c18333846118e4565b5f6116d6845490565b84546001810186555f868152602090200184905590506116f884828585611a47565b50505050565b5f80611708855490565b9050805f0361171b5761171b603161174b565b5f611726868261175c565b54905084611734875f61175c565b5561174286835f8888611783565b95945050505050565b634e487b715f52806020526024601cfd5b5f8261177b61177884611774845f9081526020902090565b0190565b90565b949350505050565b6001600160ff1b0383101561081a5760028381026001810191018581101561182c575f6117b0888461175c565b5490505f6117be898461175c565b5490506117cf828763ffffffff8816565b806117e357506117e381878763ffffffff16565b15611825575f61180985856117fc86868b63ffffffff16565b1515918118919091021890565b90506118168a8983611a8f565b6118238a8a838a8a611783565b505b505061186f565b8582101561186f575f61183f888461175c565b549050611850818663ffffffff8716565b1561186d57611860888785611a8f565b61186d8888858888611783565b505b50505050505050565b5f8060205f8451602086015f885af180611897576040513d5f823e3d81fd5b50505f513d915081156118ae5780600114156118bb565b6001600160a01b0384163b155b156116f857604051635274afe760e01b81526001600160a01b0385166004820152602401610864565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa15801561192a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061194e9190612015565b90506119656001600160a01b038616853086611ab7565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa1580156119a9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119cd9190612015565b9050818111611a335760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610864565b611a3d82826120fd565b9695505050505050565b82156116f85760025f198401045f611a5f868361175c565b549050611a70818563ffffffff8616565b15611a7c5750506116f8565b611a87868684611a8f565b509250611a47565b5f611a9a848461175c565b90505f611aa7858461175c565b8054835490915590915550505050565b6040516001600160a01b0384811660248301528381166044830152606482018390526116f89186918216906323b872dd9060840161168a565b5080545f8255905f5260205f2090810190611b0b9190611ba8565b50565b508054611b1a9061202c565b5f825580601f10611b29575050565b601f0160209004905f5260205f2090810190611b0b9190611ba8565b828054828255905f5260205f20908101928215611b98579160200282015b82811115611b9857825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611b63565b50611ba4929150611ba8565b5090565b5b80821115611ba4575f8155600101611ba9565b803563ffffffff81168114611bcf575f80fd5b919050565b5f60208284031215611be4575f80fd5b6115c182611bbc565b5f60208284031215611bfd575f80fd5b5035919050565b5f5b83811015611c1e578181015183820152602001611c06565b50505f910152565b5f8151808452611c3d816020860160208601611c04565b601f01601f19169290920160200192915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611cac5784516001600160a01b03168252938301936001929092019190830190611c83565b509695505050505050565b60018060a01b038716815285602082015260c060408201525f611cdd60c0830187611c26565b8281036060840152611cef8187611c26565b90508281036080840152611d038186611c51565b905082810360a0840152611d178185611c51565b9998505050505050505050565b5f805f805f60a08688031215611d38575f80fd5b853561ffff81168114611d49575f80fd5b945060208601356001600160401b0381168114611d64575f80fd5b94979496505050506040830135926060810135926080909101359150565b634e487b7160e01b5f52602160045260245ffd5b6020810160038310611db657634e487b7160e01b5f52602160045260245ffd5b91905290565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611df257611df2611dbc565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611e2057611e20611dbc565b604052919050565b5f82601f830112611e37575f80fd5b81356001600160401b03811115611e5057611e50611dbc565b611e63601f8201601f1916602001611df8565b818152846020838601011115611e77575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60408284031215611ea3575f80fd5b611eab611dd0565b9050611eb682611bbc565b81526020808301356001600160401b0380821115611ed2575f80fd5b818501915085601f830112611ee5575f80fd5b813581811115611ef757611ef7611dbc565b8060051b9150611f08848301611df8565b8181529183018401918481019088841115611f21575f80fd5b938501935b83851015611f5657843592506001600160a01b0383168314611f46575f80fd5b8282529385019390850190611f26565b808688015250505050505092915050565b5f805f805f60a08688031215611f7b575f80fd5b8535945060208601356001600160401b0380821115611f98575f80fd5b611fa489838a01611e28565b95506040880135915080821115611fb9575f80fd5b611fc589838a01611e28565b94506060880135915080821115611fda575f80fd5b611fe689838a01611e93565b93506080880135915080821115611ffb575f80fd5b5061200888828901611e93565b9150509295509295909350565b5f60208284031215612025575f80fd5b5051919050565b600181811c9082168061204057607f821691505b60208210810361205e57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561078957610789612064565b5f8083546120988161202c565b600182811680156120b057600181146120c5576120f1565b60ff19841687528215158302870194506120f1565b875f526020805f205f5b858110156120e85781548a8201529084019082016120cf565b50505082870194505b50929695505050505050565b8181038181111561078957610789612064565b60a081525f61212260a0830188611c26565b82810360208401526121348188611c26565b905082810360408401526121488187611c51565b9050828103606084015261215c8186611c51565b9150506001600160401b03831660808301529695505050505050565b5f8251612189818460208701611c04565b9190910192915050565b601f8211156116bc57805f5260205f20601f840160051c810160208510156121b85750805b601f840160051c820191505b8181101561081a575f81556001016121c4565b81516001600160401b038111156121f0576121f0611dbc565b612204816121fe845461202c565b84612193565b602080601f831160018114612237575f84156122205750858301515b5f19600386901b1c1916600185901b17855561228e565b5f85815260208120601f198616915b8281101561226557888601518255948401946001909101908401612246565b508582101561228257878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b634e487b7160e01b5f52603260045260245ffd5b602081525f6115c16020830184611c26565b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220c966e1ec41fa2c144e1f144f19a48252d0b1defcac911d4828f0e68516c344ab64736f6c63430008190033",
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

// ERC20 is a free data retrieval call binding the contract method 0xcc4aa204.
//
// Solidity: function ERC20() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) ERC20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "ERC20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC20 is a free data retrieval call binding the contract method 0xcc4aa204.
//
// Solidity: function ERC20() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) ERC20() (common.Address, error) {
	return _ERC20TokenSlotAuctionManager.Contract.ERC20(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// ERC20 is a free data retrieval call binding the contract method 0xcc4aa204.
//
// Solidity: function ERC20() view returns(address)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) ERC20() (common.Address, error) {
	return _ERC20TokenSlotAuctionManager.Contract.ERC20(&_ERC20TokenSlotAuctionManager.CallOpts)
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

// AuctionState is a free data retrieval call binding the contract method 0x7fb45099.
//
// Solidity: function auctionState() view returns(uint8)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCaller) AuctionState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenSlotAuctionManager.contract.Call(opts, &out, "auctionState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AuctionState is a free data retrieval call binding the contract method 0x7fb45099.
//
// Solidity: function auctionState() view returns(uint8)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerSession) AuctionState() (uint8, error) {
	return _ERC20TokenSlotAuctionManager.Contract.AuctionState(&_ERC20TokenSlotAuctionManager.CallOpts)
}

// AuctionState is a free data retrieval call binding the contract method 0x7fb45099.
//
// Solidity: function auctionState() view returns(uint8)
func (_ERC20TokenSlotAuctionManager *ERC20TokenSlotAuctionManagerCallerSession) AuctionState() (uint8, error) {
	return _ERC20TokenSlotAuctionManager.Contract.AuctionState(&_ERC20TokenSlotAuctionManager.CallOpts)
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
