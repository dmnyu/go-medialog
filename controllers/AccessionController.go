package controllers

import (
	"github.com/dmnyu/go-medialog/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func InsertAccession(input db.CreateAspaceObject) (int, error) {

	asAccession, err := FindAspaceAccession(input.RepositoryID, input.ObjectID)
	if err != nil {
		return http.StatusBadRequest, err
	}

	accession := db.Accession{
		Model:        gorm.Model{},
		AspaceID:     input.ObjectID,
		RepositoryID: input.RepositoryID,
		Title:        asAccession.Title,
		Identifiers:  asAccession.MergeIDs(),
		State:        "not_started",
		CreatedBy:    0,
		UpdatedBy:    0,
	}

	if err := db.DB.Create(&accession).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func CreateAccessionAPI(c *gin.Context) {
	var input = db.CreateAspaceObject{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	code, err := InsertAccession(input)
	if err != nil {
		c.JSON(code, err.Error())
		return
	}

	c.JSON(code, nil)
}

func MigrateAccession(c *gin.Context) {
	var accession db.Accession
	if err := c.ShouldBindJSON(&accession); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&accession).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, accession)
}

func FindAccessions(c *gin.Context) {
	var accessions []db.Accession

	if err := db.DB.Find(&accessions).Error; err != nil {
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

func FindAccession(c *gin.Context) (db.Accession, error) {
	var accession db.Accession
	if err := db.DB.Where("id = ?", c.Param("id")).First(&accession).Error; err != nil {
		log.Println(err)
		return accession, err
	}

	return accession, nil
}

func DeleteAccession(c *gin.Context) {
	var accession = db.Accession{}

	if err := db.DB.Where("id = ?", c.Param("id")).First(&accession).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
	}

	if err := db.DB.Delete(&accession); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
	}

	c.JSON(http.StatusOK, gin.H{"result": true})
}

func FindRecentAccessions() []db.Accession {
	accessions := []db.Accession{}
	db.DB.Limit(20).Find(&accessions)
	return accessions
}
