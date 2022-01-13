package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

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

	if err := db.AutoMigrate(&MediaOpticalDisc{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&MediaEntry{}); err != nil {
		panic(err)
	}

	log.Println("[INFO] migrations run")
}
