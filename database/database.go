package database

import (
	"github.com/dmnyu/go-medialog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func ConnectDatabase() {
	log.Println("[INFO] connecting to database")
	var err error
	db, err = gorm.Open(sqlite.Open("gomedialog.db"), &gorm.Config{})

	if err != nil {
		panic("[FATAL] Failed to connect to database!")
	}
}

func MigrateDatabase() {
	ConnectDatabase()
	log.Println("[INFO] migrating database")

	if err := db.AutoMigrate(&models.Repository{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.Resource{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.Accession{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.MediaOpticalDisc{}); err != nil {
		panic(err)
	}

	log.Println("[INFO] migrations run")
}
