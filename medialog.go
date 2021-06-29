package main

import (
	"github.com/dmnyu/go-medialog/controllers"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	models.ConnectDataBase()

	//General Routes
	router.GET("/", func(context *gin.Context) {
		context.HTML(200,
			"index.html",
			gin.H{
				"title": "go-medialog",
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

	//Start the router
	router.Run()

}
