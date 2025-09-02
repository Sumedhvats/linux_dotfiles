package golearning
import "testing"
func testWallet(t *testing.T){
	wallet:=Wallet{}
	wallet.Deposit(10)
	got:= wallet.Balance()
	want:= 10
	if got != want{
		t.Errorf("")
	}
}