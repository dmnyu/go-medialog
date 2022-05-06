package main

import (
	"flag"
	"fmt"
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/routes"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	migrate        bool
	goAspaceConfig string
	logFileLoc     string
)

func init() {
	flag.BoolVar(&migrate, "migrate", false, "")
	flag.StringVar(&goAspaceConfig, "config", "", "")
	flag.StringVar(&logFileLoc, "log-file", "gomedialog.log", "")
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

func getRepoCode(i int) string {
	switch i {
	case 2:
		return "tamwag"
	case 3:
		return "fales"
	case 6:
		return "archives"
	case 100:
		return "abudhabi"
	}
	return "unkown"
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
	log.Printf("[INFO] logging to %s", logFileLoc)

	if migrate == true {
		database.MigrateDatabase()
		os.Exit(0)
	}
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
		"getRepoCode":  getRepoCode,
		"add":          add,
		"subtract":     subtract,
	})

	router.LoadHTMLGlob("templates/**/*.html")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.SetTrustedProxies([]string{"127.0.0.1"})
	//Load Application Routes
	routes.LoadRoutes(router)

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

func add(a int, b int) int      { return a + b }
func subtract(a int, b int) int { return a - b }
