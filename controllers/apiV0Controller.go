package controllers

import (
	"fmt"
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAccessionAPI(c *gin.Context) {
	var accession = models.Accession{}
	if err := c.Bind(&accession); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	repository, err := database.FindRepository(accession.RepositoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	asAccession, err := FindAspaceAccession(repository.AspaceID, accession.AspaceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accession.Identifiers = asAccession.MergeIDs()
	accession.Title = asAccession.Title

	_, err = database.InsertAccession(accession)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, accession)
}

func CreateResourceAPI(c *gin.Context) {
	var resource = models.Resource{}
	if err := c.Bind(&resource); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	repository, err := database.FindRepository(resource.RepositoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	asResource, err := FindAspaceResource(repository.AspaceID, resource.AspaceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("repository %d resource %d does not exist", repository.AspaceID, resource.AspaceID))
		return
	}

	identifiers := asResource.MergeIDs(".")
	resource.Identifiers = identifiers
	resource.Title = asResource.Title

	_, err = database.InsertResource(resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, resource)
}

func CreateOpticalDiscAPI(c *gin.Context) {}
