package api

import (
	"github.com/gin-gonic/gin"

	"github.com/muhammetandic/go-backend/main/api/controllers"
	"github.com/muhammetandic/go-backend/main/utils/middleware"
)

func Routes(router *gin.Engine) {
	auth := router.Group("auth")
	auth.POST("login", controllers.Login)
	auth.POST("register", controllers.Register)

	api := router.Group("api")

	todos := api.Group("todos", middleware.Authorize)
	todos.GET("", controllers.GetAllTodos)
	todos.GET("/:id", controllers.GetTodo)
	todos.POST("", controllers.PostTodo)

	user := api.Group("users")
	user.POST("", controllers.PostUser)
}
