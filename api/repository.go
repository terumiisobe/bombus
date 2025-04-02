package repository

import (
	"github.com/terumiisobe/bombus/api/models"
	"github.com/terumiisobe/bombus/db"
	"gorm.io/gorm"
)

func GetColmeias() ([]models.Colmeia, error) {
	var colmeias []models.Colmeia
	result := db.DB.Find(&colmeias)
	return colmeias, result.Error 
}

func GetColmeia(id uint) (models.Colmeia, error) {
	var colmeia models.Colmeia
	result := db.DB.First(&colmeia, id)
	return colmeia, result.Error 
}

func CreateColmeias(colmeia models.Colmeia) error {
	result := db.DB.Create(&colmeia)
	return result.Error 
}

func DeleteColmeias(id uint) error {
	result := db.DB.Delete(&models.Colmeia{}, id)
	return result.Error 
}
