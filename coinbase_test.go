package coinbase

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {

	wantCoin := "SHIB"
	wantCurrency := "EUR"

	respon