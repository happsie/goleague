package main

import (
	"fmt"
	goleague "goleague"
	"goleague/riot/config"

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

	summonerByName, err := client.Summoner().GetSummonerByName("Robert Chase")
	if err != nil {
		logrus.Fatal(err)
	}
	summonerByAccountID, err := client.Summoner().GetSummonerByAccountID(summonerByName.AccountID)
	if err != nil {
		logrus.Fatal(err)
	}
	summonerByPUUID, err := client.Summoner().GetSummonerByPUUID(summonerByAccountID.PUUID)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("%+v\n", summonerByName)
	fmt.Printf("%+v\n", summonerByAccountID)
	fmt.Printf("%+v\n", summonerByPUUID)
}
