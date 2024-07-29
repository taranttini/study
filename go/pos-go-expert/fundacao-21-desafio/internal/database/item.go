package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Item struct {
	db          *sql.DB
	Id          string
	Description string
	Qty         int
	Value       float64
	OrderId     string
}

func NewItem(db *sql.DB) *Item {
	return &Item{db: db}
}

func (i *Item) Create(orderId string, description string, qty int, value float64) (Item, error) {
	id := uuid.New().String()
	query := "INSERT INTO items (ID, Description, Qty, Value, OrderId) VALUES ($1, $2, $3, $4, $5)"
	_, err := i.db.Exec(query, id, description, qty, value, orderId)

	if err != nil {
		return Item{}, err
	}

	return Item{
		Id:          id,
		Description: description,
		Qty:         qty,
		Value:       value,
		OrderId:     orderId,
	}, nil
}
