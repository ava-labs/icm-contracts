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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"BidEvicted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endtime\",\"type\":\"uint256\"}],\"name\":\"InitiatedAuctionValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"validatorSlots\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validationTimeLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionEndTime\",\"type\":\"uint256\"}],\"name\":\"NewValidatorAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"SuccessfulBidPlaced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"bidderInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"validatorslots\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"auctionLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validationLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumbid\",\"type\":\"uint256\"}],\"name\":\"initiateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peekTop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validationTimeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSlots\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"validatorsByNodeID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"Weight\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"validatorsByValidationID\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"Weight\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b5060405161259c38038061259c833981016040819052602b916074565b5f80546001600160a01b039384166001600160a01b0319918216179091556001805492909316911617905560a0565b80516001600160a01b0381168114606f575f80fd5b919050565b5f80604083850312156084575f80fd5b608b83605a565b9150609760208401605a565b90509250929050565b6124ef806100ad5f395ff3fe608060405234801561000f575f80fd5b5060043610610126575f3560e01c80639681d940116100a9578063b6e6a2ca1161006e578063b6e6a2ca146102a4578063d3a86386146102b7578063da4312a4146102c0578063df393075146102d3578063fe67a54b146102f7575f80fd5b80639681d940146102455780639c91ddbb14610258578063a3a65e481461026b578063a476f6751461027e578063ae9483e014610291575f80fd5b806336339388116100ef57806336339388146101d25780634b449cba146101fc57806372d49319146102055780637a7df9321461022a5780638c3c78e31461023d575f80fd5b8062d841f81461012a57806305af8256146101605780630d5daf3b146101755780631077830a1461019a5780631931f2c7146101bb575b5f80fd5b600354610143906201000090046001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b61017361016e366004611ce0565b6102ff565b005b610188610183366004611cf9565b610374565b60405161015796959493929190611dc3565b6003546101a89061ffff1681565b60405161ffff9091168152602001610157565b6101c460045481565b604051908152602001610157565b5f546101e4906001600160a01b031681565b6040516001600160a01b039091168152602001610157565b6101c460025481565b610218610213366004611cf9565b6105a9565b60405161015796959493929190611e30565b610173610238366004611e8b565b6106fd565b6101c461082a565b6101c4610253366004611ce0565b610872565b610218610266366004611fc0565b6108ea565b6101c4610279366004611ce0565b610927565b61017361028c366004611cf9565b61095d565b6001546101e4906001600160a01b031681565b6101736102b2366004611cf9565b6109ba565b6101c460055481565b6101736102ce3660046120c5565b610b2d565b6001546102e790600160a01b900460ff1681565b6040519015158152602001610157565b610173611232565b60015460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af115801561034c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103709190612173565b5050565b60096020525f90815260409020805460018201546002830180546001600160a01b039093169391926103a59061218a565b80601f01602080910402602001604051908101604052809291908181526020018280546103d19061218a565b801561041c5780601f106103f35761010080835404028352916020019161041c565b820191905f5260205f20905b8154815290600101906020018083116103ff57829003601f168201915b5050505050908060030180546104319061218a565b80601f016020809104026020016040519081016040528092919081815260200182805461045d9061218a565b80156104a85780601f1061047f576101008083540402835291602001916104a8565b820191905f5260205f20905b81548152906001019060200180831161048b57829003601f168201915b505060408051808201825260048701805463ffffffff16825260058801805484516020828102820181019096528181529899989397509195508387019450919283018282801561051f57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610501575b50505091909252505060408051808201825260068501805463ffffffff1682526007860180548451602082810282018101909652818152969796939550919380860193929083018282801561059b57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161057d575b505050505081525050905086565b60086020525f90815260409020805460018201546002830180546001600160a01b039093169391926105da9061218a565b80601f01602080910402602001604051908101604052809291908181526020018280546106069061218a565b80156106515780601f1061062857610100808354040283529160200191610651565b820191905f5260205f20905b81548152906001019060200180831161063457829003601f168201915b5050505050908060030180546106669061218a565b80601f01602080910402602001604051908101604052809291908181526020018280546106929061218a565b80156106dd5780601f106106b4576101008083540402835291602001916106dd565b820191905f5260205f20905b8154815290600101906020018083116106c057829003601f168201915b5050505060048301546005909301549192916001600160401b0316905086565b600154600160a01b900460ff161561075c5760405162461bcd60e51b815260206004820152601d60248201527f41756374696f6e2063757272656e746c7920696e2070726f677265737300000060448201526064015b60405180910390fd5b6003805469ffffffffffffffff00001916620100006001600160401b0387160217905561078983426121d6565b6002556003805461ffff191661ffff8716179055600482905560058190555f600a8190556107b990600b90611bfc565b6001805460ff60a01b1916600160a01b1790556003546040805161ffff90921682526001600160401b03861660208301528101839052606081018490527fe399236a37805e37555e18792e9216f1d9cae2fae31f4b96ad5c2d200166bf839060800160405180910390a15050505050565b6001545f90600160a01b900460ff166108555760405162461bcd60e51b8152600401610753906121e9565b600b545f0361086357505f90565b61086d600b6118f8565b905090565b60015460405163025a076560e61b815263ffffffff831660048201525f916001600160a01b031690639681d940906024015b6020604051808303815f875af11580156108c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e49190612173565b92915050565b80516020818301810180516007825292820191909301209152805460018201546002830180546001600160a01b039093169391926105da9061218a565b600154604051631474cbc960e31b815263ffffffff831660048201525f916001600160a01b03169063a3a65e48906024016108a4565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015b5f604051808303815f87803b1580156109a1575f80fd5b505af11580156109b3573d5f803e3d5ffd5b5050505050565b5f818152600860205260409020600101544211610a255760405162461bcd60e51b815260206004820152602360248201527f56616c69646174696f6e2074696d65206c696d697420686173206e6f7420656e60448201526219195960ea1b6064820152608401610753565b600760085f8381526020019081526020015f20600201604051610a489190612220565b90815260405190819003602001902080546001600160a01b03191681555f60018201819055610a7a6002830182611c1a565b610a87600383015f611c1a565b505f600482018190556005909101805467ffffffffffffffff1916905581815260086020526040812080546001600160a01b03191681556001810182905590610ad36002830182611c1a565b610ae0600383015f611c1a565b505f6004828101919091556005909101805467ffffffffffffffff19169055600154604051635b73516560e11b81529182018390526001600160a01b03169063b6e6a2ca9060240161098a565b600154600160a01b900460ff16610b565760405162461bcd60e51b8152600401610753906121e9565b5f6001600160a01b0316600785604051610b709190612292565b908152604051908190036020019020546001600160a01b031614610bd65760405162461bcd60e51b815260206004820152601b60248201527f4e6f646520697320616c726561647920612076616c696461746f7200000000006044820152606401610753565b600684604051610be69190612292565b9081526040519081900360200190205460ff1615610c545760405162461bcd60e51b815260206004820152602560248201527f4e6f646520697320616c726561647920696e20612077696e6e696e6720706f7360448201526434ba34b7b760d91b6064820152608401610753565b5f858152600960205260409020546001600160a01b031615610cb85760405162461bcd60e51b815260206004820152601860248201527f4475706c69636174652062696420696e2072756e6e696e6700000000000000006044820152606401610753565b600554851015610d0a5760405162461bcd60e51b815260206004820152601b60248201527f426964206c6f776572207468616e206d696e696d756d206269642000000000006044820152606401610753565b60035461ffff16610d1a600b5490565b1015610def575f546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015610d73573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d9791906122ad565b610ddf5760405162461bcd60e51b8152602060048201526019602482015278125b9cdd59999a58da595b9d08199d5b991cc81d1bc8189a59603a1b6044820152606401610753565b610dea600b8661191d565b6110b5565b600a54158015610e07575084610e05600b6118f8565b115b15610e2957600a85905560405162461bcd60e51b8152610753906004016122cc565b84610e34600b6118f8565b101561109d575f546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015610e8d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610eb191906122ad565b610ef95760405162461bcd60e51b8152602060048201526019602482015278125b9cdd59999a58da595b9d08199d5b991cc81d1bc8189a59603a1b6044820152606401610753565b5f610f05600b8761192a565b600a8190555f81815260096020526040908190209051919250610f2d91600290910190612220565b6040519081900381209082907fdbb7ff16cd6e3cc5c7aca1a615bf0d56f746d9f5708e12d52ebe243518b536cc905f90a35f8054828252600960205260409182902054915163a9059cbb60e01b81526001600160a01b0392831660048201526024810184905291169063a9059cbb906044016020604051808303815f875af1158015610fbb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fdf91906122ad565b50600660095f8381526020019081526020015f206002016040516110039190612220565b9081526040805160209281900383019020805460ff191690555f8381526009909252812080546001600160a01b0319168155600181018290559061104a6002830182611c1a565b611057600383015f611c1a565b60048201805463ffffffff191681555f6110746005850182611bfc565b505060068201805463ffffffff191681555f6110936007850182611bfc565b50505050506110b5565b60405162461bcd60e51b8152600401610753906122cc565b6040805160c0810182523381526020808201888152828401888152606084018890526080840187905260a084018690525f8a8152600990935293909120825181546001600160a01b0319166001600160a01b0390911617815590516001820155915190919060028201906111299082612357565b506060820151600382019061113e9082612357565b506080820151805160048301805463ffffffff191663ffffffff90921691909117815560208083015180516111799260058701920190611c51565b50505060a0820151805160068301805463ffffffff191663ffffffff90921691909117815560208083015180516111b69260078701920190611c51565b50505090505060016006856040516111ce9190612292565b908152604051908190036020018120805492151560ff19909316929092179091556111fa908590612292565b6040519081900381209086907f864a2110bb41c315e04bd6122b668384bca4b0d8509e55850db8347e86a95029905f90a35050505050565b600154600160a01b900460ff1661125b5760405162461bcd60e51b8152600401610753906121e9565b60025442116112ac5760405162461bcd60e51b815260206004820152601b60248201527f41756374696f6e20656e6474696d65206e6f74207265616368656400000000006044820152606401610753565b6001805460ff60a01b191690555f600255600a541580156112ce5750600b5415155b156112e1576112dd600b6118f8565b600a555b5f6112eb600b5490565b11156118e7575f6112fc600b61193f565b5f818152600960209081526040808320815160c08101835281546001600160a01b03168152600182015493810193909352600281018054959650939492939092918401916113499061218a565b80601f01602080910402602001604051908101604052809291908181526020018280546113759061218a565b80156113c05780601f10611397576101008083540402835291602001916113c0565b820191905f5260205f20905b8154815290600101906020018083116113a357829003601f168201915b505050505081526020016003820180546113d99061218a565b80601f01602080910402602001604051908101604052809291908181526020018280546114059061218a565b80156114505780601f1061142757610100808354040283529160200191611450565b820191905f5260205f20905b81548152906001019060200180831161143357829003601f168201915b505050918352505060408051808201825260048401805463ffffffff1682526005850180548451602082810282018101909652818152958501959394929385840193909291908301828280156114cd57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116114af575b50505091909252505050815260408051808201825260068401805463ffffffff16825260078501805484516020828102820181019096528181529585019593949293858401939092919083018282801561154e57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611530575b505050919092525050509052505f548151600a549293506001600160a01b039091169163a9059cbb91906115829086612416565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156115ca573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115ee91906122ad565b506001546040828101516060840151608085015160a08601516003549451634e5bb12760e11b81525f966001600160a01b031695639cb7624e9561164c9590949093909290916201000090046001600160401b031690600401612429565b6020604051808303815f875af1158015611668573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061168c9190612173565b9050815f01516001600160a01b0316817fccb9a8514540a67c432230e6261f90893e22e8e31f65c40e7629b5a47374dbbe6002546004546116cd91906121d6565b60405190815260200160405180910390a36040518060c00160405280835f01516001600160a01b0316815260200160025460045461170b91906121d6565b81526040848101805160208401526060808701518385015283018590526003546201000090046001600160401b03166080909301929092529051905160079161175391612292565b9081526040805160209281900383019020835181546001600160a01b0319166001600160a01b03909116178155918301516001830155820151600282019061179b9082612357565b50606082015160038201906117b09082612357565b506080820151816004015560a0820151816005015f6101000a8154816001600160401b0302191690836001600160401b031602179055509050506040518060c00160405280835f01516001600160a01b0316815260200160025460045461181791906121d6565b81526040848101516020808401919091526060808701518385015283018590526003546201000090046001600160401b03166080909301929092525f84815260088352819020835181546001600160a01b0319166001600160a01b0390911617815591830151600183015582015160028201906118949082612357565b50606082015160038201906118a99082612357565b506080820151600482015560a0909101516005909101805467ffffffffffffffff19166001600160401b039092169190911790555050600a556112e1565b6003805461ffff191690555f600455565b5f815f015f8154811061190d5761190d612491565b905f5260205f2001549050919050565b610370828261194c611950565b5f611938838361194c611981565b9392505050565b5f6108e48261194c6119ce565b1190565b5f611959845490565b84546001810186555f8681526020902001849055905061197b84828585611a5f565b50505050565b5f8061198b855490565b9050805f0361199e5761199e6031611aa7565b5f6119a98682611ab8565b549050846119b7875f611ab8565b556119c586835f8888611adf565b95945050505050565b5f806119d8845490565b9050805f036119eb576119eb6031611aa7565b5f6119f68582611ab8565b5490505f611a07865f198501611ab8565b548654909150869080611a1c57611a1c6124a5565b600190038181905f5260205f20015f9055905580611a455f885f01611ab890919063ffffffff16565b55611a56865f1985015f8489611adf565b50949350505050565b821561197b5760025f198401045f611a778683611ab8565b549050611a88818563ffffffff8616565b15611a9457505061197b565b611a9f868684611bd4565b509250611a5f565b634e487b715f52806020526024601cfd5b5f82611ad7611ad484611ad0845f9081526020902090565b0190565b90565b949350505050565b6001600160ff1b038310156109b357600283810260018101910185811015611b88575f611b0c8884611ab8565b5490505f611b1a8984611ab8565b549050611b2b828763ffffffff8816565b80611b3f5750611b3f81878763ffffffff16565b15611b81575f611b658585611b5886868b63ffffffff16565b1515918118919091021890565b9050611b728a8983611bd4565b611b7f8a8a838a8a611adf565b505b5050611bcb565b85821015611bcb575f611b9b8884611ab8565b549050611bac818663ffffffff8716565b15611bc957611bbc888785611bd4565b611bc98888858888611adf565b505b50505050505050565b5f611bdf8484611ab8565b90505f611bec8584611ab8565b8054835490915590915550505050565b5080545f8255905f5260205f2090810190611c179190611cb4565b50565b508054611c269061218a565b5f825580601f10611c35575050565b601f0160209004905f5260205f2090810190611c179190611cb4565b828054828255905f5260205f20908101928215611ca4579160200282015b82811115611ca457825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611c6f565b50611cb0929150611cb4565b5090565b5b80821115611cb0575f8155600101611cb5565b803563ffffffff81168114611cdb575f80fd5b919050565b5f60208284031215611cf0575f80fd5b61193882611cc8565b5f60208284031215611d09575f80fd5b5035919050565b5f5b83811015611d2a578181015183820152602001611d12565b50505f910152565b5f8151808452611d49816020860160208601611d10565b601f01601f19169290920160200192915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611db85784516001600160a01b03168252938301936001929092019190830190611d8f565b509695505050505050565b60018060a01b038716815285602082015260c060408201525f611de960c0830187611d32565b8281036060840152611dfb8187611d32565b90508281036080840152611e0f8186611d5d565b905082810360a0840152611e238185611d5d565b9998505050505050505050565b60018060a01b038716815285602082015260c060408201525f611e5660c0830187611d32565b8281036060840152611e688187611d32565b9150508360808301526001600160401b03831660a0830152979650505050505050565b5f805f805f60a08688031215611e9f575f80fd5b853561ffff81168114611eb0575f80fd5b945060208601356001600160401b0381168114611ecb575f80fd5b94979496505050506040830135926060810135926080909101359150565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611f1f57611f1f611ee9565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611f4d57611f4d611ee9565b604052919050565b5f82601f830112611f64575f80fd5b81356001600160401b03811115611f7d57611f7d611ee9565b611f90601f8201601f1916602001611f25565b818152846020838601011115611fa4575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215611fd0575f80fd5b81356001600160401b03811115611fe5575f80fd5b611ad784828501611f55565b5f60408284031215612001575f80fd5b612009611efd565b905061201482611cc8565b81526020808301356001600160401b0380821115612030575f80fd5b818501915085601f830112612043575f80fd5b81358181111561205557612055611ee9565b8060051b9150612066848301611f25565b818152918301840191848101908884111561207f575f80fd5b938501935b838510156120b457843592506001600160a01b03831683146120a4575f80fd5b8282529385019390850190612084565b808688015250505050505092915050565b5f805f805f60a086880312156120d9575f80fd5b8535945060208601356001600160401b03808211156120f6575f80fd5b61210289838a01611f55565b95506040880135915080821115612117575f80fd5b61212389838a01611f55565b94506060880135915080821115612138575f80fd5b61214489838a01611ff1565b93506080880135915080821115612159575f80fd5b5061216688828901611ff1565b9150509295509295909350565b5f60208284031215612183575f80fd5b5051919050565b600181811c9082168061219e57607f821691505b6020821081036121bc57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156108e4576108e46121c2565b60208082526017908201527f41756374696f6e206e6f7420696e2070726f6772657373000000000000000000604082015260600190565b5f80835461222d8161218a565b60018281168015612245576001811461225a57612286565b60ff1984168752821515830287019450612286565b875f526020805f205f5b8581101561227d5781548a820152908401908201612264565b50505082870194505b50929695505050505050565b5f82516122a3818460208701611d10565b9190910192915050565b5f602082840312156122bd575f80fd5b81518015158114611938575f80fd5b60208082526022908201527f426964206e6f74206869676820656e6f75676820746f2077696e20617563746960408201526137b760f11b606082015260800190565b601f82111561235257805f5260205f20601f840160051c810160208510156123335750805b601f840160051c820191505b818110156109b3575f815560010161233f565b505050565b81516001600160401b0381111561237057612370611ee9565b6123848161237e845461218a565b8461230e565b602080601f8311600181146123b7575f84156123a05750858301515b5f19600386901b1c1916600185901b17855561240e565b5f85815260208120601f198616915b828110156123e5578886015182559484019460019091019084016123c6565b508582101561240257878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b818103818111156108e4576108e46121c2565b60a081525f61243b60a0830188611d32565b828103602084015261244d8188611d32565b905082810360408401526124618187611d5d565b905082810360608401526124758186611d5d565b9150506001600160401b03831660808301529695505050505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220f0472b40bbec8baa4cfcab040a159e49867f5547d670a34a02ed879d7ad19ac964736f6c63430008190033",
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

// MinimumBid is a free data retrieval call binding the contract method 0xd3a86386.
//
// Solidity: function minimumBid() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCaller) MinimumBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "minimumBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumBid is a free data retrieval call binding the contract method 0xd3a86386.
//
// Solidity: function minimumBid() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerSession) MinimumBid() (*big.Int, error) {
	return _SlotAuctionManager.Contract.MinimumBid(&_SlotAuctionManager.CallOpts)
}

// MinimumBid is a free data retrieval call binding the contract method 0xd3a86386.
//
// Solidity: function minimumBid() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) MinimumBid() (*big.Int, error) {
	return _SlotAuctionManager.Contract.MinimumBid(&_SlotAuctionManager.CallOpts)
}

// PeekTop is a free data retrieval call binding the contract method 0x8c3c78e3.
//
// Solidity: function peekTop() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCaller) PeekTop(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "peekTop")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeekTop is a free data retrieval call binding the contract method 0x8c3c78e3.
//
// Solidity: function peekTop() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerSession) PeekTop() (*big.Int, error) {
	return _SlotAuctionManager.Contract.PeekTop(&_SlotAuctionManager.CallOpts)
}

// PeekTop is a free data retrieval call binding the contract method 0x8c3c78e3.
//
// Solidity: function peekTop() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) PeekTop() (*big.Int, error) {
	return _SlotAuctionManager.Contract.PeekTop(&_SlotAuctionManager.CallOpts)
}

// ValidationTimeLimit is a free data retrieval call binding the contract method 0x1931f2c7.
//
// Solidity: function validationTimeLimit() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCaller) ValidationTimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "validationTimeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidationTimeLimit is a free data retrieval call binding the contract method 0x1931f2c7.
//
// Solidity: function validationTimeLimit() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerSession) ValidationTimeLimit() (*big.Int, error) {
	return _SlotAuctionManager.Contract.ValidationTimeLimit(&_SlotAuctionManager.CallOpts)
}

// ValidationTimeLimit is a free data retrieval call binding the contract method 0x1931f2c7.
//
// Solidity: function validationTimeLimit() view returns(uint256)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) ValidationTimeLimit() (*big.Int, error) {
	return _SlotAuctionManager.Contract.ValidationTimeLimit(&_SlotAuctionManager.CallOpts)
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
// Solidity: function validatorsByNodeID(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 Weight)
func (_SlotAuctionManager *SlotAuctionManagerCaller) ValidatorsByNodeID(opts *bind.CallOpts, nodeID []byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID [32]byte
	Weight       uint64
}, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "validatorsByNodeID", nodeID)

	outstruct := new(struct {
		Addr         common.Address
		EndTime      *big.Int
		NodeID       []byte
		BlsPublicKey []byte
		ValidationID [32]byte
		Weight       uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NodeID = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.BlsPublicKey = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.ValidationID = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.Weight = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// ValidatorsByNodeID is a free data retrieval call binding the contract method 0x9c91ddbb.
//
// Solidity: function validatorsByNodeID(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 Weight)
func (_SlotAuctionManager *SlotAuctionManagerSession) ValidatorsByNodeID(nodeID []byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID [32]byte
	Weight       uint64
}, error) {
	return _SlotAuctionManager.Contract.ValidatorsByNodeID(&_SlotAuctionManager.CallOpts, nodeID)
}

// ValidatorsByNodeID is a free data retrieval call binding the contract method 0x9c91ddbb.
//
// Solidity: function validatorsByNodeID(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 Weight)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) ValidatorsByNodeID(nodeID []byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID [32]byte
	Weight       uint64
}, error) {
	return _SlotAuctionManager.Contract.ValidatorsByNodeID(&_SlotAuctionManager.CallOpts, nodeID)
}

// ValidatorsByValidationID is a free data retrieval call binding the contract method 0x72d49319.
//
// Solidity: function validatorsByValidationID(bytes32 validationID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 Weight)
func (_SlotAuctionManager *SlotAuctionManagerCaller) ValidatorsByValidationID(opts *bind.CallOpts, validationID [32]byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID [32]byte
	Weight       uint64
}, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "validatorsByValidationID", validationID)

	outstruct := new(struct {
		Addr         common.Address
		EndTime      *big.Int
		NodeID       []byte
		BlsPublicKey []byte
		ValidationID [32]byte
		Weight       uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NodeID = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.BlsPublicKey = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.ValidationID = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.Weight = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// ValidatorsByValidationID is a free data retrieval call binding the contract method 0x72d49319.
//
// Solidity: function validatorsByValidationID(bytes32 validationID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 Weight)
func (_SlotAuctionManager *SlotAuctionManagerSession) ValidatorsByValidationID(validationID [32]byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID [32]byte
	Weight       uint64
}, error) {
	return _SlotAuctionManager.Contract.ValidatorsByValidationID(&_SlotAuctionManager.CallOpts, validationID)
}

// ValidatorsByValidationID is a free data retrieval call binding the contract method 0x72d49319.
//
// Solidity: function validatorsByValidationID(bytes32 validationID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 Weight)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) ValidatorsByValidationID(validationID [32]byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID [32]byte
	Weight       uint64
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
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 auctionLength, uint256 validationLength, uint256 minimumbid) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) InitiateAuction(opts *bind.TransactOpts, validatorslots uint16, weight uint64, auctionLength *big.Int, validationLength *big.Int, minimumbid *big.Int) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "initiateAuction", validatorslots, weight, auctionLength, validationLength, minimumbid)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x7a7df932.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 auctionLength, uint256 validationLength, uint256 minimumbid) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) InitiateAuction(validatorslots uint16, weight uint64, auctionLength *big.Int, validationLength *big.Int, minimumbid *big.Int) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateAuction(&_SlotAuctionManager.TransactOpts, validatorslots, weight, auctionLength, validationLength, minimumbid)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x7a7df932.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 auctionLength, uint256 validationLength, uint256 minimumbid) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) InitiateAuction(validatorslots uint16, weight uint64, auctionLength *big.Int, validationLength *big.Int, minimumbid *big.Int) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateAuction(&_SlotAuctionManager.TransactOpts, validatorslots, weight, auctionLength, validationLength, minimumbid)
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
	ValidationID [32]byte
	OwnerAddress common.Address
	Endtime      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitiatedAuctionValidatorRegistration is a free log retrieval operation binding the contract event 0xccb9a8514540a67c432230e6261f90893e22e8e31f65c40e7629b5a47374dbbe.
//
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 endtime)
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
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 endtime)
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
// Solidity: event InitiatedAuctionValidatorRegistration(bytes32 indexed validationID, address indexed ownerAddress, uint256 endtime)
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
	ValidatorSlots      uint16
	ValidatorWeight     uint64
	ValidationTimeLimit *big.Int
	AuctionEndTime      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewValidatorAuction is a free log retrieval operation binding the contract event 0xe399236a37805e37555e18792e9216f1d9cae2fae31f4b96ad5c2d200166bf83.
//
// Solidity: event NewValidatorAuction(uint16 validatorSlots, uint64 validatorWeight, uint256 validationTimeLimit, uint256 auctionEndTime)
func (_SlotAuctionManager *SlotAuctionManagerFilterer) FilterNewValidatorAuction(opts *bind.FilterOpts) (*SlotAuctionManagerNewValidatorAuctionIterator, error) {

	logs, sub, err := _SlotAuctionManager.contract.FilterLogs(opts, "NewValidatorAuction")
	if err != nil {
		return nil, err
	}
	return &SlotAuctionManagerNewValidatorAuctionIterator{contract: _SlotAuctionManager.contract, event: "NewValidatorAuction", logs: logs, sub: sub}, nil
}

// WatchNewValidatorAuction is a free log subscription operation binding the contract event 0xe399236a37805e37555e18792e9216f1d9cae2fae31f4b96ad5c2d200166bf83.
//
// Solidity: event NewValidatorAuction(uint16 validatorSlots, uint64 validatorWeight, uint256 validationTimeLimit, uint256 auctionEndTime)
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

// ParseNewValidatorAuction is a log parse operation binding the contract event 0xe399236a37805e37555e18792e9216f1d9cae2fae31f4b96ad5c2d200166bf83.
//
// Solidity: event NewValidatorAuction(uint16 validatorSlots, uint64 validatorWeight, uint256 validationTimeLimit, uint256 auctionEndTime)
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
