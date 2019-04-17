package main

import (
	"github.com/devsolmu/devsolmu-core/app"
	"github.com/devsolmu/devsolmu-core/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
