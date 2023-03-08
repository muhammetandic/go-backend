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
    router.GET("grocery/:id", grocery.GetGrocery)
    router.POST("grocery", grocery.PostGrocery)
    router.PUT("grocery/:id", grocery.UpdateGrocery)
    router.DELETE("grocery/:id", grocery.DeleteGrocery)
    log.Fatal(router.Run(":4444"))
}