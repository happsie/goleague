package api

import (
	"encoding/json"
	"github.com/Jepzter/goleague/riot"
	"github.com/Jepzter/goleague/riot/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSummonerAPI_GetSummonerByName(t *testing.T) {
	srv := mockSummonerServer()
	defer srv.Close()
	summonerAPI := NewSummonerAPI(&riot.RiotHttpClient{
		Config: config.RiotConfig{
			Token:        "API_TOKEN",
			URL:          srv.URL,
			Region:       "euw1",
			RetryDelayMS: 10,
			Retries:      1,
		},
		HTTP: http.Client{},
	})

	summoner, err := summonerAPI.GetSummonerByName("test")
	if err != nil {
		t.Error(err)
	}
	if summoner.ID != "qwerty" {
		t.Fail()
	}
	if summoner.Name != "Test Summoner" {
		t.Fail()
	}
	if summoner.PUUID != "123abc" {
		t.Fail()
	}
	if summoner.AccountID != "abc123" {
		t.Fail()
	}
	if summoner.ProfileIconID != 1 {
		t.Fail()
	}
	if summoner.SummonerLevel != 1337 {
		t.Fail()
	}
}

func TestSummonerAPI_GetSummonerByAccountID(t *testing.T) {
	srv := mockSummonerServer()
	defer srv.Close()
	summonerAPI := NewSummonerAPI(&riot.RiotHttpClient{
		Config: config.RiotConfig{
			Token:        "API_TOKEN",
			URL:          srv.URL,
			Region:       "euw1",
			RetryDelayMS: 10,
			Retries:      1,
		},
		HTTP: http.Client{},
	})

	summoner, err := summonerAPI.GetSummonerByAccountID("test")
	if err != nil {
		t.Error(err)
	}
	if summoner.ID != "ytrewq" {
		t.Fail()
	}
	if summoner.Name != "Test Account" {
		t.Fail()
	}
	if summoner.PUUID != "abc123" {
		t.Fail()
	}
	if summoner.AccountID != "123abc" {
		t.Fail()
	}
	if summoner.ProfileIconID != 2 {
		t.Fail()
	}
	if summoner.SummonerLevel != 1327 {
		t.Fail()
	}
}

func TestSummonerAPI_GetSummonerByPUUID(t *testing.T) {
	srv := mockSummonerServer()
	defer srv.Close()
	summonerAPI := NewSummonerAPI(&riot.RiotHttpClient{
		Config: config.RiotConfig{
			Token:        "API_TOKEN",
			URL:          srv.URL,
			Region:       "euw1",
			RetryDelayMS: 10,
			Retries:      1,
		},
		HTTP: http.Client{},
	})

	summoner, err := summonerAPI.GetSummonerByPUUID("test")
	if err != nil {
		t.Error(err)
	}
	if summoner.ID != "abcdef" {
		t.Fail()
	}
	if summoner.Name != "Test puuid" {
		t.Fail()
	}
	if summoner.PUUID != "abcdef123" {
		t.Fail()
	}
	if summoner.AccountID != "456abc" {
		t.Fail()
	}
	if summoner.ProfileIconID != 4 {
		t.Fail()
	}
	if summoner.SummonerLevel != 1347 {
		t.Fail()
	}
}

func mockSummonerServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/lol/summoner/v4/summoners/by-name/test", summonerNameMock)
	handler.HandleFunc("/lol/summoner/v4/summoners/by-account/test", summonerAccountMock)
	handler.HandleFunc("/lol/summoner/v4/summoners/by-puuid/test", summonerPuuidMock)
	srv := httptest.NewServer(handler)
	return srv
}

func summonerNameMock(w http.ResponseWriter, r *http.Request) {
	summoner := Summoner{
		ID:            "qwerty",
		AccountID:     "abc123",
		PUUID:         "123abc",
		Name:          "Test Summoner",
		ProfileIconID: 1,
		SummonerLevel: 1337,
	}
	body, err := json.Marshal(summoner)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(500)
	}
}


func summonerAccountMock(w http.ResponseWriter, r *http.Request) {
	summoner := Summoner{
		ID:            "ytrewq",
		AccountID:     "123abc",
		PUUID:         "abc123",
		Name:          "Test Account",
		ProfileIconID: 2,
		SummonerLevel: 1327,
	}
	body, err := json.Marshal(summoner)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(500)
	}
}

func summonerPuuidMock(w http.ResponseWriter, r *http.Request) {
	summoner := Summoner{
		ID:            "abcdef",
		AccountID:     "456abc",
		PUUID:         "abcdef123",
		Name:          "Test puuid",
		ProfileIconID: 4,
		SummonerLevel: 1347,
	}
	body, err := json.Marshal(summoner)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(500)
	}
}