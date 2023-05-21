package agents

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type agentsClient struct {
	httpClient *http.Client
}

type agentOpts func(*agentsClient)

var (
	defaultOpts = []agentOpts{
		WithHTTPClient(http.DefaultClient),
	}
)

func WithHTTPClient(httpClient *http.Client) agentOpts {
	return func(a *agentsClient) {
		a.httpClient = httpClient
	}
}

func NewAgents(opts ...agentOpts) *agentsClient {
	a := &agentsClient{}

	opts = append(defaultOpts, opts...)

	for _, opt := range opts {
		opt(a)
	}
	return a
}

func (a *agentsClient) NewAgent(ctx context.Context, req *NewAgentRequest) (*NewAgentResponse, error) {
	reqBody, err := json.Marshal(req)
	fmt.Println(string(reqBody))
	if err != nil {
		return nil, err
	}
	// TODO: add context to request
	// TODO: use constants for URL and content type
	httpResp, err := http.Post("https://api.spacetraders.io/v2/register", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body) // response body is []byte

	resp := &NewAgentResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *agentsClient) GetAgent(ctx context.Context, req *GetAgentRequest) (*GetAgentResponse, error) {
	// TODO: add context to request
	// TODO: use constants for URL and content type
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/my/agent", nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", req.Token))

	httpResp, err := a.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body) // response body is []byte
	if err != nil {
		return nil, err
	}

	resp := &GetAgentResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
