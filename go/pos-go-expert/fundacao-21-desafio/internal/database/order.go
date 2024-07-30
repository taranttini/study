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

	query := `INSERT INTO "Orders" ("Id", "Data") VALUES ($1, $2)`
	_, err := o.db.Exec(query, id, data)

	if err != nil {
		return Order{}, err
	}

	return Order{
		Id:    id,
		Data:  data,
		Items: []Item{},
	}, nil

}

func (o *Order) FindAll() ([]Order, error) {
	query := `SELECT "Id", "Data" FROM "Orders"`
	rows, err := o.db.Query(query)

	if err != nil {
		return []Order{}, err
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		var id, data string
		if err := rows.Scan(&id, &data); err != nil {
			return nil, err
		}
		orders = append(orders, Order{Id: id, Data: data})
	}

	return orders, nil

}
