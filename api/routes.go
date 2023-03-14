package api

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/api/controllers"
	"github.com/muhammetandic/go-backend/main/services"
)

func Routes(router *gin.Engine) {
	auth := router.Group("auth")
	auth.POST("login", controllers.Login)
	auth.POST("register", controllers.Register)
	
	api := router.Group("api")
	
	todos := api.Group("todos")
	todos.GET("", services.GetAllTodos)
	todos.GET("/:id", services.GetTodo)
	todos.POST("", services.PostTodo)

	user := api.Group("users")
	user.POST("", controllers.PostUser)

	return
}
