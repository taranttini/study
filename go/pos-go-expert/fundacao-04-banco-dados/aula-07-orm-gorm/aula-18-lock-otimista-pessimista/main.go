package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	db.Create(&Manufacturer{Name: "abarth"})

	// lock otimista cria versao
	// lock pessimista, lock na tabela
	tx := db.Begin()
	var c Manufacturer
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}

	c.Name = "Novo dado"
	tx.Debug().Save(&c)
	tx.Commit()
}
