package main

import (
	"fmt"
	"github.com/Jepzter/goleague"
	"github.com/Jepzter/goleague/riot/config"
	"github.com/sirupsen/logrus"
)

func main() {
	client := goleague.NewAPIClient(config.RiotConfig{
		Token:        "API_TOKEN",
		URL:          "https://eun1.api.riotgames.com",
		Region:       "eun1",
		RetryDelayMS: 5000,
		Retries:      3,
	})

	rankedInfo, err := client.League().GetRankedInfo("test123")
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("%+v", rankedInfo)
}
