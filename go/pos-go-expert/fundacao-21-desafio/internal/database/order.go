package database

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/entity"
)

/*
	type Order struct {
		db    *sql.DB
		Id    string
		Data  string
		Items []Item
	}
*/
type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o *OrderRepository) AddItem(orderId string, descrition string, qty int, value float64) (entity.Item, error) {

	println(orderId)
	item, err := NewItemRepository(o.db).Create(orderId, descrition, qty, value)
	if err != nil {
		return entity.Item{}, err
	}
	return entity.Item{
		Id:          item.Id,
		OrderId:     item.OrderId,
		Description: item.Description,
		Qty:         item.Qty,
		Value:       item.Value,
	}, nil
}

// func (o *Order) Create(data string) (Order, error) {
func (o *OrderRepository) Create(data string) (entity.Order, error) {
	id := uuid.New().String()

	query := `
		INSERT INTO "Orders" ("Id", "Data")
		VALUES ($1, $2)
	`
	_, err := o.db.Exec(query, id, data)

	if err != nil {
		return entity.Order{}, err
	}

	return entity.Order{
		Id:    id,
		Data:  data,
		Items: []entity.Item{},
	}, nil

}

// func (o *Order) FindAll() ([]Order, error) {
func (o *OrderRepository) FindAll() ([]entity.Order, error) {
	query := `
		SELECT
			"Id",
			"Data"
		FROM
			"Orders"
	`
	rows, err := o.db.Query(query)

	if err != nil {
		return []entity.Order{}, err
	}
	defer rows.Close()

	orders := []entity.Order{}
	for rows.Next() {
		var id, data string
		if err := rows.Scan(&id, &data); err != nil {
			return nil, err
		}
		orders = append(orders, entity.Order{Id: id, Data: data})
	}

	return orders, nil
}

// func (o *Order) FindByItemId(itemId string) (Order, error) {
func (o *OrderRepository) FindByItemId(itemId string) (entity.Order, error) {
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
		ORDER BY "o"."Id", "i"."Id"
	`
	var id, data string
	err := o.db.QueryRow(query, itemId).Scan(&id, &data)

	if err != nil {
		return entity.Order{}, err
	}

	return entity.Order{Id: id, Data: data}, nil
}

type NullableString struct{ sql.NullString }
type NullableInt32 struct{ sql.NullInt32 }
type NullableFloat64 struct{ sql.NullFloat64 }

func (ns *NullableString) ValueOrDefault(defaultValue string) string {
	if ns.Valid {
		return ns.String
	}
	return defaultValue
}
func (ns *NullableInt32) ValueOrDefault(defaultValue int32) int32 {
	if ns.Valid {
		return ns.Int32
	}
	return defaultValue
}
func (ns *NullableFloat64) ValueOrDefault(defaultValue float64) float64 {
	if ns.Valid {
		return ns.Float64
	}
	return defaultValue
}

// func (o *Order) FindAllIncludeItems() ([]Order, error) {
func (o *OrderRepository) FindAllIncludeItems() ([]entity.Order, error) {
	query := `
	SELECT
		"o"."Id" AS "OrderId",
		"o"."Data",
		"i"."Id",
		"i"."Description",
		"i"."Qty",
		"i"."Value"
	FROM
		"Orders" "o"
		LEFT JOIN
			"Items" "i" ON "i"."OrderId" = "o"."Id"
`
	rows, err := o.db.Query(query)

	if err != nil {
		return []entity.Order{}, err
	}
	defer rows.Close()

	orders := []entity.Order{}
	items := []entity.Item{}
	var lastItemId string
	for rows.Next() {
		var orderId, data string
		var _id, _description NullableString
		var _qty NullableInt32
		var _value NullableFloat64
		if err := rows.Scan(&orderId, &data, &_id, &_description, &_qty, &_value); err != nil {
			return nil, err
		}

		id := _id.ValueOrDefault("")
		description := _description.ValueOrDefault("")
		qty := _qty.ValueOrDefault(0)
		value := _value.ValueOrDefault(0)

		if lastItemId != orderId {
			//fmt.Printf(".. novo item %v orderId %v \n", id, orderId)
			items = []entity.Item{}
			if len(id) > 0 {
				items = append(items, entity.Item{Id: id, Description: description, Qty: int(qty), Value: value})
			}
			orders = append(orders, entity.Order{Id: orderId, Data: data, Items: items})
			lastItemId = orderId

		} else {
			//fmt.Printf(">>  add item %v orderId %v \n", id, orderId)
			items = append(items, entity.Item{Id: id, Description: description, Qty: int(qty), Value: value})

			if len(id) > 0 {
				orders[len(orders)-1].Items = items
			}
		}
	}

	return orders, nil
}

// func (o *Order) FindByOrderIdIncludeItems(orderId string) ([]Order, error) {
func (o *OrderRepository) FindByOrderIdIncludeItems(orderId string) ([]entity.Order, error) {
	query := `
	SELECT
		"o"."Id" AS "OrderId",
		"o"."Data",
		"i"."Id",
		"i"."Description",
		"i"."Qty",
		"i"."Value"
	FROM
		"Orders" "o"
		LEFT JOIN
			"Items" "i" ON "i"."OrderId" = "o"."Id"
	WHERE
		"o"."Id" = $1
`
	rows, err := o.db.Query(query)

	if err != nil {
		return []entity.Order{}, err
	}
	defer rows.Close()

	orders := []entity.Order{}
	items := []entity.Item{}
	lastItemId := ""
	for rows.Next() {
		var orderId, id, data, description string
		var qty int32
		var value float64
		if err := rows.Scan(&orderId, &data, &id, &description, &qty, &value); err != nil {
			return nil, err
		}

		if lastItemId != orderId {
			//fmt.Printf(".. novo item %v orderId %v \n", id, orderId)
			items = []entity.Item{}
			items = append(items, entity.Item{Id: id, Description: description, Qty: int(qty), Value: value})

			orders = append(orders, entity.Order{Id: orderId, Data: data, Items: items})
			lastItemId = orderId

		} else {
			//fmt.Printf(">>  add item %v orderId %v \n", id, orderId)
			items = append(items, entity.Item{Id: id, Description: description, Qty: int(qty), Value: value})

			orders[len(orders)-1].Items = items
		}
	}

	return orders, nil
}
