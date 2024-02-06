package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manufacturer struct {
	ID   int `gorm:primaryKey`
	Name string
	Cars []Car
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

	// create manufacture
	manufacturerFerrari := &Manufacturer{Name: "ferrari"}
	db.Create(&manufacturerFerrari)

	manufacturerGM := &Manufacturer{Name: "gm"}
	db.Create(&manufacturerGM)

	// create car
	cars := []Car{
		{Name: "F50", Price: 999999.99, ManufacturerID: manufacturerFerrari.ID},
		{Name: "Enzo", Price: 599999.99, ManufacturerID: manufacturerFerrari.ID},
		{Name: "Corvette", Price: 199999.99, ManufacturerID: manufacturerGM.ID},
		{Name: "Onix", Price: 99999.99, ManufacturerID: manufacturerGM.ID},
	}
	//fmt.Println(cars)
	db.Create(&cars)

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
