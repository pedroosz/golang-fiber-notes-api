package config

import (
	"log"

	"github.com/pedroosz/golang-fiber-notes-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error on db open.")
	}

	db.AutoMigrate(&models.Note{})

	Database = db

}
