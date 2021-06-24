package controllers

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindRepositories(c *gin.Context) {
	var repositories []models.Repository
	models.DB.Find(&repositories)
	c.JSON(http.StatusOK, gin.H{"data": repositories})
}

func FindRepository(c *gin.Context) { // Get model if exist
	var repository models.Repository

	if err := models.DB.Where("id = ?", c.Param("id")).First(&repository).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": repository})
}

func CreateRepository(c *gin.Context) {
	// Validate input
	var input models.CreateRepository
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	repository := models.Repository{AspaceID: input.AspaceID, Name: input.Name}
	models.DB.Create(&repository)

	c.JSON(http.StatusOK, gin.H{"data": repository})
}

func DeleteRepository(c *gin.Context) {
	// Get model if exist
	var repository models.Repository
	if err := models.DB.Where("id = ?", c.Param("id")).First(&repository).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&repository)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
