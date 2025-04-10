package db

import (
	"github.com/terumiisobe/bombus/api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("db/database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to the database:", err)
	}

	// Auto-Migrate Tables
	DB.AutoMigrate(&models.ColmeiaModel{}, &models.Inspection{})

	log.Println("✅ Database Connected & Migrated Successfully!")
}
