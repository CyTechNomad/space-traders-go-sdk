package systems

import (
	"context"
	"spacetradersgo/v2/factions"
	"spacetradersgo/v2/utils"
	"time"
)

type System struct {
	Symbol       string             `json:"symbol"`
	SectorSymbol string             `json:"sectorSymbol"`
	Type         string             `json:"type"`
	X            int                `json:"x"`
	Y            int                `json:"y"`
	Waypoints    []Waypoint         `json:"waypoints"`
	Factions     []factions.Faction `json:"factions"`
}

type Waypoint struct {
	Symbol       string           `json:"symbol"`
	Type         string           `json:"type"`
	SystemSymbol string           `json:"systemSymbol"`
	X            int              `json:"x"`
	Y            int              `json:"y"`
	Orbitals     []WaypointObital `json:"orbitals"`
	Faction      factions.Faction `json:"faction"`
	Traits       []utils.Traits   `json:"traits"`
	Chart        Chart            `json:"chart"`
}

type WaypointObital struct {
	Symbol string `json:"symbol"`
}

type Chart struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	SubmittedBy    string    `json:"submittedBy"`
	SubmittedOn    time.Time `json:"submittedOn"`
}

type Market struct {
	Symbol       string         `json:"symbol"`
	Exports      []exports      `json:"exports"`
	Imports      []imports      `json:"imports"`
	Exchange     []exchange     `json:"exchange"`
	Transactions []Transactions `json:"transactions"`
	TradeGoods   []tradeGoods   `json:"tradeGoods"`
}

type exports struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type imports struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type exchange struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Transactions struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	TradeSymbol    string    `json:"tradeSymbol"`
	Type           string    `json:"type"`
	Units          int       `json:"units"`
	PricePerUnit   int       `json:"pricePerUnit"`
	TotalPrice     int       `json:"totalPrice"`
	Timestamp      time.Time `json:"timestamp"`
}
type tradeGoods struct {
	Symbol        string `json:"symbol"`
	TradeVolume   int    `json:"tradeVolume"`
	Supply        string `json:"supply"`
	PurchasePrice int    `json:"purchasePrice"`
	SellPrice     int    `json:"sellPrice"`
}
type SystemsClient interface {
	// List all systems.
	ListSystems(ctx context.Context, req *ListSystemsRequest) (*ListSystemsResponse, error)
	// Get information about a specific system.
	GetSystem(ctx context.Context, req *GetSystemRequest) (*GetSystemResponse, error)
	// List all waypoints in a system.
	ListWaypoints(ctx context.Context, req *ListWaypointsRequest) (*ListWaypointsResponse, error)
	// Get information about a specific waypoint in a system.
	GetWaypoint(ctx context.Context, req *GetWaypointRequest) (*GetWaypointResponse, error)
	// Get market information about a specific waypoint in a system.
	GetMarket(ctx context.Context, req *GetMarketRequest) (*GetMarketResponse, error)
}

type ListSystemsRequest struct {
	Token      string
	NumPerPage int
	Page       int
}
type ListSystemsResponse struct {
	Systems []System   `json:"data"`
	Meta    utils.Meta `json:"meta"`
}

type GetSystemRequest struct {
	Token    string
	SystemID string
}
type GetSystemResponse struct {
	System System `json:"data"`
}

type ListWaypointsRequest struct {
	Token      string
	SystemID   string
	NumPerPage int
	Page       int
}
type ListWaypointsResponse struct {
	Waypoints []Waypoint `json:"data"`
	Meta      utils.Meta `json:"meta"`
}

type GetWaypointRequest struct {
	Token      string
	SystemID   string
	WaypointID string
}
type GetWaypointResponse struct {
	Waypoint Waypoint `json:"data"`
}

type GetMarketRequest struct {
	Token      string
	SystemID   string
	WaypointID string
}
type GetMarketResponse struct {
	Market Market `json:"data"`
}
