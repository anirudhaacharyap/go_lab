package main

import "fmt"

type BankAccount struct {
	holder  string
	balance float64
}

func (b *BankAccount) Deposit(a float64)  { b.balance += a; fmt.Printf("Deposited: %.2f\n", a) }
func (b *BankAccount) Withdraw(a float64) {
	if a > b.balance {
		fmt.Println("Insufficient funds!")
		return
	}
	b.balance -= a
	fmt.Printf("Withdrawn: %.2f\n", a)
}
func (b *BankAccount) Balance() float64 { return b.balance }

func main() {
	acc := BankAccount{"John Doe", 1000}
	acc.Deposit(500)
	acc.Withdraw(300)
	fmt.Printf("Balance: %.2f\n", acc.Balance())
	acc.Withdraw(1500)
	fmt.Printf("Balance: %.2f\n", acc.Balance())
}
