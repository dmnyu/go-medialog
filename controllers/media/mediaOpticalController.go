package mediacontrollers

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

	repository, err := database.FindRepository(entry.RepositoryID)
	if err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resource, err := database.FindResource(entry.ResourceID)
	if err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	accession, err := database.FindAccession(entry.AccessionID)
	if err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "optical-create.html", gin.H{
		"repository":   repository,
		"resource":     resource,
		"accession":    accession,
		"subtypes":     models.OpticalSubtypes,
		"units":        models.MediaUnit,
		"contentTypes": models.OpticalContentTypes,
	})

}

func CreateOpticalDisc(c *gin.Context) {
	var optical = models.MediaOpticalDisc{}
	if err := c.Bind(&optical); err != nil {
		log.Printf("[ERROR] [MEDIALOG] cannot create optical struct from request: %s", err.Error())
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
	resp, err := index.AddToIndex(entry, nil)

	if err != nil {
		log.Printf("[ERROR] [INDEX] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("[INFO] [INDEX] %s", resp)

	c.Redirect(http.StatusFound, fmt.Sprintf("/accessions/%d/show", optical.AccessionID))

}

func showOpticalDisc(c *gin.Context, entry *models.MediaEntry, docID string) {
	repResAcc, err := GetRepResAcc(entry)
	if err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	optical, err := database.FindOpticalDisc(int(entry.DatabaseID))
	if err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.HTML(http.StatusOK, "optical-show.html", gin.H{
		"optical":   optical,
		"docID":     docID,
		"repResAcc": repResAcc,
	})
}

func editOpticalDisc(c *gin.Context, entry *models.MediaEntry, docID string) {
	repResAcc, err := GetRepResAcc(entry)
	if err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	optical, err := database.FindOpticalDisc(int(entry.DatabaseID))
	if err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.HTML(http.StatusOK, "optical-edit.html", gin.H{
		"optical":      optical,
		"docID":        docID,
		"repResAcc":    repResAcc,
		"subtypes":     models.OpticalSubtypes,
		"units":        models.MediaUnit,
		"contentTypes": models.OpticalContentTypes,
	})
}

func UpdateOpticalDisc(c *gin.Context) {

	var optical = models.MediaOpticalDisc{}

	if err := c.Bind(&optical); err != nil {
		log.Printf("[ERROR] [MEDIALOG] cannot create optical struct from request: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//look up the doc in the index
	docID := c.Param("id")
	doc, err := index.FindDoc(docID)
	if err != nil {
		log.Printf("[ERROR] [INDEX] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	dbID := doc.DatabaseID
	optical.ID = dbID

	if err := database.UpdateOpticalDisc(&optical); err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	deleteMSG, err := index.DeleteFromIndex(docID)
	if err != nil {
		log.Printf("[ERROR] [INDEX] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	log.Printf("[INFO] [INDEX] %v", deleteMSG.Json)

	createResponse, err := index.AddToIndex(optical.GetMediaEntry(), &docID)
	if err != nil {
		log.Printf("[ERROR] [INDEX] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	log.Printf("[INFO] [INDEX] %v", createResponse.Json)

	c.Redirect(http.StatusFound, fmt.Sprintf("/media/%s/show", createResponse.ID))

}

func deleteOpticalDisc(c *gin.Context, docID string, entry *models.MediaEntry) {

	delete, err := index.DeleteFromIndex(docID)
	if err != nil {
		log.Printf("[ERROR] [INDEX] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Printf("[INFO] [INDEX] %s", delete)

	if err := database.DeleteOpticalDisc(entry.DatabaseID); err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/accessions/%d/show", entry.AccessionID))
}
