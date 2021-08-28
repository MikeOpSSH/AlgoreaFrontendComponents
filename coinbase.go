
package coinbase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	ApiUrl = "https://api.coinbase.com/v2/prices"
	TYPE   = "spot"
)

type Response struct {
	Amount   float64
	Base     string
	Currency string
}

func buildUrl(coin string) string {
	return fmt.Sprintf("%s/%s/%s", ApiUrl, coin, TYPE)
}

func buildUrlWithDate(coin, date string) string {
	return fmt.Sprintf("%s/%s/%s?date=%s", ApiUrl, coin, TYPE, date)
}

func get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err