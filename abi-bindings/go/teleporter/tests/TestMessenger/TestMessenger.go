// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testmessenger

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

// TestMessengerMetaData contains all meta data concerning the TestMessenger contract.
var TestMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTeleporterVersion\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ReceiveMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"SendMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TELEPORTER_REGISTRY_APP_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getCurrentMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051611ff4380380611ff483398101604081905261002e9161062e565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff1615906001600160401b03165f811580156100775750825b90505f826001600160401b031660011480156100925750303b155b9050811580156100a0575080155b156100be5760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b031916600117855583156100ec57845460ff60401b1916680100000000000000001785555b6100f4610152565b6100ff888888610164565b831561014557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505061067e565b61015a610184565b6101626101d2565b565b61016c610184565b6101768382610200565b61017f82610226565b505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661016257604051631afcd79f60e31b815260040160405180910390fd5b6101da610184565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b610208610184565b610210610152565b61021861023a565b6102228282610242565b5050565b61022e610184565b610237816103d1565b50565b610162610184565b61024a610184565b6001600160a01b0382166102cb5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b5f7fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0090505f8390505f816001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610330573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103549190610667565b116103a95760405162461bcd60e51b815260206004820152603260248201525f80516020611fd4833981519152604482015271656c65706f7274657220726567697374727960701b60648201526084016102c2565b81546001600160a01b0319166001600160a01b0382161782556103cb8361040b565b50505050565b6103d9610184565b6001600160a01b03811661040257604051631e4fbdf760e01b81525f60048201526024016102c2565b610237816105a3565b7fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0080546040805163301fd1f560e21b815290515f926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa158015610472573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104969190610667565b6002830154909150818411156104f55760405162461bcd60e51b815260206004820152603160248201525f80516020611fd483398151915260448201527032b632b837b93a32b9103b32b939b4b7b760791b60648201526084016102c2565b80841161056a5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e0060648201526084016102c2565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d905f90a350505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b80516001600160a01b0381168114610629575f80fd5b919050565b5f805f60608486031215610640575f80fd5b61064984610613565b925061065760208501610613565b9150604084015190509250925092565b5f60208284031215610677575f80fd5b5051919050565b6119498061068b5f395ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c8063973142971161006e5780639731429714610159578063b33fead41461017c578063c868efaa1461019d578063d2cc7a70146101b0578063f2fde38b146101d7578063f63d09d7146101ea575f80fd5b80632b0d8f18146100b55780634511243e146100ca5780635eb99514146100dd578063715018a6146100f05780638da5cb5b146100f8578063909a6ac014610137575b5f80fd5b6100c86100c336600461130a565b6101fd565b005b6100c86100d836600461130a565b6102ff565b6100c86100eb366004611325565b6103ee565b6100c8610402565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b0390911681526020015b60405180910390f35b61014b5f805160206118f483398151915281565b60405190815260200161012e565b61016c61016736600461130a565b610415565b604051901515815260200161012e565b61018f61018a366004611325565b610435565b60405161012e929190611389565b6100c86101ab3660046113f9565b610507565b7fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d025461014b565b6100c86101e536600461130a565b6106e5565b61014b6101f8366004611451565b61071f565b5f805160206118f483398151915261021361087b565b6001600160a01b0382166102425760405162461bcd60e51b8152600401610239906114d1565b60405180910390fd5b61024c8183610883565b156102af5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610239565b6001600160a01b0382165f81815260018381016020526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a25050565b5f805160206118f483398151915261031561087b565b6001600160a01b03821661033b5760405162461bcd60e51b8152600401610239906114d1565b6103458183610883565b6103a35760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472794170703a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610239565b6001600160a01b0382165f818152600183016020526040808220805460ff19169055517f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c39190a25050565b6103f661087b565b6103ff816108a7565b50565b61040a610a3f565b6104135f610a9a565b565b5f5f805160206118f483398151915261042e8184610883565b9392505050565b5f81815260208181526040808320815180830190925280546001600160a01b0316825260018101805460609486949392908401916104729061151f565b80601f016020809104026020016040519081016040528092919081815260200182805461049e9061151f565b80156104e95780601f106104c0576101008083540402835291602001916104e9565b820191905f5260205f20905b8154815290600101906020018083116104cc57829003601f168201915b5050505050815250509050805f015181602001519250925050915091565b61050f610b0a565b5f5f805160206118f483398151915260028101548154919250906001600160a01b0316634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa15801561057a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061059e9190611557565b10156106055760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610239565b61060f8133610883565b156106755760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610239565b6106b5858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610b5492505050565b506106df60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050565b6106ed610a3f565b6001600160a01b03811661071657604051631e4fbdf760e01b81525f6004820152602401610239565b6103ff81610a9a565b5f610728610b0a565b5f851561073c576107398787610c09565b90505b876001600160a01b0316897fa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca898489898960405161077e959493929190611596565b60405180910390a36108446040518060c001604052808b81526020018a6001600160a01b0316815260200160405180604001604052808b6001600160a01b031681526020018581525081526020018781526020015f67ffffffffffffffff8111156107eb576107eb6115c3565b604051908082528060200260200182016040528015610814578160200160208202803683370190505b508152602001868660405160200161082d9291906115d7565b604051602081830303815290604052815250610c15565b91505061087060017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b979650505050505050565b610413610a3f565b6001600160a01b0381165f90815260018301602052604090205460ff165b92915050565b5f805160206118f483398151915280546040805163301fd1f560e21b815290515f926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa1580156108fb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061091f9190611557565b6002830154909150818411156109915760405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610239565b808411610a065760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610239565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d905f90a350505050565b33610a717f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146104135760405163118cdaa760e01b8152336004820152602401610239565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901610b4e57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f81806020019051810190610b6991906115ea565b6040805180820182526001600160a01b03868116825260208083018581525f8a815291829052939020825181546001600160a01b03191692169190911781559151929350916001820190610bbd90826116d2565b50905050826001600160a01b0316847f1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa83604051610bfb9190611792565b60405180910390a350505050565b5f61042e833384610d30565b5f80610c1f610e93565b60408401516020015190915015610cc4576040830151516001600160a01b0316610ca15760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610239565b604083015160208101519051610cc4916001600160a01b03909116908390610f83565b604051630624488560e41b81526001600160a01b03821690636244885090610cf09086906004016117e7565b6020604051808303815f875af1158015610d0c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061042e9190611557565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015610d76573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d9a9190611557565b9050610db16001600160a01b03861685308661100a565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015610df5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e199190611557565b9050818111610e7f5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610239565b610e898282611878565b9695505050505050565b5f805160206118f483398151915280546040805163d820e64f60e01b815290515f939284926001600160a01b039091169163d820e64f916004808201926020929091908290030181865afa158015610eed573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f11919061188b565b9050610f1d8282610883565b156108a15760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610239565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301525f919085169063dd62ed3e90604401602060405180830381865afa158015610fd0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ff49190611557565b90506106df848461100585856118a6565b611071565b6040516001600160a01b0384811660248301528381166044830152606482018390526106df9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506110fc565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790526110c28482611162565b6106df576040516001600160a01b0384811660248301525f60448301526110f691869182169063095ea7b39060640161103f565b6106df84825b5f6111106001600160a01b03841683611203565b905080515f1415801561113457508080602001905181019061113291906118b9565b155b1561115d57604051635274afe760e01b81526001600160a01b0384166004820152602401610239565b505050565b5f805f846001600160a01b03168460405161117d91906118d8565b5f604051808303815f865af19150503d805f81146111b6576040519150601f19603f3d011682016040523d82523d5f602084013e6111bb565b606091505b50915091508180156111e55750805115806111e55750808060200190518101906111e591906118b9565b80156111fa57505f856001600160a01b03163b115b95945050505050565b606061042e83835f845f80856001600160a01b0316848660405161122791906118d8565b5f6040518083038185875af1925050503d805f8114611261576040519150601f19603f3d011682016040523d82523d5f602084013e611266565b606091505b5091509150610e8986838360608261128657611281826112cd565b61042e565b815115801561129d57506001600160a01b0384163b155b156112c657604051639996b31560e01b81526001600160a01b0385166004820152602401610239565b508061042e565b8051156112dd5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b6001600160a01b03811681146103ff575f80fd5b5f6020828403121561131a575f80fd5b813561042e816112f6565b5f60208284031215611335575f80fd5b5035919050565b5f5b8381101561135657818101518382015260200161133e565b50505f910152565b5f815180845261137581602086016020860161133c565b601f01601f19169290920160200192915050565b6001600160a01b03831681526040602082018190525f906113ac9083018461135e565b949350505050565b5f8083601f8401126113c4575f80fd5b50813567ffffffffffffffff8111156113db575f80fd5b6020830191508360208285010111156113f2575f80fd5b9250929050565b5f805f806060858703121561140c575f80fd5b84359350602085013561141e816112f6565b9250604085013567ffffffffffffffff811115611439575f80fd5b611445878288016113b4565b95989497509550505050565b5f805f805f805f60c0888a031215611467575f80fd5b873596506020880135611479816112f6565b95506040880135611489816112f6565b9450606088013593506080880135925060a088013567ffffffffffffffff8111156114b2575f80fd5b6114be8a828b016113b4565b989b979a50959850939692959293505050565b6020808252602e908201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b600181811c9082168061153357607f821691505b60208210810361155157634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215611567575f80fd5b5051919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60018060a01b0386168152846020820152836040820152608060608201525f61087060808301848661156e565b634e487b7160e01b5f52604160045260245ffd5b602081525f6113ac60208301848661156e565b5f602082840312156115fa575f80fd5b815167ffffffffffffffff80821115611611575f80fd5b818401915084601f830112611624575f80fd5b815181811115611636576116366115c3565b604051601f8201601f19908116603f0116810190838211818310171561165e5761165e6115c3565b81604052828152876020848701011115611676575f80fd5b61087083602083016020880161133c565b601f82111561115d57805f5260205f20601f840160051c810160208510156116ac5750805b601f840160051c820191505b818110156116cb575f81556001016116b8565b5050505050565b815167ffffffffffffffff8111156116ec576116ec6115c3565b611700816116fa845461151f565b84611687565b602080601f831160018114611733575f841561171c5750858301515b5f19600386901b1c1916600185901b17855561178a565b5f85815260208120601f198616915b8281101561176157888601518255948401946001909101908401611742565b508582101561177e57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b602081525f61042e602083018461135e565b5f815180845260208085019450602084015f5b838110156117dc5781516001600160a01b0316875295820195908201906001016117b7565b509495945050505050565b60208152815160208201525f602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c08401526118476101008401826117a4565b905060a0840151601f198483030160e08501526111fa828261135e565b634e487b7160e01b5f52601160045260245ffd5b818103818111156108a1576108a1611864565b5f6020828403121561189b575f80fd5b815161042e816112f6565b808201808211156108a1576108a1611864565b5f602082840312156118c9575f80fd5b8151801515811461042e575f80fd5b5f82516118e981846020870161133c565b919091019291505056fede77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d00a264697066735822122062fd8850632ccfe83eea1845a044d69129346188301e664b16db39ae9a7db3cb64736f6c6343000819003354656c65706f7274657252656769737472794170703a20696e76616c69642054",
}

// TestMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use TestMessengerMetaData.ABI instead.
var TestMessengerABI = TestMessengerMetaData.ABI

// TestMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestMessengerMetaData.Bin instead.
var TestMessengerBin = TestMessengerMetaData.Bin

// DeployTestMessenger deploys a new Ethereum contract, binding an instance of TestMessenger to it.
func DeployTestMessenger(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, teleporterManager common.Address, minTeleporterVersion *big.Int) (common.Address, *types.Transaction, *TestMessenger, error) {
	parsed, err := TestMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestMessengerBin), backend, teleporterRegistryAddress, teleporterManager, minTeleporterVersion)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestMessenger{TestMessengerCaller: TestMessengerCaller{contract: contract}, TestMessengerTransactor: TestMessengerTransactor{contract: contract}, TestMessengerFilterer: TestMessengerFilterer{contract: contract}}, nil
}

// TestMessenger is an auto generated Go binding around an Ethereum contract.
type TestMessenger struct {
	TestMessengerCaller     // Read-only binding to the contract
	TestMessengerTransactor // Write-only binding to the contract
	TestMessengerFilterer   // Log filterer for contract events
}

// TestMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestMessengerSession struct {
	Contract     *TestMessenger    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestMessengerCallerSession struct {
	Contract *TestMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestMessengerTransactorSession struct {
	Contract     *TestMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestMessengerRaw struct {
	Contract *TestMessenger // Generic contract binding to access the raw methods on
}

// TestMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestMessengerCallerRaw struct {
	Contract *TestMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// TestMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestMessengerTransactorRaw struct {
	Contract *TestMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestMessenger creates a new instance of TestMessenger, bound to a specific deployed contract.
func NewTestMessenger(address common.Address, backend bind.ContractBackend) (*TestMessenger, error) {
	contract, err := bindTestMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestMessenger{TestMessengerCaller: TestMessengerCaller{contract: contract}, TestMessengerTransactor: TestMessengerTransactor{contract: contract}, TestMessengerFilterer: TestMessengerFilterer{contract: contract}}, nil
}

// NewTestMessengerCaller creates a new read-only instance of TestMessenger, bound to a specific deployed contract.
func NewTestMessengerCaller(address common.Address, caller bind.ContractCaller) (*TestMessengerCaller, error) {
	contract, err := bindTestMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestMessengerCaller{contract: contract}, nil
}

// NewTestMessengerTransactor creates a new write-only instance of TestMessenger, bound to a specific deployed contract.
func NewTestMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*TestMessengerTransactor, error) {
	contract, err := bindTestMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestMessengerTransactor{contract: contract}, nil
}

// NewTestMessengerFilterer creates a new log filterer instance of TestMessenger, bound to a specific deployed contract.
func NewTestMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*TestMessengerFilterer, error) {
	contract, err := bindTestMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestMessengerFilterer{contract: contract}, nil
}

// bindTestMessenger binds a generic wrapper to an already deployed contract.
func bindTestMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestMessenger *TestMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestMessenger.Contract.TestMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestMessenger *TestMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMessenger.Contract.TestMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestMessenger *TestMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestMessenger.Contract.TestMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestMessenger *TestMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestMessenger *TestMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestMessenger *TestMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestMessenger.Contract.contract.Transact(opts, method, params...)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_TestMessenger *TestMessengerCaller) TELEPORTERREGISTRYAPPSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestMessenger.contract.Call(opts, &out, "TELEPORTER_REGISTRY_APP_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_TestMessenger *TestMessengerSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _TestMessenger.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_TestMessenger.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_TestMessenger *TestMessengerCallerSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _TestMessenger.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_TestMessenger.CallOpts)
}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 sourceBlockchainID) view returns(address, string)
func (_TestMessenger *TestMessengerCaller) GetCurrentMessage(opts *bind.CallOpts, sourceBlockchainID [32]byte) (common.Address, string, error) {
	var out []interface{}
	err := _TestMessenger.contract.Call(opts, &out, "getCurrentMessage", sourceBlockchainID)

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 sourceBlockchainID) view returns(address, string)
func (_TestMessenger *TestMessengerSession) GetCurrentMessage(sourceBlockchainID [32]byte) (common.Address, string, error) {
	return _TestMessenger.Contract.GetCurrentMessage(&_TestMessenger.CallOpts, sourceBlockchainID)
}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 sourceBlockchainID) view returns(address, string)
func (_TestMessenger *TestMessengerCallerSession) GetCurrentMessage(sourceBlockchainID [32]byte) (common.Address, string, error) {
	return _TestMessenger.Contract.GetCurrentMessage(&_TestMessenger.CallOpts, sourceBlockchainID)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_TestMessenger *TestMessengerCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestMessenger.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_TestMessenger *TestMessengerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _TestMessenger.Contract.GetMinTeleporterVersion(&_TestMessenger.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_TestMessenger *TestMessengerCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _TestMessenger.Contract.GetMinTeleporterVersion(&_TestMessenger.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_TestMessenger *TestMessengerCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _TestMessenger.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_TestMessenger *TestMessengerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _TestMessenger.Contract.IsTeleporterAddressPaused(&_TestMessenger.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_TestMessenger *TestMessengerCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _TestMessenger.Contract.IsTeleporterAddressPaused(&_TestMessenger.CallOpts, teleporterAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestMessenger *TestMessengerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestMessenger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestMessenger *TestMessengerSession) Owner() (common.Address, error) {
	return _TestMessenger.Contract.Owner(&_TestMessenger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestMessenger *TestMessengerCallerSession) Owner() (common.Address, error) {
	return _TestMessenger.Contract.Owner(&_TestMessenger.CallOpts)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.PauseTeleporterAddress(&_TestMessenger.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.PauseTeleporterAddress(&_TestMessenger.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_TestMessenger *TestMessengerTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_TestMessenger *TestMessengerSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _TestMessenger.Contract.ReceiveTeleporterMessage(&_TestMessenger.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_TestMessenger *TestMessengerTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _TestMessenger.Contract.ReceiveTeleporterMessage(&_TestMessenger.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestMessenger *TestMessengerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestMessenger *TestMessengerSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestMessenger.Contract.RenounceOwnership(&_TestMessenger.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestMessenger *TestMessengerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestMessenger.Contract.RenounceOwnership(&_TestMessenger.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(bytes32)
func (_TestMessenger *TestMessengerTransactor) SendMessage(opts *bind.TransactOpts, destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "sendMessage", destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(bytes32)
func (_TestMessenger *TestMessengerSession) SendMessage(destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _TestMessenger.Contract.SendMessage(&_TestMessenger.TransactOpts, destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(bytes32)
func (_TestMessenger *TestMessengerTransactorSession) SendMessage(destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _TestMessenger.Contract.SendMessage(&_TestMessenger.TransactOpts, destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestMessenger *TestMessengerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestMessenger *TestMessengerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.TransferOwnership(&_TestMessenger.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestMessenger *TestMessengerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.TransferOwnership(&_TestMessenger.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.UnpauseTeleporterAddress(&_TestMessenger.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.UnpauseTeleporterAddress(&_TestMessenger.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_TestMessenger *TestMessengerTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_TestMessenger *TestMessengerSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _TestMessenger.Contract.UpdateMinTeleporterVersion(&_TestMessenger.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_TestMessenger *TestMessengerTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _TestMessenger.Contract.UpdateMinTeleporterVersion(&_TestMessenger.TransactOpts, version)
}

// TestMessengerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TestMessenger contract.
type TestMessengerInitializedIterator struct {
	Event *TestMessengerInitialized // Event containing the contract specifics and raw log

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
func (it *TestMessengerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerInitialized)
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
		it.Event = new(TestMessengerInitialized)
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
func (it *TestMessengerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerInitialized represents a Initialized event raised by the TestMessenger contract.
type TestMessengerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TestMessenger *TestMessengerFilterer) FilterInitialized(opts *bind.FilterOpts) (*TestMessengerInitializedIterator, error) {

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TestMessengerInitializedIterator{contract: _TestMessenger.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TestMessenger *TestMessengerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TestMessengerInitialized) (event.Subscription, error) {

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerInitialized)
				if err := _TestMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TestMessenger *TestMessengerFilterer) ParseInitialized(log types.Log) (*TestMessengerInitialized, error) {
	event := new(TestMessengerInitialized)
	if err := _TestMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the TestMessenger contract.
type TestMessengerMinTeleporterVersionUpdatedIterator struct {
	Event *TestMessengerMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *TestMessengerMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerMinTeleporterVersionUpdated)
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
		it.Event = new(TestMessengerMinTeleporterVersionUpdated)
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
func (it *TestMessengerMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the TestMessenger contract.
type TestMessengerMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_TestMessenger *TestMessengerFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*TestMessengerMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerMinTeleporterVersionUpdatedIterator{contract: _TestMessenger.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_TestMessenger *TestMessengerFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *TestMessengerMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerMinTeleporterVersionUpdated)
				if err := _TestMessenger.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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

// ParseMinTeleporterVersionUpdated is a log parse operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_TestMessenger *TestMessengerFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*TestMessengerMinTeleporterVersionUpdated, error) {
	event := new(TestMessengerMinTeleporterVersionUpdated)
	if err := _TestMessenger.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestMessenger contract.
type TestMessengerOwnershipTransferredIterator struct {
	Event *TestMessengerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestMessengerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerOwnershipTransferred)
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
		it.Event = new(TestMessengerOwnershipTransferred)
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
func (it *TestMessengerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerOwnershipTransferred represents a OwnershipTransferred event raised by the TestMessenger contract.
type TestMessengerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestMessenger *TestMessengerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestMessengerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerOwnershipTransferredIterator{contract: _TestMessenger.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestMessenger *TestMessengerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestMessengerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerOwnershipTransferred)
				if err := _TestMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestMessenger *TestMessengerFilterer) ParseOwnershipTransferred(log types.Log) (*TestMessengerOwnershipTransferred, error) {
	event := new(TestMessengerOwnershipTransferred)
	if err := _TestMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerReceiveMessageIterator is returned from FilterReceiveMessage and is used to iterate over the raw logs and unpacked data for ReceiveMessage events raised by the TestMessenger contract.
type TestMessengerReceiveMessageIterator struct {
	Event *TestMessengerReceiveMessage // Event containing the contract specifics and raw log

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
func (it *TestMessengerReceiveMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerReceiveMessage)
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
		it.Event = new(TestMessengerReceiveMessage)
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
func (it *TestMessengerReceiveMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerReceiveMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerReceiveMessage represents a ReceiveMessage event raised by the TestMessenger contract.
type TestMessengerReceiveMessage struct {
	SourceBlockchainID  [32]byte
	OriginSenderAddress common.Address
	Message             string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReceiveMessage is a free log retrieval operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, string message)
func (_TestMessenger *TestMessengerFilterer) FilterReceiveMessage(opts *bind.FilterOpts, sourceBlockchainID [][32]byte, originSenderAddress []common.Address) (*TestMessengerReceiveMessageIterator, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "ReceiveMessage", sourceBlockchainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerReceiveMessageIterator{contract: _TestMessenger.contract, event: "ReceiveMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveMessage is a free log subscription operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, string message)
func (_TestMessenger *TestMessengerFilterer) WatchReceiveMessage(opts *bind.WatchOpts, sink chan<- *TestMessengerReceiveMessage, sourceBlockchainID [][32]byte, originSenderAddress []common.Address) (event.Subscription, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "ReceiveMessage", sourceBlockchainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerReceiveMessage)
				if err := _TestMessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
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

// ParseReceiveMessage is a log parse operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, string message)
func (_TestMessenger *TestMessengerFilterer) ParseReceiveMessage(log types.Log) (*TestMessengerReceiveMessage, error) {
	event := new(TestMessengerReceiveMessage)
	if err := _TestMessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerSendMessageIterator is returned from FilterSendMessage and is used to iterate over the raw logs and unpacked data for SendMessage events raised by the TestMessenger contract.
type TestMessengerSendMessageIterator struct {
	Event *TestMessengerSendMessage // Event containing the contract specifics and raw log

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
func (it *TestMessengerSendMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerSendMessage)
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
		it.Event = new(TestMessengerSendMessage)
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
func (it *TestMessengerSendMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerSendMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerSendMessage represents a SendMessage event raised by the TestMessenger contract.
type TestMessengerSendMessage struct {
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	FeeTokenAddress         common.Address
	FeeAmount               *big.Int
	RequiredGasLimit        *big.Int
	Message                 string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSendMessage is a free log retrieval operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationBlockchainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_TestMessenger *TestMessengerFilterer) FilterSendMessage(opts *bind.FilterOpts, destinationBlockchainID [][32]byte, destinationAddress []common.Address) (*TestMessengerSendMessageIterator, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "SendMessage", destinationBlockchainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerSendMessageIterator{contract: _TestMessenger.contract, event: "SendMessage", logs: logs, sub: sub}, nil
}

// WatchSendMessage is a free log subscription operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationBlockchainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_TestMessenger *TestMessengerFilterer) WatchSendMessage(opts *bind.WatchOpts, sink chan<- *TestMessengerSendMessage, destinationBlockchainID [][32]byte, destinationAddress []common.Address) (event.Subscription, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "SendMessage", destinationBlockchainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerSendMessage)
				if err := _TestMessenger.contract.UnpackLog(event, "SendMessage", log); err != nil {
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

// ParseSendMessage is a log parse operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationBlockchainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_TestMessenger *TestMessengerFilterer) ParseSendMessage(log types.Log) (*TestMessengerSendMessage, error) {
	event := new(TestMessengerSendMessage)
	if err := _TestMessenger.contract.UnpackLog(event, "SendMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the TestMessenger contract.
type TestMessengerTeleporterAddressPausedIterator struct {
	Event *TestMessengerTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *TestMessengerTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerTeleporterAddressPaused)
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
		it.Event = new(TestMessengerTeleporterAddressPaused)
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
func (it *TestMessengerTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the TestMessenger contract.
type TestMessengerTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_TestMessenger *TestMessengerFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*TestMessengerTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerTeleporterAddressPausedIterator{contract: _TestMessenger.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_TestMessenger *TestMessengerFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *TestMessengerTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerTeleporterAddressPaused)
				if err := _TestMessenger.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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

// ParseTeleporterAddressPaused is a log parse operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_TestMessenger *TestMessengerFilterer) ParseTeleporterAddressPaused(log types.Log) (*TestMessengerTeleporterAddressPaused, error) {
	event := new(TestMessengerTeleporterAddressPaused)
	if err := _TestMessenger.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the TestMessenger contract.
type TestMessengerTeleporterAddressUnpausedIterator struct {
	Event *TestMessengerTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *TestMessengerTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerTeleporterAddressUnpaused)
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
		it.Event = new(TestMessengerTeleporterAddressUnpaused)
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
func (it *TestMessengerTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the TestMessenger contract.
type TestMessengerTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_TestMessenger *TestMessengerFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*TestMessengerTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerTeleporterAddressUnpausedIterator{contract: _TestMessenger.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_TestMessenger *TestMessengerFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *TestMessengerTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerTeleporterAddressUnpaused)
				if err := _TestMessenger.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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

// ParseTeleporterAddressUnpaused is a log parse operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_TestMessenger *TestMessengerFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*TestMessengerTeleporterAddressUnpaused, error) {
	event := new(TestMessengerTeleporterAddressUnpaused)
	if err := _TestMessenger.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
