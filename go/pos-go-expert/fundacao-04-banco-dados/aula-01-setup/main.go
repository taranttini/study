package main

import (
	"database/sql"
	// o _ serve para manter o pacote, pois estamos usando ele indiretamente
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
