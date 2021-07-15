package controllers

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAccession(c *gin.Context) {
	var accession models.Accession
	if err := c.ShouldBindJSON(&accession); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&accession).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, accession)
}

func FindAccessions(c *gin.Context) {
	var accessions []models.Accession

	if err  := models.DB.Find(&accessions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, accessions)
}

func FindAccession(c *gin.Context) {
	var accession models.Accession
	if err := models.DB.Where("id = ?", c.Param("id")).First(&accession).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	}

	c.JSON(http.StatusOK, accession)
}

func DeleteAccession(c *gin.Context) {
	var accession = models.Accession{}

	if err := models.DB.Where("id = ?", c.Param("id")).First(&accession).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
	}

	if err := models.DB.Delete(&accession); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
	}

	c.JSON(http.StatusOK, gin.H{"result": true})
}

func FindRecentAccessions() []models.Accession {
	accessions := []models.Accession{}
	models.DB.Limit(20).Find(&accessions)
	return accessions
}
