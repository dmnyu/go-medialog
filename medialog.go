package main

import (
	"fmt"
	"github.com/dmnyu/go-medialog/controllers"
	"github.com/dmnyu/go-medialog/models"
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

func main() {

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
		"getRepoCode":  getRepoCode,
	})

	router.LoadHTMLGlob("templates/*")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	models.ConnectDataBase()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//Index
	router.GET("/", func(c *gin.Context) {

		var entries = controllers.FindRecent()

		c.HTML(http.StatusOK, "index.html", gin.H{
			"entries": entries,
			"title":   "go-medialog",
		})
	})

	//Entry Routes
	router.GET("/entries/:id", controllers.FindEntry)
	router.GET("/entries", controllers.FindEntries)
	router.POST("/entries", controllers.CreateEntry)

	//User Routes
	router.GET("/users", controllers.FindUsers)
	router.POST("/users", controllers.CreateUser)
	router.POST("/users/validate", controllers.ValidateCredentials)

	//Repository Routes
	router.POST("/repositories", controllers.CreateRepository)
	router.GET("/repositories", controllers.FindRepositories)
	router.GET("/repositories/:id", controllers.FindRepository)
	router.DELETE("/repositories/:id", controllers.DeleteRepository)

	//Resource Routes
	router.POST("/resources", controllers.CreateResource)
	router.GET("/resources", controllers.FindResources)
	router.GET("/resources/:id", controllers.FindResource)
	router.DELETE("/resources/:id", controllers.DeleteResource)

	//Accession Routes
	router.GET("/accessions", controllers.FindAccessions)
	router.POST("/accessions", controllers.CreateAccession)
	router.GET("/accessions/:id", controllers.FindAccession)

	models.MigrateDatabase()
	//Start the router
	s.ListenAndServe()
	router.Run()

}
