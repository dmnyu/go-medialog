package routes

import (
	"github.com/dmnyu/go-medialog/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRoutes(router *gin.Engine) {
	var repoRoutes = router.Group("/repositories")
	repoRoutes.GET("", func(c *gin.Context) { controllers.GetRepositories(c) })
	repoRoutes.GET("/:id/show", func(c *gin.Context) { controllers.GetRepository(c) })
	repoRoutes.GET("/:id/edit", func(c *gin.Context) { NullRoute(c) })
	repoRoutes.GET("/:id/delete", func(c *gin.Context) { NullRoute(c) })
	repoRoutes.POST("/preview", func(c *gin.Context) { controllers.PreviewRepository(c) })
	repoRoutes.POST("/create", func(c *gin.Context) { controllers.CreateRepository(c) })

	var resourceRoutes = router.Group("/resources")
	resourceRoutes.GET("", func(c *gin.Context) { controllers.GetResources(c) })
	resourceRoutes.GET("/:id/show", func(c *gin.Context) { controllers.GetResource(c) })
	resourceRoutes.POST("/preview", func(c *gin.Context) { controllers.PreviewResource(c) })
	resourceRoutes.POST("/create", func(c *gin.Context) { controllers.CreateResource(c) })
	resourceRoutes.GET("/:id/edit", func(c *gin.Context) { NullRoute(c) })
	resourceRoutes.GET("/:id/delete", func(c *gin.Context) { controllers.DeleteResource(c) })

	var accessionRoutes = router.Group("/accessions")
	accessionRoutes.GET("", func(c *gin.Context) { controllers.GetAccessions(c) })
	accessionRoutes.POST("/preview", func(c *gin.Context) { controllers.PreviewAccession(c) })
	accessionRoutes.POST("/create", func(c *gin.Context) { controllers.CreateAccession(c) })
	accessionRoutes.GET("/:id/show", func(c *gin.Context) { controllers.GetAccession(c) })
	accessionRoutes.GET("/:id/delete", func(c *gin.Context) { controllers.DeleteAccession(c) })

	var mediaRoutes = router.Group("/media")
	mediaRoutes.POST("/new", func(c *gin.Context) { controllers.NewMedia(c) })
	mediaRoutes.POST("/create/optical", func(c *gin.Context) { controllers.CreateOpticalDisc(c) })
	mediaRoutes.GET("/:id/show", func(c *gin.Context) { controllers.ShowMedia(c) })
	mediaRoutes.GET("/:id/delete", func(c *gin.Context) { controllers.DeleteMedia(c) })

	//API Routes
	var apiV0Routes = router.Group("/api/v0")
	apiV0Routes.POST("/create-optical", func(c *gin.Context) { controllers.CreateOpticalDiscAPI(c) })
	apiV0Routes.POST("/create-resource", func(c *gin.Context) { controllers.CreateResourceAPI(c) })
	apiV0Routes.POST("/create-accession", func(c *gin.Context) { controllers.CreateAccessionAPI(c) })
}

func NullRoute(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, "Unsupported Route")
}
