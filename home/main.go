package main

import (
	"bufio"
	"fmt"
	"krakenApi"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {

	var krakenPairs = ReadPairsFromFile()

	tick := time.Tick(5 * time.Second)
	for range tick {
		var result []string
		for key, value := range krakenPairs {
			result = append(result, krakenApi.GetKrakenData(key, value[0], value[1]))
		}
		sort.Strings(result)
		for _, crypto := range result {
			if strings.Contains(crypto, "-") {
				fmt.Println(color.RedString(crypto))
			} else {
				fmt.Println(color.GreenString(crypto))
			}
		}

	}
}

func ReadPairsFromFile() map[string][]string {
	file, err := os.Open(`../pairs.txt`)
	if err != nil {
		fmt.Println(`File with pairs cannot be open`)
	}

	krakenPairs := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var pairs []string
		var row string
		row = scanner.Text()
		pairs = strings.Split(row, ",")
		krakenPairs[pairs[2]] = []string{pairs[0], pairs[1]}
	}
	file.Close()
	return krakenPairs
}
