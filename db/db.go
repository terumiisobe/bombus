package db

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/terumiisobe/bombus/api/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("db/database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to the database:", err)
	}

	// Auto-Migrate Tables
	DB.AutoMigrate(&models.Colmeia{}, &models.Inspection{})

	log.Println("✅ Database Connected & Migrated Successfully!")
}
