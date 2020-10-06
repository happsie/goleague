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
	MiniSeries struct {
		Losses int32
		Progress string
		Target int32
		Wins int32
	}
}

func NewLeagueAPI(client *riot.RiotHttpClient) *LeagueAPI {
	return &LeagueAPI{client:client}
}

func (api *LeagueAPI) GetRankedInfo(encryptedSummonerID string) (*[]Ranked, error) {
	rankedInfo := &[]Ranked{}
	err := api.client.GET(fmt.Sprintf("league/v4/entries/by-summoner/%s", encryptedSummonerID), rankedInfo)
	if err != nil {
		return nil, err
	}
	return rankedInfo, nil
}
