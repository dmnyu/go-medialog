package routes

import (
	"github.com/dmnyu/go-medialog/controllers"
	mediacontrollers "github.com/dmnyu/go-medialog/controllers/media"
	"github.com/dmnyu/go-medialog/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	accessionIDs map[int]string
	resourceIDs  map[int]string
)

func LoadRoutes(router *gin.Engine) {

	//Main Index
	router.GET("", func(c *gin.Context) { controllers.Index(c) })

	//Repository Routes
	var repoRoutes = router.Group("/repositories")
	repoRoutes.GET("", func(c *gin.Context) { controllers.GetRepositories(c) })
	repoRoutes.GET("/:id/show", func(c *gin.Context) { controllers.GetRepository(c) })
	repoRoutes.GET("/:id/edit", func(c *gin.Context) { NullRoute(c) })
	repoRoutes.GET("/:id/delete", func(c *gin.Context) { NullRoute(c) })
	repoRoutes.GET("/new", func(c *gin.Context) { controllers.AddRepository(c) })
	repoRoutes.POST("/preview", func(c *gin.Context) { controllers.PreviewRepository(c) })
	repoRoutes.POST("/create", func(c *gin.Context) { controllers.CreateRepository(c) })

	//Resource Routes
	var resourceRoutes = router.Group("/resources")
	resourceRoutes.GET("", func(c *gin.Context) { controllers.GetResources(c) })
	resourceRoutes.GET("/:id/show", func(c *gin.Context) { controllers.GetResource(c) })
	resourceRoutes.POST("/preview", func(c *gin.Context) { controllers.PreviewResource(c) })
	resourceRoutes.POST("/create", func(c *gin.Context) { controllers.CreateResource(c) })
	resourceRoutes.GET("/:id/edit", func(c *gin.Context) { NullRoute(c) })
	resourceRoutes.GET("/:id/delete", func(c *gin.Context) { controllers.DeleteResource(c) })
	resourceRoutes.GET("/:id/new", func(c *gin.Context) { controllers.AddResource(c) })

	var accessionRoutes = router.Group("/accessions")
	accessionRoutes.GET("", func(c *gin.Context) { controllers.GetAccessions(c) })
	accessionRoutes.POST("/preview", func(c *gin.Context) { controllers.PreviewAccession(c) })
	accessionRoutes.POST("/create", func(c *gin.Context) { controllers.CreateAccession(c) })
	accessionRoutes.GET("/:id/show", func(c *gin.Context) { controllers.GetAccession(c) })
	accessionRoutes.GET("/:id/delete", func(c *gin.Context) { controllers.DeleteAccession(c) })
	accessionRoutes.GET("/:id/new", func(c *gin.Context) { controllers.AddAccession(c) })

	var mediaRoutes = router.Group("/media")
	mediaRoutes.GET("/entries", func(c *gin.Context) { mediacontrollers.GetEntries(c) })
	mediaRoutes.POST("/new", func(c *gin.Context) { mediacontrollers.NewMedia(c) })
	mediaRoutes.GET("/:id/show", func(c *gin.Context) { mediacontrollers.ShowMedia(c) })
	mediaRoutes.GET("/:id/edit", func(c *gin.Context) { mediacontrollers.EditMedia(c) })
	mediaRoutes.GET("/:id/delete", func(c *gin.Context) { mediacontrollers.DeleteMedia(c) })

	//Optical Disc Routes
	mediaRoutes.POST("/create/optical", func(c *gin.Context) { mediacontrollers.CreateOpticalDisc(c) })
	mediaRoutes.POST("/:id/update/optical", func(c *gin.Context) { mediacontrollers.UpdateOpticalDisc(c) })

	//Hard Disk Drive routes
	mediaRoutes.POST("/create/hard-disk-drive", func(c *gin.Context) { mediacontrollers.CreateHardDiskDrive(c) })

	//User Routes
	var userRoutes = router.Group("/users")
	userRoutes.GET("", func(c *gin.Context) { controllers.GetUsers(c) })
	userRoutes.GET("/new", func(c *gin.Context) { controllers.NewUser(c) })
	userRoutes.POST("/create", func(c *gin.Context) { controllers.CreateUser(c) })
	userRoutes.GET("/login", func(c *gin.Context) { controllers.UserLogin(c) })
	userRoutes.GET("/logout", func(c *gin.Context) { controllers.UserLogout(c) })
	userRoutes.POST("/authenticate", func(c *gin.Context) { controllers.UserAuthenticate(c) })

	//Session Routes
	var sessionRoutes = router.Group("/sessions")
	sessionRoutes.GET("/create", func(c *gin.Context) { controllers.CreateSession(c) })
	sessionRoutes.DELETE("/:id/destroy", func(c *gin.Context) { controllers.CreateSession(c) })
	sessionRoutes.GET("", func(c *gin.Context) { controllers.CreateSession(c) })

	//Search Routes
	var searchRoutes = router.Group("/search")
	searchRoutes.GET("", func(c *gin.Context) { controllers.NewSearch(c) })
	searchRoutes.POST("", func(c *gin.Context) {
		getIdentifiers()
		controllers.SearchIndex(c)
	})
}

func NullRoute(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, "Unsupported Route")
}

func getIdentifiers() {
	accessionIDs = *database.GetAccessionIdentifiers()
	resourceIDs = *database.GetResourceIdentifiers()
}
