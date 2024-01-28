package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProductsGorm struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// auto migration - criar migracoes automaticas
	db.AutoMigrate(ProductsGorm{})

	// insert
	db.Create(&ProductsGorm{
		Name:  "macbook",
		Price: 9999.99,
	})

	// insert many
	products := []ProductsGorm{
		{Name: "linuxbook", Price: 499.99},
		{Name: "windowsbook", Price: 2999.99},
	}
	db.Create(&products)

	fmt.Println("=== select first by key ==========")

	// select one
	var product ProductsGorm
	db.First(&product, 1)
	fmt.Println(product)

	fmt.Println("=== select first by param name ===")

	// select one by param
	var productByName ProductsGorm
	db.First(&productByName, "name=?", "windowsbook")
	fmt.Println(productByName)

	fmt.Println("=== select all ===================")

	// select all
	var allProducts []ProductsGorm
	db.Find(&allProducts)
	for _, product := range allProducts {
		fmt.Println(product)
	}

}
