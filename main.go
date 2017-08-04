package main

import "fmt"

func main(){
    err := CreatePersonTable()
    if err.Error() != "Error 1050: Table 'person' already exists" {
        fmt.Println(err)
        return
    }
}
