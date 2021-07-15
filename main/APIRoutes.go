package main

import "github.com/dmnyu/go-medialog/controllers"

//API Routes
func loadAPI() {
	router.GET("/api/v0/accessions/", controllers.FindAccessions)
	router.POST("/api/v0/accessions", controllers.CreateAccession)
	router.GET("/api/v0/accessions/:id", controllers.FindAccession)
}