package main

import (
	"fmt"
	"oop/accounts"
	"oop/customers"
	"oop/interfaces"
)

func paySlip(account interfaces.AccountInterface, amount float64) {
	account.Draw(amount)
}

func main() {
	brunoProfile := customers.Owner{Name: "Bruno", Cpf: "096.999.899-63", Profession: "Dev"}
	brunoAccount := accounts.CurrentAccount{Owner: brunoProfile, Number: 1122, Balance: 120.00}
	fmt.Println(brunoAccount)

	brunoAccount.Draw(20.)
	fmt.Println(brunoAccount)
}
