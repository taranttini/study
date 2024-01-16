package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	// se nao usar o json com caixa baixa ele vai exibir a propriedade com primeira letra em caixa alta
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	// Amount int    `json:"-"` // omitir esse dado
	// Amount int `json:"amount" validate:"gt=0"` // omitir esse dado
}

func main() {
	account := Account{
		Name: "zero", Amount: -1000,
	}

	fmt.Printf("Saldo %d\n", account.Amount)

	response, err := json.Marshal(account)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(response))

	//encoder := json.NewEncoder(os.Stdout)
	//encoder.Encode(account)
	err = json.NewEncoder(os.Stdout).Encode(account)

	if err != nil {
		println(err)
	}

	jsonFromByte := []byte(`{"name":"two","amount":-50}`)
	var jsonParsed Account
	err = json.Unmarshal(jsonFromByte, &jsonParsed)

	if err != nil {
		println(err)
	}

	fmt.Println(jsonParsed)
}
