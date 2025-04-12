package repository

import (
	"log"

	"github.com/terumiisobe/bombus/api/models"
	"github.com/terumiisobe/bombus/db"
)

func GetColmeias() ([]models.ColmeiaModel, error) {
	var colmeias []models.ColmeiaModel
	result := db.DB.Find(&colmeias)
	return colmeias, result.Error
}

func GetColmeia(id int) (models.ColmeiaModel, error) {
	var colmeia models.ColmeiaModel
	result := db.DB.First(&colmeia, id)
	log.Println(colmeia)
	return colmeia, result.Error
}

func CreateColmeia(colmeia models.ColmeiaModel) error {
	result := db.DB.Create(&colmeia)
	return result.Error
}

func DeleteColmeia(id int) (int, error) {
	result := db.DB.Delete(&models.ColmeiaModel{}, id)
	return int(result.RowsAffected), result.Error
}
