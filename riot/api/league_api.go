package api

import (
	"fmt"
	"github.com/google/uuid"
	"goleague/riot"
)

type LeagueAPI struct {
	client *riot.RiotHttpClient
}

type Ranked struct {
	LeagueID uuid.UUID
	QueueType string
	Tier string
	Rank string
	SummonerID string
	SummonerName string
	LeaguePoints int32
	Wins int32
	Losses int32
	HotStreak bool
	Veteran bool
	FreshBlood bool
	Inactive bool
	MiniSeries MiniSeries
}

type MiniSeries  struct {
	Losses int32
	Progress string
	Target int32
	Wins int32
}

// NewLeagueAPI creates a new league api
func NewLeagueAPI(client *riot.RiotHttpClient) *LeagueAPI {
	return &LeagueAPI{client:client}
}

// GetRankedInfo fetches ranked information from the league endpoint, comes as slice with different queue types
// Mapped against https://eun1.api.riotgames.com/lol/league/v4/entries/by-summoner/{encryptedSummonerId}
func (api *LeagueAPI) GetRankedInfo(encryptedSummonerID string) ([]Ranked, error) {
	var rankedInfo []Ranked
	err := api.client.GET(fmt.Sprintf("league/v4/entries/by-summoner/%s", encryptedSummonerID), &rankedInfo)
	if err != nil {
		return nil, err
	}
	return rankedInfo, nil
}
