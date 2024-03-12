package main

import (
	"api-assignmet/lib"
	"fmt"
)

func main() {
	_ , err := lib.InitDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("db berjalan")
}