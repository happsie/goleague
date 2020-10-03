package main

import (
	"fmt"
	"go-league/lol"
	"go-league/lol/config"

	"github.com/sirupsen/logrus"
)

func main() {
	client := lol.NewAPIClient(config.RiotConfig{
		Token:        "API_TOKEN",
		URL:          "api.riotgames.com/lol",
		Schema:       "https",
		Region:       "euw1",
		RetryDelayMS: 5000,
		Retries:      3,
	})

	summoner, err := client.Summoner().GetSummoner("Robert Chase")
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Printf("%+v", summoner)
}
