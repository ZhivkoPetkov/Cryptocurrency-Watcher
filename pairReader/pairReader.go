package pairReader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CryptoPairs() map[string][]string {
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
