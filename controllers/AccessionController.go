package controllers

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateAccession(c *gin.Context) {
	var input models.CreateAccession
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aspaceAccession, err  := client.GetAccession(input.AspaceRepositoryID, input.AspaceAccessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	accession := models.Accession{
		Model: gorm.Model{},
		ID: 0,
		AccessionID: input.AspaceAccessionID,
		RepositoryID: input.AspaceRepositoryID,
		Title: aspaceAccession.Title,
		Identifier: getAccessionIndentifierString(aspaceAccession),
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
