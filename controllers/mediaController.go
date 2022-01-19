package controllers

import (
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/index"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateMedia(c *gin.Context) {
	var input = database.MediaEntry{}
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

func newOpticalDisc(c *gin.Context) {
	var entry = database.MediaEntry{}
	if err := c.Bind(&entry); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	repository, err := database.FindRepository(entry.RepositoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	resource, err := database.FindResource(entry.ResourceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	accession, err := database.FindAccession(entry.AccessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	subtypes := database.SubTypes[int(entry.ModelID)]

	c.HTML(http.StatusOK, "optical-create.html", gin.H{
		"repository": repository,
		"resource":   resource,
		"accession":  accession,
		"entry":      entry,
		"subtypes":   subtypes,
	})

}

func CreateOpticalDisc(c *gin.Context) {
	var o = database.MediaOpticalDisc{}
	if err := c.Bind(&o); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	mediaID, err := database.GetNextMediaIDForResource(o.ResourceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, mediaID)
	}

	o.MediaID = mediaID
	err = database.InsertOpticalDisc(&o)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	entry := o.GetMediaEntry()
	entry.ObjectID = int(o.ID)
	entry.MediaID = mediaID

	msg, err := index.AddToIndex(entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(200, msg)

}

func deleteOpticalDisc(c *gin.Context, entry database.MediaEntry) {
	//delete the disc
	err := database.DeleteOpticalDisc(entry.ObjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func newHardDiskDrive(c *gin.Context) {
	c.JSON(http.StatusOK, "Not Implemented")
}
