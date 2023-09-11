package main

import (
	"flag"
	"fmt"
	"github.com/dmnyu/go-medialog/controllers"
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/index"
	"github.com/dmnyu/go-medialog/models"
	"github.com/dmnyu/go-medialog/routes"
	"github.com/dmnyu/go-medialog/shared"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"html/template"
	"log"
	"os"
	"time"
)

var (
	migrate         bool
	reindex         bool
	config          string
	accessionIDs    map[int]string
	resourceIDs     map[int]string
	logFileLocation string
	router          = gin.Default()
)

func init() {
	flag.BoolVar(&migrate, "migrate", false, "")
	flag.BoolVar(&reindex, "reindex", false, "")
	flag.StringVar(&config, "config", "config/go-medialog.yml", "")
}

type GoMedialogConfig struct {
	Log          string `yaml:"log"`
	Database     string `yaml:"database"`
	AspaceConfig string `yaml:"aspace_config"`
	AspaceEnv    string `yaml:"aspace_env"`
}

func configure() {

	//ensure the config file exists
	if err := shared.FileExists(config); err != nil {
		panic(err)
	}

	configBytes, err := os.ReadFile(config)
	if err != nil {
		panic(err)
	}

	goMedialogConfig := GoMedialogConfig{}

	if err := yaml.Unmarshal(configBytes, &goMedialogConfig); err != nil {
		panic(err)
	}

	logFileLocation = goMedialogConfig.Log

	//set the database
	database.DatabaseLoc = goMedialogConfig.Database

	//setup aspace
	controllers.AspaceConfig = goMedialogConfig.AspaceConfig
	controllers.AspaceEnv = goMedialogConfig.AspaceEnv

}

func main() {
	flag.Parse()
	configure()

	logfile, err := os.OpenFile(logFileLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.Printf("[INFO] [APP] logging to %s", logFileLocation)
	log.Println("[INFO] [APP] starting go-medialog ☮")

	//migrate the database if `migrate` flag is set
	if migrate == true {
		if err := database.MigrateDatabase(); err != nil {
			log.Printf("[FATAL] [DATABASE] database migration failed")
			os.Exit(2)
		}
		log.Printf("[INFO] [APP] shutting down medialog")
		os.Exit(0)
	}

	//reindex if `reindex` flag is set
	if reindex == true {

		if err := database.ConnectDatabase(); err != nil {
			log.Printf("[FATAL] [DATABASE] database connection failed")
			os.Exit(1)
		}

		//delete the index entries
		if err := index.DeleteAll(); err != nil {
			log.Printf("[FATAL] [INDEX] shutting down medialog")
			os.Exit(3)
		}

		if err := index.Reindex(); err != nil {
			log.Printf("[FATAL] [INDEX] shutting down medialog")
			os.Exit(3)
		}

		log.Printf("[INFO] [MEDIALOG] shutting down medialog")
		os.Exit(0)

	}

	//connect to the database
	if err := database.ConnectDatabase(); err != nil {
		log.Printf("[FATAL] [DATABASE] database connection failed")
		os.Exit(1)
	}
	log.Printf("[INFO] [DATABASE] connected to database")

	//test archivesspace connection
	if err := controllers.GetClient(); err != nil {
		log.Printf("[FATAL] [ASPACE] archivesspace connection failed")
		os.Exit(4)
	}
	log.Printf("[INFO] [ASPACE] connected to archivesspace instance")

	//test index connection

	//load functions
	router.SetFuncMap(template.FuncMap{
		"formatAsDate":           formatAsDate,
		"getRepoName":            getRepoName,
		"add":                    add,
		"subtract":               subtract,
		"getMediaType":           getMediaType,
		"getAccessionIdentifier": getAccessionIdentifier,
		"getResourceIdentifier":  getResourceIdentifier,
		"isEqual":                isEqual,
	})

	//configure router
	router.LoadHTMLGlob("templates/**/*.html")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.SetTrustedProxies([]string{"127.0.0.1"})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//Load Application Routes
	routes.LoadRoutes(router)

	//run the application
	if err := router.Run(); err != nil {
		panic(err)
	}

}

// global functions
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

// turn this into a function that uses the db
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

func getResourceIdentifier(resourceID int) string { return resourceIDs[resourceID] }

func isEqual(a string, b string) bool { return a == b }

// Routes
