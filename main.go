package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/terumiisobe/bombus/api/models"
	"github.com/terumiisobe/bombus/db"
	"log"	
)

func main() {

	// Connects to dabatase
	db.ConnectDB()

	// Create a new router
	r := gin.Default() // or gin.New()

	// Define routes
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})

	// Start server
	r.Run(":8080")

	log.Println("✅ Application running")
}

