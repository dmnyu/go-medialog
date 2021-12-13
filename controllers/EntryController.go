package controllers

import (
	"github.com/dmnyu/go-medialog/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateEntry(c *gin.Context) {
	input := db.Entry{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := db.DB.Create(&input).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, input.ID)
}

func FindEntries(c *gin.Context) {
	entries := []db.Entry{}
	if err := db.DB.Find(&entries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, entries)
}

func FindEntry(c *gin.Context) {
	entry := db.Entry{}
	if err := db.DB.Where("id = ?", c.Param("id")).First(&entry).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, entry)
}
