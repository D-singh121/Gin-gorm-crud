package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// var err error

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		log.Fatal("Connection failed :")
	}
	err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	DB = database

}
