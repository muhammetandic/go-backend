package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/api"
	"github.com/muhammetandic/go-backend/main/db"
)

func main() {
	db, error := db.Connect()
	if error != nil {
		log.Println(error)
	}

	db.DB()

	router := gin.Default()

	api.Routes(router)
	log.Fatal(router.Run(":4444"))
}
