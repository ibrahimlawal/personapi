package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect() (err error) {
	d := DBParams()

	Db, err = sql.Open("mysql", d.String())
	if err != nil {
		return
	}
	fmt.Println("Database connection successful...")

	// make sure connection is available
	err = Db.Ping()
	if err != nil {
		return
	}
	fmt.Println("Database ping successful...")
	return
}
