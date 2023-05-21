package agents

import (
	"context"
)

type AgentsClient interface {
	// NewAgent creates a new agent.
	NewAgent(ctx context.Context, req *NewAgentRequest) (*NewAgentResponse, error)
	// GetAgent returns the agent associated with token.
	GetAgent(ctx context.Context, req *GetAgentRequest) (*GetAgentResponse, error)
}

type Agent struct {
	// AccountID is the unique identifier of the parent Account.
	// >= 1 characters
	AccountID string `json:"accountId"`
	// Symbol is the unique identifier of the agent.
	// >= 1 characters
	Symbol string `json:"symbol"`
	// Headquaters The headquarters of the agent.
	// >= 1 characters
	Headquarters string `json:"headquarters"`
	// Credits The number of credits the agent has available. Credits can be negative if funds have been overdrawn.
	Credits int64 `json:"credits"`
	// StartingFaction The faction the agent started with.
	// >= 1 characters
	StartingFaction string `json:"startingFaction"`
}

type NewAgentRequest struct {
	Faction string `json:"faction"`
	Symbol  string `json:"symbol"`
	Email   string `json:"email"`
}

type NewAgentResponse struct {
	Data struct {
		Agent    Agent    `json:"agent"`
		Contract struct{} `json:"contract"`
		Faction  struct{} `json:"faction"`
		Ship     struct{} `json:"ship"`
		Token    string   `json:"token"`
	} `json:"data"`
}

type GetAgentRequest struct {
	Token string
}

type GetAgentResponse struct {
	Agent Agent `json:"data"`
}
