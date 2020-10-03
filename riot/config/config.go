package config

// RiotConfig is default config required for accessing the riotgames api.
type RiotConfig struct {
	Token        string `json:"token" yaml:"token"`
	URL          string `json:"url" yaml:"url"`
	Region       string `json:"region" yaml:"region"`
	RetryDelayMS int64  `json:"retryDelayMS" yaml:"retryDelayMS"`
	Retries      int    `json:"retries" yaml:"retries"`
}
