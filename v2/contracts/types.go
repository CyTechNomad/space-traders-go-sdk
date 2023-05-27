package contracts

import (
	"context"
	"spacetradersgo/v2/agents"
	"spacetradersgo/v2/fleets"
	"spacetradersgo/v2/utils"
	"time"
)

type ContractsClient interface {
	// View all contracts.
	ListContracts(ctx context.Context, req *ListContractsRequest) (*ListContractsResponse, error)
	// View a specific contract.
	GetContract(ctx context.Context, req *GetContractRequest) (*GetContractResponse, error)
	// Accept a contract.
	AcceptContract(ctx context.Context, req *AcceptContractRequest) (*AcceptContractResponse, error)
}

type Contract struct {
	ID               string    `json:"id"`
	FactionSymbol    string    `json:"factionSymbol"`
	Type             string    `json:"type"`
	Terms            terms     `json:"terms"`
	Accepted         bool      `json:"accepted"`
	Fulfilled        bool      `json:"fulfilled"`
	Expiration       time.Time `json:"expiration"`
	DeadlineToAccept time.Time `json:"deadlineToAccept"`
}

type terms struct {
	Deadline time.Time `json:"deadline"`
	Payment  payment   `json:"payment"`
	Deliver  []deliver `json:"deliver"`
}

type payment struct {
	OnAccepted  int `json:"onAccepted"`
	OnFulfilled int `json:"onFulfilled"`
}

type deliver struct {
	TradeSymbol       string `json:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol"`
	UnitsRequired     int    `json:"unitsRequired"`
	UnitsFulfilled    int    `json:"unitsFulfilled"`
}

type ListContractsRequest struct {
	Token      string
	NumPerPage int
	Page       int
}
type ListContractsResponse struct {
	Contracts []Contract `json:"data"`
	Meta      utils.Meta `json:"meta"`
}

type GetContractRequest struct {
	Token      string
	ContractID string
}
type GetContractResponse struct {
	Contract Contract `json:"data"`
}

type AcceptContractRequest struct {
	Token      string
	ContractID string
}
type AcceptContractResponse struct {
	Data struct {
		Agent    agents.Agent `json:"agent"`
		Contract Contract     `json:"contract"`
	} `json:"data"`
}

type DeliverContractRequest struct {
	Token string
	// The ID of the contract to deliver
	ContractID string
	// Symbol of the Ship to use to deliver the contract
	ShipSymbol string
	// Trade Symbol of the good to deliver
	TradeSymbol string
	// The number of units to deliver
	Units int
}

type DeliverContractResponse struct {
	Data struct {
		Contract Contract     `json:"contract"`
		Cargo    fleets.Cargo `json:"cargo"`
	} `json:"data"`
}

type FulfillContractRequest struct {
	Token      string
	ContractID string
}
type FulfillContractResponse struct {
	Data struct {
		Agent    agents.Agent `json:"agent"`
		Contract Contract     `json:"contract"`
	} `json:"data"`
}
