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

	todos := api.Group("todos", middleware.Authenticate, middleware.Authorize)
	todos.POST("", controllers.CreateTodo)
	todos.GET("", controllers.GetAllTodos)
	todos.GET("/:id", controllers.GetTodo)
	todos.PUT("/:id", controllers.UpdateTodo)
	todos.DELETE("/:id", controllers.DeleteTodo)

	user := api.Group("users", middleware.Authenticate, middleware.Authorize)
	user.POST("", controllers.CreateUser)
	user.GET("", controllers.GetAllUsers)
	user.GET("/:id", controllers.GetUser)
	user.PUT("/:id", controllers.UpdateUser)
	user.DELETE("/:id", controllers.DeleteUser)
	user.POST("/add-role", controllers.AddRole)
	user.POST("/remove-role", controllers.RemoveRole)

	role := api.Group("roles", middleware.Authenticate, middleware.Authorize)
	role.POST("", controllers.CreateRole)
	role.GET("", controllers.GetAllRoles)
	role.GET("/:id", controllers.GetRole)
	role.PUT("/:id", controllers.UpdateRole)
	role.DELETE("/:id", controllers.DeleteRole)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
