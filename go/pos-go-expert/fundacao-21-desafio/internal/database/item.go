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

func (i *Item) Create(orderId string, description string, qty int, value float64) (*Item, error) {
	id := uuid.New().String()
	query := `INSERT INTO "Items" ("Id", "Description", "Qty", "Value", "OrderId") VALUES ($1, $2, $3, $4, $5)`
	_, err := i.db.Exec(query, id, description, qty, value, orderId)

	if err != nil {
		return nil, err
	}

	return &Item{
		Id:          id,
		Description: description,
		Qty:         qty,
		Value:       value,
		OrderId:     orderId,
	}, nil
}

func (i *Item) FindAll() ([]Item, error) {
	query := `SELECT "Id", "Description", "Qty", "Value", "OrderId" FROM "Items"`
	rows, err := i.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var id, description, orderId string
		var qty int
		var value float64
		if err := rows.Scan(&id, &description, &qty, &value, &orderId); err != nil {
			return nil, err
		}
		items = append(items, Item{Id: id, Description: description, Qty: qty, Value: value, OrderId: orderId})
	}

	return items, nil
}

func (i *Item) FindByOrderId(orderId string) ([]Item, error) {
	query := `SELECT "Id", "Description", "Qty", "Value", "OrderId" FROM "Items" WHERE "OrderId" = $1`
	rows, err := i.db.Query(query, orderId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var id, description, orderId string
		var qty int
		var value float64
		if err := rows.Scan(&id, &description, &qty, &value, &orderId); err != nil {
			return nil, err
		}
		items = append(items, Item{Id: id, Description: description, Qty: qty, Value: value, OrderId: orderId})
	}

	return items, nil
}
