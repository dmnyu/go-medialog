package controllers

import (
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/index"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"github.com/nyudlts/bytemath"
	"log"
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
		showOpticalDisc(c, entry)
	default:
		c.JSON(http.StatusInternalServerError, "Not Implemented Yet")

	}
}

func newOpticalDisc(c *gin.Context, entry models.MediaEntry) {

	c.HTML(http.StatusOK, "optical-create.html", gin.H{
		"entry":    entry,
		"subtypes": models.OpticalSubTypes,
		"units":    models.MediaUnit,
	})

}

func showOpticalDisc(c *gin.Context, entry models.MediaEntry) {

}

func deleteOpticalDisc(c *gin.Context, entry models.MediaEntry) {}

func CreateOpticalDisc(c *gin.Context) {
	var optical = models.MediaOpticalDisc{}
	if err := c.Bind(&optical); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	optical.SizeInBytes = bytemath.ConvertToBytes(float64(optical.StockSize), bytemath.MB)

	if err := database.InsertOpticalDisc(&optical); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	entry := optical.GetMediaEntry()
	resp, err := index.AddToIndex(entry)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, resp)

}
