package contracts

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type contractsClient struct {
	httpClient *http.Client
}

type contractsCientOpt func(*contractsClient)

var (
	defaultOptions = []contractsCientOpt{
		WithHTTPClient(http.DefaultClient),
	}
)

func WithHTTPClient(httpClient *http.Client) contractsCientOpt {
	return func(c *contractsClient) {
		c.httpClient = httpClient
	}
}

func NewContracts(opts ...contractsCientOpt) *contractsClient {
	c := &contractsClient{
		httpClient: http.DefaultClient,
	}

	opts = append(defaultOptions, opts...)

	for _, o := range opts {
		o(c)
	}

	return c
}

// View all contracts.
func (c *contractsClient) ListContracts(ctx context.Context, req *ListContractsRequest) (*ListContractsResponse, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.spacetraders.io/v2/my/contracts", nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", req.Token))

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	resp := &ListContractsResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// View a specific contract.
func (c *contractsClient) GetContract(ctx context.Context, req *GetContractRequest) (*GetContractResponse, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://api.spacetraders.io/v2/my/contracts/%s", req.ContractID), nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", req.Token))

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	resp := &GetContractResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Accept a contract.
func (c *contractsClient) AcceptContract(ctx context.Context, req *AcceptContractRequest) (*AcceptContractResponse, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("https://api.spacetraders.io/v2/my/contracts/%s/accept", req.ContractID), nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", req.Token))

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	resp := &AcceptContractResponse{}
	err = json.NewDecoder(httpResp.Body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
