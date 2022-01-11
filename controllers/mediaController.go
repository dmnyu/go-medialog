package controllers

import (
	"github.com/dmnyu/go-medialog/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateMedia(c *gin.Context) {
	var input = database.MediaCore{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	switch input.ModelID {
	case 0:
		createOpticalDisc(c)
	case 1:
		createHardDiskDrive(c)
	}
}

func createOpticalDisc(c *gin.Context) {
	var mediaObject = database.MediaCore{}
	if err := c.Bind(&mediaObject); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	repository, err := database.FindRepository(mediaObject.RepositoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	resource, err := database.FindResource(mediaObject.ResourceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	accession, err := database.FindAccession(mediaObject.AccessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	subtypes := database.SubTypes[0]

	c.HTML(http.StatusOK, "optical-create.html", gin.H{
		"repository":  repository,
		"resource":    resource,
		"accession":   accession,
		"mediaObject": mediaObject,
		"subtypes":    subtypes,
	})

}

func createHardDiskDrive(c *gin.Context) {
	c.JSON(http.StatusOK, "Not Implemented")
}
