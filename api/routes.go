package api

import (
	"github.com/labstack/echo/v4"
	"github.com/muhammetandic/go-backend/main/api/controllers"
	"github.com/muhammetandic/go-backend/main/utils/middleware"
)

func Routes(router *echo.Echo) {
	auth := router.Group("auth")

	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)
	auth.POST("/refresh-token", controllers.RefreshToken)

	api := router.Group("api")
	api.Use(middleware.Authenticate)
	api.Use(middleware.Authorize)

	todos := api.Group("todos")
	todos.POST("", controllers.CreateTodo)
	todos.GET("", controllers.GetAllTodos)
	todos.GET("/:id", controllers.GetTodo)
	todos.PUT("/:id", controllers.UpdateTodo)
	todos.DELETE("/:id", controllers.DeleteTodo)

	// user := api.Group("users")
	// user.POST("", controllers.CreateUser)
	// user.GET("", controllers.GetAllUsers)
	// user.GET("/:id", controllers.GetUser)
	// user.PUT("/:id", controllers.UpdateUser)
	// user.DELETE("/:id", controllers.DeleteUser)

	// role := api.Group("roles")
	// role.POST("", controllers.CreateRole)
	// role.GET("", controllers.GetAllRoles)
	// role.GET("/:id", controllers.GetRole)
	// role.PUT("/:id", controllers.UpdateRole)
	// role.DELETE("/:id", controllers.DeleteRole)
}
