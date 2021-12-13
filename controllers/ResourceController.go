package controllers

import (
	"github.com/dmnyu/go-medialog/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetResources() []db.Resource {
	var resources = []db.Resource{}
	db.DB.Find(&resources)
	return resources
}

func FindResources(c *gin.Context) {
	var resources []db.Resource
	db.DB.Find(&resources)
	c.JSON(http.StatusOK, gin.H{"data": resources})
}

func GetResource(c *gin.Context) (db.Resource, error) {
	resource := db.Resource{}
	resourceId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return resource, err
	}

	resource, err = db.FindResource(resourceId)
	if err != nil {
		return resource, err
	}

	return resource, nil
}

func FindResourcesByRepository(repositoryID int) ([]db.Resource, error) {
	var resources []db.Resource
	if err := db.DB.Where("repository_id = ?", repositoryID).Find(&resources).Error; err != nil {
		return resources, err
	}
	return resources, nil
}

func CreateResource(resource db.Resource) error {

	if err := db.DB.Create(&resource).Error; err != nil {
		return err
	}

	return nil
}

func DeleteResource(c *gin.Context) {
	// Get model if exist
	var resource db.Resource
	if err := db.DB.Where("id = ?", c.Param("id")).First(&resource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	if err := db.DB.Delete(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func FindRecent() []db.Entry {
	entries := []db.Entry{}
	db.DB.Limit(20).Find(&entries)
	return entries
}
