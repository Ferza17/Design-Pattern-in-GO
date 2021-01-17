package main

import "fmt"

var overDraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposit ", amount, "\b , balance is now ", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overDraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew ", amount, "\b , balance is now ", b.balance)
		return true
	}
	return false
}

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(val bool)
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{
		account: account,
		amount:  amount,
		action:  action,
	}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeeded = true
	case Withdraw:
		b.succeeded = b.account.Withdraw(b.amount)
	}
}

func (b *BankAccountCommand) Succeeded() bool {
	return b.succeeded
}

func (b *BankAccountCommand) SetSucceeded(val bool) {
	b.succeeded = val
}

func (b *BankAccountCommand) Undo() {
	if !b.succeeded {
		return
	}

	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

// Composite Command
type CompositeBankAccountCommand struct {
	commands []Command
}

func (c CompositeBankAccountCommand) Call() {
	for _, command := range c.commands {
		command.Call()
	}
}

func (c CompositeBankAccountCommand) Undo() {
	for index := range c.commands {
		c.commands[len(c.commands)-index-1].Undo()
	}
}

func (c CompositeBankAccountCommand) Succeeded() bool {
	for _, command := range c.commands {
		if !command.Succeeded() {
			return false
		}
	}
	return true
}

func (c CompositeBankAccountCommand) SetSucceeded(val bool) {
	for _, command := range c.commands {
		command.SetSucceeded(val)
	}
}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from *BankAccount, to *BankAccount, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{
		from:   from,
		to:     to,
		amount: amount,
	}
	c.commands = append(c.commands, NewBankAccountCommand(from, Withdraw, amount))
	c.commands = append(c.commands, NewBankAccountCommand(to, Deposit, amount))

	return c
}

func (m *MoneyTransferCommand) Call() {
	ok := true
	for _, command := range m.commands {
		if ok {
			command.Call()
			ok = command.Succeeded()
		} else {
			command.SetSucceeded(false)
		}
	}
}

func main() {
	from := BankAccount{100}
	to := BankAccount{0}
	mtc := NewMoneyTransferCommand(&from, &to, 25)
	mtc.Call()
	fmt.Println(from, to)
	mtc.Undo()
	fmt.Println(from, to)

}
