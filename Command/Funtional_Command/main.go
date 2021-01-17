package main

import "fmt"

var overDraftLimit = -500

type BankAccount struct {
	balance int
}

func Deposit(b *BankAccount, amount int) {
	b.balance += amount
	fmt.Println("Deposit ", amount, "\b , balance is now ", b.balance)
}

func Withdraw(b *BankAccount, amount int) bool {
	if b.balance-amount >= overDraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew ", amount, "\b , balance is now ", b.balance)
		return true
	}
	return false
}

func main() {
	ba := &BankAccount{0}
	var commands []func()
	commands = append(commands, func() {
		Deposit(ba, 100)
	})
	commands = append(commands, func() {
		Withdraw(ba, 25)
	})

	for _, command := range commands {
		command()
	}

	fmt.Println(ba)
}
