package utils

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/libevm/log"

	. "github.com/onsi/gomega"
)

const (
	DEFAULT_API_PORT = 8080
	SIG_AGG_API_PATH = "/aggregate-signatures"
)

// This is a wrapper around a signature aggregator binary instead of importing the package directly
// to avoid cyclic dependencies
type SignatureAggregator struct {
	cmd        *exec.Cmd
	cancelFunc context.CancelFunc
}

type SignatureAggregatorConfig struct {
	PChainAPI       ApiConfig `json:"p-chain-api"`
	InfoAPI         ApiConfig `json:"info-api"`
	SubnetIDs       []string  `json:"tracked-subnet-ids"`
	APIPort         int       `json:"api-port"`
	AllowPrivateIPs bool      `json:"allow-private-ips"`
}

type ApiConfig struct {
	BaseURL     string            `json:"base-url"`
	QueryParams map[string]string `json:"query-parameters"`
	HTTPHeaders map[string]string `json:"http-headers"`
}

type AggregateSignaturesRequest struct {
	Message          string `json:"message"`
	Justification    string `json:"justification,omitempty"`
	SigningSubnetID  string `json:"signing-subnet-id,omitempty"`
	QuorumPercentage uint64 `json:"quorum-percentage,omitempty"`
}

type SignatureAggregatorResponse struct {
	SignedMessage string `json:"signed-message"`
}

func (s *SignatureAggregator) Shutdown() {
	s.cancelFunc()
}

// Aggregator utils
func NewSignatureAggregator(baseURL string, subnetIDs []ids.ID) *SignatureAggregator {
	sigAggPath := os.Getenv("SIG_AGG_PATH")
	Expect(sigAggPath).ShouldNot(BeEmpty())
	subnetIDStrings := make([]string, 0, len(subnetIDs))
	for _, subnetID := range subnetIDs {
		subnetIDStrings = append(subnetIDStrings, subnetID.String())
	}
	cfg := SignatureAggregatorConfig{
		PChainAPI: ApiConfig{
			BaseURL: baseURL,
		},
		InfoAPI: ApiConfig{
			BaseURL: baseURL,
		},
		SubnetIDs:       subnetIDStrings,
		APIPort:         DEFAULT_API_PORT,
		AllowPrivateIPs: true,
	}
	// write config to a JSON file in /tmp directory
	configFile, err := os.CreateTemp("/tmp", "sig_agg_config_*.json")
	Expect(err).Should(BeNil())
	defer configFile.Close()

	encoder := json.NewEncoder(configFile)
	err = encoder.Encode(cfg)
	Expect(err).Should(BeNil())

	ctx, cancel := context.WithCancel(context.Background())

	cmd := exec.CommandContext(ctx, sigAggPath, "--config-file", configFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	Expect(err).Should(BeNil())
	go func() {
		err := cmd.Wait()
		// Context cancellation is the only expected way for the process to exit, otherwise log an error
		// Don't panic to allow for easier cleanup
		if !errors.Is(ctx.Err(), context.Canceled) {
			log.Error("Signature aggregator exited abnormally", err)
		}
	}()

	// TODO: when the signature aggregator health check endpoint is improved to not return
	// before ready to serve requests replace this sleep.
	time.Sleep(time.Second * 5)
	return &SignatureAggregator{
		cancelFunc: cancel,
		cmd:        cmd,
	}
}

func (s *SignatureAggregator) CreateSignedMessage(
	unsignedMessage *avalancheWarp.UnsignedMessage,
	justification []byte,
	inputSigningSubnet ids.ID,
	quorumPercentage uint64,
) (*avalancheWarp.Message, error) {
	var err error
	var signedMessage *avalancheWarp.Message
	for i := 0; i < 3; i++ {
		signedMessage, err = s.createSignedMessage(unsignedMessage, justification, inputSigningSubnet, quorumPercentage)
		if err == nil {
			return signedMessage, nil
		}
		time.Sleep(time.Second * 1)
	}
	return nil, err
}

func (s *SignatureAggregator) createSignedMessage(
	unsignedMessage *avalancheWarp.UnsignedMessage,
	justification []byte,
	inputSigningSubnet ids.ID,
	quorumPercentage uint64,
) (*avalancheWarp.Message, error) {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	requestURL := fmt.Sprintf("http://localhost:%d%s", DEFAULT_API_PORT, SIG_AGG_API_PATH)
	reqBody := AggregateSignaturesRequest{
		Message:          hex.EncodeToString(unsignedMessage.Bytes()),
		Justification:    hex.EncodeToString(justification),
		SigningSubnetID:  inputSigningSubnet.String(),
		QuorumPercentage: quorumPercentage,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(b)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status code 200, got %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response SignatureAggregatorResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	decodedMessage, err := hex.DecodeString(response.SignedMessage)
	if err != nil {
		return nil, err
	}

	signedMessage, err := avalancheWarp.ParseMessage(decodedMessage)
	if err != nil {
		return nil, err
	}

	return signedMessage, nil
}
