package main

import (
	"fmt"
	"krakenApi"
	"pairReader"
	"time"
)

func main() {
	var krakenPairs = pairReader.CryptoPairs()

	tick := time.Tick(5 * time.Second)
	for range tick {
		krakenApi.NewList()
		result := make(chan string)
		for key, value := range krakenPairs {
			go krakenApi.GetKrakenData(key, value[0], value[1], result)
			fmt.Println(<-result)
		}
		close(result)
	}
}
