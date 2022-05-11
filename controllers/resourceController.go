package controllers

import (
	"fmt"
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetResources(c *gin.Context) {
	resources, err := database.FindResources()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "resources-index.html", gin.H{
		"title":     "go-medialog - resources",
		"resources": resources,
	})
}

func GetResource(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resource, err := database.FindResource(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	repository, err := database.FindRepository(resource.RepositoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accessions, err := database.FindAccessionsByResourceID(resource.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "resources-show.html", gin.H{
		"title":      "go-medialog - resources",
		"repository": repository,
		"resource":   resource,
		"accessions": accessions,
	})
}

func PreviewResource(c *gin.Context) {
	log.Println("PREVIEW")
	var input = models.CreateAspaceObject{}
	if err := c.Bind(&input); err != nil {
		log.Println("[ERROR] %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
	}

	repository, err := database.FindRepository(input.RepositoryID)
	if err != nil {
		log.Println("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resource, err := FindAspaceResource(repository.AspaceID, input.ResourceID)
	if err != nil {
		log.Println("[ERROR] [ASPACE] %s", err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	identifiers := resource.MergeIDs(".")

	c.HTML(http.StatusOK, "resources-preview.html", gin.H{
		"title":       "go-medialog - resources",
		"repository":  repository,
		"resource":    resource,
		"identifiers": identifiers,
		"input":       input,
	})

}

func CreateResource(c *gin.Context) {

	var input = models.CreateAspaceObject{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	repository, err := database.FindRepository(input.RepositoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	asResource, err := FindAspaceResource(repository.AspaceID, input.ResourceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	asURI := asResource.URI
	asIDSplit := strings.Split(asURI, "/")
	asID, err := strconv.Atoi(asIDSplit[len(asIDSplit)-1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	identifiers := asResource.MergeIDs(".")

	resource := models.Resource{
		Model:        gorm.Model{},
		AspaceID:     asID,
		RepositoryID: input.RepositoryID,
		Title:        asResource.Title,
		Identifiers:  identifiers,
	}

	id, err := database.InsertResource(resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/resources/%d/show", id))

}

func DeleteResource(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	resource, err := database.FindResource(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DeleteResource(&resource); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/repositories/%d/show", resource.RepositoryID))
}
