package api

import (
	"fmt"
	"go-league/lol/internal"
)

type SummonerAPI struct {
	client *internal.RiotHttpClient
}

type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	PUUID         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int32  `json:"profileIconId"`
	SummonerLevel int32  `json:"summonerLevel"`
}

// NewSummonerAPI creates a new summoner api
func NewSummonerAPI(client *internal.RiotHttpClient) *SummonerAPI {
	return &SummonerAPI{client: client}
}

// GetSummoner fetches a summoner by their summoner name
func (api *SummonerAPI) GetSummoner(name string) (*Summoner, error) {
	summoner := &Summoner{}
	err := api.client.GET(fmt.Sprintf("summoner/v4/summoners/by-name/%s", name), summoner)
	return summoner, err
}
