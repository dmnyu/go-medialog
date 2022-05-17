package controllers

import (
	"github.com/dmnyu/go-medialog/index"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetEntries(c *gin.Context) { c.HTML(http.StatusOK, "entries-index.html", gin.H{}) }

func NewMedia(c *gin.Context) {
	var entry = models.MediaEntry{}
	if err := c.Bind(&entry); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	switch entry.ModelID {
	case models.OpticalDisc:
		newOpticalDisc(c, entry)
	case models.HardDiskDrive:
		newHardDiskDrive(c, entry)
	default:
		c.JSON(http.StatusBadRequest, "Mediatype not supported")
	}
}

func DeleteMedia(c *gin.Context) {
	docID := c.Param("id")
	entry, err := index.FindDoc(docID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Printf("[ERROR] [INDEX] can not locate document %s in index", docID)
	}

	switch entry.ModelID {
	case models.OpticalDisc:
		deleteOpticalDisc(c, docID, entry)
	case models.HardDiskDrive:
		deleteHardDiskDrive(c, docID, entry)
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
