package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Car struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// auto migration - criar migracoes automaticas
	db.AutoMigrate(Car{})

	db.Exec("TRUNCATE TABLE cars")

	// insert
	db.Create(&Car{
		Name:  "ferrari",
		Price: 999999.99,
	})

	// insert many
	cars := []Car{
		{Name: "bugatti", Price: 599999.99},
		{Name: "corvette", Price: 199999.99},
	}
	db.Create(&cars)

	fmt.Println("=== select first by key ==========")

	// soft delete
	var car Car
	db.Find(&car, "name = ? ", "ferrari")
	db.Delete(&car)

}
