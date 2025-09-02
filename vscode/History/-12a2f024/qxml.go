package golearning
import "fmt"
type Bitcoin int
type Wallet struct {
	money Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.money
}

func (w *Wallet) Deposit(mo Bitcoin) {
	w.money += mo
	
}

func (b Bitcoin)String() string{
	return fmt.Sprintf("%d BTC",b);
}