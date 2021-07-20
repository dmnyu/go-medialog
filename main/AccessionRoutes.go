package main

import (
	"github.com/dmnyu/go-medialog/controllers"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var accRoutes = router.Group("/accessions")

func loadAccessions() {
	accRoutes.GET("", func(c *gin.Context) {

		var accessions = controllers.FindRecentAccessions()

		c.HTML(http.StatusOK, "accessions-index.html", gin.H {
			"accessions": accessions,
			"title":   "go-medialog - accessions",
		})
	})

	accRoutes.GET("/new", func(c *gin.Context) {

		c.HTML(http.StatusOK, "accessions-create.html", gin.H {
			"title":   "go-medialog-accessions -- create",
		})
	})

	accRoutes.POST("/preview", func(c *gin.Context) {
		var input = models.CreateAccession{}
		if err := c.Bind(&input); err !=nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		accession, err := controllers.FindAspaceAccession(input.AspaceRepositoryID, input.AspaceID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		identifiers := accession.MergeIDS()

		c.HTML(http.StatusOK, "accessions-preview.html", gin.H {
			"input": input,
			"accession": accession,
			"identifiers": identifiers,
			"title":   "go-medialog-accessions -- create",
		})
	})

	accRoutes.POST("/create", func(c *gin.Context) {

		var input = models.CreateAccession{}

		if err := c.Bind(&input); err !=nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		code, err := controllers.InsertAccession(input)
		if err != nil {
			c.JSON(code, err.Error())
			return
		}
		c.Redirect(http.StatusFound, "/accessions")
	})

	accRoutes.GET("/show/:id",func(c *gin.Context) {
		accession, err  := controllers.FindAccession(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.HTML(http.StatusOK, "accessions-show.html", gin.H {
			"accession": accession,
			"title":   "go-medialog - accessions",
		})
	})
}
