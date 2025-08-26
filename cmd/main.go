package main

import (
	"gin-app/config"

	"gin-app/internal/pkg/router"
)

func main() {
	// Initialize Database
	config.InitDB()

	// Initial route
	r := router.SetupRouter()
	// Run server
	r.Run(":8080")
}
