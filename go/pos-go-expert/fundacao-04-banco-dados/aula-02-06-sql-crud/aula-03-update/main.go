package main

import (
	"database/sql"
	"fmt"

	// o _ serve para manter o pacote, pois estamos usando ele indiretamente
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("nootebook", 1000)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Salvo o registro %v - %v - %v \n", product.ID, product.Name, product.Price)

	product.Price = 1999.99
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Atualizado o registro %v - %v - %v \n", product.ID, product.Name, product.Price)
}

func insertProduct(db *sql.DB, product *Product) error {

	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// atencao com a ordem dos campos
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ? , price = ? WHERE id = ? ")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// atencao com a ordem dos campos
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}
