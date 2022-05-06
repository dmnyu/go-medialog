package database

import (
	"github.com/dmnyu/go-medialog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var databaseLoc = "gomedialog.db"

func ConnectDatabase() {

	log.Printf("[INFO] [DATABASE] connecting to %s", databaseLoc)
	var err error
	db, err = gorm.Open(sqlite.Open(databaseLoc), &gorm.Config{})
	if err != nil {
		log.Fatalf("[FATAL] [DATABASE] Failed to connect to %s", databaseLoc)
	}
	log.Printf("[INFO] [DATABASE] successfully connected to %s", databaseLoc)
}

func MigrateDatabase() {
	ConnectDatabase()
	log.Printf("[INFO] [DATABASE] migrating %s", databaseLoc)

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

	log.Printf("[INFO] [DATABASE] successfully migrated %s", databaseLoc)
}
