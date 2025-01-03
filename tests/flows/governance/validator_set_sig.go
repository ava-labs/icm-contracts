package governance

import (
	"context"
	"math/big"

	validatorsetsig "github.com/ava-labs/icm-contracts/abi-bindings/go/governance/ValidatorSetSig"
	exampleerc20 "github.com/ava-labs/icm-contracts/abi-bindings/go/mocks/ExampleERC20"
	localnetwork "github.com/ava-labs/icm-contracts/tests/network"
	"github.com/ava-labs/icm-contracts/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

func ValidatorSetSig(network *localnetwork.LocalNetwork) {
	// ************************************************************************************************
	// Setup
	// ************************************************************************************************
	// Deploy ValidatorSetSig expecting signatures from L1B instances to both L1s
	// Deploy exampleERC20 instance to both L1s
	// Construct ValidatorSetSig message with mock ERC20 as the target contract
	// Create off-chain ICM messages using the ValidatorSetSig message to be signed by the L1B
	// ************************************************************************************************
	// Test Case 1: validatorChain (L1B) != targetChain (L1A)
	// ************************************************************************************************
	// Send the off-chain message to L1A instance of ValidatorSetSig and confirm it is accepted.
	// Confirm the event is emitted
	// Retry the same message and confirm it fails
	// Send a new message with incremented nonce and confirm it is accepted.
	// ************************************************************************************************
	// Test Case 2: validatorChain (L1B) == targetChain (L1B)
	// ************************************************************************************************
	// Send a new message to L1B instance of ValidatorSetSig and confirm it is accepted

	// ************************************************************************************************
	// Setup
	// ************************************************************************************************
	L1A, L1B := network.GetTwoL1s()
	_, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy a ValidatorSetSigContract to L1A
	validatorSetSigContractAddress, validatorSetSig := utils.DeployValidatorSetSig(
		ctx,
		fundedKey,
		L1A,
		L1B,
	)
	// Deploy a ValidatorSetSigContract to L1B
	validatorSetSigContractAddress2, validatorSetSig2 := utils.DeployValidatorSetSig(
		ctx,
		fundedKey,
		L1B,
		L1B,
	)

	// Deploy a mock ERC20 contract to L1A
	exampleERC20ContractAddressA, exampleERC20A := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		L1A,
	)

	// Deploy a new example ERC20 contract this time to the same L1 as the validator.
	exampleERC20ContractAddressB, exampleERC20B := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		L1B,
	)

	erc20ABI, err := exampleerc20.ExampleERC20MetaData.GetAbi()
	Expect(err).Should(BeNil())

	// Confirm that the validatorContract has a balance of 0 on the example erc20 contracts on both L1s
	startingBalanceA, err := exampleERC20A.BalanceOf(
		&bind.CallOpts{}, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(startingBalanceA.Cmp(big.NewInt(0))).Should(BeZero())

	startingBalanceB, err := exampleERC20B.BalanceOf(
		&bind.CallOpts{}, validatorSetSigContractAddress2)
	Expect(err).Should(BeNil())
	Expect(startingBalanceB.Cmp(big.NewInt(0))).Should(BeZero())

	// Construct a ValidatorSetSig message with mock ERC20 as the target contract
	// and mint 100 tokens as the TxPayload
	callData, err := erc20ABI.Pack("mint", validatorSetSigContractAddress, big.NewInt(100))
	Expect(err).Should(BeNil())

	vssMessage1 := validatorsetsig.ValidatorSetSigMessage{
		ValidatorSetSigAddress: validatorSetSigContractAddress,
		TargetContractAddress:  exampleERC20ContractAddressA,
		TargetBlockchainID:     L1A.BlockchainID,
		Nonce:                  big.NewInt(0),
		Value:                  big.NewInt(0),
		Payload:                callData,
	}

	// Construct a second ValidatorSetSig message with mock ERC20 as the target contract
	// and mint 50 tokens as the TxPayload
	callData2, err := erc20ABI.Pack("mint", validatorSetSigContractAddress, big.NewInt(50))
	Expect(err).Should(BeNil())
	vssMessage2 := validatorsetsig.ValidatorSetSigMessage{
		ValidatorSetSigAddress: validatorSetSigContractAddress,
		TargetContractAddress:  exampleERC20ContractAddressA,
		TargetBlockchainID:     L1A.BlockchainID,
		Nonce:                  big.NewInt(1),
		Value:                  big.NewInt(0),
		Payload:                callData2,
	}

	// Create a message for the case where validatorSetSig contract and targetContract are on the same chain.
	// Construct a ValidatorSetSig message with mock ERC20 as the target contract
	// and mint 100 tokens as the TxPayload
	callData3, err := erc20ABI.Pack("mint", validatorSetSigContractAddress2, big.NewInt(100))
	Expect(err).Should(BeNil())

	vssMessage3 := validatorsetsig.ValidatorSetSigMessage{
		ValidatorSetSigAddress: validatorSetSigContractAddress2,
		TargetContractAddress:  exampleERC20ContractAddressB,
		TargetBlockchainID:     L1B.BlockchainID,
		Nonce:                  big.NewInt(0),
		Value:                  big.NewInt(0),
		Payload:                callData3,
	}

	// Create chain config file with off-chain validatorsetsig message
	networkID := network.GetNetworkID()
	offchainMessages, icmEnabledChainConfigWithMsg := utils.InitOffChainMessageChainConfigValidatorSetSig(
		networkID,
		L1B,
		validatorSetSigContractAddress,
		[]validatorsetsig.ValidatorSetSigMessage{vssMessage1, vssMessage2, vssMessage3},
	)

	// Create chain config with off-chain messages
	chainConfigs := make(utils.ChainConfigMap)
	chainConfigs.Add(L1B, icmEnabledChainConfigWithMsg)

	// Restart nodes with new chain config
	network.SetChainConfigs(chainConfigs)

	// ************************************************************************************************
	// Test Case 1: validatorChain (L1B) != targetChain (L1A)
	// ************************************************************************************************

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Execute the ValidatorSetSig executeCall and wait for acceptance
	receipt := utils.ExecuteValidatorSetSigCallAndVerify(
		ctx,
		L1B,
		L1A,
		validatorSetSigContractAddress,
		fundedKey,
		&offchainMessages[0],
		aggregator,
		true,
	)

	// Confirm that the Delivered event is emitted and that validatorSetSig contract has a balance of 100 of ERC20
	deliveredEvent, err := utils.GetEventFromLogs(receipt.Logs, validatorSetSig.ParseDelivered)
	Expect(err).Should(BeNil())
	Expect(deliveredEvent.TargetContractAddress).Should(Equal(exampleERC20ContractAddressA))
	Expect(deliveredEvent.Nonce.Cmp(big.NewInt(0))).Should(BeZero())

	endingBalance, err := exampleERC20A.BalanceOf(nil, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(endingBalance).Should(Equal(big.NewInt(100)))

	// Resend the same message again and it should fail due to nonce being consumed

	_ = utils.ExecuteValidatorSetSigCallAndVerify(
		ctx,
		L1B,
		L1A,
		validatorSetSigContractAddress,
		fundedKey,
		&offchainMessages[0],
		aggregator,
		false,
	)

	// Confirm that the validatorSetSig contract still has a balance of 100 of ERC20
	endingBalance, err = exampleERC20A.BalanceOf(nil, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(endingBalance).Should(Equal(big.NewInt(100)))

	// Send another valid transaction with the incremented nonce
	receipt2 := utils.ExecuteValidatorSetSigCallAndVerify(
		ctx,
		L1B,
		L1A,
		validatorSetSigContractAddress,
		fundedKey,
		&offchainMessages[1],
		aggregator,
		true,
	)

	// Confirm that the Delivered event is emitted and that validatorSetSig contract has a balance of 150 of ERC20
	deliveredEvent2, err := utils.GetEventFromLogs(receipt2.Logs, validatorSetSig.ParseDelivered)
	Expect(err).Should(BeNil())
	Expect(deliveredEvent2.TargetContractAddress).Should(Equal(exampleERC20ContractAddressA))
	Expect(deliveredEvent2.Nonce).Should(Equal(big.NewInt(1)))

	endingBalance, err = exampleERC20A.BalanceOf(nil, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(endingBalance).Should(Equal(big.NewInt(150)))

	startingBalanceB, err = exampleERC20B.BalanceOf(
		&bind.CallOpts{}, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(startingBalanceB.Cmp(big.NewInt(0))).Should(BeZero())

	// ************************************************************************************************
	// Test Case 2: validatorChain (L1B) == targetChain (L1B)
	// ************************************************************************************************

	// Send the third transaction where the validatorSetSig contract expects validator signatures
	// from the same chain that it is deployed on.
	receipt3 := utils.ExecuteValidatorSetSigCallAndVerify(
		ctx,
		L1B,
		L1B,
		validatorSetSigContractAddress2,
		fundedKey,
		&offchainMessages[2],
		aggregator,
		true,
	)

	// Confirm that the Delivered event is emitted and that validatorSetSig2 contract has a balance
	// of 100 of ERC20
	deliveredEvent3, err := utils.GetEventFromLogs(receipt3.Logs, validatorSetSig2.ParseDelivered)
	Expect(err).Should(BeNil())
	Expect(deliveredEvent3.TargetContractAddress).Should(Equal(exampleERC20ContractAddressB))
	Expect(deliveredEvent.Nonce.Cmp(big.NewInt(0))).Should(BeZero())

	endingBalance, err = exampleERC20B.BalanceOf(nil, validatorSetSigContractAddress2)
	Expect(err).Should(BeNil())
	Expect(endingBalance).Should(Equal(big.NewInt(100)))
}
