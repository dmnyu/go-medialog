package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDataBase() {
	dsn := "medialog:medialog@tcp(127.0.0.1:3306)/gomedialog?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db = d
}

func MigrateDatabase() {
	ConnectDataBase()
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
