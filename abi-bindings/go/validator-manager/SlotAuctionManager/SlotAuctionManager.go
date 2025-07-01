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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"bidderInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"currentValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"validationID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"validatorslots\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"auctionLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validationLength\",\"type\":\"uint256\"}],\"name\":\"initiateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"nodeHasBid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"hasBid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peekTop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validationTimeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSlots\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001805460ff60a01b191690555f6002819055600380546001600160501b03191690556004553480156034575f80fd5b506040516120c63803806120c6833981016040819052605191609a565b5f80546001600160a01b039384166001600160a01b0319918216179091556001805492909316911617905560c6565b80516001600160a01b03811681146095575f80fd5b919050565b5f806040838503121560aa575f80fd5b60b1836080565b915060bd602084016080565b90509250929050565b611ff3806100d35f395ff3fe608060405234801561000f575f80fd5b506004361061011b575f3560e01c80638c3c78e3116100a9578063b6e6a2ca1161006e578063b6e6a2ca1461025f578063da4312a414610285578063df39307514610298578063f826c832146102bc578063fe67a54b146102ea575f80fd5b80638c3c78e3146102315780639681d94014610239578063a3a65e481461024c578063a476f6751461025f578063ae9483e014610272575f80fd5b80631931f2c7116100ef5780631931f2c7146101b057806336339388146101c75780633e89e29a146101f15780634b449cba146102045780635fcf25381461020d575f80fd5b8062d841f81461011f57806305af8256146101555780630d5daf3b1461016a5780631077830a1461018f575b5f80fd5b600354610138906201000090046001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b6101686101633660046118ba565b6102f2565b005b61017d6101783660046118da565b610367565b60405161014c969594939291906119a4565b60035461019d9061ffff1681565b60405161ffff909116815260200161014c565b6101b960045481565b60405190815260200161014c565b5f546101d9906001600160a01b031681565b6040516001600160a01b03909116815260200161014c565b6101686101ff366004611a11565b61059c565b6101b960025481565b61022061021b366004611b3c565b610665565b60405161014c959493929190611b6d565b6101b96107bb565b6101b96102473660046118ba565b6107f4565b6101b961025a3660046118ba565b61086c565b61016861026d3660046118da565b6108a2565b6001546101d9906001600160a01b031681565b610168610293366004611c92565b6108fe565b6001546102ac90600160a01b900460ff1681565b604051901515815260200161014c565b6102ac6102ca366004611b3c565b805160208183018101805160058252928201919093012091525460ff1681565b6102ac610d87565b60015460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af115801561033f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103639190611d40565b5050565b60076020525f90815260409020805460018201546002830180546001600160a01b0390931693919261039890611d57565b80601f01602080910402602001604051908101604052809291908181526020018280546103c490611d57565b801561040f5780601f106103e65761010080835404028352916020019161040f565b820191905f5260205f20905b8154815290600101906020018083116103f257829003601f168201915b50505050509080600301805461042490611d57565b80601f016020809104026020016040519081016040528092919081815260200182805461045090611d57565b801561049b5780601f106104725761010080835404028352916020019161049b565b820191905f5260205f20905b81548152906001019060200180831161047e57829003601f168201915b505060408051808201825260048701805463ffffffff16825260058801805484516020828102820181019096528181529899989397509195508387019450919283018282801561051257602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116104f4575b50505091909252505060408051808201825260068501805463ffffffff1682526007860180548451602082810282018101909652818152969796939550919380860193929083018282801561058e57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610570575b505050505081525050905086565b600154600160a01b900460ff16156105fb5760405162461bcd60e51b815260206004820152601760248201527f41756374696f6e20616c72656164792072756e6e696e6700000000000000000060448201526064015b60405180910390fd5b6003805469ffffffffffffffff00001916620100006001600160401b038616021790556106288242611da3565b6002556003805461ffff191661ffff8616179055600481905561064c60085f6117d6565b50506001805460ff60a01b1916600160a01b1790555050565b80516020818301810180516006825292820191909301209152805460018201546002830180546001600160a01b039093169391926106a290611d57565b80601f01602080910402602001604051908101604052809291908181526020018280546106ce90611d57565b80156107195780601f106106f057610100808354040283529160200191610719565b820191905f5260205f20905b8154815290600101906020018083116106fc57829003601f168201915b50505050509080600301805461072e90611d57565b80601f016020809104026020016040519081016040528092919081815260200182805461075a90611d57565b80156107a55780601f1061077c576101008083540402835291602001916107a5565b820191905f5260205f20905b81548152906001019060200180831161078857829003601f168201915b5050506004909301549192505063ffffffff1685565b5f6107c560085490565b15806107db5750600154600160a01b900460ff16155b156107e557505f90565b6107ef6008611469565b905090565b60015460405163025a076560e61b815263ffffffff831660048201525f916001600160a01b031690639681d940906024015b6020604051808303815f875af1158015610842573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108669190611d40565b92915050565b600154604051631474cbc960e31b815263ffffffff831660048201525f916001600160a01b03169063a3a65e4890602401610826565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b1580156108e5575f80fd5b505af11580156108f7573d5f803e3d5ffd5b5050505050565b600154600160a01b900460ff166109505760405162461bcd60e51b815260206004820152601660248201527541756374696f6e206973206e6f742072756e6e696e6760501b60448201526064016105f2565b60016005856040516109629190611db6565b908152604051908190036020019020805491151560ff1990921691909117905561098b60085490565b5f0361099f5761099f60088661148e611492565b846109aa6008611469565b1180156109d457506003546109c49061ffff166001611dd1565b61ffff166109d160085490565b10155b6108f7575f546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015610a2b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a4f9190611df3565b610a9b5760405162461bcd60e51b815260206004820152601f60248201527f496e73756666696369656e742066756e6473206261736564206f6e206269640060448201526064016105f2565b5f858152600760205260409020546001600160a01b0316158015610acd57505f85815260076020526040902060010154155b610b275760405162461bcd60e51b815260206004820152602560248201527f42696420616c726561647920706c616365642c20796f75206d6179206269642060448201526430b3b0b4b760d91b60648201526084016105f2565b600354610b399061ffff166001611dd1565b61ffff16610b4660085490565b10610c6c575f610b5a60088761148e6114c3565b5f805482825260076020526040918290208054600190910154925163a9059cbb60e01b81526001600160a01b0391821660048201526024810193909352929350919091169063a9059cbb906044016020604051808303815f875af1158015610bc4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610be89190611df3565b505f81815260076020526040812080546001600160a01b03191681556001810182905590610c1960028301826117f4565b610c26600383015f6117f4565b60048201805463ffffffff191681555f610c4360058501826117d6565b505060068201805463ffffffff191681555f610c6260078501826117d6565b5050505050610c7a565b610c7a60088661148e611492565b6040805160c0810182523381526020808201888152828401888152606084018890526080840187905260a084018690525f8a8152600790935293909120825181546001600160a01b0319166001600160a01b039091161781559051600182015591519091906002820190610cee9082611e5b565b5060608201516003820190610d039082611e5b565b506080820151805160048301805463ffffffff191663ffffffff9092169190911781556020808301518051610d3e926005870192019061182b565b50505060a0820151805160068301805463ffffffff191663ffffffff9092169190911781556020808301518051610d7b926007870192019061182b565b50505050505050505050565b6001545f90600160a01b900460ff16610de25760405162461bcd60e51b815260206004820152601760248201527f41756374696f6e206e6f7420696e2070726f677265737300000000000000000060448201526064016105f2565b6002544211610e335760405162461bcd60e51b815260206004820152601e60248201527f4265666f72652073657420656e64206f662061756374696f6e2074696d65000060448201526064016105f2565b6001805460ff60a01b191681555f600255600354610e579161ffff90911690611dd1565b61ffff16610e6460085490565b101561110a575f610e756008611469565b5f818152600760209081526040808320815160c08101835281546001600160a01b0316815260018201549381019390935260028101805495965093949293909291840191610ec290611d57565b80601f0160208091040260200160405190810160405280929190818152602001828054610eee90611d57565b8015610f395780601f10610f1057610100808354040283529160200191610f39565b820191905f5260205f20905b815481529060010190602001808311610f1c57829003601f168201915b50505050508152602001600382018054610f5290611d57565b80601f0160208091040260200160405190810160405280929190818152602001828054610f7e90611d57565b8015610fc95780601f10610fa057610100808354040283529160200191610fc9565b820191905f5260205f20905b815481529060010190602001808311610fac57829003601f168201915b505050918352505060408051808201825260048401805463ffffffff16825260058501805484516020828102820181019096528181529585019593949293858401939092919083018282801561104657602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611028575b50505091909252505050815260408051808201825260068401805463ffffffff1682526007850180548451602082810282018101909652818152958501959394929385840193909291908301828280156110c757602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116110a9575b5050509190925250505090525060408101516060820151608083015160a0840151600354949550611106946201000090046001600160401b0316611510565b5050505b60085415611454575f61111d6008611469565b90506111296008611597565b505f6111356008611469565b90505f6111428383611f1a565b5f838152600760209081526040808320815160c08101835281546001600160a01b031681526001820154938101939093526002810180549596509394929390929184019161118f90611d57565b80601f01602080910402602001604051908101604052809291908181526020018280546111bb90611d57565b80156112065780601f106111dd57610100808354040283529160200191611206565b820191905f5260205f20905b8154815290600101906020018083116111e957829003601f168201915b5050505050815260200160038201805461121f90611d57565b80601f016020809104026020016040519081016040528092919081815260200182805461124b90611d57565b80156112965780601f1061126d57610100808354040283529160200191611296565b820191905f5260205f20905b81548152906001019060200180831161127957829003601f168201915b505050918352505060408051808201825260048401805463ffffffff16825260058501805484516020828102820181019096528181529585019593949293858401939092919083018282801561131357602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116112f5575b50505091909252505050815260408051808201825260068401805463ffffffff16825260078501805484516020828102820181019096528181529585019593949293858401939092919083018282801561139457602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611376575b505050919092525050509052505f54815160405163a9059cbb60e01b81526001600160a01b03918216600482015260248101869052929350169063a9059cbb906044016020604051808303815f875af11580156113f3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114179190611df3565b5061144a8160400151826060015183608001518460a00151600360029054906101000a90046001600160401b0316611510565b505050505061110a565b506003805461ffff191690555f600455600190565b5f815f015f8154811061147e5761147e611f2d565b905f5260205f2001549050919050565b1090565b5f61149b845490565b84546001810186555f868152602090200184905590506114bd848285856115a4565b50505050565b5f806114cd855490565b9050805f036114e0576114e060316115ec565b5f6114eb86826115fd565b549050846114f9875f6115fd565b5561150786835f8888611624565b95945050505050565b600154604051634e5bb12760e11b81525f9182916001600160a01b0390911690639cb7624e9061154c908a908a908a908a908a90600401611f41565b6020604051808303815f875af1158015611568573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061158c9190611d40565b979650505050505050565b5f6108668261171961171d565b82156114bd5760025f198401045f6115bc86836115fd565b5490506115cd818563ffffffff8616565b156115d95750506114bd565b6115e48686846117ae565b5092506115a4565b634e487b715f52806020526024601cfd5b5f8261161c61161984611615845f9081526020902090565b0190565b90565b949350505050565b6001600160ff1b038310156108f7576002838102600181019101858110156116cd575f61165188846115fd565b5490505f61165f89846115fd565b549050611670828763ffffffff8816565b80611684575061168481878763ffffffff16565b156116c6575f6116aa858561169d86868b63ffffffff16565b1515918118919091021890565b90506116b78a89836117ae565b6116c48a8a838a8a611624565b505b5050611710565b85821015611710575f6116e088846115fd565b5490506116f1818663ffffffff8716565b1561170e576117018887856117ae565b61170e8888858888611624565b505b50505050505050565b1190565b5f80611727845490565b9050805f0361173a5761173a60316115ec565b5f61174585826115fd565b5490505f611756865f1985016115fd565b54865490915086908061176b5761176b611fa9565b600190038181905f5260205f20015f90559055806117945f885f016115fd90919063ffffffff16565b556117a5865f1985015f8489611624565b50949350505050565b5f6117b984846115fd565b90505f6117c685846115fd565b8054835490915590915550505050565b5080545f8255905f5260205f20908101906117f1919061188e565b50565b50805461180090611d57565b5f825580601f1061180f575050565b601f0160209004905f5260205f20908101906117f1919061188e565b828054828255905f5260205f2090810192821561187e579160200282015b8281111561187e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611849565b5061188a92915061188e565b5090565b5b8082111561188a575f815560010161188f565b803563ffffffff811681146118b5575f80fd5b919050565b5f602082840312156118ca575f80fd5b6118d3826118a2565b9392505050565b5f602082840312156118ea575f80fd5b5035919050565b5f5b8381101561190b5781810151838201526020016118f3565b50505f910152565b5f815180845261192a8160208601602086016118f1565b601f01601f19169290920160200192915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b808310156119995784516001600160a01b03168252938301936001929092019190830190611970565b509695505050505050565b60018060a01b038716815285602082015260c060408201525f6119ca60c0830187611913565b82810360608401526119dc8187611913565b905082810360808401526119f0818661193e565b905082810360a0840152611a04818561193e565b9998505050505050505050565b5f805f8060808587031215611a24575f80fd5b843561ffff81168114611a35575f80fd5b935060208501356001600160401b0381168114611a50575f80fd5b93969395505050506040820135916060013590565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611a9b57611a9b611a65565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611ac957611ac9611a65565b604052919050565b5f82601f830112611ae0575f80fd5b81356001600160401b03811115611af957611af9611a65565b611b0c601f8201601f1916602001611aa1565b818152846020838601011115611b20575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215611b4c575f80fd5b81356001600160401b03811115611b61575f80fd5b61161c84828501611ad1565b60018060a01b038616815284602082015260a060408201525f611b9360a0830186611913565b8281036060840152611ba58186611913565b91505063ffffffff831660808301529695505050505050565b5f60408284031215611bce575f80fd5b611bd6611a79565b9050611be1826118a2565b81526020808301356001600160401b0380821115611bfd575f80fd5b818501915085601f830112611c10575f80fd5b813581811115611c2257611c22611a65565b8060051b9150611c33848301611aa1565b8181529183018401918481019088841115611c4c575f80fd5b938501935b83851015611c8157843592506001600160a01b0383168314611c71575f80fd5b8282529385019390850190611c51565b808688015250505050505092915050565b5f805f805f60a08688031215611ca6575f80fd5b8535945060208601356001600160401b0380821115611cc3575f80fd5b611ccf89838a01611ad1565b95506040880135915080821115611ce4575f80fd5b611cf089838a01611ad1565b94506060880135915080821115611d05575f80fd5b611d1189838a01611bbe565b93506080880135915080821115611d26575f80fd5b50611d3388828901611bbe565b9150509295509295909350565b5f60208284031215611d50575f80fd5b5051919050565b600181811c90821680611d6b57607f821691505b602082108103611d8957634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561086657610866611d8f565b5f8251611dc78184602087016118f1565b9190910192915050565b61ffff818116838216019080821115611dec57611dec611d8f565b5092915050565b5f60208284031215611e03575f80fd5b815180151581146118d3575f80fd5b601f821115611e5657805f5260205f20601f840160051c81016020851015611e375750805b601f840160051c820191505b818110156108f7575f8155600101611e43565b505050565b81516001600160401b03811115611e7457611e74611a65565b611e8881611e828454611d57565b84611e12565b602080601f831160018114611ebb575f8415611ea45750858301515b5f19600386901b1c1916600185901b178555611f12565b5f85815260208120601f198616915b82811015611ee957888601518255948401946001909101908401611eca565b5085821015611f0657878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b8181038181111561086657610866611d8f565b634e487b7160e01b5f52603260045260245ffd5b60a081525f611f5360a0830188611913565b8281036020840152611f658188611913565b90508281036040840152611f79818761193e565b90508281036060840152611f8d818661193e565b9150506001600160401b03831660808301529695505050505050565b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220c2e2c7a3922356f809925ca98255f85f242fa166abd2b40dd88d965be521f60264736f6c63430008190033",
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
// Solidity: function currentValidators(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, uint32 validationID)
func (_SlotAuctionManager *SlotAuctionManagerCaller) CurrentValidators(opts *bind.CallOpts, nodeID []byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID uint32
}, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "currentValidators", nodeID)

	outstruct := new(struct {
		Addr         common.Address
		EndTime      *big.Int
		NodeID       []byte
		BlsPublicKey []byte
		ValidationID uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NodeID = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.BlsPublicKey = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.ValidationID = *abi.ConvertType(out[4], new(uint32)).(*uint32)

	return *outstruct, err

}

// CurrentValidators is a free data retrieval call binding the contract method 0x5fcf2538.
//
// Solidity: function currentValidators(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, uint32 validationID)
func (_SlotAuctionManager *SlotAuctionManagerSession) CurrentValidators(nodeID []byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID uint32
}, error) {
	return _SlotAuctionManager.Contract.CurrentValidators(&_SlotAuctionManager.CallOpts, nodeID)
}

// CurrentValidators is a free data retrieval call binding the contract method 0x5fcf2538.
//
// Solidity: function currentValidators(bytes nodeID) view returns(address addr, uint256 endTime, bytes nodeID, bytes blsPublicKey, uint32 validationID)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) CurrentValidators(nodeID []byte) (struct {
	Addr         common.Address
	EndTime      *big.Int
	NodeID       []byte
	BlsPublicKey []byte
	ValidationID uint32
}, error) {
	return _SlotAuctionManager.Contract.CurrentValidators(&_SlotAuctionManager.CallOpts, nodeID)
}

// NodeHasBid is a free data retrieval call binding the contract method 0xf826c832.
//
// Solidity: function nodeHasBid(bytes nodeID) view returns(bool hasBid)
func (_SlotAuctionManager *SlotAuctionManagerCaller) NodeHasBid(opts *bind.CallOpts, nodeID []byte) (bool, error) {
	var out []interface{}
	err := _SlotAuctionManager.contract.Call(opts, &out, "nodeHasBid", nodeID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NodeHasBid is a free data retrieval call binding the contract method 0xf826c832.
//
// Solidity: function nodeHasBid(bytes nodeID) view returns(bool hasBid)
func (_SlotAuctionManager *SlotAuctionManagerSession) NodeHasBid(nodeID []byte) (bool, error) {
	return _SlotAuctionManager.Contract.NodeHasBid(&_SlotAuctionManager.CallOpts, nodeID)
}

// NodeHasBid is a free data retrieval call binding the contract method 0xf826c832.
//
// Solidity: function nodeHasBid(bytes nodeID) view returns(bool hasBid)
func (_SlotAuctionManager *SlotAuctionManagerCallerSession) NodeHasBid(nodeID []byte) (bool, error) {
	return _SlotAuctionManager.Contract.NodeHasBid(&_SlotAuctionManager.CallOpts, nodeID)
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
// Solidity: function endAuction() returns(bool)
func (_SlotAuctionManager *SlotAuctionManagerTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns(bool)
func (_SlotAuctionManager *SlotAuctionManagerSession) EndAuction() (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.EndAuction(&_SlotAuctionManager.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns(bool)
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
