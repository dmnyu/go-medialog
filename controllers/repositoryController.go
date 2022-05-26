package controllers

import (
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/models"
	"github.com/dmnyu/go-medialog/shared"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func CreateRepository(c *gin.Context) {
	var input = models.Repository{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	asRepository, err := FindAspaceRepository(input.AspaceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	repository := models.Repository{
		Model:    gorm.Model{},
		AspaceID: input.AspaceID,
		Slug:     asRepository.Slug,
		Name:     asRepository.Name,
	}

	if err := database.InsertRepository(repository); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusFound, "/repositories")

}

func GetRepository(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	//pagination
	var p int = 1
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

	// get the repository
	repository, err := database.FindRepository(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}

	//get the resources
	resources, err := database.FindResourcesByRepoID(id, pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	//render the page
	c.HTML(http.StatusOK, "repositories-show.html", gin.H{
		"title":      "go-medialog - repositories",
		"repository": repository,
		"resources":  resources,
		"page":       p,
	})
}

func GetRepositories(c *gin.Context) {
	repositories := database.FindRepositories()
	c.HTML(http.StatusOK, "repositories-index.html", gin.H{
		"title":        "go-medialog - repositories",
		"repositories": repositories,
	})
}

func AddRepository(c *gin.Context) {
	repositories, err := GetASpaceRepositories()
	if err != nil {
		log.Printf("[ERROR] [ASPACE] %s", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "repositories-new.html", gin.H{
		"repositories": repositories,
	})
}

func PreviewRepository(c *gin.Context) {
	var input = models.Repository{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	repository, err := FindAspaceRepository(input.AspaceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.HTML(http.StatusOK, "repositories-preview.html", gin.H{
		"repository": repository,
		"input":      input,
		"title":      "go-medialog-repositories -- create",
	})
}
