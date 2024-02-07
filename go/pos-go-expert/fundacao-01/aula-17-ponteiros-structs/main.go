package main

import "fmt"

// ponteiros e structs

type Account struct {
	Amount int
}

func NewAccount() *Account {
	return &Account{Amount: 0}
}

func (a *Account) ChangeAmount(value int) {
	fmt.Printf("saldo anterior %d\n", a.Amount)
	a.Amount = value
	fmt.Printf("novo saldo %d\n", a.Amount)
}

func main() {
	account := Account{Amount: 10}
	//println(account.Amount)
	account.ChangeAmount(20)
	//println(account.Amount)

	println("--------")

	account2 := NewAccount()
	account2.ChangeAmount(10)
	//println(account2.Amount)
}
