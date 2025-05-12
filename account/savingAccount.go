package account

import (
	"fmt"
	"oop-bank-manager/clients"
)

type SavingAccount struct {
	Assignee      clients.Assignee
	AgencyNumber  int
	AccountNumber int
	Operation     int
	balance       float64
}

func (c *SavingAccount) Deposit(value float64) string {
	if value <= 0 {
		return "deposit amount must be greater than zero"
	}
	c.balance += value
	return fmt.Sprintf("deposit successful! New balance: %.2f", c.balance)
}

func (c *SavingAccount) Transfer(value float64, target *CheckingAccount) string {
	canTransfer := value <= c.balance && value > 0
	if value <= 0 {
		return "transfer amount must be greater than zero"
	}
	if canTransfer {
		c.balance -= value
		target.balance += value
		return fmt.Sprintf("transfer successful! New balance: %.2f", c.balance)
	}
	return fmt.Sprintf("insufficient funds! Current balance: %.2f", c.balance)
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
