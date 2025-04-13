package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/terumiisobe/bombus/api/services"
	"net/http"
	"strconv"
)

func GetColmeias(c *gin.Context) {
	colmeias := services.FetchColmeias()
	c.JSON(http.StatusOK, colmeias)
}

func GetColmeia(c *gin.Context) {
	paramID := c.Param("id")
	colmeiaID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	colmeia, err := services.GetColmeia(colmeiaID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, *colmeia)
}

func CreateColmeia(c *gin.Context) {
	var colmeia services.ColmeiaService
	if err := c.ShouldBindJSON(&colmeia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.CreateColmeia(colmeia)
	c.JSON(http.StatusCreated, gin.H{"message": "Colmeia created"})
}

func DeleteColmeia(c *gin.Context) {
	paramID := c.Param("id")
	colmeiaID, _ := strconv.Atoi(paramID)
	// todo: create validation when receivin from api to avoid conversion
	err := services.DeleteColmeia(colmeiaID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Colmeia deleted"})
}
