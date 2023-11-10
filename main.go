package main

import (
	"github.com/saffrondigits/apis/database"
	"github.com/saffrondigits/apis/logger"
	"github.com/saffrondigits/apis/route"
)

func main() {
	log := logger.NewLogger()
	// Connection to the postgres db
	sql, err := database.ConnectToDatabase()
	if err != nil {
		log.Errorf("cannot connect to the db: %v", err)
		return
	}

	// Initialized the handler type
	handler := route.NewHandler(sql, log)

	r := route.Route(handler)

	log.Info("The server is started...")
	err = r.Run("0.0.0.0:8080")
	if err != nil {
		log.Errorf("cannot start the server: %v", err)
	}
}
