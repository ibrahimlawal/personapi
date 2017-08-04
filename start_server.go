package main

import (
	_ "bytes"
	_ "database/sql"
	_ "fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func StartServer() {
	type Person struct {
		Id         int
		First_Name string
		Last_Name  string
	}
	router := gin.Default()
	// Add API handlers here
	router.Run(":3000")
}
