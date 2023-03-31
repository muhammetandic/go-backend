package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/api"
	"github.com/muhammetandic/go-backend/main/db"
)

func main() {
	db.Connect()
	db.Migrate()

	router := gin.New()

	api.Routes(router)
	log.Fatal(router.Run(":4444"))
}
