package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/muhammetandic/go-backend/main/db/model"
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
	err := Instance.AutoMigrate(&model.Todo{}, &model.User{}, &model.Role{}, &model.UserToRole{}, &model.Privilege{})
	if err != nil {
		log.Println("database migration is failed")
	}
	log.Println("database migration completed")
}
