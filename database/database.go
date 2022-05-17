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

	//core models
	if err := db.AutoMigrate(&models.Repository{}); err != nil {
		log.Fatalf("[FATAL] [DATABASE] %s", err.Error())
	}

	if err := db.AutoMigrate(&models.Resource{}); err != nil {
		log.Fatalf("[FATAL] [DATABASE] %s", err.Error())
	}

	if err := db.AutoMigrate(&models.Accession{}); err != nil {
		log.Fatalf("[FATAL] [DATABASE] %s", err.Error())
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("[FATAL] [DATABASE] %s", err.Error())
	}

	//media models
	if err := db.AutoMigrate(&models.MediaOpticalDisc{}); err != nil {
		log.Fatalf("[FATAL] [DATABASE] %s", err.Error())
	}

	if err := db.AutoMigrate(&models.MediaHardDrive{}); err != nil {
		log.Fatalf("[FATAL] [DATABASE] %s", err.Error())
	}

	log.Printf("[INFO] [DATABASE] migration of %s complete", databaseLoc)

}
