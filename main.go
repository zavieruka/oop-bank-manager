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
		Balance:       100}

	fmt.Println(jonhAccount)

}
