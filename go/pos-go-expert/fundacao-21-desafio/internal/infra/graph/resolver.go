package graph

import "github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ItemDB  *database.ItemRepository
	OrderDB *database.OrderRepository
}
