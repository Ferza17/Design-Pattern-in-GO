package main

import "fmt"

type Momento struct {
	Balance int
}

type BankAccount struct {
	balance int
	// Save history of Struct
	changes []*Momento
	current int
}

func NewBankAccount(balance int) *BankAccount {
	b := &BankAccount{
		balance: balance,
	}

	b.changes = append(b.changes, &Momento{balance})

	return b
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("Balance = $", b.balance, " current = ", b.current)
}

func (b *BankAccount) Deposit(amount int) *Momento {
	b.balance += amount
	m := Momento{b.balance}
	b.changes = append(b.changes, &m)
	b.current++
	fmt.Println("Deposited ", amount, " Balance is now ", b.balance)
	return &m
}

func (b *BankAccount) Restore(m *Momento) {
	if m != nil {
		b.balance = m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount) Undo() *Momento {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}

	return nil
}

func (b *BankAccount) Redo() *Momento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}

	return nil
}

func main() {
	ba := NewBankAccount(100)
	ba.Deposit(50)
	ba.Deposit(25)

	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1:", ba)
	ba.Undo()
	fmt.Println("Undo 2:", ba)
	ba.Redo()
	fmt.Println("Redo :", ba)
}
