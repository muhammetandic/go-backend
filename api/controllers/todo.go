package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/muhammetandic/go-backend/main/services"
	"github.com/muhammetandic/go-backend/main/utils/helpers"
)

func GetAllTodos(c *gin.Context) {
	data, err := services.GetAllTodos()
	if err != nil {
		errResponse := helpers.StatusInternalServerError(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
	}
	c.JSON(http.StatusOK, data)
}

func GetTodo(c *gin.Context) {
	// c.JSON(http.StatusOK, todo)
}

func PostTodo(c *gin.Context) {
	// c.JSON(http.StatusCreated, todo)
}

func Put(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
