package coinbase

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {

	wantCoin := "SHIB"
	wantCurrency := "EUR"

	response, _ := Get(fmt.Sprintf("%s-%s", wantCoin, wantCurrency))
	if response.Base != wantCoin {
		t.Errorf("Translate() = %q, want %q", response.Base, wantCoin)
	}

	if response.Currency != wantCurrency {
		t.Errorf("TestGet() = %q, want %q", response.Currency, wantCoin)
	}

}

func TestGetWithDate(t *testing.T) {

	wantCoin := "SHIB"
	wantCurrency := "EUR"
	wantAmount := 0.0000190

	response, _ := GetWithDate(fmt.Sprintf("%s-%s", wantCoin, wantCurrency