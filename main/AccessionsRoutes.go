package main

import (
	"github.com/dmnyu/go-medialog/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func loadAccessions() {
	router.GET("/accessions", func(c *gin.Context) {

		var accessions = controllers.FindRecentAccessions()
		log.Println(len(accessions))
		c.HTML(http.StatusOK, "accessions-index.html", gin.H {
			"accessions": accessions,
			"title":   "go-medialog-accessions",
		})
	})
}
