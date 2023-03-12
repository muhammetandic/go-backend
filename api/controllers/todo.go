package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, todo)
}

func Post(c *gin.Context) {
	c.JSON(http.StatusCreated, todo)
}

func Put(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
