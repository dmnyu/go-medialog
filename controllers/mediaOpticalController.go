package controllers

import (
	"fmt"
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/index"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"github.com/nyudlts/bytemath"
	"log"
	"net/http"
)

func newOpticalDisc(c *gin.Context, entry models.MediaEntry) {

	c.HTML(http.StatusOK, "optical-create.html", gin.H{
		"entry":    entry,
		"subtypes": models.OpticalSubtypes,
		"units":    models.MediaUnit,
	})

}

func CreateOpticalDisc(c *gin.Context) {
	var optical = models.MediaOpticalDisc{}
	if err := c.Bind(&optical); err != nil {
		log.Printf("[ERROR] [MEDIALOG] cannot create hard drive struct from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
	}

	optical.SizeInBytes = int64(bytemath.ConvertToBytes(float64(optical.StockSize), models.ByteMathSuffix[optical.StockUnit]))

	//get the next ID if the media ID is 0
	if optical.MediaID == 0 {
		nextMediaID, err := index.FindNextMediaIDInResource(optical.ResourceID)
		if err != nil {
			log.Printf("[ERROR] [DATABASE] %s", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		optical.MediaID = *nextMediaID
	}

	if err := database.InsertOpticalDisc(&optical); err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	entry := optical.GetMediaEntry()
	resp, err := index.AddToIndex(entry)

	if err != nil {
		log.Printf("[ERROR] [INDEX] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("[INFO] [INDEX] %s", resp)

	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/accessions/%d/show", optical.AccessionID))

}

func showOpticalDisc(c *gin.Context, entry *models.MediaEntry) {
	c.HTML(http.StatusOK, "optical-show.html", gin.H{
		"optical": entry,
	})
}

func deleteOpticalDisc(c *gin.Context, docID string, entry *models.MediaEntry) {

	if err := index.DeleteFromIndex(docID); err != nil {
		log.Printf("[ERROR] [INDEX] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := database.DeleteOpticalDisc(entry.DatabaseID); err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/accessions/%d/show", entry.AccessionID))
}
