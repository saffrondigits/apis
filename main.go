package main

import (
	"log"

	"github.com/saffrondigits/apis/database"
	"github.com/saffrondigits/apis/route"
)

func main() {
	// Connection to the postgres db
	sql, err := database.ConnectToDatabase()
	if err != nil {
		log.Println("cannot connect to the db")
		return
	}

	// Initialized the handler type
	handler := route.NewHandler(sql, nil)

	r := route.Route(handler)

	err = r.Run("0.0.0.0:8080")
	if err != nil {
		log.Println("cannot start the server!")
	}
}
