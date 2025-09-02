package golearning

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	want :=Bitcoin(10)
	if got != want {
		t.Errorf("%#v got %s instead of %s", wallet, got, want)
	}
}

