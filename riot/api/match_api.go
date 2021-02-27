package api

import (
	"fmt"
	"github.com/Jepzter/goleague/riot"
)

type MatchAPI struct {
	client *riot.RiotHttpClient
}

type MatchList struct {
	StartIndex, TotalGames, EndIndex int
	Matches []MatchReference `json:"matches"`
}

type MatchReference struct {
	GameID, Timestamp int64
	Role, PlatformID, Lane string
	Queue, Champion, Season int
}

type Filters struct {
	EndIndex, BeginIndex int
}

func NewMatchAPI(client *riot.RiotHttpClient) *MatchAPI {
	return &MatchAPI{ client: client }
}

func (api *MatchAPI) ListMatches(accountID string, filters Filters) (MatchList, error) {
	query := make(map[string]interface{})
	query["beginIndex"] = filters.BeginIndex
	query["endIndex"] = filters.EndIndex

	var matchList MatchList
	err := api.client.GET(fmt.Sprintf("match/v4/matchlists/by-account/%s", accountID), query, &matchList)
	if err != nil {
		return MatchList{}, err
	}
	return matchList, nil
}
