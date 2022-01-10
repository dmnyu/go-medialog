package database

import (
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
	if err := db.AutoMigrate(&Repository{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Resource{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Accession{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&MediaObject{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&MediaOpticalDisc{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&MediaHardDiskDrive{}); err != nil {
		panic(err)
	}
}
