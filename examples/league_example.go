package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goleague"
	"goleague/riot/config"
)

func main() {
	client := goleague.NewAPIClient(config.RiotConfig{
		Token:        "RGAPI-f18b65a5-b6c5-4f9a-83fc-a1ecdde52de7",
		URL:          "https://eun1.api.riotgames.com",
		Region:       "eun1",
		RetryDelayMS: 5000,
		Retries:      3,
	})

	rankedInfo, err := client.League().GetRankedInfo("MY-JdSQUP_QyQ3oxM669uZi2fBPVWHuNGwq8-lufdU3qNV0")
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("%+v", rankedInfo)
}
