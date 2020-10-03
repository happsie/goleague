package riot

import (
	"github.com/Jepzter/goleague/riot/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRiotHttpClient_GET(t *testing.T) {
	srv := mockServer()
	http := NewRiotHTTPClient(http.Client{},
		config.RiotConfig{
			Token:        "API_TOKEN",
			URL:          srv.URL,
			Region:       "euw1",
			RetryDelayMS: 10,
			Retries:      1,
		})


	err := http.GET("test", nil)
	if err != nil {
		t.Error(err)
	}
}

func mockServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/lol/test", testGETMock)
	srv := httptest.NewServer(handler)
	return srv
}

func testGETMock(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}