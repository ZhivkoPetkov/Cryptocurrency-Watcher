package krakenApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

var clear map[string]func()

func GetKrakenData(cryptoName string, pair string, symbol string) string {
	var url string = fmt.Sprintf("https://api.kraken.com/0/public/Ticker?pair=%s", pair)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response KrakenResponse
	json.Unmarshal(responseData, &response)

	var currentValues = getPairValues(response.Result[symbol]["c"].([]interface{})[0].(string), response.Result[symbol]["o"].(string))
	return fmt.Sprintf(`%s: %.2f $, Difference: %.2f percents`, cryptoName, currentValues[0], currentValues[1])

}

type KrakenResponse struct {
	Error  []string                          `json:"error"`
	Result map[string]map[string]interface{} `json: "result"`
}

type PairMapper map[string]struct {
}

func getPairValues(lastTrade string, openingPrice string) [2]float64 {
	callClear()
	var result [2]float64
	parsedLastTrade, errTrade := strconv.ParseFloat(lastTrade, 2)
	parsedOpeningPrice, errOpening := strconv.ParseFloat(openingPrice, 2)

	if errTrade != nil || errOpening != nil {
		fmt.Println("Error in the values, when converting")
	}
	result[0] = parsedLastTrade
	result[1] = ((parsedLastTrade - parsedOpeningPrice) / parsedLastTrade) * 100
	return result
}

func callClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func init() {

	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
