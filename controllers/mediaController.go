package controllers

import (
	"github.com/dmnyu/go-medialog/index"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewMedia(c *gin.Context) {
	var input = models.MediaEntry{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	switch input.ModelID {
	case models.OpticalDisc:
		newOpticalDisc(c, input)
	case models.HardDiskDrive:
		newHardDiskDrive(c)
	default:
		c.JSON(http.StatusBadRequest, "Mediatype not supported")
	}
}

func newHardDiskDrive(c *gin.Context) {
	c.JSON(http.StatusOK, "Not Implemented")
}

func DeleteMedia(c *gin.Context) {
	docID := c.Param("id")
	entry, err := index.FindDoc(docID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	switch entry.ModelID {
	case models.OpticalDisc:
		deleteOpticalDisc(c, docID, entry)
	default:
		c.JSON(http.StatusInternalServerError, "Not Implemented Yet")

	}
}

func ShowMedia(c *gin.Context) {
	docID := c.Param("id")
	entry, err := index.FindDoc(docID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	switch entry.ModelID {
	case models.OpticalDisc:
		showOpticalDisc(c, entry)
	default:
		c.JSON(http.StatusInternalServerError, "Not Implemented Yet")

	}
}
