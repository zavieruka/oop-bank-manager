package main

import (
	"fmt"
	"oop-bank-manager/account"
	"oop-bank-manager/clients"
)

func main() {
	jonhAssignee := clients.Assignee{
		Name:       "jonh doe",
		CPF:        "123.456.789-00",
		Profession: "worker"}

	jonhAccount := account.CheckingAccount{
		Assignee:      jonhAssignee,
		AgencyNumber:  123,
		AccountNumber: 456789,
	}

	jonhAccount.Deposit(1000)
	fmt.Println(jonhAccount.GetBalance())

	janeAssignee := clients.Assignee{
		Name:       "jane doe",
		CPF:        "987.654.321-00",
		Profession: "worker"}
	janeAccount := account.SavingAccount{
		Assignee:      janeAssignee,
		AgencyNumber:  123,
		AccountNumber: 456789,
		Operation:     1,
	}

	janeAccount.Deposit(1000)
	fmt.Println(janeAccount.GetBalance())

}
