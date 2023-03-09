package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/grocery"
	"github.com/muhammetandic/go-backend/main/model"
	"github.com/muhammetandic/go-backend/main/todo"
)

func main() {
    db, error := model.DbInit()
    if error != nil {
        log.Println(error)
    }
    
    db.DB()
    
    router := gin.Default()
    
    groceries := router.Group("grocery")
    
    groceries.GET("", grocery.GetAllGroceries)
    groceries.GET("/:id", grocery.GetGrocery)
    groceries.POST("", grocery.PostGrocery)
    groceries.PUT("/:id", grocery.UpdateGrocery)
    groceries.DELETE("/:id", grocery.DeleteGrocery)
    
    todos := router.Group("todo")
    todos.GET("", todo.GetAllTodos)
    todos.GET("/:id", todo.GetTodo)
    todos.POST("", todo.PostTodo)
    todos.PUT("/:id", todo.UpdateTodo)
    todos.DELETE("/:id", todo.DeleteTodo)
    log.Fatal(router.Run(":4444"))
}