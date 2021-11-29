package main

import "github.com/dmnyu/go-medialog/controllers"

var apiRoutes = router.Group("/api/v0")
//API Routes

func loadAPIRoutes() {
	loadAccessionsAPIRoutes()
	loadEntryAPIRoutes()
	loadRepoAPIRoutes()
	loadResourceAPIRoutes()
	loadUserAPIRoutes()
}

func loadAccessionsAPIRoutes() {
	apiRoutes.GET("/accessions", controllers.FindAccessions)
	apiRoutes.POST("accessions", controllers.CreateAccessionAPI)
	apiRoutes.POST("/accessions/migrate", controllers.MigrateAccession)
	apiRoutes.GET("/accessions/:id", controllers.FindAccessionAPI)
}

func loadEntryAPIRoutes() {
	router.GET("/api/v0/entries/:id", controllers.FindEntry)
	router.GET("/api/v0/entries", controllers.FindEntries)
	router.POST("/api/v0/entries", controllers.CreateEntry)
}

func loadUserAPIRoutes() {
	router.GET("/api/v0/users", controllers.FindUsers)
	router.POST("/api/v0/users", controllers.CreateUser)
	router.POST("/api/v0/users/validate", controllers.ValidateCredentials)
}

func loadRepoAPIRoutes() {
	router.POST("/api/v0/repositories", controllers.CreateRepositoryAPI)
	router.GET("/api/v0/repositories", controllers.FindRepositories)
	router.GET("/api/v0/repositories/:id", controllers.FindRepositoryAPI)
	router.DELETE("/api/v0/repositories/:id", controllers.DeleteRepository)
}

func loadResourceAPIRoutes() {
	router.POST("/api/v0/resources", controllers.CreateResource)
	router.GET("/api/v0/resources", controllers.FindResources)
	router.GET("/api/v0/resources/:id", controllers.FindResource)
	router.DELETE("/api/v0/resources/:id", controllers.DeleteResource)
}