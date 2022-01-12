package coinbase

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {

	wantCoin := "SHIB"
	wantCurrency := "EUR"

	response, _ := Get(fmt.Sprintf("%s-%s", wantCoin