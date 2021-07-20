package controllers

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindRepositories(c *gin.Context) {
	var repositories []models.Repository

	if err := DB.Find(&repositories); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": repositories})
}

func FindRepository(c *gin.Context) { // Get model if exist
	var repository models.Repository

	if err := DB.Where("id = ?", c.Param("id")).First(&repository).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
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
	if err := DB.Create(&repository); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": repository})
}

func DeleteRepository(c *gin.Context) {
	// Get model if exist
	var repository models.Repository
	if err := DB.Where("id = ?", c.Param("id")).First(&repository).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := DB.Delete(&repository).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"result": true})
}
