package goleague

import (
	"goleague/riot"
	"goleague/riot/api"
	"goleague/riot/config"
	"net/http"
)

type Client struct {
	config *config.RiotConfig
	httpClient *riot.RiotHttpClient
}

// NewAPIClient Creates a new API client using RiotConfig.
func NewAPIClient(config config.RiotConfig) *Client {
	return &Client{
		config: &config,
		httpClient: riot.NewRiotHTTPClient(http.Client{}, config),
	}
}

// Summoner gives access to the summoner api (v4)
func (c *Client) Summoner() *api.SummonerAPI {
	return api.NewSummonerAPI(c.httpClient)
}

// Champion gives access to the champion api (v3)
func (c *Client) Champion() *api.ChampionAPI {
	return api.NewChampionAPI(c.httpClient)
}