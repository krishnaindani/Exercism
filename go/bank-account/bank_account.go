package account

import "sync"

//Account struct
type Account struct {
	sync.Mutex
	balance int64
	open    bool
}

//Open opens the account
func Open(initialDeposit int64) *Account {
	var a Account

	if initialDeposit < 0 {
		return nil
	}
	a = Account{
		balance: initialDeposit,
		open:    true,
	}
	return &a
}

//Close closes the account
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if !a.open {
		return a.balance, false
	}
	payout = a.balance
	a.balance = 0
	a.open = false
	return payout, true
}

//Balance returns the balance
func (a *Account) Balance() (balance int64, ok bool) {
	if a.open == false {
		a.balance = 0
		return 0, false
	}
	return a.balance, true
}

//Deposit deposits the amount
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if !a.open {
		return 0, false
	}
	value := a.balance + amount

	if value < 0 {
		return a.balance, false
	}
	a.balance = value
	return a.balance, true
}
