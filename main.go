package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/services"
)

func main() {
	db, error := db.Connect()
	if error != nil {
		log.Println(error)
	}

	db.DB()

	router := gin.Default()

	groceries := router.Group("grocery")

	groceries.GET("", services.GetAllGroceries)
	groceries.GET("/:id", services.GetGrocery)
	groceries.POST("", services.PostGrocery)
	groceries.PUT("/:id", services.UpdateGrocery)
	groceries.DELETE("/:id", services.DeleteGrocery)

	todos := router.Group("todo")
	todos.GET("", services.GetAllTodos)
	todos.GET("/:id", services.GetTodo)
	todos.POST("", services.PostTodo)
	todos.PUT("/:id", services.UpdateTodo)
	todos.DELETE("/:id", services.DeleteTodo)
	log.Fatal(router.Run(":4444"))
}
