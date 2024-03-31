package poinerr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		exp := Bitcoin(10.00)
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.00))

		assertBalance(t, wallet, exp)
	})
	t.Run("withdraw with funds", func(t *testing.T) {
		exp := Bitcoin(10.00)
		wallet := Wallet{balance: Bitcoin(20.00)}
		err := wallet.Withdraw(Bitcoin(10.00))

		assertNoError(t, err)
		assertBalance(t, wallet, exp)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20.00)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100.00))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertError(t testing.TB, err, exp error) {
	t.Helper()

	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err != exp {
		t.Errorf("expected %q actual %q", err, exp)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, exp Bitcoin) {
	t.Helper()
	actual := wallet.Balance()

	assert.Equal(t, exp, actual)
	if exp != actual {
		t.Errorf("expected %s and got %s", exp, actual)
	}
}
