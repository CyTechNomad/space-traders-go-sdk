package factions

import (
	"context"
	"spacetradersgo/v2/utils"
)

type FactionsClient interface {
	// View the details of a faction by symbol.
	GetFaction(ctx context.Context, req *GetFactionRequest) (*GetFactionResponse, error)
	// View all factions.
	ListFactions(ctx context.Context, req *ListFactionsRequest) (*ListFactionsResponse, error)
}

type Faction struct {
	Symbol       string         `json:"symbol"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Headquarters string         `json:"headquarters"`
	Traits       []utils.Traits `json:"traits"`
	IsRecruiting bool           `json:"isRecruiting"`
}

type GetFactionRequest struct {
	Token         string
	FactionSymbol string
}
type GetFactionResponse struct {
	Faction Faction `json:"data"`
}

type ListFactionsRequest struct {
	Token string
	// How many entries to return per page
	// >= 1 <= 20
	NumPerPage int
	// What entry offset to request
	// >= 1
	Page int
}
type ListFactionsResponse struct {
	Factions []Faction  `json:"data"`
	Meta     utils.Meta `json:"meta"`
}
