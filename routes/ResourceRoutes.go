package routes

import (
	"github.com/dmnyu/go-medialog/controllers"
	"github.com/dmnyu/go-medialog/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

func LoadResourceRoutes(router *gin.Engine) {
	var resourceRoutes = router.Group("/resources")

	//show all resources
	resourceRoutes.GET("", func(c *gin.Context) {
		resources := controllers.GetResources()
		repositories := controllers.GetRepositories()

		c.HTML(http.StatusOK, "resources-index.html", gin.H{
			"title":        "go-medialog -- resources",
			"repositories": repositories,
			"resources":    resources,
		})
	})

	//show a resource
	resourceRoutes.GET("/:id", func(c *gin.Context) {
		resource, err := controllers.GetResource(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		repository, err := controllers.GetRepositoryByID(resource.RepositoryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.HTML(http.StatusOK, "resources-show.html", gin.H{
			"title":      "go-medialog -- resource",
			"repository": repository,
			"resource":   resource,
		})
	})

	//preview a new resource
	resourceRoutes.POST("/preview", func(c *gin.Context) {
		var input = db.CreateAspaceObject{}
		if err := c.Bind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		repository, err := controllers.GetRepositoryByID(input.RepositoryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		aspaceResource, err := controllers.FindAspaceResource(repository.AspaceID, input.ObjectID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		identifiers := aspaceResource.MergeIDs()

		c.HTML(http.StatusOK, "resources-preview.html", gin.H{
			"input":       input,
			"identifiers": identifiers,
			"repository":  repository,
			"resource":    aspaceResource,
		})
	})

	//create a new resource
	resourceRoutes.POST("/create", func(c *gin.Context) {
		var input = db.CreateAspaceObject{}
		if err := c.Bind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		repository, err := controllers.GetRepositoryByID(input.RepositoryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		aspaceResource, err := controllers.FindAspaceResource(repository.AspaceID, input.ObjectID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		uriSplit := strings.Split(aspaceResource.URI, "/")
		aspaceResourceID, err := strconv.Atoi(uriSplit[len(uriSplit)-1])
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		resource := db.Resource{
			Model:                     gorm.Model{},
			AspaceResourceID:          aspaceResourceID,
			RepositoryID:              input.RepositoryID,
			AspaceResourceTitle:       aspaceResource.Title,
			AspaceResourceIdentifiers: aspaceResource.MergeIDs(),
		}

		if err = controllers.CreateResource(resource); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.Redirect(http.StatusFound, "/resources")
	})

}
