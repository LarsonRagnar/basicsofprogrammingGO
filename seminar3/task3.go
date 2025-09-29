package main

import "fmt"

type BankAccount struct {
	owner   string
	balance float64
}

func (b *BankAccount) Deposit(addedBalace float64) {
	b.balance += addedBalace
}

func (b *BankAccount) Withdraw(withdrawBalance float64) {
	if withdrawBalance > b.balance {
		fmt.Println("Недостаточно средств")
	} else {
		b.balance -= withdrawBalance
	}
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}
func main() {
	account := &BankAccount{
		owner:   "Иван",
		balance: 1000,
	}

	account.Deposit(500)
	fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())
	//
	account.Withdraw(200)
	fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())
	//
	account.Withdraw(1500)
	fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())
}
