package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/db/model"
)

func PostUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
