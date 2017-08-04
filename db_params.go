package main

import (
	"fmt"
	"os"
)

type DB struct {
	User string
	Pass string
	Host string
	Name string
}

func (d DB) String() string {
	return fmt.Sprintf("%v:%v@tcp(%v:3306)/%v", d.User, d.Pass, d.Host, d.Name)
}

func (d *DB) loadEnv() {
	d.User = os.Getenv("PERSON_API_DBUSER")
	d.Pass = os.Getenv("PERSON_API_DBPASS")
	d.Host = os.Getenv("PERSON_API_DBHOST")
	d.Name = os.Getenv("PERSON_API_DBNAME")
}

func DBParams() (d DB) {
	d = DB{}
	d.loadEnv()
	return
}
