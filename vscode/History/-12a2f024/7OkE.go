package golearning

type Wallet struct {
	money int
}

func (w Wallet) Balance() int {
	return w.money
}

func (w Wallet) Deposit(mo int) {
	w.money = mo

}
