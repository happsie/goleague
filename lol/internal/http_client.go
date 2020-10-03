package internal

import (
	"encoding/json"
	"fmt"
	"go-league/lol/config"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	TokenHeader = "X-Riot-Token"
	GET         = "GET"
)

type RiotHttpClient struct {
	http   http.Client
	config config.RiotConfig
}

// NewRiotHttpClient creates a http client that handles requests to the API
func NewRiotHTTPClient(client http.Client, config config.RiotConfig) *RiotHttpClient {
	return &RiotHttpClient{http: client, config: config}
}

// GET performs GET request to Riots API.
func (c *RiotHttpClient) GET(endpoint string, output interface{}) error {
	resp, err := c.do(GET, endpoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(output)
}

func (c *RiotHttpClient) do(method, endpoint string, body io.Reader) (*http.Response, error) {
	url := c.formatEndpoint(endpoint)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(TokenHeader, c.config.Token)
	logrus.Infof("%s: %s", method, url)
	resp, err := c.http.Do(req)
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
		retryAfter := resp.Header.Get("Retry-After")
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
		return nil, fmt.Errorf("error response: %d", resp.StatusCode)
	}

	return resp, nil
}

func (c *RiotHttpClient) retry(req *http.Request) (*http.Response, error) {
	for retryCount := 0; retryCount < c.config.Retries; retryCount++ {
		time.Sleep(time.Duration(c.config.RetryDelayMS) * time.Millisecond)
		resp, err := c.http.Do(req)
		if err == nil {
			return resp, err
		}
	}
	return nil, fmt.Errorf("retry failed")
}

func (c *RiotHttpClient) formatEndpoint(endpoint string) string {
	if strings.HasPrefix(endpoint, "/") {
		return fmt.Sprintf("%s://%s.%s%s", c.config.Schema, c.config.Region, c.config.URL, endpoint)
	}
	return fmt.Sprintf("%s://%s.%s/%s", c.config.Schema, c.config.Region, c.config.URL, endpoint)
}
