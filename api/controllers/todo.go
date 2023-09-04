package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"

	"github.com/muhammetandic/go-backend/main/db/model"
	"github.com/muhammetandic/go-backend/main/services"
)

func CreateTodo(c echo.Context) error {
	var todoData model.Todo
	if err := c.Bind(&todoData); err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	todoService := services.NewTodoService()
	todo, err := todoService.Add(&todoData, context.Background())
	if err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, todo)
}

func GetAllTodos(c echo.Context) error {
	todoService := services.NewTodoService()
	entities, err := todoService.GetAll(context.Background())
	if err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, entities)
}

func GetTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	todoService := services.NewTodoService()
	entity, err := todoService.Get(id, context.Background())
	if err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, entity)
}

func UpdateTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var todo model.Todo
	err = c.Bind(&todo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	todoService := services.NewTodoService()
	err = todoService.Update(id, &todo, context.Background())
	if err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return c.JSON(http.StatusNoContent, gin.H{})
}

func DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	todoService := services.NewTodoService()
	err = todoService.Delete(id, context.Background())
	if err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return c.JSON(http.StatusNoContent, gin.H{})
}
