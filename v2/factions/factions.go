package factions

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type factionsClient struct {
	httpClient *http.Client
}

type factionsClientOpts func(*factionsClient)

var (
	defaultOpts = []factionsClientOpts{
		WithHTTPClient(http.DefaultClient),
	}
)

func WithHTTPClient(httpClient *http.Client) factionsClientOpts {
	return func(c *factionsClient) {
		c.httpClient = httpClient
	}
}

func NewFactions(opts ...factionsClientOpts) *factionsClient {
	c := &factionsClient{}

	opts = append(defaultOpts, opts...)

	for _, opt := range opts {
		opt(c)
	}
	return c
}

// View the details of a faction by symbol.
func (c *factionsClient) GetFaction(ctx context.Context, req *GetFactionRequest) (*GetFactionResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/factions/"+req.FactionSymbol, nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+req.Token)

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	resp := &GetFactionResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// View all factions.
func (c *factionsClient) ListFactions(ctx context.Context, req *ListFactionsRequest) (*ListFactionsResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/factions", nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+req.Token)
	values := httpReq.URL.Query()
	if req.NumPerPage != 0 {
		values.Set("limit", strconv.Itoa(req.NumPerPage))
	}
	if req.Page != 0 {
		values.Set("page", strconv.Itoa(req.Page))
	}
	httpReq.URL.RawQuery = values.Encode()

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	resp := &ListFactionsResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
