package riot

import (
	"encoding/json"
	"fmt"
	"github.com/Jepzter/goleague/riot/config"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	TokenHeader = "X-Riot-Token"
	RetryAfter = "Retry-After"
	GET         = "GET"
)

type RiotHttpClient struct {
	HTTP   http.Client
	Config config.RiotConfig
}

// NewRiotHttpClient creates a HTTP client that handles requests to the API
func NewRiotHTTPClient(client http.Client, config config.RiotConfig) *RiotHttpClient {
	return &RiotHttpClient{HTTP: client, Config: config}
}

// GET performs GET request to Riots API.
func (c *RiotHttpClient) GET(endpoint string, output interface{}) error {
	resp, err := c.do(GET, endpoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(output)
	if err != nil && err.Error() == "EOF" {
		logrus.Errorf("GET: %s returned nothing in response body", endpoint)
		return err
	}
	return nil
}

func (c *RiotHttpClient) do(method, endpoint string, body io.Reader) (*http.Response, error) {
	url := c.formatEndpoint(endpoint)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(TokenHeader, c.Config.Token)
	logrus.Infof("%s: %s", method, url)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		logrus.Errorf("failed to make request: %v", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusServiceUnavailable {
		logrus.Errorf("Riot Games API is unavailable, retrying...")
		resp, err = c.retry(req)
		if err != nil {
			return nil, err
		}
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		retryAfter := resp.Header.Get(RetryAfter)
		seconds, err := strconv.Atoi(retryAfter)
		if err != nil {
			logrus.Errorf("failed to convert retry: %v", err)
			return nil, err
		}
		logrus.Debug("Reached rate-limit. Waiting %i seconds before continuing", seconds)
		time.Sleep(time.Duration(seconds) * time.Second)
		return c.do(method, endpoint, body)
	}

	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		logrus.Warnf("%d %s: %s", resp.StatusCode, method, url)
		return nil, fmt.Errorf("error response: %d", resp.StatusCode)
	}

	return resp, nil
}

func (c *RiotHttpClient) retry(req *http.Request) (*http.Response, error) {
	for retryCount := 0; retryCount < c.Config.Retries; retryCount++ {
		time.Sleep(time.Duration(c.Config.RetryDelayMS) * time.Millisecond)
		resp, err := c.HTTP.Do(req)
		if err == nil {
			return resp, err
		}
	}
	return nil, fmt.Errorf("retry failed")
}

func (c *RiotHttpClient) formatEndpoint(endpoint string) string {
	if strings.HasPrefix(endpoint, "/") {
		return fmt.Sprintf("%s/lol%s", c.Config.URL, endpoint)
	}
	return fmt.Sprintf("%s/lol/%s", c.Config.URL, endpoint)
}
