package pointers

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds error message. Also referenced in the test case
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin to wrap an int
type Bitcoin int

// Wallet cuz I said so
type Wallet struct {
	balance Bitcoin
}

// Deposit takes a deposit
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns a balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw takes money out of the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// String implements the Stringer interface on Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
