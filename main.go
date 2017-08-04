package main

import "fmt"

func main() {
	err := Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer Db.Close()

	err = CreatePersonTable()

	if err != nil {
		// Silence error thrown when we try to re-migrate the Table
		if err.Error() != "Error 1050: Table 'person' already exists" {
			// Stop executing otherwise
			fmt.Println(err)
			return
		}
	}

	StartServer()
}
