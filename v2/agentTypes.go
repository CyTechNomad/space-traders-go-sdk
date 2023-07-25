package v2

import (
	"spacetradersgo/v2/agent"
	"spacetradersgo/v2/factions"
)

type NewAgentRequest struct {
	Faction string `json:"faction"`
	Symbol  string `json:"symbol"`
	Email   string `json:"email"`
}

type NewAgentResponse struct {
	Data struct {
		Agent    agent.Agent      `json:"agent"`
		Contract struct{}         `json:"contract"`
		Faction  factions.Faction `json:"faction"`
		Ship     struct{}         `json:"ship"`
		Token    string           `json:"token"`
	} `json:"data"`
}

type GetAgentRequest struct {
	Token string
}

type GetAgentResponse struct {
	Agent agent.Agent `json:"data"`
}
