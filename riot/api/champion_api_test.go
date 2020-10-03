package api

import (
	"encoding/json"
	"github.com/Jepzter/goleague/riot"
	"github.com/Jepzter/goleague/riot/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChampionAPI_GetChampionRotation(t *testing.T) {
	srv := mockChampionServer()
	defer srv.Close()
	championAPI := NewChampionAPI(&riot.RiotHttpClient{
		Config: config.RiotConfig{
			Token:        "API_TOKEN",
			URL:          srv.URL,
			Region:       "euw1",
			RetryDelayMS: 10,
			Retries:      1,
		},
		HTTP: http.Client{},
	})

	rotation, err := championAPI.GetChampionRotation()
	if err != nil {
		t.Error(err)
	}
	if len(rotation.FreeChampionIDs) != 5 {
		t.Fail()
	}
	if len(rotation.FreeChampionIDsForNewPlayer) != 4 {
		t.Fail()
	}
	if rotation.MaxNewPlayerLevel != 10 {
		t.Fail()
	}
}

func mockChampionServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/lol/platform/v3/champion-rotations", championRotationMock)
	srv := httptest.NewServer(handler)
	return srv
}

func championRotationMock(w http.ResponseWriter, r *http.Request) {
	rotation := ChampionRotation{
		FreeChampionIDs:             []int32{1, 2, 3, 4, 5},
		FreeChampionIDsForNewPlayer: []int32{6, 7, 8, 9},
		MaxNewPlayerLevel:           10,
	}
	body, err := json.Marshal(rotation)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(500)
	}
}
