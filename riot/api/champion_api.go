package api

import (
	"goleague/riot"
)

type ChampionAPI struct {
	client *riot.RiotHttpClient
}

type ChampionRotation struct {
	FreeChampionIDs []int32 `json:"freeChampionIds"`
	FreeChampionIDsForNewPlayer []int32 `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel int32 `json:"maxNewPlayerLevel"`
}

// NewChampionAPI creates a new champion api
func NewChampionAPI(client *riot.RiotHttpClient) *ChampionAPI {
	return &ChampionAPI{client: client}
}

// GetChampionRotation returns the current chamption rotation for both new and old players and new players max level.
// this api is mapped aganinst (https://region.api.riotgames.com/lol/platform/v3/champion-rotations)
func (api *ChampionAPI) GetChampionRotation() (*ChampionRotation, error) {
	champRotation := &ChampionRotation{}
	err := api.client.GET("platform/v3/champion-rotations", champRotation)
	if err != nil {
		return nil, err
	}
	return champRotation, nil
}
