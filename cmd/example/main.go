package main

import (
	"github.com/efectn/fiber-boilerplate/internal"
	"github.com/efectn/fiber-boilerplate/internal/config"
)

// TODO: Register routes in main.go and DRY fixes.
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

	// Run webserver
	ws.ListenWebServer()
}
