package golearning
import "fmt"
type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(mo Bitcoin) {
	w.balance += mo
	
}

func (b Bitcoin)String() string{
	return fmt.Sprintf("%d BTC",b);
}