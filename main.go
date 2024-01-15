	// main.go
package main

import (
	"github.com/Rahmatdev030605/app-coffee/db"
	"github.com/Rahmatdev030605/app-coffee/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	// Setup regular routes
	routes.SetupRoutes(r)

	// Setup authentication-related routes
	routes.SetupAuthRoutes(r)

	r.Run(":8080")
}
