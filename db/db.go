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

func Connect(connectionStrings ...string) {
	if len(connectionStrings) > 0 {
		Instance, dbError = gorm.Open(sqlite.Open(connectionStrings[0]), &gorm.Config{})
	} else {
		Instance, dbError = gorm.Open(sqlite.Open("./antpos.db"), &gorm.Config{})
	}

	if dbError != nil {
		log.Fatal(dbError)
		panic("cannot connect to DB")
	}

	log.Println("connected to DB")
}

func Migrate() {
	err := Instance.AutoMigrate(&model.Todo{}, &model.User{}, &model.Role{}, &model.UserToRole{}, &model.Privilege{})
	if err != nil {
		log.Println("database migration failed:", err.Error())
		return
	}
	log.Println("database migration completed")
}

func SeedDatabase() {
	user := model.User{Email: "admin@antpos.dev", FullName: "admin", Password: "admin"}
	if err := user.HashPassword(user.Password); err != nil {
		log.Println(err.Error())
	}
	if err := Instance.First(&model.User{}, "email = ?", user.Email).Error; err != nil {
		Instance.Create(&user)
	}

	role := model.Role{Name: "admin"}
	if err := Instance.First(&model.Role{}, "name = ?", role.Name).Error; err != nil {
		Instance.Create(&role)
	}

	userToRole := model.UserToRole{UserID: 1, RoleID: 1}
	if err := Instance.First(&model.UserToRole{}, "user_id = ? AND role_id = ?", userToRole.UserID, userToRole.RoleID).Error; err != nil {
		Instance.Create(&userToRole)
	}

	privileges := []*model.Privilege{
		{RoleID: 1, Endpoint: "users", CanRead: true, CanInsert: true, CanUpdate: true, CanDelete: true},
		{RoleID: 1, Endpoint: "roles", CanRead: true, CanInsert: true, CanUpdate: true, CanDelete: true},
		{RoleID: 1, Endpoint: "todos", CanRead: true, CanInsert: true, CanUpdate: true, CanDelete: true},
		{RoleID: 1, Endpoint: "privileges", CanRead: true, CanInsert: true, CanUpdate: true, CanDelete: true},
	}
	for _, privilege := range privileges {
		if err := Instance.First(&model.Privilege{}, "role_id = ? AND endpoint = ?", privilege.RoleID, privilege.Endpoint).Error; err != nil {
			Instance.Create(privilege)
		}
	}
}
