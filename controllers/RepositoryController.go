package controllers

import (
	"github.com/dmnyu/go-medialog/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func FindRepositories(c *gin.Context) {
	var repositories []db.Repository

	if err := db.DB.Find(&repositories); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": repositories})
}

func FindRepositoryAPI(c *gin.Context) { // Get model if exist
	var repository db.Repository

	if err := db.DB.Where("id = ?", c.Param("id")).First(&repository).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": repository})
}

func CreateRepositoryAPI(c *gin.Context) {
	// Validate input
	var input db.Repository
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	repository := db.Repository{AspaceID: input.AspaceID, Slug: input.Slug, Name: input.Name}
	if err := db.DB.Create(&repository); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": repository})
}

func DeleteRepository(c *gin.Context) {
	// Get model if exist
	var repository db.Repository
	if err := db.DB.Where("id = ?", c.Param("id")).First(&repository).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := db.DB.Delete(&repository).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"result": true})
}

func GetRepositories() []db.Repository {
	repositories := []db.Repository{}
	db.DB.Find(&repositories)
	return repositories
}

func GetRepositoryByID(id int) (db.Repository, error) { return db.FindRepository(id) }

func CreateRepository(repo db.Repository) (int, error) {
	asRepository, err := FindAspaceRepository(repo.AspaceID)
	if err != nil {
		return 0, err
	}

	repository := db.Repository{
		Model:    gorm.Model{},
		AspaceID: repo.AspaceID,
		Slug:     asRepository.Slug,
		Name:     asRepository.Name,
	}

	if err := db.DB.Create(&repository).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil

}
