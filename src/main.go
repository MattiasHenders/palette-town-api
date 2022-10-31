package main

import (
	"fmt"

	"github.com/MattiasHenders/palette-town-api/config"
	"github.com/MattiasHenders/palette-town-api/src/db"
	"github.com/MattiasHenders/palette-town-api/src/server"
)

func main() {

	fmt.Println("Getting config...")
	config := config.GetConfig()

	fmt.Println("Connecting to DB...")
	db.Connect()
	mongoErr := db.Ping()
	if mongoErr != nil {
		panic("Failed to connect to MongoDB")
	}

	fmt.Println("Starting the server...")
	server.Start(config)
}
