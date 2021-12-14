package controllers

import (
	"github.com/dmnyu/go-medialog/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetResources(c *gin.Context) {
	resources := database.FindResources()
	c.HTML(http.StatusOK, "resources-index.html", gin.H{
		"title":        "go-medialog - resources",
		"repositories": resources,
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

	c.HTML(http.StatusOK, "resources-index.html", gin.H{
		"title":      "go-medialog - resources",
		"repository": repository,
		"resource":   resource,
	})
}

func PreviewResource(c *gin.Context) {}

func CreateResource(c *gin.Context) {}
