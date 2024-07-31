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

	query := `
		INSERT INTO "Orders" ("Id", "Data") 
		VALUES ($1, $2)
	`
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
	query := `
		SELECT 
			"Id", 
			"Data" 
		FROM 
			"Orders"
	`
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

func (o *Order) FindByItemId(itemId string) (Order, error) {
	query := `
		SELECT 
			"o"."Id", 
			"o"."Data" 
		FROM 
			"Orders" "o" 
		INNER JOIN 
			"Items" "i" 
				ON "i"."OrderId" = "o"."Id" 
		WHERE 
			"i"."Id" = $1
	`
	var id, data string
	err := o.db.QueryRow(query, itemId).Scan(&id, &data)

	if err != nil {
		return Order{}, err
	}

	return Order{Id: id, Data: data}, nil
}
