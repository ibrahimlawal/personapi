package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func CreatePersonTable() (err error) {
	d := DBParams()

	db, err := sql.Open("mysql", d.String())
	if err != nil {
		return
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		return
	}

	stmt, err := db.Prepare("CREATE TABLE person (id int NOT NULL AUTO_INCREMENT, first_name varchar(40), last_name varchar(40), PRIMARY KEY (id));")
	if err != nil {
		return
	}

	_, err = stmt.Exec()
	if err == nil {
		fmt.Println("Person Table successfully migrated....")
	}
	return
}
