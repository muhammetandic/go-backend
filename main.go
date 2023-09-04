package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muhammetandic/go-backend/main/api"
	"github.com/muhammetandic/go-backend/main/db"
	_ "github.com/muhammetandic/go-backend/main/docs"
)

func main() {
	db.Connect()
	db.Migrate()

	router := echo.New()
	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}
	router.Use(middleware.CORSWithConfig(corsConfig))

	api.Routes(router)
	log.Fatal(router.Start(":4444"))
}
