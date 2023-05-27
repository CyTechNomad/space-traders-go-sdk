package systems

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type systemsClient struct {
	httpClient *http.Client
}

type systemsClientOpts func(*systemsClient)

var (
	defaultOpts = []systemsClientOpts{
		WithHTTPClient(http.DefaultClient),
	}
)

func WithHTTPClient(httpClient *http.Client) systemsClientOpts {
	return func(c *systemsClient) {
		c.httpClient = httpClient
	}
}

func NewSystems(opts ...systemsClientOpts) *systemsClient {
	c := &systemsClient{}

	opts = append(defaultOpts, opts...)

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *systemsClient) ListSystems(ctx context.Context, req *ListSystemsRequest) (*ListSystemsResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/systems", nil)
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

	resp := &ListSystemsResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *systemsClient) GetSystem(ctx context.Context, req *GetSystemRequest) (*GetSystemResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/systems/"+req.SystemID, nil)
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

	resp := &GetSystemResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *systemsClient) ListWaypoints(ctx context.Context, req *ListWaypointsRequest) (*ListWaypointsResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/systems/"+req.SystemID+"/waypoints", nil)
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

	resp := &ListWaypointsResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *systemsClient) GetWaypoint(ctx context.Context, req *GetWaypointRequest) (*GetWaypointResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/systems/"+req.SystemID+"/waypoints/"+req.WaypointID, nil)
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

	resp := &GetWaypointResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *systemsClient) GetMarket(ctx context.Context, req *GetMarketRequest) (*GetMarketResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/systems/"+req.SystemID+"/waypoints/"+req.WaypointID+"/market", nil)
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

	resp := &GetMarketResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
