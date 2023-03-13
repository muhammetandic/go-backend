package api

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/api/controllers"
	"github.com/muhammetandic/go-backend/main/services"
)

func Routes(router *gin.Engine) {
	todos := router.Group("todos")
	todos.GET("", services.GetAllTodos)
	todos.GET("/:id", services.GetTodo)
	todos.POST("", services.PostTodo)

	auth := router.Group("auth")
	auth.POST("login", controllers.Login)
	auth.POST("register")

	return
}
