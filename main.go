package main

import (
	"example.com/hexagonal-auth/app"
	"example.com/hexagonal-auth/logger"
)

func main() {
	logger.Info("Starting the application...")

	app.Start()
}
