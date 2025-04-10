package controllers

import (
	"github.com/terumiisobe/bombus/api/services"
	"github.com/gin-gonic/gin"
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
	colmeia := services.GetColmeia(colmeiaID)
	c.JSON(http.StatusOK, colmeia)
}

func CreateColmeia(c *gin.Context) {
	var colmeia services.Colmeia
	if err := c.ShouldBindJSON(&colmeia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.CreateColmeia(colmeia)
	c.JSON(http.StatusCreated, gin.H{"message": "Colmeia created"})
}

func DeleteColmeia(c *gin.Context) {
	paramID := c.Param("id")
	colmeiaID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	services.DeleteColmeia(colmeiaID)
	c.JSON(http.StatusOK, gin.H{"message": "Colmeia deleted"})
}
