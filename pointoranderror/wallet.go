package pointoranderror

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (bitcoin Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", bitcoin)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of wallet in Deposit is", &w)
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type Stringer interface {
	String() string
}
