package controllers

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"github.com/nyudlts/go-aspace"
	"net/http"
)

var (
	client *aspace.ASClient
	err error
)

func init() {
	client, err = aspace.NewClient("go-aspace.yml", "dev", 20)
	if err != nil {
		panic(err)
	}
}

func CreateAccession(c *gin.Context) {
	var input models.Accession
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, input)
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
