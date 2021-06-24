package controllers

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindResources(c *gin.Context) {
	var resources []models.Resource
	models.DB.Find(&resources)
	c.JSON(http.StatusOK, gin.H{"data": resources})
}

func FindResource(c *gin.Context) { // Get model if exist
	var resource models.Resource

	if err := models.DB.Where("id = ?", c.Param("id")).First(&resource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resource})
}

func CreateResource(c *gin.Context) {
	// Validate input
	var input models.CreateResource
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	resource := models.Resource{AspaceID: input.AspaceID, RepositoryID: input.RepositoryID, Name: input.Name}
	models.DB.Create(&resource)

	c.JSON(http.StatusOK, gin.H{"data": resource})
}

func DeleteResource(c *gin.Context) {
	// Get model if exist
	var resource models.Resource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&resource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&resource)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
