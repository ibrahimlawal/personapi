package main

import "fmt"

func main() {
	err := Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer Db.Close()

	// Silence error thrown when we try to re-migrate the Table
	// Stop executing otherwise
	err = CreatePersonTable()
	if err != nil {
		if err.Error() != "Error 1050: Table 'person' already exists" {
			fmt.Println(err)
			return
		}
	}
}
