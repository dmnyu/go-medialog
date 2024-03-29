package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetResources(c *gin.Context) {
	session := sessions.Default(c)
	auth := session.Get("auth-key")
	fmt.Println(auth)
	if len(fmt.Sprintf("%v", auth)) != 32 {
		fmt.Println(session)
		c.JSON(http.StatusForbidden, "Must Authenticate")
		return
	}

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

	relatedAccessions, err := GetAccessionListForResource(repository.AspaceID, resource.AspaceID)
	if err != nil {
		log.Printf("[ERROR] [ASPACE] %s", strings.ReplaceAll(err.Error(), "\n", ""))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
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
	/*
		pagination := shared.Pagination{
			Limit: 10,
			Page:  p,
			Sort:  "id asc",
		}


			entries, err := index.FindByType(int(resource.ID), index.Accession, pagination)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
	*/

	c.HTML(http.StatusOK, "resources-show.html", gin.H{
		"title":             "go-medialog - resources",
		"repository":        repository,
		"resource":          resource,
		"relatedAccessions": relatedAccessions,
		"accessions":        accessions,
		"entries":           nil,
	})
}

func AddResource(c *gin.Context) {
	repo_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("[ERROR] [MEDIALOG] %s", err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	repository, err := database.FindRepository(repo_id)
	if err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	aspaceResources, err := GetResourceList(repository.AspaceID)
	if err != nil {
		log.Printf("[ERROR] [ASPACE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "resources-new.html", gin.H{
		"repository":       repository,
		"aspace_resources": aspaceResources,
	})
}

func PreviewResource(c *gin.Context) {
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
