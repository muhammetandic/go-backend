package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/muhammetandic/go-backend/main/api/controllers"
	"github.com/muhammetandic/go-backend/main/utils/middleware"
)

func Routes(router *gin.Engine) {
	auth := router.Group("auth")
	auth.POST("login", controllers.Login)
	auth.POST("register", controllers.Register)
	auth.POST("refresh-token", controllers.RefreshToken)

	api := router.Group("api")

	todos := api.Group("todos", middleware.Authorize)
	todos.GET("", controllers.GetAllTodos)
	todos.GET("/:id", controllers.GetTodo)
	todos.POST("", controllers.PostTodo)

	user := api.Group("users", middleware.Authorize)
	user.POST("", controllers.PostUser)
	user.GET("", controllers.GetAllUsers)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
