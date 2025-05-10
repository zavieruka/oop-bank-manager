package main

import (
	"fmt"
	"oop-bank-manager/account"
)

func main() {
	account1 := &account.CheckingAccount{
		Assignee:      "John Doe",
		AgencyNumber:  1234,
		AccountNumber: 5678,
		Balance:       1000.50,
	}

	account2 := &account.CheckingAccount{
		Assignee:      "Jane Smith",
		AgencyNumber:  4321,
		AccountNumber: 8765,
		Balance:       500.75,
	}

	fmt.Println("Initial balances:")
	fmt.Printf("Account 1: %.2f\n", account1.Balance)
	fmt.Printf("Account 2: %.2f\n", account2.Balance)

	fmt.Println("deposit function:")
	fmt.Println(account1.Deposit(200))
	fmt.Println(account1.Deposit(-50))

	fmt.Println("cashOut function:")
	fmt.Println(account1.CashOut(300))
	fmt.Println(account1.CashOut(2000))
	fmt.Println(account1.CashOut(-100))

	fmt.Println("transfer function:")
	fmt.Println(account1.Transfer(400, account2))
	fmt.Println(account1.Transfer(2000, account2))
	fmt.Println(account1.Transfer(-50, account2))

	fmt.Println("Final balances:")
	fmt.Printf("Account 1: %.2f\n", account1.Balance)
	fmt.Printf("Account 2: %.2f\n", account2.Balance)

}
