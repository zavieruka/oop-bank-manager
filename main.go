package main

import "fmt"

type checkingAccount struct {
	assignee      string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	account1 := &checkingAccount{
		assignee:      "John Doe",
		agencyNumber:  1234,
		accountNumber: 5678,
		balance:       1000.50,
	}

	account2 := &checkingAccount{
		assignee:      "John Doe",
		agencyNumber:  1234,
		accountNumber: 5678,
		balance:       1000.50,
	}

	fmt.Println("Account 1:", account1)
	fmt.Println("Account 2:", account2)

	fmt.Println("Same memory location?", account1 == account2)
	fmt.Println("Same content?", *account1 == *account2)
	fmt.Printf("Account 1 pointer: %p\n", account1)
	fmt.Printf("Account 2 pointer: %p\n", account2)
}
