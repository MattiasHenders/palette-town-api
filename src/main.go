package main

import (
	"fmt"

	"github.com/MattiasHenders/palette-town-api/config"
	"github.com/MattiasHenders/palette-town-api/src/server"
)

func Main() {

	fmt.Println("Getting config...")
	config := config.GetConfig()

	fmt.Println("Starting the server...")
	server.Start(config)
}
