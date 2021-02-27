package main

import (
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
		EndIndex:   10,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("%+v", matchList)
}
