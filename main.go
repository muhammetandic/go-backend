package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/model"
	"github.com/muhammetandic/go-backend/main/grocery"
)

func main() {
    db, error := model.DbInit()
    if error != nil {
        log.Println(error)
    }
    
    db.DB()
    
    router := gin.Default()
    
    router.GET("grocery", grocery.GetAllGroceries)
    log.Fatal(router.Run(":4444"))
}