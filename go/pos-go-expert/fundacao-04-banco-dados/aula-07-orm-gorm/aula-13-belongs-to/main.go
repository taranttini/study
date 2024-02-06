package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manufacturer struct {
	ID   int `gorm:primaryKey`
	Name string
}

type Car struct {
	ID             int `gorm:"primaryKey"`
	Name           string
	Price          float64
	ManufacturerID int
	Manufacturer   Manufacturer
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Exec("drop table if exists serial_numbers")
	db.Exec("drop table if exists cars_manufacturers")
	db.Exec("drop table if exists cars")
	db.Exec("drop table if exists manufacturers")
	db.Exec("drop table if exists product_gorms")

	// auto migration - criar migracoes automaticas
	db.AutoMigrate(Car{}, Manufacturer{})

	manufacturerFerrari := Manufacturer{Name: "ferrari"}
	db.Create(&manufacturerFerrari)

	manufacturerBugatti := Manufacturer{Name: "bugatti"}
	db.Create(&manufacturerBugatti)

	manufacturerGM := Manufacturer{Name: "gm"}
	db.Create(&manufacturerGM)

	// insert many
	cars := []Car{
		{Name: "ferrari", Price: 999999.99, ManufacturerID: manufacturerFerrari.ID},
		{Name: "bugatti", Price: 599999.99, ManufacturerID: manufacturerBugatti.ID},
		{Name: "corvette", Price: 199999.99, ManufacturerID: manufacturerGM.ID},
	}
	//fmt.Println(cars)
	db.Create(&cars)

	// select all
	var carsManufacturers []Car
	db.Preload("Manufacturer").Find(&carsManufacturers)
	for _, car := range carsManufacturers {
		fmt.Println(car.Name, car.Manufacturer.Name)
	}

}
