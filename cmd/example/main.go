package main

import (
	"github.com/efectn/fiber-boilerplate/internal"
	"github.com/efectn/fiber-boilerplate/internal/config"
	"github.com/efectn/fiber-boilerplate/internal/routes"
)

func main() {
	config, err := config.ParseConfig("example")
	if err != nil {
		panic(err)
	}

	// Setup webserver
	ws, err := internal.SetupWebServer(config)
	if err != nil {
		panic(err)
	}

	// Register Routes
	routes.RegisterAPIRoutes(ws.App())

	// Run webserver
	ws.ListenWebServer()
}
