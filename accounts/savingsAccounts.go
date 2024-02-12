package accounts

import (
	"fmt"
	"oop/customers"
)

type SavingsAccount struct {
	Owner                           customers.Owner
	agencyNumber, number, operation int
	Balance                         float64
}

func (c *SavingsAccount) Draw(value float64) {
	if value > c.Balance || value < 0 {
		fmt.Println("Nao foi possível realizar o saque.")
	} else {
		c.Balance -= value
		fmt.Println("Saque feito com sucesso.")
	}
}

func (c *SavingsAccount) Deposit(value float64) {
	if value < 0 {
		fmt.Println("Nao foi possível realizar o depósito.")
	} else {
		c.Balance += value
		fmt.Println("Depósito feito com sucesso.")
	}
}

func (b *SavingsAccount) Transfer(receiver *SavingsAccount, amount float64) bool {
	if b.Balance < amount || amount < 0 {
		fmt.Println("Não foi possível realizar a transferência.")
		return false
	} else {
		b.Balance -= amount
		receiver.Balance += amount
		fmt.Println("Transferência realizada com sucesso!")
		return true
	}
}
