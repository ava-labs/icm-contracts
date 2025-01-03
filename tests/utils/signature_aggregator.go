package utils

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"

	. "github.com/onsi/gomega"
)

const (
	DEFAULT_SIG_AGG_PATH = "~/.teleporter-deps/icm-services/signature-aggregator"
	DEFAULT_API_PORT     = 9090
	SIG_AGG_API_PATH     = "/aggregate-signatures"
)

// This is a wrapper around a signature aggregator binar instead of importing the package directly
// to avoid cyclic dependencies
type SignatureAggregator struct {
	cmd        *exec.Cmd
	cancelFunc context.CancelFunc
}

type SignatureAggregatorConfig struct {
	PChainAPI       ApiConfig `json:"p-chain-api"`
	InfoAPI         ApiConfig `json:"info-api"`
	L1IDs           []string  `json:"tracked-l1s"`
	ApiPort         int       `json:"api-port"`
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

func (s *SignatureAggregator) Shutdown() error {
	s.cancelFunc()
	return s.cmd.Wait()
}

// Aggregator utils
func NewSignatureAggregator(apiUri string, l1IDs []ids.ID) *SignatureAggregator {
	sigAggPath := os.Getenv("SIG_AGG_PATH")
	if sigAggPath == "" {
		sigAggPath = DEFAULT_SIG_AGG_PATH
	}
	l1IdStrings := make([]string, 0, len(l1IDs))
	for _, l1Id := range l1IDs {
		l1IdStrings = append(l1IdStrings, l1Id.String())
	}
	cfg := SignatureAggregatorConfig{
		PChainAPI: ApiConfig{
			BaseURL: apiUri,
		},
		InfoAPI: ApiConfig{
			BaseURL: apiUri,
		},
		L1IDs:           l1IdStrings,
		ApiPort:         DEFAULT_API_PORT,
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
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status code 200, got %d", res.StatusCode)
	}

	defer res.Body.Close()
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
