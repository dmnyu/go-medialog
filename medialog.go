package main

import (
	"fmt"
	"github.com/dmnyu/go-medialog/database"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

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

	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
		"getRepoCode":  getRepoCode,
	})

	router.LoadHTMLGlob("templates/**/*.html")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//Load Application Routes
	loadRoutes(router)

	//Index
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "go-medialog",
		})
	})

	//Start the router
	database.ConnectDataBase()
	database.MigrateDatabase()
	s.ListenAndServe()
	router.Run()

}
