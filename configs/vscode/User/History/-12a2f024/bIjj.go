package golearning
type Wallet struct{
money int
}

func (w Wallet)Balance() int{
	return w.money
}