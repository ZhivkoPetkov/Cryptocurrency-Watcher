package main

import (
	"fmt"
	"krakenApi"
	"pairReader"
	"time"
)

func main() {
	var krakenPairs = pairReader.CryptoPairs()
	channels := make([]chan string, len(krakenPairs))

	for i := range krakenPairs {
		channels[i] = make(chan string)
	}

	tick := time.Tick(5 * time.Second)
	for range tick {
		krakenApi.NewList()

		for i, pair := range krakenPairs {
			go krakenApi.GetKrakenData(pair.Name, pair.ShortSymbol, pair.LongSymbol, channels[i])
		}

		for _, value := range channels {
			fmt.Println(<-value)
		}
	}
}
