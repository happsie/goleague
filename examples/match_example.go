package main

import (
	"encoding/json"
	"fmt"

	"github.com/Jepzter/goleague"
	"github.com/Jepzter/goleague/riot/api"
	"github.com/Jepzter/goleague/riot/config"
	"github.com/sirupsen/logrus"
)

func main() {
	client := goleague.NewAPIClient(config.RiotConfig{
		Token:        "API_TOKEN",
		URL:          "https://euw1.api.riotgames.com",
		Region:       "euw1",
		RetryDelayMS: 5000,
		Retries:      3,
	})

	summoner, err := client.Summoner().GetSummonerByName("Matt Donovan")
	if err != nil {
		logrus.Fatal(err)
	}

	matchList, err := client.Match().ListMatches(summoner.AccountID, api.Filters{
		EndIndex: 1,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	match, err := client.Match().Match(matchList.Matches[0].GameID)
	b, _ := json.Marshal(match)
	fmt.Printf("%v", string(b))
}
