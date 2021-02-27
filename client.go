package goleague

import (
	"github.com/Jepzter/goleague/riot"
	"github.com/Jepzter/goleague/riot/api"
	"github.com/Jepzter/goleague/riot/config"
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

// League gives access to the league api (v4)
func (c *Client) League() *api.LeagueAPI {
	return api.NewLeagueAPI(c.httpClient)
}

func (c *Client) Match() *api.MatchAPI {
	return api.NewMatchAPI(c.httpClient)
}