package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", func(context *gin.Context) {
		context.HTML(
			200,
			"index.html",
			gin.H{
				"title": "Medialog",
			},
		)
	})
	router.Run()
}
