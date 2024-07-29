package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Order struct {
	db    *sql.DB
	Id    string
	Data  string
	Items []Item
}

func NewOrder(db *sql.DB) *Order {
	return &Order{db: db}
}

func (o *Order) Create(data string) (Order, error) {
	id := uuid.New().String()
	query := "INSERT INTO orders (ID, data) VALUES ($1, $2)"
	_, err := o.db.Exec(query, data)

	if err != nil {
		return Order{}, err
	}

	return Order{
		Id:    id,
		Data:  data,
		Items: []Item{},
	}, nil

}
