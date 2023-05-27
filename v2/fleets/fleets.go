package fleets

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func (c *fleetClient) GetShip(ctx context.Context, req *GetShipRequest) (*GetShipResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/my/ships/"+req.ShipID, nil)
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

	resp := &GetShipResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *fleetClient) GetShipCargo(ctx context.Context, req *GetShipCargoRequest) (*GetShipCargoResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/my/ships/"+req.ShipID+"/cargo", nil)
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

	resp := &GetShipCargoResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *fleetClient) GetShipNav(ctx context.Context, req *GetShipNavRequest) (*GetShipNavResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/my/ships/"+req.ShipID+"/nav", nil)
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

	resp := &GetShipNavResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *fleetClient) GetShipCooldown(ctx context.Context, req *GetShipCooldownRequest) (*GetShipCooldownResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, "https://api.spacetraders.io/v2/my/ships/"+req.ShipID+"/cooldown", nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+req.Token)

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	if httpResp.StatusCode == http.StatusNoContent {
		return &GetShipCooldownResponse{
			IsOnCooldown: false,
			Cooldown:     Cooldown{},
		}, nil
	}

	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	resp := &GetShipCooldownResponse{
		IsOnCooldown: true,
	}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *fleetClient) OrbitShip(ctx context.Context, req *OrbitShipRequest) (*OrbitShipResponse, error) {
	httpReq, err := http.NewRequest(http.MethodPost, "https://api.spacetraders.io/v2/my/ships/"+req.ShipID+"/orbit", nil)
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

	resp := &OrbitShipResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *fleetClient) DockShip(ctx context.Context, req *DockShipRequest) (*DockShipResponse, error) {
	httpReq, err := http.NewRequest(http.MethodPost, "https://api.spacetraders.io/v2/my/ships/"+req.ShipID+"/dock", nil)
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

	resp := &DockShipResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *fleetClient) CreateChart(ctx context.Context, req *CreateChartRequest) (*CreateChartResponse, error) {
	httpReq, err := http.NewRequest(http.MethodPost, "https://api.spacetraders.io/v2/my/ships/"+req.ShipID+"/chart", nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+req.Token)
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	resp := &CreateChartResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *fleetClient) CreateSurvey(ctx context.Context, req *CreateSurveyRequest) (*CreateSurveyResponse, error) {
	httpReq, err := http.NewRequest(http.MethodPost, "https://api.spacetraders.io/v2/my/ships/"+req.ShipID+"/survey", nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+req.Token)
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	resp := &CreateSurveyResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *fleetClient) NavigateShip(ctx context.Context, req *NavagateShipRequest) (*NavagateShipResponse, error) {
	reqBody, err := json.Marshal(map[string]string{"waypointSymbol": req.WaypointSymbol})
	if err != nil {
		return nil, err
	}
	fmt.Println(string(reqBody))
	httpReq, err := http.NewRequest(http.MethodPost, "https://api.spacetraders.io/v2/my/ships/"+req.ShipID+"/navigate", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+req.Token)
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	resp := &NavagateShipResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *fleetClient) ExtractResource(ctx context.Context, req *ExtractResourceRequest) (*ExtractResourceResponse, error) {
	reqBody, err := json.Marshal(req.Survey)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(reqBody))
	httpReq, err := http.NewRequest(http.MethodPost, "https://api.spacetraders.io/v2/my/ships/"+req.ShipID+"/extract", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+req.Token)
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	resp := &ExtractResourceResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
