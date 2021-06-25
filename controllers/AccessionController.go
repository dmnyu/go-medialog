package controllers

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAccession(c *gin.Context) {
	var input models.CreateAccession
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accession := models.Accession{
		ID:           0,
		RepositoryID: input.RepositoryID,
		AccessionID:  input.AccessionID,
	}

	models.DB.Create(&accession)
	c.JSON(http.StatusOK, accession)
}

func FindAccessions(c *gin.Context) {
	var accessions []models.Accession
	models.DB.Find(&accessions)
	c.JSON(http.StatusOK, accessions)
}

func FindAccession(c *gin.Context) {
	var accession models.Accession
	if err := models.DB.Where("id = ?", c.Param("id")).First(&accession).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//call the go-aspace api
	accessionRecord, err := client.GetAccession(accession.RepositoryID, accession.AccessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Did not receive a response from aspace instance")
	}

	c.JSON(http.StatusOK, accessionRecord)

}
