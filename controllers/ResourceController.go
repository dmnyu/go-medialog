package controllers

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	resourceRecord, err := FindAspaceResource(resource.RepositoryID, resource.AspaceResourceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{"data": resourceRecord})
}

func CreateResource(c *gin.Context) {
	// Validate input
	var input models.CreateResource
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a resource
	aspaceResource, err := FindAspaceResource(input.AspaceRepositoryID, input.AspaceResourceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resource := models.Resource{
		Model:                     gorm.Model{},
		AspaceResourceID:          input.AspaceResourceID,
		RepositoryID:              input.AspaceRepositoryID,
		AspaceResourceTitle:       aspaceResource.Title,
		AspaceResourceIdentifiers: aspaceResource.MergeIDs(),
	}

	if err := models.DB.Create(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": resource})
}

func DeleteResource(c *gin.Context) {
	// Get model if exist
	var resource models.Resource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&resource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	if err := models.DB.Delete(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func FindRecent() []models.Entry {
	entries := []models.Entry{}
	models.DB.Limit(20).Find(&entries)
	return entries
}
