package controllers

import (
	"fmt"
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/index"
	"github.com/dmnyu/go-medialog/models"
	"github.com/dmnyu/go-medialog/shared"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetAccessions(c *gin.Context) {
	accessions := database.FindAccessions()
	c.HTML(http.StatusOK, "accessions-index.html", gin.H{
		"title":      "go-medialog -- accessions",
		"accessions": accessions,
	})
}

func GetAccession(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	//pagination
	var p = 1
	page := c.Request.URL.Query()["page"]

	if len(page) > 0 {
		p, err = strconv.Atoi(page[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		if p == 0 {
			p = 1
		}
	}

	pagination := shared.Pagination{
		Limit: 10,
		Page:  p,
		Sort:  "id asc",
	}

	accession, err := database.FindAccession(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	repository, err := database.FindRepository(accession.RepositoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	//reposourceIDs := *database.GetResourceIdentifiers()

	resource, err := database.FindResource(accession.ResourceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	entries, err := index.SearchByAccessionID(int(accession.ID), pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "accessions-show.html", gin.H{
		"accession":  accession,
		"repository": repository,
		"resource":   resource,
		"entries":    entries,
		"mediaTypes": models.MediaNames,
		"page":       p,
	})
}

func CreateAccession(c *gin.Context) {
	var input = models.CreateAspaceObject{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	repository, err := database.FindRepository(input.RepositoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	asAccession, err := FindAspaceAccession(repository.AspaceID, input.AccessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	accession := models.Accession{
		Model:        gorm.Model{},
		AspaceID:     input.AccessionID,
		RepositoryID: input.RepositoryID,
		ResourceID:   input.ResourceID,
		Title:        asAccession.Title,
		Identifiers:  asAccession.MergeIDs("."),
		State:        "",
	}

	id, err := database.InsertAccession(accession)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/accessions/%d/show?page=1", id))

}

func PreviewAccession(c *gin.Context) {
	var input = models.CreateAspaceObject{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	resource, err := database.FindResource(input.ResourceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	repository, err := database.FindRepository(input.RepositoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	accession, err := FindAspaceAccession(repository.AspaceID, input.AccessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.HTML(http.StatusOK, "accessions-preview.html", gin.H{
		"title":       "go-medialog - accessions",
		"repository":  repository,
		"resource":    resource,
		"accession":   accession,
		"identifiers": accession.MergeIDs("."),
		"input":       input,
	})

}

func DeleteAccession(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accession, err := database.FindAccession(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err = database.DeleteAccession(&accession); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/resources/%d/show", accession.ResourceID))
}
