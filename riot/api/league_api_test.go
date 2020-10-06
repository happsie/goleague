package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"goleague/riot"
	"goleague/riot/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLeagueAPI_GetRankedInfo(t *testing.T) {
	srv := mockLeagueServer()
	defer srv.Close()
	leagueAPI := NewLeagueAPI(&riot.RiotHttpClient{
		Config: config.RiotConfig{
			Token:        "API_TOKEN",
			URL:          srv.URL,
			Region:       "euw1",
			RetryDelayMS: 10,
			Retries:      1,
		},
		HTTP: http.Client{},
	})
	ranked, err := leagueAPI.GetRankedInfo("test123")
	if err != nil {
		t.Error(err)
	}

	if ranked[0].QueueType != "RANKED_SOLO_5x5" {
		t.Fail()
	}
	if ranked[0].Tier != "IRON" {
		t.Fail()
	}
	if ranked[0].Rank != "IV" {
		t.Fail()
	}
	if ranked[0].SummonerID != "Cool ID" {
		t.Fail()
	}
	if ranked[0].SummonerName != "Cool name" {
		t.Fail()
	}
	if ranked[0].LeaguePoints != 0 {
		t.Fail()
	}
	if ranked[0].Wins != 10 {
		t.Fail()
	}
	if ranked[0].Losses != 5 {
		t.Fail()
	}
}

func mockLeagueServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/lol/league/v4/entries/by-summoner/test123", leagueEntriesBySummoner)
	srv := httptest.NewServer(handler)
	return srv
}

func leagueEntriesBySummoner(w http.ResponseWriter, r *http.Request) {
	ranked := []Ranked{
		{
			LeagueID:     uuid.New(),
			QueueType:    "RANKED_SOLO_5x5",
			Tier:         "IRON",
			Rank:         "IV",
			SummonerID:   "Cool ID",
			SummonerName: "Cool name",
			LeaguePoints: 0,
			Wins:         10,
			Losses:       5,
			HotStreak:    false,
			Veteran:      false,
			FreshBlood:   false,
			Inactive:     false,
			MiniSeries:   MiniSeries{},
		},
	}
	body, err := json.Marshal(ranked)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(500)
	}
}