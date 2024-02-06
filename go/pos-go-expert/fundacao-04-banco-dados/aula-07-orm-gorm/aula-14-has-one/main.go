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
	SerialNumber   SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID     int `gorm:primaryKey`
	Number string
	CarID  int
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
	err = db.AutoMigrate(&Car{}, &Manufacturer{}, &SerialNumber{})
	if err != nil {
		panic(err)
	}

	// create manufacture
	manufacturerFerrari := &Manufacturer{Name: "ferrari"}
	db.Create(&manufacturerFerrari)

	// create car
	car := &Car{
		Name:           "458 italia",
		Price:          999999.99,
		ManufacturerID: manufacturerFerrari.ID,
	}
	db.Create(car)

	// create serial number
	db.Create(&SerialNumber{
		Number: "12345678",
		CarID:  car.ID,
	})

	// select all
	var carsManufacturers []Car
	db.Preload("Manufacturer").Preload("SerialNumber").Find(&carsManufacturers)
	for _, car := range carsManufacturers {
		fmt.Println(car.Name, car.Manufacturer.Name, car.SerialNumber.Number)
	}

}
