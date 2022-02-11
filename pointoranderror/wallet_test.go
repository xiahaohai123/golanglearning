package pointoranderror

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		fmt.Println("address of wallet in test is", &wallet)
		fmt.Println("address of balance in test is", &wallet.balance)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(10)
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Error("Wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf(" got %s want %s", got, want)
	}
}
