package main

import (
	"fmt"
)

func CreatePersonTable() (err error) {
	stmt, err := Db.Prepare("CREATE TABLE person (id int NOT NULL AUTO_INCREMENT, first_name varchar(40), last_name varchar(40), PRIMARY KEY (id));")
	if err != nil {
		return
	}

	_, err = stmt.Exec()
	if err == nil {
		fmt.Println("Person Table successfully migrated....")
	}
	return
}
