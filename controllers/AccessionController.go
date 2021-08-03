package controllers

import (
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func InsertAccession(input models.CreateAccession) (int, error) {

	asAccession, err := FindAspaceAccession(input.AspaceRepositoryID, input.AspaceID)
	if err !=nil {
		return http.StatusBadRequest, err
	}

	accession := models.Accession{
		Model:        	gorm.Model{},
		AspaceID:     	input.AspaceID,
		RepositoryID: 	input.AspaceRepositoryID,
		ResourceID:   	asAccession.GetParentResourceID(),
		Title:			asAccession.Title,
		Identifiers:  	asAccession.MergeIDs(),
		State:        	"not_started",
		CreatedBy:    	0,
		UpdatedBy:    	0,
	}

	if err := models.DB.Create(&accession).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func CreateAccessionAPI(c *gin.Context) {
	var input = models.CreateAccession{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	code,err := InsertAccession(input);
	if err != nil {
		c.JSON(code, err.Error())
		return
	}

	c.JSON(code, nil)
}

func MigrateAccession(c *gin.Context) {
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

func FindAccessionAPI(c *gin.Context) {
	accession, err := FindAccession(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, accession)
}

func FindAccession(c *gin.Context) (models.Accession, error) {
	var accession models.Accession
	if err := models.DB.Where("id = ?", c.Param("id")).First(&accession).Error; err != nil {
		log.Println(err)
		return accession, err
	}

	return accession, nil
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
