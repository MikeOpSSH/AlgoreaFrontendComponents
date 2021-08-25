
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