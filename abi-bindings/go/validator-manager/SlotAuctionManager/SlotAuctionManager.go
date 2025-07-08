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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"bidderInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"currentValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"Weight\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"validatorslots\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"auctionLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validationLength\",\"type\":\"uint256\"}],\"name\":\"initiateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peekTop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validationTimeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSlots\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50604051612078380380612078833981016040819052602b916074565b5f80546001600160a01b039384166001600160a01b0319918216179091556001805492909316911617905560a0565b80516001600160a01b0381168114606f575f80fd5b919050565b5f80604083850312156084575f80fd5b608b83605a565b9150609760208401605a565b90509250929050565b611fcb806100ad5f395ff3fe608060405234801561000f575f80fd5b5060043610610110575f3560e01c80638c3c78e31161009e578063ae9483e01161006e578063ae9483e014610268578063b6e6a2ca14610255578063da4312a41461027b578063df3930751461028e578063fe67a54b146102b2575f80fd5b80638c3c78e3146102275780639681d9401461022f578063a3a65e4814610242578063a476f67514610255575f80fd5b80631931f2c7116100e45780631931f2c7146101a557806336339388146101bc5780633e89e29a146101e65780634b449cba146101f95780635fcf253814610202575f80fd5b8062d841f81461011457806305af82561461014a5780630d5daf3b1461015f5780631077830a14610184575b5f80fd5b60035461012d906201000090046001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b61015d61015836600461183f565b6102ba565b005b61017261016d366004611858565b61032f565b60405161014196959493929190611922565b6003546101929061ffff1681565b60405161ffff9091168152602001610141565b6101ae60045481565b604051908152602001610141565b5f546101ce906001600160a01b031681565b6040516001600160a01b039091168152602001610141565b61015d6101f436600461198f565b610564565b6101ae60025481565b610215610210366004611aba565b610634565b60405161014196959493929190611aeb565b6101ae610794565b6101ae61023d36600461183f565b6107cd565b6101ae61025036600461183f565b610845565b61015d610263366004611858565b61087b565b6001546101ce906001600160a01b031681565b61015d610289366004611c1a565b6108d7565b6001546102a290600160a01b900460ff1681565b6040519015158152602001610141565b61015d610e9a565b60015460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af1158015610307573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061032b9190611cc8565b5050565b60076020525f90815260409020805460018201546002830180546001600160a01b0390931693919261036090611cdf565b80601f016020809104026020016040519081016040528092919081815260200182805461038c90611cdf565b80156103d75780601f106103ae576101008083540402835291602001916103d7565b820191905f5260205f20905b8154815290600101906020018083116103ba57829003601f168201915b5050505050908060030180546103ec90611cdf565b80601f016020809104026020016040519081016040528092919081815260200182805461041890611cdf565b80156104635780601f1061043a57610100808354040283529160200191610463565b820191905f5260205f20905b81548152906001019060200180831161044657829003601f168201915b505060408051808201825260048701805463ffffffff1682526005880180548451602082810282018101909652818152989998939750919550838701945091928301828280156104da57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116104bc575b50505091909252505060408051808201825260068501805463ffffffff1682526007860180548451602082810282018101909652818152969796939550919380860193929083018282801561055657602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610538575b505050505081525050905086565b600154600160a01b900460ff16156105c35760405162461bcd60e51b815260206004820152601760248201527f41756374696f6e20616c72656164792072756e6e696e6700000000000000000060448201526064015b60405180910390fd5b6003805469ffffffffffffffff00001916620100006001600160401b038616021790556105f08242611d2b565b6002556003805461ffff191661ffff861617905560048190555f600881905561061b9060099061175b565b50506001805460ff60a01b1916600160a01b1790555050565b80516020818301810180516006825292820191909301209152805460018201546002830180546001600160a01b0390931693919261067190611cdf565b80601f016020809104026020016040519081016040528092919081815260200182805461069d90611cdf565b80156106e85780601f106106bf576101008083540402835291602001916106e8565b820191905f5260205f20905b8154815290600101906020018083116106cb57829003601f168201915b5050505050908060030180546106fd90611cdf565b80601f016020809104026020016040519081016040528092919081815260200182805461072990611cdf565b80156107745780601f1061074b57610100808354040283529160200191610774565b820191905f5260205f20905b81548152906001019060200180831161075757829003601f168201915b5050505060048301546005909301549192916001600160401b0316905086565b5f61079e60095490565b15806107b45750600154600160a01b900460ff16155b156107be57505f90565b6107c860096113d0565b905090565b60015460405163025a076560e61b815263ffffffff831660048201525f916001600160a01b031690639681d940906024015b6020604051808303815f875af115801561081b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061083f9190611cc8565b92915050565b600154604051631474cbc960e31b815263ffffffff831660048201525f916001600160a01b03169063a3a65e48906024016107ff565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b1580156108be575f80fd5b505af11580156108d0573d5f803e3d5ffd5b5050505050565b600154600160a01b900460ff166109295760405162461bcd60e51b815260206004820152601660248201527541756374696f6e206973206e6f742072756e6e696e6760501b60448201526064016105ba565b5f6001600160a01b03166006856040516109439190611d3e565b908152604051908190036020019020546001600160a01b0316146109a95760405162461bcd60e51b815260206004820152601b60248201527f4e6f646520697320616c726561647920612076616c696461746f72000000000060448201526064016105ba565b6005846040516109b99190611d3e565b9081526040519081900360200190205460ff1615610a275760405162461bcd60e51b815260206004820152602560248201527f4e6f646520697320616c726561647920696e20612077696e6e696e6720706f7360448201526434ba34b7b760d91b60648201526084016105ba565b60035461ffff16610a3760095490565b1015610b0c575f546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015610a90573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ab49190611d59565b610afc5760405162461bcd60e51b8152602060048201526019602482015278125b9cdd59999a58da595b9d08199d5b991cc81d1bc8189a59603a1b60448201526064016105ba565b610b076009866113f5565b610d5a565b600854158015610b24575084610b2260096113d0565b115b15610b33576008859055610d5a565b84610b3e60096113d0565b1015610d55575f546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015610b97573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bbb9190611d59565b610c035760405162461bcd60e51b8152602060048201526019602482015278125b9cdd59999a58da595b9d08199d5b991cc81d1bc8189a59603a1b60448201526064016105ba565b5f610c0f600987611402565b60088190555f8054828252600760205260409182902054915163a9059cbb60e01b81526001600160a01b03928316600482015260248101849052929350169063a9059cbb906044016020604051808303815f875af1158015610c73573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c979190611d59565b50600560075f8381526020019081526020015f20600201604051610cbb9190611d78565b9081526040805160209281900383019020805460ff191690555f8381526007909252812080546001600160a01b03191681556001810182905590610d026002830182611779565b610d0f600383015f611779565b60048201805463ffffffff191681555f610d2c600585018261175b565b505060068201805463ffffffff191681555f610d4b600785018261175b565b5050505050610d5a565b6108d0565b6040805160c0810182523381526020808201888152828401888152606084018890526080840187905260a084018690525f8a8152600790935293909120825181546001600160a01b0319166001600160a01b039091161781559051600182015591519091906002820190610dce9082611e33565b5060608201516003820190610de39082611e33565b506080820151805160048301805463ffffffff191663ffffffff9092169190911781556020808301518051610e1e92600587019201906117b0565b50505060a0820151805160068301805463ffffffff191663ffffffff9092169190911781556020808301518051610e5b92600787019201906117b0565b5050509050506001600585604051610e739190611d3e565b908152604051908190036020019020805491151560ff199092169190911790555050505050565b600154600160a01b900460ff16610ef35760405162461bcd60e51b815260206004820152601760248201527f41756374696f6e206e6f7420696e2070726f677265737300000000000000000060448201526064016105ba565b6002544211610f445760405162461bcd60e51b815260206004820152601b60248201527f41756374696f6e20656e6474696d65206e6f742072656163686564000000000060448201526064016105ba565b6001805460ff60a01b191690555f600255600854158015610f66575060095415155b15610f7957610f7560096113d0565b6008555b5f610f8360095490565b11156113bf575f610f946009611417565b5f818152600760209081526040808320815160c08101835281546001600160a01b0316815260018201549381019390935260028101805495965093949293909291840191610fe190611cdf565b80601f016020809104026020016040519081016040528092919081815260200182805461100d90611cdf565b80156110585780601f1061102f57610100808354040283529160200191611058565b820191905f5260205f20905b81548152906001019060200180831161103b57829003601f168201915b5050505050815260200160038201805461107190611cdf565b80601f016020809104026020016040519081016040528092919081815260200182805461109d90611cdf565b80156110e85780601f106110bf576101008083540402835291602001916110e8565b820191905f5260205f20905b8154815290600101906020018083116110cb57829003601f168201915b505050918352505060408051808201825260048401805463ffffffff16825260058501805484516020828102820181019096528181529585019593949293858401939092919083018282801561116557602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611147575b50505091909252505050815260408051808201825260068401805463ffffffff1682526007850180548451602082810282018101909652818152958501959394929385840193909291908301828280156111e657602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116111c8575b5050505050815250508152505090505f600854836112049190611ef2565b5f54835160405163a9059cbb60e01b81526001600160a01b03918216600482015260248101849052929350169063a9059cbb906044016020604051808303815f875af1158015611256573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061127a9190611d59565b505f6112ae8360400151846060015185608001518660a00151600360029054906101000a90046001600160401b0316611424565b90506040518060c00160405280845f01516001600160a01b03168152602001426004546112db9190611d2b565b81526040858101805160208401526060808801518385015283018590526003546201000090046001600160401b03166080909301929092529051905160069161132391611d3e565b9081526040805160209281900383019020835181546001600160a01b0319166001600160a01b03909116178155918301516001830155820151600282019061136b9082611e33565b50606082015160038201906113809082611e33565b506080820151600482015560a0909101516005909101805467ffffffffffffffff19166001600160401b03909216919091179055505050600855610f79565b6003805461ffff191690555f600455565b5f815f015f815481106113e5576113e5611f05565b905f5260205f2001549050919050565b61032b82826114ab6114af565b5f61141083836114ab6114e0565b9392505050565b5f61083f826114ab61152d565b600154604051634e5bb12760e11b81525f9182916001600160a01b0390911690639cb7624e90611460908a908a908a908a908a90600401611f19565b6020604051808303815f875af115801561147c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114a09190611cc8565b979650505050505050565b1190565b5f6114b8845490565b84546001810186555f868152602090200184905590506114da848285856115be565b50505050565b5f806114ea855490565b9050805f036114fd576114fd6031611606565b5f6115088682611617565b54905084611516875f611617565b5561152486835f888861163e565b95945050505050565b5f80611537845490565b9050805f0361154a5761154a6031611606565b5f6115558582611617565b5490505f611566865f198501611617565b54865490915086908061157b5761157b611f81565b600190038181905f5260205f20015f90559055806115a45f885f0161161790919063ffffffff16565b556115b5865f1985015f848961163e565b50949350505050565b82156114da5760025f198401045f6115d68683611617565b5490506115e7818563ffffffff8616565b156115f35750506114da565b6115fe868684611733565b5092506115be565b634e487b715f52806020526024601cfd5b5f826116366116338461162f845f9081526020902090565b0190565b90565b949350505050565b6001600160ff1b038310156108d0576002838102600181019101858110156116e7575f61166b8884611617565b5490505f6116798984611617565b54905061168a828763ffffffff8816565b8061169e575061169e81878763ffffffff16565b156116e0575f6116c485856116b786868b63ffffffff16565b1515918118919091021890565b90506116d18a8983611733565b6116de8a8a838a8a61163e565b505b505061172a565b8582101561172a575f6116fa8884611617565b54905061170b818663ffffffff8716565b156117285761171b888785611733565b611728888885888861163e565b505b50505050505050565b5f61173e8484611617565b90505f61174b8584611617565b8054835490915590915550505050565b5080545f8255905f5260205f20908101906117769190611813565b50565b50805461178590611cdf565b5f825580601f10611794575050565b601f0160209004905f5260205f20908101906117769190611813565b828054828255905f5260205f20908101928215611803579160200282015b8281111561180357825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906117ce565b5061180f929150611813565b5090565b5b8082111561180f575f8155600101611814565b803563ffffffff8116811461183a575f80fd5b919050565b5f6020828403121561184f575f80fd5b61141082611827565b5f60208284031215611868575f80fd5b5035919050565b5f5b83811015611889578181015183820152602001611871565b50505f910152565b5f81518084526118a881602086016020860161186f565b601f01601f19169290920160200192915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b808310156119175784516001600160a01b031682529383019360019290920191908301906118ee565b509695505050505050565b60018060a01b038716815285602082015260c060408201525f61194860c0830187611891565b828103606084015261195a8187611891565b9050828103608084015261196e81866118bc565b905082810360a084015261198281856118bc565b9998505050505050505050565b5f805f80608085870312156119a2575f80fd5b843561ffff811681146119b3575f80fd5b935060208501356001600160401b03811681146119ce575f80fd5b93969395505050506040820135916060013590565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611a1957611a196119e3565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611a4757611a476119e3565b604052919050565b5f82601f830112611a5e575f80fd5b81356001600160401b03811115611a7757611a776119e3565b611a8a601f8201601f1916602001611a1f565b818152846020838601011115611a9e575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215611aca575f80fd5b81356001600160401b03811115611adf575f80fd5b61163684828501611a4f565b60018060a01b038716815285602082015260c060408201525f611b1160c0830187611891565b8281036060840152611b238187611891565b9150508360808301526001600160401b03831660a0830152979650505050505050565b5f60408284031215611b56575f80fd5b611b5e6119f7565b9050611b6982611827565b81526020808301356001600160401b0380821115611b85575f80fd5b818501915085601f830112611b98575f80fd5b813581811115611baa57611baa6119e3565b8060051b9150611bbb848301611a1f565b8181529183018401918481019088841115611bd4575f80fd5b938501935b83851015611c0957843592506001600160a01b0383168314611bf9575f80fd5b8282529385019390850190611bd9565b808688015250505050505092915050565b5f805f805f60a08688031215611c2e575f80fd5b8535945060208601356001600160401b0380821115611c4b575f80fd5b611c5789838a01611a4f565b95506040880135915080821115611c6c575f80fd5b611c7889838a01611a4f565b94506060880135915080821115611c8d575f80fd5b611c9989838a01611b46565b93506080880135915080821115611cae575f80fd5b50611cbb88828901611b46565b9150509295509295909350565b5f60208284031215611cd8575f80fd5b5051919050565b600181811c90821680611cf357607f821691505b602082108103611d1157634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561083f5761083f611d17565b5f8251611d4f81846020870161186f565b9190910192915050565b5f60208284031215611d69575f80fd5b81518015158114611410575f80fd5b5f808354611d8581611cdf565b60018281168015611d9d5760018114611db257611dde565b60ff1984168752821515830287019450611dde565b875f526020805f205f5b85811015611dd55781548a820152908401908201611dbc565b50505082870194505b50929695505050505050565b601f821115611e2e57805f5260205f20601f840160051c81016020851015611e0f5750805b601f840160051c820191505b818110156108d0575f8155600101611e1b565b505050565b81516001600160401b03811115611e4c57611e4c6119e3565b611e6081611e5a8454611cdf565b84611dea565b602080601f831160018114611e93575f8415611e7c5750858301515b5f19600386901b1c1916600185901b178555611eea565b5f85815260208120601f198616915b82811015611ec157888601518255948401946001909101908401611ea2565b5085821015611ede57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b8181038181111561083f5761083f611d17565b634e487b7160e01b5f52603260045260245ffd5b60a081525f611f2b60a0830188611891565b8281036020840152611f3d8188611891565b90508281036040840152611f5181876118bc565b90508281036060840152611f6581866118bc565b9150506001600160401b03831660808301529695505050505050565b634e487b7160e01b5f52603160045260245ffdfea264697066735822122038bfb7f59b1f719c5046284380279bc160d8583204c19630335c35f0109652e964736f6c63430008190033",
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

// CurrentValidators is a free data retrieval call binding the contract method 0x5fcf2538.
//
// Solidity: function currentValidators(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 Weight)
func (_SlotAuctionManager *SlotAuctionManagerCaller) CurrentValidators(opts *bind.CallOpts, nodeID []byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID [32]byte
	Weight       uint64
}, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "currentValidators", nodeID)

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

// CurrentValidators is a free data retrieval call binding the contract method 0x5fcf2538.
//
// Solidity: function currentValidators(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 Weight)
func (_SlotAuctionManager *SlotAuctionManagerSession) CurrentValidators(nodeID []byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID [32]byte
	Weight       uint64
}, error) {
	return _SlotAuctionManager.Contract.CurrentValidators(&_SlotAuctionManager.CallOpts, nodeID)
}

// CurrentValidators is a free data retrieval call binding the contract method 0x5fcf2538.
//
// Solidity: function currentValidators(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, bytes32 validationID, uint64 Weight)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) CurrentValidators(nodeID []byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID [32]byte
	Weight       uint64
}, error) {
	return _SlotAuctionManager.Contract.CurrentValidators(&_SlotAuctionManager.CallOpts, nodeID)
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

// InitiateAuction is a paid mutator transaction binding the contract method 0x3e89e29a.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 auctionLength, uint256 validationLength) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactor) InitiateAuction(opts *bind.TransactOpts, validatorslots uint16, weight uint64, auctionLength *big.Int, validationLength *big.Int) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "initiateAuction", validatorslots, weight, auctionLength, validationLength)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x3e89e29a.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 auctionLength, uint256 validationLength) returns()
func (_SlotAuctionManager *SlotAuctionManagerSession) InitiateAuction(validatorslots uint16, weight uint64, auctionLength *big.Int, validationLength *big.Int) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateAuction(&_SlotAuctionManager.TransactOpts, validatorslots, weight, auctionLength, validationLength)
}

// InitiateAuction is a paid mutator transaction binding the contract method 0x3e89e29a.
//
// Solidity: function initiateAuction(uint16 validatorslots, uint64 weight, uint256 auctionLength, uint256 validationLength) returns()
func (_SlotAuctionManager *SlotAuctionManagerTransactorSession) InitiateAuction(validatorslots uint16, weight uint64, auctionLength *big.Int, validationLength *big.Int) (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.InitiateAuction(&_SlotAuctionManager.TransactOpts, validatorslots, weight, auctionLength, validationLength)
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
