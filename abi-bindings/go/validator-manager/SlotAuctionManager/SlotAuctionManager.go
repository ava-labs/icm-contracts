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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vmAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER\",\"outputs\":[{\"internalType\":\"contractIValidatorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"bidderInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRemoval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"currentValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"Weight\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"validatorslots\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"auctionLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validationLength\",\"type\":\"uint256\"}],\"name\":\"initiateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateRemoveInitialValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"nodeHasBid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"hasBid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peekTop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validationTimeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSlots\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50604051612592380380612592833981016040819052602b916074565b5f80546001600160a01b039384166001600160a01b0319918216179091556001805492909316911617905560a0565b80516001600160a01b0381168114606f575f80fd5b919050565b5f80604083850312156084575f80fd5b608b83605a565b9150609760208401605a565b90509250929050565b6124e5806100ad5f395ff3fe608060405234801561000f575f80fd5b506004361061011b575f3560e01c80638c3c78e3116100a9578063b6e6a2ca1161006e578063b6e6a2ca14610260578063da4312a414610286578063df39307514610299578063f826c832146102bd578063fe67a54b146102eb575f80fd5b80638c3c78e3146102325780639681d9401461023a578063a3a65e481461024d578063a476f67514610260578063ae9483e014610273575f80fd5b80631931f2c7116100ef5780631931f2c7146101b057806336339388146101c75780633e89e29a146101f15780634b449cba146102045780635fcf25381461020d575f80fd5b8062d841f81461011f57806305af8256146101555780630d5daf3b1461016a5780631077830a1461018f575b5f80fd5b600354610138906201000090046001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b610168610163366004611d46565b610300565b005b61017d610178366004611d5f565b610375565b60405161014c96959493929190611e29565b60035461019d9061ffff1681565b60405161ffff909116815260200161014c565b6101b960045481565b60405190815260200161014c565b5f546101d9906001600160a01b031681565b6040516001600160a01b03909116815260200161014c565b6101686101ff366004611e96565b6105aa565b6101b960025481565b61022061021b366004611fc1565b610673565b60405161014c96959493929190611ff2565b6101b96107d3565b6101b9610248366004611d46565b61080c565b6101b961025b366004611d46565b610884565b61016861026e366004611d5f565b6108ba565b6001546101d9906001600160a01b031681565b610168610294366004612121565b610916565b6001546102ad90600160a01b900460ff1681565b604051901515815260200161014c565b6102ad6102cb366004611fc1565b805160208183018101805160058252928201919093012091525460ff1681565b6102f3610ef9565b60405161014c91906121cf565b60015460405163025a076560e61b815263ffffffff831660048201526001600160a01b0390911690639681d940906024016020604051808303815f875af115801561034d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103719190612212565b5050565b60076020525f90815260409020805460018201546002830180546001600160a01b039093169391926103a690612229565b80601f01602080910402602001604051908101604052809291908181526020018280546103d290612229565b801561041d5780601f106103f45761010080835404028352916020019161041d565b820191905f5260205f20905b81548152906001019060200180831161040057829003601f168201915b50505050509080600301805461043290612229565b80601f016020809104026020016040519081016040528092919081815260200182805461045e90612229565b80156104a95780601f10610480576101008083540402835291602001916104a9565b820191905f5260205f20905b81548152906001019060200180831161048c57829003601f168201915b505060408051808201825260048701805463ffffffff16825260058801805484516020828102820181019096528181529899989397509195508387019450919283018282801561052057602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610502575b50505091909252505060408051808201825260068501805463ffffffff1682526007860180548451602082810282018101909652818152969796939550919380860193929083018282801561059c57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161057e575b505050505081525050905086565b600154600160a01b900460ff16156106095760405162461bcd60e51b815260206004820152601760248201527f41756374696f6e20616c72656164792072756e6e696e6700000000000000000060448201526064015b60405180910390fd5b6003805469ffffffffffffffff00001916620100006001600160401b038616021790556106368242612275565b6002556003805461ffff191661ffff8616179055600481905561065a60085f611c62565b50506001805460ff60a01b1916600160a01b1790555050565b80516020818301810180516006825292820191909301209152805460018201546002830180546001600160a01b039093169391926106b090612229565b80601f01602080910402602001604051908101604052809291908181526020018280546106dc90612229565b80156107275780601f106106fe57610100808354040283529160200191610727565b820191905f5260205f20905b81548152906001019060200180831161070a57829003601f168201915b50505050509080600301805461073c90612229565b80601f016020809104026020016040519081016040528092919081815260200182805461076890612229565b80156107b35780601f1061078a576101008083540402835291602001916107b3565b820191905f5260205f20905b81548152906001019060200180831161079657829003601f168201915b5050505060048301546005909301549192916001600160401b0316905086565b5f6107dd60085490565b15806107f35750600154600160a01b900460ff16155b156107fd57505f90565b61080760086118c8565b905090565b60015460405163025a076560e61b815263ffffffff831660048201525f916001600160a01b031690639681d940906024015b6020604051808303815f875af115801561085a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061087e9190612212565b92915050565b600154604051631474cbc960e31b815263ffffffff831660048201525f916001600160a01b03169063a3a65e489060240161083e565b600154604051635b73516560e11b8152600481018390526001600160a01b039091169063b6e6a2ca906024015f604051808303815f87803b1580156108fd575f80fd5b505af115801561090f573d5f803e3d5ffd5b5050505050565b600154600160a01b900460ff166109685760405162461bcd60e51b815260206004820152601660248201527541756374696f6e206973206e6f742072756e6e696e6760501b6044820152606401610600565b5f6001600160a01b03166006856040516109829190612288565b908152604051908190036020019020546001600160a01b0316146109e85760405162461bcd60e51b815260206004820152601b60248201527f4e6f646520697320616c726561647920612076616c696461746f7200000000006044820152606401610600565b6005846040516109f89190612288565b9081526040519081900360200190205460ff1615610a4f5760405162461bcd60e51b8152602060048201526014602482015273139bd919481a185cc8185b1c9958591e48189a5960621b6044820152606401610600565b5f858152600760205260409020546001600160a01b0316158015610a8157505f85815260076020526040902060010154155b610acd5760405162461bcd60e51b815260206004820181905260248201527f4475706c6963617465206269642c20796f75206d61792062696420616761696e6044820152606401610600565b6001600585604051610adf9190612288565b908152604051908190036020019020805491151560ff19909216919091179055610b0860085490565b1580610b305750600354610b219061ffff1660016122a3565b61ffff16610b2e60085490565b105b15610c04575f546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015610b88573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bac91906122c5565b610bf45760405162461bcd60e51b8152602060048201526019602482015278125b9cdd59999a58da595b9d08199d5b991cc81d1bc8189a59603a1b6044820152606401610600565b610bff6008866118ed565b610dec565b84610c0f60086118c8565b1015610de7575f546040516323b872dd60e01b8152336004820152306024820152604481018790526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015610c68573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c8c91906122c5565b610cd45760405162461bcd60e51b8152602060048201526019602482015278125b9cdd59999a58da595b9d08199d5b991cc81d1bc8189a59603a1b6044820152606401610600565b5f610ce06008876118fa565b5f8054828252600760205260409182902054915163a9059cbb60e01b81526001600160a01b03928316600482015260248101849052929350169063a9059cbb906044016020604051808303815f875af1158015610d3f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d6391906122c5565b505f81815260076020526040812080546001600160a01b03191681556001810182905590610d946002830182611c80565b610da1600383015f611c80565b60048201805463ffffffff191681555f610dbe6005850182611c62565b505060068201805463ffffffff191681555f610ddd6007850182611c62565b5050505050610dec565b61090f565b6040805160c0810182523381526020808201888152828401888152606084018890526080840187905260a084018690525f8a8152600790935293909120825181546001600160a01b0319166001600160a01b039091161781559051600182015591519091906002820190610e60908261232d565b5060608201516003820190610e75908261232d565b506080820151805160048301805463ffffffff191663ffffffff9092169190911781556020808301518051610eb09260058701920190611cb7565b50505060a0820151805160068301805463ffffffff191663ffffffff9092169190911781556020808301518051610eed9260078701920190611cb7565b50505050505050505050565b600154606090600160a01b900460ff16610f555760405162461bcd60e51b815260206004820152601760248201527f41756374696f6e206e6f7420696e2070726f67726573730000000000000000006044820152606401610600565b6002544211610fa65760405162461bcd60e51b815260206004820152601b60248201527f41756374696f6e20656e6474696d65206e6f74207265616368656400000000006044820152606401610600565b6003545f90610fc19061ffff16610fbc60085490565b61190f565b6001600160401b03811115610fd857610fd8611eea565b604051908082528060200260200182016040528015611001578160200160208202803683370190505b506001805460ff60a01b191681555f60028190556003549293509161102c9161ffff909116906122a3565b61ffff1661103960085490565b108015611047575060085415155b15611427575f61105760086118c8565b5f818152600760209081526040808320815160c08101835281546001600160a01b03168152600182015493810193909352600281018054959650939492939092918401916110a490612229565b80601f01602080910402602001604051908101604052809291908181526020018280546110d090612229565b801561111b5780601f106110f25761010080835404028352916020019161111b565b820191905f5260205f20905b8154815290600101906020018083116110fe57829003601f168201915b5050505050815260200160038201805461113490612229565b80601f016020809104026020016040519081016040528092919081815260200182805461116090612229565b80156111ab5780601f10611182576101008083540402835291602001916111ab565b820191905f5260205f20905b81548152906001019060200180831161118e57829003601f168201915b505050918352505060408051808201825260048401805463ffffffff16825260058501805484516020828102820181019096528181529585019593949293858401939092919083018282801561122857602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161120a575b50505091909252505050815260408051808201825260068401805463ffffffff1682526007850180548451602082810282018101909652818152958501959394929385840193909291908301828280156112a957602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161128b575b5050505050815250508152505090505f6112eb8260400151836060015184608001518560a00151600360029054906101000a90046001600160401b031661191e565b90506040518060c00160405280835f01516001600160a01b03168152602001426004546113189190612275565b81526040848101805160208401526060808701518385015283018590526003546201000090046001600160401b03166080909301929092529051905160069161136091612288565b9081526040805160209281900383019020835181546001600160a01b0319166001600160a01b0390911617815591830151600183015582015160028201906113a8908261232d565b50606082015160038201906113bd908261232d565b506080820151600482015560a0909101516005909101805467ffffffffffffffff19166001600160401b0390921691909117905584518190869061ffff871690811061140b5761140b6123ec565b60209081029190910101528361142081612400565b9450505050505b600161143260085490565b11156118b3575f61144360086118c8565b905061144f60086119a5565b505f61145b60086118c8565b90505f6114688383612420565b5f838152600760209081526040808320815160c08101835281546001600160a01b03168152600182015493810193909352600281018054959650939492939092918401916114b590612229565b80601f01602080910402602001604051908101604052809291908181526020018280546114e190612229565b801561152c5780601f106115035761010080835404028352916020019161152c565b820191905f5260205f20905b81548152906001019060200180831161150f57829003601f168201915b5050505050815260200160038201805461154590612229565b80601f016020809104026020016040519081016040528092919081815260200182805461157190612229565b80156115bc5780601f10611593576101008083540402835291602001916115bc565b820191905f5260205f20905b81548152906001019060200180831161159f57829003601f168201915b505050918352505060408051808201825260048401805463ffffffff16825260058501805484516020828102820181019096528181529585019593949293858401939092919083018282801561163957602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161161b575b50505091909252505050815260408051808201825260068401805463ffffffff1682526007850180548451602082810282018101909652818152958501959394929385840193909291908301828280156116ba57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161169c575b505050919092525050509052505f54815160405163a9059cbb60e01b81526001600160a01b03918216600482015260248101869052929350169063a9059cbb906044016020604051808303815f875af1158015611719573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061173d91906122c5565b505f6117718260400151836060015184608001518560a00151600360029054906101000a90046001600160401b031661191e565b90506040518060c00160405280835f01516001600160a01b031681526020014260045461179e9190612275565b81526040848101805160208401526060808701518385015283018590526003546201000090046001600160401b0316608090930192909252905190516006916117e691612288565b9081526040805160209281900383019020835181546001600160a01b0319166001600160a01b03909116178155918301516001830155820151600282019061182e908261232d565b5060608201516003820190611843908261232d565b506080820151600482015560a0909101516005909101805467ffffffffffffffff19166001600160401b0390921691909117905586518190889061ffff8916908110611891576118916123ec565b6020908102919091010152856118a681612400565b9650505050505050611427565b506003805461ffff191690555f600455919050565b5f815f015f815481106118dd576118dd6123ec565b905f5260205f2001549050919050565b61037182826119b26119b6565b5f61190883836119b26119e7565b9392505050565b5f828218828410028218611908565b600154604051634e5bb12760e11b81525f9182916001600160a01b0390911690639cb7624e9061195a908a908a908a908a908a90600401612433565b6020604051808303815f875af1158015611976573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061199a9190612212565b979650505050505050565b5f61087e826119b2611a34565b1190565b5f6119bf845490565b84546001810186555f868152602090200184905590506119e184828585611ac5565b50505050565b5f806119f1855490565b9050805f03611a0457611a046031611b0d565b5f611a0f8682611b1e565b54905084611a1d875f611b1e565b55611a2b86835f8888611b45565b95945050505050565b5f80611a3e845490565b9050805f03611a5157611a516031611b0d565b5f611a5c8582611b1e565b5490505f611a6d865f198501611b1e565b548654909150869080611a8257611a8261249b565b600190038181905f5260205f20015f9055905580611aab5f885f01611b1e90919063ffffffff16565b55611abc865f1985015f8489611b45565b50949350505050565b82156119e15760025f198401045f611add8683611b1e565b549050611aee818563ffffffff8616565b15611afa5750506119e1565b611b05868684611c3a565b509250611ac5565b634e487b715f52806020526024601cfd5b5f82611b3d611b3a84611b36845f9081526020902090565b0190565b90565b949350505050565b6001600160ff1b0383101561090f57600283810260018101910185811015611bee575f611b728884611b1e565b5490505f611b808984611b1e565b549050611b91828763ffffffff8816565b80611ba55750611ba581878763ffffffff16565b15611be7575f611bcb8585611bbe86868b63ffffffff16565b1515918118919091021890565b9050611bd88a8983611c3a565b611be58a8a838a8a611b45565b505b5050611c31565b85821015611c31575f611c018884611b1e565b549050611c12818663ffffffff8716565b15611c2f57611c22888785611c3a565b611c2f8888858888611b45565b505b50505050505050565b5f611c458484611b1e565b90505f611c528584611b1e565b8054835490915590915550505050565b5080545f8255905f5260205f2090810190611c7d9190611d1a565b50565b508054611c8c90612229565b5f825580601f10611c9b575050565b601f0160209004905f5260205f2090810190611c7d9190611d1a565b828054828255905f5260205f20908101928215611d0a579160200282015b82811115611d0a57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611cd5565b50611d16929150611d1a565b5090565b5b80821115611d16575f8155600101611d1b565b803563ffffffff81168114611d41575f80fd5b919050565b5f60208284031215611d56575f80fd5b61190882611d2e565b5f60208284031215611d6f575f80fd5b5035919050565b5f5b83811015611d90578181015183820152602001611d78565b50505f910152565b5f8151808452611daf816020860160208601611d76565b601f01601f19169290920160200192915050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611e1e5784516001600160a01b03168252938301936001929092019190830190611df5565b509695505050505050565b60018060a01b038716815285602082015260c060408201525f611e4f60c0830187611d98565b8281036060840152611e618187611d98565b90508281036080840152611e758186611dc3565b905082810360a0840152611e898185611dc3565b9998505050505050505050565b5f805f8060808587031215611ea9575f80fd5b843561ffff81168114611eba575f80fd5b935060208501356001600160401b0381168114611ed5575f80fd5b93969395505050506040820135916060013590565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715611f2057611f20611eea565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611f4e57611f4e611eea565b604052919050565b5f82601f830112611f65575f80fd5b81356001600160401b03811115611f7e57611f7e611eea565b611f91601f8201601f1916602001611f26565b818152846020838601011115611fa5575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215611fd1575f80fd5b81356001600160401b03811115611fe6575f80fd5b611b3d84828501611f56565b60018060a01b038716815285602082015260c060408201525f61201860c0830187611d98565b828103606084015261202a8187611d98565b9150508360808301526001600160401b03831660a0830152979650505050505050565b5f6040828403121561205d575f80fd5b612065611efe565b905061207082611d2e565b81526020808301356001600160401b038082111561208c575f80fd5b818501915085601f83011261209f575f80fd5b8135818111156120b1576120b1611eea565b8060051b91506120c2848301611f26565b81815291830184019184810190888411156120db575f80fd5b938501935b8385101561211057843592506001600160a01b0383168314612100575f80fd5b82825293850193908501906120e0565b808688015250505050505092915050565b5f805f805f60a08688031215612135575f80fd5b8535945060208601356001600160401b0380821115612152575f80fd5b61215e89838a01611f56565b95506040880135915080821115612173575f80fd5b61217f89838a01611f56565b94506060880135915080821115612194575f80fd5b6121a089838a0161204d565b935060808801359150808211156121b5575f80fd5b506121c28882890161204d565b9150509295509295909350565b602080825282518282018190525f9190848201906040850190845b81811015612206578351835292840192918401916001016121ea565b50909695505050505050565b5f60208284031215612222575f80fd5b5051919050565b600181811c9082168061223d57607f821691505b60208210810361225b57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561087e5761087e612261565b5f8251612299818460208701611d76565b9190910192915050565b61ffff8181168382160190808211156122be576122be612261565b5092915050565b5f602082840312156122d5575f80fd5b81518015158114611908575f80fd5b601f82111561232857805f5260205f20601f840160051c810160208510156123095750805b601f840160051c820191505b8181101561090f575f8155600101612315565b505050565b81516001600160401b0381111561234657612346611eea565b61235a816123548454612229565b846122e4565b602080601f83116001811461238d575f84156123765750858301515b5f19600386901b1c1916600185901b1785556123e4565b5f85815260208120601f198616915b828110156123bb5788860151825594840194600190910190840161239c565b50858210156123d857878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b634e487b7160e01b5f52603260045260245ffd5b5f61ffff80831681810361241657612416612261565b6001019392505050565b8181038181111561087e5761087e612261565b60a081525f61244560a0830188611d98565b82810360208401526124578188611d98565b9050828103604084015261246b8187611dc3565b9050828103606084015261247f8186611dc3565b9150506001600160401b03831660808301529695505050505050565b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220c794d015b2b0b2cbeaaa174a51d63829f5db7508c55ae56dfadb809869f569db64736f6c63430008190033",
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
// Solidity: function endAuction() returns(bytes32[])
func (_SlotAuctionManager *SlotAuctionManagerTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlotAuctionManager.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns(bytes32[])
func (_SlotAuctionManager *SlotAuctionManagerSession) EndAuction() (*types.Transaction, error) {
	return _SlotAuctionManager.Contract.EndAuction(&_SlotAuctionManager.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns(bytes32[])
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
