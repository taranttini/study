package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manufacturer struct {
	ID   int `gorm:primaryKey`
	Name string
	Cars []Car `gorm:"many2many:cars_manufacturers;"`
}

type Car struct {
	ID            int `gorm:"primaryKey"`
	Name          string
	Price         float64
	Manufacturers []Manufacturer `gorm:"many2many:cars_manufacturers;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// remover validacao da chave foreign
	db.Exec("drop table if exists serial_numbers")
	db.Exec("drop table if exists cars_manufacturers")
	db.Exec("drop table if exists cars")
	db.Exec("drop table if exists manufacturers")
	db.Exec("drop table if exists product_gorms")

	// auto migration - criar migracoes automaticas
	err = db.AutoMigrate(&Car{}, &Manufacturer{})
	if err != nil {
		panic(err)
	}

	//db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	//db.Exec("TRUNCATE cars_manufacturies ")
	//db.Exec("TRUNCATE cars ")
	//db.Exec("TRUNCATE manufacturers ")
	//db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// create manufacture
	manufacturerAbarth := Manufacturer{Name: "abarth"}
	db.Create(&manufacturerAbarth)

	manufacturerFiat := Manufacturer{Name: "fiat"}
	db.Create(&manufacturerFiat)

	// create cars
	carAbarth := &Car{Name: "abarth", Price: 999.99, Manufacturers: []Manufacturer{
		manufacturerFiat, manufacturerAbarth}}

	db.Create(carAbarth)

	// mapear carros das marcas
	var manufacturers []Manufacturer
	err = db.Model(&Manufacturer{}).Preload("Cars").Find(&manufacturers).Error
	if err != nil {
		panic(err)
	}

	for _, manufacture := range manufacturers {
		fmt.Println(manufacture.Name, ":")
		for _, car := range manufacture.Cars {
			fmt.Printf("- %v %.2f \n", car.Name, car.Price)
		}
	}

}
