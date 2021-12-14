package controllers

import (
	"fmt"
	"github.com/dmnyu/go-medialog/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetResources(c *gin.Context) {
	resources := database.FindResources()
	log.Println(resources)
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
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.HTML(http.StatusOK, "resources-show.html", gin.H{
		"title":      "go-medialog - resources",
		"repository": repository,
		"resource":   resource,
	})
}

func PreviewResource(c *gin.Context) {
	var input = database.CreateAspaceObject{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	resource, err := FindAspaceResource(input.RepositoryID, input.ObjectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	identifiers := resource.MergeIDs()

	repository, err := database.FindRepository(input.RepositoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "resources-preview.html", gin.H{
		"title":       "go-medialog - resources",
		"repository":  repository,
		"resource":    resource,
		"identifiers": identifiers,
		"input":       input,
	})
}

func CreateResource(c *gin.Context) {

	var input = database.CreateAspaceObject{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	asResource, err := FindAspaceResource(input.RepositoryID, input.ObjectID)
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

	identifiers := asResource.MergeIDs()

	resource := database.Resource{
		Model:                     gorm.Model{},
		AspaceResourceID:          asID,
		RepositoryID:              input.RepositoryID,
		AspaceResourceTitle:       asResource.Title,
		AspaceResourceIdentifiers: identifiers,
	}

	id, err := database.InsertResource(resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/resources/%d/show", id))

}
