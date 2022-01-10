package main

import (
	"flag"
	"fmt"
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/routes"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
	"time"
)

var migrate bool

func init() {
	flag.BoolVar(&migrate, "migrate", false, "")
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

func getRepoCode(i int) string {
	switch i {
	case 2:
		return "tamwag"
	case 3:
		return "fales"
	case 6:
		return "archives"
	case 100:
		return "abudhabi"
	}
	return "unkown"
}

var router = gin.Default()

func main() {
	flag.Parse()
	if migrate == true {
		database.MigrateDatabase()
		os.Exit(0)
	}
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
		"getRepoCode":  getRepoCode,
	})

	router.LoadHTMLGlob("templates/**/*.html")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	//Load Application Routes
	routes.LoadRoutes(router)

	//Index
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "go-medialog",
		})
	})

	//Start the router
	database.ConnectDatabase()
	router.Run()

}
