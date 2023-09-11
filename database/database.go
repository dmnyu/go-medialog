package database

import (
	"github.com/dmnyu/go-medialog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var DatabaseLoc = "gomedialog.db"

func ConnectDatabase() error {

	var err error
	db, err = gorm.Open(sqlite.Open(DatabaseLoc), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func MigrateDatabase() error {
	ConnectDatabase()
	log.Printf("[INFO] [DATABASE] migrating %s", DatabaseLoc)

	//core models
	if err := db.AutoMigrate(&models.Repository{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Resource{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Accession{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	//media models
	if err := db.AutoMigrate(&models.MediaOpticalDisc{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.MediaHardDrive{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Session{}); err != nil {
		return err
	}

	return nil

}
