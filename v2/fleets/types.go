package fleets

import (
	"context"
	"spacetradersgo/v2/utils"
	"time"
)

type Ship struct {
	ID           string
	Registration registration
	Nav          nav
	Crew         crew
	Frame        frame
	Reactor      reactor
	Engine       engine
	Modules      []module
	Mounts       []mount
	Cargo        Cargo
	Fule         fule
}

type registration struct {
	Name          string `json:"name"`
	FactionSymbol string `json:"factionSymbol"`
	Role          string `json:"role"`
}

type nav struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Route          route  `json:"route"`
	Status         string `json:"status"`
	FlightMode     string `json:"flightMode"`
}

type route struct {
	Destination   destination `json:"destination"`
	Departure     departure   `json:"departure"`
	DepartureTime time.Time   `json:"departureTime"`
	Arrival       time.Time   `json:"arrivalTime"`
}

type destination struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type departure struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type crew struct {
	Current  int    `json:"current"`
	Required int    `json:"required"`
	Capacity int    `json:"capacity"`
	Rotation string `json:"rotation"`
	Morale   int    `json:"morale"`
	Wages    int    `json:"wages"`
}

type frame struct {
	Symbol         string       `json:"symbol"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Condition      int          `json:"condition"`
	ModuleSlots    int          `json:"moduleSlots"`
	MountingPoints int          `json:"mountingPoints"`
	FuleCapacity   int          `json:"fuleCapacity"`
	Requirements   requirements `json:"requirements"`
}

type requirements struct {
	Power int `json:"power"`
	Crew  int `json:"crew"`
	Slots int `json:"slots"`
}

type reactor struct {
	Symbol       string       `json:"symbol"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Condition    int          `json:"condition"`
	PowerOutput  int          `json:"powerOutput"`
	Requirements requirements `json:"requirements"`
}

type engine struct {
	Symbol       string       `json:"symbol"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Condition    int          `json:"condition"`
	Speed        int          `json:"speed"`
	Requirements requirements `json:"requirements"`
}

type module struct {
	Symbol       string       `json:"symbol"`
	Capacity     int          `json:"capacity"`
	Range        int          `json:"range"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Requirements requirements `json:"requirements"`
}

type mount struct {
	Symbol       string       `json:"symbol"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Strength     int          `json:"strength"`
	Deposits     []string     `json:"deposits"`
	Requirements requirements `json:"requirements"`
}

type Cargo struct {
	Capacity  int         `json:"capacity"`
	Units     int         `json:"units"`
	Inventory []cargoItem `json:"inventory"`
}

type cargoItem struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Units       int    `json:"units"`
}

type fule struct {
	Current  int `json:"current"`
	Capacity int `json:"capacity"`
	Consumed struct {
		Amount    int       `json:"amount"`
		TimeStamp time.Time `json:"timeStamp"`
	} `json:"consumed"`
}

type FleetsClient interface {
	// list ships
	ListShips(ctx context.Context, req *ListShipsRequest) (*ListShipsResponse, error)
}

type ListShipsRequest struct {
	Token      string
	NumPerPage int
	Page       int
}
type ListShipsResponse struct {
	Ships []Ship     `json:"data"`
	Meta  utils.Meta `json:"meta"`
}

type GetShipRequest struct {
	Token  string
	ShipID string
}
type GetShipResponse struct {
	Ship Ship `json:"data"`
}
