package main

import (
	"flag"
	"fmt"
	"github.com/dmnyu/go-medialog/controllers"
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/index"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	migrate        bool
	reindex        bool
	goAspaceConfig string
	logFileLoc     string
	accessionIDs   map[int]string
	resourceIDs    map[int]string
)

func init() {
	flag.BoolVar(&migrate, "migrate", false, "")
	flag.BoolVar(&reindex, "reindex", false, "")
	flag.StringVar(&goAspaceConfig, "config", "", "")
	flag.StringVar(&logFileLoc, "log-file", "gomedialog.log", "")
}

var router = gin.Default()

func main() {
	flag.Parse()
	logfile, err := os.OpenFile(logFileLoc, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.Println("[INFO] [APP] starting go-medialog ☮ ☮")
	log.Printf("[INFO] [APP] logging to %s", logFileLoc)

	if migrate == true {
		database.MigrateDatabase()
		log.Printf("[INFO] [APP] shutting down medialog")
		os.Exit(0)
	}

	if reindex == true {
		database.ConnectDatabase()
		index.Reindex()
		os.Exit(0)
	}

	router.SetFuncMap(template.FuncMap{
		"formatAsDate":           formatAsDate,
		"getRepoName":            getRepoName,
		"add":                    add,
		"subtract":               subtract,
		"getMediaType":           getMediaType,
		"getAccessionIdentifier": getAccessionIdentifier,
		"getResourceIdentifier":  getResourceIdentifier,
	})

	router.LoadHTMLGlob("templates/**/*.html")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.SetTrustedProxies([]string{"127.0.0.1"})
	//Load Application Routes
	loadRoutes(router)

	//Index
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "go-medialog",
		})
	})

	//Start the router
	database.ConnectDatabase()

	if err := router.Run(); err != nil {
		panic(err)
	}

}

//global functions

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

func getRepoName(i int) string {
	switch i {
	case 1:
		return "tamwag"
	case 2:
		return "fales"
	case 3:
		return "nyu archives"
	case 100:
		return "abudhabi"
	}
	return "unknown"
}

func add(a int, b int) int { return a + b }

func subtract(a int, b int) int { return a - b }

func getMediaType(id models.MediaModel) string {
	return models.MediaNames[id]
}

func getAccessionIdentifier(accessionID int) string { return accessionIDs[accessionID] }

func getResourceIdentifier(resourceID int) string {
	resourceIDs := *database.GetResourceIdentifiers()
	return resourceIDs[resourceID]
}

func getIdentifiers() {
	accessionIDs = *database.GetAccessionIdentifiers()
	resourceIDs = *database.GetResourceIdentifiers()
}

//Routes
func loadRoutes(router *gin.Engine) {
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
	mediaRoutes.GET("/entries", func(c *gin.Context) { controllers.GetEntries(c) })
	mediaRoutes.POST("/new", func(c *gin.Context) { controllers.NewMedia(c) })
	mediaRoutes.GET("/:id/show", func(c *gin.Context) { controllers.ShowMedia(c) })
	mediaRoutes.GET("/:id/delete", func(c *gin.Context) { controllers.DeleteMedia(c) })
	mediaRoutes.POST("/create/optical", func(c *gin.Context) { controllers.CreateOpticalDisc(c) })
	mediaRoutes.POST("/create/hard-disk-drive", func(c *gin.Context) { controllers.CreateHardDiskDrive(c) })

	//API Routes
	var apiV0Routes = router.Group("/api/v0")
	apiV0Routes.POST("/create-optical", func(c *gin.Context) { controllers.CreateOpticalDiscAPI(c) })
	apiV0Routes.POST("/create-resource", func(c *gin.Context) { controllers.CreateResourceAPI(c) })
	apiV0Routes.POST("/create-accession", func(c *gin.Context) { controllers.CreateAccessionAPI(c) })

	//User Routes
	var userRoutes = router.Group("/users")
	userRoutes.GET("", func(c *gin.Context) { controllers.GetUsers(c) })
	userRoutes.GET("/new", func(c *gin.Context) { controllers.NewUser(c) })
	userRoutes.POST("/create", func(c *gin.Context) { controllers.CreateUser(c) })
	userRoutes.GET("/login", func(c *gin.Context) { controllers.UserLogin(c) })
	userRoutes.POST("/authenticate", func(c *gin.Context) { controllers.UserAuthenticate(c) })

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
