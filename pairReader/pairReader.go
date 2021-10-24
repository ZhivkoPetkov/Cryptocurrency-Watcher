package pairReader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CryptoPairs() []Pair {
	file, err := os.Open(`../pairs.txt`)
	if err != nil {
		fmt.Println(`File with pairs cannot be open`)
	}

	var krakenPairs []Pair

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var pairs []string
		var row string
		row = scanner.Text()
		pairs = strings.Split(row, ",")
		krakenPairs = append(krakenPairs, Pair{Name: pairs[2], ShortSymbol: pairs[0], LongSymbol: pairs[1]})
	}
	file.Close()
	return krakenPairs
}

type Pair struct {
	Name        string
	ShortSymbol string
	LongSymbol  string
}
