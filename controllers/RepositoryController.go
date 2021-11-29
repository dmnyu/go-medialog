package controllers

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func FindRepositories(c *gin.Context) {
	var repositories []models.Repository

	if err := models.DB.Find(&repositories); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": repositories})
}

func FindRepositoryAPI(c *gin.Context) { // Get model if exist
	var repository models.Repository

	if err := models.DB.Where("id = ?", c.Param("id")).First(&repository).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": repository})
}

func CreateRepositoryAPI(c *gin.Context) {
	// Validate input
	var input models.CreateRepository
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	repository := models.Repository{AspaceID: input.AspaceID, Slug: input.Slug, Name: input.Name}
	if err := models.DB.Create(&repository); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": repository})
}

func DeleteRepository(c *gin.Context) {
	// Get model if exist
	var repository models.Repository
	if err := models.DB.Where("id = ?", c.Param("id")).First(&repository).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := models.DB.Delete(&repository).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"result": true})
}

func GetRepositories() []models.Repository {
	repositories := []models.Repository{}
	models.DB.Find(&repositories)
	return repositories
}

func FindRepository(id int) (models.Repository, error) {
	repository := models.Repository{}
	if err = models.DB.Find(&repository, "id = ?", id).Error; err != nil {
		return repository, err
	}

	return repository, nil
}

func CreateRepository(repo models.CreateRepository) (int, error) {
	asRepository, err := FindAspaceRepository(repo.AspaceID)
	if err != nil {
		return 0, err
	}

	repository := models.Repository{
		Model:    gorm.Model{},
		ID:       0,
		AspaceID: repo.AspaceID,
		Slug:     asRepository.Slug,
		Name:     asRepository.Name,
	}

	if err := models.DB.Create(&repository).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil

}

