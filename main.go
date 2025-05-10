package main

import "fmt"

type checkingAccount struct {
	assignee      string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (c *checkingAccount) cashOut(value float64) string {
	canCashOut := value <= c.balance
	if canCashOut {
		c.balance -= value
		return fmt.Sprintf("cash out successful! New balance: %.2f", c.balance)
	}
	return fmt.Sprintf("insufficient funds! Current balance: %.2f", c.balance)

}

func (c *checkingAccount) deposit(value float64) string {
	if value <= 0 {
		return "deposit amount must be greater than zero"
	}
	c.balance += value
	return fmt.Sprintf("deposit successful! New balance: %.2f", c.balance)
}

func main() {
	account1 := &checkingAccount{
		assignee:      "John Doe",
		agencyNumber:  1234,
		accountNumber: 5678,
		balance:       1000.50,
	}

	account1.cashOut(200.00)
	account1.deposit(500.00)
}
