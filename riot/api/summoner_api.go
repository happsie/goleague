package api

import (
	"fmt"
	"github.com/Jepzter/goleague/riot"
)

type SummonerAPI struct {
	client *riot.RiotHttpClient
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
func NewSummonerAPI(client *riot.RiotHttpClient) *SummonerAPI {
	return &SummonerAPI{client: client}
}

// GetSummonerByName fetches a summoner by their summoner name
// This api is mapped against (https://region.api.riotgames.com/lol/summoner/v4/summoners/by-name/:name)
func (api *SummonerAPI) GetSummonerByName(name string) (*Summoner, error) {
	return api.getSummoner(fmt.Sprintf("summoner/v4/summoners/by-name/%s", name))
}

// GetSummonerByAccountID fetches a summoner by their account id.
// This api is mapped against (https://region.api.riotgames.com/lol/summoner/v4/summoners/by-account/:id)
func (api *SummonerAPI) GetSummonerByAccountID(accountID string) (*Summoner, error) {
	return api.getSummoner(fmt.Sprintf("summoner/v4/summoners/by-account/%s", accountID))
}

// GetSummonerByPUUID fetches a summoner by their puuid.
// This api is mapped against (https://region.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/:puuid)
func (api *SummonerAPI) GetSummonerByPUUID(PUUID string) (*Summoner, error) {
	return api.getSummoner(fmt.Sprintf("summoner/v4/summoners/by-puuid/%s", PUUID))
}

func (api *SummonerAPI) getSummoner(endpoint string) (*Summoner, error) {
	summoner := &Summoner{}
	err := api.client.GET(endpoint, summoner)
	if err != nil {
		return nil, err
	}
	return summoner, nil
}