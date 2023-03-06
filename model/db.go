package model

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbInit() (*gorm.DB, error) {
    db, error := gorm.Open(sqlite.Open("./antpos.db"), &gorm.Config{})
    
    if error != nil {
        log.Fatal(error.Error())
    }
    
    if error = db.AutoMigrate(&Grocery{}); error != nil {
        log.Println(error)
    }
    
    return db, error
}