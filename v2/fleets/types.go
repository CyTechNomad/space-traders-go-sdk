package fleets

import (
	"context"
	"spacetradersgo/v2/agents"
	"spacetradersgo/v2/systems"
	"spacetradersgo/v2/utils"
	"time"
)

type Ship struct {
	ID           string       `json:"symbol"`
	Registration registration `json:"registration"`
	Nav          nav          `json:"nav"`
	Crew         crew         `json:"crew"`
	Frame        frame        `json:"frame"`
	Reactor      reactor      `json:"reactor"`
	Engine       engine       `json:"engine"`
	Modules      []module     `json:"modules"`
	Mounts       []mount      `json:"mounts"`
	Cargo        Cargo        `json:"cargo"`
	Fule         fule         `json:"fule"`
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

type Cooldown struct {
	ShipID          string    `json:"shipSymbol"`
	TotalSeconds    int       `json:"totalSeconds"`
	RemaningSeconds int       `json:"remainingSeconds"`
	Experation      time.Time `json:"expiraton"`
}

type Survey struct {
	Signature  string    `json:"signature"`
	Symbol     string    `json:"symbol"`
	Deposits   []Deposit `json:"deposits"`
	Expiration time.Time `json:"expiration"`
	Size       string    `json:"size"`
}

type Deposit struct {
	Symbol string `json:"symbol"`
}

type extraction struct {
	ShipSymbol string `json:"shipSymbol"`
	Yield      struct {
		Symbol string `json:"symbol"`
		Units  int    `json:"units"`
	} `json:"yield"`
}

type FleetsClient interface {
	// list ships
	ListShips(ctx context.Context, req *ListShipsRequest) (*ListShipsResponse, error)
	// get ship info
	GetShip(ctx context.Context, req *GetShipRequest) (*GetShipResponse, error)
	// get ship's cargo
	GetShipCargo(ctx context.Context, req *GetShipCargoRequest) (*GetShipCargoResponse, error)
	// get ship's nav
	GetShipNav(ctx context.Context, req *GetShipNavRequest) (*GetShipNavResponse, error)
	// get ship's cooldown
	GetShipCooldown(ctx context.Context, req *GetShipCooldownRequest) (*GetShipCooldownResponse, error)
	// Orbit Ship
	OrbitShip(ctx context.Context, req *OrbitShipRequest) (*OrbitShipResponse, error)
	// Dock Ship
	DockShip(ctx context.Context, req *DockShipRequest) (*DockShipResponse, error)
	// Creates a Chart for the current system
	CreateChart(ctx context.Context, req *CreateChartRequest) (*CreateChartResponse, error)
	// Create Survey for the current system
	CreateSurvey(ctx context.Context, req *CreateSurveyRequest) (*CreateSurveyResponse, error)
	// Navigate to a waypoint
	NavigateShip(ctx context.Context, req *NavagateShipRequest) (*NavagateShipResponse, error)
	// Extract a Resource
	ExtractResource(ctx context.Context, req *ExtractResourceRequest) (*ExtractResourceResponse, error)
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

type GetShipCargoRequest struct {
	Token  string
	ShipID string
}
type GetShipCargoResponse struct {
	Cargo Cargo `json:"data"`
}

type GetShipNavRequest struct {
	Token  string
	ShipID string
}
type GetShipNavResponse struct {
	Nav nav `json:"data"`
}

type GetShipCooldownRequest struct {
	Token  string
	ShipID string
}
type GetShipCooldownResponse struct {
	IsOnCooldown bool
	Cooldown     Cooldown `json:"data"`
}

type OrbitShipRequest struct {
	Token  string
	ShipID string
}
type OrbitShipResponse struct {
	Ship Ship `json:"data"`
}

type DockShipRequest struct {
	Token  string
	ShipID string
}
type DockShipResponse struct {
	Ship Ship `json:"data"`
}

type CreateChartRequest struct {
	Token  string
	ShipID string
}
type CreateChartResponse struct {
	Data struct {
		Chart    systems.Chart    `json:"chart"`
		Waypoint systems.Waypoint `json:"waypoint"`
	} `json:"data"`
}

type CreateSurveyRequest struct {
	Token  string
	ShipID string
}
type CreateSurveyResponse struct {
	Data struct {
		Cooldown Cooldown `json:"cooldown"`
		Survey   Survey   `json:"survey"`
	} `json:"data"`
}

type NavagateShipRequest struct {
	Token          string
	ShipID         string
	WaypointSymbol string
}
type NavagateShipResponse struct {
	Data struct {
		Fule fule `json:"fule"`
		Nav  nav  `json:"nav"`
	} `json:"data"`
}

type ExtractResourceRequest struct {
	Token  string
	ShipID string
	Survey Survey
}
type ExtractResourceResponse struct {
	Data struct {
		Cooldown   Cooldown   `json:"cooldown"`
		Extraction extraction `json:"extraction"`
		Cargo      Cargo      `json:"cargo"`
	} `json:"data"`
}

type SellCargoRequest struct {
	Token       string
	ShipID      string
	CargoSymbol string
	Units       int
}
type SellCargoResponse struct {
	Data struct {
		Agent       agents.Agent         `json:"agent"`
		Cargo       Cargo                `json:"cargo"`
		Transaction systems.Transactions `json:"transaction"`
	} `json:"data"`
}
