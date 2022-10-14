package main

import (
	"github.com/MattiasHenders/palette-town-api/app"
	"github.com/MattiasHenders/palette-town-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
