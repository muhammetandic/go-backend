package api

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/services"
)

func Routes(router *gin.Engine) {
	todos := router.Group("todos")
	todos.GET("", services.GetAllTodos)
	todos.GET("/:id", services.GetTodo)
	todos.POST("", services.PostTodo)
	return
}
