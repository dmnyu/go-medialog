package controllers

import (
	"github.com/dmnyu/go-medialog/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateRepository(c *gin.Context) {
	var input = database.Repository{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	asRepository, err := FindAspaceRepository(input.AspaceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	repository := database.Repository{
		Model:    gorm.Model{},
		AspaceID: input.AspaceID,
		Slug:     asRepository.Slug,
		Name:     asRepository.Name,
	}

	if err := database.InsertRepository(repository); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.Redirect(http.StatusFound, "/repositories")

}

func GetRepository(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	repository, err := database.FindRepository(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}

	/*
		resources, err := controllers.FindResourcesByRepository(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	*/

	c.HTML(http.StatusOK, "repositories-show.html", gin.H{
		"title":      "go-medialog - repositories",
		"repository": repository,
		//"resources":  resources,
	})
}

func GetRepositories(c *gin.Context) {
	repositories := database.FindRepositories()
	c.HTML(http.StatusOK, "repositories-index.html", gin.H{
		"title":        "go-medialog - repositories",
		"repositories": repositories,
	})
}

func PreviewRepository(c *gin.Context) {
	var input = database.Repository{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	repository, err := FindAspaceRepository(input.AspaceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.HTML(http.StatusOK, "repositories-preview.html", gin.H{
		"repository": repository,
		"input":      input,
		"title":      "go-medialog-repositories -- create",
	})
}
