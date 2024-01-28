package main

import (
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
}
