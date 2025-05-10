package account

import "fmt"

type CheckingAccount struct {
	Assignee      string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *CheckingAccount) CashOut(value float64) string {
	canCashOut := value <= c.Balance && value > 0
	if value <= 0 {
		return "cash out amount must be greater than zero"
	}
	if canCashOut {
		c.Balance -= value
		return fmt.Sprintf("cash out successful! New balance: %.2f", c.Balance)
	}
	return fmt.Sprintf("insufficient funds! Current balance: %.2f", c.Balance)

}

func (c *CheckingAccount) Deposit(value float64) string {
	if value <= 0 {
		return "deposit amount must be greater than zero"
	}
	c.Balance += value
	return fmt.Sprintf("deposit successful! New balance: %.2f", c.Balance)
}

func (c *CheckingAccount) Transfer(value float64, target *CheckingAccount) string {
	canTransfer := value <= c.Balance && value > 0
	if value <= 0 {
		return "transfer amount must be greater than zero"
	}
	if canTransfer {
		c.Balance -= value
		target.Balance += value
		return fmt.Sprintf("transfer successful! New balance: %.2f", c.Balance)
	}
	return fmt.Sprintf("insufficient funds! Current balance: %.2f", c.Balance)
}
