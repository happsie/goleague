package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	goleague "goleague"
	"goleague/riot/config"
)

func main() {
	client := goleague.NewAPIClient(config.RiotConfig{
		Token:        "API_TOKEN",
		URL:          "https://euw1.api.riotgames.com",
		Region:       "euw1",
		RetryDelayMS: 5000,
		Retries:      3,
	})

	championRotation, err := client.Champion().GetChampionRotation()
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("%+v\n", championRotation)
}
