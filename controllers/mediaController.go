package controllers

import (
	"github.com/dmnyu/go-medialog/index"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateMedia(c *gin.Context) {
	var input = models.MediaEntry{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	switch input.ModelID {
	case 0:
		newOpticalDisc(c)
	case 1:
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
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	msg, err := index.DeleteFromIndex(docID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Println("[INFO] ", msg)

	switch entry.ModelID {
	case models.OpticalDisc:
		deleteOpticalDisc(c, entry)
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
		showOptical(c, entry)
	default:
		c.JSON(http.StatusInternalServerError, "Not Implemented Yet")

	}
}
