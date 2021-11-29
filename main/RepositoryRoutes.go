package main

import (
	"github.com/dmnyu/go-medialog/controllers"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var repoRoutes = router.Group("/repositories")

func loadRepositoryRoutes() {

	repoRoutes.GET("", func(c *gin.Context){
		var repositories = controllers.GetRepositories()
		c.HTML(http.StatusOK, "repositories-index.html", gin.H{
			"title": "go-medialog - repositories",
			"repositories": repositories,
		})
	})

	repoRoutes.GET("/view/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		log.Println(id)

		repository, err := controllers.FindRepository(id)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
		}

		c.HTML(http.StatusOK, "repositories-show.html", gin.H{
			"title": "go-medialog - repositories",
			"repository": repository,
		})
	})

	repoRoutes.GET("/new", func(c *gin.Context){
		c.HTML(http.StatusOK, "repositories-create.html", gin.H{
			"title": "go-medialog - create a repository",
		})
	})

	repoRoutes.POST("/create", func(c *gin.Context) {
		var input = models.CreateRepository{}
		if err := c.Bind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		code, err := controllers.CreateRepository(input)
		if err != nil {
			c.JSON(code, err.Error())
		}

		c.Redirect(http.StatusFound, "/repositories")
	})

	repoRoutes.POST("/preview", func(c *gin.Context){
		var input = models.CreateRepository{}
		if err := c.Bind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}


		repository, err := controllers.FindAspaceRepository(input.AspaceID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		c.HTML(http.StatusOK, "repositories-preview.html", gin.H{
			"repository": repository,
			"input": input,
			"title": "go-medialog-repositories -- create",
		})

	})
}
