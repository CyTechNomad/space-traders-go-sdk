package fleets

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type fleetClient struct {
	httpClient *http.Client
}

type fleetClientOpts func(*fleetClient)

var (
	defaultOpts = []fleetClientOpts{
		WithHTTPClient(http.DefaultClient),
	}
)

func WithHTTPClient(httpClient *http.Client) fleetClientOpts {
	return func(c *fleetClient) {
		c.httpClient = httpClient
	}
}

func NewFleets(opts ...fleetClientOpts) *fleetClient {
	c := &fleetClient{}

	opts = append(defaultOpts, opts...)

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *fleetClient) ListShips(ctx context.Context, req *ListShipsRequest) (*ListShipsResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/my/ships", nil)
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

	resp := &ListShipsResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
