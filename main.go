package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/muhammetandic/go-backend/main/api"
	"github.com/muhammetandic/go-backend/main/db"
	_ "github.com/muhammetandic/go-backend/main/docs"
)

// @title AntPos API
// @version 1.0
// @description AntPos Basic Pos App API Project

// @contact.name   Muhammet Andiç
// @contact.email  muhammetandic@gmail.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @BasePath /
func main() {
	db.Connect()
	db.Migrate()

	router := gin.New()

	api.Routes(router)
	log.Fatal(router.Run(":4444"))
}
