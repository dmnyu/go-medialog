package routes

import (
	"github.com/dmnyu/go-medialog/controllers"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(router *gin.Engine) {
	var repoRoutes = router.Group("/repositories")
	repoRoutes.GET("", func(c *gin.Context) { controllers.GetRepositories(c) })
	repoRoutes.GET("/:id/show", func(c *gin.Context) { controllers.GetRepository(c) })
	repoRoutes.GET("/:id/edit", func(c *gin.Context) { NullRoute() })
	repoRoutes.GET("/:id/delete", func(c *gin.Context) { NullRoute() })
	repoRoutes.POST("/preview", func(c *gin.Context) { controllers.PreviewRepository(c) })
	repoRoutes.POST("/create", func(c *gin.Context) { controllers.CreateRepository(c) })

	var resourceRoutes = router.Group("/resources")
	resourceRoutes.GET("", func(c *gin.Context) { controllers.GetResources(c) })
	resourceRoutes.GET("/:id/show", func(c *gin.Context) { controllers.GetResource(c) })
	resourceRoutes.POST("/preview", func(c *gin.Context) { controllers.PreviewResource(c) })
	resourceRoutes.POST("/create", func(c *gin.Context) { controllers.CreateResource(c) })
	resourceRoutes.GET("/:id/edit", func(c *gin.Context) { NullRoute() })
	resourceRoutes.GET("/:id/delete", func(c *gin.Context) { NullRoute() })

	var accessionRoutes = router.Group("/accessions")
	accessionRoutes.GET("", func(c *gin.Context) { controllers.GetAccessions(c) })
	accessionRoutes.POST("/preview", func(c *gin.Context) { controllers.PreviewAccession(c) })
	accessionRoutes.POST("/create", func(c *gin.Context) { controllers.CreateAccession(c) })
	accessionRoutes.GET("/:id/show", func(c *gin.Context) { controllers.GetAccession(c) })
	//accessionRoutes.GET("/:id/delete", func(c *gin.Context) { controllers.DeleteAccession(c) })

	var mediaRoutes = router.Group("/media")
	mediaRoutes.POST("/create", func(c *gin.Context) { controllers.CreateMedia(c) })
	mediaRoutes.POST("/create/optical", func(c *gin.Context) { controllers.CreateOpticalDisc(c) })
	mediaRoutes.GET("/:id/show", func(c *gin.Context) { controllers.ShowMedia(c) })
	mediaRoutes.GET("/:id/delete", func(c *gin.Context) { controllers.DeleteMedia(c) })
}

func NullRoute() {}
