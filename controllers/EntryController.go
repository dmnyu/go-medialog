package controllers

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateEntry(c *gin.Context) {
	input := models.Entry{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	models.DB.Create(input)
	c.JSON(http.StatusOK, input)
}

func FindEntries(c *gin.Context) {
	entries := []models.Entry{}
	models.DB.Find(&entries)
	c.JSON(http.StatusOK, entries)
}

func FindEntry(c *gin.Context) {
	entry := models.Entry{}
	if err := models.DB.Where("id = ?", c.Param("id")).First(&entry).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, entry)
}
