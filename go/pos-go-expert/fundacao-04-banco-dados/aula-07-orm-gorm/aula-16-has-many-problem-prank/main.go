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

	// remover validacao da chave foreign
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

	manufacturerGM := &Manufacturer{Name: "gm"}
	db.Create(&manufacturerGM)

	// create cars
	carF50 := &Car{Name: "F50", Price: 999999.99, ManufacturerID: manufacturerFerrari.ID}
	carEnzo := &Car{Name: "Enzo", Price: 599999.99, ManufacturerID: manufacturerFerrari.ID}
	carCorvette := &Car{Name: "Corvette", Price: 199999.99, ManufacturerID: manufacturerGM.ID}
	carOnix := &Car{Name: "Onix", Price: 99999.99, ManufacturerID: manufacturerGM.ID}

	db.Create(carF50)
	db.Create(carEnzo)
	db.Create(carCorvette)
	db.Create(carOnix)

	// create serial numbers
	db.Create(&SerialNumber{Number: "123", CarID: carF50.ID})
	db.Create(&SerialNumber{Number: "456", CarID: carEnzo.ID})
	db.Create(&SerialNumber{Number: "789", CarID: carCorvette.ID})
	db.Create(&SerialNumber{Number: "147", CarID: carOnix.ID})

	// mapear carros das marcas
	var manufacturers []Manufacturer
	// not load
	// err = db.Model(&Manufacturer{}).Preload("Cars").Preload("SerialNumber").Find(&manufacturers).Error

	// carrega corretamente
	err = db.Model(&Manufacturer{}).Preload("Cars").Preload("Cars.SerialNumber").Find(&manufacturers).Error

	// ou
	// err = db.Model(&Manufacturer{}).Preload("Cars.SerialNumber").Find(&manufacturers).Error
	if err != nil {
		panic(err)
	}

	for _, manufacture := range manufacturers {
		fmt.Println(manufacture.Name, ":")
		for _, car := range manufacture.Cars {
			//fmt.Printf("- %v %.2f %v \n", car.Name, car.Price, car.SerialNumber.Number)
			fmt.Printf("- %v %.2f %v \n", car.Name, car.Price, car.SerialNumber.Number)
		}
	}

}
