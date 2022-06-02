package controllers

import (
	"fmt"
	"github.com/dmnyu/go-medialog/database"
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
		showOpticalDisc(c, entry, docID)
	default:
		c.JSON(http.StatusInternalServerError, "Not Implemented Yet")
	}
}

func EditMedia(c *gin.Context) {
	docID := c.Param("id")
	entry, err := index.FindDoc(docID)
	if err != nil {
		log.Printf("[ERROR] [INDEX] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	switch entry.ModelID {
	case models.OpticalDisc:
		editOpticalDisc(c, entry, docID)
	default:
		log.Printf("[ERROR] [MEDIALOG] no route available for type %d", models.OpticalDisc)
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("no route available for type %d", models.OpticalDisc))
		return
	}
}

type RepResAcc struct {
	RepositoryID         int
	RepositoryName       string
	ResourceID           int
	ResourceIdentifiers  string
	ResourceTitle        string
	AccessionID          int
	AccessionIdentifiers string
	AccessionTitle       string
}

func GetRepResAcc(m *models.MediaEntry) (*RepResAcc, error) {
	repository, err := database.FindRepository(m.RepositoryID)
	if err != nil {
		return nil, err
	}

	resource, err := database.FindResource(m.ResourceID)
	if err != nil {
		return nil, err
	}

	accession, err := database.FindAccession(m.AccessionID)
	if err != nil {
		return nil, err
	}

	return &RepResAcc{
		RepositoryID:         int(repository.ID),
		RepositoryName:       repository.Name,
		ResourceID:           int(resource.ID),
		ResourceIdentifiers:  resource.Identifiers,
		ResourceTitle:        resource.Title,
		AccessionID:          int(accession.ID),
		AccessionIdentifiers: accession.Identifiers,
		AccessionTitle:       accession.Title,
	}, nil
}
