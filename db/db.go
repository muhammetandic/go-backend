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

func Create() {
	user := model.User{Email: "admin@antpos.dev", FullName: "admin", Password: "admin"}
	if err := user.HashPassword(user.Password); err != nil {
		log.Println(err.Error())
	}
	Instance.Create(&user)

	role := model.Role{Name: "admin"}
	Instance.Create(&role)

	userToRole := model.UserToRole{UserID: 1, RoleID: 1}
	Instance.Create(&userToRole)

	privileges := []*model.Privilege{
		{RoleID: 1, Endpoint: "users", CanRead: true, CanInsert: true, CanUpdate: true, CanDelete: true},
		{RoleID: 1, Endpoint: "roles", CanRead: true, CanInsert: true, CanUpdate: true, CanDelete: true},
		{RoleID: 1, Endpoint: "todos", CanRead: true, CanInsert: true, CanUpdate: true, CanDelete: true},
	}
	Instance.Create(privileges)
}
