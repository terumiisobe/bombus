package main

import (
	"github.com/gin-gonic/gin"
	"github.com/terumiisobe/bombus/db"
	"github.com/terumiisobe/bombus/api/routes"
	"log"	
)

func main() {

	// Connects to dabatase
	db.ConnectDB()

	// Create a new router
	r := gin.Default() // or gin.New()
	routes.RegisterRoutes(r)

	// Start server
	r.Run(":8080")

	log.Println("✅ Application running")
}

