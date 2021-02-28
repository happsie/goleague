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

type Match struct {
	GameID, GameDuration, GameCreation int64
	QueueID, SeasonID, MapID int
	GameType, PlatformID, GameVersion, GameMode string
	ParticipantIdentities []ParticipantIdentities
	Participants []Participant

}

type ParticipantIdentities struct {
	ParticipantID int
	Player Player
}

type Player struct {
	ProfileIcon int
	AccountID, MatchHistoryUri, CurrentAccountID, CurrentPlatformID, SummonerName, SummonerID, PlatformID string
}

type Participant struct {
	ParticipantID, ChampionID, TeamID, Spell1ID, Spell2ID int
	HighestAchievedSeasonTier string
	Stats ParticipantStats `json:"stats"`
}

type ParticipantStats struct {
	Item0, Item1, Item2, Item3, Item4, Item5, Item6, Deaths, Kills, Assists, GoldEarned, ChampLevel, ParticipantID int
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

func (api *MatchAPI) Match(matchID int64) (Match, error) {
	var match Match
	err := api.client.GET(fmt.Sprintf("match/v4/matches/%d", matchID), nil, &match)
	if err != nil {
		return Match{}, err
	}
	return match, nil
}
