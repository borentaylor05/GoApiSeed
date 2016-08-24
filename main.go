package main

import (
	"GoApiSeed/db"
	"GoApiSeed/server"
	"fmt"
)

func main() {
	db, err := db.Open()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	server.Init()
}
