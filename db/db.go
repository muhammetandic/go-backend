package db

import (
	"log"

	"github.com/muhammetandic/go-backend/main/db/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, error := gorm.Open(sqlite.Open("./antpos.db"), &gorm.Config{})

	if error != nil {
		log.Fatal(error.Error())
	}

	if error = db.AutoMigrate(&model.Grocery{}, &model.Todo{}, &model.User{}); error != nil {
		log.Println(error)
	}

	return db, error
}
