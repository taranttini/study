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

	productSelected, err := selectOneProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Selecionado o registro %v - %v - %v \n", productSelected.ID, productSelected.Name, productSelected.Price)

	products, err := selectAllProduct(db)
	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("Produto: %v - %v valor %.2f \n", p.ID, p.Name, p.Price)
	}

	err = deleteProduct(db, productSelected.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Removido o registro %v \n", product.ID)
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

func selectOneProduct(db *sql.DB, productID string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name , price  FROM products WHERE id = ? ")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product

	// atencao com a ordem dos campos
	err = stmt.QueryRow(productID).Scan(&p.ID, &p.Name, &p.Price)
	// caso tivesse um contexto, usaria essa forma para executar a rotina
	// err = stmt.QueryRowContext(ctx, productID).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProduct(db *sql.DB) ([]Product, error) {

	rows, err := db.Query("SELECT id, name, price FROM products ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		// atencao com a ordem dos campos
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, productID string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ? ")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(productID)
	if err != nil {
		return err
	}

	return nil
}
