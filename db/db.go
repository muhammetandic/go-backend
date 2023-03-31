package db

import (
	"log"

	"github.com/muhammetandic/go-backend/main/db/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Instance *gorm.DB
	dbError  error
)

func Connect() {
	Instance, dbError = gorm.Open(sqlite.Open("./antpos.db"), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
		panic("cannot connect to DB")
	}

	log.Println("connected to DB")
}

func Migrate() {
	Instance.AutoMigrate(&model.Grocery{}, &model.Todo{}, &model.User{})
	log.Println("database migration completed")
}
