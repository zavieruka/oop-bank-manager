package main

import "fmt"

type checkingAccount struct {
	assignee      string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	account := checkingAccount{
		assignee:      "John Doe",
		agencyNumber:  1234,
		accountNumber: 567890,
		balance:       1000.50,
	}

	fmt.Println("Account Holder:", account.assignee)
	fmt.Println("Agency Number:", account.agencyNumber)
	fmt.Println("Account Number:", account.accountNumber)
	fmt.Println("Balance:", account.balance)
}
