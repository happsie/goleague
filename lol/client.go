package lol

import (
	"go-league/lol/api"
	"go-league/lol/config"
	"go-league/lol/internal"
	"net/http"
)

type Client struct {
	*config.RiotConfig
	httpClient *internal.RiotHttpClient
}

// NewAPIClient Creates a new API client using RiotConfig.
func NewAPIClient(config config.RiotConfig) *Client {
	return &Client{
		RiotConfig: &config,
		httpClient: internal.NewRiotHTTPClient(http.Client{}, config),
	}
}

// Summoner gives access to the summoner api (v4)
func (c *Client) Summoner() *api.SummonerAPI {
	return api.NewSummonerAPI(c.httpClient)
}
