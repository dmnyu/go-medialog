package controllers

import (
	"github.com/dmnyu/go-medialog/index"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewSearch(c *gin.Context) {
	c.HTML(http.StatusOK, "search-index.html", gin.H{})
}

type IndexQuery struct {
	Query string `json:"query" form:"query"`
}

func SearchIndex(c *gin.Context) {
	var query = IndexQuery{}
	if err := c.Bind(&query); err != nil {
		log.Printf("[ERROR] [APP] %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("[DEBUG] [CONTROLLER] %s", query.Query)
	entries, err := index.KeywordSearch(query.Query)
	if err != nil {
		log.Printf("[ERROR] [INDEX] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "search-results.html", gin.H{
		"entries": entries,
	})
}
