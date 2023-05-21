package v2

import (
	"spacetradersgo/v2/agents"
	"spacetradersgo/v2/contracts"
	"spacetradersgo/v2/factions"
)

type SpcaeTradersClient struct {
	Agents    agents.AgentsClient
	Factions  factions.FactionsClient
	Contracts contracts.ContractsClient
}

type spaceTraderClientOpts func(*SpcaeTradersClient)

var (
	defaultOpts = []spaceTraderClientOpts{
		WithAgentsClient(agents.NewAgents()),
		WithFactionsClient(factions.NewFactions()),
		WithContractsClient(contracts.NewContracts()),
	}
)

func WithAgentsClient(agentsClient agents.AgentsClient) spaceTraderClientOpts {
	return func(c *SpcaeTradersClient) {
		c.Agents = agentsClient
	}
}

func WithFactionsClient(factionsClient factions.FactionsClient) spaceTraderClientOpts {
	return func(c *SpcaeTradersClient) {
		c.Factions = factionsClient
	}
}

func WithContractsClient(contractsClient contracts.ContractsClient) spaceTraderClientOpts {
	return func(c *SpcaeTradersClient) {
		c.Contracts = contractsClient
	}
}

func NewSpaceTradersClient(opts ...spaceTraderClientOpts) *SpcaeTradersClient {
	c := &SpcaeTradersClient{}
	opts = append(defaultOpts, opts...)
	for _, opt := range opts {
		opt(c)
	}

	return c
}
