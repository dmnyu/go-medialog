package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := "medialog:medialog@tcp(127.0.0.1:3306)/gomedialog?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}

func MigrateDatabase() {
	ConnectDataBase()
	DB.AutoMigrate(&Repository{})
	DB.AutoMigrate(&Resource{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Accession{})
	DB.AutoMigrate(&Entry{})
}
