package utils

import (
	"context"
	"crypto/ecdsa"

	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	validatorsetsig "github.com/ava-labs/icm-contracts/abi-bindings/go/governance/ValidatorSetSig"
	"github.com/ava-labs/icm-contracts/tests/interfaces"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func DeployValidatorSetSig(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	contractL1 interfaces.L1TestInfo,
	validatorL1 interfaces.L1TestInfo,
) (common.Address, *validatorsetsig.ValidatorSetSig) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, contractL1.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, validatorSetSig, err := validatorsetsig.DeployValidatorSetSig(
		opts,
		contractL1.RPCClient,
		validatorL1.BlockchainID,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, contractL1, tx.Hash())

	return address, validatorSetSig
}

// Returns Receipt for the transaction unlike TeleporterRegistry version since this is a non-teleporter case
// and we don't want to add the ValidatorSetSig ABI to the L1Info
func ExecuteValidatorSetSigCallAndVerify(
	ctx context.Context,
	source interfaces.L1TestInfo,
	destination interfaces.L1TestInfo,
	validatorSetSigAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	signatureAggregator *SignatureAggregator,
	expectSuccess bool,
) *types.Receipt {
	signedWarpMsg := GetSignedMessage(source, destination, unsignedMessage, nil, signatureAggregator)
	log.Info("Got signed warp message", "messageID", signedWarpMsg.ID())

	signedPredicateTx := CreateExecuteCallPredicateTransaction(
		ctx,
		signedWarpMsg,
		validatorSetSigAddress,
		senderKey,
		destination,
	)

	// Wait for tx to be accepted and verify events emitted
	if expectSuccess {
		return SendTransactionAndWaitForSuccess(ctx, destination, signedPredicateTx)
	}
	return SendTransactionAndWaitForFailure(ctx, destination, signedPredicateTx)
}

func InitOffChainMessageChainConfigValidatorSetSig(
	networkID uint32,
	l1 interfaces.L1TestInfo,
	validatorSetSigAddress common.Address,
	validatorSetSigMessages []validatorsetsig.ValidatorSetSigMessage,
) ([]avalancheWarp.UnsignedMessage, string) {
	unsignedMessages := []avalancheWarp.UnsignedMessage{}
	for _, message := range validatorSetSigMessages {
		unsignedMessage := CreateOffChainValidatorSetSigMessage(networkID, l1, message)
		unsignedMessages = append(unsignedMessages, *unsignedMessage)
		log.Info("Adding validatorSetSig off-chain message to Warp chain config",
			"messageID", unsignedMessage.ID(),
			"blockchainID", l1.BlockchainID.String())
	}
	return unsignedMessages, GetChainConfigWithOffChainMessages(unsignedMessages)
}

// Creates an off-chain Warp message pointing to a function, contract and payload to be executed
// if the validator set signs this message
func CreateOffChainValidatorSetSigMessage(
	networkID uint32,
	l1 interfaces.L1TestInfo,
	message validatorsetsig.ValidatorSetSigMessage,
) *avalancheWarp.UnsignedMessage {
	sourceAddress := []byte{}
	payloadBytes, err := message.Pack()
	Expect(err).Should(BeNil())

	addressedPayload, err := payload.NewAddressedCall(sourceAddress, payloadBytes)
	Expect(err).Should(BeNil())

	unsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		l1.BlockchainID,
		addressedPayload.Bytes(),
	)
	Expect(err).Should(BeNil())

	return unsignedMessage
}
