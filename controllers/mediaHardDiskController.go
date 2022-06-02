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

func newHardDiskDrive(c *gin.Context, entry models.MediaEntry) {

	c.HTML(http.StatusOK, "hard-drive-create.html", gin.H{
		"entry":    entry,
		"subtypes": models.HardDriveSubtypes,
		"units":    models.MediaUnit,
	})

}

func CreateHardDiskDrive(c *gin.Context) {
	var hardDiskDrive = models.MediaHardDrive{}
	if err := c.Bind(&hardDiskDrive); err != nil {
		log.Printf("[ERROR] [MEDIALOG] cannot create hard drive struct from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
	}

	//get the size in bytes
	hardDiskDrive.SizeInBytes = int64(bytemath.ConvertToBytes(float64(hardDiskDrive.StockSize), models.ByteMathSuffix[hardDiskDrive.StockUnit]))

	//get the next media id
	if hardDiskDrive.MediaID == 0 {
		nextMediaID, err := index.FindNextMediaIDInResource(hardDiskDrive.ResourceID)
		if err != nil {
			log.Printf("[ERROR] [INDEX] %s", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		hardDiskDrive.MediaID = *nextMediaID
	}

	//insert into database
	if err := database.InsertHardDiskDrive(&hardDiskDrive); err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	//insert into index
	resp, err := index.AddToIndex(hardDiskDrive.GetMediaEntry(), nil)
	if err != nil {
		log.Printf("[ERROR] [INDEX]] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("[INFO] [INDEX] %s", resp)
	c.Redirect(http.StatusFound, fmt.Sprintf("/accessions/%d/show", hardDiskDrive.AccessionID))

}

func deleteHardDiskDrive(c *gin.Context, docID string, entry *models.MediaEntry) {

	//remove entry from index
	deleteMsg, err := index.DeleteFromIndex(docID)
	if err != nil {
		log.Printf("[ERROR] [INDEX] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("[INFO] [INDEX] %s", deleteMsg)

	//remove entry from database
	if err := database.DeleteHardDiskDrive(entry.DatabaseID); err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("[INFO] [DATABASE] removed hard disk drive %d from database", entry.DatabaseID)

	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/accessions/%d/show", entry.AccessionID))
}
